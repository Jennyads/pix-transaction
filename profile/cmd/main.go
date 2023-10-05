package main

import "profile/platform/dynamo"

func main() {
	dynamo.NewClient().Connect()
}
