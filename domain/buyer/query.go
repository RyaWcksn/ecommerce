package buyer

// Query for buyer
const (
	GetPasswordByEmailQuery = "SELECT id,email,password from ecommerce.buyer WHERE email = ?"
	GetBuyerDataById        = "SELECT name, email, alamat_pengiriman from ecommerce.buyer WHERE id = ?"
)
