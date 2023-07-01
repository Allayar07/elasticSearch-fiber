package elasticSearch_client

import "github.com/elastic/go-elasticsearch/v8"

func SetUpElasticSearch() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"}, // Elasticsearch server address
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return es, nil
}
