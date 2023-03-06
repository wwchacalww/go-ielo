package db_test

import (
	"database/sql"
	"log"
	"testing"
	"wwchacalww/go-psyc/adapters/db"
	"wwchacalww/go-psyc/domain/model"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var DB *sql.DB

func setUp() {
	DB, _ = sql.Open("sqlite3", ":memory:")

	createTable(DB)
	createUsers(DB)
}

func createTable(db *sql.DB) {
	user_table := `CREATE TABLE users (
		"id" string,
		"name" string,
		"email" string,
		"password" string,
		"role" string,
		"status" boolean
	);
	`

	rt_table := `CREATE TABLE refresh_tokens (
		"id" INTEGER PRIMARY KEY,
		"user_id" string,
		"token" string,
		"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		"expired_at" TIMESTAMP
	);`
	stmt, err := db.Prepare(rt_table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
	stmt.Close()
	rtExec, err := db.Prepare(user_table)
	if err != nil {
		log.Fatal(err.Error())
	}
	rtExec.Exec()
	rtExec.Close()
}

func createUsers(db *sql.DB) {
	insert := `insert into users values
	("id-1", "Fulano", "fulano@gmail.com", "password", "role-test", 1),
	("id-2", "Siclano", "siclano@gmail.com", "password", "role-test", 1);`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestUserDB_Get(t *testing.T) {
	setUp()
	defer DB.Close()
	userDB := db.NewUserDB(DB)
	user, err := userDB.FindById("id-1")
	require.Nil(t, err)
	require.Equal(t, "Fulano", user.GetName())
}

func TestUserDB_GetEmail(t *testing.T) {
	setUp()
	defer DB.Close()
	userDB := db.NewUserDB(DB)
	user, err := userDB.FindByEmail("fulano@gmail.com")
	require.Nil(t, err)
	require.Equal(t, "Fulano", user.GetName())
}

func TestUserDB_List(t *testing.T) {
	setUp()
	defer DB.Close()
	userDB := db.NewUserDB(DB)
	users, err := userDB.List()
	require.Nil(t, err)
	require.Equal(t, len(users), 2)
}

func TestUserDB_Create(t *testing.T) {
	setUp()
	defer DB.Close()
	userDB := db.NewUserDB(DB)

	user := model.NewUser()
	user.Name = "Beltrano"
	user.Email = "beltrano@mail.com"
	user.Password = "12345"
	user.Role = "atendente"
	user.Status = true

	err := userDB.Create(user)
	require.Nil(t, err)
}

func TestAuthDB_CreateRT(t *testing.T) {
	rt := model.NewRefreshToken()
	acc := model.Account{
		Name:   "Fulano",
		UserID: "id-1",
		Email:  "fulano@gmail.com",
		Role:   "role-test",
	}
	rt.Account = acc
	setUp()
	defer DB.Close()
	authDb := db.NewAuthDB(DB)
	refreshToken, err := authDb.CreateRefreshToken(rt)
	log.Println(rt.GetExpiredAt(), rt.GetExpiredAt())
	require.Nil(t, err)
	require.Equal(t, rt.GetToken(), refreshToken.GetToken())
	require.Equal(t, rt.GetAccount().Email, refreshToken.GetAccount().Email)

	log.Println(refreshToken)
}

func TestAuthDB_CheckAcc(t *testing.T) {
	setUp()
	defer DB.Close()
	authDb := db.NewAuthDB(DB)
	check, err := authDb.CheckAccount("fulano@gmail.com")
	require.Nil(t, err)
	require.Equal(t, check.GetEmail(), "fulano@gmail.com")
}
