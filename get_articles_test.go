package fortnox_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestGetArticles(t *testing.T) {
	req := client.NewGetArticlesRequest()
	// req.QueryParams().ArticleNumber = "TEST-002"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetArticlesAll(t *testing.T) {
	req := client.NewGetArticlesRequest()
	// req.QueryParams().ArticleNumber = "TEST-002"
	resp, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
