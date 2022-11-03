package api

import (
	"fmt"
	"net/http"
	"root/db"
	"root/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func saveQuery(query string) (int64, error) {
	db := db.SetupDB()
	result, err := db.Exec("INSERT INTO queries(query) VALUES($1);", query)
	utils.CheckErr(err)

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insert query: %v", err)
	}

	return id, nil
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

func FizzBuzzHandler(c *gin.Context) {
	if c.Query("int1") == "" ||
		c.Query("int2") == "" || c.Query("limit") == "" ||
		c.Query("str1") == "" || c.Query("str2") == "" {
		c.String(http.StatusBadRequest, "Bad request : Missing parameter(s).")
	} else {
		int1, err := strconv.Atoi(c.Query("int1"))
		utils.CheckAtoiErr(err, c)
		int2, err := strconv.Atoi(c.Query("int2"))
		utils.CheckAtoiErr(err, c)
		limit, err := strconv.Atoi(c.Query("limit"))
		utils.CheckAtoiErr(err, c)

		if err == nil {
			c.JSON(http.StatusOK, fizzBuzzIt(int1, int2, limit, c.Query("str1"), c.Query("str2")))
			saveQuery(c.Request.RequestURI)
		}
	}
}
