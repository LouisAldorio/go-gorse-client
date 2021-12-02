package gorseclient

// ITEM
// ==============================================================
var (
	InsertItemParam struct {
		Description string   `json:"description" gorse:"Comment"`
		ItemID      string   `json:"item_id" gorse:"ItemId"`
		Labels      []string `json:"labels" gorse:"Labels"`
		Timestamp   string   `json:"timestamp" gorse:"Timestamp"`
	}
)