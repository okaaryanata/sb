package model

import (
	"question_2/db"
	"time"
)

// LogHistory - struct for table log_history in DB
type LogHistory struct {
	SearchWord string
	Pagination int64
	Success    bool
	CreatedAt  string
}

func (l LogHistory) stmntInsert() string {
	stmnt := `INSERT INTO log_history (search_word, pagination, success, created_at)
					 VALUES (?,?,?,?)`
	return stmnt
}

// Insert - function to insert log history to DB
func (l LogHistory) Insert(db db.Connection) (int64, error) {
	db.Begin()
	stmnt := l.stmntInsert()
	now := time.Now().Format("2006-01-02 15:04:05")

	id, err := db.Exec(
		stmnt,
		l.SearchWord,
		l.Pagination,
		l.Success,
		now,
	)
	if err != nil {
		db.Tx.Rollback()
		return 0, err
	}
	db.Commit()
	return int64(id), nil
}
