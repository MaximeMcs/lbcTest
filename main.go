package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func checkAtoiErr(err error) {
	if err != nil {
		fmt.Println("error converting string to int")
		os.Exit(1)
	}
}

// strBuilt := strings.Builder{}
// if i%a == 0 {
// 	strBuilt.WriteString(str1)
// }
// if i%b == 0 {
// 	strBuilt.WriteString(str2)
// }
// if i%a == 0 || i%b == 0 {
// 	strSlice = append(strSlice, strBuilt.String())
// } else {
// 	strSlice = append(strSlice, strconv.Itoa(i))
// }
// strBuilt.Reset()

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

	int1, err := strconv.Atoi(c.Query("int1"))
	checkAtoiErr(err)
	int2, err := strconv.Atoi(c.Query("int2"))
	checkAtoiErr(err)
	limit, err := strconv.Atoi(c.Query("limit"))
	checkAtoiErr(err)

	str1 := c.Query("str1")
	str2 := c.Query("str2")

	c.JSON(http.StatusOK, fizzBuzzIt(int1, int2, limit, str1, str2))
}

func main() {
	router := gin.Default()

	router.GET("/fizzbuzz", fizzBuzzHandler)
	router.Run()
}
