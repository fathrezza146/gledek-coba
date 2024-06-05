package controllers

import (
	"gledekcoba/internal/models"

	"github.com/sev-2/raiden"
)

type ToDoController struct {
	raiden.ControllerBase
	Http  string `path:"/todo" type:"rest"`
	Model models.ToDo
}
