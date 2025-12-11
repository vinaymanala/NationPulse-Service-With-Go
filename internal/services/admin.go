package services

import (
	"net/http"

	. "github.com/nationpulse-bff/internal/utils"
)

type AdminService struct {
	// Add any dependencies like database connections here
	Configs *Configs
}

func NewAdminService(configs *Configs) *AdminService {
	return &AdminService{
		Configs: configs,
	}
}
func (as *AdminService) SetPermissions(userID string, newPermissions map[string]int) bool {
	// upate the database with new permissions
	return true
}

func (as *AdminService) GetAllPermissions(w http.ResponseWriter, r *http.Request) {
	// get all the permissions and return
	w.Write([]byte(""))
}
