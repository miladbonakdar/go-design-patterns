package main

import (
	"fmt"
	"sync"
)

var once = &sync.Once{}
var config *Configuration

type Configuration struct {
	ApplicationPort    int
	DbConnectionString string
	IsCacheAvailable   bool
}

func GetConfiguration() Configuration {
	once.Do(func() {
		//this section will run just once
		//and makes the initialization of the config object lazy.
		//config can be read from env file or db...
		config = &Configuration{
			ApplicationPort:    3333,
			DbConnectionString: "some connection string",
			IsCacheAvailable:   true,
		}
	})
	return *config
}

func main() {
	c := GetConfiguration()
	fmt.Println(c)
	// you can read again but the config has initialized before
	fmt.Println("Connection string : " + GetConfiguration().DbConnectionString)
}
