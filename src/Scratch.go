package main

import (
    "fmt"
    "strings"
)

func main() {
    var str = "{every} {1} {dec 25} {starting dec 25 2016} {do nothing} = $christmas {every} {2} {weeks} {starting jan 1 2016} {send email to list} EXCEPT {$Christmas} = biweeklynewsletter"
    fmt.Println("Example : ")
    fmt.Printf("%v\n\n",str)
    fmt.Println("Initial parsing : ")
    var words = strings.SplitAfterN(str, "}", strings.Count(str,"{"))
    fmt.Printf("%q\n", words)
    
    
    fmt.Println("Correction : ")
    for index,element := range words {
        // index is the index where we are
        // element is the element from someSlice for where we are
        if index > 0 {
            if strings.Index(element, "{") > 1 {
                words[index - 1] = words[index - 1] +  element[0:strings.Index(element, "{") - 1]
                words[index] = element[strings.Index(element, "{"):len(element)]
            }
        }
    }
    fmt.Printf("%q\n", words)
    
}


func FieldsCount(str string) int {
    return 0
}