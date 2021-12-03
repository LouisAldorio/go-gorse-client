package gorseclient

import "encoding/json"

func (a *API) CreateUser(input InsertUserParam) (*InsertUserResponse, error) {

	jsoned, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	if err := a.newRequest(endpointInsertUser(jsoned)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel InsertUserResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) GetUserByUserID(userID string) (*GetUserResponse, error) {

	if err := a.newRequest(endpointGetUserByUserID(userID)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel GetUserResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) DeleteUserByUserID(userID string) (*DeleteUserResponse, error) {

	if err := a.newRequest(endpointDeleteUserByUserID(userID)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel DeleteUserResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) UpdateUserByUserID(userID string, input UpdateUserParam) (*UpdateUserResponse, error) {

	jsoned, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	if err := a.newRequest(endpointUpdateUserByUserID(userID, jsoned)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel UpdateUserResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) GetUsers(n int, cursor string)(*GetUsersResponse, error) {

	if err := a.newRequest(endpointGetUsers(n, cursor)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel GetUsersResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) CreateUsers(input []*InsertUserParam) (*InsertUsersResponse, error) {

	jsoned, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	if err := a.newRequest(endpointInsertUsers(jsoned)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel InsertUsersResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}
