package error

import (
	errField "github.com/thomzes/field-service-booking-app/constants/error/field"
	errFieldSchedule "github.com/thomzes/field-service-booking-app/constants/error/fieldschedule"
)

func ErrorMapping(err error) bool {
	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralErrors...)                        // Fixed: added allErrors as first argument
	allErrors = append(allErrors, errField.FieldErrors...)                 // Fixed: added allErrors as first argument
	allErrors = append(allErrors, errFieldSchedule.FieldScheduleErrors...) // Fixed: added allErrors as first argument

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
