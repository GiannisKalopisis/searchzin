package http

import "github.com/gin-gonic/gin"

func Api(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")

	r.GET("/", index)

	v1 := r.Group("v1")
	v1.GET("/ping", ping)
	v1.POST("/documents", documentSave)
	v1.GET("/documents", documentList)
	v1.GET("/indexes", indexList)
}
