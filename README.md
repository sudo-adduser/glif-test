# GLIF Test

## Running The Project

The config values are environmental variables so I installed the `.env` package to make development/local setup easier. If you copy the `example.env` to a local `.env` file in the root directory of the project, fill in the correct values for your system, then the commands `go install` then `go run main.go` everything should be up and running.

## Endpoints

### GET Balance by address

**Route**: `GET /v1/wallet/{address}`

**Example Request:** `http://localhost:8080/v1/wallet/0xb83fEbc22F3C48C99290fCb4bb59F31803283E3b`

**Example Response (Converted to ETH for readability):** `"0.9999996692"`

### POST a new transaction

**Route**: `POST /v1/transaction`

**Example Request:** `http://localhost:8080/v1/transaction`

**Body:** 

```json
{
    "sender": "{sender address}",
    "receiver": "{reciever address}",
    "amount": "{value in wei}"
}
```



**Example Response:**

```json
{
    "status": "submitted",
    "txHash": "0x7d1fc965ea2bcf171df76a209ac2080346d63616a6381f1bbd0237d4e0f1a563"
}
```

### GET all transactions by address (that have been done through this backend project)

**Route:** `/v1/transactions/{address}`

**Example Request:** `http://localhost:8080/v1/transactions/0xb83fEbc22F3C48C99290fCb4bb59F31803283E3b`

**Example Response:**

```json
[
  {
    "id": "b8da02d2-26c9-4beb-a72d-3c05e20b71d8",
    "hash": "0x0ee49734d1f8f3246ac37292448937fb9c81d581fb497444cb8f02eadc07ad5c",
    "sender": "0xb83fEbc22F3C48C99290fCb4bb59F31803283E3b",
    "receiver": "0x0fB94c4e9153B44CCDF83b8dfE2fd02B4Cb36FF2",
    "amount": "1",
    "status": "pending"
  },
  {
    "id": "36a550ef-5d7e-425b-91a5-af8b2ddb2374",
    "hash": "0x3d98bc1420caee12be6e4f9cc9f008d1811eb5e578151590a362eba4d31fd671",
    "sender": "0xb83fEbc22F3C48C99290fCb4bb59F31803283E3b",
    "receiver": "0x0fB94c4e9153B44CCDF83b8dfE2fd02B4Cb36FF2",
    "amount": "1",
    "status": "pending"
  },
  {
    "id": "f03dc07d-57fe-4db0-8155-b4dd3162ea55",
    "hash": "0x3643015f2158cdd00327f06a51f110467453b7180e6de88a9fb2dbc3e0373140",
    "sender": "0xb83fEbc22F3C48C99290fCb4bb59F31803283E3b",
    "receiver": "0x0fB94c4e9153B44CCDF83b8dfE2fd02B4Cb36FF2",
    "amount": "1",
    "status": "pending"
  },
  {
    "id": "5250c597-0a39-4bd1-9d63-b55b3f25b6af",
    "hash": "0xb97e69ab32f05b875c7d3d5974a62933c1a3048bac645ce3f62feac9d9d5db36",
    "sender": "0xb83fEbc22F3C48C99290fCb4bb59F31803283E3b",
    "receiver": "0x0fB94c4e9153B44CCDF83b8dfE2fd02B4Cb36FF2",
    "amount": "1",
    "status": "pending"
  },
  {
    "id": "9a058d42-4b73-4a5e-8c9e-a3154094197a",
    "hash": "0x7d1fc965ea2bcf171df76a209ac2080346d63616a6381f1bbd0237d4e0f1a563",
    "sender": "0xb83fEbc22F3C48C99290fCb4bb59F31803283E3b",
    "receiver": "0x0fB94c4e9153B44CCDF83b8dfE2fd02B4Cb36FF2",
    "amount": "1",
    "status": "confirmed"
  }
]
```

### Notes:

For tracking transaction status I have a goroutine with exponential backoff checking the transaction status. The current initial interval is set to a minute so I didn't eat up my free API credits. I didn't make the backoff configurable but you can change the value in the `trackTransactionStatus` function in the `glif-test/blockchain/wallet.go` file if you're impatient.

