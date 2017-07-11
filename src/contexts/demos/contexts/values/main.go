package main

import (
	"context"
	"fmt"
)

// START OMIT
type key int

var RequestID key = 0

// END OMIT

func main() {
	ctx := context.Background()
	DoThing(ctx)
	ctx = context.WithValue(ctx, RequestID, "12345")
	DoThing(ctx)
}

func DoThing(ctx context.Context) {
	reqid, ok := ctx.Value(RequestID).(string)
	if !ok {
		fmt.Println("request id not found")
		return
	}
	fmt.Println("RequestID: ", reqid)
}
