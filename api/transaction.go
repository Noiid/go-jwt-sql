package api

import (
	"CaseMajoo/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

type TransactionsListSuccessResponse struct {
	Transactions repository.Transactions `json:"transactions"`
}

type TransactionsErrorResponse struct {
	Error string `json:"error"`
}

func (api *API) transactionList(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	response := TransactionsListSuccessResponse{}

	userinfo := r.Context().Value("userInfo").(jwt.MapClaims)["ID_User"].(float64)
	str_id_merchant := r.URL.Query().Get("id")
	int_id_merchant, err := strconv.Atoi(str_id_merchant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(MerchantErrorResponse{Error: err.Error()})
		return
	}
	checkMerchant, err := api.merchantsRepo.SelectByID(int64(int_id_merchant))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(MerchantErrorResponse{Error: "Merchant tidak tersedia"})
		return
	}

	if checkMerchant.ID_User != int64(userinfo) {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(MerchantErrorResponse{Error: "Tidak memiliki hak akses!"})
		return
	}

	int_page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(MerchantErrorResponse{Error: err.Error()})
		return
	}

	if int_page <= 0 || int_page > 30 {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(MerchantErrorResponse{Error: "Hanya halaman 1-30 saja!"})
		return
	}

	transaction, err := api.transactionsRepo.SelectByMerchantID(checkMerchant.ID, int64(int_page))
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(MerchantErrorResponse{Error: err.Error()})
			return
		}
	}()
	if err != nil {
		return
	}

	if transaction == nil {
		response.Transactions = repository.Transactions{
			MerchantID:   checkMerchant.ID,
			MerchantName: checkMerchant.Name,
			Date:         "2021-11-" + strconv.Itoa(int_page),
			Omzet:        0,
		}
	} else {
		response.Transactions = repository.Transactions{
			MerchantID:   transaction.MerchantID,
			MerchantName: transaction.MerchantName,
			Date:         transaction.Date,
			Omzet:        transaction.Omzet,
		}
	}

	// encoder.Encode(response)
	result, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
}
