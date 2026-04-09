package main

import "strings"


type EmergencyMsg struct {
	RawText     string
	HouseNumber string
	Priority    string // High, Medium, Low
	Source      string
	Injury        string
}


func ParseMsg(raw string) EmergencyMsg {
	msg := EmergencyMsg{
		RawText: raw,
		Injury: "",
	}

	if strings.Contains(raw, "Help from house") {
		msg.Source = "P2P"
		msg.Priority = "HIGH"

		parts := strings.Split(raw, "house") 
		if len(parts) > 1 {
			msg.HouseNumber = strings.TrimSpace(parts[1])
		}
	} else {
		msg.Source = "INTERNET"
		// Priority: High, Location: 5th Ave, Injury: Broken Leg
		parts := strings.Split(strings.TrimSpace(raw), ",")
		for _, part := range parts {
			partArr := strings.Split(strings.TrimSpace(part), ":")

			if len(partArr) != 2 {
				continue
			}

			key := strings.TrimSpace(partArr[0])
			val := strings.TrimSpace(partArr[1])
			
			switch key {
				case "Priority":
					msg.Priority = strings.ToUpper(val)
				case "Location":
					msg.HouseNumber = val
				case "Injury":
					msg.Injury = val
			}
		}
	}

	return msg
}
