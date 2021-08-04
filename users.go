package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	li "user/loggers"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phonenumber"`
	IsActive    bool      `json:"isactive"`
}

var DB *gorm.DB

func InitialMigration() {
	var err error
	const DNS = "root:hello@tcp(127.0.0.1:3306)/userdb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		li.NewLogger().Error("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}

func openCSV(url string) [][]string {
	var records [][]string
	csv_file, err := os.Open(url)
	if err != nil {
		li.NewLogger().Error("Error opening the file")
		return records
	}
	defer csv_file.Close()

	r := csv.NewReader(csv_file)
	records, err = r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return records
}

func readCSV(w http.ResponseWriter, r *http.Request) {
	var standardLogger = li.NewLogger()
	data, _ := ioutil.ReadAll(r.Body)
	records := openCSV(string(data))
	var user User
	type ID uuid.UUID
	visited := make(map[ID]struct{})

	for _, rec := range records {
		user.ID, _ = uuid.FromString(rec[0])
		user.Name = rec[1]
		user.Email = rec[2]
		user.PhoneNumber = rec[3]
		user.IsActive, _ = strconv.ParseBool(rec[4])
		if checkValidity(user.Name, user.PhoneNumber, user.Email, standardLogger) {
			if _, ok := visited[ID(user.ID)]; ok {
				id := uuid.NewV4()
				standardLogger.DuplicateEntry()
				user.ID = id
			} else {
				visited[ID(user.ID)] = struct{}{}
			}
			dbuser := User{
				ID:          user.ID,
				Name:        user.Name,
				Email:       user.Email,
				PhoneNumber: user.PhoneNumber,
				IsActive:    user.IsActive,
			}
			DB.Create(&dbuser)
		} else {
			continue
		}
	}
	fmt.Fprintf(w, "Validated the csv file and updated in database")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}
