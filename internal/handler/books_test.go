package handler

import (
	"bytes"
	"elasticSearch/internal/models"
	"elasticSearch/internal/services"
	mock_services "elasticSearch/internal/services/mocks"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateBook(t *testing.T) {
	type mockBehavior func(s *mock_services.MockBooks, book models.Book)

	testTable := []struct {
		name                 string
		inputBody            string
		inputBook            models.Book
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"name","pageCount":300,"author":"docs","description":["текст","Примеры"],"authorEmail":"author.email.com"}`,
			inputBook: models.Book{
				Name:        "name",
				PageCount:   300,
				Author:      "docs",
				Description: []string{"текст", "Примеры"},
				AuthorEmail: "author.email.com",
			},
			mockBehavior: mockBehavior(func(s *mock_services.MockBooks, book models.Book) {
				s.EXPECT().CreateBook(book).Return(1, nil)
			}),
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "bad Request",
			mockBehavior:         mockBehavior(func(s *mock_services.MockBooks, book models.Book) {}),
			expectedStatusCode:   400,
			expectedResponseBody: `{"code":400,"message":"Bad Request"}`,
		},
		//TODO need implement internal server error case
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			//Init app dependencies
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//create repository
			repo := mock_services.NewMockBooks(ctrl)
			test.mockBehavior(repo, test.inputBook)

			// create services
			service := &services.Service{Books: repo}
			handlers := Handler{service: service}

			//new fiber app
			app := fiber.New()
			app.Post("/create", handlers.CreateBook)

			//make request
			req := httptest.NewRequest("POST", "/create", bytes.NewBufferString(test.inputBody))
			req.Header.Set("Content-Type", "application/json")
			//fiber has own test method
			resp, err := app.Test(req, -1)
			if err != nil {
				t.Error(err)
			}
			//read response body for asserting with expected
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}
			//assert testing result
			assert.Equal(t, resp.StatusCode, test.expectedStatusCode)
			assert.Equal(t, string(body), test.expectedResponseBody)
		})
	}
}

func TestHandler_Update(t *testing.T) {
	type mockBehavior func(books *mock_services.MockBooks, inputBody models.Book)

	testTable := []struct {
		name                 string
		inputString          string
		inputBody            models.Book
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			inputString: `{"name":"update name"}`,
			inputBody:   models.Book{Name: "update name"},
			mockBehavior: mockBehavior(func(s *mock_services.MockBooks, book models.Book) {
				s.EXPECT().Update(book).Return(nil)
			}),
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `{"message":"OK"}`,
		},
		{
			name:                 "bad request",
			mockBehavior:         mockBehavior(func(s *mock_services.MockBooks, book models.Book) {}),
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"code":400,"message":"Bad Request"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			//Init dependencies
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//create repository
			repo := mock_services.NewMockBooks(ctrl)
			test.mockBehavior(repo, test.inputBody)

			//create services
			service := &services.Service{Books: repo}

			//create handler layer
			handlers := Handler{service: service}

			//create new fiber app
			app := fiber.New()
			app.Put("/update", handlers.Update)

			//create request
			req := httptest.NewRequest("PUT", "/update", bytes.NewBufferString(test.inputString))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req, -1)
			if err != nil {
				t.Error(err)
			}
			//read response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}
			//assert test results
			assert.Equal(t, resp.StatusCode, test.expectedStatusCode)
			assert.Equal(t, string(body), test.expectedResponseBody)

		})
	}
}

func TestHandler_DeleteById(t *testing.T) {
	type mockBehavior func(books *mock_services.MockBooks, ids models.DeleteIds)

	testTable := []struct {
		name                 string
		inputString          string
		inputIds             models.DeleteIds
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			inputString: `{"ids":[1,2,3]}`,
			inputIds:    models.DeleteIds{Ids: []int{1, 2, 3}},
			mockBehavior: mockBehavior(func(s *mock_services.MockBooks, ids models.DeleteIds) {
				s.EXPECT().Delete(ids).Return(nil)
			}),
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `{"message":"OK"}`,
		},
		{
			name:                 "bad Request",
			mockBehavior:         mockBehavior(func(s *mock_services.MockBooks, ids models.DeleteIds) {}),
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"code":400,"message":"Bad Request"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			//Init dependencies
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//create repository
			repo := mock_services.NewMockBooks(ctrl)
			test.mockBehavior(repo, test.inputIds)

			//create services
			service := &services.Service{Books: repo}
			handlers := Handler{service: service}

			//new fiber app
			app := fiber.New()
			app.Delete("/delete", handlers.DeleteById)

			//create request
			req := httptest.NewRequest("DELETE", "/delete", bytes.NewBufferString(test.inputString))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req, -1)
			if err != nil {
				t.Error(err)
			}

			//read response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}

			//assert got test results
			assert.Equal(t, resp.StatusCode, test.expectedStatusCode)
			assert.Equal(t, string(body), test.expectedResponseBody)
		})
	}
}

func TestHandler_Search(t *testing.T) {
	type mockBehavior func(books *mock_services.MockBooks, searchInput string)

	testTable := []struct {
		name                 string
		searchInput          string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			searchInput: "book",
			mockBehavior: mockBehavior(func(s *mock_services.MockBooks, searchInput string) {
				s.EXPECT().Search(searchInput).Return([]models.Book{
					{
						Id:          1,
						Name:        "book",
						PageCount:   300,
						Author:      "name",
						AuthorEmail: "email.com",
						Description: []string{"hello", "how", "are", "you"},
					},
				}, nil)
			}),
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `[{"id":1,"name":"book","pageCount":300,"author":"name","description":["hello","how","are","you"],"authorEmail":"email.com"}]`,
		},
		{
			name: "internal server error",
			mockBehavior: mockBehavior(func(s *mock_services.MockBooks, searchInput string) {
				s.EXPECT().Search(searchInput).Return(nil, errors.New((fiber.ErrInternalServerError).Error()))
			}),
			expectedStatusCode:   fiber.StatusInternalServerError,
			expectedResponseBody: `{"code":500,"message":"Internal Server Error"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			//init dependencies
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//crete repository
			repo := mock_services.NewMockBooks(ctrl)
			test.mockBehavior(repo, test.searchInput)

			//create service
			service := &services.Service{Books: repo}
			handlers := Handler{service: service}

			//new fiber app
			app := fiber.New()
			app.Get("/search", handlers.Search)

			//create request
			req := httptest.NewRequest("GET", fmt.Sprintf("/search?find=%s", test.searchInput), nil)
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req, -1)
			if err != nil {
				t.Error(err)
			}

			//reade response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}

			//assert got results
			assert.Equal(t, resp.StatusCode, test.expectedStatusCode)
			assert.Equal(t, string(body), test.expectedResponseBody)
		})
	}
}
