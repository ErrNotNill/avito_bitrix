package api

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB //global variable

func InitDB(url string) (err error) {

	Db, err = sqlx.Connect("mysql", url)
	if err != nil {
		fmt.Println("not connected")
		return
	}
	err = Db.Ping()
	return
}

func AddToken(auth *Auth) {
	fmt.Println("db connected")
	result, err := Db.Exec(`insert into Token (AccessToken,TokenType,ExpiresIn ) values (?, ?, ?)`,
		&auth.AccessToken, &auth.TokenType, &auth.ExpiresIn)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}

func GetToken() string {
	fmt.Println("db connected")
	auth := &Auth{}
	qry, err := Db.Query(`SELECT * FROM Token`)
	if err != nil {
		fmt.Println(`error query`)
	}
	for qry.Next() {
		if err := qry.Scan(&auth.AccessToken, &auth.ExpiresIn, &auth.TokenType); err != nil {
			fmt.Println(`error scan`)
			//log.Fatal(err)
		}
	}
	return auth.AccessToken
}
