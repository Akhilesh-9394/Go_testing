package handler

import (
	"net/http"
	"strconv"

	"github.com/fullstacker-go/practice_gin/model"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)

}

func CreateUser(c *gin.Context) {
	var newuser model.User
	c.Bind(&newuser)
	users, err := model.GetallUsers()
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err)
		return

	}
	user := model.NextID(&newuser, users)
	c.JSON(http.StatusOK, user)
}

func GetAUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := model.FindUserById(id)
	if err != nil {
		c.String(http.StatusNotFound, "%s", err)
		return

	}
	c.JSON(http.StatusOK, user)
}

func GetAll(c *gin.Context) {

	users, err := model.GetallUsers()
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err)
		return

	}
	c.JSON(200, users)
}
