# Ethereum API Service

This project is a REST API built on 
* Golang
* MySQL

For development, the project uses both Docker and docker-compose.

The assets/postman contains a Postman collection with the requests to the API.

Bonus:
* HTMX

### Third-party API

As a third-party API for the data, there has been used [Etherscan](https://docs.etherscan.io/api-endpoints/geth-parity-proxy) to fetch Ethereum blocks and transactions. The project requires an API key from Etherscan to work. It is supposed to be stored in the .env file. By default, Etherescan API allows 5 requests per second.


## Project structure

```bash

eth-api
└───src                          # Source folder containing the application code
    │
    ├───app                      # Application configuration and initialization
    │       app.go
    │
    ├───config                   # Configuration files and variables
    │       config.go
    │
    ├───controllers              # Controller files handling the requests
    │       eth_block_controller.go
    │       eth_transaction_controller.go
    │       eth_block_controller_test.go
    │       eth_transaction_controller_test.go
    │
    ├───database                 # Database configuration and initialization
    │       database.go
    │
    ├───helpers                  # Helper functions and utilities
    │   ├───logger               # Logger utility
    │   │       logger.go
    │   │
    │       eth_block_helper.go
    │       url_helper.go
    │       validator.go
    │       eth_block_helper_test.go
    │       url_helper_test.go
    │       validator_test.go
    │
    ├───middleware               # Middleware for logging requests
    │       middleware.go
    │
    ├───models                   # Data models and database access objects
    │   ├───dao                  # Data Access Objects for interacting with the database
    │   │       eth_block_dao.go
    │   │       eth_transaction_dao.go
    │   │       eth_block_dao_test.go
    │   │       eth_transaction_dao_test.go
    │   │
    │   ├───dto                  # Data Transfer Objects for transforming data between layers
    │   │       eth_block_dto.go
    │   │       eth_transaction_dto.go
    │   │       eth_block_dto_test.go
    │   │       eth_transaction_dto_test.go
    │   │
    │       eth_block_model.go
    │       eth_transaction_model.go
    │       model.go
    │       timestamp.go
    │
    ├───routes                   # Route definitions and setup
    │       routes.go
    │
    └───services                 # Service layer for handling business logic
            cro_service.go
            eth_block_service.go
            eth_transaction_service.go
            eth_block_service_test.go
            eth_transaction_service_test.go
│
│   .env                         # Environment configuration file
│   docker-compose.yaml          # Docker Compose file to define and run multi-container Docker applications
│   Dockerfile                   # Dockerfile to create Docker image for the application
│   go.mod                       # Go modules file, listing all the project dependencies
│   main.go                      # Main entry point for the application


```

## Database

Database contains six tables:
* EthBlocks
* Uncles
* Withdrawals
* EthTransactions
* AccessLists
* StorageKeys

```bash
+--------------------------+         
|       EthBlocks         |         
+--------------------------+         
| Id                   uint|         
| BaseFeePerGas        string|       
| Difficulty           string|       
| ExtraData            string|       
| GasLimit             string|       
| GasUsed              string|       
| Hash                 string|       
| LogsBloom            string|       
| Miner                string|       
| MixHash              string|       
| Nonce                string|       
| Number               string|       
| ParentHash           string|       
| ReceiptsRoot         string|       
| Sha3Uncles           string|       
| Size                 string|       
| StateRoot            string|       
| Timestamp            string|       
| TotalDifficulty      string|       
| TransactionsRoot     string|       
| WithdrawalsRoot      string|       
| CreatedAt         time.Time|       
| UpdatedAt         time.Time|       
+----------------------------+       
|                            |       
|    2  +--------------------------+
|       |       Uncles             |
|       +--------------------------+
|       | Id                   uint|
|       | EthBlockId           uint|
|       | Uncles             string|
|       | CreatedAt       time.Time|
|       | UpdatedAt       time.Time|
|       +--------------------------+
|       |                          |
|    3  +--------------------------+
|       |     Withdrawals          |
|       +--------------------------+
|       | Id                   uint|
|       | EthBlockId           uint|
|       | Address            string|
|       | Amount             string|
|       | Index              string|
|       | ValidatorIndex     string|
|       | CreatedAt       time.Time|
|       | UpdatedAt       time.Time|
|       +--------------------------+
|       |                          |
|    4  +--------------------------+
|       |     EthTransactions      |
|       +--------------------------+
|       | Id                   uint|
|       | EthBlockId           uint|
|       | BlockHash          string|
|       | BlockNumber        string|
|       | ChainId            string|
|       | From               string|
|       | Gas                string|
|       | GasPrice           string|
|       | Hash               string|
|       | Input              string|
|       | MaxFeePerGas       string|
|       | MaxPriorityFeePerGas string|
|       | Nonce              string|
|       | R                  string|
|       | S                  string|
|       | To                 string|
|       | TransactionIndex   string|
|       | Type               string|
|       | V                  string|
|       | Value             string|
|       | CreatedAt      time.Time|
|       | UpdatedAt      time.Time|
|       +--------------------------+
|       |                          |
|       |    5  +--------------------------+
|       |       |       AccessLists         |
|       |       +--------------------------+
|       |       | Id                   uint|
|       |       | EthTransactionId     uint|
|       |       | Address            string|
|       |       | CreatedAt        time.Time|
|       |       | UpdatedAt        time.Time|
|       |       +--------------------------+
|       |       |                          |
|       |       |    6  +--------------------------+
|       |       |       |       StorageKeys       |
|       |       |       +--------------------------+
|       |       |       | Id                   uint|
|       |       |       | AccessListId         uint|
|       |       |       | StorageKey         string|
|       |       |       | CreatedAt        time.Time|
|       |       |       | UpdatedAt        time.Time|
|       |       |       +--------------------------+

```

## Endpoints

### GET /v1/eth-blocks/latest

This endpoint returns the latest saved Ethereum blocks from the database.

#### Response Example:
```bash
[
    {
        "hash": "0x7a99036fb8ab128c138b31c349d072e86029761e3ee471b206cfdda44b0a9d38",
        "number": "0x118a744",
        "transactions": [
            "0x40eb421da31885c079579dd4ed21fb05ad79fb4a580229bfed014ee2f63c890d",
            "0x02230bed4895b85924b7ef2234b7316f844bba38eec8e7787b40eebaa029d787",
            "0xfa3c140140b043535f4faacb333f3f162cfae39437c49fd381937193b15c19d5"
        ]
    },
    {
        "hash": "0x031147b048492dd7663d98ae4b9da8cf42aac4dcaf5ab3a456d60f150187a146",
        "number": "0x118c226",
        "transactions": [
            "0x0074758329914d962193ea1811472206239bd3e2ec0cd2a36d9b54f01157a035",
            "0x2a37aa6fd40619ca1475e1b218fe4ae939cb443af214bdbedd93d85a52e22653",
            "0x556cb92b29441558875dca5cf5b33c4ae578c4d7562f7b6bb72db9f4c27e49ac"
        ]
    },
    {
        "hash": "0xaba8c3976a2416c395d1ae8baf471abfc4ef63ea0b4d309f3b7e010204fa3fb9",
        "number": "0x118c35a",
        "transactions": [
            "0xa0e2f8846c039e6b6885c541aae66d57bbc17c570d899892174672c011bb0c26",
            "0x0e9fc95e5a6c4d47cd4fb8061588976c4942ff0ba960f22a2b7865a74e82ca76",
            "0xb8d825a1db0e624869fc7c46bf2cfc7acf3cd65d2278af5a2fbb422492e01a47"
        ]
    },
    {
        "hash": "0xc8a6715d587133bfe251e08c3f2ac0fa4f139a5e517ba49629fd4ad069f2e8ee",
        "number": "0x118cba5",
        "transactions": [
            "0x80ac56ca3e460a01a899a5051ecae96cc895d8adc224db634817a49e1975dcbc",
            "0x94061fddac6e979320c4fe4d16a53394c03b06198762da427eee572ead71dca0",
            "0x296d8d8fa857e5c37f610bbec1314e788e3f58e014ae41d2724682f638489e54"
        ]
    }
]
```

### GET /v1/eth-blocks/:number
This endpoint returns the Ethereum block with the given block number.

Initially, presence of the requested block is checked in the database. If it is not found, the block is fetched from the Etherscan API and saved to the database.

#### Response Example:
```bash
{
    "hash": "0x2675cbbca2c51d4006d70baccde4cd3c8f6c34150f03d0a65f84fbd0414ea2b1",
    "number": "0x118d151",
    "transactions": [
        "0x341206625112dcdbde681b0fbe9705b404fe3fa6cc74652fa3569a79a36a6e15",
        "0x3fc38422c12c35822de08399de1e42dc15a18f9b2e66471bfc427f7fd9cdd3f0",
        "0x25e09d3ff6c89024454149134c442e910457089dc0a80bd30bcbc8e64e1c0f42"
    ]
}
```

### GET /v1/eth-transactions/:hash
This endpoint returns the Ethereum transaction with the given hash.

Firstly, the presence of the requested transaction is checked in the database. If it is not found, the transaction is fetched from the Etherscan API and saved to the database.

```bash
{
    "blockHash": "0xc8a6715d587133bfe251e08c3f2ac0fa4f139a5e517ba49629fd4ad069f2e8ee",
    "blockNumber": "0x118cba5",
    "chainId": "0x1",
    "from": "0xae2fc483527b8ef99eb5d9b44875f005ba1fae13",
    "hash": "0x224bfe7ff0f33d8affae809b92f75233c59af27a906f2d498ad6118e637d3c1c",
    "to": "0x6b75d8af000000e20b7a7ddf000ba900b4009a80",
    "value": "0x2c4ade1",
    "accessList": null
}
```

### GET /v1/eth-transactions/address/:address
This endpoint returns the Ethereum transactions with the given address as the sender or receiver.

```bash
[
    {
        "blockHash": "0xc8a6715d587133bfe251e08c3f2ac0fa4f139a5e517ba49629fd4ad069f2e8ee",
        "blockNumber": "0x118cba5",
        "chainId": "0x1",
        "from": "0xae2fc483527b8ef99eb5d9b44875f005ba1fae13",
        "hash": "0x224bfe7ff0f33d8affae809b92f75233c59af27a906f2d498ad6118e637d3c1c",
        "to": "0x6b75d8af000000e20b7a7ddf000ba900b4009a80",
        "value": "0x2c4ade1",
        "accessList": null
    },
    {
        "blockHash": "0xc8a6715d587133bfe251e08c3f2ac0fa4f139a5e517ba49629fd4ad069f2e8ee",
        "blockNumber": "0x118cba5",
        "chainId": "0x1",
        "from": "0xae2fc483527b8ef99eb5d9b44875f005ba1fae13",
        "hash": "0xab231aab949289c5db447b277b176f2648baa99e09121462693fa90380084633",
        "to": "0x6b75d8af000000e20b7a7ddf000ba900b4009a80",
        "value": "0x2f05a94",
        "accessList": null
    }
]
```

## Dependencies
* [Fiber](https://github.com/gofiber/fiber) — framework built on top of [Fasthttp](https://github.com/valyala/fasthttp), the fastest HTTP engine for Go
* [Gorm](https://github.com/go-gorm/gorm)  — ORM library for Golang;
* [Air](https://github.com/cosmtrek/air) — live-reloading command line utility for developing Go applications;
* [Zap](https://github.com/uber-go/zap) — fast, structured, leveled logging in Golang;
* [Godotenv](https://pkg.go.dev/github.com/joho/godotenv) — loads environment variables from .env file;
* [Cron](https://github.com/robfig/cron) — a cron library for Golang;


## Commands

Command to run the application:

```bash
docker-compose up --build
```

Run the tests
```bash
go test ./...
```

#### Example of Result 
```bash
ok  	eth-api/src/controllers	0.374s
ok  	eth-api/src/helpers	0.297s
ok  	eth-api/src/models/dao	0.552s
ok  	eth-api/src/models/dto	0.285s
ok  	eth-api/src/services	0.355s
```


#### Docker, Air and MySQL

Docker creates .dbdata folder in the root directory of the project. This folder contains the database files.

Air is a live-reloading utility for Go applications. It watches for file changes in your project directory and restarts your application when there are changes in the code.
Air creates the tmp folder in the root directory of the project. This folder contains the build files.

## Frontend

Project contains a simple frontend based on [HTMX](https://htmx.org/) and [UnoCSS](https://github.com/unocss/unocss).

[HTMX](https://www.youtube.com/watch?v=r-GSGH2RxJs) is a modern library that allows you to access AJAX, CSS Transitions, WebSockets, and Server Sent Events directly in HTML, without writing any JavaScript. When used in combination with Golang, HTMX can bring numerous advantages to web-development.


The frontend could be available on http://localhost:8080/htmx
