package validator

import "github.com/google/uuid"

func IsValidUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return err
	}
	return nil
}