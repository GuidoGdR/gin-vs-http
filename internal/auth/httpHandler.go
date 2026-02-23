package auth

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/GuidoGdR/go-speed-test/internal/platform/appErr"
	"github.com/GuidoGdR/go-speed-test/internal/platform/errorBody"
	"github.com/go-playground/validator/v10"
)

type hTTPHandler struct {
	service  *Service
	validate *validator.Validate
}

func NewHTTPHandler(s *Service, validate *validator.Validate) *hTTPHandler {
	return &hTTPHandler{
		service:  s,
		validate: validate,
	}
}

func (h *hTTPHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	if r.Method != http.MethodPost {

		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorBody.MethodNotAllowed("Only POST allowed"))
		return
	}

	var data *loginRequest

	if err := json.NewDecoder(io.LimitReader(r.Body, 1<<20)).Decode(&data); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorBody.BadRequestFormat())
		return
	}

	if err := h.validate.Struct(data); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {

			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorBody.Unauthorized())
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorBody.InternalServerError())
		return
	}

	logginData, err := h.service.Login(ctx, data.Username, data.Password)
	if err != nil {

		if errors.Is(err, appErr.Unauthorized) {

			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorBody.Unauthorized())
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorBody.InternalServerError())
		return
	}

	body := loginResponse{
		Access:    logginData.AccessTkn,
		Refresh:   logginData.RefreshTkn,
		TokenType: logginData.TokenType,
		User:      logginData.User,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)
}
func (h *hTTPHandler) Refresh(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	if r.Method != http.MethodPost {

		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorBody.MethodNotAllowed("Only POST allowed"))
		return
	}

	var data *refreshRequest

	if err := json.NewDecoder(io.LimitReader(r.Body, 1<<20)).Decode(&data); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorBody.BadRequestFormat())
		return
	}

	if err := h.validate.Struct(data); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {

			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorBody.Unauthorized())
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorBody.InternalServerError())
		return
	}

	tkns, err := h.service.Refresh(ctx, data.Refresh)
	if err != nil {

		if errors.Is(err, appErr.Unauthorized) {

			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorBody.Unauthorized())
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorBody.InternalServerError())
		return
	}

	body := refreshResponse{
		Access:  tkns.AccessTkn,
		Refresh: tkns.RefreshTkn,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}
func (h *hTTPHandler) Register(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	if r.Method != http.MethodPost {

		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorBody.MethodNotAllowed("Only POST allowed"))
		return
	}

	var data *registerRequest

	if err := json.NewDecoder(io.LimitReader(r.Body, 1<<20)).Decode(&data); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorBody.BadRequestFormat())
		return
	}

	if err := h.validate.Struct(data); err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorBody.BadRequestValidationErrors(err))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorBody.InternalServerError())
		return
	}

	usr, err := h.service.Register(ctx, data.Username, data.Password, data.Email)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorBody.InternalServerError())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usr)
}
