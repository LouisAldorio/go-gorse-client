package gorseclient

import "time"

// ITEM
// ==============================================================

type InsertItemParam struct {
	// description about the item
	Description string   `json:"Comment"`
	// ID that identifies your item in your database
	ItemID      string   `json:"ItemId"`
	// Labels helps describe your item better, and make recommendation more accurate
	Labels      []string `json:"Labels"`
	// The current time when this item is created and feed to the model
	Timestamp   time.Time   `json:"Timestamp"`
}
type UpdateItemParam struct {
	Description string   `json:"Comment"`
	Labels      []string `json:"Labels"`
	Timestamp   time.Time   `json:"Timestamp"`
}

// USER
// ==============================================================

type InsertUserParam struct {
	Description string   `json:"Comment"`
	Labels      []string `json:"Labels"`
	Subscribe   []string `json:"Subscribe"`
	UserID      string   `json:"UserId"`
}
type UpdateUserParam struct {
	Description string   `json:"Comment"`
	Labels      []string `json:"Labels"`
	Subscribe   []string `json:"Subscribe"`
}

// FEEDBACK
// ==============================================================

type FeedbackParam struct {
	Description  string `json:"Comment"`
	FeedbackType string `json:"FeedbackType"`
	ItemID       string `json:"ItemId"`
	UserID       string `json:"UserId"`
	Timestamp    time.Time `json:"Timestamp"`
}
