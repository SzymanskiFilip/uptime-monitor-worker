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


type ResponseTimeRow struct {
	Date time.Time `json:"date"`
	Avg float64 `json:"avg"`
}

func GetDailyResponseTimeAverage(id string) ([]ResponseTimeRow, error) {
	rows, err := db.Query(`
	SELECT DATE(saved_at), AVG(response_time) FROM statistics
	where url_id = $1 and saved_at >= NOW() - INTERVAL '1 month'
	group by DATE(saved_at)
	order by date
	`, id)

	if err != nil {
		fmt.Print(err)
	}


	defer rows.Close()

	data := []ResponseTimeRow{}

	var date time.Time
	var avg float64

	for rows.Next() {
		err := rows.Scan(&date, &avg)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, ResponseTimeRow{Date: date, Avg: avg})
	}

	return data, nil
}

func GetDomainById(i string) (idReturn string, urlReturn string){

	var url string
	var id string

	err := db.QueryRow(`
	SELECT * FROM urls where urls.id = $1
	`, i).Scan(&id, &url)
	if err != nil {
		log.Fatal(err)
	}

	return id, url
}

func GetMaxAndMinRespTime(id string) (min types.StatisticStored, max types.StatisticStored) {

	minStat := types.StatisticStored{}
	maxStat := types.StatisticStored{}

	err := db.QueryRow(`
	SELECT *
		FROM statistics
		where success = true and url_id = $1
		ORDER BY response_time asc
	`, id).Scan(&minStat.Id, &minStat.URL_ID, &minStat.Headers, &minStat.Success, &minStat.ResponseTime, &minStat.SavedAt); if err != nil {
		log.Fatal(err)
	}

	er := db.QueryRow(`
		SELECT *
		FROM statistics
		where success = true and url_id = $1
		ORDER BY response_time desc
	`, id).Scan(&maxStat.Id, &maxStat.URL_ID, &maxStat.Headers, &maxStat.Success, &maxStat.ResponseTime, &maxStat.SavedAt); if er != nil {
		log.Fatal(err)
	}

	return minStat, maxStat
}

type ResponseTimeUnionRow struct {
	Date time.Time `json:"date"`
	Avg float64 `json:"avg"`
	Period string `json:"period"`
}

func GetPrevWeeks(id string) (first []ResponseTimeRow, second []ResponseTimeRow) {
	rows, err := db.Query(`
	SELECT DATE(saved_at) AS date, AVG(response_time) AS avg_response_time, 'last7' AS period
		FROM statistics
		WHERE url_id = $1
		  AND saved_at >= NOW() - INTERVAL '7 days'
		GROUP BY DATE(saved_at)
		UNION ALL
		SELECT DATE(saved_at) AS date, AVG(response_time) AS avg_response_time, 'last14to7' AS period
		FROM statistics
		WHERE url_id = $1
		  AND saved_at < NOW() - INTERVAL '7 days'
		  AND saved_at >= NOW() - INTERVAL '14 days'
		GROUP BY DATE(saved_at)
		ORDER BY date, period;
	`, id)

	if err != nil {
		fmt.Print(err)
	}


	data := []ResponseTimeUnionRow{}


	var date time.Time
	var avg float64
	var period string

	for rows.Next() {
		err := rows.Scan(&date, &avg, &period)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, ResponseTimeUnionRow{Date: date, Avg: avg, Period: period})
	}

	firstList := []ResponseTimeRow{}
	secondList := []ResponseTimeRow{}


	for _, value := range data {
		if value.Period == "last7" {
			firstList = append(firstList, ResponseTimeRow{Avg: value.Avg, Date: value.Date})
		} else {
			secondList = append(secondList, ResponseTimeRow{Avg: value.Avg, Date: value.Date})
		}
	}

	return firstList, secondList
}

func GetOutages(id string) []types.StatisticStored {
	rows, err := db.Query(`
	SELECT * FROM statistics
	WHERE success = false and url_id = $1
	`, id); if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	data := []types.StatisticStored{}

	stat := types.StatisticStored{}
	for rows.Next() {
		err := rows.Scan(&stat.Id, &stat.URL_ID, &stat.Headers, &stat.Success, &stat.ResponseTime, &stat.SavedAt)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, stat)
	}

	return data
}