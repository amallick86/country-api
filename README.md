# country-api

Country-API Using golang

## Technology Used

1. Golang
2. Postgresql (database)
3. Docker
4. JWT

## Features of the project

1. Create User (username and password)
2. Login User (Generate Token Using JWT)
3. Renew Access Token using Refresh Token
4. Get Countries List from Third Party API and Save the data to the database (https://api.countrystatecity.in/v1/countries)
5. Get Countries List from Database with pagination
6. Get States List from Third Party API and Save the data to the database (https://api.countrystatecity.in/v1/states) [data is so big so it takes few seconds]
7. Get States List from Database with pagination
8. Get Sates List of particular Country

## Package Used

1. gin [ framework ] ( url :- https://github.com/gin-gonic/gin )
2. swagger [ api documentation ] ( url :- https://github.com/swaggo/gin-swagger )
3. ginCors [ cors ] ( url :- https://github.com/gin-contrib/cors )
4. JWT [ token ] ( url :- https://github.com/dgrijalva/jwt-go )
5. viper [ to read config file ] ( url :- https://github.com/spf13/viper )

## Doumentation

http://localhost:8080/swagger/index.html#/

## Steps to run

1. install docker on your machine
2. clone repo " git clone https://github.com/amallick86/country-api.git "
3. OPEN the project in "Goland" or "vs code"
4. For Windows ( OPEN start.sh and SELECT " LF " in end of line sequence )
5. RUN command " make dcup " OR (" docker-compose up ")[windows] || (" sudo docker-compose up ")[linux] on vs code terminal for up your composer and wait for it
6. Hit http://localhost:8080/swagger/index.html#/ url in chrome for the documentation
7. You can use above documentation to test all REST API 
8. In above documentation click on " Authorize " button and paste your token that you get from login api, paste toke as " Bearer your_token "

## Config File
### On project level "app.env"
###### DB_DRIVER=postgres
###### DB_SOURCE=
###### SERVER_ADDRESS=
###### TOKEN_SYMMETRIC_KEY=
###### ACCESS_TOKEN_DURATION=15m
###### REFRESH_TOKEN_DURATION=24h
###### COUNTRY_STATE_API_TOKEN=

## Steps to close the project

1. RUN command " make dcdown " OR " docker-compose down " for down your composer
2. RUN command " make drmi " OR " docker rmi country-api-api-1 " for remove image from your computer
