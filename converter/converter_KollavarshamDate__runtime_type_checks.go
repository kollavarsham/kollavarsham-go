//go:build !no_runtime_type_checking

// Convert Gregorian date to Kollavarsham date and vice versa
package converter

import (
	"fmt"
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateKollavarshamDate_GetMasaNameParameters(masaNumber *float64) error {
	if masaNumber == nil {
		return fmt.Errorf("parameter masaNumber is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetAharganaParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetDateParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetGregorianDateParameters(val *time.Time) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetJulianDayParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetMonthParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetNaksatraParameters(val *Naksatra) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetSakaDateParameters(val SakaDate) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetSauraDivasaParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetSauraMasaParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetYearParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

