package btc

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sunjiangjun/supernode/blockchain"
	chainConfig "github.com/sunjiangjun/supernode/blockchain/config"
	"github.com/sunjiangjun/supernode/blockchain/service"
	"github.com/sunjiangjun/supernode/task"
	"github.com/sunjiangjun/supernode/task/config"
	"github.com/sunjiangjun/xlog"
	"github.com/tidwall/gjson"
)

type Btc struct {
	log        *xlog.XLog
	api        blockchain.API
	blockChain int64
}

func (e *Btc) CreateNodeTask(nodeId string, blockChain int64, number string) (*task.NodeTask, error) {
	t := &task.NodeTask{
		NodeId:      nodeId,
		BlockNumber: number,
		BlockChain:  blockChain,
		TaskType:    2,
		TaskStatus:  0,
		CreateTime:  time.Now(),
		LogTime:     time.Now(),
		Id:          time.Now().UnixNano(),
	}
	return t, nil
}

func NewBtc(log *xlog.XLog, v *config.BlockConfig) *Btc {
	clusters := make([]*chainConfig.NodeCluster, 0, 2)
	for _, v := range v.Cluster {
		c := &chainConfig.NodeCluster{NodeUrl: v.NodeHost, NodeToken: v.NodeKey, Weight: v.Weight}
		clusters = append(clusters, c)
	}
	api := service.NewApi(v.BlockChainCode, clusters, log)
	return &Btc{
		blockChain: v.BlockChainCode,
		log:        log,
		api:        api,
	}
}

func (e *Btc) GetLatestBlockNumber() (int64, error) {
	log := e.log.WithFields(logrus.Fields{
		"id":         time.Now().UnixMilli(),
		"model":      "GetLatestBlockNumber",
		"blockChain": e.blockChain,
	})

	/**
	  {\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x1019b4c\"}
	*/
	var lastNumber int64
	jsonResult, err := e.api.LatestBlock(e.blockChain)
	if err != nil {
		log.Errorf("Eth_GetBlockNumber|err=%v", err)
		return 0, err
	} else {
		log.Printf("Eth_GetBlockNumber|resp=%v", jsonResult)
	}

	lastNumber = gjson.Parse(jsonResult).Get("result").Int()

	if lastNumber > 1 {
		//_ = s.UpdateLastNumber(v.BlockChainCode, lastNumber)
		return lastNumber, nil
	} else {
		return 0, errors.New("GetLastBlockNumber is fail")
	}
}
