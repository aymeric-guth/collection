package repo

import (
	"github.com/nleof/goyesql"
	"github.com/rqlite/gorqlite"
)

type Repo struct {
	Conn    *gorqlite.Connection
	queries goyesql.Queries
}

var conn *gorqlite.Connection

type fileRepo struct {
	Repo
}

var File = new(fileRepo)

// var err error
// var f *os.File

func init() {
	// f, err = os.OpenFile("/tmp/deep_insights.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// 	panic(err)
	// }
	// gorqlite.TraceOn(f)
	// gorqlite.TraceOff()
	db, err := gorqlite.Open("http://localhost:4001/")
	if err != nil {
		panic(err)
	}
	conn = &db

	File.Conn = conn
	File.queries = goyesql.MustParseFile("sql/file.sql")
}

func Deinit() {
	conn.Close()
	// gorqlite.TraceOff()
	// f.Close()
}
