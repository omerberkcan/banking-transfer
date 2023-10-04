# Banking Transfer Money API

This is a simple bank transfer system implemented in golang.This project serves as a basic example and should not be used for actual financial transactions.

## Features

- Login with id number and password
- Create bank accounts with id number. 
- There is JWT Authentication
- Transfer money from one account to another.
- Check the balance of your account.

## Requirements/dependencies
- Docker
- Docker-compose

## Env


| Key        | Default          | Description       |
| --------------- | :---------------------: | :-----------------: |
| `BT_PORT` | `9295`                | `Api server port` |
| `BT_AC_TOKEN_TIME` | `5m`                 | `JWT Expire Time m:min h:hour d:day`   |
| `BT_JWT_SECRET`   | `jwt-token-secret-key`                |    `Jwt Secret` |
| `BT_MYSQL_HOST`| `localhost`                | `Mysql Host Address` | 
| `BT_MYSQL_DBNAME`| `bank`                 | `Mysql DB Name`  |
| `BT_MYSQL_USERNAME`| `root`                 | `Mysql User Name`  |
| `BT_MYSQL_PASSWORD`| `1234`                 | `Mysql Password`  |
| `BT_MYSQL_PORT`| `3306`                 | `Mysql Password`  |
| `BT_REDIS_HOST`| `localhost`                 | `Redis Host Address`  |
| `BT_REDIS_DB`| `0`                 | `Redis DB`  |
| `BT_REDIS_PASSWORD`|                 | `Redis Password`  |

## Clone Github Repo

```bash
 git clone https://github.com/omerberkcan/banking-transfer.git
 cd banking-transfer
 ```


## Run banking-tranfer api

```bash
docker compose up -d
 ```


Open your browser and write http://localhost:8080/v1/health. 



## API Request

| Endpoint        | HTTP Method           | Description       |
| --------------- | :---------------------: | :-----------------: |
| `/v1/register` | `POST`                | `Create accounts` |
| `/v1/login` | `POST`                 | `Login User`   |
| `/v1/accounts/profile`   | `GET`                |    `Get Account Info` |
| `/v1/transfers`| `POST`                | `Create transfer` | 
| `/v1/health`| `GET`                 | `Health check`  |

## Test endpoints API using Postman

Download the Postman collection file (`Bank APP.postman_collection.json`) from this repository.


## Test endpoints API using curl

- #### Creating new account

`Request`

```bash
curl --location 'localhost:8080/v1/register' \
--header 'Content-Type: application/json' \
--data '{"id_no":"11111111111","name":"Omer","balance":250,"password":"1234"}'
```

`Response`
```json
{
    "code": 200,
    "status": "Success",
    "data": null,
    "message": "Success"
}
```


`Request`

```bash
curl --location 'localhost:8080/v1/register' \
--header 'Content-Type: application/json' \
--data '{"id_no":"44444444444","name":"Berkcan","balance":150,"password":"1234"}'
```

`Response`
```json
{
    "code": 200,
    "status": "Success",
    "data": null,
    "message": "Success"
}
```
 

- #### Login

```bash
curl --location 'localhost:8080/v1/login' \
--header 'Content-Type: application/json' \
--data '{"id_no":"11111111111","password":"12345"}'
```


`Response`
```json
{
    "code": 200,
    "status": "Success",
    "data": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTYzODc1OTQsInVzZXJfaWQiOjMsInV1aWQiOiIyZWIyODQ5NC1jNDJjLTQ4ZjgtYmI2Yy0zMGI1ZWQ3ZjExNTUifQ.6CeG-6wycbFX9BiWt6VfQhHt71LDax-opXx8PdewTus"
    },
    "message": "Success"
}
```

- #### Fetching account info

`Request`
```bash
curl --location 'localhost:8080/v1/accounts/profile' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTYzODc1OTQsInVzZXJfaWQiOjMsInV1aWQiOiIyZWIyODQ5NC1jNDJjLTQ4ZjgtYmI2Yy0zMGI1ZWQ3ZjExNTUifQ.6CeG-6wycbFX9BiWt6VfQhHt71LDax-opXx8PdewTus'
```

`Response`
```json
{
    "code": 200,
    "status": "Success",
    "data": {
        "id_no": "11111111111",
        "name": "Berkcan",
        "balance": "39.4"
    },
    "message": "success"
}
```

- #### Creating new transfer

`Request`
```bash
curl --location 'localhost:8080/v1/transfers' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTYzODc1OTQsInVzZXJfaWQiOjMsInV1aWQiOiIyZWIyODQ5NC1jNDJjLTQ4ZjgtYmI2Yy0zMGI1ZWQ3ZjExNTUifQ.6CeG-6wycbFX9BiWt6VfQhHt71LDax-opXx8PdewTus' \
--header 'Content-Type: application/json' \
--data '{"id_no":"44444444444",
"amount":60,
"description":"kira"
}'
```

`If the amount is insufficient Response`
```json
{
    "code": 400,
    "status": "Error",
    "data": null,
    "message": "insufficient balance in your wallet for this transfer"
}
```
`Request`
```bash
curl --location 'localhost:8080/v1/transfers' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTYzODc1OTQsInVzZXJfaWQiOjMsInV1aWQiOiIyZWIyODQ5NC1jNDJjLTQ4ZjgtYmI2Yy0zMGI1ZWQ3ZjExNTUifQ.6CeG-6wycbFX9BiWt6VfQhHt71LDax-opXx8PdewTus' \
--header 'Content-Type: application/json' \
--data '{"id_no":"44444444444",
"amount":10,
"description":"kira"
}'
```
`Response`
```json
{
    "code": 200,
    "status": "Success",
    "data": null,
    "message": "Success"
}
```
 
