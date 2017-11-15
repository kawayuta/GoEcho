package handlers

import (
	"net/http"
	"strconv"
)

func insertUser(c echo.Context) error {
	u := new(userinfoJSON)
	if err := c.Bind(u); err != nil {
		return err
	}

	sess.InsertInto(tablename).Columns("id", "email", "first_name", "last_name").Values(u.ID, u.Email, u.Firstname, u.Lastname).Exec()

	return c.NoContent(http.StatusOK)
}

func selectUsers(c echo.Context) error {
	var u []userinfo

	sess.Select("*").From(tablename).Load(&u)
	response := new(responseData)
	response.Users = u
	return c.JSON(http.StatusOK, response)
}
func selectUser(c echo.Context) error {
	var m userinfo
	id := c.Param("id")
	sess.Select("*").From(tablename).Where("id = ?", id).Load(&m)
	return c.JSON(http.StatusOK, m)

}

func updateUser(c echo.Context) error {
	u := new(userinfoJSON)
	if err := c.Bind(u); err != nil {
		return err
	}

	attrsMap := map[string]interface{}{"id": u.ID, "email": u.Email, "first_name": u.Firstname, "last_name": u.Lastname}
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
