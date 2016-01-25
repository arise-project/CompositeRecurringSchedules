package main

import (
    "fmt"
    "strings"
    "encoding/json"
    "time"
)



type WordConfiguration struct {
		TypeName string
		Values []string
		Script bool
	}



const(
    every = 1
    once = 2
)



const(
    year = 1
    month = 2
    week = 3
    day = 4
)
	
	
type ScheduleElement struct {
		SeriesType int // every|once
		SeriesNumber int // 1, 2...
		SeriesDate time.Time() // dec 25
		SeriesPeriod int // week|day..
		fromDate time.Time() // starting dec 25 2016
		toDate time.Time() // ending dec 25 2017
		ExpressionName string //do nothing
		ExpressionText string //$Christmas
	}



func main() {
    var str = "{every} {1} {dec 25} {starting dec 25 2016} {do nothing} = $christmas {every} {2} {weeks} {starting jan 1 2016} {send email to list} EXCEPT {$Christmas} = biweeklynewsletter"
    fmt.Println("Example : ")
    fmt.Printf("%v\n\n",str)
    fmt.Println("Initial parsing : ")
    var wordsCount = FieldsCount(str)
    var words = strings.SplitAfterN(str, "}", wordsCount)
    fmt.Printf("%q\n", words)
    
    fmt.Println("Correction : ")
    
    var scriptsCount = CorrectScriptPart(words)
    
    fmt.Printf("%q\n Corrected count %v\n", words, scriptsCount)
    
    Configuration()
}



func CorrectScriptPart(words []string) int {
    var scriptsCount = 0;
    for index,element := range words {
        // index is the index where we are
        // element is the element from someSlice for where we are
        if index > 0 {
            if strings.Index(element, "{") > 1 {
                words[index - 1] = words[index - 1] +  element[0:strings.Index(element, "{") - 1]
                words[index] = element[strings.Index(element, "{"):len(element)]
                scriptsCount = scriptsCount + 1
            }
        }
    }
    
    return scriptsCount
}



func Validate(str string) bool {
    return true
}



func FieldsCount(str string) int {
    return  strings.Count(str,"{")
}



func Configuration() {
    var jsonBlob = []byte(`[
		{"TypeName": "Platypus", "Values": ["Crimson","Red","Ruby","Maroon"], "Script": true},
		{"TypeName": "Quoll", "Values": ["Crimson","Red"], "Script": false}
	]`)
	
	var configs []WordConfiguration
	err := json.Unmarshal(jsonBlob, &configs)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("\nConfiguration:")
	fmt.Printf("%+v\n", configs)

}



//begin schedule 
func IsOccuring(event string, aDate time.Time) bool{
	return false
}


		
func Dates(event string, fromDate time.Time, toDate time.Time) [] time.Time {
	return nil
}
		
		
		
func NextOccurance(event string, aDate time.Time) time.Time{
	return time.Now()
}

//end schedule 