package validator

import (
	"fmt"
	"regexp"
	"time"
)

func ValidateDates(firstDate, secondDate string) error {

	layout := "2006-01-02"

	productionDate, err := time.Parse(layout, firstDate)
	if err != nil {
		return fmt.Errorf("error parsing time - %v", err)
	}

	bestBefore, err := time.Parse(layout, secondDate)
	if err != nil {
		return fmt.Errorf("error parsing time - %v", err)
	}

	if productionDate.After(bestBefore) || productionDate.Equal(bestBefore) {
		return fmt.Errorf("production date is greater or equal than best before")
	}

	if productionDate.After(time.Now()) {
		return fmt.Errorf("your product came from the future")
	}

	if bestBefore.Before(time.Now()) {
		return fmt.Errorf("you are forcing an expired product on me")
	}

	return nil

}

func ValidateEmail(email string) error {

	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if !regexp.MustCompile(emailPattern).MatchString(email) {
		return fmt.Errorf("incorrect email format")
	}

	return nil

}
