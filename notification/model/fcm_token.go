package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type FCMToken struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    UserID   string             `bson:"user_id"`
    Token    string             `bson:"token"`
}
