package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/search", search)
	r.Run("localhost:8080")
}

func search(c *gin.Context) {
	q := c.Query("q")

	// 環境変数からElasticsearchの設定を取得
	esConfig := getElasticsearchConfig()

	// Elasticsearchのパスとリクエストボディを作成
	esPath := fmt.Sprintf("%s/%s/_search/template", esConfig.URL, esConfig.Index)
	searchTemplateBody := createSearchTemplateBody(q)

	// Elasticsearchにリクエストを送信し、レスポンスを取得
	esSearchResponse, err := searchElasticsearch(esPath, searchTemplateBody, esConfig)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// レスポンスを変換してクライアントに返却
	searchResponse := toSearchResponse(esSearchResponse)
	c.JSON(http.StatusOK, searchResponse)
}

type ElasticsearchConfig struct {
	Index    string `json:"index"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Elasticsearchの設定を取得する関数
func getElasticsearchConfig() ElasticsearchConfig {
	return ElasticsearchConfig{
		Index:    os.Getenv("ELASTICSEARCH_INDEX"),
		URL:      os.Getenv("ELASTICSEARCH_URL"),
		Username: os.Getenv("ELASTICSEARCH_USERNAME"),
		Password: os.Getenv("ELASTICSEARCH_PASSWORD"),
	}
}

type SearchTemplateBody struct {
	ID     string `json:"id"`
	Params struct {
		Q    string `json:"q"`
		From int    `json:"from"`
		Size int    `json:"size"`
	} `json:"params"`
}

// 検索用テンプレートボディを作成する関数
func createSearchTemplateBody(q string) SearchTemplateBody {
	return SearchTemplateBody{
		ID: "donguri-search-template",
		Params: struct {
			Q    string "json:\"q\""
			From int    "json:\"from\""
			Size int    "json:\"size\""
		}{
			Q:    q,
			From: 0,
			Size: 20,
		},
	}
}

type Hit struct {
	Index  string  `json:"_index"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source struct {
		Episode int `json:"episode"`
	} `json:"_source"`
}

type Hits struct {
	Total struct {
		Value    int    `json:"value"`
		Relation string `json:"relation"`
	} `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hit   `json:"hits"`
}

type EsSearchResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits Hits `json:"hits"`
}

func searchElasticsearch(path string, searchTemplateBody SearchTemplateBody, config ElasticsearchConfig) (EsSearchResponse, error) {
	// リクエストの作成
	b, err := json.Marshal(searchTemplateBody)
	if err != nil {
		return EsSearchResponse{}, err
	}
	req, err := http.NewRequest("POST", path, bytes.NewBuffer(b))
	if err != nil {
		return EsSearchResponse{}, err
	}

	// ヘッダーの設定
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(config.Username, config.Password)

	// クライアントの作成
	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	// リクエストの送信
	res, err := client.Do(req)
	if err != nil {
		return EsSearchResponse{}, err
	}
	defer res.Body.Close()

	// レスポンスの読み込み
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return EsSearchResponse{}, err
	}

	// レスポンスのパース
	var esSearchResponse EsSearchResponse
	if err := json.Unmarshal(body, &esSearchResponse); err != nil {
		return EsSearchResponse{}, err
	}

	return esSearchResponse, nil
}

type SearchResponse struct {
	Total    int
	Episodes []struct {
		Episode int
	}
}

func toSearchResponse(esSearchResponse EsSearchResponse) SearchResponse {
	searchResponse := SearchResponse{
		Total:    esSearchResponse.Hits.Total.Value,
		Episodes: make([]struct{ Episode int }, 0, len(esSearchResponse.Hits.Hits)),
	}

	for _, hit := range esSearchResponse.Hits.Hits {
		searchResponse.Episodes = append(searchResponse.Episodes, struct{ Episode int }{
			Episode: hit.Source.Episode,
		})
	}
	return searchResponse
}
