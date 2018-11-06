package books

import (
	"github.com/labstack/echo"
	"test/db/models/book"
		"net/http"
	"test/db"
	"io/ioutil"
		"encoding/json"
	"log"
	"bytes"
)

type EmptyResponse struct {
	Status bool `json:"status"`
}

type ErrorResponse struct {
	Status bool `json:"status"`
	Reason string `json:"reason"`
}

type CreateBookPayload struct {
	Author string  `json:"author"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
}

func CreateBook(ctx echo.Context) error {

	var bodyBytes []byte

	if ctx.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request().Body)
	}

	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	payload := &CreateBookPayload{}

	unmarshalErr := json.Unmarshal(bodyBytes, payload)
	if unmarshalErr != nil {
		log.Println("Parsing error", unmarshalErr.Error())
		return ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Status: false,
			Reason: unmarshalErr.Error(),
		})
	}

	newBook := book.BookModel{
		Author: payload.Author,
		Title:  payload.Title,
		Price:  payload.Price,
	}

	db.Connection.Model(book.BookModel{}).Create(&newBook).Save(&newBook)
	return ctx.JSON(http.StatusOK, &EmptyResponse{
		Status: true,
	})
}
