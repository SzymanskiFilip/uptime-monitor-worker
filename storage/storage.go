package storage

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SzymanskiFilip/uptime-monitoring-go/types"
)

func PersistRequest(r *http.Response){
	stat := types.Statistic {
		URL: `http://localhost:3000`,
		Headers: "any",
		Status: int16(r.StatusCode),
		Success: true,
		ResponseTime: 200,
	}

	sqlStatement := `INSERT INTO statistics (url, headers, status, success, response_time) 
	values ($1, $2, $3, $4, $5)`

	_, err := db.Exec(sqlStatement, stat.URL, stat.Headers, stat.Status, stat.Success, stat.ResponseTime)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("persisted")

}

