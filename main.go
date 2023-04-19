package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func getFileTxt(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	return string(content), err
}

func main() {
	router := gin.Default()

	DATAFILE := os.Getenv("DATAFILE")
	if len(DATAFILE) == 0 {
		DATAFILE = "data/"
	}

	api := router.Group("/api")
	{

		api.GET("/data", func(c *gin.Context) {
			content, err := os.ReadFile(DATAFILE + "data.json")
			if err != nil {
				log.Println(err)
			}
			c.Writer.WriteString(string(content))
		})

		api.GET("/config", func(c *gin.Context) {
			content, err := os.ReadFile(DATAFILE + "config.json")
			if err != nil {
				log.Println(err)
			}
			c.Writer.WriteString(string(content))
		})

		api.PUT("/key", func(c *gin.Context) {
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

		api.PUT("/data", func(c *gin.Context) {
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

		api.PUT("/config", func(c *gin.Context) {
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
