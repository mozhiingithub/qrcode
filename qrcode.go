package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	qrcode "github.com/skip2/go-qrcode"
)

const imgName = "qr.png"

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Input text:")
	input, _ := inputReader.ReadString('\n')
	input = input[:len(input)-1]
	qrcode.WriteFile(input, qrcode.Medium, 256, imgName)
	exec.Command("xdg-open", imgName).Start()
}
