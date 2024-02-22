package readinter

import (
	"bufio"
	"os"
	"strings"
)


func ReadString() (string,error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}