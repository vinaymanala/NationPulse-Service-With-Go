package repos

import (
	"encoding/json"
	"log"
	"strconv"

	. "github.com/nationpulse-bff/internal/utils"
)

type UtilsRepo struct {
	Configs *Configs
}

func NewUtilsRepo(configs *Configs) *UtilsRepo {
	return &UtilsRepo{
		Configs: configs,
	}
}

var utilsID = "utils:"

func (ur *UtilsRepo) GetPermissions(userID string) (any, error) {
	var userPermissions []UserPermissions
	var permissions []int
	data, err := GetDataFromCache(ur.Configs, utilsID+"permissions", &permissions)
	if err != nil {
		log.Println("Cache Get Failed. Trying DB.")
	} else {
		return *data, nil
	}

	sqlStatement := `SELECT * FROM get_user_permissions($1);`
	id, err := strconv.Atoi(userID)
	if err != nil {
		log.Fatal("Error converting userId to int")
		return nil, err
	}
	rows, err := ur.Configs.Db.Client.Query(ur.Configs.Context, sqlStatement, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var userPermission UserPermissions
		if err := rows.Scan(
			&userPermission.Name,
			&userPermission.Email,
			&userPermission.RoleId,
			&userPermission.RoleName,
			&userPermission.RoleDescription,
			&userPermission.ModuleID,
			&userPermission.ModuleName,
			&userPermission.ModuleValue,
			&userPermission.PermissionID,
			&userPermission.PermissionName,
			&userPermission.PermissionValue,
		); err != nil {
			log.Fatalf("Error scanning a row: %v\n", err)
			return nil, err
		}
		userPermissions = append(userPermissions, userPermission)
	}

	for _, permission := range userPermissions {
		permissions = append(permissions, permission.PermissionValue)
	}
	defer rows.Close()
	marshalledData, err := json.Marshal(permissions)
	if err != nil {
		log.Println("Error marshalling data", err)
	}
	if err := ur.Configs.Cache.SetData(ur.Configs.Context, utilsID+"permissions", marshalledData); err != nil {
		log.Println("Error Set Cache Data", err)
	}
	return permissions, nil
}
