package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

func WriteJSON(w http.ResponseWriter, status int, data any, success bool, err any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var response = &ApiResponse{
		Data:      data,
		IsSuccess: success,
		Error:     err,
	}
	log.Printf("Message: %s, isSuccess: %t, Error:%v \n", "Data received", success, err)
	return json.NewEncoder(w).Encode(response)
}

func GetUserDetailsFromCache(r *http.Request, configs *Configs) {
	fmt.Printf("Form Details: %s\n", r.Form.Get("userID"))
}

func GetDataFromCache[T any](configs *Configs, key string, mappedStruct T) (*T, error) {
	//var zero T
	data, err := configs.Cache.GetData(configs.Context, key)
	fmt.Println("CACHE DATA:", data, err)
	if err != nil {
		log.Printf("Error fetching data from cache %s\n", err)
		return nil, errors.New("error fetching data from cache")
	}
	if err := json.Unmarshal([]byte(data), &mappedStruct); err != nil {
		log.Println("Error unmarshalling data from cache.")
		return nil, errors.New("error unmarshalling data from cache")
	}
	fmt.Println("==================================")
	// fmt.Println("Unmarshal data", &mappedStruct)
	fmt.Println("Fetched Data from Cache!!")
	return &mappedStruct, nil
}

func checkModulePermission(permissions []UserPermissions, moduleID int) bool {
	log.Printf("HEALTH_ID %d", HEALTH_ID)
	log.Printf("Permissions LOG: %v", permissions)
	for _, p := range permissions {
		log.Printf("p.ModuleID: %d, moduleID: %d", p.ModuleValue, moduleID)
		if p.ModuleValue == moduleID {
			return true
		}
	}
	return false
}

func HasPermissions(requestPath string, permissions *[]UserPermissions) bool {
	log.Printf("PERMISSIONS ARG %v", *permissions)
	switch {
	case strings.HasPrefix(requestPath, ADMIN_PERMISSION):
		return checkModulePermission(*permissions, PERMISSION_ID)
	case strings.HasPrefix(requestPath, DASHBOARD) || strings.HasPrefix(requestPath, PERMISSION):
		return checkModulePermission(*permissions, DASHBOARD_ID)
	case strings.HasPrefix(requestPath, POPULATION) || strings.HasPrefix(requestPath, POPULATION_GROWTH):
		return checkModulePermission(*permissions, POPULATION_ID)
	case strings.HasPrefix(requestPath, HEALTH):
		log.Println("Touched Health Case")
		return checkModulePermission(*permissions, HEALTH_ID)
	case strings.HasPrefix(requestPath, ECONOMY):
		return checkModulePermission(*permissions, ECONOMY_ID)
	case strings.HasPrefix(requestPath, GROWTH):
		return checkModulePermission(*permissions, GROWTH_ID)
	case strings.HasPrefix(requestPath, REPORTING):
		return checkModulePermission(*permissions, REPORTING_ID)
	default:
		return false
	}
}

func FetchPermissionsFromDB(configs *Configs, sqlStatement string, userPermissions []UserPermissions, id int) ([]UserPermissions, error) {

	rows, err := configs.Db.Client.Query(configs.Context, sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// for rows.Next() {
	// 	var userPermission UserPermissions
	// if err := rows.Scan(
	// 	&userPermission.Name,
	// 	&userPermission.Email,
	// 	&userPermission.RoleId,
	// 	&userPermission.RoleName,
	// 	&userPermission.RoleDescription,
	// 	&userPermission.ModuleID,
	// 	&userPermission.ModuleName,
	// 	&userPermission.ModuleValue,
	// 	&userPermission.PermissionID,
	// 	&userPermission.PermissionName,
	// 	&userPermission.PermissionValue,
	// );
	data, err := pgx.CollectRows(rows, pgx.RowToStructByPos[UserPermissions])
	if err != nil {
		log.Fatalf("Error scanning a row: %v\n", err)
		return nil, err
	}
	// userPermissions = append(userPermissions, userPermission)
	return data, nil

}

func GetModulePermissionsFromCache(configs *Configs, userID int, key string, permissions []UserPermissions, w http.ResponseWriter, r *http.Request) ([]UserPermissions, error) {
	cacheData, err := configs.Cache.GetData(configs.Context, key)
	if err != nil {
		// cache miss, fetch db
		sqlStatement := `SELECT * FROM get_user_permissions($1);`
		data, err := FetchPermissionsFromDB(configs, sqlStatement, permissions, userID)
		log.Println("Fetching permissions from DB.")
		if err != nil {
			log.Println("Error: UnAuthorized")
			WriteJSON(w, http.StatusUnauthorized, nil, false, "Cannot fetch permissions, User not authroized.")
		}
		marshalledData, err := json.Marshal(data)
		if err != nil {
			log.Println("Error marshalling data")
			return nil, err
		}
		if err := configs.Cache.SetData(configs.Context, "utils:modulePermissions:"+strconv.Itoa(userID), marshalledData); err != nil {
			log.Println("Error Set Cache Data for route level permissions", err)
			return data, nil
		}
		return data, nil
	}
	if err := json.Unmarshal([]byte(cacheData), &permissions); err != nil {
		log.Println("Error unmarshalling permissions dataa", err)
		return nil, err
	}
	return permissions, nil
}
