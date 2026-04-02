package fortnox_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestGetArticleByNumber(t *testing.T) {
	req := client.NewGetArticleByNumberRequest()
	req.PathParams().ArticleNumber = "TEST-001"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
