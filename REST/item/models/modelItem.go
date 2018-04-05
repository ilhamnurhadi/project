package models

type Item struct {
	tblItemID int    `json:"tblItemID"`
	ItemName  string `json:"ItemName"`
	ItemCode  string `json:"ItemCode"`
}
