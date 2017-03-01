package assignment

import (
	"com/dbs/service"
	"testing"
	"runtime"
	"sync"
	"com/dbs/util"
)
func init() {
	//assuming you are using 4 cores cpu
	runtime.GOMAXPROCS(4)
}


func TestSearchByCriteria(t *testing.T) {
	sliceArr := service.SearchByCriteria("Starbucks", "Singapore")
	util.AssertTrue(t, len(sliceArr) > 0, "Valid test data should have entries inside the list")
}

func TestSearchByCriteriaNeg(t *testing.T) {
	sliceArr := service.SearchByCriteria("", "Dubai")
	util.AssertTrue(t, len(sliceArr) == 0, "Invalid to search with empty query")
	sliceArr = service.SearchByCriteria("Starbucks", "")
	util.AssertTrue(t, len(sliceArr) == 0, "Invalid to search with empty query")
}

/*
	Testing the function by just invoking a parallel code running
	and see the performance
*/

func TestSearchByCriteriaParallel(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	
	for i:=0; i < 10; i++ {
		go func() {
			defer wg.Done()
			service.SearchByCriteria("Starbucks", "Singapore")
		}()
	}
	t.Logf("%s", "Waiting To Finish")
	wg.Wait()
}

/*
	Benchmarking with Go language testing capability
	but this is sequential, it does the aggregation
	of test results for you
*/
func BenchmarkSearchByCriteria(b *testing.B) {
	for i := 0; i < b.N; i++ {
		service.SearchByCriteria("Starbucks", "Singapore")
	}
}
