package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/tumeraltunbas/go-blog/constants/errors"
)

func DtoValidation[T any](next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var dto T

			if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
				errors.NewInternalServerError(w)
				return
			}

			validate := validator.New()

			if err := validate.Struct(dto); err != nil {
				errors.NewValidationError(w, err)
				return
			}

			ctx := context.WithValue(r.Context(), "dto", dto)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
