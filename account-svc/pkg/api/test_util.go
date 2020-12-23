package api

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

func randomEmail() string {
	return fmt.Sprintf("%s%s@my_bank.com", randomString(5), uuid.New().String())
}

//RandomString generates a random String of length n
func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomFullName generates a random owner name
func randomFullName() string {
	return randomString(6)
}

func randomPassword() string {
	return fmt.Sprintf("%s%v", randomString(5), randomInt(1, 5))
}

//RandomAddress generates a random owner' address name
func randomAddress() string {
	return fmt.Sprintf("India Pin - %v", randomInt(1, 5))
}

func randomKycID() string {
	return fmt.Sprintf("%s-%v", randomString(4), randomInt(1, 5))
}

//RandomInt Generates a random integer between min and max
func randomInt(min, max int64) int64 {
	return min * rand.Int63n(max-min+1)
}

//RandomRole generates a random currency Code
func randomRole() string {
	role := []string{"admin", "staff"}
	n := len(role)
	return role[rand.Intn(n)]
}

//RandomKycType generates a random kyc type
func randomKycType() string {
	kyc := []string{"Pan Card", "Aadhar Card", "Voter Card", "Password"}
	n := len(kyc)
	return kyc[rand.Intn(n)]
}

//RandomCurrency generates a random currency Code
func randomCurrency() string {
	currency := []string{"USD", "EUR", "CAD", "INR"}
	n := len(currency)
	return currency[rand.Intn(n)]
}

//RandomCurrency generates a random currency Code
func randomAccountType() string {
	accountType := []string{"savings", "salary", "loan", "current"}
	n := len(accountType)
	return accountType[rand.Intn(n)]
}

//RandomMoney generates a random amount of money
func randomMoney() string {
	return fmt.Sprintf("%v", randomInt(0, 1000))
}
