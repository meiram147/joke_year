package example

var Year int

func YearSet(s int) { // экспортируемая функция
	Year = s
}

func GetSay() string { //
	var say string
	if Year >= 1946 && Year <= 1964 {
		say = "«Привет, бумер!».\n"
	}
	if Year >= 1965 && Year <= 1980 {
		say = "«Привет, представитель X!».\n"
	}
	if Year >= 1981 && Year <= 1996 {
		say = "«Привет, миллениал!».\n"
	}
	if Year >= 1997 && Year <= 2012 {
		say = "«Привет, зумер!».\n"
	}
	if Year == 2013 {
		say = "«Привет, альфа!».\n"
	}
	if Year != 2013 && Year >= 2013 {
		say = "Для твоего года мы еще не придумали название)\n"
	}
	return say
}
