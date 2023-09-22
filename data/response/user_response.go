package response

type UserResponse struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	DOB         string `json:"dob"`
	Role        string `json:"role"`
	CreateAt    string `json:"createAt"`
	UpdateAt    string `json:"updateAt"`
}
