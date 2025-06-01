package redisCash

type SaveAccount struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Username   string `json:"username"`
	SenderCode string `json:"sender_code"`
}

type Account struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Username   string `json:"username"`
	SenderCode string `json:"sender_code"`
}
