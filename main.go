package main

import (
	"Huang/Lijun"
	"os"
	"strings"
)

func main() {
	args := os.Args
	refreshTokens := args[1]
	if refreshTokens != "null" {
		refreshTokenList := strings.Split(refreshTokens, ",")
		HuangLijun := HuangLijun{}
		for _, refreshToken := range refreshTokenList {
			HuangLijun.Run(refreshToken)
		}
	}
}
