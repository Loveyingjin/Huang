package main

import (
	"Huang/LiJun"
	"os"
	"strings"
)

func main() {
	args := os.Args
	refreshTokens := args[1]
	if refreshTokens != "null" {
		refreshTokenList := strings.Split(refreshTokens, ",")
		HuangLijun := LiJun.HuangLiJun{}
		for _, refreshToken := range refreshTokenList {
			HuangLijun.Run(refreshToken)
		}
	}
}
