package queries

import (
	"database/sql"
	"time"
)


func ReadURLLatest(db *sql.DB, url string) (checked_at time.Time, is_up bool, err error) {
	err = db.QueryRow(
		`SELECT checked_at, is_up FROM checks WHERE url = $1 ORDER BY checked_at DESC LIMIT 1;`,
		url,
	).Scan(&checked_at, &is_up)

	return checked_at, is_up, err
}

func CheckURLResponseRates(db *sql.DB, url string) (min int, max int, median float64, err error) {
	err = db.QueryRow(
		`SELECT MIN(response_time), MAX(response_time), percentile_cont(0.5) WITHIN GROUP (ORDER BY response_time) FROM checks WHERE url = $1;`,
	url).Scan(&min, &max, &median)

	return min, max, median, err
}
// func ReadOutOfRangePeriods() {}
// func ReadDownHistory() {}
