package main

import "fmt"

type sqlContainer struct {
	connectionString string
}

func (s *sqlContainer) Get(id string) string {
	// do the query and return the result
	return "data from sql container with id " + id
}

type memoryContainer struct {
	bucket map[string]string
}

func (s *memoryContainer) Get(id string) string {
	// find the data stored in the bucket
	return "data from inMemory container with id " + id
}

type Container interface {
	Get(id string) string
}

// this is the factory interface. it's making the decision for witch container should be created
func NewContainer(connection string) Container {
	if connection == "" {
		return &memoryContainer{bucket: map[string]string{}}
	}
	return &sqlContainer{connection}
}

func main() {
	container1 := NewContainer("")
	container2 := NewContainer("some sql connection string")
	// user doesn't know about the container type and cannot change the properties like
	// bucket or connection string. it's just exposing the methods available in the interface

	fmt.Println(container1.Get("data id"))
	fmt.Println(container2.Get("some other id"))
}
