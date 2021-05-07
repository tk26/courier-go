package examples

import (
	"context"

	"github.com/trycourier/courier-go/v2"
)

func restoreList() {
	client := courier.CreateClient("<YOUR_AUTH_TOKEN>", nil)

	client.RestoreList(context.Background(), "my-list")
}
