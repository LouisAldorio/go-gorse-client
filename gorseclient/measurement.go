package gorseclient

// Get the measurements of your recommendation model engine
// name for the measurements name
// n for the number of returned items
func (a *API) GetMeasurementsByName(name string, n int) (*GetMeasurementResponse, error) {
	
	if err := a.newRequest(endpointGetMeasurements(name, n)).Error; err != nil {
		return nil, err
	}
	body, err := a.do()
	if err != nil {
		return nil, err
	}
	var responseModel GetMeasurementResponse
	err = a.jsonUnmarshal(body, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}