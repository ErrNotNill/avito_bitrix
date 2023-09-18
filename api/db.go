package api

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB //global variable

func InitDB(url string) (err error) {
	urlSettings := FindSettings("DatabaseUrl")
	fmt.Println("urlSettings", urlSettings)
	Db, err = sqlx.Connect("mysql", urlSettings)
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
	qry, err := Db.Query(`SELECT AccessToken FROM Token`)
	if err != nil {
		fmt.Println(`error query`)
	}
	for qry.Next() {
		if err := qry.Scan(&auth.AccessToken); err != nil {
			fmt.Println(`error scan`)
			//log.Fatal(err)
		}
	}
	return auth.AccessToken

	//todo add in dbase token date created
	//and check between date_created and expires date
}
