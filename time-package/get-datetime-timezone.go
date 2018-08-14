package main

import (
	"fmt"
	"time"
)

// DateTimeTimezone :
// syntax: func LoadLocation(name string) (*Location, error)
// LoadLocation returns the Location with the given name.
func DateTimeTimezone() {

	// Get DateTime based on TimeZone
	InTimeZone()

	// Get DateTime based on TimeStandard
	InTimeStandard()

}

func InTimeZone() {
	t := time.Now()
	fmt.Println("Location : ", t.Location(), " Time : ", t) //Get Local Time

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Location : ", location, " Time : ", t.In(location)) //Get America Time Zone

	loc, _ := time.LoadLocation("Asia/Kuala_Lumpur")
	now := time.Now().In(loc)
	fmt.Println("Location : ", loc, " Time : ", now) //Get Asia Zone
}

func InTimeStandard() {
	t := time.Now()
	z, _ := t.Zone()
	fmt.Println("ZONE : ", z, " Time : ", t) // local time

	location, err := time.LoadLocation("EST")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ZONE : ", location, " Time : ", t.In(location)) // EST (Eastern Standard Time)

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	fmt.Println("ZONE : ", loc, " Time : ", now) // UTC (Coordinated Universal Time)

	loc, _ = time.LoadLocation("MST")
	now = time.Now().In(loc)
	fmt.Println("ZONE : ", loc, " Time : ", now) // MST
}
