//go:build no_runtime_type_checking

// Convert Gregorian date to Kollavarsham date and vice versa
package converter

// Building without runtime type checking enabled, so all the below just return nil

func validateKollavarshamDate_GetMasaNameParameters(masaNumber *float64) error {
	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetAharganaParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetDateParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetGregorianDateParameters(val *time.Time) error {
	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetJulianDayParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetMonthParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetNaksatraParameters(val *Naksatra) error {
	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetSakaDateParameters(val SakaDate) error {
	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetSauraDivasaParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetSauraMasaParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_KollavarshamDate) validateSetYearParameters(val *float64) error {
	return nil
}

