package users

import (
	"database/sql"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

type User struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	Name        string       `json:"name" db:"name"`
	Email       string       `json:"email" db:"email"`
	Color       nulls.String `json:"color" db:"color"`
	NameInitial string       `json:"initial" db:"-"`
}

func All(tx *sql.DB, userID uuid.UUID) ([]User, error) {
	rows, err := tx.Query(`
	SELECT 
		users.id,
		users.name,
		users.color,
		users.email
	FROM users
	WHERE users.id != $1;
	`, userID)
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Color, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (user User) Initial() string {
	if user.Name != "" {
		return string(user.Name[0])
	}
	return "?"
}
