package main

import (
	"autoSign/platform"
	"os"
	"strings"
)

func main() {
	args := os.Args
	refreshTokens := args[1]
	if refreshTokens != "null" {
		refreshTokenList := strings.Split(refreshTokens, ",")
		HuangLijun := platform.HuangLijun{}
		for _, refreshToken := range refreshTokenList {
			HuangLijun.Run(pushPlusToken, refreshToken)
		}
	}

}
