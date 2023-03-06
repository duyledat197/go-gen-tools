package elastic_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/duyledat197/go-gen-tools/config"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	es7 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"go.uber.org/zap"
)

type ElasticClient struct {
	Client   *es7.Client
	configs  elasticsearch.Config
	Indexes  []string
	Logger   *zap.Logger
	Address  string
	APIKey   string
	Database *config.Database
}

func (c *ElasticClient) Connect(ctx context.Context) error {
	c.configs.Addresses = append(c.configs.Addresses, c.Address)
	c.configs.APIKey = c.APIKey
	c.configs.EnableDebugLogger = true
	c.configs.Username = c.Database.UserName
	c.configs.Password = c.Database.Password
	c.configs.Logger = &elastictransport.ColorLogger{
		Output:             os.Stdin,
		EnableRequestBody:  true,
		EnableResponseBody: true,
	}

	client, err := es7.NewClient(c.configs)
	if err != nil {
		return fmt.Errorf("connect elastic error: %w", err)
	}
	if _, err := c.Client.Ping(); err != nil {
		return fmt.Errorf("ping elastic error: %w", err)
	}
	c.Client = client
	return nil
}

func (c *ElasticClient) Stop(ctx context.Context) error {
	return nil
}

func NewClient() *ElasticClient {
	return &ElasticClient{}
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
