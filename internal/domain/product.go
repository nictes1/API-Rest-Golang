package domain

type Product struct {
	ID             int    `json:"id"`
	Description    string `json:"description"`
	ProductCode    string `json:"product_code"`
	Name           string `json:"name"`
	TypeProduct    string `json:"type_product"`
	ExpirationRate string `json:"expiration"`
}
