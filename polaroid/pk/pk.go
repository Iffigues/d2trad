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
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL primary key,
		login VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(50) UNIQUE NOT NULL,
		mail VARCHAR(50) UNIQUE NOT NULL
	);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS option (
		id SERIAL,
		user_id INT references users 
	);`)
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(`INSERT INTO users(login, password, mail)  VALUES ('iffigues','Petassia01', 'iffigues@vivaldi.net');`)
	db.Exec(`INSERT INTO users (login , password, mail)VALUES ('css', 'Mince1234', 'iffigues@vivaldi.net');`);
}

func (a *Pk) Connect() (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", a.Host, a.Port, a.User, a.Password, a.Dbname)
	if db, err = sql.Open("postgres", psqlInfo); err != nil {
		return
	}
	return db, db.Ping()
}
