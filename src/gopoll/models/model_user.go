package models

import (
	_ "database/sql"
)

type UserData struct {
	SurName   string `json:"surename"`
	GivenName string `json:"givenname"`
}
