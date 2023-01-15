// Convert Gregorian date to Kollavarsham date and vice versa
package converter

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kollavarsham/kollavarsham-go/converter/v2/jsii"
)

// Represents a Kollavarsham date's year, month and date.
type KollavarshamDate interface {
	BaseDate
	// The `Ahargana` corresponding to this instance of the date. **Set separately after an instance is created**.
	//
	// In Sanskrit `ahoratra` means one full day and `gana` means count.
	// Hence, the Ahargana on any given day stands for the number of lunar days that have elapsed starting from an epoch.
	//
	// _Source_: http://cs.annauniv.edu/insight/Reading%20Materials/astro/sharptime/ahargana.htm
	Ahargana() *float64
	SetAhargana(val *float64)
	// The date corresponding to this instance of the date.
	Date() *float64
	SetDate(val *float64)
	// The gregorian date corresponding to this instance of the date.
	//
	// **Set separately after an instance is created**.
	GregorianDate() *time.Time
	SetGregorianDate(val *time.Time)
	// The `Julian Day` corresponding to this instance of the date.
	//
	// **Set separately after an instance is created**
	// Julian day is the continuous count of days since the beginning of the Julian Period used primarily by astronomers.
	//
	// _Source_: https://en.wikipedia.org/wiki/Julian_day
	JulianDay() *float64
	SetJulianDay(val *float64)
	// Returns the Kollavarsham month name (in English) for this instance of date.
	MasaName() *string
	// Returns the Kollavarsham month name (in Malayalam) for this instance of date.
	MlMasaName() *string
	// Returns the Kollavarsham Naksatra name (in Malayalam) for this instance of date.
	MlNaksatraName() *string
	// Returns the weekday (in Malayalam) for the current instance of date.
	MlWeekdayName() *string
	// The month corresponding to this instance of the date.
	Month() *float64
	SetMonth(val *float64)
	Naksatra() *Naksatra
	SetNaksatra(val *Naksatra)
	// Returns the Kollavarsham Naksatra name (in English) for this instance date.
	NaksatraName() *string
	SakaDate() SakaDate
	SetSakaDate(val SakaDate)
	// The `Saura Divasa` (Solar Calendar Day) for this instance of the date.
	//
	// **Set separately after an instance is created**.
	SauraDivasa() *float64
	SetSauraDivasa(val *float64)
	// The `Saura Masa` (Solar Calendar Month) for this instance of the date.
	//
	// **Set separately after an instance is created**.
	SauraMasa() *float64
	SetSauraMasa(val *float64)
	// Returns the Saura Masa name for the current instance of date.
	SauraMasaName() *string
	// Returns the weekday (in English) for the current instance of date.
	WeekdayName() *string
	// The year corresponding to this instance of the date.
	Year() *float64
	SetYear(val *float64)
	// Converts the Date to a nicely formatted string with year, month and date.
	ToString() *string
}

// The jsii proxy struct for KollavarshamDate
type jsiiProxy_KollavarshamDate struct {
	jsiiProxy_BaseDate
}

func (j *jsiiProxy_KollavarshamDate) Ahargana() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ahargana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) Date() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"date",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) GregorianDate() *time.Time {
	var returns *time.Time
	_jsii_.Get(
		j,
		"gregorianDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) JulianDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"julianDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) MasaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) MlMasaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlMasaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) MlNaksatraName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlNaksatraName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) MlWeekdayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlWeekdayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) Month() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"month",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) Naksatra() *Naksatra {
	var returns *Naksatra
	_jsii_.Get(
		j,
		"naksatra",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) NaksatraName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"naksatraName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) SakaDate() SakaDate {
	var returns SakaDate
	_jsii_.Get(
		j,
		"sakaDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) SauraDivasa() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sauraDivasa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) SauraMasa() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sauraMasa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) SauraMasaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sauraMasaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) WeekdayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weekdayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KollavarshamDate) Year() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"year",
		&returns,
	)
	return returns
}


func NewKollavarshamDate(year *float64, month *float64, day *float64) KollavarshamDate {
	_init_.Initialize()

	j := jsiiProxy_KollavarshamDate{}

	_jsii_.Create(
		"kollavarsham.KollavarshamDate",
		[]interface{}{year, month, day},
		&j,
	)

	return &j
}

func NewKollavarshamDate_Override(k KollavarshamDate, year *float64, month *float64, day *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"kollavarsham.KollavarshamDate",
		[]interface{}{year, month, day},
		k,
	)
}

func (j *jsiiProxy_KollavarshamDate)SetAhargana(val *float64) {
	if err := j.validateSetAharganaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ahargana",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate)SetDate(val *float64) {
	if err := j.validateSetDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"date",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate)SetGregorianDate(val *time.Time) {
	if err := j.validateSetGregorianDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gregorianDate",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate)SetJulianDay(val *float64) {
	if err := j.validateSetJulianDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"julianDay",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate)SetMonth(val *float64) {
	if err := j.validateSetMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"month",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate)SetNaksatra(val *Naksatra) {
	if err := j.validateSetNaksatraParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"naksatra",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate)SetSakaDate(val SakaDate) {
	if err := j.validateSetSakaDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sakaDate",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate)SetSauraDivasa(val *float64) {
	if err := j.validateSetSauraDivasaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sauraDivasa",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate)SetSauraMasa(val *float64) {
	if err := j.validateSetSauraMasaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sauraMasa",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate)SetYear(val *float64) {
	if err := j.validateSetYearParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"year",
		val,
	)
}

// Returns the month names object that has Saka, Saura and Kollavarsham (English & Malayalam) month names for the specified index `masaNumber`.
func KollavarshamDate_GetMasaName(masaNumber *float64) *MasaName {
	_init_.Initialize()

	if err := validateKollavarshamDate_GetMasaNameParameters(masaNumber); err != nil {
		panic(err)
	}
	var returns *MasaName

	_jsii_.StaticInvoke(
		"kollavarsham.KollavarshamDate",
		"getMasaName",
		[]interface{}{masaNumber},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KollavarshamDate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

