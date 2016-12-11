package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	alice, _ := reader.ReadString('\n')
	bob, _ := reader.ReadString('\n')

	aliceArr := strings.Split(alice, " ")
	bobArr := strings.Split(bob, " ")

	aliceScore := 0
	bobScore := 0
	for index, _ := range aliceArr {

		a, _ := strconv.Atoi(strings.TrimSpace(aliceArr[index]))
		b, _ := strconv.Atoi(strings.TrimSpace(bobArr[index]))

		if a > b {
			aliceScore++
		}

		if a < b {
			bobScore++
		}


	}

	fmt.Printf("%d %d\n", aliceScore, bobScore)
	
}