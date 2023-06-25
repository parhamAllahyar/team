package postgres

import (
	"customer/configs"
	"database/sql"
	"fmt"
	"strings"
	"text/template"
	"time"

	_ "github.com/lib/pq"

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

func NewPostgres(cnf configs.PostgresDBConfig) (*sql.DB, error) {

	conn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cnf.Host, cnf.Port, cnf.User, cnf.Pass, cnf.DbName)

	// connString := "postgres://" + cnf.User + ":" + cnf.Pass + "+@" + cnf.Host + ":" + cnf.Port + "/" + cnf.DbName

	// fmt.Println("connection", connString)

	db, err := sql.Open("postgres", conn)
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
