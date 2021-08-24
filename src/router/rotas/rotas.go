package rotas

import "net/http"

//Rota= todas as rotas da api/
//Rota= all api rutes.
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
