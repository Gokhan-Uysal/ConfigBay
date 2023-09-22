package port

import "database/sql"

type (
	BaseRepo interface {
		Exec(*sql.Tx, string, ...interface{}) (sql.Result, error)
		Query(*sql.Tx, string, ...interface{}) (*sql.Rows, error)
		QueryRow(
			*sql.Tx,
			string,
			...interface{},
		) *sql.Row
		CloseRows(rows *sql.Rows)
		Begin() (*sql.Tx, error)
		CommitOrRollback(*sql.Tx, error) error
	}
)
