package response

type FarmerResponse struct {
	ID                 uint         `json:"id"`
	User               UserResponse `json:"user"`
	CompanyName        string       `json:"companyName"`
	CompanyTradingName string       `json:"companyTradingName"`
	CompanyEmail       string       `json:"companyEmail"`
	CreateAt           string       `json:"createAt"`
	UpdateAt           string       `json:"updateAt"`
}
