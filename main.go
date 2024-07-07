package main

import (
	"log"

	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	power, err := IsPowerConnected()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Power connected:", power)
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		newPower, err := IsPowerConnected()
		if err != nil {
			log.Fatalln(err)
		}
		if newPower != power {
			power = newPower
			log.Println("Power connected:", power)
			postPowerStatus(power)
		}
	}
}

func postPowerStatus(power bool) {
	powerStatus := "battery"
	if power {
		powerStatus = "charger"
	}
	payload := map[string]string{"power": powerStatus}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln("Unable to marshal JSON", err)
	}

	resp, err := http.Post(config.Url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Failed to post power status", err)
	} else {
		defer resp.Body.Close()
		log.Println("Response code:", resp.StatusCode)
	}
}
