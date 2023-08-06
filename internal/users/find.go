package users

import (
	"database/sql"

	"github.com/gofrs/uuid"
)

func Find(tx *sql.DB, id uuid.UUID) (User, error) {
	var user User
	err := tx.QueryRow(`
	SELECT 
		users.id,
		users.name,
		users.color,
		users.email
	FROM users
	WHERE users.id = $1;
	`, id).Scan(&user.ID, &user.Name, &user.Color, &user.Email)
	if err != nil {
		return user, err
	}

	user.NameInitial = user.Initial()
	return user, nil
}
