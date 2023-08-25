package db

import (
	"log"
	"testing"
	"time"

	"github.com/0xcregis/easynode/store"
	"github.com/0xcregis/easynode/store/config"
	"github.com/sunjiangjun/xlog"
)

func Init() store.DbStoreInterface {
	cfg := config.LoadConfig("./../../cmd/store/config.json")
	return NewChService(&cfg, xlog.NewXLogger())
}

func TestClickhouseDb_AddMonitorAddress(t *testing.T) {
	s := Init()
	log.Println(s.AddMonitorAddress(0, &store.MonitorAddress{Address: "0xe5cB067E90D5Cd1F8052B83562Ae670bA4A211a8", Token: "5fe5f231-7051-4caf-9b52-108db92edbb4", BlockChain: 0, TxType: "0x2", Id: time.Now().UnixMilli()}))
}

func TestClickhouseDb_GetAddressByToken(t *testing.T) {
	s := Init()
	log.Println(s.GetAddressByToken(200, "5fe5f231-7051-4caf-9b52-108db92edbb4"))
}

func TestClickhouseDb_NewToken(t *testing.T) {
	s := Init()
	log.Println(s.NewToken(&store.NodeToken{
		Token: "1231token",
		Email: "123@qq.com",
		Id:    time.Now().UnixMilli(),
	}))
}

func TestClickhouseDb_UpdateToken(t *testing.T) {

	s := Init()
	log.Println(s.UpdateToken("1231token", &store.NodeToken{
		Token: "1231token",
		Email: "125@qq.com",
		Id:    time.Now().UnixMilli(),
	}))
}

func TestClickhouseDb_GetNodeTokenByEmail(t *testing.T) {
	s := Init()
	log.Println(s.GetNodeTokenByEmail("125@qq.com"))
}

func TestClickhouseDb_GetAddressByToken2(t *testing.T) {
	s := Init()
	log.Println(s.GetAddressByToken2("36ee0ad5-f4bc-4bca-a1dc-c51db006e249"))
}

func TestClickhouseDb_DelMonitorAddress(t *testing.T) {
	s := Init()
	log.Println(s.DelMonitorAddress(200, "36ee0ad5-f4bc-4bca-a1dc-c51db006e249", "0x28c6c06298d514db089934071355e5743bf21d61"))
}

func TestClickhouseDb_NewSubFilter(t *testing.T) {
	s := Init()
	t.Log(s.NewSubFilter([]*store.SubFilter{{Token: "token", BlockChain: 205, TxCode: "1"}}))
}

func TestClickhouseDb_NewBackupTx(t *testing.T) {
	s := Init()
	t.Log(s.NewBackupTx(200, []*store.BackupTx{{ChainCode: 200, ID: uint64(time.Now().UnixMilli()), From: "0x123", To: "0x456", Signed: "0xdddd", Status: 1, Response: "{\"id\":\"12312321\"}"}}))
}

func TestClickhouseDb_GetSubFilter(t *testing.T) {
	s := Init()
	t.Log(s.GetSubFilter("token", 205, "1"))
}

func TestClickhouseDb_DelSubFilter2(t *testing.T) {
	s := Init()
	t.Log(s.DelSubFilter2(&store.SubFilter{Token: "token"}))
}

func TestClickhouseDb_DelSubFilter(t *testing.T) {
	s := Init()
	t.Log(s.DelSubFilter(1691999490145550000))
}
