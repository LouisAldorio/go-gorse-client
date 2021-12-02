package gorseclient

import (
	"fmt"
	"net/http"
)


// ============================
// ITEM
// ============================
func endpointInsertItem(data []byte) (string, string, []byte) {
	return http.MethodPost, "/api/item", data
}

func endpointGetItemByItemID(itemID string) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/item/%s", itemID), nil
}

func endpointDeleteItemByItemID(itemID string) (string, string, []byte) {
	return http.MethodDelete, fmt.Sprintf("/api/item/%s", itemID), nil
}

func endpointUpdateItemByItemID(itemID string, data []byte) (string, string, []byte) {
	return http.MethodPatch, fmt.Sprintf("/api/item/%s", itemID), data
}

//cursors
func endpointGetItems(n int, cursor string) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/items?n=%d&cursor=%s", n, cursor), nil
}

func endpointInsertItems(data []byte) (string, string, []byte) {
	return http.MethodPost, "/api/items", data
}

// ============================
// USER
// ============================
func endpointInsertUser(data []byte) (string, string, []byte) {
	return http.MethodPost, "/api/user", data
}

func endpointGetUserByItemID(userID string) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/user/%s", userID), nil
}

func endpointDeleteUserByItemID(userID string) (string, string, []byte) {
	return http.MethodDelete, fmt.Sprintf("/api/user/%s", userID), nil
}

func endpointUpdateUserByItemID(userID string, data []byte) (string, string, []byte) {
	return http.MethodPatch, fmt.Sprintf("/api/user/%s", userID), data
}

// cursor
func endpointGetUsers(n int, cursor string) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/users?n=%d&cursor=%s", n, cursor), nil
}

func endpointInsertUsers(data []byte) (string, string, []byte) {
	return http.MethodPost, "/api/users", data
}

// ============================
//FEEDBACK
// ============================
func endpointGetMultipleFeedbacks(n int, cursor string) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/feedback?cursor=%s&n=%d", cursor, n), nil
}

func endpointUpsertMultipleFeedbacks(data []byte) (string, string, []byte) {
	return http.MethodPut, "/api/feedback", data
}

func endpointInsertMultipleFeedbacks(data []byte) (string, string, []byte) {
	return http.MethodPost, "/api/feedback", data
}

func endpointGetMultipleFeedbacksByFeedbackType(feedbackType,cursor string, n int) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/feedback/%s?cursor=%s&n=%d", feedbackType, cursor, n), nil
}

func endpointGetFeedbackByUserItemAndFeedbackType(userID, itemID, feedbackType string) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/feedback/%s/%s/%s", feedbackType, userID, itemID), nil
}

func endpointDeleteFeedbackByUserItemAndFeedbackType(userID, itemID, feedbackType string) (string, string, []byte) {
	return http.MethodDelete, fmt.Sprintf("/api/feedback/%s/%s/%s", feedbackType, userID, itemID), nil
}

func endpointGetFeedbackByUserAndItem(userID, itemID string) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/feedback/%s/%s", userID, itemID), nil
}

func endpointDeleteFeedbackByUserAndItem(userID, itemID string) (string, string, []byte) {
	return http.MethodDelete, fmt.Sprintf("/api/feedback/%s/%s", userID, itemID), nil
}

func endpointGetFeedbackByItem(itemID string) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/item/%s/feedback", itemID), nil
}

func endpointGetFeedbackByItemAndFeedbackType(itemID, feedbackType string) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/item/%s/feedback/%s", itemID, feedbackType), nil
}

func endpointGetFeedbackByUserID(userID string) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/user/%s/feedback", userID), nil
}

func endpointGetFeedbackByUserIDAndFeedbackType(userID, feedbackType string) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/user/%s/feedback/%s", userID, feedbackType), nil
}

// ============================
// RECOMMENDATION
// ============================
func endpointGetItemNeighbors(itemID string, n, offset int) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/item/%s/neighbors?n=%d&offset=%d", itemID, n, offset), nil
}

func endpointGetLatestItems(n, offset int) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/latest?n=%d&offset=%d", n, offset), nil
}

func endpointGetPopularItems(n, offset int) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/latest?n=%d&offset=%d", n, offset), nil
}

func endpointGetUserRecommendation(userID string, n, offset int) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/recommend/%s?n=%d&offset=%d", userID, n, offset), nil
}
	
func endpointGetUserNeighbors(userID string, n, offset int) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/user/%s/neighbors?n=%d&offset=%d", userID, n, offset), nil
}


// ============================
// MEASUREMENT
// ============================
func endpointGetMeasurements(name string, n int) (string, string, []byte) {
	return http.MethodGet, fmt.Sprintf("/api/measurements/%s?n=%d", name, n), nil
}