package repository

import (
	"context"
	"log"
	"todo/common/configs"
	"todo/model"
	pb "todo/pb"

	"google.golang.org/grpc"
)

func SendPendingReminders(ctx context.Context) error {
    reminders, err := GetNeedSentReminders(ctx)
    if err != nil {
        return err
    }

    // Connect to the notification service
    conn, err := grpc.NewClient(configs.Yaml.Notification.URL, grpc.WithInsecure())


    if err != nil {
        return err
    }
    defer conn.Close()

    // Create a client instance
    client := pb.NewNotificationServiceClient(conn)

    sucessReminders := []model.Reminder{}
    for _, reminder := range reminders {
        // Prepare the request
        req := &pb.MessageRequest{
            UserId: reminder.UserID,
            Title:  "Reminder!",
            Body:  reminder.Message + "\n "+ reminder.Start.String() + " is the time to start!+ ",
        }

        // Call the notification service
        _, err := client.SendMessage(ctx, req)
        if err != nil {
            log.Printf("Failed to send notification for user %s: %v", reminder.UserID, err)
            continue
        }

        sucessReminders = append(sucessReminders, reminder)
        log.Printf("Notification sent successfully to user %s", reminder.UserID)

    }

    if len(sucessReminders) == 0 {
        return nil
    }
    UpdateBatchReminderSentStatus(ctx, sucessReminders)

    return nil
}

