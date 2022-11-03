package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"root/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func saveQuery(query string) (int64, error) {
	db, err := db.SetupDB()
	if err != nil {
		return 0, fmt.Errorf("db error: %w", err)
	}

	result, err := db.Exec("INSERT INTO queries(query) VALUES($1);", query)
	if err != nil {
		return 0, fmt.Errorf("insert query: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insert query: %w", err)
	}

	return id, nil
}

func fizzBuzz(a int, b int, limit int, str1 string, str2 string) []string {
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

// FizzBuzzHandler transforms an initial string and responds with a JSON array.
// QueriesHandler	godoc
// @Summary      FizzBuzz a string and get a JSON array !
// @Description  Take 5 parameters : int 1, int 2, limit, str1, str2. It plays with it then responds with a JSON array of the modified string.
// @Tags         fizzbuzz
// @Produce      json
// @Success      200
// @Router       /fizzbuzz [get]
func FizzBuzzHandler(c *gin.Context) {
	if c.Query("int1") == "" ||
		c.Query("int2") == "" || c.Query("limit") == "" ||
		c.Query("str1") == "" || c.Query("str2") == "" {
		c.String(http.StatusBadRequest, "Bad request : Missing parameter(s).")
		return
	}
	int1, err1 := strconv.Atoi(c.Query("int1"))
	if err1 != nil || int1 <= 0 {
		c.String(http.StatusBadRequest, "Bad request : int1, int2 and limit should be valid numbers and different from zero.")
		return
	}
	int2, err2 := strconv.Atoi(c.Query("int2"))
	if err2 != nil || int2 <= 0 {
		c.String(http.StatusBadRequest, "Bad request : int1, int2 and limit should be valid numbers and different from zero.")
		return
	}
	limit, err3 := strconv.Atoi(c.Query("limit"))
	if err3 != nil || limit <= 0 || limit > 1000000 {
		c.String(http.StatusBadRequest, "Bad request : int1, int2 and limit should be valid numbers and different from zero.")
		return
	}

	u, err := url.Parse(c.Request.RequestURI)
	if err != nil {
		return
	}
	saveQuery(u.RequestURI())
	c.JSON(http.StatusOK, fizzBuzz(int1, int2, limit, c.Query("str1"), c.Query("str2")))
}
