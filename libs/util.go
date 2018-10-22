package libs

import "fmt"
import "bufio"
import "os"
import "strings"

func inputText(mensaje string) string {
	res, reader := "", bufio.NewReader(os.Stdin)

	fmt.Print(mensaje)
	text, err := reader.ReadString('\n')

	if err == nil {
		res = strings.TrimRight(text, "\r\n")
	}

	return res
}

func inputDecimal(message string) float32 {
	var res float32

	fmt.Print(message)
	fmt.Scan(&res)

	return res
}

func inputNumber(message string) uint32 {
	return uint32(inputDecimal(message))
}
