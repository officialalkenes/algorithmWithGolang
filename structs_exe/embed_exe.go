package main

import "fmt"

type Sender struct {
	rateLimit int
	User
}

type User struct {
	name   string
	number int
}

func test_rate_limit(s Sender) {
	fmt.Println("Sender name: ", s.name)
	fmt.Println("Sender Number: ", s.number)
	fmt.Println("Sender Ratelimit: ", s.rateLimit)
}

func main() {
	test_rate_limit(
		Sender{
			rateLimit: 1000,
			User: User{
				name:   "Olamide",
				number: 9123456789,
			},
		})
}
