//fazer build para app
go build -o app

//correr a app
./app

//parar a app
ctrl + c

//estrutura da app
main -> cria web server
routes -> contém a gestão das rotas e respectivos handlers
handlers -> contém os handlers chamados pelas rotas

//ver app
localhost:8080