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

func GetIdentifiers(ch chan Identifier, m_id, m_type string) {
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
	client.Db = 1
	val, _ = client.Get(or.Mappings[0].OrgRefrenceId)
	var id Identifier
	formatJson(val, &id)
	ch <- id
}
