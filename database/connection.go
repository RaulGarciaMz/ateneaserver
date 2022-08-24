package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(usuario, password, dbName, host string) (*gorm.DB, error) {
	uri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s sslmode=disable", host, usuario, dbName, password)

	//dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: uri,  // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	  }), &gorm.Config{})

	//db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
