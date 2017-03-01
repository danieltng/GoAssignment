package service

import (
	"log"
	"github.com/PuerkitoBio/goquery"
	"com/dbs/model"
	"strconv"
	"strings"
	"com/dbs/util"
	"time"
	"com/dbs/dao"
)

func init() {
	log.Println("Initialized service successfully")
	//runtime.GOMAXPROCS(runtime.NumCPU())
}


/*
	Using slice is a performance consideration as it is copy the memory address only.
	It will be better when we do passing of values through function call
*/
func xmlProcessing(doc *goquery.Document) (searchResultList []*model.SearchResult) {
	searchResultList = make([]*model.SearchResult, 0)
	
	var searchResult *model.SearchResult
	doc.Find(".search-result-title").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".biz-name").First().Text()
		searchResult = model.NewSearchResult(title, 0, 0, "")
		searchResultList = append(searchResultList, searchResult)
	})
	
	doc.Find(".rating-large").Each(func(i int, s *goquery.Selection) {
		fullRatingStr := s.AttrOr("title", "")
		//log.Println("Index val: " + strconv.Itoa(i));
		if len(fullRatingStr) > 0 {
			fullRatingStr = fullRatingStr[0:3]
			rating,err := strconv.ParseFloat(fullRatingStr,64)
			if (err != nil) {
				log.Fatal(err)
			}
			searchResultList[i].SetRating(rating)
		} else {
			//do nothing as default has initialized as 0 rating
			//or we can throw HTML parsing exception to denotes something wrong with the HTML page itself
		}
	})
	
	doc.Find(".review-count").Each(func(i int, s *goquery.Selection) {
		reviewNumStr := strings.TrimSpace(s.Text())
		reviewNumStr = strings.TrimSuffix(reviewNumStr, " reviews")
		reviewNumStr = strings.TrimSuffix(reviewNumStr, " review")
		reviewNum,err := strconv.Atoi(reviewNumStr)
		if (err != nil) {
			log.Fatal(err)
		} else {
			searchResultList[i].SetNumReviews(reviewNum)
		}
	})
	
	doc.Find("div[class='secondary-attributes']").Each(func(i int, s *goquery.Selection) {
		addressStr := strings.TrimSpace(s.Find("address").Text())
		searchResultList[i].SetAddress(addressStr)
	})
	
	return
}


func SearchByCriteria(query string, location string) (searchResultList []*model.SearchResult) {
	if (len(query) > 0 && len(location) > 0) {
		log.Println("Validated input successfully")
	} else {
		log.Println("Either one input values are empty")
		return
	}
	
	start := time.Now()
	url := "https://www.yelp.com/search?find_desc="+ query + "&find_loc=" + location + "&ns=1"
	doc, err := goquery.NewDocument(url)
	log.Printf("GoRoutineID:%d HTTPCallingTiming:%s", util.GetGoRoutineID(), time.Since(start))
	if err != nil {
		log.Fatal(err)
		return
	}
	
	start = time.Now()
	searchResultList = xmlProcessing(doc)
	log.Printf("GoRoutineID:%d XMLProcessingTime:%s", util.GetGoRoutineID(), time.Since(start))
	log.Println("Length of slice: " + strconv.Itoa(len(searchResultList)))
	
	start = time.Now()
	persistSearchResult(searchResultList)
	log.Printf("GoRoutineID:%d DBProcessingTime:%s", util.GetGoRoutineID(), time.Since(start))
	
	return
}

func persistSearchResult(searchResultList []*model.SearchResult) {
	log.Println("Store in DB now")
	searchResultDAO := dao.NewSearchResultDAO();
	for _, value := range searchResultList {
		searchResultDAO.AddSearchResult(value)
	}
}

