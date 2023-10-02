package repository

import (
	"log"

	"github.com/omerberkcan/banking-transfer/internal/config"
	"github.com/omerberkcan/banking-transfer/internal/repository/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db    *gorm.DB
	users UserRepository
}

func ConnectMysqlServer(dbcfg *config.DBConfiguration) (*gorm.DB, error) {
	conStr := dbcfg.Username + ":" + dbcfg.Password + "@tcp(" + dbcfg.Host + ")/" + dbcfg.Dbname
	db, err := gorm.Open(mysql.Open(conStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Printf("%s \n", "Mysql Connection Succesful \n")
	return db, nil
}

func New(db *gorm.DB) *Database {
	return &Database{
		db:    db,
		users: user.NewRepository(db),
	}
}

func (db *Database) Users() UserRepository {
	return db.users
}

// Ping checks if database is up
func (db *Database) Ping() error {
	sqlDB, err := db.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}
