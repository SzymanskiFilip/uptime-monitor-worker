package worker

import (
	"fmt"
	"time"
)

//pobrać adresy do pingowania
//zacząć goroutine który pinguje każdy url co 2min i zapisuje staty


func StartPinging(){
	var adresses = getPingingAdresses()

	for _, item := range adresses {
		go ping(item)
	}
}

func ping(ad string){
	for {
		fmt.Printf("Hello from item -> " + ad)
		time.Sleep(5 * time.Second)
	}
}

func getPingingAdresses() []string{
	adresses := []string{"http://localhost:3000"}
	return adresses
}

