package repositories

import (
	"backend/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *users {
	return &users{db}
}

func (repo users) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare("insert into users (name, nick, email, password) values ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(user.Name, user.Nick, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (repo users) Update(ID uint64, user models.User) error {
	statement, err := repo.db.Prepare("UPDATE users SET name = $1, nick = $2, email = $3 WHERE id = $4")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Query(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repo users) UpdatePassword(ID uint64, password string) error {
	statement, err := repo.db.Prepare("UPDATE users SET password = $1 WHERE id = $2")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Query(password, ID); err != nil {
		return err
	}

	return nil
}

func (repo users) Delete(ID uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM users WHERE id = $1")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Query(ID); err != nil {
		return err
	}

	return nil
}

func (repo users) FindAll(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := repo.db.Query("SELECT id, name, nick, email, created_at FROM users WHERE name ILIKE $1 OR nick LIKE $2", nameOrNick, nameOrNick)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo users) FindOne(ID uint64) (models.User, error) {
	lines, err := repo.db.Query("SELECT id, name, nick, email, created_at FROM users WHERE id = $1", ID)

	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo users) FindByEmail(email string) (models.User, error) {
	lines, err := repo.db.Query("SELECT id, password FROM users WHERE email = $1", email)

	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo users) FindById(userId uint64) (string, error) {
	lines, err := repo.db.Query("SELECT password FROM users WHERE id = $1", userId)

	if err != nil {
		return "", err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repo users) Follow(userId, followerId uint64) error {
	statement, err := repo.db.Prepare("INSERT INTO followers (user_id, follower_id) VALUES ($1, $2) ON CONFLICT(user_id, follower_id) DO NOTHING")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Query(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repo users) Unfollow(userId, followerId uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM followers WHERE user_id = $1 AND follower_id = $2")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Query(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repo users) FindFollowers(userId uint64) ([]models.User, error) {
	lines, err := repo.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM users u
		INNER JOIN followers f ON (u.id = f.follower_id)
		WHERE f.user_id = $1
	`, userId)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo users) FindFollowing(userId uint64) ([]models.User, error) {
	lines, err := repo.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM users u
		INNER JOIN followers f ON (u.id = f.user_id)
		WHERE f.follower_id = $1
	`, userId)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
