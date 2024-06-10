package firebase

import (
    "context"
    "firebase.google.com/go"
    "firebase.google.com/go/messaging"
	"google.golang.org/api/option"
    "notification/common/configs"
    "log"
)


var Client *messaging.Client

func Init() error {
    // Set up Firebase
    opt := option.WithCredentialsFile(configs.Yaml.Firebase.GoogleApplicationCredentials)
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        log.Println("Failed to create Firebase app: %v", err)
        return err
    }

    // Initialize the messaging client
    Client, err = app.Messaging(context.Background())
    if err != nil {
        log.Println("Failed to create Firebase messaging client: %v", err)
        return err
    }

    return nil
}

