package request

type CreateUserRequest struct {
	Username    string `validate:"required,max=200,min=1" json:"username"`
	FirstName   string `validate:"required,max=200,min=1" json:"firstName"`
	LastName    string `validate:"required,max=200,min=1" json:"lastName"`
	Email       string `validate:"required,max=200,min=1" json:"email"`
	Address     string `validate:"required,max=200,min=1" json:"address"`
	PhoneNumber string `validate:"required,max=200,min=1" json:"phoneNumber"`
	DOB         string `validate:"required,max=200,min=1" json:"dob"`
	Role        string `validate:"required,max=200,min=1" json:"role"`
}

type UpdateUserRequest struct {
	Id          int    `validate:"required"`
	Username    string `validate:"required,max=200,min=1" json:"username"`
	FirstName   string `validate:"required,max=200,min=1" json:"firstName"`
	LastName    string `validate:"required,max=200,min=1" json:"lastName"`
	Email       string `validate:"required,max=200,min=1" json:"email"`
	Address     string `validate:"required,max=200,min=1" json:"address"`
	PhoneNumber string `validate:"required,max=200,min=1" json:"phoneNumber"`
	DOB         string `validate:"required,max=200,min=1" json:"dob"`
	Role        string `validate:"required,max=200,min=1" json:"role"`
}
