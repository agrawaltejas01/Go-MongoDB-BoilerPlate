package userRepo

import (
	"shive-app/database"
	"shive-app/database/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userModel *mongo.Collection

func init() {
	userModel, _ = database.GetCollection("users")
}

func findOne(query bson.M, projection bson.M) (models.User, error) {
	result := database.FindOne(userModel, query, projection)

	var user models.User
	err := result.Decode(&user)
	return user, err
}

func AddUser(user models.User) (primitive.ObjectID, error) {
	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	return database.InsertOne(userModel, user)
}

func FindByEmailOrUserName(email string, userName string) (models.User, error) {
	query := []interface{}{}

	if email != "" {
		query = append(query, bson.M{"email": email})
	}
	if userName != "" {
		query = append(query, bson.M{"userName": userName})
	}

	return findOne(
		bson.M{
			"$or": query,
		},
		bson.M{})
}

func UpdateTokens(_id primitive.ObjectID, token string, refreshToken string) error {
	// return errors.New("test")

	update := bson.M{
		"$set": bson.M{
			"token":         token,
			"refresh_token": refreshToken,
			"updated_at":    time.Now(),
		},
	}
	_, err := database.UpdateOne(userModel, bson.M{"_id": _id}, update, nil)

	return err

}

func FindByUserId(userId string) (models.User, error) {
	return findOne(
		bson.M{
			"user_id": userId,
		},
		bson.M{})
}

func GetAllUsersForAdmin() ([]models.User, error) {
	groupStage := bson.D{{
		Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
			{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
			{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
		},
	}}

	pipeline := mongo.Pipeline{
		groupStage,
	}

	cursor, context, err := database.Aggregate(userModel, pipeline)

	if err != nil {
		return nil, err
	}

	var allUsers []bson.M
	if err = cursor.All(context, &allUsers); err != nil {
		return nil, err
	}

	var userSlice []models.User

	for _, elem := range allUsers[0]["data"].(primitive.A) {
		if doc, ok := elem.(bson.M); ok { // Check if it's a bson.M document
			var user models.User
			bsonBytes, _ := bson.Marshal(doc)
			if err := bson.Unmarshal(bsonBytes, &user); err == nil {
				userSlice = append(userSlice, user)
			}
		}
	}

	return userSlice, nil
}
