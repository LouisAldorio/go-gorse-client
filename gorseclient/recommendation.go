package gorseclient

// return Item IDs that is neighboring the current itemID
func (a *API) GetNeighborsOfItems(itemID string, n, offset int) ([]string, error) {

	if err := a.newRequest(endpointGetItemNeighbors(itemID, n, offset)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel []string
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (a *API) GetLatestItems(n, offset int) ([]*GetLatestItemResponse, error) {

	if err := a.newRequest(endpointGetLatestItems(n, offset)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel []*GetLatestItemResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (a *API) GetPopularItems(n, offset int) ([]string, error) {

	if err := a.newRequest(endpointGetPopularItems(n, offset)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel []string
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (a *API) GetRecommendationForUser(userID string, n, offset int) ([]string, error) {

	if err := a.newRequest(endpointGetUserRecommendation(userID, n, offset)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel []string
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (a *API) GetNeighborsOfUser(userID string, n, offset int) ([]string, error) {

	if err := a.newRequest(endpointGetUserNeighbors(userID, n, offset)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel []string
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}

func (a *API) GetCollaborativeFilteringRecommendation(userID string, n, offset int) ([]*GetLatestItemResponse, error) {

	if err := a.newRequest(endpointGetCollaborativeFilteringForUser(userID, n, offset)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel []*GetLatestItemResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return responseModel, nil
}
