package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/nationpulse-bff/internal/auth"
	"github.com/nationpulse-bff/internal/store"
)

type UserService struct {
	Cache   *store.Redis
	Context context.Context
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var demoUser = &User{ID: "1", Name: "admin", Password: "secret"}

func (us *UserService) HandleLogin(w http.ResponseWriter, r *http.Request) {
	rds := us.Cache
	log.Println("Handling Login Route..")
	var in struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Printf("User %s attempting to log in", in.Name)
	// check in db if user exists
	if in.Name != demoUser.Name || in.Password != demoUser.Password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokens, err := auth.IssueTokens(demoUser.ID)
	if err != nil {
		http.Error(w, "Failed to issue tokens", http.StatusInternalServerError)
		return
	}
	fmt.Println("TOKENS", tokens)
	if err := auth.Persist(r.Context(), rds, tokens); err != nil {
		http.Error(w, "Failed to persist tokens", http.StatusInternalServerError)
		return
	}
	auth.SetAuthCookies(w, tokens)
	w.Write([]byte("Login successful"))
}

func (us *UserService) HandleLogout(w http.ResponseWriter, r *http.Request) {
	rds := us.Cache
	acc, _ := r.Cookie("access_token")
	ref, _ := r.Cookie("refresh_token")

	if acc.Value != "" {
		if claims, err := auth.ParseAccess(acc.Value); err == nil {
			_ = us.Cache.DelJTI(r.Context(), "access"+claims.ID)
		}
	}

	if ref.Value != "" {
		if claims, err := auth.ParseAccess(ref.Value); err == nil {
			_ = rds.DelJTI(r.Context(), "ref"+claims.ID)
		}
	}
	auth.ClearAuthCookies(w)
	w.Write([]byte("Ok:true"))
}

func (us *UserService) HandleRefreshToken(w http.ResponseWriter, r *http.Request) {
	rds := us.Cache
	ref, err := auth.MustCookie(r, "refresh_token")
	if err != nil {
		log.Println(http.StatusUnauthorized, errors.New("missing refresh token"))
		return
	}
	claims, err := auth.ParseRefresh(ref)
	if err != nil {
		log.Println(http.StatusUnauthorized, errors.New("invalid refresh token"))
	}
	ctx := context.Background()
	if _, err := rds.GetUserByJTI(ctx, "refresh:"+claims.ID); err != nil {
		log.Println(http.StatusUnauthorized, errors.New("refresh revoked"))
		return
	}
	_ = rds.DelJTI(ctx, "refresh:"+claims.ID)

	toks, err := auth.IssueTokens(claims.Subject)
	if err != nil {
		log.Println(http.StatusInternalServerError, errors.New("could not issue new tokens"))
		return
	}
	if err := auth.Persist(ctx, rds, toks); err != nil {
		log.Println(http.StatusInternalServerError, errors.New("could not persist new tokens"))
		return
	}
	auth.SetAuthCookies(w, toks)
	log.Println(http.StatusCreated, "{ok: true}")

}

func (us *UserService) getPermissions(w http.ResponseWriter, r *http.Request) {
	// get the user specific permissions from db
	w.Write([]byte("permissions"))
}
