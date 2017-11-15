package main

type (
	userinfo struct {
		ID         int    `db:"id"`
		Email      string `db:"email"`
		First_name string `db:"first_name"`
		Last_name  string `db:"last_name"`
	}

	userinfoJSON struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		Firstname string `json:"firstName"`
		Lastname  string `json:"lastName"`
	}

	responseData struct {
		Users []userinfo `json:"users"`
	}
)

var (
	tablename = "userinfo"
	seq       = 1
	conn, _   = dbr.Open("mysql", "root:@tcp(127.0.0.1:3306)/ws2", nil)
	sess      = conn.NewSession(nil)
)
