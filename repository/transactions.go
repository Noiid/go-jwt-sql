package repository

import (
	"database/sql"
	"strconv"
)

type TransactionsRepository struct {
	db *sql.DB
}

func NewTransactionsRepository(db *sql.DB) TransactionsRepository {
	return TransactionsRepository{db}
}

func (t *TransactionsRepository) SelectByMerchantID(merchant_id, page int64) (*Transactions, error) {
	str_page := "2021-11-" + strconv.Itoa(int(page))

	var result = Transactions{}
	err := t.db.QueryRow("SELECT t.merchant_id, m.merchant_name, DATE(t.created_at) as dates, SUM(t.bill_total) as Omzet FROM transactions as t INNER JOIN merchants as m ON t.merchant_id = m.id INNER JOIN outlets as o ON t.outlet_id = o.id GROUP BY DATE(t.created_at), t.merchant_id HAVING dates = ? && t.merchant_id = ?", str_page, merchant_id).Scan(&result.MerchantID, &result.MerchantName, &result.Date, &result.Omzet)
	if err != nil {
		return nil, nil
	}
	return &result, nil
}

func (t *TransactionsRepository) SelectAll() ([]Transactions, error) {
	preparedStatement := `
		SELECT t.merchant_id, m.merchant_name, DATE(t.created_at), SUM(t.bill_total) as Omzet FROM transactions as t
		INNER JOIN merchants as m ON t.merchant_id = m.id
		INNER JOIN outlets as o ON t.outlet_id = o.id
		GROUP BY DATE(t.created_at), t.merchant_id
		ORDER BY t.merchant_id, t.created_at
		`

	rows, err := t.db.Query(preparedStatement)

	if err != nil {
		return nil, err
	}

	var list_trs []Transactions

	for rows.Next() {
		var trs Transactions
		err := rows.Scan(&trs.MerchantID, &trs.MerchantName, &trs.Date, &trs.Omzet)
		if err != nil {
			return nil, err
		}
		list_trs = append(list_trs, trs)
	}
	return list_trs, nil
}
