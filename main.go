// main package
package main

import (
    "fmt"	
    "strings"
)

func main() {
    str := "10:228:31:49"
    fmt.Printf("%v", strings.Split(str, ":"))
}
