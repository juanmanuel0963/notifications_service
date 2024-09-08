package main

import (
	"fmt"

	"github.com/juanmanuel0963/notification_service/m/database"
	"github.com/juanmanuel0963/notification_service/m/notifications"
)

func init() {

	//Initialize DB conn
	resultErr := database.OpenDBConn()

	if resultErr != nil {
		fmt.Println("Failed to connect to database. " + resultErr.Error())
	} /* else {
		fmt.Println("Connection successfull to database.")
	}
	*/
	//Populate tables
	resultErr = database.PopulateDB()

	if resultErr != nil {
		fmt.Println("Failed to populate the database. " + resultErr.Error())
	} /* else {
		fmt.Println("Database populated successfully.")
	}
	*/
}

func main() {
	// Variables to store user input
	var user string
	var typeId int
	fmt.Println()
	// Request input from the user
	fmt.Print("Enter User: ")
	fmt.Scan(&user)

	fmt.Print("Enter TypeId: 1-STATUS | 2-NEWS | 3-MARKETING ")
	fmt.Scan(&typeId)

	fmt.Println()

	resultError := notifications.Send(user, typeId, "message")

	if resultError != nil {
		fmt.Println(resultError)
	} else {
		fmt.Println("Notification sent successfully.")
	}
}
