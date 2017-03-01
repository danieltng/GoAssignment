package dao

import (
	mgo "gopkg.in/mgo.v2"
	"com/dbs/model"
	//uuid "github.com/nu7hatch/gouuid"
	"log"
)

type SearchResultDAO struct {
	datastore *Datastore
}

func NewSearchResultDAO() (searchResultDAO *SearchResultDAO) {
	searchResultDAO = &SearchResultDAO{
		datastore: NewDatastore("localhost:27017", "mydb"),
	}
	return
}

func (searchResultDAO *SearchResultDAO) AddSearchResult(searchResult *model.SearchResult) {
	var f func(database *mgo.Database)
	f = func (database *mgo.Database) {
		coll := database.C("mycoll")
		err := coll.Insert(searchResult.ConvertToDTO())
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	searchResultDAO.datastore.RunStatement(f)
}
/*
func (searchResultDAO *SearchResultDAO) GetSearchResultByTitle(title string) (searchResultDAO *model.SearchResult) {
	var f func(database *mgo.Database)
	f = func (database *mgo.Database) {
		coll := database.C("mycoll")
		err := coll.Find()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	searchResultDAO.datastore.RunStatement(f)
}
*/
