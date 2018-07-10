package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"
	services "MA.Content.Services.OrgMapper/src/services"
	. "MA.Content.Services.OrgMapper/src/models"
	"fmt"
	"time"
	"errors"
)
func newSentinelPool() *redis.Pool {
	sntnl := &sentinel.Sentinel{
		Addrs: []string{"APC-LBMDCRED201:5000"},
		MasterName: "redis-eai-mongo",
		Dial: func(addr string) (redis.Conn, error) {
			timeout := 500 * time.Millisecond
			c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   64,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			masterAddr, err := sntnl.MasterAddr()
			if err != nil {
				return nil, err
			}
			c, err := redis.Dial("tcp", masterAddr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if !sentinel.TestRole(c, "master") {
				return errors.New("Role check failed")
			} else {
				return nil
			}
		},
	}
}

func createClient() redis.Conn {
	client, _ := newSentinelPool().Dial()
	return client
}

func loopGetIndentifiers(ch chan Identifier, m_id string, types []string, dbIdx int) {
	client := createClient()
	services.GetIdentifiersAll(client, ch, m_id, types, dbIdx)
	for i := 0; i < len(types); i++ {
		go services.GetIdentifiers(client, ch, m_id, types[i], dbIdx)
	}
}

func GetIdentifiersViaId(c *gin.Context) {
	types := GetOrgReferencePrefix()
	m_id := c.Param("id")
	ch := make(chan Identifier, len(types))
	loopGetIndentifiers(ch, m_id, types, 1)
	for i := range ch {
		fmt.Println(i)
	}
	close(ch)
}

func IsFuzzySearch(_type string) bool {
	return _type == "all"
}

func appendIdentifiers(origin Identifier, target Identifier) Identifier {
	origin.Bank_Identifier_Code = append(origin.Bank_Identifier_Code, target.Bank_Identifier_Code...)
	origin.BvD_ID = append(origin.BvD_ID, target.BvD_ID...)
	origin.Central_Index_Key = append(origin.Central_Index_Key, target.Central_Index_Key...)
	origin.Cmor_Company_Number = append(origin.Cmor_Company_Number, target.Cmor_Company_Number...)
	origin.Cu_Number = append(origin.Cu_Number, target.Cu_Number...)
	origin.Cuisp = append(origin.Cuisp, target.Cuisp...)
	origin.Dtc_Sales_Agent_Part_Num = append(origin.Dtc_Sales_Agent_Part_Num, target.Dtc_Sales_Agent_Part_Num...)
	origin.Duns_Number = append(origin.Duns_Number, target.Duns_Number...)
	origin.Equity_SEDOL = append(origin.Equity_SEDOL, target.Equity_SEDOL...)
	origin.Equity_Ticker = append(origin.Equity_Ticker, target.Equity_Ticker...)
	origin.FIGI = append(origin.FIGI, target.FIGI...)
	origin.FIGI_3 = append(origin.FIGI_3, target.FIGI_3...)
	origin.FIGI_4 = append(origin.FIGI_4, target.FIGI_4...)
	origin.FIGI_5 = append(origin.FIGI_5, target.FIGI_5...)
	origin.FIGI_Previous = append(origin.FIGI_Previous, target.FIGI_Previous...)
	origin.Ibm_Number = append(origin.Ibm_Number, target.Ibm_Number...)
	origin.Isin = append(origin.Isin, target.Isin...)
	origin.ISIN_Nubmer = append(origin.ISIN_Nubmer, target.ISIN_Nubmer...)
	origin.Legacy_Deal_Id = append(origin.Legacy_Deal_Id, target.Legacy_Deal_Id...)
	origin.Lei = append(origin.Lei, target.Lei...)
	origin.Lloyds_Syndicate_Performance = append(origin.Lloyds_Syndicate_Performance, target.Lloyds_Syndicate_Performance...)
	origin.Moodys_Deal_Number = append(origin.Moodys_Deal_Number, target.Moodys_Deal_Number...)
	origin.Rssd_Id = append(origin.Rssd_Id, target.Rssd_Id...)
	origin.Salesforce_Account_Id = append(origin.Salesforce_Account_Id, target.Salesforce_Account_Id...)
	origin.Salesforce_Opportunity_Id = append(origin.Salesforce_Opportunity_Id, target.Salesforce_Opportunity_Id...)
	origin.Sedol = append(origin.Sedol, target.Sedol...)
	origin.SIC_Code = append(origin.SIC_Code, target.SIC_Code...)
	origin.Ticker = append(origin.Ticker, target.Ticker...)
	origin.Tokyo_Stock_Exchange_Ticker_Symbol = append(origin.Tokyo_Stock_Exchange_Ticker_Symbol, target.Tokyo_Stock_Exchange_Ticker_Symbol...)
	return origin
}

func GetAllIdentifiers(m_id string, types []string, dbIdx int) Identifier {
	count := len(types)
	ch := make(chan Identifier, count)
	loopGetIndentifiers(ch, m_id, types, dbIdx)
	concatedId := Identifier{}
	for i := 0; i < count; i++ {
		c := <- ch
		concatedId = appendIdentifiers(concatedId, c)
	}
	return concatedId
}

func GetIdentifiers1(c *gin.Context) { 
	m_id := c.Param("id")
	var result Identifier
	if m_type := c.Param("type"); !IsFuzzySearch(m_type) {
		ch := make(chan Identifier, 1)
		client := createClient()
		go services.GetIdentifiers(client, ch, m_id, m_type, 1)
		for i := 0; i < 1; i++ {
			result = <- ch
		}
		close(ch)
	} else {
		result = GetAllIdentifiers(m_id, GetOrgReferencePrefix(), 1)
		// id = GetAllIdentifiers(m_id, GetInstReferencePrefix(), 2)
		// fmt.Println(id)
	}

	c.JSON(200, result)
}