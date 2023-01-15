//go:build !no_runtime_type_checking

// Convert Gregorian date to Kollavarsham date and vice versa
package converter

import (
	"fmt"
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateBaseDate_GetMasaNameParameters(masaNumber *float64) error {
	if masaNumber == nil {
		return fmt.Errorf("parameter masaNumber is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseDate) validateSetAharganaParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseDate) validateSetDateParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseDate) validateSetGregorianDateParameters(val *time.Time) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseDate) validateSetJulianDayParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseDate) validateSetMonthParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseDate) validateSetNaksatraParameters(val *Naksatra) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_BaseDate) validateSetSauraDivasaParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseDate) validateSetSauraMasaParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BaseDate) validateSetYearParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

