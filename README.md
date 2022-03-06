# raven-go
Go SDK for Raven
## Installation
```
go get github.com/ravenappdev/raven-go
```
## Documentation

### 1. Introduction
Go SDK helps you send notifications through Raven
APIs supported:
-	Send Event
-	Send Event Bulk 

### 3. Usage
Create a NewAPIClient using your apiKey as shown below:

``` golang
client := raven.NewAPIClient(apiKey)
```

Here is an example for sending a notification event using Raven API.

```golang
package main

import (
    "context"
    "fmt"

    raven "github.com/ravenappdev/raven-go"
    "github.com/ravenappdev/raven-go/data"
)
func main() {
    apiKey := "Your api key"
    
    client := raven.NewAPIClient(apiKey)

    EventApis := client.EventApi

    appId := "<your raven app_id>"

    payload := data.SendEvent{
        Event : "event_101",
        Email : "user@mail.com",
        Mobile : "+911234567890",
    }
    
    resp, err := EventApis.SendEvent(context.Background(), appId, payload)

    if err != nil {
        fmt.Print("error: ", resp.Error)
    } else {
        fmt.Print("success: ", resp.Id)
    }
```
