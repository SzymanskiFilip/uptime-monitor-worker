package storage

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/SzymanskiFilip/uptime-monitoring-go/types"
)

func PersistRequest(r *http.Response, t time.Duration){
	stat := types.Statistic {
		URL: `http://localhost:3000`,
		Headers: "any",
		Status: int16(r.StatusCode),
		Success: true,
		ResponseTime: t.Milliseconds(),
	}

	sqlStatement := `INSERT INTO statistics (url, headers, status, success, response_time) 
	values ($1, $2, $3, $4, $5)`

	_, err := db.Exec(sqlStatement, stat.URL, stat.Headers, stat.Status, stat.Success, stat.ResponseTime)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("stat persisted")
}


func GetDomains() []types.URLStored{
	db := GetDB();

	data := []types.URLStored{}

	rows, err := db.Query(`
		SELECT * FROM urls
	`)
	if err != nil {
		log.Fatal(err)
	}

	var id string
	var domain string

	for rows.Next() {
		err := rows.Scan(&id, &domain)
		if err != nil {
			log.Fatal()
		}

		data = append(data, types.URLStored{Id: id, Domain: domain})
	}

	defer rows.Close()

	return data
}

//1 = success, 2 = already exists
func SaveDomain(domain string) int {
	_, err := db.Exec(`
	INSERT INTO urls (url) VALUES ($1)
	`, domain)

	if err != nil {
		if strings.Contains(err.Error(), `pq: duplicate key value violates unique constraint "urls_url_key"`) {
			fmt.Println("domain not persisted, already exists")
			return 2
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("domain persisted")
	return 1
}

func DeleteDomain(id string) bool {
	_, err := db.Exec(`
		DELETE FROM urls WHERE urls.id = $1
	`, id)
	fmt.Println(err)
	return err == nil
}