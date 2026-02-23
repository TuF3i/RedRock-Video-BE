package models

type AddVideoReq struct {
	Rvid        int64  `json:"rvid"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
