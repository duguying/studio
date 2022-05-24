package models

type TrojanUsers struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Quota    int64  `json:"quota"`
	Download int64  `json:"download"`
	Upload   int64  `json:"upload"`
}
