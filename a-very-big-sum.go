package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	arrSizeInput, _ := reader.ReadString('\n')
	arrInput, _ := reader.ReadString('\n')


	arrSize, _ := strconv.Atoi(strings.TrimSpace(arrSizeInput))

	var bigSum uint64

	for index, value := range (strings.Split(arrInput, " ")) {

		if index > arrSize {
			break
		}

		value = strings.TrimSpace(value)
		if current, err := strconv.ParseUint(value, 10, 32); err == nil {
			bigSum = bigSum + current	
		}
	}

	fmt.Printf("%d\n", bigSum)

}
