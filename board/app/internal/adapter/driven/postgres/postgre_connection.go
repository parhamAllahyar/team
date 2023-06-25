package postgres

import (
	"board/configs"
	"database/sql"
	"fmt"
	"strings"
	"text/template"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

const connString = "postgres://{{.User}}:{{.Pass}}@{{.Host}}:{{.Port}}/{{.DbName}}@sslmode=disable"

func buildConnectionStringOrPanic(cnf configs.PostgresDBConfig) string {

	sb := strings.Builder{}
	temp := template.Must(template.New("ConnString").Parse(connString))

	fmt.Println(temp)

	if err := temp.Execute(&sb, cnf); err != nil {
		panic(err)
	}

	return sb.String()
}

func NewPostgre(cnf configs.PostgresDBConfig) (*sql.DB, error) {

	conn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cnf.Host, cnf.Port, cnf.User, cnf.Pass, cnf.DbName)

	fmt.Println("connection is ", conn)

	db, err := sql.Open("postgres", conn)
	fmt.Println("err is ", err)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Print("Failed to ping the database")
		return db, fmt.Errorf("Could not ping the database %w", err)
	}

	db.SetConnMaxLifetime(time.Second)
	db.SetConnMaxIdleTime(30 * time.Second)

	return db, nil
}

func callGorm(db *sql.DB) *gorm.DB {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return gormDB
}
