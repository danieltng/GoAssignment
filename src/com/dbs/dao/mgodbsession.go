package dao

import (
	"log"
	mgo "gopkg.in/mgo.v2"
)

type fn func(database *mgo.Database)

type Datastore struct {
	endpoint		string
	databaseName	string
}

func NewDatastore(_endpoint string, _databaseName string) (dataStore *Datastore) {
	dataStore = &Datastore{
		databaseName: _databaseName,
		endpoint: _endpoint,
	}
	return
}

func (dataStore *Datastore) RunStatement(f fn) {
	session, err := mgo.Dial(dataStore.endpoint)
	if err != nil {
		log.Fatal("Database cannot be connected sucessfully");
		return
	}
	defer session.Close()
	databaseWithName := session.DB(dataStore.databaseName)
	f(databaseWithName)
}
