package greetings

import (
	"errors"
	"fmt"

	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("nombre vacio")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for i, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, fmt.Errorf("%v en la posición %d", err, i+1)
		}

		messages[name] = message
	}

	return messages, nil

}

func randomFormat() string {
	formats := []string{
		"Hola, %v. Bienvenido!",
		"¡Qué bueno verte, %v!",
		"¡Saludo, %v! !Encantado!",
	}
	return formats[rand.Intn(len(formats))]
}
