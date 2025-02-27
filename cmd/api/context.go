package main

import (
	"context"
	"greenlight.kakani21.net/internal/data_1"
	"net/http"
)

type contextKey string

const userContextKey = contextKey("user")

func (app *application) contextSetUser(r *http.Request, user *data_1.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *data_1.User {
	user, ok := r.Context().Value(userContextKey).(*data_1.User)

	if !ok {
		panic("missing user value in request context")
	}

	return user
}
