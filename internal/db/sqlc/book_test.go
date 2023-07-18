package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vivekdubey/fiber-api/internal/util"
)

func addRandomBook(t *testing.T) Book {
	arg := AddBookParams{
		Title:       util.RandomBookTitle(),
		Author:      util.RandomBookAuthor(),
		Description: util.RandomDescription(),
	}

	book, err := testQueries.AddBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)
	require.Equal(t, arg.Author, book.Author)
	require.Equal(t, arg.Title, book.Title)
	require.Equal(t, arg.Description, book.Description)
	require.NotZero(t, book.ID)
	require.NotZero(t, book.CreatedAt)

	return book
}
func TestAddBook(t *testing.T) {
	addRandomBook(t)
}

func TestGetBook(t *testing.T) {
	book1 := addRandomBook(t)
	book2, err := testQueries.GetBook(context.Background(), book1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, book2)
	require.Equal(t, book1.Author, book2.Author)
	require.Equal(t, book1.Title, book2.Title)
	require.Equal(t, book1.Description, book2.Description)
	require.Equal(t, book1.ID, book2.ID)

}

func TestUpdateBook(t *testing.T) {
	book1 := addRandomBook(t)
	arg := UpdateBookParams{
		ID:          book1.ID,
		Title:       util.RandomBookTitle(),
		Description: util.RandomDescription(),
		Author:      util.RandomBookAuthor(),
	}

	updatedBook, err := testQueries.UpdateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedBook)
	require.Equal(t, arg.Author, updatedBook.Author)
	require.Equal(t, arg.Title, updatedBook.Title)
	require.Equal(t, arg.Title, updatedBook.Title)
	require.Equal(t, book1.ID, updatedBook.ID)

}

func TestDeleteBook(t *testing.T) {
	book1 := addRandomBook(t)
	err := testQueries.DeleteBook(context.Background(), book1.ID)
	require.NoError(t, err)

	book2, err := testQueries.GetBook(context.Background(), book1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, book2)
}

func TestGetBooks(t *testing.T) {
	for i := 0; i < 10; i++ {
		addRandomBook(t)
	}

	arg := GetBooksParams{
		Limit:  5,
		Offset: 5,
	}

	books, err := testQueries.GetBooks(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, books, 5)

	for _, book := range books {
		require.NotEmpty(t, book)
	}
}
