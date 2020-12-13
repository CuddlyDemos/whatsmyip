package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting app...")
	r := gin.Default()
	r.GET("/", index)
	r.GET("/ping", ping)
	r.Run()
}

func index(c *gin.Context) {
	ips := strings.Split(c.Request.Header.Get("X-FORWARDED-FOR"), ",")
	ip := ips[len(ips)-1]
	if ip == "" {
		ip = c.ClientIP()
	}
	c.JSON(http.StatusOK, gin.H{"ip": ip})
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
