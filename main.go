package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB_NAME = os.Getenv("DB_NAME")
)

type Employee struct {
	Id            string
	Name          string
	DOB           string
	Sex           string
	Address       string
	Religion      string
	MarriedStatus string
	Work          string
	BloodType     string
}

func main() {
	var (
		// ADDRESS     = os.Getenv("ADDRESS")
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
	)

	//Connect to database
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, errorOpenDB := sql.Open("postgres", dbInfo)
	if errorOpenDB != nil {
		log.Fatalf(errorOpenDB.Error())
	}

	defer db.Close()

	//Show database data
	var employees []Employee

	rows, errorQuery := db.Query("SELECT * FROM employee")
	if errorQuery != nil {
		log.Fatalf(errorQuery.Error())
	}

	for rows.Next() {
		var employee Employee

		errorScan := rows.Scan(
			&employee.Id,
			&employee.Name,
			&employee.DOB,
			&employee.Sex,
			&employee.Address,
			&employee.Religion,
			&employee.MarriedStatus,
			&employee.Work,
			&employee.BloodType,
		)
		if errorScan != nil {
			log.Fatalf(errorScan.Error())
		}

		employees = append(employees, employee)
	}

	fmt.Println(employees)
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
