package api

import (
	"database/sql"
	"net/http"
	"root/db"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getLastPartOfCutString(str string, sep string) string {
	_, str, _ = strings.Cut(str, sep)
	return str
}

func buildMostPopularQuery(query string, total string) string {
	cleanQuery := getLastPartOfCutString(query, "?")
	querySlice := strings.Split(cleanQuery, "&")
	mostPopularQuery := "Les paramètres les plus utilisés sont : " + querySlice[0] + ", " + querySlice[1] + ", " + querySlice[2] + ", " + querySlice[3] + ", " + querySlice[4] + ", la requête a été appelée " + total + " fois."
	return mostPopularQuery
}

func getMostPopularQuery(c *gin.Context) *sql.Row {
	db := db.SetupDB()

	return db.QueryRowContext(c, "SELECT query, COUNT( query ) total FROM queries GROUP BY query ORDER BY total DESC LIMIT 1")
}

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

	row := getMostPopularQuery(c)
	switch err := row.Scan(&query, &total); err {
	case sql.ErrNoRows:
		c.String(http.StatusOK, "No rows were returned!")
	case nil:
		c.String(http.StatusOK, buildMostPopularQuery(query, strconv.Itoa(total)))
	default:
		panic(err)
	}
}
