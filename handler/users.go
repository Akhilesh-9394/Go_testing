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
	users, err := model.GetallUsers()
	if err != nil {
		c.String(http.StatusNotFound, "%s", err)
		return

	}
	user, _, err := model.FindUserById(id, users)
	if err != nil {
		c.String(http.StatusNotFound, "%s", err)
		return

	}
	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {

	users, err := model.GetallUsers()
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", err)
		return

	}
	c.JSON(200, users)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.User
	c.Bind(&user)
	user.Id = id
	users, err := model.GetallUsers()
	if err != nil {
		c.String(http.StatusNotFound, "%s", err)
		return

	}
	_, userId, err := model.FindUserById(id, users)
	if err != nil {
		c.String(http.StatusNotFound, "%s", err)
		return

	}
	users = append(users[:userId-1], users[userId:]...)
	users = append(users, user)
	c.JSON(200, users)
}
