package repo

import (
	"database/sql"
	"github.com/Gokhan-Uysal/ConfigBay.git/internal/core/domain/common/errorx"

	"github.com/Gokhan-Uysal/ConfigBay.git/internal/lib/logger"
)

type (
	Scanner interface {
		Scan(dest ...any) error
	}

	baseRepo struct {
		db *sql.DB
	}
)

func newBaseRepo(db *sql.DB) (*baseRepo, error) {
	if db == nil {
		return nil, errorx.NilPointerErr{Item: "db"}
	}
	return &baseRepo{db: db}, nil
}

func (br baseRepo) Exec(tx *sql.Tx, query string, args ...any) (sql.Result, error) {
	var (
		result sql.Result
		err    error
	)

	if tx == nil {
		result, err = br.db.Exec(query, args...)
	} else {
		result, err = tx.Exec(query, args...)
	}

	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}
	return result, nil
}

func (br baseRepo) Query(tx *sql.Tx, query string, args ...any) (*sql.Rows, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if tx == nil {
		rows, err = br.db.Query(query, args...)
	} else {
		rows, err = tx.Query(query, args...)
	}

	if err != nil {
		logger.ERR.Println(err)
		return nil, err
	}
	return rows, nil
}

func (br baseRepo) QueryRow(
	tx *sql.Tx,
	query string,
	args ...any,
) *sql.Row {
	var (
		row *sql.Row
	)

	if tx == nil {
		row = br.db.QueryRow(query, args...)
	} else {
		row = tx.QueryRow(query, args...)
	}
	return row
}

func (br baseRepo) CloseRows(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		logger.ERR.Println("failed to close rows:", err)
		panic(err)
	}
}

func (br baseRepo) Begin() (*sql.Tx, error) {
	return br.db.Begin()
}

func (br baseRepo) CommitOrRollback(tx *sql.Tx, err error) error {
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			logger.ERR.Println("failed to rollback ", rollbackErr)
			return rollbackErr
		}
		return err
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		logger.ERR.Println("failed to commit ", commitErr)
		return commitErr
	}
	return nil
}
