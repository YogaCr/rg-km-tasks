package repository

import (
	"database/sql"
	"errors"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (User, error) {
	user := &User{}
	row := u.db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
	return *user, nil // TODO: replace this
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	users := make([]User, 0)
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, *user)
	}
	return users, nil // TODO: replace this
}

func (u *UserRepository) Login(username string, password string) (*string, error) {
	sqlStatement := "SELECT count(*) FROM users WHERE username = ? AND password = ?"
	res := u.db.QueryRow(sqlStatement, username, password)
	count := 0
	res.Scan(&count)
	if count > 0 {
		return &username, nil
	}
	return nil, errors.New("Login Failed") // TODO: replace this
}

func (u *UserRepository) InsertUser(username string, password string, role string, loggedin bool) error {
	sqlStatement := "INSERT INTO users (username, password, role, loggedin) VALUES(?,?,?,?)"
	_, err := u.db.Exec(sqlStatement, username, password, role, loggedin)
	return err // TODO: replace this
}

func (u *UserRepository) FetchUserRole(username string) (*string, error) {
	sqlStatement := "SELECT role FROM users WHERE username=?"
	role := ""
	res := u.db.QueryRow(sqlStatement, username)
	err := res.Scan(&role)
	return &role, err // TODO: replace this
}
