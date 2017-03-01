package model

import (
	"fmt"
	"strconv"
	"strings"
	dto "com/dbs/dao/dto"
)

type SearchResult struct {
	title      string
	rating     float64
	numReviews int
	address    string
	blob   *[1e6]int
}


func NewSearchResult(title string, rating float64, numReviews int, address string) (searchResult *SearchResult) {
	return &SearchResult{
		title:      title,
		rating:     rating,
		numReviews: numReviews,
		address:    address,
	}
}

func (searchResult *SearchResult) SetTitle(title string) {
	searchResult.title = title
}

func (searchResult *SearchResult) SetRating(rating float64) {
	searchResult.rating = rating
}

func (searchResult *SearchResult) SetNumReviews(numReviews int) {
	searchResult.numReviews = numReviews
}

func (searchResult *SearchResult) SetAddress(address string) {
	searchResult.address = address
}

func (searchResult SearchResult) Address() string {
	return searchResult.address
}

func (searchResult SearchResult) Rating() float64 {
	return searchResult.rating
}

func (searchResult SearchResult) NumReviews() int {
	return searchResult.numReviews
}

func (searchResult SearchResult) Title() string {
	return searchResult.title
}

/*
func (searchResult *SearchResult) SetImage(imageBlob *[1e6]int) {
	searchResult.Blob = imageBlob
	searchResult.Blob[2] = 4;
	searchResult.Blob[3] = 5;
}
*/

func (searchResult *SearchResult) Print() {
	fmt.Println("Title: " + searchResult.title + "; Rating: " + 
		strconv.FormatFloat(searchResult.rating, 'f', 1, 64) + "; NumReviews: " + 
		strconv.Itoa(searchResult.numReviews) + "; Address: " + searchResult.address)
}

func (searchResult *SearchResult) IsSameLocation(location string) (result bool) {
	if (strings.Compare(searchResult.address, location) == 0) {
		result = true;
	} else {
		result = false;
	}
	return
}

func (searchResult *SearchResult) ConvertToDTO() (searchResultDTO *dto.SearchResultDTO) {
	searchResultDTO = &dto.SearchResultDTO{
		Title: searchResult.title,
		Rating: searchResult.rating,
		NumReviews: searchResult.numReviews,
		Address: searchResult.address,
	}
	return
}
