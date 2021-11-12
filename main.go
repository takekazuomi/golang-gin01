package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func requestDump(c *gin.Context) {

	requestDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(requestDump))

	s := ""
	if c.Request.Body != nil {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		s = string(body)
	}
	c.JSON(http.StatusOK, gin.H{"body": s})
}

func main() {
	router := gin.Default()

	router.GET("/dump", requestDump)

	router.Run(":8088")
}
