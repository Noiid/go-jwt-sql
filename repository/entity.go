package repository

type User struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Username string `db:"user_name"`
	Password string `db:"password"`
}

type Merchant struct {
	ID        int64  `db:"id"`
	Name      string `db:"merchant_name"`
	ID_User   int64  `db:"user_id"`
	User_Name string `db:"username"`
}

type Outlets struct {
}

type Transactions struct {
	MerchantID   int64  `db:"merchant_id"`
	MerchantName string `db:"merchant_name"`
	Date         string `db:"date"`
	Omzet        int64  `db:"omzet"`
}
