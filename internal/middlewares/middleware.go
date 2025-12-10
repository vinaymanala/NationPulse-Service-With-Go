package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/nationpulse-bff/internal/auth"
	"github.com/nationpulse-bff/internal/utils"
)

type Middleware func(*utils.Configs, http.Handler) http.Handler
type WithAuthMiddleware func(*utils.Configs, http.Handler) http.Handler

func allowCors(configs *utils.Configs, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow-Control-Access-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		log.Default().Println("CORS middleware executed")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func logging(configs *utils.Configs, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func panicRecovery(configs *utils.Configs, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func bearerFromHandler(r *http.Request) string {
	h := r.Header.Get("Authorization")
	if strings.HasPrefix(h, "Bearer ") {
		return strings.TrimPrefix(h, "Bearer ")
	}
	return ""
}

func authMiddleware(configs *utils.Configs, next http.Handler) http.Handler {
	rd := configs.Cache
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement authentication logic here
		var token string
		if c, err := r.Cookie("access_token"); err == nil && c != nil && c.Value != "" {
			fmt.Println("--------COOKIE----------", c.Value)
			token = c.Value
		} else {
			fmt.Println("No cookie found, checking Authorization header")
			token = bearerFromHandler(r)
		}
		fmt.Println("--------TOKEN----------", token)
		if token == "" {
			log.Println(http.StatusUnauthorized, "missing token")
			return
		}

		claims, err := auth.ParseAccess(token)
		if err != nil {
			log.Println(http.StatusUnauthorized, err, "invalid token")
			return
		}

		ctx := context.Background()
		if _, err := rd.GetUserByJTI(ctx, "access:"+claims.ID); err != nil {
			log.Println(http.StatusUnauthorized, err, "invalid token jti")
			return
		}
		fmt.Println("")
		fmt.Printf("CLAIMS: %v+\n", claims)
		fmt.Println("PATH:", r.URL.Path)
		r.ParseForm()
		r.Form.Set("userID", claims.Subject)
		next.ServeHTTP(w, r)
	})
}

func MustCookie(r *http.Request, name string) (string, error) {
	val, err := r.Cookie(name)
	if err != nil || val == nil {
		return "", errors.New("missing cookie: " + name)
	}
	return val.Value, nil
}

func executeMiddlewares(configs *utils.Configs, ms []Middleware, next http.Handler) http.Handler {
	h := next
	for _, m := range ms {
		h = m(configs, h)
	}
	return h
}

func DefaultMiddlewares(configs *utils.Configs, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middlewares := []Middleware{
			allowCors,
			logging,
			panicRecovery,
		}
		h := executeMiddlewares(configs, middlewares, next)
		h.ServeHTTP(w, r)
	})
}

func WithAuthMiddlewares(configs *utils.Configs, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middlewares := []Middleware{
			DefaultMiddlewares,
		}
		middlewares = append(middlewares, func(configs *utils.Configs, next http.Handler) http.Handler {
			return authMiddleware(configs, next)
		})
		h := executeMiddlewares(configs, middlewares, next)
		h.ServeHTTP(w, r)
	})
}
