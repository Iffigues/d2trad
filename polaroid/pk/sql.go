package pk

import (
	"database/sql"
	_ "github.com/lib/pq"
)


func InsertUsers(st *sql.DB, email, login, pwd string) (err error){
	i := `INSERT INTO userss (login, password, mail) VALUES ($1, $2, $3) RETURNING id`
	id := 0;
	return st.QueryRow(i,login,pwd,email).Scan(&id)
}
