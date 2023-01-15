//go:build no_runtime_type_checking

// Convert Gregorian date to Kollavarsham date and vice versa
package converter

// Building without runtime type checking enabled, so all the below just return nil

func validateJulianDate_GetMasaNameParameters(masaNumber *float64) error {
	return nil
}

func (j *jsiiProxy_JulianDate) validateSetAharganaParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_JulianDate) validateSetDateParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_JulianDate) validateSetGregorianDateParameters(val *time.Time) error {
	return nil
}

func (j *jsiiProxy_JulianDate) validateSetJulianDayParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_JulianDate) validateSetMonthParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_JulianDate) validateSetNaksatraParameters(val *Naksatra) error {
	return nil
}

func (j *jsiiProxy_JulianDate) validateSetSauraDivasaParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_JulianDate) validateSetSauraMasaParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_JulianDate) validateSetYearParameters(val *float64) error {
	return nil
}

