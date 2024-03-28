package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/satorunooshie/example-error-handling/apperr"
)

func respond(_ context.Context, w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if code == http.StatusNoContent {
		return
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return
	}
}

func respondError(_ context.Context, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		code := httpStatusCode(err)
		log.Printf("status_code: %d, error_code: %d, err: %v, error_message: %q,  description: %q", code, apperr.GetSystemCode(err), err, apperr.GetErrorMessage(err), apperr.GetDescription(err))
		w.WriteHeader(code)
		_, _ = fmt.Fprintf(w, `{"erorr_code":%d,"error_message":%q}`, apperr.GetSystemCode(err), apperr.GetErrorMessage(err))
	}
}

func httpStatusCode(err error) int {
	switch err.(type) {
	case *apperr.ClientError:
		return http.StatusBadRequest
	case *apperr.InternalError:
		return http.StatusInternalServerError
	case *apperr.NotFoundError:
		return http.StatusNotFound
	case *apperr.CanceledError:
		return http.StatusConflict
	case *apperr.ForbiddenError:
		return http.StatusForbidden
	case *apperr.UnauthorizedError:
		return http.StatusUnauthorized
	case *apperr.DeadlineExceededError:
		return http.StatusGatewayTimeout
	case *apperr.RateLimitExceededError:
		return http.StatusTooManyRequests
	case *apperr.PreconditionFailedError:
		return http.StatusPreconditionFailed
	default:
		return http.StatusInternalServerError
	}
}
