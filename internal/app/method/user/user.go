package user

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserDBRepository struct {
	db sql.DB
}

func HandleError(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func OpenDB(db *sql.DB) (sql.DB){
	db, err := sql.Open(
		"mysql",
		"rinat:jojo1337@tcp(127.0.0.1:3306)/website",
	)
	HandleError(err)
	return *db
}

func GetLastID(db sql.DB) int{
	var id int
	rows, err := db.Query("SELECT user_id FROM users")
	HandleError(err)
	for rows.Next(){
		err := rows.Scan(&id)
		if err != nil{
			return 0
		}
	}
	return id
}

func HashPassword (password string) (string, error){
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}
	return string(encryptedPassword), nil
}
