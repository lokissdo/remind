package firebase

import (
    "context"
    "firebase.google.com/go/messaging"
)

func SendNotification(ctx context.Context, token string, title string, body string) error {
    message := &messaging.Message{
        Token: token,
        Notification: &messaging.Notification{
            Title: title,
            Body:  body,
        },
    }

    // Send a message to the device corresponding to the provided registration token
    _, err := Client.Send(ctx, message)
    if err != nil {
        return err
    }

    return nil
}
