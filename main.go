package main

import (
	"HuangLiJun/LiJun"
	"os"
)

func main() {
	args := os.Args
	refreshToken := args[1]
	pushPlusToken := args[2]
	if refreshToken != "null" {
		HuangLijun := LiJun.HuangLijun{}
		HuangLijun.Run(refreshToken, pushPlusToken)
	}

}
