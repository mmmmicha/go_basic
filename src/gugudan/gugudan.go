package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("몇 단을 출력하시겠습니까?")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n, _ := strconv.Atoi(line)

	for i := 1; i < 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}
