package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/leafney/meiju-web/models"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

func Home_Index(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	//查询所有数据
	persons := []models.Person{}
	db.C(models.CollectionPerson).Find(nil).All(&persons)

	c.JSON(200, gin.H{
		"persons": persons,
	})
}
