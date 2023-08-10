package chain

import (
	"fmt"
	"time"

	"github.com/0xcregis/easynode/collect"
	"github.com/0xcregis/easynode/collect/config"
	"github.com/0xcregis/easynode/collect/service/cmd/chain/ether"
	"github.com/0xcregis/easynode/collect/service/cmd/chain/filecoin"
	"github.com/0xcregis/easynode/collect/service/cmd/chain/polygonpos"
	"github.com/0xcregis/easynode/collect/service/cmd/chain/tron2"
	"github.com/sunjiangjun/xlog"
	"github.com/tidwall/gjson"
)

func GetBlockchain(blockchain int, c *config.Chain, store collect.StoreTaskInterface, logConfig *config.LogConfig, nodeId string) collect.BlockChainInterface {
	x := xlog.NewXLogger().BuildOutType(xlog.FILE).BuildFormatter(xlog.FORMAT_JSON).BuildFile(fmt.Sprintf("%v/chain_info", logConfig.Path), 24*time.Hour)
	var srv collect.BlockChainInterface
	if blockchain == 200 {
		srv = ether.NewService(c, x, store, nodeId, collect.EthTopic)
	} else if blockchain == 205 {
		srv = tron2.NewService(c, x, store, nodeId, collect.TronTopic)
	} else if blockchain == 201 {
		srv = polygonpos.NewService(c, x, store, nodeId, collect.PolygonTopic)
	} else if blockchain == 301 {
		srv = filecoin.NewService(c, x, store, nodeId, collect.PolygonTopic)
	}

	return srv
}

func GetTxHashFromKafka(blockchain int, msg []byte) string {
	r := gjson.ParseBytes(msg)
	var txHash string
	if blockchain == 200 {
		txHash = r.Get("hash").String()
	} else if blockchain == 205 {
		tx := r.Get("tx").String()
		txHash = gjson.Parse(tx).Get("txID").String()
	} else if blockchain == 201 {
		txHash = r.Get("hash").String()
	} else if blockchain == 301 {
		txHash = r.Get("hash").String()
	}

	return txHash
}

func GetBlockHashFromKafka(blockchain int, msg []byte) string {
	r := gjson.ParseBytes(msg)
	var blockHash string
	if blockchain == 200 {
		blockHash = r.Get("hash").String()
	} else if blockchain == 205 {
		blockHash = r.Get("blockID").String()
	} else if blockchain == 201 {
		blockHash = r.Get("hash").String()
	} else if blockchain == 301 {
		blockHash = r.Get("blockHash").String()
	}
	return blockHash
}

func GetReceiptHashFromKafka(blockchain int, msg []byte) string {
	r := gjson.ParseBytes(msg)
	var txHash string
	if blockchain == 200 {
		txHash = r.Get("transactionHash").String()
	} else if blockchain == 205 {
		txHash = r.Get("id").String()
	} else if blockchain == 201 {
		txHash = r.Get("transactionHash").String()
	} else if blockchain == 301 {
		txHash = r.Get("transactionHash").String()
	}
	return txHash
}
