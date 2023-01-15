//go:build no_runtime_type_checking

// Convert Gregorian date to Kollavarsham date and vice versa
package converter

// Building without runtime type checking enabled, so all the below just return nil

func validateBaseDate_GetMasaNameParameters(masaNumber *float64) error {
	return nil
}

func (j *jsiiProxy_BaseDate) validateSetAharganaParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_BaseDate) validateSetDateParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_BaseDate) validateSetGregorianDateParameters(val *time.Time) error {
	return nil
}

func (j *jsiiProxy_BaseDate) validateSetJulianDayParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_BaseDate) validateSetMonthParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_BaseDate) validateSetNaksatraParameters(val *Naksatra) error {
	return nil
}

func (j *jsiiProxy_BaseDate) validateSetSauraDivasaParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_BaseDate) validateSetSauraMasaParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_BaseDate) validateSetYearParameters(val *float64) error {
	return nil
}

