package handler

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	. "../model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

var (
	Tablename = "users"
	Seq       = 1
	Conn, _   = dbr.Open("mysql", "my_app:secret@/my_app", nil)
	Sess      = Conn.NewSession(nil)
)


func InsertUser(c echo.Context) error {
	u := new(UsersJSON)
	if err := c.Bind(u); err != nil {
		return err
	}
	Sess.InsertInto(Tablename).Columns("id", "email", "username", "viewname").Values(u.ID, u.Email, u.Username, u.Viewname).Exec()

	return c.NoContent(http.StatusOK)
}

func SelectUsers(c echo.Context) error {
	var u []Users
	Sess.Select("*").From(Tablename).Load(&u)
	response := new(ResponseData)
	response.Users = u
	return c.JSON(http.StatusOK, response)
}

func SelectUser(c echo.Context) error {
	var m Users
	id := c.Param("id")
	Sess.Select("*").From(Tablename).Where("id = ?", id).Load(&m)
	return c.JSON(http.StatusOK, m)
}

func UpdateUser(c echo.Context) error {
	u := new(UsersJSON)
	if err := c.Bind(u); err != nil {
		return err
	}

	attrsMap := map[string]interface{}{"id": u.ID, "email": u.Email, "username": u.Username, "viewname": u.Viewname}
	Sess.Update(Tablename).SetMap(attrsMap).Where("id = ?", u.ID).Exec()
	return c.NoContent(http.StatusOK)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	Sess.DeleteFrom(Tablename).
		Where("id = ?", id).
		Exec()

	return c.NoContent(http.StatusNoContent)
}
