package cache

import (
	db "om-tool/src/db"
	. "om-tool/src/config"
	. "om-tool/src/models"
	"strconv"
	"fmt"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"time"
	"strings"
	"gopkg.in/mgo.v2"
	"log"
	"github.com/gomodule/redigo/redis"
)

var (
	loopCount = 5000
	loopTake = 10000
	curBP db.BreakPoint
)

func init() {
	_io := lastTransInfo()
	curBP = getBreakpoint(_io)
}

func createMongoInfo(mongoURL string) *mgo.DialInfo{
	// mongodb://eagle_app_user:eagleappuser@ftc-lbeagmdb306:27017,ftc-lbeagmdb307:27017,ftc-lbeagmdb308:27017/ODS
	infos := strings.Split(mongoURL, "@")
	auth := strings.Split(infos[0], "//")[1]
	user := strings.Split(auth, ":")[0]
	pwd := strings.Split(auth, ":")[1]
	dbinfos := strings.Split(infos[1], "/")
	dbs := strings.Split(dbinfos[0], ",")
	collection := dbinfos[1]
	
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    dbs,
		Timeout:  60 * time.Second,
		Database: collection,
		Username: user,
		Password: pwd,
	}
	return mongoDBDialInfo
}

func lastTransInfo() db.IO {
	_io, err := db.Io.Read()
	if err != nil {
		log.Println("Read io.yml error!")
	}
	return _io
}

func getBreakpoint(_io db.IO) db.BreakPoint {
	var bp db.BreakPoint
	if _io.Org > -1 {
		bp.Collection = 0
		bp.Number = _io.Org
		return bp
	}
	if _io.Inst > -1 {
		bp.Collection = 1
		bp.Number = _io.Inst
		return bp
	}
	if _io.Sale > -1 {
		bp.Collection = 2
		bp.Number = _io.Sale
		return bp
	}
	return db.BreakPoint{Collection: 2, Number: -1}
}

func SetupCache(config Config, mongoURL string) error {
	log.Println("Populating cache.", config.Name)
	_db := db.DB{}

	mongoInfo := createMongoInfo(mongoURL)
	mongoDb, err := _db.SetUp(mongoInfo)
	if err != nil {
		log.Println("Failed to connect to dbs", err)
		return err
	}
	defer mongoDb.Session.Close()
	for i := curBP.Collection; i < 2; i++ {
		if err := loopHandleCollections(i, config, mongoDb); err != nil {
			return err
		}
	}
	isAllCompleted := curBP.Collection == 2 && curBP.Number == -1
	if !isAllCompleted {
		chs := make(chan bool, 1)
		go func() {
			handleSaleCollection(mongoDb, chs)
		}()
		<- chs
	}
	db.Io.Write("sale_reference", -1)
	return nil
}

func getRedisHashKeys(conn redis.Conn) (string, error) {
	tempKey, err := redis.Int64(conn.Do("INCR", "index"))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("maps:%d", tempKey), nil
}

func setRedis(_id []IdentifiersObj, key, id string) error {
	val, err := json.Marshal(_id)
	key = strings.Replace(key, " ", "", -1)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	if err := db.Set(key, val); err != nil {
		log.Printf("Could not send SET %v", err)
		return err
	}
	id = strings.Replace(id, " ", "", -1)
	setData(id, key)
	return nil
}

func dedulplicateKey(_id []IdentifiersObj, v []byte, id IDs) []IdentifiersObj {
  	if err := json.Unmarshal(v, &_id); err != nil {
		log.Println(err)	  
	}
	var mergedIds []IdentifiersObj
	var newId []IdentifiersObj // remove empty id
	for _, _cId := range id.Identifiers {
		exist := false
		for _, __id := range _id {
			if __id.Value == _cId.Value && __id.Label == _cId.Label {
				exist = true
				break
			}
		}
		if !exist {
			mergedIds = append(mergedIds, _cId)
		}
	}
	_id = append(mergedIds, _id...)
	for _, __id := range _id {
		if __id.Label != "" {
			newId = append(newId, __id)
		}
	}
	return newId
}

func generateKey(ident IdentifiersObj) string {
	return strings.Replace(ident.Label, " ", "", -1) + "-" + strings.Replace(ident.Value, " ", "", -1)
}

func setData(key, value string) error {
	var valStr string
	if ok, _ := db.Exists(key); ok {
		v, _:= db.Get(key)
		if !strings.Contains(string(v), value) {
			valStr = string(v) + "," + value
		}
	} else {
		valStr = value
	}
	if valStr != "" {
		if err := db.Set(key, []byte(valStr)); err != nil {
			log.Printf("Could not send SET %v", err)
			return err
		}
	}
	return nil
}

func createOrgRoleIdentifier(s Sale, _key string) []KeyValue {
	var kvs []KeyValue
	for _, s_org := range s.SaleOrgs {
		_label := ""
		switch s_org.OrganizationRole {
			case "Issuer":
				_label = "Issuer_Org_Id"
			break
			case "Obligor":
				_label = "Obligor_Org_Id"
			break
		}
		kv := createIdentifierKV(_key, _label, strconv.Itoa(s_org.OrgId))
		kvs = append(kvs, kv)
	}
	return kvs
}

func createIdentifierKV(_key, label, value string) KeyValue {
	var id IdentifiersObj
	id.Source = []string{}
	id.Value = value
	id.Label = label
	idByte, err := json.Marshal(id)
	if err != nil {
		log.Println("Marshal identifier error.")
	}
	return KeyValue{Key: _key, Value: idByte }
}

func createOrgIdentifier(s Sale, _key string) []KeyValue {
	var kvs []KeyValue
	for _, oid := range s.OrgIds {
		kv := createIdentifierKV(_key, "org_id", strconv.Itoa(oid))
		kvs = append(kvs, kv)
	}
	return kvs
}

func createSaleKVS(_key string, s Sale) []KeyValue {
	var kvs []KeyValue
	kvs = append(kvs, createOrgRoleIdentifier(s, _key)...)
	kvs = append(kvs, createOrgIdentifier(s, _key)...)
	kvs = append(kvs, createIdentifierKV(_key, "sale_id", strconv.Itoa(s.InstId)))
	return kvs
}

func transferSale(sale []Sale) error {
	var allkvs []KeyValue
	for _, s := range sale {
		var kvs []KeyValue
		// add sale_id key
		_key := "sale_id-" + strconv.Itoa(s.InstId)
		kvs = append(kvs, createSaleKVS(_key, s)...)
		// add id key
		_key = strconv.Itoa(s.InstId)
		kvs = append(kvs, createSaleKVS(_key, s)...)
		// integrate both keys
		allkvs = append(allkvs, kvs...)
	}
	db.MSet(allkvs)
	return nil
}

func unique(stringSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{} 
    for _, entry := range stringSlice {
        if _, value := keys[entry]; !value && strings.Replace(entry, " ", "", -1) != "" {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func loopSADDIdentifiers(keys []string, ids []IdentifiersObj) []KeyValue {
	var newKVS []KeyValue
	for _, id := range ids {
		val, _ := json.Marshal(id)
		for _, k := range keys {
			newKVS = append(newKVS, KeyValue{Key: k, Value: val})
		}
	}
	return newKVS
}

func transferDataOnlyInsert(ids []IDs) error {
	var newKVS []KeyValue
	for _, id := range ids {
		var allKeys []string
		for _, ident := range id.Identifiers {
			key := generateKey(ident)
			idKey := strings.Replace(ident.Value, " ", "", -1)
			if key != "" && idKey != ""  {
				allKeys = append(allKeys, key)
				allKeys = append(allKeys, idKey)
			}
		}
		newKVS = append(newKVS, loopSADDIdentifiers(allKeys, id.Identifiers)...)
	}
	db.MSet(newKVS)
	return nil
}

func handleSaleCollection(mongoDb *mgo.Database,  _ch chan bool) {
	collection := "sale_reference"
	log.Printf("Processing: get mongodb: sale_reference data...")
	count, _ := mongoDb.C(collection).Find(bson.M{"inst_id": bson.M{"$exists": 1}}).Count()
	log.Printf("mongodb sale_reference count is " + strconv.Itoa(count) + " transferring...")
	gorouteNum := int(count / loopCount)
	chs := make(chan bool, gorouteNum)
	for i := 0; i < gorouteNum; i++ {
		go func(idx int) {
			var sale []Sale
			mongoDb.C(collection).Find(nil).Skip(idx * loopCount).Select(bson.M{"inst_id": 1, "org_id": 1, "sale_org": 1}).Limit(loopCount).All(&sale)
			if err := transferSale(sale); err != nil {
				log.Printf("Error transfer data %v", err)
			}
			chs <- true
		}(i)
	}
	for i := 0; i < gorouteNum; i++ {
		if !<- chs {
			log.Printf("Error transfer")
		}
	}
	db.Io.Write(collection, -1) 
	log.Printf("Processed " + collection + " all documents for query")
	_ch <- true
}

func handleCollection(collection string, mongoDb *mgo.Database,  _ch chan bool, skip int) {
	go func () { 
		db.Io.Write(collection, skip) 
	}()
	log.Printf("Processing: get mongodb: " + collection + " data, skip " + strconv.Itoa(skip) + "...")
	var count int
	if skip < 0 {
		count, _ = mongoDb.C(collection).Find(bson.M{"identifiers": bson.M{"$exists": 1}}).Select(bson.M{"identifiers": 1}).Count()
		skip = 0;
	} else {
		count, _ = mongoDb.C(collection).Find(bson.M{"identifiers": bson.M{"$exists": 1}}).Select(bson.M{"identifiers": 1}).Skip(skip).Limit(loopTake).Count()
	}
	log.Printf("mongodb " + collection + " count is " + strconv.Itoa(count) + " transferring...")
	gorouteNum := int(count / loopCount)
	chs := make(chan bool, gorouteNum)
	for i := 0; i < gorouteNum; i++ {
		go func(idx int) {
			var ids []IDs
			mongoDb.C(collection).Find(bson.M{"identifiers": bson.M{"$exists": 1}}).Skip(idx * loopCount + skip).Select(bson.M{"identifiers": 1}).Limit(loopCount).All(&ids)
			if err := transferDataOnlyInsert(ids); err != nil {
				log.Printf("Error transfer data %v", err)
			}
			chs <- true
		}(i)
	}
	for i := 0; i < gorouteNum; i++ {
		if !<- chs {
			log.Printf("Error transfer")
		}
	}
	log.Printf("Processed " + collection + " data, skip " + strconv.Itoa(skip) + " all documents for query")
	_ch <- true
}

func loopHandleCollections(idx int, config Config, mongoDb *mgo.Database) error {
	collection := config.Collections[idx].Collection
	log.Printf("Processing " + collection + "...")
	count, _ := mongoDb.C(collection).Find(bson.M{"identifiers": bson.M{"$exists": 1}}).Count()
	takes := int((count - curBP.Number) / loopTake)
	for i := 0; i < takes; i++ {
		skip := i * loopTake + curBP.Number
		chs := make(chan bool, 1)
		go func() {
			handleCollection(collection, mongoDb, chs, skip)
		}()
		<- chs
	}
	db.Io.Write(collection, -1)
	curBP = getBreakpoint(lastTransInfo())
	return nil
}