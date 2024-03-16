package taskcreate

import (
	"github.com/sunjiangjun/supernode/common/chain"
	"github.com/sunjiangjun/supernode/task"
	"github.com/sunjiangjun/supernode/task/config"
	"github.com/sunjiangjun/supernode/task/service/taskcreate/bnb"
	"github.com/sunjiangjun/supernode/task/service/taskcreate/btc"
	"github.com/sunjiangjun/supernode/task/service/taskcreate/ether"
	"github.com/sunjiangjun/supernode/task/service/taskcreate/filecoin"
	"github.com/sunjiangjun/supernode/task/service/taskcreate/polygonpos"
	"github.com/sunjiangjun/supernode/task/service/taskcreate/tron"
	"github.com/sunjiangjun/supernode/task/service/taskcreate/xrp"
	"github.com/sunjiangjun/xlog"
)

func NewApi(blockchain int64, log *xlog.XLog, v *config.BlockConfig) task.BlockChainInterface {
	if chain.GetChainCode(blockchain, "ETH", log) {
		return ether.NewEther(log, v)
	} else if chain.GetChainCode(blockchain, "TRON", log) {
		return tron.NewTron(log, v)
	} else if chain.GetChainCode(blockchain, "POLYGON", log) {
		return polygonpos.NewPolygonPos(log, v)
	} else if chain.GetChainCode(blockchain, "BSC", log) {
		return bnb.NewBnb(log, v)
	} else if chain.GetChainCode(blockchain, "FIL", log) {
		return filecoin.NewFileCoin(log, v)
	} else if chain.GetChainCode(blockchain, "XRP", log) {
		return xrp.NewXRP(log, v)
	} else if chain.GetChainCode(blockchain, "BTC", log) {
		return btc.NewBtc(log, v)
	}
	return nil
}
