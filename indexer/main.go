package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

type Document struct {
	Episode int    `json:"episode"`
	Script  string `json:"script"`
}

func NewDocument(episode int, script string) *Document {
	return &Document{
		Episode: episode,
		Script:  script,
	}
}

func main() {
	start := time.Now().UTC()

	// 環境変数を取得
	s3Bucket := os.Getenv("S3_BUCKET")
	s3Key := os.Getenv("S3_KEY")
	esIndex := os.Getenv("ELASTICSEARCH_INDEX")
	esUrl := os.Getenv("ELASTICSEARCH_URL")
	esUsername := os.Getenv("ELASTICSEARCH_USERNAME")
	esPassword := os.Getenv("ELASTICSEARCH_PASSWORD")

	// S3クライアントの作成
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	s3Client := s3.NewFromConfig(cfg)

	// CSVファイルをロード
	r, err := s3Client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &s3Bucket,
		Key:    &s3Key,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	episodes := []Document{}
	records, err := csv.NewReader(r.Body).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records[1:] {
		episodeNumber, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		episodes = append(episodes, Document{
			Episode: episodeNumber,
			Script:  record[1],
		})
	}

	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			esUrl,
		},
		Username: esUsername,
		Password: esPassword,
	})
	if err != nil {
		log.Fatal(err)
	}

	// インデックス作成
	indexName := esIndex + time.Now().Format("20060102150405")
	res, err := esClient.Indices.Create(indexName)
	if err != nil {
		log.Println(indexName)
		log.Fatalf("Cannot create index: %s", err)
	}
	res.Body.Close()

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  indexName,
		Client: esClient,
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}
	for _, e := range episodes {
		data, err := json.Marshal(e)
		if err != nil {
			log.Fatalf("Cannot encode episode %d: %s", e.Episode, err)
		}

		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action: "index",
				Body:   bytes.NewReader(data),
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
	}
	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	// ログ出力
	biStats := bi.Stats()
	dur := time.Since(start)
	if biStats.NumFailed > 0 {
		log.Fatalf(
			"Indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			humanize.Comma(int64(biStats.NumFailed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	} else {
		log.Printf(
			"Sucessfuly indexed [%s] documents in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	}
}
