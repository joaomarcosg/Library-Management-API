package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcosg/Library-Management-API/internal/configuration/resterror"
	"github.com/joaomarcosg/Library-Management-API/internal/controller/model/request"
)

func SignupUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := resterror.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error = %s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

}
