# plan-generator

This is the solution to the coding task provided by Lendico.

## Install

Extract the content of the zip file to a directory of your choice.
The project uses go mod, so to start please do a:

`go get ./...`

## Running the server

To start the server please do a:

`go run cmd/generator/main.go`

## Project Structure

The project implements the recommended package structure:

- cmd contains runnable code
- internal contains code which is used inside the application but shouldnt be used elsewere

## API

The server runs a single Endpoint which you can reach under

`/rest/api/v1/plan/generate` - Method Post

the Payload of the request should look like this:
    
    {
    "loanAmount": "7500",
    "nominalRate": "6.0",
    "duration": 36,
    "startDate": "2022-01-01T00:00:01Z"
    }

## Response

The Response is a payment plan in json format. Under the key borrowerPayments you find
a list of payments with its corresponding data.

    {
        "borrowerPayments": [
            {
                "date": "2022-01-01T00:00:01Z",
                "borrowerPaymentAmount": "228.17",
                "principal": "190.67",
                "interest": "37.50",
                "initialOutstandingPrincipal": "7500.00",
                "remainingOutstandingPrincipal": "7309.33"
            },
            ...
        ]
    }