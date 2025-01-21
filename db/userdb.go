package db

import "github.com/xarick/golang-cassandra-example/models"

func CreateUser(user *models.User) error {
	return Session.Query(
		"INSERT INTO users (id, name, email) VALUES (?, ?, ?)",
		user.ID, user.Name, user.Email,
	).Exec()
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	iter := Session.Query("SELECT id, name, email FROM users").Iter()

	var user models.User
	for iter.Scan(&user.ID, &user.Name, &user.Email) {
		users = append(users, user)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := Session.Query(
		"SELECT id, name, email FROM users WHERE id = ?",
		id,
	).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id string, user *models.User) error {
	return Session.Query(
		"UPDATE users SET name = ?, email = ? WHERE id = ?",
		user.Name, user.Email, id,
	).Exec()
}

func DeleteUser(id string) error {
	return Session.Query(
		"DELETE FROM users WHERE id = ?",
		id,
	).Exec()
}
