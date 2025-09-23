package error

import (
	errField "github.com/thomzes/field-service-booking-app/constants/error/field"
	errFieldSchedule "github.com/thomzes/field-service-booking-app/constants/error/fieldschedule"
	errTime "github.com/thomzes/field-service-booking-app/constants/error/time"
)

func ErrorMapping(err error) bool {

	var (
		GeneralErrors       = GeneralErrors
		FieldErrors         = errField.FieldErrors
		FieldScheduleErrors = errFieldSchedule.FieldScheduleErrors
		TimeErrors          = errTime.TimeErrors
	)

	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralErrors...)       // Fixed: added allErrors as first argument
	allErrors = append(allErrors, FieldErrors...)         // Fixed: added allErrors as first argument
	allErrors = append(allErrors, FieldScheduleErrors...) // Fixed: added allErrors as first argument
	allErrors = append(allErrors, TimeErrors...)          // Fixed: added allErrors as first argument

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
