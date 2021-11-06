package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/AdairHdz/auth-server/entity"
)

var (	
	databaseUser string
	databasePassword string	
	databaseName string
	databaseContainerName string
)



func Query(emailAddress string, user *entity.User) (err error) {
	db, err := newConnection()
	
	if err != nil {
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT users.id as user_id, users.names, users.lastname," +
	"states.id as state_id," +
	"accounts.email_address, accounts.user_type, accounts.verified, accounts.password," +
    "IF(accounts.user_type = 1, (SELECT id from otw.service_providers WHERE user_id = users.id), (SELECT id from otw.service_requesters WHERE user_id = users.id)) AS id" +
	" from otw.users" +
	" inner join otw.states on states.id = users.state_id" +
    " inner join otw.accounts on users.id = accounts.user_id" +
	" where accounts.email_address = ?")
	
	if err != nil {
		return
	}

	defer stmt.Close()
	
	stmt.QueryRow(emailAddress).Scan(
		&user.UserID,
		&user.Names,
		&user.LastName,
		&user.StateID,
		&user.EmailAddress,
		&user.UserType,
		&user.Verified,
		&user.Password,
		&user.ID,
	)

	return
}

func newConnection() (db *sql.DB, err error) {	
	databaseUser = os.Getenv("DB_USER")
	databasePassword = os.Getenv("DB_PASSWORD")	
	databaseName = os.Getenv("DB_NAME")
	databaseContainerName = os.Getenv("DB_CONTAINER_NAME")
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	databaseUser, databasePassword, databaseContainerName, databaseName))
	
	if err != nil {		
		return
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	
	return
}