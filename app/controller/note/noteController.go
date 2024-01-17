package note

import (
	"net/http"
	"strconv"

	"github.com/SojebSikder/goframe/app/model/note"
	orm "github.com/SojebSikder/goframe/system/core/ORM"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var result any
	result, _ = note.FindAll(orm.Ctx, orm.Client)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func Create(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	note.Create(orm.Ctx, orm.Client, title, content)

	var result any
	result, _ = note.FindAll(orm.Ctx, orm.Client)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func Update(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	note.Update(orm.Ctx, orm.Client, title, content)

	var result any
	result, _ = note.FindAll(orm.Ctx, orm.Client)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	note.Delete(orm.Ctx, orm.Client, id)

	var result any
	result, _ = note.FindAll(orm.Ctx, orm.Client)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}
