package repository

import (
	"database/sql"
	"fmt"
)

type MerchantRepository struct {
	db *sql.DB
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return MerchantRepository{db}
}

func (m *MerchantRepository) SelectByUserID(user_id int64) ([]Merchant, error) {
	rows, err := m.SelectAll()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var list_merchant []Merchant
	for _, merchant := range rows {
		if merchant.ID_User == user_id {
			list_merchant = append(list_merchant, merchant)
		}
	}

	return list_merchant, nil
}

func (m *MerchantRepository) SelectAll() ([]Merchant, error) {
	preparedStatement := `
		SELECT m.id, m.merchant_name, m.user_id, u.user_name 
		FROM merchants m
		INNER JOIN users u 
		ON m.user_id = u.id`

	rows, err := m.db.Query(preparedStatement)

	if err != nil {
		return nil, err
	}

	var list_merchant []Merchant

	for rows.Next() {
		var mer Merchant
		err := rows.Scan(&mer.ID, &mer.Name, &mer.ID_User, &mer.User_Name)
		if err != nil {
			return nil, err
		}
		list_merchant = append(list_merchant, mer)
	}
	return list_merchant, nil
}

func (m *MerchantRepository) SelectByID(id int64) (*Merchant, error) {
	rows, err := m.SelectAll()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	for _, merchant := range rows {
		if merchant.ID == id {
			return &merchant, nil
		}
	}

	return nil, fmt.Errorf("merchant tidak tersedia")
}
