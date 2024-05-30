package tests

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"user_service/pkg/proto/user_service_proto"
	"user_service/pkg/repository"
	"user_service/pkg/service/db"
	"user_service/pkg/tests/auxilary_funcs"
)

func TestRepoGetByNickname(t *testing.T) {
	db.ConnectBD()
	defer db.DisconnectDB()
	err := tests.CreateTestUser()
	defer tests.DeleteTestUser()
	if err != nil {
		t.Error("Test failed: couldn't create user")
	}
	user, err := repository.GetByNickname(tests.TestUser.Nickname)
	if err != nil {
		t.Error("couldn't get user by nickname")
	}
	tests.TestUser.Id = user.Id
	if !tests.CheckUser(user, tests.TestUser) {
		t.Error("Test failed: user data is invalid")
	}
}

func TestRepoGetById(t *testing.T) {
	db.ConnectBD()
	defer db.DisconnectDB()
	err := tests.CreateTestUser()
	defer tests.DeleteTestUser()
	if err != nil {
		t.Error("Test failed: couldn't create user")
	}
	user, _ := repository.GetByNickname(tests.TestUser.Nickname)
	tests.TestUser.Id = user.Id
	user, err = repository.GetById(user.Id)
	if err != nil {
		t.Error(fmt.Sprintf("Test failed: couldn't get user with id %s", user.Id))
	}
	if !tests.CheckUser(user, tests.TestUser) {
		t.Error("Test failed: invalid user data")
	}
}

func TestRepoUpdating(t *testing.T) {
	db.ConnectBD()
	defer db.DisconnectDB()
	err := tests.CreateTestUser()
	defer tests.DeleteTestUser()
	if err != nil {
		t.Error("Test failed: couldn't create user")
	}
	user, err := repository.GetByNickname(tests.TestUser.Nickname)
	oid, _ := primitive.ObjectIDFromHex(user.Id)
	filter := bson.D{{Key: "_id", Value: oid}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "nickname", Value: "hello"}}}}
	err = repository.Update(filter, update)
	if err != nil {
		t.Error("Test Failed: couldn't update user record")
	}
	tests.TestUser.Nickname = "hello"
	user, err = repository.GetByNickname("hello")
	tests.TestUser.Id = user.Id
	if err != nil {
		t.Error("Test failed: couldn't get user with new name")
	}
	if !tests.CheckUser(user, tests.TestUser) {
		t.Error("Test failed: invalid user data")
	}
}

func TestRepoSaving(t *testing.T) {
	db.ConnectBD()
	defer db.DisconnectDB()
	for _, testData := range tests.GenerateSaveUserData() {
		u := testData.Arg.(user_service_proto.AddRequest)
		resultUser, err := repository.Save(&u)
		if err != nil {
			t.Error("Test failed: couldn't save user")
		}
		user, err := tests.GetUserByObjectId(*resultUser)
		tests.TestUser.Id = user.Id
		if err != nil {
			t.Error("Test failed: couldn't get user with such id")
		}
		if !tests.CheckUser(user, tests.TestUser) {
			t.Error("Test failed: user has invalid data")
		}
		tests.DeleteTestUser()
	}
}

func TestRepoDeleting(t *testing.T) {
	db.ConnectBD()
	defer db.DisconnectDB()
	err := tests.CreateTestUser()
	if err != nil {
		t.Error("Test failed: couldn't create user")
	}
	user, err := repository.GetByNickname(tests.TestUser.Nickname)
	oid, _ := primitive.ObjectIDFromHex(user.Id)
	filter := bson.D{{Key: "_id", Value: oid}}
	err = repository.Delete(filter)
	if err != nil {
		t.Error("Test failed: couldn't delete user with such id")
	}
}
