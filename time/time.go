package main

import (
	"fmt"
	"time"
)

func main() {
	//now := time.Now().UTC()
	//fmt.Println("Current time is ", now, "\n")
timeUtc()
spaceLine()
spaceLine()
dateTime()
spaceLine()
spaceLine()
dateFunc()
}

func timeUtc(){
	now := time.Now().UTC()
	fmt.Print("Current time is: ", now)
	


}

func dateTime(){
now :=time.Now().UTC()

fmt.Print("Year: ", now.Year())
fmt.Print(" Day ", now.Day())
fmt.Print(" Hour: ", now.Hour())
fmt.Print(" Minutes: ", now.Minute())
fmt.Print(" Seconds: ", now.Second())
fmt.Print(" Nanoseconds: ", now.Nanosecond())
fmt.Print(" Location: ", now.Location())

}

// this function will center on date function in golang
 func dateFunc(){
t :=time.Date(2022, time.December, 10, 23, 30, 60, 0, time.UTC)
//here year is 2022, month is December, day is 10, hour is 23, minute is 0, seconds is 0
fmt.Print(t)


 }
 
func spaceLine(){
	fmt.Println("_________________________________________")
}