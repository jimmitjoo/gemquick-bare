package data

import (
	"database/sql"
	"fmt"
	"os"

	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper db2.Session

type Models struct {
}

func New(databasePool *sql.DB) Models {
	db = databasePool

	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":
		upper, _ = mysql.New(db)
	case "postgres", "postgresql":
		upper, _ = postgresql.New(db)
	default:
		// no database
	}

	return Models{}
}

func getInsertID(i db2.ID) int {

	if i == nil {
		return 0
	}

	idType := fmt.Sprintf("%T", i)

	if idType == "int64" {
		return int(i.(int64))
	} else if idType == "func() db.ID" {
		return getInsertID(i.(func() db2.ID)())
	}

	return int(i.(int))
}
