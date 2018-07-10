package services

import (
	"github.com/gomodule/redigo/redis"
	. "MA.Content.Services.OrgMapper/src/models"
	"encoding/json"
	"fmt"
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

func GetIdentifiersAll(client redis.Conn, ch chan Identifier, m_id string, m_type[] string, dbIdx int) {
	var args []interface{}
	for _, k := range m_type {
		args = append(args, k + "-" + m_id)
	}
	values, _ := redis.Strings(client.Do("MGET", args...))
	var ors []interface{}
	for _, v := range values {
		if v != "" {
			var or OrgRefrence
			val := []byte(v)
			formatJson(val, &or)
			ors = append(ors, getReferenceId(1, or))
		}
	}
	values, _ = redis.Strings(client.Do("MGET", ors...))
	for _, v := range values {
		if v != "" {
			var id Identifier
			val := []byte(v)
			formatJson(val, &id)
			fmt.Println(v)
			ch <- id
		}
	}
}

func GetIdentifiers(client redis.Conn, ch chan Identifier, m_id, m_type string, dbIdx int) {
	if _, err1 := client.Do("SELECT", 0); err1 != nil {
		return
	  }
	result, _ := client.Do("GET", m_type + "-" + m_id)
	val, err := redis.Bytes(result, nil)
	if err != nil {
		ch <- Identifier{}
		return
	}
	var or OrgRefrence
	formatJson(val, &or)
	client.Do("Select", 1)
	val, _ = redis.Bytes(client.Do("GET", getReferenceId(dbIdx, or)))
	var id Identifier
	if val != nil && len(val) > 0 {
		formatJson(val, &id)
	}
	ch <- id
}