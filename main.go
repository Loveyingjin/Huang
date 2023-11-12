package main

import (
	"HuangLiJun/LiJun"
	"os"
)

func main() {
	args := os.Args
	refreshToken := args[1]
	if refreshToken != "null" {
		HuangLijun := LiJun.HuangLijun{}
		HuangLijun.Run(refreshToken)
	}

}
