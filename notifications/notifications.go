package notifications

import (
	"errors"
	"fmt"
	"time"

	"github.com/juanmanuel0963/notification_service/m/database"
	"github.com/juanmanuel0963/notification_service/m/models"
	"github.com/juanmanuel0963/notification_service/m/types"
)

func SendNotification(newRecord models.Notification) error {

	// ---------- Populate notifications table ----------
	var notificationsList = []models.Notification{
		newRecord,
	}

	// Create new entry in the notifications table
	result := database.DB.Create(&notificationsList)
	if result.Error != nil {
		return errors.New("Error adding entry into notifications table: " + result.Error.Error())
	}

	//TODO: Send notification by email or sms

	return nil
}

// Validate the set limits for the type of message and user are not surpassed yet
func ValidateFrequencyLimits(newRecord models.Notification) error {

	//Get the type of message configuration
	resultType, resultError := types.GetTypeById(newRecord.TypeID)

	if resultError != nil {
		return resultError
	}

	//Get the count of messages already sent for the user and type of message
	resultCount, resultError := GetNotificationsCount(newRecord.User, resultType.Window, newRecord.TypeID)

	if resultError != nil {
		return resultError
	}

	//fmt.Println(resultCount)

	//If we reached the limit number for the user and type of message
	if int(resultCount) >= resultType.Frequency {
		print := fmt.Sprintf("The set limit of %d messages per %s for the %s type was surpassed for the user %s.", resultType.Frequency, resultType.Window, resultType.Name, newRecord.User)
		return errors.New(print)
	}

	return nil
}
func GetNotificationsCount(User string, Window string, TypeId int) (int64, error) {
	var count int64

	// Get the current time
	currentTime := time.Now()

	// Define the time condition based on the window parameter
	var condition string
	var args []interface{}

	switch Window {
	case "MINUTE":
		// Get the start of the current minute
		startOfMinute := currentTime.Truncate(time.Minute).Format("2006-01-02 15:04:05")
		condition = "stamp > ? AND user = ? AND type_id = ?"
		args = append(args, startOfMinute, User, TypeId)

	case "HOUR":
		// Get the start of the current hour
		startOfHour := currentTime.Truncate(time.Hour).Format("2006-01-02 15:04:05")
		condition = "stamp > ? AND user = ? AND type_id = ?"
		args = append(args, startOfHour, User, TypeId)

	case "DAY":
		// Set the start time to midnight of the current day
		startOfDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Format("2006-01-02 15:04:05")
		condition = "stamp > ? AND user = ? AND type_id = ?"
		args = append(args, startOfDay, User, TypeId)

	default:
		return 0, errors.New("invalid window parameter: must be 'MINUTE', 'HOUR', or 'DAY'")
	}

	//fmt.Println(condition)
	//fmt.Println(args)

	// Query the count of notifications based on the condition
	result := database.DB.Model(&models.Notification{}).Where(condition, args...).Count(&count)
	if result.Error != nil {
		return 0, errors.New("Error counting notifications: " + result.Error.Error())
	}

	return count, nil
}
