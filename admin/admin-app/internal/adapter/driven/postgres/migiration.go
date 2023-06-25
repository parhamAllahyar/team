package postgres

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"sync"
	"fmt"

)

var migrateOnce sync.Once = sync.Once{}

func Migrate(ctx context.Context, db *sql.DB) (err error) {

	migrateOnce.Do(func() {
		fpath, err := filepath.Abs(filepath.Join("../assets", "tables.sql"))
		if err != nil {
			return
		}

		fbyte, err := os.ReadFile(fpath)
		if err != nil {
			return
		}

		if _, err := db.ExecContext(ctx, string(fbyte)); err != nil {
			return
		}
	})

	fmt.Println("databse migrated")

	return
}
