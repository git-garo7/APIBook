package router

import "github.com/gorilla/mux"

//gerar retorna um router com rotas configuradas/
//Gerar=generate returns a router with configured routes.
func Gerar() *mux.Router {
	return mux.NewRouter()
}
