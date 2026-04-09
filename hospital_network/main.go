package main

import (
	"fmt"
)

func myPrint(raw string, msg EmergencyMsg, num string) {
	fmt.Printf("%s: %s\n", num, raw)
	fmt.Printf("House: %s\n", msg.HouseNumber)
	fmt.Printf("Priority: %s\n", msg.Priority)
	fmt.Printf("Source: %s\n------\n", msg.Source)
}

func main() {
	raw := "Help from house 22"

	msg := ParseMsg(raw)

	myPrint(raw, msg, "1")

	raw2 := "Priority: High, Location: 5th Ave, Injury: Broken Leg"
	
	msg2 := ParseMsg(raw2)

	myPrint(raw2, msg2, "2")
}


