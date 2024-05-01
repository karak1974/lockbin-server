package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"lockbin_server/types"
	"lockbin_server/vars"
	"time"
)

// init connect to database
func GetRecord(uuid string) (types.Record, error) {
	db, err := sql.Open("postgres", vars.ConnectionString)
	if err != nil {
		return types.Record{}, err
	}
	defer db.Close()

	sqlStatement := `
		SELECT masterKey, unlockTime, deleteTime, message
		FROM lockbin
		WHERE uuid = $1`
	var masterKey string
	var unlockTime time.Time
	var deleteTime time.Time
	var message string
	err = db.QueryRow(sqlStatement, uuid).Scan(&masterKey, &unlockTime, &deleteTime, &message)
	if err != nil {
		return types.Record{}, err
	}

	record := types.Record{
		MasterKey:  masterKey,
		UnlockTime: unlockTime.Unix(),
		DeleteTime: deleteTime.Unix(),
		Message:    message,
	}
	return record, nil
}

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
