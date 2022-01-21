package main

import (
	"fmt"
	"strings"
	"os"
	
)

func main() {
    str := os.Args[1]
    if str=="" || str=="''"{
    	fmt.Println(0)
    } else{
		masStr := strings.Split(str, " ")
		fmt.Println(len(masStr))
 	}
	
		
}

