package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func getFileTxt(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	return string(content), err
}

var (
	VERSION    string
	BUILD_TIME string
	GO_VERSION string
)

func main() {
	ginModel := os.Getenv("GIN_MODE")
	fmt.Printf("%s\n%s\n%s\n%s\n", VERSION, BUILD_TIME, GO_VERSION, ginModel)
	router := gin.Default()

	DATAFILE := os.Getenv("DATAFILE")
	if len(DATAFILE) == 0 {
		DATAFILE = "data/"
	}

	v1 := router.Group("/v1")
	{

		v1.GET("/data", func(c *gin.Context) {
			content, err := os.ReadFile(DATAFILE + "data.json")
			if err != nil {
				log.Println(err)
			}
			c.Writer.WriteString(string(content))
		})

		v1.GET("/config", func(c *gin.Context) {
			content, err := os.ReadFile(DATAFILE + "config.json")
			if err != nil {
				log.Println(err)
			}
			c.Writer.WriteString(string(content))
		})

		v1.PUT("/key", func(c *gin.Context) {
			oldkey := c.PostForm("oldkey")
			newkey := c.PostForm("newkey")
			content, err := getFileTxt(DATAFILE + "key.txt")
			if string(content) != oldkey {
				log.Println(err)
				c.JSON(401, gin.H{
					"message": "no auth",
				})
				return
			}

			err = ioutil.WriteFile(DATAFILE+"key.txt", []byte(newkey), 0666)
			if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{
					"message": "write key error",
				})
				return
			}
			c.JSON(200, gin.H{
				"message": "ok",
			})
		})

		v1.PUT("/data", func(c *gin.Context) {
			key := c.PostForm("key")
			content, err := getFileTxt(DATAFILE + "key.txt")
			if string(content) != key {
				log.Println(err)
				c.JSON(401, gin.H{
					"message": "no auth",
				})
				return
			}
			data := c.PostForm("data")
			err = ioutil.WriteFile(DATAFILE+"data.json", []byte(data), 0666)
			if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{
					"message": "write data error",
				})
				return
			}
			c.Writer.WriteString(string(content))
		})

		v1.PUT("/config", func(c *gin.Context) {
			key := c.PostForm("key")
			content, err := getFileTxt(DATAFILE + "key.txt")
			if string(content) != key {
				log.Println(err)
				c.JSON(401, gin.H{
					"message": "no auth",
				})
				return
			}
			data := c.PostForm("config")
			err = ioutil.WriteFile(DATAFILE+"config.json", []byte(data), 0666)
			if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{
					"message": "write config error",
				})
				return
			}
			c.Writer.WriteString(string(content))
		})
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
