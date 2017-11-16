package model


type (
	Users struct {
		ID        int    `db:"id"`
		Email     string `db:"email"`
		Username  string `db:"username"`
		Viewname  string `db:"viewname"`
	}

	UsersJSON struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		Viewname  string `json:"viewname"`
	}

	ResponseData struct {
		Users []Users `json:"users"`
	}
)
