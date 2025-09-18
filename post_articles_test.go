package fortnox_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestPostArticle(t *testing.T) {
	req := client.NewPostArticleRequest()
	article := req.RequestBody().Article
	article.Description = "Test from API #2"
	article.ArticleNumber = "TEST-002"
	req.RequestBody().Article = article
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
