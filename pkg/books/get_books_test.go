package books

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"github.com/vivekdubey/fiber-api/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (h handler) assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func TestGetBooks(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	ht := &handler{
		DB: db,
	}

	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	// req, err := http.NewRequest("http://localhost/books", nil)
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected while creating request", err)
	// }
	// w := httptest.NewRecorder()

	// mock.
	rows := sqlmock.NewRows([]string{"id", "title", "author", "description"}).
		AddRow(1, "book 1", "author 1", "description 1").
		AddRow(2, "book 2", "author 2", "description 2")
	// mock.ExpectQuery("^SELECT (.+) FROM books$").WillReturnRows(rows)
	ht.GetBooks(c)

	data := struct {
		Books []*models.Book
	}{Books: []*models.Book{
		{Id: 1, Title: "book 1", Author: "author 1", Description: "description 1"},
		{Id: 2, Title: "book 2", Author: "author 2", Description: "description 2"},
	}}

	ht.assertJSON(c.Body(), data, t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
