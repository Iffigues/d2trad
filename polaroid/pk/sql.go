package pk

import (
	"database/sql"
	_ "github.com/lib/pq"
)


func InsertUsers(st *sql.DB, email, login, pwd string) (err error){
	println(login)
	i := `INSERT INTO users (login, password, mail) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err = st.QueryRow(i,login,pwd,email).Scan(&id)
	if err != nil {
		return err
	}
	println("id=",id)
	return
}

func UserExists(st *sql.DB, login, pd string) (t bool, err error) {
	var lgn, pwd string
	i := `SELECT login, password FROM users WHERE login = $1 LIMIT 1`;
	err = st.QueryRow(i, login).Scan(&lgn, &pwd)
	if err != nil  {
		return false, err
	}
	return true, nil
}
