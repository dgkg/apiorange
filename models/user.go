package models

import (
	"encoding/json"
)

type User struct {
	UID      int    `json:"uid"`
	UserName string `json:"user_name"`
	Password string `json:"pass"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func (u User) MarshalJSON() ([]byte, error) {
	aux := struct {
		UID      int    `json:uid`
		UserName string `json:user_name`
		Name     string `json:name`
		Email    string `json:email`
	}{
		UID:      u.UID,
		UserName: u.UserName,
		Name:     u.Name,
		Email:    u.Email,
	}

	return json.Marshal(aux)
}