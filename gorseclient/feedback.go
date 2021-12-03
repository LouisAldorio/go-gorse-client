package gorseclient

import "encoding/json"

func (a *API) GetFeedbacks(n int, cursor string) (*GetMultipleFeedbacksResponse, error) {

	if err := a.newRequest(endpointGetMultipleFeedbacks(n, cursor)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel GetMultipleFeedbacksResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

// Insert multiple feedback. Existed feedback would be overwritten.
func (a *API) UpsertFeedbacks(input []*FeedbackParam) (*InsertFeedbackResponse, error) {

	jsoned, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	if err := a.newRequest(endpointUpsertMultipleFeedbacks(jsoned)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel InsertFeedbackResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) InsertFeedbacks(input []*FeedbackParam) (*InsertFeedbackResponse, error) {

	jsoned, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	if err := a.newRequest(endpointInsertMultipleFeedbacks(jsoned)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel InsertFeedbackResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) GetFeedbacksByFeedbackType(feedbackType, cursor string, n int) (*GetFeedbacksWithFeedbackTypeResponse, error) {

	if err := a.newRequest(endpointGetMultipleFeedbacksByFeedbackType(feedbackType, cursor, n)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel GetFeedbacksWithFeedbackTypeResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}

	return &responseModel, nil
}

func (a *API) GetFeedbackByFeedbackTypeUserIDAndItemID(feedbackType, itemID, userID string) (*FeedbackResponse, error) {

	if err := a.newRequest(endpointGetFeedbackByUserItemAndFeedbackType(userID, itemID, feedbackType)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel FeedbackResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) DeleteFeedbackByFeedbackTypeUserIDAndItemID(feedbackType, itemID, userID string) (*FeedbackResponse, error) {

	if err := a.newRequest(endpointDeleteFeedbackByUserItemAndFeedbackType(userID, itemID, feedbackType)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel FeedbackResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) GetFeedbacksByUserIDAndItemID(itemID, userID string) ([]*FeedbackResponse, error) {

	if err := a.newRequest(endpointGetFeedbackByUserAndItem(userID, itemID)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel []*FeedbackResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (a *API) DeleteFeedbacksByUserIDAndItemID(itemID, userID string) (*DeleteFeedbackResponse, error) {

	if err := a.newRequest(endpointDeleteFeedbackByUserAndItem(userID, itemID)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel DeleteFeedbackResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (a *API) GetFeedbacksByItemID(itemID string) ([]*FeedbackResponse, error) {

	if err := a.newRequest(endpointGetFeedbackByItem(itemID)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel []*FeedbackResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (a *API) GetFeedbacksByItemIDWithFeedbackType(itemID, feedbackType string) ([]*FeedbackResponse, error) {

	if err := a.newRequest(endpointGetFeedbackByItemAndFeedbackType(itemID, feedbackType)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel []*FeedbackResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (a *API) GetFeedbacksByUserID(userID string) ([]*FeedbackResponse, error) {

	if err := a.newRequest(endpointGetFeedbackByUserID(userID)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel []*FeedbackResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (a *API) GetFeedbacksByUserIDWithFeedbackType(userID, feedbackType string) ([]*FeedbackResponse, error) {

	if err := a.newRequest(endpointGetFeedbackByUserIDAndFeedbackType(userID, feedbackType)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel []*FeedbackResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}
