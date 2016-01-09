package main

import (
    "fmt"
    "strings"
)

func main() {
    var str = "{every} {1} {dec 25} {starting dec 25 2016} {do nothing} = $christmas {every} {2} {weeks} {starting jan 1 2016} {send email to list} EXCEPT {$Christmas} = biweeklynewsletter"
    fmt.Println("Example : ")
    fmt.Printf("%v\n\n",str)
    var words = strings.SplitAfterN(str, "}", strings.Count(str,"{"))
    fmt.Printf("%q\n", words)
}


func FieldsCount(str string) int {
    return 0
}