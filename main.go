package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	passwordLength := 0

	fmt.Println("Enter the length of the password")
	fmt.Scanf("%d", &passwordLength)

	includeNumber := flag.Bool("num", false, "Include number in password")
	flag.Parse()

	password := ""
	stringDomainLow := "abcdefghijklmnopqrstuvxyz"
	stringDomainHigh := "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	numbers := "0123456789"

	finalPoolOfDomain := stringDomainLow + stringDomainHigh

	switch {
	case *includeNumber:
		finalPoolOfDomain += numbers

	}

	for passwordLength > 0 {
		index := rand.Intn(len(finalPoolOfDomain))
		password += string(finalPoolOfDomain[index])
		passwordLength--

	}

	fmt.Println("Generated password is: ", password)
}
