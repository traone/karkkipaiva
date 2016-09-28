package main

import "fmt"
import "time"
//import "github.com/go-redis/redis"

import _ "net/http"

type Dude struct {
	year  int
	month int
	day   int
	name  string
}

func isKarkkipaiva() (bool, string) {

	now := time.Now()
	wkint := uint(now.Weekday())
	canHas := bool(false)
	allowedDays := []uint{5, 6}
	message := string("")

	dudes := []Dude{
		{2009, 1, 17, "Poe"},
		{2012, 7, 30, "Milo"},
		{1979, 8, 9, "Henryk"},
		{1980, 8, 17, "Hannah"},
		{1090, 9, 28, "Mestari"},
	}

	// Check for Friday and Saturday
	for i := 0; i < len(allowedDays); i++ {
		cur := allowedDays[i]
		if wkint == cur {
			canHas = true
		}
	}

	// Bdays
	for i := 0; i < len(dudes); i++ {
		if isBirthday(dudes[i]) {
			canHas = true
			message += "Hyvää synttäriä " + dudes[i].name
		}
	}

	return canHas, message

}

func isBirthday(dude Dude) bool {
	day := int(time.Now().Day())
	month := int(time.Now().Month())

	if dude.day == day && dude.month == month {
		return true
	}
	return false
}

func main() {
	canHasCandy, message := isKarkkipaiva()

	if canHasCandy {
		fmt.Println("HYVÄÄ KARKKIPÄIVÄÄ!!!\n")
		if len(message) > 0 {
			fmt.Println(message)
		}
	} else {
		fmt.Println("Tänään EI ole karkkipäivä.")
	}

}
