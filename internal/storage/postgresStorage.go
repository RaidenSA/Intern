package storage

import (
	"database/sql"
	"log"
)

type DataBase struct {
	ConnStr string
}

func (db DataBase) TokenToValue(token string) (string, bool) {
	//open connection
	connection, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		return "connection error", false
	}
	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Fatal(err, "defer error")
		}
	}(connection)
	row := connection.QueryRow("select longurl from dbase.public.tokens where token = $1", token)
	longUrl := ""
	err = row.Scan(&longUrl)
	println(err)
	if err == sql.ErrNoRows {
		return "connection error", false
	}
	//close connection
	return longUrl, true
}
func (db DataBase) ValueToToken(value string) (string, bool) {
	//open connection

	connection, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		return "connection error select", false
	}
	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Fatal(err, "defer error")
		}
	}(connection)
	row := connection.QueryRow("select token from dbase.public.tokens where longurl = $1", value)
	token := ""
	err = row.Scan(&token)
	if err == sql.ErrNoRows {
		return "connection error select url", false
	}
	//do things (select)
	//close connection
	return token, true
}
func (db DataBase) SetToken(token string, value string) {
	//open connection
	connection, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		log.Fatal(err, "Connection error")
		return
	}
	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Fatal(err, "defer error")
		}
	}(connection)
	_, err = connection.Exec("insert into dbase.public.tokens (token, longurl) values ($1, $2)", token, value)

	if err != nil {
		log.Fatal(err, "connection insert error")
	}
	//do things (select)-> insert
	//close connection
	return
}
