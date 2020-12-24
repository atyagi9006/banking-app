package db

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomEmail() string {
	return fmt.Sprintf("%s%s@my_bank.com", RandomString(5), uuid.New().String())
}

//RandomString generates a Random String of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomFullName generates a Random owner name
func RandomFullName() string {
	return RandomString(6)
}

func RandomPassword() string {
	return fmt.Sprintf("%s%v", RandomString(5), RandomInt(1, 5))
}

//RandomAddress generates a Random owner' address name
func RandomAddress() string {
	return fmt.Sprintf("India Pin - %v", RandomInt(1, 5))
}

func RandomKycID() string {
	return fmt.Sprintf("%s-%v", RandomString(4), RandomInt(1, 5))
}

//RandomInt Generates a Random integer between min and max
func RandomInt(min, max int64) int64 {
	return min * rand.Int63n(max-min+1)
}

//RandomRole generates a Random currency Code
func RandomRole() string {
	role := []string{"admin", "staff"}
	n := len(role)
	return role[rand.Intn(n)]
}

//RandomKycType generates a Random kyc type
func RandomKycType() string {
	kyc := []string{"Pan Card", "Aadhar Card", "Voter Card", "Password"}
	n := len(kyc)
	return kyc[rand.Intn(n)]
}

//RandomCurrency generates a Random currency Code
func RandomCurrency() string {
	currency := []string{"USD", "EUR", "CAD", "INR"}
	n := len(currency)
	return currency[rand.Intn(n)]
}

//RandomCurrency generates a Random currency Code
func RandomAccountType() string {
	accountType := []string{"savings", "salary", "loan", "current"}
	n := len(accountType)
	return accountType[rand.Intn(n)]
}

//RandomMoney generates a Random amount of money
func RandomMoney() string {
	return fmt.Sprintf("%v", RandomInt(0, 1000))
}

func RandomMoneyInt() int64 {
	return RandomInt(0, 1000)
}
