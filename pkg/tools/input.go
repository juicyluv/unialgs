package tools

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func ReadInt() int {
	str := ReadString()

	n, err := strconv.Atoi(str)

	if err != nil {
		panic(str + " is not a number")
	}

	return n
}

func ReadString() string {
	text, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text
}
