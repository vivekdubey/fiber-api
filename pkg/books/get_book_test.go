package books

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Test_handler_GetBook(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				DB: tt.fields.DB,
			}
			if err := h.GetBook(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("handler.GetBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
