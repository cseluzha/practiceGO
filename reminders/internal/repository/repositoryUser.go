package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (	
	schemaUser = "users"
	tableUser  = "users"
)

type User struct {
	IdUser string `db:"user_id"`
	Email  string `db:"email"`
}

type userRepository struct {
	db *sql.DB
}

type UsersRepository interface {
	NewUser(user User) string
	UpdateUser(user User) int64
	DeleteUser(userId string) int64
	ListUsers() ([]User, error)
}

func (ur *userRepository) NewUser(user User) string {
	// close database
	defer ur.db.Close()
	insertStmt := `INSERT INTO ` + schemaUser + `.` + tableUser + `(user_id, email) VALUES ($1, $2) RETURNING user_id`
	var id string

	// Scan function will save the insert id in the id
	err := ur.db.QueryRow(insertStmt, user.IdUser, user.Email).Scan(&id)
	CheckError(err)
	fmt.Printf("Inserted %v in %v\n", id, tableUser)
	return id
}

func (ur *userRepository) UpdateUser(user User) int64 {
	// close database
	defer ur.db.Close()

	// create the update sql query
	updateStmt := `UPDATE ` + schemaUser + `.` + tableUser + ` SET email=$2  WHERE user_id=$1`

	// execute the sql statement
	res, err := ur.db.Exec(updateStmt, user.IdUser, user.Email)
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}

func (ur *userRepository) DeleteUser(userId string) int64 {
	// create the postgres db connection
	//ur.db = CreateConnection(dbnameUser)

	// close database
	defer ur.db.Close()

	// create the delete sql query
	deleteStmt := `DELETE FROM ` + schemaUser + `.` + tableUser + ` WHERE user_id=$1`
	// execute the sql statement
	res, err := ur.db.Exec(deleteStmt, userId)
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func (ur *userRepository) ListUsers() ([]User, error) {
	// create the postgres db connection
	//ur.db = CreateConnection(dbnameUser)

	// close database
	defer ur.db.Close()

	var users []User

	// create the select sql query
	sqlStatement := `SELECT * FROM ` + schemaUser + `.` + tableUser
	// execute the sql statement
	rows, err := ur.db.Query(sqlStatement)
	CheckError(err)
	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user User

		// unmarshal the row object to user
		err = rows.Scan(&user.IdUser, &user.Email)

		CheckError(err)

		// append the user in the users slice
		users = append(users, user)
	}
	// return empty users on error
	return users, err
}

func NewUserRepository() UsersRepository {
	return &userRepository{db: CreateConnection()}
}
