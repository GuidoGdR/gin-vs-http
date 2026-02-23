package adapter

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GuidoGdR/go-speed-test/internal/platform/errorBody"
)

func HTTP(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer r.Body.Close()
		}
		w.Header().Set("Content-Type", "application/json")

		body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20)) // Max 1MB
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorBody.BadRequestFormat())
		}

		ctx := r.Context()

		req := newRequest(r.Method, body, nil)

		resp, err := handler(ctx, req)
		if err != nil {
			// any
		}

		for k, v := range resp.Headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(resp.StatusCode)
		json.NewEncoder(w).Encode(resp.Body)
	}
}
