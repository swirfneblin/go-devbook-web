package rotas

import "webapp/src/controllers"

var rotasLogin = []Rota{
	{
		URI:                "/",
		Metodo:             "GET",
		Funcao:             controllers.CarregarTelaDeLogin,
		RequerAutenticacao: false,
	},
	{
		URI:                "/login",
		Metodo:             "GET",
		Funcao:             controllers.CarregarTelaDeLogin,
		RequerAutenticacao: false,
	},
}
