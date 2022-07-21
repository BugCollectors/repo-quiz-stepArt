package repos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"study/internal/types"
)

const (
	tableName = "users"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (a *UserRepo) CreateUser(user types.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username) values ($1, $2)", tableName)
	row := a.db.QueryRow(query, user.Name, user.Username)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *UserRepo) GetUser(id interface{}) (*types.User, error) {
	var user types.User

	strct, ok := id.(struct {
		ID string `json:"id" db:"id"`
	})
	if !ok {
		return nil, fmt.Errorf("error ept")
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", tableName)
	row := a.db.QueryRow(query, strct.ID)

	if err := row.Scan(&user.Id, &user.Name, &user.Username); err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserRepo) DeleteUser() {}
