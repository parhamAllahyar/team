package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sync"

)

var migrateOnce sync.Once = sync.Once{}

func Migrate(ctx context.Context, db *sql.DB) (err error) {

	fmt.Println("start migration")

	migrateOnce.Do(func() {
		
		fpath, err := filepath.Abs("../internal/adapter/driven/postgres/schema/customers_table.sql")

		if err != nil {
			return
		}

		fbyte, err := os.ReadFile(fpath)
		fmt.Println(fbyte)
		if err != nil {
			fmt.Println(err)
			return
		}

		if _, err := db.ExecContext(ctx, string(fbyte)); err != nil {
			fmt.Println(err)
			return
		}
	})

	fmt.Println("end migration")

	return
}
