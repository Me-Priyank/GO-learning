package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to pizza app")
	fmt.Println("Rate our pizza b/w 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	// fmt.Println("Rating",input)

	numRate, err := strconv.ParseFloat(strings.TrimSpace(input),64)// because it takes "\n" also

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating :",numRate+1)
	}
}
