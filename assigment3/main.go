package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	// Set up Gin router
	router := gin.Default()

	// Serve static files from the web directory
	router.StaticFS("/web", http.Dir("./web"))

	// Route to display status
	router.GET("/status", getStatusHandler)

	// Start updating file periodically
	go updateFilePeriodically()

	// Run the server
	router.Run(":8080")
}

func getStatusHandler(c *gin.Context) {
	status := readStatusFromFile()

	// Determine status for water and wind
	waterStatus := getStatusStatus(status.Water)
	windStatus := getStatusStatus(status.Wind)

	// Prepare response
	response := gin.H{
		"status":       status,
		"water_status": waterStatus,
		"wind_status":  windStatus,
	}

	c.JSON(http.StatusOK, response)
}

func readStatusFromFile() Status {
	// Read status from JSON file
	file, err := ioutil.ReadFile("status.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	var status Status
	err = json.Unmarshal(file, &status)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
	}

	return status
}

func updateFilePeriodically() {
	for {
		// Generate random water and wind values
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		// Create status object
		status := Status{
			Water: water,
			Wind:  wind,
		}

		// Write status to JSON file
		file, err := json.MarshalIndent(status, "", "  ")
		if err != nil {
			fmt.Printf("Error marshalling JSON: %v\n", err)
		}

		err = ioutil.WriteFile("status.json", file, 0644)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		}

		// Sleep for 15 seconds
		time.Sleep(15 * time.Second)
	}
}

func getStatusStatus(value int) string {
	// Determine status for water and wind based on value
	if value < 5 {
		return "Aman"
	} else if value >= 6 && value <= 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}
