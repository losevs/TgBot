package help

func Help() string {
	return "/light + *text*\n" +
		"***Highlight your text to read faster***\n" +
		"/qr + *text*\n" +
		"***Make your own qr-code***\n" +
		"/weather + *city/country*\n" +
		"***Daily forecast for any city***\n" +
		"/dice or /slot\n" +
		"***Throw a dice or spin the slots***\n" +
		"/roll + *number*\n" +
		"***Random number in the selected range***\n" +
		"/curr + *your currency* + *need currency* + *amount*\n" +
		"***Convert to the chosen currency***\n" +
		"/allcur\n" +
		"***See all the available currencies***\n" +
		"/quote\n" +
		"***Random quote***\n" +
		"/love + *name* + *name*\n" +
		"***Check love chances between 2 people***\n"
}
