package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "host.docker.internal"
	DB_PORT     = "1234"
	DB_USER     = "root"
	DB_PASSWORD = "root"
	DB_NAME     = "lbc"
)

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func fizzBuzzIt(a int, b int, limit int, str1 string, str2 string) []string {
	strSlice := make([]string, limit)
	for i := 1; i <= limit; i++ {
		switch {
		case i%a == 0 && i%b == 0:
			strSlice[i-1] = str1 + str2
		case i%a == 0:
			strSlice[i-1] = str1
		case i%b == 0:
			strSlice[i-1] = str2
		default:
			strSlice[i-1] = strconv.Itoa(i)
		}
	}
	return strSlice
}

func fizzBuzzHandler(c *gin.Context) {
	if c.Query("int1") == "" ||
		c.Query("int2") == "" || c.Query("limit") == "" ||
		c.Query("str1") == "" || c.Query("str2") == "" {
		c.String(http.StatusBadRequest, "Parameter(s) missing.")
	} else {
		int1, err := strconv.Atoi(c.Query("int1"))
		checkErr(err)
		int2, err := strconv.Atoi(c.Query("int2"))
		checkErr(err)
		limit, err := strconv.Atoi(c.Query("limit"))
		checkErr(err)

		c.JSON(http.StatusOK, fizzBuzzIt(int1, int2, limit, c.Query("str1"), c.Query("str2")))
		saveQuery(c.Request.RequestURI)
	}
}

func saveQuery(query string) (int64, error) {
	db := setupDB()
	result, err := db.Exec("INSERT INTO queries(query) VALUES($1);", query)
	checkErr(err)

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insert query: %v", err)
	}

	return id, nil
}

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

func queriesListHandler(c *gin.Context) {
	db := setupDB()

	var query string
	var total int
	row := db.QueryRowContext(c, "SELECT query, COUNT( query ) total FROM queries GROUP BY query ORDER BY total DESC LIMIT 1")
	switch err := row.Scan(&query, &total); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		c.String(http.StatusOK, buildMostPopularQuery(query, strconv.Itoa(total)))
	default:
		panic(err)
	}
}

func main() {

	router := gin.Default()

	router.GET("/fizzbuzz", fizzBuzzHandler)
	router.GET("/queries", queriesListHandler)

	router.Run()
}
