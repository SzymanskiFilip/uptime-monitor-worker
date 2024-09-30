package worker

import (
	"fmt"
	"net/http"
	"time"
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
		go performRequest(ad)
	}
}

func performRequest(ad string) {
	fmt.Println("SENDING REQUEST " + time.Now().Format("15:04:05"))
	resp, err := client.Get(ad)
	if err != nil {
		//persistence
		fmt.Println("error")
	} else {
		//persistence
		fmt.Println("good " + resp.Status)
	}
}

func getPingingAdresses() []string{
	adresses := []string{"http://localhost:3000/api/service/status"}
	return adresses
}

