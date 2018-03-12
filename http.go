package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

func ginEngine() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Use(location.Default())
	r.GET("/:urlKey", func(c *gin.Context) {
		urlKey := c.Param("urlKey")
		entry, err := getMinifiedURL(urlKey)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if entry == nil {
			c.String(http.StatusNotFound, "URL not found")
			return
		}
		c.Redirect(301, entry.URL)
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	r.POST("/", func(c *gin.Context) {
		url := c.PostForm("url")
		loc := location.Get(c)
		cleanURL, err := validateURL(url)
		if err != nil {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"url":             url,
				"validationError": err,
			})
			return
		}
		entry, err := createMinifiedURL(cleanURL)
		if err != nil {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"url":             url,
				"validationError": err,
			})
			return
		}
		c.HTML(http.StatusCreated, "index.tmpl", gin.H{
			"entry": entry,
			"root":  fmt.Sprintf("%s://%s", loc.Scheme, loc.Host),
		})
	})
	return r
}
