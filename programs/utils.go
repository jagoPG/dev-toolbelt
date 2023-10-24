package main

import (
	"fmt"
	"jagoPG/utils/colors"
	"jagoPG/utils/encoding"
	"jagoPG/utils/strings"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("error: no operation specified")
		os.Exit(1)
	}
	if !existsOperation(args) {
		fmt.Printf("error: invalid operation")
		os.Exit(1)
	}

	result, err := functions[args[1]](args[2:])

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println(result)
}

func existsOperation(args []string) bool {
	return functions[args[1]] != nil
}

var functions = map[string]func([]string) (string, error){
	"rgb2hex":      colors.ConvertRGBToHex,
	"hex2rgb":      colors.ConvertHexToRGB,
	"count":        strings.Count,
	"base64Encode": encoding.EncodeBase64,
	"base64Decode": encoding.DecodeBase64,
}
