package worker

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SzymanskiFilip/uptime-monitoring-go/storage"
	"github.com/google/uuid"
)

var client = &http.Client{
	Timeout: 15 * time.Second,
}

func StartPinging(){
	var adresses = getPingingAdresses()

	for _, item := range adresses {
		ping(item)
	}
}

func ping(ad string){
	for {
		time.Sleep(1 * time.Second)
		go performRequest(ad, uuid.New())
	}
}

func performRequest(ad string, id uuid.UUID) {
	fmt.Println("SENDING REQUEST " + id.String() + time.Now().Format("15:04:05"))
	now := time.Now()
	resp, err := client.Get(ad)
	if err != nil {
		//persistence
		fmt.Println("error")
	} else {
		//persistence
		fmt.Println("good " + resp.Status)

		elapsed := time.Since(now)
		storage.PersistRequest(resp, elapsed)
	}
}

func getPingingAdresses() []string{
	adresses := []string{"http://localhost:3000/api/service/status"}
	return adresses
}

