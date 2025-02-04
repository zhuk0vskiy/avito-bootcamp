package validator

import (
	"backend/internal/model"
)

func IsValidApartmentStatus(status string) bool {
	return status == model.CreatedStatus || status == model.ApprovedStatus ||
		status == model.DeclinedStatus || status == model.ModeratingStatus
}
