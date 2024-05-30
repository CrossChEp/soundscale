// Package transport Пакет, который содержит хэндлеры, связанные с crud пользователей
package transport

import (
	"context"
	"fmt"
	user_service_proto2 "user_service/pkg/proto/user_service_proto"
	"user_service/pkg/repository"
	"user_service/pkg/service/response"
	"user_service/pkg/service/user"
)

// UserService Структура, содержащая методы для работы с пользователем
type UserService struct {
	user_service_proto2.UserServiceServer
}

// Add Функция добавляет пользователя в базу данных
//
// # Принимает аргумент типа AddRequest:
//
//	type AddRequest struct {
//		Nickname    string
//		Email       string
//		PhoneNumber string
//		Password    string
//	}
//
// # Возвращает объект типа GetPrivateResponse:
//
//	type GetPrivateResponse struct {
//		Id          string
//		Nickname    string
//		Email       string
//		PhoneNumber string
//		Password    string
//		Error       string
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *UserService) Add(_ context.Context, request *user_service_proto2.AddRequest) (*user_service_proto2.GetPrivateResponse, error) {
	resultUser, err := funcs.AddUser(request)
	if err != nil {
		return response.CreateErrorGetPrivateResponse(fmt.Sprintf("Error: couldn't save user. Details: %v", err)), nil
	}
	return response.CreateGetPrivateResponse(resultUser), nil
}

// GetById Функция получает пользователя из БД, используя его ID
//
// # Принимает аргумент типа AddRequest:
//
//	type GetByIdRequest struct {
//		Id string
//	}
//
// # Возвращает объект типа GetResponse:
//
//	type GetResponse struct {
//		Id          string
//		Nickname    string
//		Email       string
//		PhoneNumber string
//		Error       string
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *UserService) GetById(_ context.Context, request *user_service_proto2.GetByIdRequest) (*user_service_proto2.GetResponse, error) {
	resultUser, err := repository.GetById(request.GetId())
	if err != nil {
		return response.CreateErrorGetResponse(fmt.Sprintf("Error: couldn't get user by id. Details: %v", err)), nil
	}
	return response.CreateGetResponse(resultUser), nil
}

// GetByNickname Функция получает пользователя из БД, используя его Никнейм
//
// # Принимает аргумент типа AddRequest:
//
//	type GetByNicknameRequest struct {
//		Nickname string
//	}
//
// # Возвращает объект типа GetResponse:
//
//	type GetResponse struct {
//		Id          string
//		Nickname    string
//		Email       string
//		PhoneNumber string
//		Error       string
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *UserService) GetByNickname(_ context.Context, request *user_service_proto2.GetByNicknameRequest) (*user_service_proto2.GetResponse, error) {
	resultUser, err := repository.GetByNickname(request.Nickname)
	if err != nil {
		return response.CreateErrorGetResponse(fmt.Sprintf("Error: couldn't get user by nickname. Details: %v", err)), nil
	}
	return response.CreateGetResponse(resultUser), nil
}

// GetByNicknamePrivate Функция получает модель пользователя с приватными данными, используя никнейм пользователя
//
// # Принимает аргумент типа AddRequest:
//
//	type GetByNicknameRequest struct {
//		Nickname string
//	}
//
// # Возвращает объект типа GetResponse:
//
//	type GetPrivateResponse struct {
//		Id          string
//		Nickname    string
//		Email       string
//		PhoneNumber string
//		Password    string
//		Error       string
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *UserService) GetByNicknamePrivate(_ context.Context, request *user_service_proto2.GetByNicknameRequest) (*user_service_proto2.GetPrivateResponse, error) {
	resultUser, err := repository.GetByNickname(request.Nickname)
	if err != nil {
		return response.CreateErrorGetPrivateResponse(fmt.Sprintf("Error: couldn't get user by nickname. Details: %v", err)), nil
	}
	return response.CreateGetPrivateResponse(resultUser), nil
}

// Update Функция Update обновляет данные пользователя в БД
//
// # Принимает аргумент типа UpdateRequest:
//
//	type UpdateRequest struct {
//		Token       string
//		Nickname    string
//		Login       string
//		Email       string
//		PhoneNumber string
//		Password    string
//	}
//
// # Возвращает объект типа GetResponse:
//
//	type GetPrivateResponse struct {
//		Id          string
//		Nickname    string
//		Email       string
//		PhoneNumber string
//		Error       string
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *UserService) Update(_ context.Context, request *user_service_proto2.UpdateRequest) (*user_service_proto2.GetResponse, error) {
	resultUser, err := funcs.Update(request)
	if err != nil {
		return response.CreateErrorGetResponse(fmt.Sprintf("Error: couldn't update user. Details: %v", err)), nil
	}
	return response.CreateGetResponse(resultUser), nil
}

// Delete Функция Delete удаляет пользователя из БД
//
// # Принимает аргумент типа DeleteRequest:
//
//	type DeleteRequest struct {
//		Token       string
//	}
//
// # Возвращает объект типа GetResponse:
//
//	type GetPrivateResponse struct {
//		Id          string
//		Nickname    string
//		Email       string
//		PhoneNumber string
//		Error       string
//	}
//
// *Аргумент контекста не используется, поэтому этой функции можно передавать любой контекст
func (s *UserService) Delete(_ context.Context, request *user_service_proto2.DeleteRequest) (*user_service_proto2.GetResponse, error) {
	resultUser, err := funcs.Delete(request)
	if err != nil {
		return response.CreateErrorGetResponse(fmt.Sprintf("Error: couldn't delete user. Details: %v", err)), nil
	}
	return response.CreateGetResponse(resultUser), nil
}
