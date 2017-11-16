package model

import "github.com/gocraft/dbr"

type (
	users struct {
		ID        int    `db:"id"`
		Email     string `db:"email"`
		Username  string `db:"username"`
		Viewname  string `db:"viewname"`
	}

	usersJSON struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		Viewname  string `json:"viewname"`
	}

	responseData struct {
		Users []users `json:"users"`
	}
)

var (
	tablename = "users"
	seq       = 1
	conn, _   = dbr.Open("mysql", "my_app:secret@/my_app", nil)
	sess      = conn.NewSession(nil)
)
