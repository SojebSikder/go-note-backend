package routes

import (
	"github.com/SojebSikder/goframe/app/controller/example"
	"github.com/SojebSikder/goframe/app/controller/note"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", example.Index)
	router.GET("/test", example.Hello)

	// note
	router.GET("/note", note.Index)
	router.POST("/note", note.Create)
	router.PUT("/note", note.Update)
	router.DELETE("/note/:id", note.Delete)
}
