package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckAtoiErr(err error, c *gin.Context) {
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request : int1, int2 and limit must be numbers.")
	}
}
