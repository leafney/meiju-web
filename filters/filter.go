package filters

import (
	"github.com/Leafney/meiju-web/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Connect(c *gin.Context) {
	s := db.Session.Clone()
	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	// TODO: Handle it in a better way
	if len(c.Errors) > 0 {
		c.HTML(http.StatusBadRequest, "400.html", gin.H{
			"errors": c.Errors,
		})
	}
}
