package main

import (
	"fmt"
	"strings"
	"time"
)
import "github.com/golang-module/carbon"

func main() {
	fmt.Printf("Today's date using carbon  is: %s : Type: %T\n", carbon.Now().ToDateTimeString(), carbon.Now().ToDateTimeString())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Printf("Today's date using time package is: %s: Type %T\n", time.Date(2000, time.April, 16, 21, 5, 0, 0, time.UTC).Local(), time.Date(2000, time.April, 16, 21, 5, 0, 0, time.UTC).Local())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Printf("Today's date using time package is: %s: Type %T\n", time.Now(), time.Now())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Printf("Today's date using carbon  is: %s : Type: %T\n", carbon.Now(), carbon.Now())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Printf("Today's date using carbon is: %s : Type: %T\n", carbon.Now().ToString(), carbon.Now().ToString())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Printf("Today's date using carbon is: %s : Type: %T\n", carbon.Now().ToDateString(), carbon.Now().ToDateString())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Printf("Current time using carbon is: %s : Type: %T\n", carbon.Now().ToTimeString(), carbon.Now().ToTimeString())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	//Returns date time of today in given timezone
	fmt.Printf("Today's date using carbon is in america/chicago timezone: %s : Type: %T\n", carbon.Now(carbon.Chicago).ToDateTimeString(), carbon.Now().ToDateTimeString())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Printf("Today's date using carbon in asia/bangkok timezone  is: %s : Type: %T\n", carbon.Now(carbon.Bangkok).ToDateTimeString(), carbon.Now().ToDateTimeString())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Printf("Today's date using carbon is: %s : Type: %T\n", carbon.Now().ToDateTimeString(), carbon.Now().ToDateTimeString())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Printf("Today's timestamp using carbon is: %d : Type: %T\n", carbon.Now().Timestamp(), carbon.Now().Timestamp())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Printf("Yesterday's date time using carbon is: %s : Type: %T\n", carbon.Yesterday(), carbon.Yesterday())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
	fmt.Printf("Tomorrow's date time using carbon is: %s : Type: %T\n", carbon.Tomorrow(), carbon.Yesterday())
	fmt.Printf("%s\n", strings.Repeat("-", 20))
}
