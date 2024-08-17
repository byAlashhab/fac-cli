package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
  arg :=	os.Args[1]
  
  num,err := strconv.Atoi(arg)

  if err!= nil{
	panic(err)
  }
	
  fmt.Println(factorial(num))
}


func factorial(n int) int  {
	if n == 0 || n == 1 {
		return 1
	}

	return n*factorial(n-1)
}