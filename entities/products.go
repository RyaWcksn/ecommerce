package entities

type ProductListEntity struct {
	Id          int    `json:"id"`
	ProductName string `json:"productName"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Seller      int    `json:"seller"`
}
