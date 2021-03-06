package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusduboli/searchzin/index"
)

func searchTerm(c *gin.Context) {
	field := c.Query("field")
	term := c.Query("term")
	if term == "" || field == "" {
		c.JSON(400, gin.H{"status": "Missing query parameters"})
	}
	ids := index.GetIndexTerm(field, term)
	documents, err := index.GetDocuments(ids)
	if err != nil {
		c.JSON(500, err)
	}
	if len(documents) != 0 {
		c.JSON(200, documents)
	} else {
		c.JSON(404, []string{})
	}
}
