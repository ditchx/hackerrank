package main
import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main() {


	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')

	input1 = strings.TrimSpace(input1)

	arrSize, _ := strconv.Atoi(input1)

	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strArr := strings.Split(str, " ")


	if len(strArr) < arrSize {
		fmt.Println("Given array length is smaller than given array size")
		return
	}

	sum := 0
	for ct :=0; ct < arrSize; ct++ {
		cur, _ := strconv.Atoi(strArr[ct])
		sum = sum + cur
	}

	fmt.Println(sum)

}