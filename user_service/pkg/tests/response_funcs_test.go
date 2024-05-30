package tests

import (
	"testing"
	"user_service/pkg/service/response"
	"user_service/pkg/tests/auxilary_funcs"
)

func TestGetResponse(t *testing.T) {
	tests.TestUser.Id = "1"
	r := response.CreateGetResponse(tests.TestUser)
	if !tests.CheckGetResponse(tests.TestUser, r) {
		t.Error("Test failed: response data is invalid")
	}
}

func TestGetPrivateResponse(t *testing.T) {
	tests.TestUser.Id = "1"
	r := response.CreateGetPrivateResponse(tests.TestUser)
	if !tests.CheckGetPrivateResponse(tests.TestUser, r) {
		t.Error("Test failed: response data is invalid")
	}
}

func TestGetErrorResponse(t *testing.T) {
	r := response.CreateErrorGetResponse("some error")
	if r.Error == "" {
		t.Error("Test failed: invalid error response")
	}
}

func TestGetPrivateErrorResponse(t *testing.T) {
	r := response.CreateErrorGetPrivateResponse("some error")
	if r.Error == "" {
		t.Error("Test failed: invalid error response")
	}
}
