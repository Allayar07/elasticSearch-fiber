package services

import (
	"bytes"
	"context"
	"elasticSearch/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"log"
	"strconv"
	"strings"
)

func IndexingToElasticSearch(ctx context.Context, id string, book models.Book, es *elasticsearch.Client) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(book); err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      "books",
		Body:       &buf,
		DocumentID: id,
		Refresh:    "true",
	}

	resp, err := req.Do(ctx, es)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.IsError() {
		return errors.New("error: " + resp.Status())
	}

	return nil
}

func SearchBook(searchInput string, es *elasticsearch.Client) ([]models.Book, error) {
	var (
		query map[string]interface{}
		buf   bytes.Buffer
	)
	query = map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":     searchInput,
				"fields":    []string{"name", "description"},
				"fuzziness": 7,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return []models.Book{}, err
	}

	req := esapi.SearchRequest{
		Index: []string{"books"},
		Body:  &buf,
	}
	resp, err := req.Do(context.Background(), es)
	if err != nil {
		return []models.Book{}, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return []models.Book{}, errors.New(resp.Status())
	}
	var hits struct {
		Hits struct {
			Hits []struct {
				Source models.Book `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&hits); err != nil {
		fmt.Println("Error here", err)
		return []models.Book{}, err
	}
	result := make([]models.Book, len(hits.Hits.Hits))

	for i, hit := range hits.Hits.Hits {
		result[i].Id = hit.Source.Id
		result[i].Description = hit.Source.Description
		result[i].Name = hit.Source.Name
		result[i].Author = hit.Source.Author
		result[i].AuthorEmail = hit.Source.AuthorEmail
		result[i].PageCount = hit.Source.PageCount
	}
	return result, nil
}

func DeleteFromElastic(input models.DeleteIds, es *elasticsearch.Client) error {
	var (
		req strings.Builder
	)
	for _, id := range input.Ids {
		req.WriteString(fmt.Sprintf(`{"delete":{"_index":"%s","_id":"%d"}}%s`, "books", id, "\n"))
	}
	// Send the Bulk API request, it used for delete multiple document in one request
	bulkRequest := esapi.BulkRequest{
		Body: strings.NewReader(req.String()),
	}
	resp, err := bulkRequest.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return errors.New("error when deleting: " + resp.Status())
	}
	return nil
}

func UpdateFRomElastic(book models.Book, es *elasticsearch.Client) error {
	input := map[string]interface{}{}
	if book.Name != "" {
		input = map[string]interface{}{
			"doc": map[string]interface{}{
				"name": book.Name,
			},
		}
	}
	if book.PageCount != 0 {
		pageInput := map[string]interface{}{
			"doc": map[string]interface{}{
				"pageCount": book.PageCount,
			},
		}
		mergeMaps(input, pageInput)
	}
	if book.Author != "" {
		authorInput := map[string]interface{}{
			"doc": map[string]interface{}{
				"author": book.Author,
			},
		}
		mergeMaps(input, authorInput)
	}
	if book.AuthorEmail != "" {
		emailInput := map[string]interface{}{
			"doc": map[string]interface{}{
				"authorEmail": book.AuthorEmail,
			},
		}
		mergeMaps(input, emailInput)
	}
	if book.Description != nil {
		descriptionInput := map[string]interface{}{
			"doc": map[string]interface{}{
				"description": book.Description,
			},
		}
		mergeMaps(input, descriptionInput)
	}
	updateRequest := esapi.UpdateRequest{
		Index:      "books",
		DocumentID: strconv.Itoa(book.Id),
		Body:       esutil.NewJSONReader(input),
	}
	res, err := updateRequest.Do(context.Background(), es)
	if err != nil {
		return fmt.Errorf("error sending update request: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return errors.New("Elasticsearch update request failed: " + res.Status())
	}
	return nil
}

func mergeMaps(firstMap, secondMap map[string]interface{}) {
	for key, value := range secondMap {
		if _, ok := firstMap[key]; ok {
			// If the key exists in the first map, merge the values recursively
			if secondMapSubMap, ok := value.(map[string]interface{}); ok {
				if firstMapSubMap, ok := firstMap[key].(map[string]interface{}); ok {
					mergeMaps(firstMapSubMap, secondMapSubMap)
				}
			}
		} else {
			// If the key doesn't exist in the first map, add it with the corresponding value
			firstMap[key] = value
		}
	}
}

func SyncWithDB(books []models.Book, esClient *elasticsearch.Client) error {
	var buff bytes.Buffer
	for _, book := range books {
		buff.WriteString(fmt.Sprintf(`{"index":{"_index":"%s","_id":"%d"}}%s`, "books", book.Id, "\n"))
		if err := json.NewEncoder(&buff).Encode(book); err != nil {
			log.Printf("encode error: %v", err)
			return err
		}
	}
	// send bulk request for best performance :)
	bulkRequest := esapi.BulkRequest{
		Body: &buff,
	}
	resp, err := bulkRequest.Do(context.Background(), esClient)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return errors.New("error when inserting: " + resp.Status())
	}
	return nil
}
