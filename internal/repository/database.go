package repository

import (
	"log"

	"github.com/omerberkcan/banking-transfer/internal/config"
	"github.com/omerberkcan/banking-transfer/internal/repository/user"
	"github.com/omerberkcan/banking-transfer/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db    *gorm.DB
	users UserRepository
}

func ConnectMysqlServer(dbcfg *config.DBConfiguration) (*gorm.DB, error) {
	conStr := dbcfg.Username + ":" + dbcfg.Password + "@tcp(" + dbcfg.Host + ")/" + dbcfg.Dbname + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(conStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Printf("%s \n", "Mysql Connection Succesful \n")

	MigrateTables(db)

	return db, nil
}

func MigrateTables(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Printf("failed to migrate users: %v", err)
		return err
	}
	return nil
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
