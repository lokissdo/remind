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
    // test(app)
    return nil
}


func test( app *firebase.App) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "dEjEeY-RQLiCKcm5Poh0Lo:APA91bGVvKwhqd9cii9nrER9LVAkmkrIlykXtjk2z8PVxqfAPCtgfwIxH-2gqJJJYmK2lP7Q44ee9UPvTFDVotnH3RAvZxXoqNUxYSkC5KjE-HA54N5kukzoG_CNw4xtivWvZMY_jnmg"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	log.Println("Successfully sent message:", response)
}

