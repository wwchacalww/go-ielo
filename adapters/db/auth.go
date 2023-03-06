package db

import (
	"database/sql"
	"fmt"
	"time"
	"wwchacalww/go-psyc/domain/model"
)

type AuthDB struct {
	db *sql.DB
}

func NewAuthDB(db *sql.DB) *AuthDB {
	return &AuthDB{db: db}
}

func (a *AuthDB) CreateRefreshToken(rt model.RefreshTokenInterface) (model.RefreshTokenInterface, error) {
	_, err := a.db.Exec("DELETE FROM refresh_tokens WHERE user_id=$1", rt.GetAccount().UserID)
	if err != nil {
		return nil, err
	}
	stmt, err := a.db.Prepare(`INSERT INTO refresh_tokens (
		user_id,
		token,
		expired_at
	) values ($1, $2, $3)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		rt.GetAccount().UserID,
		rt.GetToken(),
		rt.GetExpiredAt(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return rt, nil
}

func (a *AuthDB) CheckAccount(email string) (model.UserInterface, error) {
	var user model.User
	stmt, err := a.db.Prepare("select id, name, email, password, status, role from users where email=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *AuthDB) CheckRefreshToken(user_id, refresh_token string) error {
	var ID int
	var Exp time.Time
	a.db.QueryRow("SELECT id, expired_at FROM refresh_tokens WHERE token = $1 AND user_id = $2", refresh_token, user_id).Scan(
		&ID,
		&Exp,
	)

	if ID == 0 {
		return fmt.Errorf("Token invalid")
	}

	if Exp.Unix() < time.Now().Unix() {
		a.db.QueryRow("DELETE FROM refresh_tokens WHERE id=$1", ID)
		return fmt.Errorf("Token expired")
	}

	return nil
}
