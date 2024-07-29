package repository

import (
	"context"
	"fmt"
	"log"
	"remind/notification/pb"
	"remind/todo/common/configs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SendPendingReminders(ctx context.Context) error {
    reminders, err := GetNeedSentReminders(ctx)
    if err != nil {
        return err
    }

    // Connect to the notification service
    conn, err := grpc.NewClient(configs.Yaml.Notification.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))


    if err != nil {
        return err
    }
    defer conn.Close()

    // Create a client instance
    client := pb.NewNotificationServiceClient(conn)

    sucessReminders := []string{}
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

        sucessReminders = append(sucessReminders, reminder.ID)
        log.Printf("Notification sent successfully to user %s", reminder.UserID)

    }

    if len(sucessReminders) == 0 {
        return nil
    }
    UpdateBatchReminderSentStatus(ctx, sucessReminders)

    return nil
}

func CreateOrUpdateFCMToken(ctx context.Context, userID, token string) error {
     // Connect to the notification service
     conn, err := grpc.NewClient(configs.Yaml.Notification.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))


     if err != nil {
         return err
     }
     defer conn.Close()
 
     // Create a client instance
    client := pb.NewNotificationServiceClient(conn)
     // Prepare the request
     req := &pb.TokenRequest{
        UserId: userID,
        Token: token,
    }
    resp, err := client.AddOrUpdateToken(ctx, req)
    if err != nil {
        return err
    }

    if !resp.Success {
        fmt.Println("Failed to updated token %s", resp.ErrorMessage)
        return fmt.Errorf("Failed to add or update token for user %s", userID)
    }
    return nil
}