package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"lockbin_server/types"
	"lockbin_server/vars"
	"time"
)

// init connect to database

func CreateRecord(record types.Record) (string, error) {
	db, err := sql.Open("postgres", vars.ConnectionString)
	if err != nil {
		return "", err
	}
	defer db.Close()

	sqlStatement := `
        INSERT INTO lockbin (uuid, masterKey, unlockTime, deleteTime, message)
        VALUES (gen_random_uuid (), $1, $2, $3, $4)
        RETURNING uuid`

	var uuid string
	err = db.QueryRow(sqlStatement,
		record.MasterKey,
		time.Unix(record.UnlockTime, 0),
		time.Unix(record.DeleteTime, 0),
		record.Message,
	).Scan(&uuid)
	if err != nil {
		return "", err
	}

	return uuid, nil
}
