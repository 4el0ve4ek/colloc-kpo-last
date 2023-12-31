package sleep

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *sql.DB
}

func (r *repository) AddSleepInfo(duration time.Duration) error {
	_, err := r.db.Exec(`INSERT INTO sleep(duration) VALUES ($1)`, duration.Nanoseconds())
	return errors.Wrap(err, "exec query")
}

func (r *repository) GetSumTime() (time.Duration, error) {
	row := r.db.QueryRow(`SELECT COALESCE(SUM(duration), 0) FROM sleep`)
	if err := row.Err(); err != nil {
		return 0, errors.Wrap(err, "unable to do query")
	}
	var duration int64
	if err := row.Scan(&duration); err != nil {
		return 0, errors.Wrap(err, "unable to scan activityTime")
	}
	return time.Duration(duration), nil
}
