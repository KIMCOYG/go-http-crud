package main

import "http-collector/pkg/crud"

func main() {
	crud.CreateData()
	crud.HandleRequests()
}
