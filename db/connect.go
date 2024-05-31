package db

import (
	"EJawTest/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	user := "myuser"
	password := "mypassword"
	dbname := "postgres"
	host := "postgres-container"

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s", user, password, dbname, host)
	var err error
	// connect to database
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	return nil
}

func GetUser(login string) (models.User, error) {
	var user models.User

	err := db.QueryRow("SELECT id, name, role, login, password, phone_number FROM users WHERE login = $1", login).Scan(&user.ID, &user.Name, &user.Role, &user.Login, &user.Password, &user.PhoneNumber)
	return user, err
}
