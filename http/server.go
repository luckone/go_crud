package http

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
	"os"
	"io"
	"test/http/routes/books"
)

func Connect() {
	e := echo.New()

	e.GET("/", hello)
	e.POST("/book/create", books.CreateBook)

	e.Logger.Fatal(e.Start(":3000"))
}

type SuccessUploadedFile struct {
	Reason    string `json:"reason"`
	FileName  string `json:"file_name"`
	CreatedAt int64  `json:"created_at"`
}

type MissedFileInRequest struct {
	Reason string `json:"reason"`
}

type ServerHandlingError struct {
	Reason string `json:"reason"`
}

func hello(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, &MissedFileInRequest{
			Reason: "fail",
		})
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &SuccessUploadedFile{
		Reason:    "Ok",
		FileName:  "penis.png",
		CreatedAt: time.Now().Unix(),
	})
}
