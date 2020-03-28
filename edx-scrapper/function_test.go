package edxscrapper

import (
	"fmt"
	"testing"
)

func TestScrapper(t *testing.T) {
	res, err := searchForCourse([]string{"web"})
	if err != nil {
		t.Errorf("fail for error %v", err)
	}

	fmt.Println(len(res))
}
