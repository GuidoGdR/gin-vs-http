package auth

import (
	"errors"
	"net/http"

	"github.com/GuidoGdR/go-speed-test/internal/platform/appErr"
	"github.com/GuidoGdR/go-speed-test/internal/platform/errorBody"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ginHandler struct {
	service *Service
}

func NewGinHandler(s *Service) *ginHandler {
	return &ginHandler{
		service: s,
	}
}
func (h *ginHandler) Login(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, errorBody.MethodNotAllowed("Only POST allowed"))
		return
	}

	var data *loginRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusUnauthorized, errorBody.Unauthorized())
			return
		}

		c.JSON(http.StatusBadRequest, errorBody.BadRequestFormat())
		return
	}

	logginData, err := h.service.Login(c.Request.Context(), data.Username, data.Password)
	if err != nil {

		if errors.Is(err, appErr.Unauthorized) {

			c.JSON(http.StatusUnauthorized, errorBody.Unauthorized())
			return
		}

		c.JSON(http.StatusInternalServerError, errorBody.InternalServerError())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"access":     logginData.AccessTkn,
		"refresh":    logginData.RefreshTkn,
		"token_type": logginData.TokenType,
		"user":       logginData.User,
	})
}
func (h *ginHandler) Refresh(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, errorBody.MethodNotAllowed("Only POST allowed"))
		return
	}

	var data *refreshRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusUnauthorized, errorBody.Unauthorized())
			return
		}

		c.JSON(http.StatusBadRequest, errorBody.BadRequestFormat())
		return
	}

	tkns, err := h.service.Refresh(c, data.Refresh)
	if err != nil {

		if errors.Is(err, appErr.Unauthorized) {

			c.JSON(http.StatusUnauthorized, errorBody.Unauthorized())
			return
		}

		c.JSON(http.StatusInternalServerError, errorBody.InternalServerError())
		return
	}

	c.JSON(http.StatusOK, gin.H{"access": tkns.AccessTkn, "refresh": tkns.RefreshTkn})
}
func (h *ginHandler) Register(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, errorBody.MethodNotAllowed("Only POST allowed"))
		return
	}

	var data *registerRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, errorBody.BadRequestValidationErrors(err))
			return
		}

		c.JSON(http.StatusBadRequest, errorBody.BadRequestFormat())
		return
	}

	usr, err := h.service.Register(c, data.Username, data.Password, data.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorBody.InternalServerError())
		return
	}

	c.JSON(http.StatusCreated, usr)
}
