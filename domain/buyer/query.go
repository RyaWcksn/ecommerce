package buyer

// Query for buyer
const (
	GetPasswordByEmailQuery = "SELECT email,password from ecommerce.buyer where email = ?"
)
