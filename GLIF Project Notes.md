# GLIF Project Notes

My goal was to meet all the project requirements and create a functional service. I definitely cut corners in places but I feel like the bones are there. I treated it like building an MVP for an early stage startup. There is definitely tech debt but it's functional.

I wish I had noticed the issue with Infura sooner. I wasted a bit of time debugging my code since I was getting 200s with incosistant balances. I should have moved over to Alchemy as soon as I noticed the inconsistencies.

## GORM

used gorm for quickly setting up schema, for the sake of time I didn't include migrations and instead just initialize the schema on server start

## ENV Package

I used environmental variables for configuration for quick development (I didn't want to deal with a JSON loader given the scope of this project for example) and added the ENV package for ease of development and to make it easier to get up and running locally 

## Alchemy

I swapped to Alchemy since the free tier of Infura was giving me inconsistent results with no HTTP status codes when querying the wallet balance endpoint.

## Backoff/Transaction Status Tracking

I used the backoff package so I could quickly implement configurable exponential backoff for my transaction confirmation goroutine. In a real world application I would lean towards making that it's own service/going more event driven (for example something like NATS Jetstream) so we could horizontally scale better, establish a queue system, and not lose track of transaction status if the server went down or we had a memory issue.

I would also have liked to make the backoff inital time configurable. The goroutine also makes a call immedietly which is an unessesary call prior to the configured exponential backoff start time.

## Nice To Haves

### Endpoint Validation/HTTP Status Codes

I would like to have endpoint validation/better HTTP error response handlers.

### Log/Error Captures

I don't a standard across the board for logging/error handling in my project. Some of my logs are a bit confusing compared to the error output. For example in my SubmitTransaction function I'm logging that there was an error converting a string to bigint which is accurate but not very informative and then I'm returning a 500 error when I should definitely be more granular. If it's a malformed request it should be a 400 instead. 

### DB Queries/Endpoints

Transaction Retrieval:

* Transactions by date range
* Pagination/Query retrieval limiting
* Query by sender (currently can only do combined)
* Query by reciever
* Transaction by status?

Transaction Submission:

* Add AES encryption for handling private keys sent via HTTPS
* User configurable gas cost
* Endpoint for getting an estimated gas cost would help support a frontend for a wallet

## AES Encryption for handling private keys programatically

I currently have the private key set as an environmental variable that's being read in the `SubmitTransaction` function. I would have liked to have set it up so one could use the API to submit transactions from any wallet rather than have it set.

## Testing

I would have liked to have added some unit tests for the wallet and database methods if I had more time. Full integration tests were definitely out of the question given the scope of the assignment