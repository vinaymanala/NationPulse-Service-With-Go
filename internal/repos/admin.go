package repos

import (
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"
	. "github.com/nationpulse-bff/internal/utils"
)

type AdminRepo struct {
	Configs *Configs
}

func NewAdminRepo(configs *Configs) *AdminRepo {
	return &AdminRepo{
		Configs: configs,
	}
}

func (ar *AdminRepo) SetUserPermissions(updatePermissions UpdatePermissions) error {
	// upate the database with new permissions
	tx, err := ar.Configs.Db.Client.Begin(ar.Configs.Context)
	if err != nil {
		return err
	}
	// rollback if not committed
	defer tx.Rollback(ar.Configs.Context)
	// Call stored procedure
	_, err = tx.Exec(ar.Configs.Context, "CALL update_user_permissions($1, $2, $3, $3)", updatePermissions.UserID, updatePermissions.RoleID, updatePermissions.Modules, updatePermissions.Permissions)
	if err != nil {
		log.Println("failed to update permissions:")
		return err
	}

	if err := tx.Commit(ar.Configs.Context); err != nil {
		return err
	}

	err = ar.Configs.Cache.DelData(ar.Configs.Context, "utils:modulePermissions:"+strconv.Itoa(updatePermissions.UserID))
	if err != nil {
		log.Printf("Failed to invalidate cache for user %d: %v", updatePermissions.UserID, err)
	}

	return nil
}

func (ar *AdminRepo) GetUsers() (any, error) {
	var users []Users
	sqlStatement := `SELECT * from users`
	rows, err := ar.Configs.Db.Client.Query(ar.Configs.Context, sqlStatement)
	if err != nil {
		log.Println("Error fetching users from DB", err)
		return nil, err
	}

	users, err = pgx.CollectRows(rows, pgx.RowToStructByPos[Users])
	if err != nil {
		log.Printf("Error scanning rows: %v\n", err)
	}

	return users, nil
}
