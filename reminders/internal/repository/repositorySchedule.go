package repository

import (
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/lib/pq"
)

const (
	schemaSchedule = "schedules"
	tableSchedule  = "scheduled_items"
)

type Schedule struct {
	Id          uuid.UUID   `json:"id" db:"id"`
	Description string      `json:"description" db:"description"`
	Users       []uuid.UUID `json:"users" db:"users" pg:"array"`
}

type scheduleRepository struct {
	db *sql.DB
}

type SchedulesRepository interface {
	NewSchedule(schedule Schedule) string
	UpdateSchedule(schedule Schedule) int64
	DeleteSchedule(id string) int64
	ListSchedules() ([]Schedule, error)
}

func (ur *scheduleRepository) NewSchedule(schedule Schedule) string {
	//TODO: validate the new schedule not exist into data base.
	// close database
	defer ur.db.Close()

	insertStmt := `INSERT INTO ` + schemaSchedule + `.` + tableSchedule + ` (id, description, users) VALUES ($1, $2, $3) RETURNING id`
	var id string

	// Scan function will save the insert id in the id
	err := ur.db.QueryRow(insertStmt, schedule.Id, schedule.Description, pq.Array(schedule.Users)).Scan(&id)
	CheckError(err)
	fmt.Printf("Inserted %v in %v\n", id, tableSchedule)
	return id
}

func (ur *scheduleRepository) UpdateSchedule(schedule Schedule) int64 {
	// close database
	defer ur.db.Close()

	// create the update sql query
	updateStmt := `UPDATE ` + schemaSchedule + `.` + tableSchedule + ` SET description=$2, users=$3 WHERE id=$1`

	// execute the sql statement
	res, err := ur.db.Exec(updateStmt, schedule.Id, schedule.Description, schedule.Users)
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v in %v\n", rowsAffected, tableSchedule)
	return rowsAffected
}

func (ur *scheduleRepository) DeleteSchedule(id string) int64 {
	// close database
	defer ur.db.Close()

	// create the delete sql query
	deleteStmt := `DELETE FROM ` + schemaSchedule + `.` + tableSchedule + ` WHERE id=$1`
	// execute the sql statement
	res, err := ur.db.Exec(deleteStmt, id)
	CheckError(err)
	// check how many rows affected
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func (ur *scheduleRepository) ListSchedules() ([]Schedule, error) {
	// close database
	defer ur.db.Close()

	var schedules []Schedule

	// create the select sql query
	sqlStatement := `SELECT * FROM ` + schemaSchedule + `.` + tableSchedule
	// execute the sql statement
	rows, err := ur.db.Query(sqlStatement)
	CheckError(err)
	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var schedule Schedule

		// unmarshal the row object to schedule
		err = rows.Scan(&schedule.Id, &schedule.Description, &schedule.Users)

		CheckError(err)

		// append the schedule in the schedules slice
		schedules = append(schedules, schedule)
	}
	// return empty schedules on error
	return schedules, err
}

func NewScheduleRepository() SchedulesRepository {
	return &scheduleRepository{db: CreateConnection()}
}
