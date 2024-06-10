package repository

import (
    "context"
    "log"
    "todo/firebase"
    "todo/model"
    "time"
)

func SendPendingReminders(ctx context.Context) error {
    reminders, err := GetNeedSentReminders(ctx)
    if err != nil {
        return err
    }

    for _, reminder := range reminders {
        // Here you need to get the user's FCM token; this is just a placeholder
        userToken := getUserToken(reminder.UserID)
        if userToken == "" {
            log.Printf("No token found for user %s", reminder.UserID)
            continue
        }

        err = firebase.SendNotification(ctx, userToken, "Reminder", reminder.Message)
        if err != nil {
            log.Printf("Failed to send notification for reminder %s: %v", reminder.ID, err)
            continue
        }

        // Update the reminder as sent (this part depends on your model)
        reminder.Sent = true
        _, err := UpdateReminderByID(reminder.ID, reminder)
        if err != nil {
            log.Printf("Failed to update reminder %s: %v", reminder.ID, err)
            continue
        }
    }

    return nil
}

func getUserToken(userID string) string {
    // Placeholder for getting the user token from the database or another service
    // This function should return the FCM token for the given user ID
    return "user-fcm-token"
}
