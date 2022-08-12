package api

import (
	"CaseMajoo/repository"
	"fmt"
	"net/http"
)

type API struct {
	usersRepo        repository.UserRepository
	merchantsRepo    repository.MerchantRepository
	transactionsRepo repository.TransactionsRepository
	mux              *http.ServeMux
}

func NewApi(usersRepo repository.UserRepository, merchantsRepo repository.MerchantRepository, transactionsRepo repository.TransactionsRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo,
		merchantsRepo,
		transactionsRepo,
		mux,
	}

	mux.Handle("/api/user/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/user/logout", api.POST(http.HandlerFunc(api.logout)))

	// mux.Handle("/api/user/all", api.GET(http.HandlerFunc(api.userList)))

	mux.Handle("/api/merchant/all", api.GET(api.MiddlewareJWTAuthorization(http.HandlerFunc(api.merchantList))))
	mux.Handle("/api/transaction/list", api.GET(api.MiddlewareJWTAuthorization(http.HandlerFunc(api.transactionList))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
