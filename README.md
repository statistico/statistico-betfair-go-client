# Statistico Betfair Go Client

[![Documentation](https://godoc.org/github.com/statistico/statistico-betfair-go-client?status.svg)](http://godoc.org/github.com/statistico/statistico-betfair-go-client)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/statistico/statistico-betfair-go-client/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/statistico/statistico-betfair-go-client/tree/master)

This library is a Golang wrapper around the Betfair Accounts and Betting API. Full documentation and API reference can be found here:

[Documentation](https://docs.developer.betfair.com/)

## Installation
```.env
$ go get -u github.com/statistico/statistico-betfair-go-client
```
## Usage
To instantiate the required Client struct and retrieve a Competition resources:
```go
package main

import (
    "context"
    "fmt"
    "github.com/statistico/statistico-betfair-go-client"
)

func main() {
    creds := betfair.InteractiveCredentials{
        Username:   "user@email.com",
        Password:   "my-secret-password-1",
        Key :       "thUjaEEdBy",
    }
	
	client := *http.Client{}
	
	store := myimplementation.Store{}
    
    bfClient := betfair.NewClient(client, creds, store) 
    
    competitions, err := bfClient.ListCompetitions(context.Background(), ListCompetitionsRequest{}) 

    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }

    // Do something with competitions variable
}
```
## Contributing
You are more than welcome to contribute to this project. Fork and make a Pull Request, or create an Issue if you notice 
any problems or would like to suggest improvements.