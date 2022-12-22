package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(

		//returns true/false wheather the substring is present in the string or not : true
		strings.Contains("test", "te"),

		//Count the number of substr in str : 2
		strings.Count("test", "t"),

		//
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),

		//returns the starting index, if the substr occur multiple times then it will return the index where it first occurs
		strings.Index("terter", "er"),

		/*takes a slice of strings and makes a new string by joining them, the sep string is placed between thees strings to separate them*/
		strings.Join([]string{"ax", "bx", "cx"}, "--"),

		// replace the string and n is the number of substring it will replace from the beginning in the main string
		strings.Replace("tester", "te", "pe", 1), // pester
		strings.ReplaceAll("tester", "te", "pe"), // pesper

		// Split the main string and create a slice of strings
		strings.Split("abcdebcsa", "bc"), //[a de sa]
		// Split the main string into n strings with the help of separator sep, split from the begining
		strings.SplitN("abcdebcsa", "b", 2), // [a cdebcsa]

		//strings.SplitAfter()
		//strings.SplitAfterN()

		//strings.ToLower()
		//strings.ToUpper()
		
	)
}
