package postgres

import (
	"context"
	"elasticSearch/internal/models"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
)

type BooksRepo struct {
	db *pgxpool.Pool
}

func NewBooksRepo(db *pgxpool.Pool) *BooksRepo {
	return &BooksRepo{
		db: db,
	}
}

func (r *BooksRepo) Create(book models.Book) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (name, page_count, author, author_email, description) VALUES ($1, $2, $3, $4, $5) RETURNING id`, "books")
	row := r.db.QueryRow(context.Background(), query, book.Name, book.PageCount, book.Author, book.AuthorEmail, book.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BooksRepo) DeleteBooks(input models.DeleteIds) error {
	deleteIds := strings.Trim(strings.Replace(fmt.Sprint(input.Ids), " ", ",", -1), "[]")
	query := fmt.Sprintf(`DELETE FROM %s WHERE id IN (%s)`, "books", deleteIds)
	_, err := r.db.Exec(context.Background(), query)
	if err != nil {
		return err
	}
	return nil
}
func (r *BooksRepo) UpdateBook(book models.Book) error {
	setValues := make([]string, 0)
	arguments := make([]interface{}, 0)
	argumentId := 1
	if book.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argumentId))
		arguments = append(arguments, book.Name)
		argumentId++
	}
	if book.PageCount != 0 {
		setValues = append(setValues, fmt.Sprintf("page_count=$%d", argumentId))
		arguments = append(arguments, book.PageCount)
		argumentId++
	}
	if book.Author != "" {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argumentId))
		arguments = append(arguments, book.Author)
		argumentId++
	}
	if book.AuthorEmail != "" {
		setValues = append(setValues, fmt.Sprintf("author_email=$%d", argumentId))
		arguments = append(arguments, book.AuthorEmail)
		argumentId++
	}
	if book.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argumentId))
		arguments = append(arguments, book.Description)
		argumentId++
	}
	arguments = append(arguments, book.Id)
	updateValues := strings.Join(setValues, ",")
	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d `, "books", updateValues, argumentId)
	tag, err := r.db.Exec(context.Background(), query, arguments...)
	if err != nil {
		return err
	}
	if tag.RowsAffected() != 0 {
		fmt.Println("update failed")
	}
	return nil
}

func (r *BooksRepo) GetForSync() (books []models.Book, err error) {
	var book models.Book
	query := fmt.Sprintf(`SELECT id, name, page_count, author, author_email, description FROM %s`, "books")
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&book.Id, &book.Name, &book.PageCount, &book.Author, &book.AuthorEmail, &book.Description); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
