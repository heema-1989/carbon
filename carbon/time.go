package main

import (
	"fmt"
	"github.com/golang-module/carbon"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	newTime := time.Date(2021, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	fmt.Printf("Today's date using time package is: %s: Type %T\n", newTime, newTime)
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	c := carbon.Now().SetYear(2021)
	fmt.Printf("Today's date using carbon package is: %s: Type %T\n", c, c)
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	//whether it is current time or not
	fmt.Print("\n", carbon.Now().IsNow())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Print("\n", carbon.Tomorrow().IsWednesday())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	// Whether is April
	carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
	fmt.Print("\n", carbon.Parse("2023-04-05 13:14:15").IsApril())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	// Whether is Monday
	carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
	fmt.Println(carbon.Parse("2023-04-05 13:14:15").IsMonday())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	// just now
	fmt.Println(carbon.Parse("2023-08-05 13:14:15").DiffForHumans())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	// 2 years from now
	fmt.Println(carbon.Parse("2023-05-19 13:14:15").DiffForHumans())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
}
