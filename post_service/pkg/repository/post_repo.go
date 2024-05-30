package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"post_service/pkg/config/global_vars_config"
	"post_service/pkg/model"
	"post_service/pkg/proto/post_service_proto"
	"post_service/pkg/service/logger"
	"time"
)

func Save(r *post_service_proto.AddPostReq) (*mongo.InsertOneResult, error) {
	result, err := insert(r)
	dir, _ := os.Getwd()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't save user %s post.", r.UserId))
		logger.DebugLog(fmt.Sprintf("%v(Get): %v", dir, err))
		return nil, err
	}
	return result, nil
}

func insert(r *post_service_proto.AddPostReq) (*mongo.InsertOneResult, error) {
	result, err := global_vars_config.DBCollection.InsertOne(
		global_vars_config.DBContext,
		bson.D{
			{Key: "author_id", Value: r.UserId},
			{Key: "content", Value: r.Content},
			{Key: "songs", Value: r.Songs},
			{Key: "playlists", Value: r.Playlists},
			{Key: "albums", Value: r.Albums},
			{Key: "creation_date", Value: time.Now().UTC().String()},
			{Key: "liked", Value: []string{}},
			{Key: "disliked", Value: []string{}},
		},
	)
	return result, err
}

func Get(postId string) (*model.PostModel, error) {
	oid, err := primitive.ObjectIDFromHex(postId)
	dir, _ := os.Getwd()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get object id from id %s.", postId))
		logger.DebugLog(fmt.Sprintf("%v(Get): %v", dir, err))
		return nil, err
	}
	return GetByObjectId(oid)
}

func GetByObjectId(oid primitive.ObjectID) (*model.PostModel, error) {
	dir, _ := os.Getwd()
	filter := bson.D{{Key: "_id", Value: oid}}
	cursor := global_vars_config.DBCollection.FindOne(global_vars_config.DBContext, filter)
	var post model.PostModel
	if err := cursor.Decode(&post); err != nil {
		logger.ErrorLog("Error: couldn't convert record to model.")
		logger.DebugLog(fmt.Sprintf("%v(GetByObjectId): %v", dir, err))
		return nil, err
	}
	return &post, nil
}

func GetUserPosts(userId string) (*model.PostsModel, error) {
	dir, _ := os.Getwd()
	filter := bson.D{{Key: "author_id", Value: userId}}
	cursor, err := global_vars_config.DBCollection.Find(global_vars_config.DBContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't get user %s posts.", userId))
		logger.DebugLog(fmt.Sprintf("%v(GetUserPosts): %v", dir, err))
		return nil, err
	}
	var posts []model.PostModel
	if err := cursor.All(global_vars_config.DBContext, &posts); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode user %s posts.", userId))
		logger.DebugLog(fmt.Sprintf("%v(GetUserPosts): %v", dir, err))
		return nil, err
	}
	return &model.PostsModel{Posts: posts}, nil
}

func Update(r *post_service_proto.UpdatePostReq) error {
	oid, err := primitive.ObjectIDFromHex(r.PostId)
	dir, _ := os.Getwd()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert post id %s to object id.", r.PostId))
		logger.DebugLog(fmt.Sprintf("%v(Update): %v", dir, err))
		return err
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	updateContent := updatePostData(r)
	if updateContent == nil {
		logger.InfoLog(fmt.Sprintf("updated data is nil"))
		return nil
	}
	updateReq := bson.D{{Key: "$set", Value: updateContent}}
	_, err = global_vars_config.DBCollection.UpdateOne(global_vars_config.DBContext, filter, updateReq)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update post %s.", r.PostId))
		logger.DebugLog(fmt.Sprintf("%v(Update): %v", dir, err))
		return err
	}
	return nil
}

func updatePostData(r *post_service_proto.UpdatePostReq) bson.D {
	var updateData bson.D
	if r.Content != "" {
		update(&updateData, "content", r.Content)
	}
	if r.Songs != nil {
		update(&updateData, "songs", r.Songs)
	}
	if r.Albums != nil {
		update(&updateData, "albums", r.Albums)
	}
	if r.Playlists != nil {
		update(&updateData, "playlists", r.PostId)
	}
	return updateData
}

func update(update *bson.D, key string, value interface{}) {
	*update = append(*update, bson.E{Key: key, Value: value})
}

func Delete(postId string) error {
	oid, err := primitive.ObjectIDFromHex(postId)
	dir, _ := os.Getwd()
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't convert post id %s to object id.", postId))
		logger.DebugLog(fmt.Sprintf("%v(Delete): %v", dir, err))
		return err
	}
	return deletePost(oid)
}

func deletePost(postId primitive.ObjectID) error {
	dir, _ := os.Getwd()
	filter := bson.D{{Key: "_id", Value: postId}}
	_, err := global_vars_config.DBCollection.DeleteOne(global_vars_config.DBContext, filter)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't delete post %s.", postId))
		logger.DebugLog(fmt.Sprintf("%v(deletePost): %v", dir, err))
		return err
	}
	return nil
}

func Like(userId string, post *model.PostModel) error {
	post.Liked = append(post.Liked, userId)
	return updatePost(post)
}

func Dislike(userId string, post *model.PostModel) error {
	post.Disliked = append(post.Disliked, userId)
	return updatePost(post)
}

func updatePost(post *model.PostModel) error {
	oid, err := primitive.ObjectIDFromHex(post.Id)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't create object id from post id %s", post.Id))
		logger.DebugLog(fmt.Sprintf("%v", err))
	}
	filter := bson.D{{Key: "_id", Value: oid}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "liked", Value: post.Liked},
			{Key: "disliked", Value: post.Disliked},
		}},
	}
	return updateRecord(filter, update)
}

func updateRecord(filter bson.D, update bson.D) error {
	_, err := global_vars_config.DBCollection.UpdateOne(global_vars_config.DBContext, filter, update)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't update record"))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return err
	}
	return nil
}

func GetUserLiked(userId string) (*model.PostsModel, error) {
	filter := bson.D{
		{Key: "liked", Value: bson.D{
			{Key: "$all", Value: bson.A{userId}},
		}},
	}
	return getPosts(filter)
}

func GetUserDisliked(userId string) (*model.PostsModel, error) {
	filter := bson.D{
		{Key: "disliked", Value: bson.D{
			{Key: "$all", Value: bson.A{userId}},
		}},
	}
	return getPosts(filter)
}

func getPosts(filter bson.D) (*model.PostsModel, error) {
	var posts []model.PostModel
	cursor, err := global_vars_config.DBCollection.Find(global_vars_config.DBContext, filter)
	if err != nil {
		logger.InfoLog(fmt.Sprintf("Error: couldn't get posts by filter %v", filter))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	if err := cursor.All(global_vars_config.DBContext, &posts); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode posts to model"))
		logger.DebugLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return &model.PostsModel{Posts: posts}, nil
}
