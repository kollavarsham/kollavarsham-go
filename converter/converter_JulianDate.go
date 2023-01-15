// Convert Gregorian date to Kollavarsham date and vice versa
package converter

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kollavarsham/kollavarsham-go/converter/v2/jsii"
)

// Represents a Julian date's year, month and day `toGregorianDateFromSaka` method of the {@link Kollavarsham} class returns an instance of this type for dates older than `1st January 1583 AD`.
//
// **INTERNAL/PRIVATE**.
type JulianDate interface {
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
	// Returns the weekday (in Malayalam) for the current instance of date.
	MlWeekdayName() *string
	// The month corresponding to this instance of the date.
	Month() *float64
	SetMonth(val *float64)
	Naksatra() *Naksatra
	SetNaksatra(val *Naksatra)
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

// The jsii proxy struct for JulianDate
type jsiiProxy_JulianDate struct {
	jsiiProxy_BaseDate
}

func (j *jsiiProxy_JulianDate) Ahargana() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ahargana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JulianDate) Date() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"date",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JulianDate) GregorianDate() *time.Time {
	var returns *time.Time
	_jsii_.Get(
		j,
		"gregorianDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JulianDate) JulianDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"julianDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JulianDate) MlWeekdayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlWeekdayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JulianDate) Month() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"month",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JulianDate) Naksatra() *Naksatra {
	var returns *Naksatra
	_jsii_.Get(
		j,
		"naksatra",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JulianDate) SauraDivasa() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sauraDivasa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JulianDate) SauraMasa() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sauraMasa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JulianDate) SauraMasaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sauraMasaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JulianDate) WeekdayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weekdayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JulianDate) Year() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"year",
		&returns,
	)
	return returns
}


func NewJulianDate(year *float64, month *float64, day *float64) JulianDate {
	_init_.Initialize()

	j := jsiiProxy_JulianDate{}

	_jsii_.Create(
		"kollavarsham.JulianDate",
		[]interface{}{year, month, day},
		&j,
	)

	return &j
}

func NewJulianDate_Override(j JulianDate, year *float64, month *float64, day *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"kollavarsham.JulianDate",
		[]interface{}{year, month, day},
		j,
	)
}

func (j *jsiiProxy_JulianDate)SetAhargana(val *float64) {
	if err := j.validateSetAharganaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ahargana",
		val,
	)
}

func (j *jsiiProxy_JulianDate)SetDate(val *float64) {
	if err := j.validateSetDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"date",
		val,
	)
}

func (j *jsiiProxy_JulianDate)SetGregorianDate(val *time.Time) {
	if err := j.validateSetGregorianDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gregorianDate",
		val,
	)
}

func (j *jsiiProxy_JulianDate)SetJulianDay(val *float64) {
	if err := j.validateSetJulianDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"julianDay",
		val,
	)
}

func (j *jsiiProxy_JulianDate)SetMonth(val *float64) {
	if err := j.validateSetMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"month",
		val,
	)
}

func (j *jsiiProxy_JulianDate)SetNaksatra(val *Naksatra) {
	if err := j.validateSetNaksatraParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"naksatra",
		val,
	)
}

func (j *jsiiProxy_JulianDate)SetSauraDivasa(val *float64) {
	if err := j.validateSetSauraDivasaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sauraDivasa",
		val,
	)
}

func (j *jsiiProxy_JulianDate)SetSauraMasa(val *float64) {
	if err := j.validateSetSauraMasaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sauraMasa",
		val,
	)
}

func (j *jsiiProxy_JulianDate)SetYear(val *float64) {
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
func JulianDate_GetMasaName(masaNumber *float64) *MasaName {
	_init_.Initialize()

	if err := validateJulianDate_GetMasaNameParameters(masaNumber); err != nil {
		panic(err)
	}
	var returns *MasaName

	_jsii_.StaticInvoke(
		"kollavarsham.JulianDate",
		"getMasaName",
		[]interface{}{masaNumber},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JulianDate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

