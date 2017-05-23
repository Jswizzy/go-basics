package main

import (
	"fmt"
)

func main() {
	coursesInProg := []string{"Docker Deep Dive", "Docker Clustering",
		"Docker and Kubernetes"}

	coursesCompleted := []string{"Docker Deep Dive", "Go Fundamentals",
		"Puppet Fundamentals"}

	for _, course := range coursesInProg {
		fmt.Println(course)
		for _, completed := range coursesCompleted {
			if course == completed {
				fmt.Println("\nHey we found a clash.", course, "is in both lists.")
			}
		}
	}
}
