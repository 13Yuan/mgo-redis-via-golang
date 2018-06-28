package services

import (
	"github.com/astaxie/goredis"
	. "MA.Content.Services.OrgMapper/src/models"
	"encoding/json"
	"os"
)

func formatJson(val []byte, or interface{}) {
	err := json.Unmarshal(val, &or)
	if err != nil {
		panic(err)
	}
}

func getReferenceId(dbIdx int, or OrgRefrence) string {
	var result string
	switch dbIdx {
		case 1: {
			result = or.Mappings[0].OrgRefrenceId
		}
		case 2: {
			result = or.Mappings[0].InstRefrenceId
		}
		case 3: {
			result = or.Mappings[0].SaleRefrenceId
		}
	}
	return result
}

func GetIdentifiers(ch chan Identifier, m_id, m_type string, dbIdx int) {
	var client goredis.Client
	client.Addr = os.Getenv("REDIS_CONNECTION")
	client.Db = 0
	val, err := client.Get(m_type + "-" + m_id)
	if err != nil {
		ch <- Identifier{}
		return
	}
	var or OrgRefrence
	formatJson(val, &or)
	client.Db = dbIdx
	val, _ = client.Get(getReferenceId(dbIdx, or))
	var id Identifier
	if val != nil && len(val) > 0 {
		formatJson(val, &id)
	}
	ch <- id
}

func Test() OrgRefrence {
	var client goredis.Client
	client.Addr = os.Getenv("REDIS_CONNECTION")
	client.Db = 0
	val, _ := client.Get("mdy_id-541000")
	var or OrgRefrence
	formatJson(val, &or)
	return or
}
