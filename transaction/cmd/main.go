package main

import "transaction/platform/dynamo"

func main() {
	dynamo.NewClient().Connect()
}
