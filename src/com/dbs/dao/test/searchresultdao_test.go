package assignment

import (
	"com/dbs/model"
	"com/dbs/dao"
	//"com/dbs/util"
	"testing"
	//mgo "gopkg.in/mgo.v2"
	//"log"
)

func TestSearchResultDAOAddAndFindIt(t *testing.T) {
	searchResultDAO := dao.NewSearchResultDAO()
	mockSearchResult := model.NewSearchResult("MockData2", 1.3, 14, "MockAddress1")
	searchResultDAO.AddSearchResult(mockSearchResult)
}
