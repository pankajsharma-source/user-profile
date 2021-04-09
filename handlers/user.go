package handlers

import (
	"context"
	"fmt"
	"github.com/pankajsharma-source/user-profile/data"
	"log"
	"net/http"
)

// Products is a http.Handler
type User struct {
	l *log.Logger
}

type KeyUser struct{}

// NewProducts creates a products handler with the given logger
func NewUser(l *log.Logger) *User {
	return &User{l}
}

// getProducts returns the products from the data store
func (u *User) GetUser(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle GET User")
}

func (u *User) AddUser(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle POST User")
	user := r.Context().Value(KeyUser{}).(data.User)
	data.AddUser(&user)
}

func (u User) UpdateUser(rw http.ResponseWriter, r *http.Request) {

}

func (u User) MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := data.User{}

		err := user.FromJSON(r.Body)
		if err != nil {
			u.l.Println("[ERROR] deserializing user", err)
			http.Error(rw, "Error reading user", http.StatusBadRequest)
			return
		}

		// validate the product
		err = user.Validate()
		if err != nil {
			u.l.Println("[ERROR] validating user", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating user: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)

	})
}