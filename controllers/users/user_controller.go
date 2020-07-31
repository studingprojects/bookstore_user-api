package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/studingprojects/bookstore_oauth-go/oauth"
	"github.com/studingprojects/bookstore_user-api/domain/users"
	"github.com/studingprojects/bookstore_user-api/services"
	"github.com/studingprojects/bookstore_user-api/utils/errors"
)

func TestServiceInterface() {

}

func GetUser(c *gin.Context) {
	if authErr := oauth.AuthenticateRequest(c.Request); authErr != nil {
		c.JSON(authErr.Status, authErr)
		return
	}

	userId, userErr := strconv.ParseInt(c.Param("userId"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	if result.Id == oauth.GetCallerId(c.Request) {
		c.JSON(http.StatusOK, result.Marshal(false))
		return
	}
	c.JSON(http.StatusOK, result.Marshal(oauth.IsPublic(c.Request)))
	return
}

func GetUsers(c *gin.Context) {

}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("userId"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}
	isPartial := c.Request.Method == http.MethodPatch
	user.Id = userID
	result, getErr := services.UsersService.UpdateUser(isPartial, user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("userId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	if deletedErr := services.UsersService.DeleteUser(userId); deletedErr != nil {
		c.JSON(deletedErr.Status, deletedErr)
		return
	}
	c.JSON(http.StatusNoContent, map[string]string{"status": "deleted"})
	return
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.FindByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var login users.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		restErr := errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status, restErr)
		return
	}
	var userInfo *users.User
	var restErr *errors.RestErr
	if userInfo, restErr = services.UsersService.Login(login.Email, login.Password); restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	fmt.Println(userInfo)
	c.JSON(http.StatusOK, userInfo)
	return
}

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}
	return userId, nil
}
