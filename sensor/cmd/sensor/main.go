package main

import (
	"fmt"
	"log"
	"net/http"
	i "sensor/cmd/sensor/injection"
	"sensor/pkg/db"
)

func main() {
	//Initialize data sources
	ds, err := db.InitializeDataSources()
	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	router, err := i.Injection(ds)
	if err != nil {
		log.Fatalf("Unable to initialize routes: %v\n", err)
	}

	fmt.Println("Rodando API")
	log.Fatal(http.ListenAndServe(":5001", router))

}
