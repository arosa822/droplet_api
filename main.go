package main

import (
	"droplet_api/models"
	jsonUtil "droplet_api/utils"
	"encoding/json"
	"fmt"
)

func main() {
	const FILE_LOC = "./data/sample_json.json"
	const DEVICE_NAME = "water_pump"
	var devices models.Devices

	// read in JSON data as byte stream
	fileStream := jsonUtil.ReadJsonFile(FILE_LOC)
	json.Unmarshal(fileStream, &devices)

	// do some stuff to the data
	devices.DescribeData()
	devices.UpdateState(DEVICE_NAME, 1)
	devices.DescribeData()

	fmt.Println("### DEBUG ###")
	fmt.Println(devices)
	fmt.Println(len(devices.Devices))

	// update the json file
	// TODO: this should be within 'UpdateState' interface
	jsonUtil.WriteToFile(&devices, FILE_LOC)

}
