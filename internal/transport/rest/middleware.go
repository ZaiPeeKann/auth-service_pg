package httphandler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) AuthMiddleware(c *gin.Context) {
	header := strings.Split(c.GetHeader("Authirization"), " ")
	if (len(header) != 2) || (header[0] != "Bearer") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	id, err := h.services.Authorization.ParseToken(header[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	}

	c.Set("UserId", id)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("UserId")
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return 0, errors.New("UserId not found")
	}

	intId, ok := id.(int)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return 0, errors.New("UserId not found")
	}
	return intId, nil
}
