package authentity



type (
	UserRegisterReq struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
		ConfirimPassword string `json:"confirim_password"`
		PhoneNumber string `json:"phone_number"`
	}
	User struct {
		UserId string
		UserName string `json:"user_name"`
		PasswordHash string `json:"password_hash"`
		PhoneNumber string `json:"phone_number"`
	}
	VerifyUserReq  struct {
		SecretCode int `json:"secret_code"`
	}
	SaveUserRegisterReq struct {
		UserName string `json:"user_name"`
		PasswordHash string `json:"password"`
		PhoneNumber string `json:"phone_number"`
		SecretCode int `json:"secret_code"`
	}


)