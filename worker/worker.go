package worker

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SzymanskiFilip/uptime-monitoring-go/storage"
	"github.com/SzymanskiFilip/uptime-monitoring-go/types"
)

var client = &http.Client{
	Timeout: 15 * time.Second,
}

var addresses = []types.URLStored{}

func StartPinging(){
	for {
		loopOverAdresses()
		time.Sleep(1 * time.Second)
	}
}

func loopOverAdresses(){
	for _, value := range addresses {
		go performRequest(value.Domain, value.Id)
	}
}

func UpdateAddresses(){
	addresses = getPingingAdresses()
}

func performRequest(ad string, id string) {
	fmt.Println("SENDING REQUEST " + id + time.Now().Format("15:04:05"))
	now := time.Now()
	resp, err := client.Get(ad)
	if err != nil {
		//persistence
		fmt.Println("error")
	} else {
		//persistence
		fmt.Println("good " + resp.Status)

		elapsed := time.Since(now)
		storage.PersistRequest(resp, elapsed, ad, id)
	}
}

func getPingingAdresses() []types.URLStored{
	adresses := storage.GetDomains()
	return adresses
}

//http://localhost:3000/api/service/status