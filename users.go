package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/gocraft/dbr"
)

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

func insertUser(c echo.Context) error {
	u := new(usersJSON)
	if err := c.Bind(u); err != nil {
		return err
	}
	sess.InsertInto(tablename).Columns("id", "email", "username", "viewname").Values(u.ID, u.Email, u.Username, u.Viewname).Exec()

	return c.NoContent(http.StatusOK)
}

func selectUsers(c echo.Context) error {
	var u []users
	sess.Select("*").From(tablename).Load(&u)
	response := new(responseData)
	response.Users = u
	return c.JSON(http.StatusOK, response)
}

func selectUser(c echo.Context) error {
	var m users
	id := c.Param("id")
	sess.Select("*").From(tablename).Where("id = ?", id).Load(&m)
	return c.JSON(http.StatusOK, m)
}

func updateUser(c echo.Context) error {
	u := new(usersJSON)
	if err := c.Bind(u); err != nil {
		return err
	}

	attrsMap := map[string]interface{}{"id": u.ID, "email": u.Email, "username": u.Username, "viewname": u.Viewname}
	sess.Update(tablename).SetMap(attrsMap).Where("id = ?", u.ID).Exec()
	return c.NoContent(http.StatusOK)
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	sess.DeleteFrom(tablename).
		Where("id = ?", id).
		Exec()

	return c.NoContent(http.StatusNoContent)
}
