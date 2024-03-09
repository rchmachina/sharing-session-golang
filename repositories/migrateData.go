package repositories

import (

	"log"

	_ "github.com/lib/pq"
	sql "github.com/rchmachina/sharing-session-golang/sql"
	"gorm.io/gorm"
)

// kontrak
type funcSql interface {
	MigrateTableSql()error
	MigrateFunctionSql()error
}

func RepositoryMigrate(db *gorm.DB) *repository {
	return &repository{db}
}


func (r *repository) MigrateTableSql()error{

	err := r.db.Exec(`DROP TABLE IF EXISTS users;;`).Error
    if err != nil {
        (log.Printf("failed to create table: %v", err))
		return err
    }

    err = r.db.Exec(sql.CreateTable).Error
    if err != nil {
        (log.Printf("failed to create table: %v", err))
		return err
    }
	return nil 
}

func (r *repository) MigrateFunctionSql()error{

	//drop all functions first 

	dropAllFunc := []string{
		`DROP FUNCTION IF EXISTS users_get_all;`,
		`DROP FUNCTION IF EXISTS users_delete;`,
		`DROP FUNCTION IF EXISTS users_login;`,
		`DROP FUNCTION IF EXISTS users_create_user;`,
		`DROP FUNCTION IF EXISTS users_update;`}

	for _, f := range dropAllFunc{
		err := r.db.Exec(f).Error
		if err != nil {
			(log.Printf("failed to create table: %v", err))
			return err
		}

	}

	migrateAllFunc := []string{
		sql.CreateUserFunction,
		sql.DeleteUserFunction,
		sql.IsUserExistFunction,
		sql.ReadAllUsersFunction,
		sql.UpdateUserFunction}
	
	for _, getFunc := range migrateAllFunc{
		err := r.db.Exec(getFunc).Error
		if err != nil {
			(log.Printf("failed to create table: %v", err))
			return err
		}

	}

	return nil 
}