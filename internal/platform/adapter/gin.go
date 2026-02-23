package adapter

import (
	"net/http"

	"github.com/GuidoGdR/go-speed-test/internal/platform/errorBody"
	"github.com/gin-gonic/gin"
)

func Gin(handler Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, errorBody.BadRequestFormat())
		}

		req := newRequest(c.Request.Method, body, nil)

		resp, err := handler(c, req)
		if err != nil {

			for k, v := range resp.Headers {
				c.Header(k, v)
			}
			c.JSON(resp.StatusCode, resp.Body)
			return
		}

		for k, v := range resp.Headers {
			c.Header(k, v)
		}
		c.JSON(resp.StatusCode, resp.Body)
	}
}
