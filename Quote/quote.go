package quote

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"tgBot/token"
)

type Text struct {
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

var categories = []string{"age", "alone", "amazing", "anger", "architecture", "art", "attitude", "beauty", "best", "birthday", "business", "car", "change", "communications", "computers", "cool", "courage", "dad", "dating", "death", "design", "dreams", "education", "environmental", "equality", "experience", "failure", "faith", "family", "famous", "fear", "fitness", "food", "forgiveness", "freedom", "friendship", "funny", "future", "god", "good", "government", "graduation", "great", "happiness", "health", "history", "home", "hope", "humor", "imagination", "inspirational", "intelligence", "jealousy", "knowledge", "leadership", "learning", "legal", "life", "love", "marriage", "medical", "men", "mom", "money", "morning", "movies", "success"}

func NewQuote() (string, error) {
	randomNum := rand.Intn(67)

	url := fmt.Sprintf("https://api.api-ninjas.com/v1/quotes?category=%s", categories[randomNum])
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", errors.New("request error")
	}
	req.Header.Add("X-Api-Key", token.NINJA_API)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("client error")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("error while reading the body response")
	}
	var newQuote []Text
	err = json.Unmarshal(body, &newQuote)
	if err != nil {
		return "", errors.New("json unmarshal error")
	}
	return fmt.Sprintf("***Author:*** %s\n***Quote:*** %s\n***Category:*** %s", newQuote[0].Author, newQuote[0].Quote, newQuote[0].Category), nil
}

/*
age
alone
amazing
anger
architecture
art
attitude
beauty
best
birthday
business
car
change
communications
computers
cool
courage
dad
dating
death
design
dreams
education
environmental
equality
experience
failure
faith
family
famous
fear
fitness
food
forgiveness
freedom
friendship
funny
future
god
good
government
graduation
great
happiness
health
history
home
hope
humor
imagination
inspirational
intelligence
jealousy
knowledge
leadership
learning
legal
life
love
marriage
medical
men
mom
money
morning
movies
success
*/
