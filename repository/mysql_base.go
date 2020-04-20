package repository

// import (
// 	"database/sql"
// 	"fmt"
// 	"travel/config"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var db *sql.DB
// var err error
// var isOpen bool

// func ConnectDB() {
// 	isConnect, _ := ping()
// 	if isConnect {
// 		return
// 	}

// 	db, err = sql.Open("mysql", ""+config.USERNAME_MYSQL+":"+config.PASSWORD_MYSQL+"@tcp("+config.HOST_MYSQL+":"+config.PORT+")/"+config.TABLE+"")
// 	if err != nil {
// 		fmt.Println("error ConnectDB" + err.Error())
// 		// kenlog.PrintlnInsertMySql("error connect: " + err.Error())
// 		ConnectDB()
// 	}

// 	isOpen = true

// 	isConnect, err := ping()
// 	if !isConnect {
// 		fmt.Println("error ConnectDB" + err.Error())
// 		// kenlog.PrintlnInsertMySql("error ping: " + err.Error())
// 		ConnectDB()
// 	} else {
// 		fmt.Print("error" + err.Error())
// 	}
// }

// func GetDB() *sql.DB {
// 	return db
// }

// func ping() (bool, error) {
// 	if !isOpen {
// 		return false, nil
// 	}
// 	err := db.Ping()
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// func CloseConnect() {
// 	db.Close()
// }
