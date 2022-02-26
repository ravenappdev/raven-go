# raven-go
Go SDK for Raven
## Installation
```
go get github.com/ravenappdev/raven-go
```
## Documentation

### 1. Introduction
The Raven API is organised in the REST format, our API accepts and return responses in the properly ordered JSON format. Through Raven API you can easily send the notifications and call our various services.<br/>
APIs supported:
-	Send Event
-	Send Event Bulk 
-	
### 2. Authentication
The Raven API uses api key for the authentication of API request. There are many privileges in the secret key (API Key), do not share it anywhere on public platforms. <br/>
To authenticate the request, create a NewAPIClient using your apiKey as shown below:

``` golang
client := raven.NewAPIClient(apiKey)
```

After this a client is created with this apiKey all subsequent requests made further with this client will be using this same key for authentication. 
The requests made without using HTTPs or without the API Key will fail.

### 3. Usage
Here is an example to show how to call for sending an event using Raven API.

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

    appId := "your raven app id"

    payload := data.SendEvent{}
    payload.Event = "event_101"

    payload.User = &data.User{}
    (*payload.User).Email = "user@mail.com"
    (*payload.User).Mobile = "+911234567890"
    
    resp, err := EventApis.SendEvent(context.Background(), appId, payload)

    if err != nil {
        fmt.Print("error: ", resp.Error)
    } else {
        fmt.Print("success: ", resp.Id)
    }
```
