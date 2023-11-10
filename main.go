package main

import (
	"autoSign/platform"
	"os"
	"strings"
)

func main() {
	args := os.Args
	pushPlusToken := args[1]
	refreshTokens := args[2]
	if refreshTokens != "null" {
		refreshTokenList := strings.Split(refreshTokens, ",")
		HuangLijun := platform.HuangLijun{}
		for _, refreshToken := range refreshTokenList {
			HuangLijun.Run(pushPlusToken, refreshToken)
		}
	}

}
