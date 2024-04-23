package types

type Record struct {
	UUID       string `json:"uuid"`
	MasterKey  string `json:"master_key"`
	UnlockTime int    `json:"unlock_time"`
	DeleteTime int    `json:"delete_time"`
	Message    string `json:"message"`
}
