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
	setValue := make([]string, 0)
	arguments := make([]interface{}, 0)
	argId := 1
	if book.Name != "" {
		setValue = append(setValue, fmt.Sprintf("name=$%d", argId))
		arguments = append(arguments, book.Name)
		argId++
	}
	if book.PageCount != 0 {
		setValue = append(setValue, fmt.Sprintf("page_count=$%d", argId))
		arguments = append(arguments, book.PageCount)
		argId++
	}
	if book.Author != "" {
		setValue = append(setValue, fmt.Sprintf("author=$%d", argId))
		arguments = append(arguments, book.Author)
		argId++
	}
	if book.AuthorEmail != "" {
		setValue = append(setValue, fmt.Sprintf("author_email=$%d", argId))
		arguments = append(arguments, book.AuthorEmail)
		argId++
	}
	if book.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description=$%d", argId))
		arguments = append(arguments, book.Description)
		argId++
	}
	arguments = append(arguments, book.Id)
	updateValues := strings.Join(setValue, ",")
	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d `, "books", updateValues, argId)
	tag, err := r.db.Exec(context.Background(), query, arguments...)
	if err != nil {
		return err
	}
	if tag.RowsAffected() != 0 {
		fmt.Println("update failed")
	}
	return nil
}
