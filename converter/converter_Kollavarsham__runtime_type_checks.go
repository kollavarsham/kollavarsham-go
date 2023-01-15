//go:build !no_runtime_type_checking

// Convert Gregorian date to Kollavarsham date and vice versa
package converter

import (
	"fmt"
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (k *jsiiProxy_Kollavarsham) validateFromGregorianDateParameters(date *time.Time) error {
	if date == nil {
		return fmt.Errorf("parameter date is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kollavarsham) validateToGregorianDateParameters(date KollavarshamDate) error {
	if date == nil {
		return fmt.Errorf("parameter date is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_Kollavarsham) validateToGregorianDateFromSakaParameters(sakaDate SakaDate) error {
	if sakaDate == nil {
		return fmt.Errorf("parameter sakaDate is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Kollavarsham) validateSetSettingsParameters(val *Settings) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func validateNewKollavarshamParameters(options *Settings) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

