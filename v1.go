package server

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"unsafe"

	"github.com/flyaways/golang-lru/simplelru"
	"github.com/gin-gonic/gin"
)

//Meta data
func Meta(v *gin.RouterGroup, cache simplelru.LRUCache) {
	//Keys
	v.GET("/keys", func(c *gin.Context) {
		c.JSON(http.StatusOK, cache.Keys())
	})

	//Len
	v.GET("/len", func(c *gin.Context) {
		c.JSON(http.StatusOK, cache.Len())
	})

	//Purge
	v.DELETE("/purge", func(c *gin.Context) {
		cache.Purge()
	})

	//GetOldest
	v.GET("/getoldest", func(c *gin.Context) {
		cache.GetOldest()
	})

	//RemoveOldest
	v.DELETE("/removeoldest", func(c *gin.Context) {
		cache.RemoveOldest()
	})

	//Contains
	v.GET("/contains", func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, cache.Contains(key))
	})

	//Peek
	v.GET("/peek", func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		val, ok := cache.Peek(key)
		if !ok {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, val)
	})

	//HTTP 0.9 GET
	//HTTP 1.0 GET HEAD POST
	//HTTP 1.1 GET HEAD POST OPTIONS PUT DELETE TRACE CONNECT
	v.OPTIONS("/:key", func(c *gin.Context) {
		c.Header("Allow", "OPTIONS")
		c.Header("Allow", "GET")
		c.Header("Allow", "DELETE")
	})
}

//Object ops
func Object(v *gin.RouterGroup, cache simplelru.LRUCache) {
	//Get
	v.GET("/:key", func(c *gin.Context) {
		key := c.Param("key")
		if key == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		val, ok := cache.Get(key)
		if !ok {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, val)
	})

	//Head
	v.HEAD("/:key", func(c *gin.Context) {
		key := c.Param("key")
		if key == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		val, ok := cache.Get(key)
		if !ok {
			c.Status(http.StatusNotFound)
			return
		}

		c.Header("Content-Length", strconv.Itoa(int(unsafe.Sizeof(val))))
	})

	//Add
	v.PUT("/:key", func(c *gin.Context) {
		key := c.Param("key")
		if key == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		if !cache.Add(key, data) {
			c.Status(http.StatusInternalServerError)
		}
	})

	//Add
	v.POST("/:key", func(c *gin.Context) {
		key := c.Param("key")
		if key == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		if !cache.Add(key, data) {
			c.Status(http.StatusInternalServerError)
		}
	})

	//Remove
	v.DELETE("/:key", func(c *gin.Context) {
		key := c.Param("key")
		if key == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		if !cache.Remove(key) {
			c.Status(http.StatusInternalServerError)
		}
	})

	//HTTP 0.9 GET
	//HTTP 1.0 GET HEAD POST
	//HTTP 1.1 GET HEAD POST OPTIONS PUT DELETE TRACE CONNECT
	v.OPTIONS("/:key", func(c *gin.Context) {
		c.Header("Allow", "OPTIONS")
		c.Header("Allow", "GET")
		c.Header("Allow", "POST")
		c.Header("Allow", "PUT")
		c.Header("Allow", "DELETE")
		c.Header("Allow", "HEAD")
	})
}

//Version1 api version 1
func Version1(v *gin.RouterGroup, cache simplelru.LRUCache) {
	Meta(v.Group("/meta"), cache)
	Object(v.Group("/object"), cache)
}
