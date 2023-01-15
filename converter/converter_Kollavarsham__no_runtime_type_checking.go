//go:build no_runtime_type_checking

// Convert Gregorian date to Kollavarsham date and vice versa
package converter

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_Kollavarsham) validateFromGregorianDateParameters(date *time.Time) error {
	return nil
}

func (k *jsiiProxy_Kollavarsham) validateToGregorianDateParameters(date KollavarshamDate) error {
	return nil
}

func (k *jsiiProxy_Kollavarsham) validateToGregorianDateFromSakaParameters(sakaDate SakaDate) error {
	return nil
}

func (j *jsiiProxy_Kollavarsham) validateSetSettingsParameters(val *Settings) error {
	return nil
}

func validateNewKollavarshamParameters(options *Settings) error {
	return nil
}

