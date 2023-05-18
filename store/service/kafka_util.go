package service

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/tidwall/gjson"
	"github.com/uduncloud/easynode/common/util"
	"math/big"
)

func ParseTx(blockchain int64, msg *kafka.Message) interface{} {
	if blockchain == 200 {
		return ParseTxForEther(msg)
	}
	if blockchain == 205 {
		return ParseTxForTron(msg)
	}
	return nil
}

func ParseTxForEther(msg *kafka.Message) interface{} {
	r := make(map[string]interface{}, 10)
	root := gjson.ParseBytes(msg.Value)
	blockHash := root.Get("blockHash").String()
	r["blockHash"] = blockHash
	blockNumber := root.Get("blockNumber").String()
	r["blockNumber"] = blockNumber
	hash := root.Get("hash").String()
	r["hash"] = hash
	from := root.Get("from").String()
	r["from"] = from
	to := root.Get("to").String()
	r["to"] = to
	input := root.Get("input").String()
	r["input"] = input
	value := root.Get("value").String()
	if len(input) > 5 {
		//合约调用
		r["txType"] = 1
		r["value"] = value
	} else {
		//普通资产转移
		r["txType"] = 2
		v, _ := util.HexToInt(value)
		r["value"] = v
	}

	txTime := root.Get("txTime").String()
	r["txTime"] = txTime

	gasPrice := root.Get("gasPrice").String() //单位：wei
	price, _ := util.HexToInt2(gasPrice)
	bigPrice := big.NewInt(price)

	receipt := root.Get("receipt").String()

	receiptRoot := gjson.Parse(receipt)

	gasUsed := receiptRoot.Get("gasUsed").String()
	gas, _ := util.HexToInt2(gasUsed)
	bigGas := big.NewInt(gas)

	fee := bigPrice.Mul(bigPrice, bigGas)

	r["fee"] = fmt.Sprintf("%v", fee)
	r["gasFee"] = map[string]string{"gasPrice": fmt.Sprintf("%v", price), "gasUsed": fmt.Sprintf("%v", gas)}

	status := receiptRoot.Get("status").String() //0x0:失败，0x1:成功
	if status == "0x0" {
		r["status"] = 0
	} else if status == "0x1" {
		r["status"] = 1
	}

	logs := receiptRoot.Get("logs").Array()
	contractTx := make([]interface{}, 0, 5)
	for _, v := range logs {
		contract := v.Get("address").String()
		value := v.Get("data").String()
		tps := v.Get("topics").Array()
		var from, to string
		if len(tps) >= 3 && tps[0].String() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {
			//method = tps[0].String()
			from = tps[1].String()
			to = tps[2].String()
			m := make(map[string]interface{}, 5)
			m["contract"] = contract
			m["value"], _ = util.HexToInt(value)
			m["from"], _ = util.Hex2Address(from)
			m["to"], _ = util.Hex2Address(to)
			m["method"] = "Transfer"
			contractTx = append(contractTx, m)
		}

	}
	r["contractTx"] = contractTx
	return r
}

func ParseTxForTron(msg *kafka.Message) interface{} {

	txBody := gjson.ParseBytes(msg.Value).Get("tx").String()
	if len(txBody) < 5 {
		return nil
	}
	r := make(map[string]interface{}, 10)
	txRoot := gjson.Parse(txBody)
	status := txRoot.Get("ret.0.contractRet").String()
	if status == "SUCCESS" { //交易成功
		r["status"] = 1
	} else {
		//交易失败
		r["status"] = 0
	}

	hash := txRoot.Get("txID").String()
	r["hash"] = hash
	blockHash := txRoot.Get("raw_data.ref_block_hash").String()
	r["blockHash"] = blockHash
	txTime := txRoot.Get("raw_data.timestamp").Uint()
	r["txTime"] = txTime
	txType := txRoot.Get("raw_data.contract.0.type").String()
	if txType == "TransferContract" {
		r["txType"] = 2
	} else if txType == "TriggerSmartContract" {
		r["txType"] = 1
	}
	v := txRoot.Get("raw_data.contract.0.parameter.value")
	from := v.Get("owner_address").String()
	r["from"] = from
	var to string
	if v.Get("receiver_address").Exists() {
		to = v.Get("receiver_address").String()
	}

	if v.Get("to_address").Exists() {
		to = v.Get("to_address").String()
	}

	if v.Get("contract_address").Exists() {
		to = v.Get("contract_address").String()
	}
	r["to"] = to

	var input string
	if v.Get("data").Exists() {
		input = v.Get("data").String()
	}
	r["input"] = input

	if txType == "TransferContract" {
		r["value"] = v.Get("amount").String()
	} else {
		r["value"] = v.String()
	}

	receiptBody := gjson.ParseBytes(msg.Value).Get("receipt").String()
	if len(receiptBody) > 5 {
		receiptRoot := gjson.Parse(receiptBody)
		fee := receiptRoot.Get("fee").String()
		r["fee"] = fee
		gasFee := receiptRoot.Get("receipt").Map()
		delete(gasFee, "result")
		r["gasFee"] = gasFee
		number := receiptRoot.Get("blockNumber").Uint()
		r["blockNumber"] = number

		logs := receiptRoot.Get("log").Array()
		contractTx := make([]interface{}, 0, 5)
		for _, v := range logs {
			contract := v.Get("address").String()
			value := v.Get("data").String()
			tps := v.Get("topics").Array()
			var from, to string
			if len(tps) >= 3 && tps[0].String() == "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {
				//method = tps[0].String()
				from = tps[1].String()
				to = tps[2].String()
				m := make(map[string]interface{}, 5)
				m["contract"] = contract
				m["value"], _ = util.HexToInt(value)
				m["from"], _ = util.Hex2Address(from)
				m["to"], _ = util.Hex2Address(to)
				m["method"] = "Transfer"
				contractTx = append(contractTx, m)
			}

		}
		r["contractTx"] = contractTx
	}

	return r
}