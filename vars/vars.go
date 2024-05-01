package vars

import (
	"fmt"
	"os"
)

var (
	ConnectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		getUser(),
		getPass(),
		getHost(),
		getPort(),
		getDB())
)

// TODO implement vetVar functions
func getUser() string {
	var user = os.Getenv("LOCKBIN_USER")
	if user == "" {
		user = "lockbin"
	}
	return user
}

func getPass() string {
	var pass = os.Getenv("LOCKBIN_PASS")
	if pass == "" {
		pass = "Password123"
	}
	return pass
}

func getHost() string {
	var host = os.Getenv("LOCKBIN_HOST")
	if host == "" {
		host = "localhost"
	}
	return host
}

func getPort() string {
	var port = os.Getenv("LOCKBIN_PORT")
	if port == "" {
		port = "5432"
	}
	return port
}

func getDB() string {
	var db = os.Getenv("LOCKBIN_PDB")
	if db == "" {
		db = "lockbin"
	}
	return db
}
