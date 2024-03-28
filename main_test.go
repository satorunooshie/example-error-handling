package main

import (
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/satorunooshie/e2e"
)

func name(code int, description ...string) string {
	return strings.Join(append([]string{strconv.Itoa(code)}, description...), "_")
}

func TestMain(t *testing.T) {
	e2e.RegisterRouter(router())
}

func TestUserHandler_Get(t *testing.T) {
	endpoint := "/user/"
	tests := []struct {
		description []string
		id          string
		status      int
	}{
		{[]string{"user not found", "invalid id"}, "1", http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(name(tt.status, tt.description...), func(t *testing.T) {
			r := e2e.NewRequest("GET", endpoint+tt.id, nil)
			e2e.RunTest(t, r, tt.status, e2e.PrettyJSON)
		})
	}
}

func TestUserHandler_Create(t *testing.T) {
	endpoint := "/user"
	tests := []struct {
		description []string
		body        map[string]any
		status      int
	}{
		{[]string{"validation", "missing header"}, nil, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(name(tt.status, tt.description...), func(t *testing.T) {
			r := e2e.NewRequest("POST", endpoint, e2e.JSONBody(t, tt.body))
			e2e.RunTest(t, r, tt.status, e2e.PrettyJSON)
		})
	}
}

func TestUserHandler_Delete(t *testing.T) {
	endpoint := "/user/"
	tests := []struct {
		description []string
		id          string
		status      int
	}{
		{[]string{"unexpected error"}, "1", http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(name(tt.status, tt.description...), func(t *testing.T) {
			r := e2e.NewRequest("DELETE", endpoint+tt.id, nil)
			e2e.RunTest(t, r, tt.status, e2e.PrettyJSON)
		})
	}
}
