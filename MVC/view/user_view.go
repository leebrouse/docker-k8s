package view

import (
	"fmt"
	"mvc/model"
)

func RenderUsers(users []model.User) string {
	result := "User List:\n"
	for _, user := range users {
		result+=fmt.Sprintf("ID: %d, Name: %s\n", user.Id, user.Name)
	}
	return result
}
