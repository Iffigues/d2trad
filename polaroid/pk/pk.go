package pk

import (
	"database/sql"
	"fmt"
	"log"
	"polaroid/config"

	_ "github.com/lib/pq"
)

type Pk struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func NewPk(s config.Conf) (a *Pk) {
	a = &Pk{
		Host:     s["host"].(string),
		Port:     s["port"].(int),
		User:     s["user"].(string),
		Password: s["password"].(string),
		Dbname:   s["dbname"].(string),
	}
	a.Starter()
	return

}

func (a *Pk) IsUsers() (ok bool) {
	db, err := a.Connect()
	if err != nil {
		return
	}
	defer db.Close()
	var t string
	err = db.QueryRow(`SELECT password FROM users WHERE login=$1`, "iffigues").Scan(&t)
	fmt.Println(t, err)
	return
}

func (a *Pk) Starter() {

	db, err := a.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS article (
		title VARCHAR(50) UNIQUE,
		describe VARCHAR(50), 
		logo VARCHAR(50),
		body VARCHAR(50), 
		css VARCHAR(50),
		name VARCHAR(50) UNIQUE
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS gallerie (
		title VARCHAR(50) UNIQUE,
		describe VARCHAR(50),
		logo VARCHAR(50),
		body VARCHAR(50),
		css VARCHAR(50),
		name VARCHAR(50) UNIQUE
	);`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS connect (
		token VARCHAR(128)
	);`)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		login VARCHAR(50),
		logo VARCHAR(50),
		password VARCHAR(50),
		mail VARCHAR(50)
	);`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO users  VALUES ('iffigues','','Petassia01', 'iffigues@vivaldi.net');`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO users VALUES ('css', 'Mince1234', 'iffigues@vivaldi.net');`);

	if err != nil {
		log.Fatal(err)
	}
}

func (a *Pk) Connect() (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", a.Host, a.Port, a.User, a.Password, a.Dbname)
	if db, err = sql.Open("postgres", psqlInfo); err != nil {
		return
	}
	return db, db.Ping()
}
