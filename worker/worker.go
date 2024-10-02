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

func StartPinging(){
	var adresses = getPingingAdresses()

	for _, item := range adresses {
		ping(item.Domain, item.Id)
	}
}

func ping(ad string, id string){
	for {
		time.Sleep(1 * time.Second)
		go performRequest(ad, id)
	}
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