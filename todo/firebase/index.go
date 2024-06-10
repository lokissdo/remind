package firebase

import (
    "context"
    "firebase.google.com/go"
    "firebase.google.com/go/messaging"
	"google.golang.org/api/option"
    "os"
)


var Client *messaging.Client

func InitFirebase() error {
    // Set up Firebase
    opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        return err
    }

    // Initialize the messaging client
    Client, err = app.Messaging(context.Background())
    if err != nil {
        return err
    }

    return nil
}

