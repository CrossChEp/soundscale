package repository

import (
	"comment_service/pkg/config/global_vars_config"
	model "comment_service/pkg/model/comment"
	"comment_service/pkg/proto/comment_service_proto"
	"comment_service/pkg/service/logger"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
)

func Save(r *comment_service_proto.AddCommentReq) (*mongo.InsertOneResult, error) {
	result, err := global_vars_config.DbCollection.InsertOne(
		global_vars_config.DbContext,
		createAddBson(r),
	)
	if err != nil {
		dir, _ := os.Getwd()
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't save comment of user %s", r.UserId), err, dir)
		return nil, err
	}
	return result, nil
}

func createAddBson(r *comment_service_proto.AddCommentReq) bson.D {
	releaseDate := time.Now()
	return bson.D{
		{Key: "user_id", Value: r.UserId},
		{Key: "entity_id", Value: r.EntityId},
		{Key: "entity_type", Value: r.EntityType},
		{Key: "content", Value: r.Content},
		{Key: "creation_date", Value: releaseDate.String()},
	}
}

func Update(commentId string, newContent string) error {
	oid, err := convertToOid(commentId)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "content", Value: newContent},
		}},
	}
	return updateRecord(filter, update)
}

func updateRecord(filter bson.D, update bson.D) error {
	_, err := global_vars_config.DbCollection.UpdateOne(global_vars_config.DbContext, filter, update)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update record. Details: %v", err))
		return err
	}
	return nil
}

func GetWithType(commentId string, entityType string) (*model.CommentGetModel, error) {
	dir, _ := os.Getwd()
	filter, err := generateFilter(commentId, entityType)
	if err != nil {
		return nil, err
	}
	cursor := global_vars_config.DbCollection.FindOne(global_vars_config.DbContext, filter)
	comment := model.CommentGetModel{}
	if err := cursor.Decode(&comment); err != nil {
		logger.ErrorWithDebugLog("Couldn't decode comment data", err, dir)
		return nil, err
	}
	return &comment, nil
}

func generateFilter(commentId string, entityType string) (bson.D, error) {
	oid, err := convertToOid(commentId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	if entityType != "" {
		filter = append(filter, bson.E{Key: "entity_type", Value: entityType})
	}
	return filter, nil
}

func Get(commentId string) (*model.CommentGetModel, error) {
	oid, err := convertToOid(commentId)
	if err != nil {
		return nil, err
	}
	return GetByObjectId(oid)
}

func GetByObjectId(oid primitive.ObjectID) (*model.CommentGetModel, error) {
	dir, _ := os.Getwd()
	filter := bson.D{{Key: "_id", Value: oid}}
	cursor := global_vars_config.DbCollection.FindOne(global_vars_config.DbContext, filter)
	comment := model.CommentGetModel{}
	if err := cursor.Decode(&comment); err != nil {
		logger.ErrorWithDebugLog("Couldn't decode comment data", err, dir)
		return nil, err
	}
	return &comment, nil
}

func Delete(commentId string) error {
	dir, _ := os.Getwd()
	oid, err := convertToOid(commentId)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	_, err = global_vars_config.DbCollection.DeleteOne(global_vars_config.DbContext, filter)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't delete comment %s", commentId), err, dir)
		return err
	}
	return nil
}

func GetEntityComments(entityId string, entityType string) ([]model.CommentGetModel, error) {
	dir, _ := os.Getwd()
	filter := bson.D{
		{Key: "entity_id", Value: entityId},
		{Key: "entity_type", Value: entityType},
	}
	cursor, err := global_vars_config.DbCollection.Find(global_vars_config.DbContext, filter)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't get comment of entity %s", entityId), err, dir)
		return nil, err
	}
	return decodeCursorToArr(cursor)
}

func GetUserComments(userId string) ([]model.CommentGetModel, error) {
	dir, _ := os.Getwd()
	filter := bson.D{
		{Key: "user_id", Value: userId},
	}
	cursor, err := global_vars_config.DbCollection.Find(global_vars_config.DbContext, filter)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't get comments of user %s", userId), err, dir)
		return nil, err
	}
	return decodeCursorToArr(cursor)
}

func decodeCursorToArr(cursor *mongo.Cursor) ([]model.CommentGetModel, error) {
	dir, _ := os.Getwd()
	var comments []model.CommentGetModel
	if err := cursor.All(global_vars_config.DbContext, &comments); err != nil {
		logger.ErrorWithDebugLog("Couldn't decode comment data", err, dir)
		return nil, err
	}
	return comments, nil
}

func convertToOid(id string) (primitive.ObjectID, error) {
	dir, _ := os.Getwd()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.ErrorWithDebugLog(fmt.Sprintf("Error: couldn't get object id from id %s", id), err, dir)
		return oid, err
	}
	return oid, nil
}
