package main

import (
	"droplet_api/models"
	jsonUtil "droplet_api/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const FILE_LOC = "./data/sample_json.json"

var devices models.Devices

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

func returnAllDevices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllDevices")
	// read in JSON data as byte stream
	fileStream := jsonUtil.ReadJsonFile(FILE_LOC)
	json.Unmarshal(fileStream, &devices)
	json.NewEncoder(w).Encode(devices)

	devices.DescribeData()
}

func returnSpecifiedDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fileStream := jsonUtil.ReadJsonFile(FILE_LOC)
	json.Unmarshal(fileStream, &devices)
	json.NewEncoder(w).Encode(devices.GetDevice(key))

}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllDevices)
	myRouter.HandleFunc("/devices/{id}", returnSpecifiedDevice)
	log.Fatal(http.ListenAndServe(":1000", myRouter))
}

func main() {
	//const DEVICE_NAME = "water_pump"

	// do some stuff to the data
	//devices.DescribeData()
	//devices.UpdateState(DEVICE_NAME, 1)
	//devices.DescribeData()

	//fmt.Println("### DEBUG ###")
	//fmt.Println(devices)
	//fmt.Println(len(devices.Devices))

	// update the json file
	// TODO: this should be within 'UpdateState' interface
	//jsonUtil.WriteToFile(&devices, FILE_LOC)

	// API start here
	handleRequests()

}
