# EA4WebHook
 

## INSTALAR GO


## INSTALAR EN POWERSHELL SWAGGER PARA GO
 - go install github.com/swaggo/swag/cmd/swag@latest
## INICIALIZAR SWAGGER 
 - swag init

## INICIALIZAR EL ARCHIVO DE MOD
 - go mod init EA4WebHook
- go mod tidy

## EJECUTAR

 - go run main.go

## DOCKER

- docker network create webhook_network
- docker build -t meybili/webhook-api .
- docker run -d --name webhook-api-container -p 8080:8080 --network webhook_network meybili/webhook-api
