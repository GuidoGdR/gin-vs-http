package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GuidoGdR/go-speed-test/internal/platform/adapter"
	"github.com/GuidoGdR/go-speed-test/internal/platform/appErr"
	"github.com/go-playground/validator/v10"
)

type AdapterHandler struct {
	service  *Service
	validate *validator.Validate
}

func NewAdapterHandler(s *Service, validate *validator.Validate) *AdapterHandler {
	return &AdapterHandler{
		service:  s,
		validate: validate,
	}
}

func (h *AdapterHandler) Login(ctx context.Context, req *adapter.Request) (*adapter.Response, error) {
	resp := new(adapter.Response)

	if req.Method != http.MethodPost {
		return resp.MethodNotAllowed("Only POST allowed"), appErr.MethodNotAllowed
	}

	var data *loginRequest

	if err := json.Unmarshal(req.Body, &data); err != nil {
		return resp.BadRequestFormat(), err
	}

	if err := h.validate.Struct(data); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {

			return resp.BadRequest("Bad username or password"), err
		}

		return resp.InternalServerError(), err
	}

	logginData, err := h.service.Login(ctx, data.Username, data.Password)
	if err != nil {

		if errors.Is(err, appErr.Unauthorized) {

			return resp.BadRequest("Bad username or password"), err
		}

		return resp.InternalServerError(), err
	}

	body := loginResponse{
		Access:    logginData.AccessTkn,
		Refresh:   logginData.RefreshTkn,
		TokenType: logginData.TokenType,
		User:      logginData.User,
	}

	return resp.Created(body), nil
}

func (h *AdapterHandler) Refresh(ctx context.Context, req *adapter.Request) (*adapter.Response, error) {
	resp := new(adapter.Response)

	if req.Method != http.MethodPost {
		return resp.MethodNotAllowed("Only POST allowed"), errors.New("Method NOT allowed")
	}

	var data *refreshRequest

	if err := json.Unmarshal(req.Body, &data); err != nil {
		return resp.BadRequestFormat(), err
	}

	tkns, err := h.service.Refresh(ctx, data.Refresh)
	if err != nil {

		if errors.Is(err, appErr.Unauthorized) {
			return resp.BadRequest("Bad username or password"), err
		}

		return resp.InternalServerError(), err
	}

	body := refreshResponse{
		Access:  tkns.AccessTkn,
		Refresh: tkns.RefreshTkn,
	}

	return resp.OK(body), nil
}

func (h *AdapterHandler) Register(ctx context.Context, req *adapter.Request) (*adapter.Response, error) {
	resp := new(adapter.Response)

	if req.Method != http.MethodPost {
		return resp.MethodNotAllowed("Only POST allowed"), appErr.MethodNotAllowed
	}

	var data *registerRequest

	if err := json.Unmarshal(req.Body, &data); err != nil {
		return resp.BadRequestFormat(), err
	}

	usr, err := h.service.Register(ctx, data.Username, data.Password, data.Email)
	if err != nil {

		return resp.InternalServerError(), err
	}

	return resp.Created(usr), nil
}
