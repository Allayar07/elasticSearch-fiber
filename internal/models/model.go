package models

type Book struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	PageCount   int64    `json:"pageCount"`
	Author      string   `json:"author"`
	Description []string `json:"description"`
	AuthorEmail string   `json:"authorEmail"`
}

type DeleteIds struct {
	Ids []int `json:"ids"`
}
