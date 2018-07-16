package db

import (
	"log"
	"gopkg.in/mgo.v2"
)

type DB struct {}

func (db *DB) SetUp(mongoInfo *mgo.DialInfo) (*mgo.Database, error) {
	mongoSession, err := mgo.DialWithInfo(mongoInfo)
	if err != nil {
		return nil, err
	}
	mongoSession.SetMode(mgo.Monotonic, false)
	mongoDB := mongoSession.DB("ODS")
	log.Println("Connected to mongo", mongoInfo)

	return mongoDB, nil
}