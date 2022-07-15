package main

import (
	"encoding/base64"
	"fmt"
	"mysterious-code/lib"
)

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	decoded := string(sd)
	fmt.Printf("decoded: %s\n", decoded)

	whatIsIt = lib.ReverseString(decoded)
	fmt.Printf("answer: %s", whatIsIt)
}
