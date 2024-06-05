package controllers

import (
	"raidendb/internal/models"

	"github.com/sev-2/raiden"
)

type BookRequest struct {
	Id int64 `path:"id"`
}

type BookController struct {
	raiden.ControllerBase
	Http  string `path:"/books" type:"rest"`
	Model *models.Books
}
