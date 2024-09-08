package main

import (
	"fmt"
	"time"

	"github.com/juanmanuel0963/notification_service/m/database"
	"github.com/juanmanuel0963/notification_service/m/models"
	"github.com/juanmanuel0963/notification_service/m/notifications"
)

func init() {
	//Initialize DB conn
	var responseErr = database.OpenDBConn()

	//Populate tables
	if responseErr == nil {
		database.PopulateDB()
	}
}

func main() {
	// Variables to store user input
	var user string
	var typeId int
	fmt.Println()
	// Request input from the user
	fmt.Print("Enter User: ")
	fmt.Scan(&user)

	fmt.Print("Enter TypeId. 1: STATUS | Type 2: NEWS | Type 3: MARKETING ")
	fmt.Scan(&typeId)

	NotificationsProcessing(user, typeId)
}

func NotificationsProcessing(user string, typeId int) error {

	// Populate the newRecord with the user inputs
	var newRecord = models.Notification{
		Stamp:   time.Now().Format("2006-01-02 15:04:05"), // Format the current time as a string
		User:    user,
		Subject: "subject", // Replace with actual subject if needed
		TypeID:  typeId,
	}

	fmt.Println()

	// Validate the set limits for the type of message and user are not surpassed yet
	resultError := notifications.ValidateFrequencyLimits(newRecord)

	if resultError != nil {
		fmt.Println(resultError)
		return resultError
	}

	// We are able to send a new message if the current messages count is lower than the maximum frequency limit
	resultError = notifications.SendNotification(newRecord)

	if resultError != nil {
		fmt.Println(resultError)
		return resultError
	} else {
		fmt.Println("Notification sent successfully.")
	}

	return nil
}
