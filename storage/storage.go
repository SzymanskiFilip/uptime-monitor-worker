package storage

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/SzymanskiFilip/uptime-monitoring-go/types"
	"github.com/google/uuid"
)

func PersistRequest(r *http.Response, t time.Duration, address string, id string){

	var headers[] string

	for key, values := range r.Header {
		header := fmt.Sprintf("%s: %s", key, strings.Join(values, ", "))
		headers = append(headers, header)
	}

	
	status, err := strconv.Atoi(r.Status[:3])
	if err != nil {
		log.Fatal(err)
	}

	uuid, err := uuid.Parse(id); if err != nil {
		log.Fatal(err)
	}

	stat := types.Statistic {
		Id: uuid,
		Headers: strings.Join(headers, "\n"),
		Success: status >= 200,
		ResponseTime: t.Milliseconds(),
		SavedAt: time.Now(),
	}

	sqlStatement := `INSERT INTO statistics (url_id, headers, success, response_time, saved_at) 
	values ($1, $2, $3, $4, $5)`

	_, error := db.Exec(sqlStatement, stat.Id, stat.Headers, stat.Success, stat.ResponseTime, stat.SavedAt)
	if error != nil {
		log.Fatal(error)
	}

	defer r.Body.Close()

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
			return 2
		} else {
			log.Fatal(err)
		}
	}
	return 1
}

func DeleteDomain(id string) bool {
	_, err := db.Exec(`
		DELETE FROM urls WHERE urls.id = $1
	`, id)
	return err == nil
}