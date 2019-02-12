package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/wyrdnixx/votegovue/src/gopoll/models"
)

type H map[string]interface{}

func GetPolls(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetPolls(db))
	}
}

func UpdatePoll(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var poll models.Poll
		c.Bind(&poll)

		index, _ := strconv.Atoi(c.Param("index"))
		fmt.Println("INFO: Name: ", poll.Name)
		fmt.Println("INFO: Down: ", poll.Downvotes)
		fmt.Println("INFO: Up: ", poll.Upvotes)
		id, err := models.UpdatePoll(db, index, poll.Name, poll.Upvotes, poll.Downvotes)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"affected": id,
			})
		}
		return err
	}
}

func UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var u models.UserData
		err := c.Bind(&u)
		fmt.Println(err.Error())
		fmt.Println("Info: SurName: ", u.SurName)
		fmt.Println("Info: GivenName: ", u.GivenName)
		return err
	}
}

func GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("INFO: GetUsers: ", c.JSON(http.StatusOK, models.GetUsers()))
		return c.JSON(http.StatusOK, models.GetUsers())
	}
}
