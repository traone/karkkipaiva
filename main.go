package main

import "fmt"
import "time"
import "os/exec"

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
	message := string("Today is not karkkipäivä.")

	dudes := []Dude{
		{2009, 1, 17, "Poe"},
		{2012, 7, 30, "Milo"},
		{1979, 8, 9, "Henryk"},
		{1980, 8, 17, "Hannah"},
	}

	// Check for Friday and Saturday
	for i := 0; i < len(allowedDays); i++ {
		cur := allowedDays[i]
		if wkint == cur {
			canHas = true
			message = "Have a nice karkkipäivä."
		}
	}

	// Bdays
	for i := 0; i < len(dudes); i++ {
		if isBirthday(dudes[i]) {
			canHas = true
			message = "Happy birthday " + dudes[i].name + ". Have a nice karkkipäivä!"
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

func say(str string)  {
	cmd := exec.Command("say", str)
	cmd.Start()
	msg := "done did: say \"" + str + "\""
	fmt.Println(msg)
}

func main() {
	_, message := isKarkkipaiva()

	fmt.Println(message)
	say(message)
}
