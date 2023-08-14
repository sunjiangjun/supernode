package chain

import (
	"fmt"

	"github.com/0xcregis/easynode/store"
	"github.com/0xcregis/easynode/store/chain/ether"
	"github.com/0xcregis/easynode/store/chain/filecoin"
	"github.com/0xcregis/easynode/store/chain/polygonpos"
	"github.com/0xcregis/easynode/store/chain/tron"
	"github.com/segmentio/kafka-go"
)

func GetReceiptFromKafka(value []byte, blockChain int64) (*store.Receipt, error) {
	if blockChain == 200 {
		return ether.GetReceiptFromKafka(value)
	} else if blockChain == 205 {
		return tron.GetReceiptFromKafka(value)
	} else if blockChain == 201 {
		return polygonpos.GetReceiptFromKafka(value)
	} else if blockChain == 301 {
		return filecoin.GetReceiptFromKafka(value)
	} else {
		return nil, fmt.Errorf("blockchain:%v does not support", blockChain)
	}
}

func GetBlockFromKafka(value []byte, blockChain int64) (*store.Block, error) {
	if blockChain == 200 {
		return ether.GetBlockFromKafka(value)
	} else if blockChain == 205 {
		return tron.GetBlockFromKafka(value)
	} else if blockChain == 201 {
		return polygonpos.GetBlockFromKafka(value)
	} else if blockChain == 301 {
		return filecoin.GetBlockFromKafka(value)
	} else {
		return nil, fmt.Errorf("blockchain:%v does not support", blockChain)
	}
}

func GetTxFromKafka(value []byte, blockChain int64) (*store.Tx, error) {
	if blockChain == 200 {
		return ether.GetTxFromKafka(value)
	} else if blockChain == 205 {
		return tron.GetTxFromKafka(value)
	} else if blockChain == 201 {
		return polygonpos.GetTxFromKafka(value)
	} else if blockChain == 301 {
		return filecoin.GetTxFromKafka(value)
	} else {
		return nil, fmt.Errorf("blockchain:%v does not support", blockChain)
	}
}

func ParseTx(blockchain int64, msg *kafka.Message) (*store.SubTx, error) {
	if blockchain == 200 {
		return ether.ParseTx(msg.Value, store.EthTopic, blockchain)
	}
	if blockchain == 205 {
		return tron.ParseTx(msg.Value, store.TronTopic, blockchain)
	}
	if blockchain == 201 {
		return polygonpos.ParseTx(msg.Value, store.PolygonTopic, blockchain)
	}
	if blockchain == 301 {
		return filecoin.ParseTx(msg.Value, store.PolygonTopic, blockchain)
	}
	return nil, nil
}

func GetTxType(blockchain int64, msg *kafka.Message) (uint64, error) {
	if blockchain == 200 {
		return ether.GetTxType(msg.Value)
	}
	if blockchain == 205 {
		return tron.GetTxType(msg.Value)
	}
	if blockchain == 201 {
		return polygonpos.GetTxType(msg.Value)
	}
	if blockchain == 301 {
		return filecoin.GetTxType(msg.Value)
	}
	return 0, nil
}

func GetCoreAddress(blockChain int64, address string) string {
	if blockChain == 200 {
		return ether.GetCoreAddr(address)
	} else if blockChain == 205 {
		return tron.GetCoreAddr(address)
	} else if blockChain == 201 {
		return polygonpos.GetCoreAddr(address)
	} else if blockChain == 301 {
		return filecoin.GetCoreAddr(address)
	} else {
		return address
	}
}

func CheckAddress(blockChain int64, msg *kafka.Message, list map[string]*store.MonitorAddress) bool {
	if len(list) < 1 {
		return false
	}
	if blockChain == 200 {
		return ether.CheckAddress(msg.Value, list, store.EthTopic)
	} else if blockChain == 205 {
		return tron.CheckAddress(msg.Value, list, store.TronTopic)
	} else if blockChain == 201 {
		return polygonpos.CheckAddress(msg.Value, list, store.PolygonTopic)
	} else if blockChain == 301 {
		return filecoin.CheckAddress(msg.Value, list, store.PolygonTopic)
	} else {
		return false
	}
}