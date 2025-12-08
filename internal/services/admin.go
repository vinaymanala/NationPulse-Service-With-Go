package services

import (
	"net/http"

	u "github.com/nationpulse-bff/internal/utils"
)

type AdminService struct {
	// Add any dependencies like database connections here
	Configs *u.Configs
}

func (as *AdminService) SetPermissions(userID string, newPermissions map[string]int) bool {
	// upate the database with new permissions
	return true
}

func (as *AdminService) GetAllPermissions(w http.ResponseWriter, r *http.Request) {
	// get all the permissions and return
	w.Write([]byte(""))
}
