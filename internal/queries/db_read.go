package queries

import (
	"fmt"
	"database/sql"
	"time"
)

type OutOfRangeResult struct {
	Response_time float64 
	Checked_at time.Time
}

type SiteDownResult struct {
	Status_code int
	Is_up bool
	Checked_at time.Time
}

func ReadURLLatest(db *sql.DB, url string) (checked_at time.Time, is_up bool, err error) {
	err = db.QueryRow(
		`SELECT checked_at, is_up FROM checks WHERE url = $1 ORDER BY checked_at DESC LIMIT 1;`,
		url,
	).Scan(&checked_at, &is_up)

	return checked_at, is_up, err
}

func CheckURLResponseRates(db *sql.DB, url string) (min int, max int, median float64, err error) {
	err = db.QueryRow(
		`SELECT MIN(response_time), MAX(response_time), percentile_cont(0.5) 
		WITHIN GROUP (
			ORDER BY response_time
		) 
		FROM checks WHERE url = $1;`,
	url).Scan(&min, &max, &median)

	return min, max, median, err
}
func ReadOutOfRangePeriods(db *sql.DB, url string) (results []OutOfRangeResult, err error) {
	rows, err := db.Query(
		`SELECT response_time, checked_at 
		FROM checks where url = $1 AND checked_at >= CURRENT_DATE - INTERVAL '7 days' AND response_time > (
			SELECT percentile_cont(0.9) 
			WITHIN GROUP (
				ORDER BY response_time
			) 
			FROM checks WHERE url = $1
			)`,
		url)

		if err != nil {
			fmt.Printf("\nError: %v\n", err)
			return
		}
		
		defer rows.Close()

		for rows.Next() {
			var row OutOfRangeResult
			rows.Scan(&row.Response_time, &row.Checked_at)
			results = append(results, row)
		}
		
		return
}
func ReadDownHistory(db *sql.DB, url string) (results []SiteDownResult, err error) {
	rows, err := db.Query(
		`SELECT status_code, is_up, checked_at FROM (
			SELECT status_code, is_up, checked_at, LAG(is_up) OVER (ORDER BY checked_at) AS prev_is_up
			FROM checks WHERE url = $1
		) AS subquery
		WHERE (is_up = false and prev_is_up = true) OR (is_up = true and prev_is_up = false)
		`, url,
	)

	if err != nil {
		fmt.Printf("\nError: %v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var row SiteDownResult
		rows.Scan(&row.Status_code, &row.Is_up, &row.Checked_at)
		results = append(results, row)
	}
	
	return

}
