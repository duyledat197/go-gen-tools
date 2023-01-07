package elastic_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/elastic/go-elasticsearch/v7"
	es7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type ElasticClient struct {
	Client  *es7.Client
	Configs elasticsearch.Config
	Indexes []string
}

func NewcElasticClient(configs elasticsearch.Config) *ElasticClient {
	cfg := configs
	client, err := es7.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return &ElasticClient{
		Client:  client,
		Configs: configs,
	}
}

func (c *ElasticClient) CreateDocument(index string, content interface{}) error {
	b, err := json.Marshal(content)
	if err != nil {
		return fmt.Errorf("json.Marshaler: %w", err)
	}
	req := esapi.IndexRequest{
		Index:   index,
		Body:    bytes.NewReader(b),
		Refresh: "true",
		Pretty:  true,
	}
	res, err := req.Do(context.Background(), c.Client)
	if err != nil {
		return fmt.Errorf("Error getting response: %w", err)
	}
	defer res.Body.Close()
	return nil
}

func (c *ElasticClient) Search(index string, query interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("Error encoding query: %s", err)
	}
	resp, err := c.Client.Search(
		c.Client.Search.WithContext(context.Background()),
		c.Client.Search.WithIndex(index),
		c.Client.Search.WithBody(&buf),
		c.Client.Search.WithTrackTotalHits(true),
		c.Client.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("Error getting response: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	return body, nil
}

// ? pass pointer to parameter
func Convert[T any](from []byte, to *T) error {
	if err := json.Unmarshal(from, to); err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}
	return nil
}
