package main

import "fmt"


type bucket struct {
	name string
	region string
}

type service struct {
	name string
	port int
}
func main() {
	// name := "Krishna"
	bct := bucket{
		name: "Data",
		region: "India",
	}

	serv := service{
		name: "LocalAWS",
		port: 9000,
	}
	project := "LocalAWS"
	port := 9000
    
	fmt.Println("Starting ",project,"on",port)
    fmt.Println(bct.region)
	fmt.Println(bct.name)
	fmt.Println(serv.name,serv.port)
}