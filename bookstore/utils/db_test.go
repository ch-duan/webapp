package utils

import (
	"testing"
	"webapp/bookstore/model"
)

func Test(t *testing.T) {
	book := model.Books{}
	AllValues(&book)
}
