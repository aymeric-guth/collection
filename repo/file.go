package repo

import (
	"fmt"
	"obsidian-server/types"

	"github.com/rqlite/gorqlite"
)

type IFile interface {
	CreateTable() error
	DropTable() error
	ReadAllPath() ([]string, error)
	// CreateMany([]*types.File) error
	CreateMany(<-chan *types.File) error
	FindByName(*types.File) ([]int64, error)
}

func (db *fileRepo) CreateTable() error {
	statements := make([]string, 0)
	statements = append(statements, db.queries["create-table"])
	_, err := db.Conn.Write(statements)
	if err != nil {
		return err
	}
	return nil
}

func (db *fileRepo) DropTable() error {
	statements := make([]string, 0)
	statements = append(statements, db.queries["drop-table"])
	_, err := db.Conn.Write(statements)
	if err != nil {
		return err
	}
	return nil
}

func (db *fileRepo) ReadAllPath() ([]string, error) {
	rows, err := db.Conn.QueryOne(db.queries["read-all-path"])
	if err != nil {
		return nil, err
	}

	rs := make([]string, rows.NumRows())
	for rows.Next() {
		var s string
		err := rows.Scan(&s)
		if err != nil {
			panic(err)
		}
		rs[rows.RowNumber()] = s
	}
	return rs, nil
}

func (db *fileRepo) CreateMany(ch <-chan *types.File, done chan bool) error {
	stmnt := make([]gorqlite.ParameterizedStatement, 0)
	for f := range ch {
		stmnt = append(stmnt, gorqlite.ParameterizedStatement{Query: db.queries["create-many"], Arguments: []interface{}{f.RelPath, f.Name, f.Extension}})
	}
	fmt.Println("Writing to DB")
	_, err := db.Conn.WriteParameterized(stmnt)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully commited to DB")
	done <- true
	return nil
}

func (db *fileRepo) FindByName(arg *types.File) ([]int64, error) {
	var query string
	var args []interface{}
	if arg.RelPath == "" {
		query = db.queries["find-by-name-extension"]
		args = []interface{}{arg.Name, arg.Extension}
	} else {
		query = db.queries["find-by-name-extension-path"]
		args = []interface{}{arg.Name, arg.Extension, arg.RelPath}
	}

	rows, err := db.Conn.QueryOneParameterized(
		gorqlite.ParameterizedStatement{
			Query:     query,
			Arguments: args,
		},
	)
	if err != nil {
		return nil, err
	}

	rs := make([]int64, rows.NumRows())
	for rows.Next() {
		id := &rs[rows.RowNumber()]
		err := rows.Scan(id)
		if err != nil {
			panic(err)
		}
	}
	return rs, nil
}
