package nutrition

import (
	"database/sql"

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

func (r *repository) AddNutritionInfo(dish string, size int, calories int) error {
	_, err := r.db.Exec(`INSERT INTO nutrition(dish, size, calories) VALUES ($1, $2, $3)`, dish, size, calories)
	return errors.Wrap(err, "exec query")
}

func (r *repository) GetSumCalories() (int, error) {
	row := r.db.QueryRow(`SELECT COALESCE(SUM(calories), 0) FROM nutrition`)
	if err := row.Err(); err != nil {
		return 0, errors.Wrap(err, "unable to do query")
	}
	var calories int
	if err := row.Scan(&calories); err != nil {
		return 0, errors.Wrap(err, "unable to scan sumCalories")
	}
	return calories, nil
}
