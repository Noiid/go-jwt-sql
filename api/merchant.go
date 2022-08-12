package api

import (
	"CaseMajoo/repository"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type MerchantListSuccessResponse struct {
	Merchants []repository.Merchant `json:"merchants"`
}

type MerchantErrorResponse struct {
	Error string `json:"error"`
}

func (api *API) merchantList(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	response := MerchantListSuccessResponse{}
	response.Merchants = make([]repository.Merchant, 0)

	userinfo := r.Context().Value("userInfo").(jwt.MapClaims)["ID_User"]

	merchants, err := api.merchantsRepo.SelectByUserID(int64(userinfo.(float64)))
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

	for _, merchant := range merchants {
		response.Merchants = append(response.Merchants, repository.Merchant{
			ID:        merchant.ID,
			Name:      merchant.Name,
			ID_User:   merchant.ID_User,
			User_Name: merchant.User_Name,
		})
	}

	// encoder.Encode(response)
	result, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
}
