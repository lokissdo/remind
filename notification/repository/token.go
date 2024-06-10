package repository

import (
    "context"
    "notification/database"
	"notification/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// AddOrUpdateFCMToken adds or updates an FCM token in the database.
func AddOrUpdateFCMToken(ctx context.Context, userID, token string) error {
    filter := bson.M{"user_id": userID}
    update := bson.M{"$set": bson.M{"token": token}}

    _, err := database.FCMTokenCollection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
    return err
}

func GetFCMTokenByUserID(ctx context.Context, userID string) (string, error) {
	var fcmToken model.FCMToken 
	filter := bson.M{"user_id": userID}
	err := database.FCMTokenCollection.FindOne(ctx, filter).Decode(&fcmToken)
	return fcmToken.Token, err
}