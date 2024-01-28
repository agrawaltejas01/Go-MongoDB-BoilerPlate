package reviewRepo

import (
	"context"
	"shive-app/database"
	"shive-app/database/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var reviewModel *mongo.Collection

func init() {
	reviewModel, _ = database.GetCollection("review")
}

func buildReviewsSlice(cursor *mongo.Cursor, context context.Context) ([]models.Review, error) {
	var result []models.Review

	for cursor.Next(context) {
		var user models.Review
		cursor.Decode(&user)

		result = append(result, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	defer cursor.Close(context)
	return result, nil
}

func find(query interface{}, projection bson.M) ([]models.Review, error) {
	cursor, context, err := database.Find(reviewModel, query, projection)
	if err != nil {
		return nil, err
	}
	return buildReviewsSlice(cursor, context)
}

func AddReview(review models.Review) (primitive.ObjectID, error) {

	review.Created_at = time.Now()
	review.Updated_at = time.Now()
	return database.InsertOne(reviewModel, review)
}

func DeleteReview(reviewId string) error {
	return database.DeleteOne(reviewModel, bson.M{"review_id": reviewId})
}

func FindByMovieId(movieId int) ([]models.Review, error) {
	return find(
		bson.M{
			"movie_id": movieId,
		},
		bson.M{})
}

func FindByUserId(userId string) ([]models.Review, error) {
	return find(
		bson.M{
			"user_id": userId,
		},
		bson.M{})
}
