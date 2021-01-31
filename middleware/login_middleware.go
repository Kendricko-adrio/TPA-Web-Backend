package middleware

import (
	"context"
	"fmt"
	"github.com/kendricko-adrio/gqlgen-todos/graph/jwt"
	"net/http"
	"github.com/kendricko-adrio/gqlgen-todos/graph/model"

)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

var userWriter = &writerResponse{"writter"}

type writerResponse struct {

	name string
}

//var userCookie = &userCookieCtx{"cookie"}
//
//type userCookieCtx struct {
//
//	name string
//}

func LoginMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//header := r.Header.Get("Authorization")
			//header := r.Header.

			header, err := r.Cookie("staem")
			// Allow unauthenticated users in
			//if header == ""{
			//	next.ServeHTTP(w, r)
			//	fmt.Println("masuk non auth")
			//	return
			//}
			write := context.WithValue(r.Context(), userWriter ,&w)
			r = r.WithContext(write)

			//httpreq := context.WithValue(r.Context(), userCookie, &r)
			//r = r.WithContext(httpreq)
			if err != nil {

				next.ServeHTTP(w, r)
				fmt.Println("masuk non auth")
				return
			}

			//validate jwt token
			tokenStr := header
			//username, err := jwt.ParseToken(tokenStr)
			username, err := jwt.ParseToken(tokenStr.Value)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user := model.User{
				UserName: username,
			}
			id, err := model.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = id
			// put it in context

			ctx := context.WithValue(r.Context(), userCtxKey, &user)
			//write := context.WithValue(r.Context(), userWriter ,&w)
			// and call the next with our new context
			fmt.Println("masuk auth")
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}

func ForWriter(ctx context.Context) *http.ResponseWriter {
	raw, _ := ctx.Value(userWriter).(*http.ResponseWriter)
	return raw
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
//func ForRequest(ctx context.Context) *http.Request {
//	raw, _ := ctx.Value(userCookie).(*http.Request)
//	return raw
//}

