package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"

   "github.com/atotto/clipboard"
)

const (
	defaultKey string = "you-didnt-specify-a-key"
	defaultSecret string = "to-be-provided"
)

func main() {
	key, secret := getCredentials()
	token := createToken(key, secret)
	fmt.Println(token)
	err := clipboard.WriteAll(token)
	if err != nil {
		fmt.Println("error while copying to clipboard !")
		return
	}
	fmt.Println("token successfully copied to clipboard, you can paste it with <CTRL-v>")
}

func getCredentials() (string, string) {
	key := flag.String("key", defaultKey, "credentials key")
	secret := flag.String("secret", defaultSecret, "shared secret")
	flag.Parse()
	return *key, *secret
}

func encrypt(key string, secret string, ts int64, nonce int) string {
	encryptionKey := fmt.Sprintf("%v%v%v", secret, ts, nonce)
	mac := hmac.New(sha256.New, []byte(encryptionKey))
	mac.Write([]byte(key))
	return hex.EncodeToString(mac.Sum(nil))
}

func createToken(key string, secret string) string {
	ts := time.Now().Unix()
	rand.Seed(time.Now().UnixNano())
	nonce := rand.Intn(int(math.Pow(2, 16)))
	clientToken := encrypt(key, secret, ts, nonce)
	return fmt.Sprintf("%v.%v.%v.%v", key, clientToken, ts, nonce)
}
