package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
	"github.com/richxcame/alertd/admin"
	"github.com/richxcame/alertd/server"
)

// Define a global map to track the last alert time for each server
var lastAlertTimes = make(map[string]time.Time)

// Define a cooldown duration
const alertCooldown = 30 * time.Minute

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}
}

func main() {
	// Get all servers and admins
	servers, err := server.GetAll("servers.json")
	if err != nil {
		log.Fatalf("failed to get all servers: %v", err)
	}
	admins, err := admin.GetAll("admins.json")
	if err != nil {
		log.Fatalf("failed to get all admins: %v", err)
	}

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf("failed to create new scheduler: %v", err)
	}
	scheduler.NewJob(gocron.DurationJob(30*time.Second), gocron.NewTask(task, servers, admins))
	scheduler.Start()

	// Block until shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Perform cleanup and stop scheduler
	if err := scheduler.Shutdown(); err != nil {
		log.Fatalf("failed to shutdown scheduler: %v", err)
	}
	log.Println("Scheduler stopped. Exiting application.")

}

// Define a task function that will be executed by the scheduler
// The task function will check the status of all servers
func task(servers server.Servers, admins admin.Admins) {
	// Iterate over all servers and check their status
	// If a server is not OK, send an alert to all system administrators
	for _, server := range servers {
		systemsInfo, err := server.GetSystemsStatus()
		if err != nil {
			log.Printf("failed to get systems status for %s: %v", server.Host, err)
			for _, admin := range admins {
				if time.Since(lastAlertTimes[server.Host]) < alertCooldown {
					log.Printf("Alert for %s is in cooldown. Skipping SMS.", server.Host)
					continue
				}

				msg := fmt.Sprintf("Hello, %s! The server %v didn't give response", admin.Fullname, server.Host)
				if err := admin.SendSMS(msg); err != nil {
					log.Printf("failed to send sms to %s: %v", admin.Phone, err)
				}
			}
			continue
		}

		if !systemsInfo.IsOK() {
			for _, admin := range admins {
				// Check if the server is within the cooldown period
				if time.Since(lastAlertTimes[server.Host]) < alertCooldown {
					log.Printf("Alert for %s is in cooldown. Skipping SMS.", server.Host)
					continue
				}

				msg := fmt.Sprintf("Hello, %s! The server %s is not OK", admin.Fullname, server.Host)
				if err := admin.SendSMS(msg); err != nil {
					log.Printf("failed to send sms to %s: %v", admin.Phone, err)
				}
			}
		}
	}

}
