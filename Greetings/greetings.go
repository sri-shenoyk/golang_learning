package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty Name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomformat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomformat() string {
	format := []string{
		"Hi, %v . Welcome",
		"Nice To meet you %v",
		"Good to see you %v",
	}

	return format[rand.Intn(len(format))]
}
