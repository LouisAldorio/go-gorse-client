package gorseclient

import "encoding/json"


// Create New Item , So the model can recognize new item and consider it as the candidate for recommendation
func (a *API) CreateItem(input InsertItemParam) (*InsertItemResponse, error) {
	jsoned, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	if err := a.newRequest(endpointInsertItem(jsoned)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel InsertItemResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) GetItemByItemID(itemId string) (*GetItemResponse, error) {

	if err := a.newRequest(endpointGetItemByItemID(itemId)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel GetItemResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) DeleteItemByItemID(itemId string) (*DeleteItemResponse, error) {

	if err := a.newRequest(endpointDeleteItemByItemID(itemId)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel DeleteItemResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) UpdateItemByItemID(input UpdateItemParam, itemID string) (*UpdateItemResponse, error) {

	jsoned, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	if err := a.newRequest(endpointUpdateItemByItemID(itemID, jsoned)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel UpdateItemResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) GetItems(n int, cursor string) (*GetItemsResponse, error) {

	if err := a.newRequest(endpointGetItems(n, cursor)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel GetItemsResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) InsertItems(input []*InsertItemParam) (*InsertItemsResponse, error) {

	jsoned, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	if err := a.newRequest(endpointInsertItems(jsoned)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel InsertItemsResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}
