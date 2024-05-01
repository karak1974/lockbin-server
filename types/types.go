package types

type Record struct {
	MasterKey  string `json:"master_key"`
	UnlockTime int64  `json:"unlock_time"`
	DeleteTime int64  `json:"delete_time"`
	Message    string `json:"message"`
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
