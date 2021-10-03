package main

import (
	"fmt"
	"time"
)

func main() {
	present_time := time.Now()
	fmt.Println(present_time.Format("01-02-2006  15:04:05 Monday"))

	//Printing the time only
	fmt.Println(present_time.Format("15:04:05"))

	//Printing the Day only
	fmt.Println(present_time.Format("sunday"))

	//Printing the Date only
	fmt.Println(present_time.Format("01-02-2006"))

	present_Date := time.Date(2020, time.August, 24, 12, 30, 21, 0, time.UTC)
	fmt.Println((present_Date.Format("01-02-2006 Monday")))

}


Output:

        10-03-2021  19:06:00 Sunday
        19:06:00
        sunday
        10-03-2021
        08-24-2020 Monday
