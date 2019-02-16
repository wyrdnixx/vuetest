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

func PutTask() echo.HandlerFunc {

	return func(c echo.Context) error {
		fmt.Println("in handler-PutTask...")
		var task models.Task

		c.Bind(&task)

		//surname, err := strconv.Itoa(c.Param())

		fmt.Println("task-surname: " + c.Param("surname"))
		name := models.PutTask(task.Surname)
		fmt.Println("Surname: " + name)
		return c.JSON(http.StatusCreated, H{
			"created": name,
		})
		// Handle any errors
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
