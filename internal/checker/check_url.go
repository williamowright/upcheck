package checker

import (
    "database/sql"
    "fmt"
    "net/http"
    "time"
)

func CheckURL(db *sql.DB, url string) {

	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	var statusCode int
	var isUp bool

	if err != nil {
		fmt.Printf("DOWN - %s - Error: %v\n", url, err)
		statusCode = 0
		isUp = false
	} else {
		defer resp.Body.Close()
		statusCode = resp.StatusCode
		isUp = resp.StatusCode < 400
		fmt.Printf("\nUP - %s - Status: %d - Time: %v\n", url, resp.StatusCode, duration)
	}

	_, dbErr := db.Exec(
		`INSERT INTO checks (url, status_code, response_time, is_up, checked_at) VALUES ($1, $2, $3, $4, $5)`,
		url,
		statusCode,
		duration.Milliseconds(),
		isUp,
		time.Now(),
	)
	
	if dbErr != nil {
		fmt.Printf("DB error: %v\n", dbErr)
	}
}