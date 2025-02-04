package validator

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)




func ValidateAddress(address string) error {
	basicAddressRegex := regexp.MustCompile(`^[а-яА-Яa-zA-Z0-9\s\.,/-]+$`)
	minLength := 10
	maxLength := 200

	if strings.TrimSpace(address) == "" {
		return fmt.Errorf("address is empty")
	}

	if len(address) < minLength {
		return fmt.Errorf("address is too short (minimum %d characters)", minLength)
	}
	if len(address) > maxLength {
		return fmt.Errorf("address is too long (maximum %d characters)", maxLength)
	}

	if !basicAddressRegex.MatchString(address) {
		return fmt.Errorf("address contains invalid characters")
	}

	if err := validateAddressComponents(address); err != nil {
		return err
	}

	return nil
}

func validateAddressComponents(address string) error {
	components := strings.Split(address, ",")

	if len(components) < 2 {
		return fmt.Errorf("address should contain at least street and house number")
	}

	for i, component := range components {
		component = strings.TrimSpace(component)

		if component == "" {
			return fmt.Errorf("empty address component at position %d", i+1)
		}

		if len(component) < 2 {
			return fmt.Errorf("address component too short at position %d", i+1)
		}

		if err := validateComponent(component); err != nil {
			return fmt.Errorf("invalid component at position %d: %v", i+1, err)
		}
	}

	return nil
}

func validateComponent(component string) error {
	for _, r := range component {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsSpace(r) && r != '/' && r != '-' && r != '.' {
			return fmt.Errorf("invalid character: %c", r)
		}
	}
	return nil
}
