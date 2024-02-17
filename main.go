package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/miqbals17/cryspy"
	randomizer "github.com/miqbals17/randomizer"
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
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
		KEY         = os.Getenv("KEY")
		IV          = os.Getenv("IV")
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

	//Insert encrypted data to database (with example data)
	var newEmployee = Employee{
		Id:            randomizer.RandomString(8),
		Name:          "Yusuf Dwiyanto",
		DOB:           "21 Juni 2000",
		Sex:           "Laki-laki",
		Address:       "Tlawong, Sawit, Boyolali, Jawa Tengah",
		Religion:      "Islam",
		MarriedStatus: "Belum Menikah",
		Work:          "Pegawai Swasta",
		BloodType:     "O+",
	}

	var encryptedObject = EncryptObject(newEmployee, KEY, IV)

	_, errInsert := db.Exec("INSERT INTO employee VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9);", encryptedObject.Id, encryptedObject.Name, encryptedObject.DOB, encryptedObject.Sex, encryptedObject.Address, encryptedObject.Religion, encryptedObject.MarriedStatus, encryptedObject.Work, encryptedObject.BloodType)
	if errInsert != nil {
		log.Fatalf(errInsert.Error())
	}

	fmt.Println("Data berhasil ditambahkan!")
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func EncryptObject(data Employee, KEY string, IV string) Employee {
	encryptedData := Employee{
		Id:            data.Id,
		Name:          data.Name,
		DOB:           fmt.Sprintf("%x", cryspy.EncryptCBC(data.DOB, KEY, IV)),
		Sex:           data.Sex,
		Address:       fmt.Sprintf("%x", cryspy.EncryptCBC(data.Address, KEY, IV)),
		Religion:      fmt.Sprintf("%x", cryspy.EncryptCBC(data.Religion, KEY, IV)),
		MarriedStatus: fmt.Sprintf("%x", cryspy.EncryptCBC(data.MarriedStatus, KEY, IV)),
		Work:          fmt.Sprintf("%x", cryspy.EncryptCBC(data.Work, KEY, IV)),
		BloodType:     fmt.Sprintf("%x", cryspy.EncryptCBC(data.BloodType, KEY, IV)),
	}

	return encryptedData
}
