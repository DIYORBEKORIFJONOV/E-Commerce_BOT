package accountmodel

type Account struct {
	AccountID string `json:"account_id" bson:"account_id"`
	Name      string `json:"name" bson:"name"`
	Phone     string `json:"phone" bson:"phone"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Role      string `json:"role" bson:"role"`
}

type CreateAccountRequest struct {
	Name     string `json:"name" bson:"name"`
	Phone    string `json:"phone" bson:"phone"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type CreateAccountResponse struct {
	Account Account `json:"account" bson:"account"`
}

type GetAccountRequest struct {
	AccountID string `json:"account_id" bson:"account_id"`
}

type GetAccountResponse struct {
	Account Account `json:"account" bson:"account"`
}

type ChangePasswordRequest struct {
	AccountID   string `json:"account_id" bson:"account_id"`
	NewPassword string `json:"new_password" bson:"new_password"`
}

type ChangePasswordResponse struct {
	Changed bool `json:"changed" bson:"changed"`
}

type DeleteUserRequest struct {
	AccountID string `json:"account_id" bson:"account_id"`
}

type DeleteUserResponse struct {
	Deleted bool `json:"deleted" bson:"deleted"`
}

type UserExistsRequest struct {
	Phone string `json:"phone" bson:"phone"`
}

type UserExistsResponse struct {
	Exists bool `json:"exists" bson:"exists"`
}

type LoginRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
