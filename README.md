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
Create a NewAPIClient using your secretKey as shown below:

``` golang
client := raven.NewAPIClient(secretKey)
```

Here is an example for sending a notification event using Raven API.

```golang
package main

import (
	"context"

	raven "github.com/ravenappdev/raven-go"
	"github.com/ravenappdev/raven-go/data"
)

func main() {
	secretKey := "<secret_key>"

	client := raven.NewAPIClient(secretKey)

	appId := "<app_id>"

	payload := data.SendEvent{
		Event:  "event_101",
		Email:  "user@mail.com",
		Mobile: "+911234567890",
	}

	resp, err := client.Events.SendEvent(context.Background(), appId, payload)

	if err != nil {
		log.Print("error: ", resp.Error)
	}
	
```
