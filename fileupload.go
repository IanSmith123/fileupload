package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	addr := flag.String("addr", "0.0.0.0", "listen addr")
	port := flag.Int("port", 2019, "listen port")
	flag.Parse()

	r := gin.Default()

	r.GET("/", rootGet)
	r.POST("/", rootPost)

	r.Use(static.Serve("/", static.LocalFile("./", true)))

	listenAddr := fmt.Sprintf("%s:%d", *addr, *port)
	fmt.Println("listen on: " + listenAddr)

	err := r.Run(listenAddr)

	if err != nil {
		panic(err)
	}
}

func rootGet(c *gin.Context) {
	//remoteUser := c.Request.RemoteAddr + "\t" + c.Request.UserAgent() + c.Request.Header.Get("Xr")
	remoteUser := c.Request.RemoteAddr + "\t" + c.Request.UserAgent() + "\t" + c.Request.Header.Get("X-Forwarded-For")
	log.Println(remoteUser)

	files, _ := filepath.Glob("*")
	s := "<pre>\n"
	for _, file := range files {
		s += fmt.Sprintf("<a href=%s>%s</a>\n", file, file)
	}
	s += "</pre>"

	help := "upload file: curl -F file=@a.png " + c.Request.Host
	resp := s + "\n\n" + help + "\n"

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(resp))

}

func rootPost(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}

	err = c.SaveUploadedFile(file, file.Filename)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	msg := fmt.Sprintf("%s uploaded %s.", c.Request.RemoteAddr, file.Filename)
	fmt.Println(msg)

	c.String(http.StatusOK, msg)
}
