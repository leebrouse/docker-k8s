package controller

import (
	"mvc/model"
	"mvc/view"
	"net/http"
)

func UserShowHandler(w http.ResponseWriter, r *http.Request) {
	users := model.GetName()
	result := view.RenderUsers(users)

	w.Write([]byte(result))
}
