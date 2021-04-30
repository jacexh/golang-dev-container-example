package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacexh/golang-dev-container-example/api/dto"
	"github.com/jacexh/golang-dev-container-example/internal/application"
	"github.com/jacexh/golang-dev-container-example/internal/trace"
)

func CreateUser(c *gin.Context) {
	user := new(dto.UserDTO)
	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	_ = application.User.CreateUser(trace.GenContextWithRequestIndex(c), user)
	c.JSON(http.StatusOK, nil)
}
