package handler

import (
	"errors"
	"net/http"

	"github.com/satorunooshie/example-error-handling/apperr"
)

type User struct{}

func (User) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := get(r.PathValue("id"))
	if err != nil {
		respondError(ctx, w, err)
		return
	}

	respond(ctx, w, http.StatusOK, user)
}

func get(_ string) (any, error) {
	return nil, apperr.NewNotFoundError(apperr.New(apperr.UserNotFound))
}

func (User) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := validate(); err != nil {
		respondError(ctx, w, err)
		return
	}

	respond(ctx, w, http.StatusCreated, nil)
}

func validate() error {
	return apperr.NewClientError(apperr.New(
		apperr.InvalidRequest,
		apperr.WithMessage("missing required fields in request header"),
		apperr.WithDescription("suspicous request"),
	))
}

func (User) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := del(r.PathValue("id")); err != nil {
		respondError(ctx, w, err)
		return
	}

	respond(ctx, w, http.StatusNoContent, nil)
}

func del(_ string) error {
	// not wrapped error
	return errors.New("not implemented")
}
