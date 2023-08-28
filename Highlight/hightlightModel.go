package highlight

import (
	"fmt"
	"unicode"
)

/*
1 letter	1 bold
2 letters	1 bold
3 letters	1 bold
4 letters	2 bold
5 letters	3 bold
6 letters	3 bold
7 letters	4 bold
8 letters	4 bold
9 letters	5 bold
10 letters	5 bold
11 letters	6 bold
12 letters	6 bold
13 letters	7 bold
14 letters	7 bold

len <= 3	1 bold
len == 4	2 bold
len <= 6	3 bold
len <= 8	4 bold
len <= 10	5 bold
len <= 12	6 bold
len <= 14	7 bold
default		8 bold

Attention neurodivergent community - this bionic reading method is absolutely mind blowing.
Your eyes scan the first bold letters and your brain center automatically completes the words.
It lets you read twice as fast, is less overwhelming and helps you to stay focused.
You will feel much more productive and a greater sense of achievement which will boost your confidence
and makes you overall feel more positive. Let me know in the comments if this bionic reading method works for you.
*/

func Light(text []string) string {
	resultText := ""
	for _, word := range text {
		lenWord := countLetters(word)
		switch {
		case lenWord == 0:
			continue
		case lenWord <= 3:
			resultText += fmt.Sprintf("*%s*%s ", word[:1], word[1:])
		case lenWord == 4:
			resultText += fmt.Sprintf("*%s*%s ", word[:2], word[2:])
		case lenWord <= 6:
			resultText += fmt.Sprintf("*%s*%s ", word[:3], word[3:])
		case lenWord <= 8:
			resultText += fmt.Sprintf("*%s*%s ", word[:4], word[4:])
		case lenWord <= 10:
			resultText += fmt.Sprintf("*%s*%s ", word[:5], word[5:])
		case lenWord <= 12:
			resultText += fmt.Sprintf("*%s*%s ", word[:6], word[6:])
		case lenWord <= 3:
			resultText += fmt.Sprintf("*%s*%s ", word[:7], word[7:])
		default:
			resultText += fmt.Sprintf("*%s*%s ", word[:8], word[8:])
		}
	}
	return resultText
}

func countLetters(str string) int {
	count := 0
	for _, letter := range str {
		if unicode.IsLetter(letter) {
			count++
		}
	}
	return count
}
