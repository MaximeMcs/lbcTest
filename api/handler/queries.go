package handler

import (
	"database/sql"
	"net/http"
	"root/db"

	"github.com/gin-gonic/gin"
)

// QueriesHandler responds with the most hit FizzBuzz Query.
// QueriesHandler             godoc
// @Summary      Get most popular query as string
// @Description  Responds a string of the most hit FizzBuzz Query.
// @Tags         queries
// @Produce      json
// @Success      200
// @Router       /queries [get]
func QueriesHandler(c *gin.Context) {
	var query string
	var total int

	db, err := db.SetupDB()
	if err != nil {
		c.String(http.StatusBadRequest, "Database error.")
		return
	}

	row := db.QueryRow("SELECT query, COUNT( query ) total FROM queries GROUP BY query ORDER BY total DESC LIMIT 1")
	defer db.Close()
	switch err := row.Scan(&query, &total); err {
	case sql.ErrNoRows:
		c.String(http.StatusOK, "Send a first fizzbuzz query to get a result here.")
	case nil:
		c.String(http.StatusOK, "the most popular query is : %s, hit: %d times", query, total)
	default:
		c.String(http.StatusBadRequest, "Database error.")
		return
	}
}
