package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetArticleByNumber(t *testing.T) {
	req := client.NewGetArticleByNumberRequest()
	req.PathParams().ArticleNumber = "TEST-001"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
