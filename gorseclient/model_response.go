package gorseclient

import "time"

// ITEM
// ==============================================================

type InsertItemResponse struct {
	RowAffected int `json:"RowAffected"`
}
type GetItemResponse struct {
	Description string    `json:"Comment"`
	ItemID      string    `json:"ItemId"`
	Labels      []string  `json:"Labels"`
	Timestamp   time.Time `json:"Timestamp"`
}
type DeleteItemResponse struct {
	RowAffected int `json:"RowAffected"`
}
type UpdateItemResponse struct {
	RowAffected int `json:"RowAffected"`
}
type GetItemsResponse struct {
	Cursor string             `json:"Cursor"`
	Items  []*GetItemResponse `json:"Items"`
}
type InsertItemsResponse struct {
	RowAffected int `json:"RowAffected"`
}

// USER
// ==============================================================
type InsertUserResponse struct {
	RowAffected int `json:"RowAffected"`
}

type GetUserResponse struct {
	Description string   `json:"Comment"`
	Labels      []string `json:"Labels"`
	Subscribe   []string `json:"Subscribe"`
	UserID      string   `json:"UserId"`
}

type DeleteUserResponse struct {
	RowAffected int `json:"RowAffected"`
}

type UpdateUserResponse struct {
	RowAffected int `json:"RowAffected"`
}

type GetUsersResponse struct {
	Cursor string             `json:"Cursor"`
	Users  []*GetUserResponse `json:"Users"`
}

type InsertUsersResponse struct {
	RowAffected int `json:"RowAffected"`
}

// FEEDBACK
// ==============================================================
type FeedbackResponse struct {
	Description  string    `json:"Comment"`
	FeedbackType string    `json:"FeedbackType"`
	ItemID       string    `json:"ItemId"`
	UserID       string    `json:"UserId"`
	Timestamp    time.Time `json:"Timestamp"`
}

type GetFeedbacksWithFeedbackTypeResponse struct {
	Cursor    string              `json:"Cursor"`
	Feedbacks []*FeedbackResponse `json:"Feedback"`
}

type GetMultipleFeedbacksResponse struct {
	Cursor    string              `json:"Cursor"`
	Feedbacks []*FeedbackResponse `json:"Feedback"`
}

type InsertFeedbackResponse struct {
	RowAffected int `json:"RowAffected"`
}

type DeleteFeedbackResponse struct {
	RowAffected int `json:"RowAffected"`
}

// RECOMMENDATION
// ==============================================================
type GetLatestItemResponse struct {
	ItemID string  `json:"Id"`
	Score  float64 `json:"Score"`
}

// MEASUREMENT
// ==============================================================
type GetMeasurementResponse struct {
	Description string    `json:"Comment"`
	Name        string    `json:"Name"`
	Timestamp   time.Time `json:"Timestamp"`
	Value       int       `json:"Value"`
}
