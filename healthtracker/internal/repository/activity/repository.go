package activity

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

func (r *repository) AddActivityInfo(activity string, duration time.Duration, calories int) error {
	_, err := r.db.Exec(`INSERT INTO activity(activity, duration, calories) VALUES ($1, $2, $3)`, activity, duration.Nanoseconds(), calories)
	return errors.Wrap(err, "exec query")
}

func (r *repository) GetSumDuration() (time.Duration, error) {
	row := r.db.QueryRow(`SELECT COALESCE(SUM(duration), 0) FROM activity`)
	if err := row.Err(); err != nil {
		return time.Duration(0), errors.Wrap(err, "unable to do query")
	}
	var activityTime int64
	if err := row.Scan(&activityTime); err != nil {
		return time.Duration(0), errors.Wrap(err, "unable to scan activityTime")
	}
	return time.Duration(activityTime), nil
}

func (r *repository) GetSumCalories() (int, error) {
	row := r.db.QueryRow(`SELECT COALESCE(SUM(calories), 0) FROM activity`)
	if err := row.Err(); err != nil {
		return 0, errors.Wrap(err, "unable to do query")
	}
	var calories int
	if err := row.Scan(&calories); err != nil {
		return 0, errors.Wrap(err, "unable to scan sumCalories")
	}
	return calories, nil
}
