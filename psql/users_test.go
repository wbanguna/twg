package psql

import (
	"database/sql"
	"os"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	// 00. flag.Parse() if you need flags
	// 1. Setup anything you need
	exitCode := m.Run() // have to do m.Run to run tests when you have TestMain
	os.Exit(exitCode)
}

func TestUserStore_Find(t *testing.T) {
	const (
		dropDB          = `DROP DATABASE IF EXIST test_user_store;`
		createDB        = `CREATE DATABASE test_user_store;`
		createUserTable = `CREATE TABLE users (
							id SERIAL PRIMARY KEY,
							name TEXT,
							email TEST UNIQUE NOT NULL
							);`
	)

	psql, err := sql.Open("postgres", "host=localhost port=5432 user=jon sslmode=disable")
	if err != nil {
		t.Fatalf("sql.Open() err = %s", err)
	}
	defer psql.Close()

	_, err = psql.Exec(dropDB)
	if err != nil {
		t.Fatalf("psql.Exec() err = %s", err)
	}

	_, err = psql.Exec(createDB)
	if err != nil {
		t.Fatalf("psql.Exec() err = %s", err)
	}

	// teardown
	defer func() {
		_, err = psql.Exec(dropDB)
		if err != nil {
			t.Errorf("psql.Exec() err = %s", err)
		}
	}()

	db, err := sql.Open("postgres", "host=localhost port=5432 user=jon sslmode=disable dbname=test_user_store")
	if err != nil {
		t.Fatalf("sql.Open() err = %s", err)
	}
	defer db.Close()
	_, err = db.Exec(createUserTable)
	if err != nil {
		t.Fatalf("db.Exec() err = %s", err)
	}

	us := &UserStore{
		sql: db,
	}
	t.Run("Find", testUserStore_Find(us))
	t.Run("Create", testUserStore_Find(us))
	t.Run("Delete", testUserStore_Find(us))
	t.Run("Subscribe", testUserStore_Find(us))
	// teardown
}

func testUserStore_Find(us *UserStore) func(t *testing.T) {
	return func(t *testing.T) {
		jon := &User{
			Name:  "Jon Long",
			Email: "jonathan.d.long@maersk.com",
		}
		err := us.Create(jon)
		if err != nil {
			t.Errorf("us.Create() err = %s", err)
		}
		defer func() {
			err := us.Delete(jon.ID)
			if err != nil {
				t.Errorf("us.Delete() err = %s", err)
			}
		}()

		tests := []struct {
			name    string
			id      int
			want    *User
			wantErr error
		}{
			{"Found", jon.ID, jon, nil},
			{"NotFound", -1, nil, ErrNotFound},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := us.Find(tt.id)
				if err != tt.wantErr {
					t.Errorf("us.Find() err = %s", err)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("us.Find() err = %+v, want %+v", got, tt.want)
				}
			})
		}
	}
}
