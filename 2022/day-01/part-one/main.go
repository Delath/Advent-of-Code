package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	max := 0
	sum := 0
	var temp string
	var convertedTemp int

	scanner := bufio.NewScanner(os.Stdin)

	// there is a better way to handle input, i used it in the latest challenges solved with golang
	for scanner.Scan() { // until interrupt key
		temp = scanner.Text()
		
		if strings.Compare(temp,"\n") == -1 { // end of elf detected

			if max < sum {
				max = sum
			}
			sum = 0

		}else{ // still the same elf

			convertedTemp, _ = strconv.Atoi(temp)
			sum += convertedTemp
			temp = scanner.Text()

		}
	}

	fmt.Printf("Elf is carrying %d calories", max)
}