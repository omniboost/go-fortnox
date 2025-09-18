package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetArticles(t *testing.T) {
	req := client.NewGetArticlesRequest()
	req.QueryParams().ArticleNumber = "TEST-002"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
