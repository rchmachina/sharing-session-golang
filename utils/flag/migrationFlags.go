package flags

import (
	"flag"
	"log"

	"github.com/rchmachina/sharing-session-golang/repositories"
	"github.com/rchmachina/sharing-session-golang/utils/database"
)


func MigrationFlags(){

	db := database.DatabaseConnection()
	var isMigrateFunction = flag.Bool("migrateDatabaseFunction", false, "migrate function to database?")
	var isMigrateTable = flag.Bool("migrateTable", false, "migrate database?")

	flag.Parse()

	if *isMigrateTable {
		RepositoryMigrate := repositories.RepositoryMigrate(db)
		err := RepositoryMigrate.MigrateTableSql()
		if err == nil {
			log.Println("succes migrate database")
		}else{
			panic(err)
		}
	} else {
		log.Println("skiping migration")
	}
	if *isMigrateFunction {
		RepositoryMigrate := repositories.RepositoryMigrate(db)
		err := RepositoryMigrate.MigrateFunctionSql()
		if err == nil {
			log.Println("succes migrate function database")
		}else{
			panic(err)
		}
	} else {
		log.Println("skiping migration function")
	}
}