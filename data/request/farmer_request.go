package request

type CreateFarmerRequest struct {
	UserId             int    `validate:"required,max=200,min=1" json:"userId"`
	CompanyName        string `validate:"required,max=200,min=1" json:"companyName"`
	CompanyTradingName string `validate:"required,max=200,min=1" json:"companyTradingName"`
	CompanyEmail       string `validate:"required,max=200,min=1" json:"companyEmail"`
}

type UpdateFarmerRequest struct {
	ID                 int    `validate:"required"`
	UserId             int    `validate:"required,max=200,min=1" json:"userId"`
	CompanyName        string `validate:"required,max=200,min=1" json:"companyName"`
	CompanyTradingName string `validate:"required,max=200,min=1" json:"companyTradingName"`
	CompanyEmail       string `validate:"required,max=200,min=1" json:"companyEmail"`
}
