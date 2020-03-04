package main

import (
	"bufio"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	input = input[:len(input)-1]
	qrcode.WriteFile(input, qrcode.Medium, 256, "qr.png")

}
