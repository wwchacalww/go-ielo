package db

import (
	"database/sql"
	"wwchacalww/go-psyc/domain/model"
)

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) Create(user model.UserInterface) error {
	stmt, err := u.db.Prepare(`INSERT INTO users (
		id,
		name,
		email,
		password,
		role,
		status
	) values ($1, $2, $3, $4, $5, $6)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		user.GetID(),
		user.GetName(),
		user.GetEmail(),
		user.GetPassword(),
		user.GetRole(),
		user.GetStatus(),
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDB) FindById(id string) (model.UserInterface, error) {
	var user model.User
	stmt, err := u.db.Prepare("select id, name, email, password, status, role, avatar_url from users where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status, &user.Role, &user.AvatarUrl)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDB) FindByEmail(email string) (model.UserInterface, error) {
	var user model.User
	var avatar sql.NullString
	stmt, err := u.db.Prepare("select id, name, email, password, status, role, avatar_url from users where email=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status, &user.Role, &avatar)
	user.AvatarUrl = avatar.String
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDB) List() ([]model.UserInterface, error) {
	var users []model.UserInterface
	rows, err := u.db.Query("SELECT id, name, email, password, role, status, avatar_url FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		var avatar sql.NullString
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.Status, &avatar)
		user.AvatarUrl = avatar.String
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u *UserDB) ChangePassword(id, pwd string) error {
	_, err := u.db.Exec("UPDATE users SET password=$1 WHERE id=$2", pwd, id)
	return err
}

func (u *UserDB) ChangeRole(id, role string) error {
	_, err := u.db.Exec("UPDATE users SET role=$2 WHERE id=$1", id, role)
	return err
}

func (u *UserDB) ChangeAvatarUrl(email, avatar_url string) error {
	_, err := u.db.Exec("UPDATE users SET avatar_url=$2 WHERE email=$1", email, avatar_url)
	return err
}
