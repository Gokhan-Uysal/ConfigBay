package port

import "database/sql"

type (
	BaseRepo interface {
		Exec(*sql.Tx, string, ...any) (sql.Result, error)
		Query(*sql.Tx, string, ...any) (*sql.Rows, error)
		QueryRow(
			*sql.Tx,
			string,
			...any,
		) *sql.Row
		CloseRows(rows *sql.Rows)
		Begin() (*sql.Tx, error)
		CommitOrRollback(*sql.Tx, error) error
	}
)
