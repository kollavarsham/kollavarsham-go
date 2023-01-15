// Convert Gregorian date to Kollavarsham date and vice versa
package converter

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kollavarsham/kollavarsham-go/converter/v2/jsii"
)

// Represents an Saka date's year, month and paksa and tithi.
type SakaDate interface {
	BaseDate
	// The Adhimasa (`Adhika Masa`) prefix corresponding to this instance of the date.
	//
	// **Set separately after an instance is created**.
	Adhimasa() *string
	SetAdhimasa(val *string)
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
	// The fractional `Tithi`corresponding to this instance of the date.
	//
	// **Set separately after an instance is created**.
	FractionalTithi() *float64
	SetFractionalTithi(val *float64)
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
	// The Kali year corresponding to this instance of the date.
	//
	// **Set separately after an instance is created**.
	KaliYear() *float64
	SetKaliYear(val *float64)
	// Returns the month name for this instance of SakaDate.
	MasaName() *string
	// Returns the weekday (in Malayalam) for the current instance of date.
	MlWeekdayName() *string
	// The month corresponding to this instance of the date.
	Month() *float64
	SetMonth(val *float64)
	Naksatra() *Naksatra
	SetNaksatra(val *Naksatra)
	// Returns the Saka Naksatra name for this instance of SakaDate.
	NaksatraName() *string
	// The original ahargana passed in to the celestial calculations (TODO: Not sure why we need this!?).
	OriginalAhargana() *float64
	SetOriginalAhargana(val *float64)
	// The Paksha/Paksa corresponding to this instance of the date.
	//
	// Paksha (or pakṣa: Sanskrit: पक्ष), refers to a fortnight or a lunar phase in a month of the Hindu lunar calendar.
	//
	// _Source_: https://en.wikipedia.org/wiki/Paksha
	Paksa() *string
	SetPaksa(val *string)
	// Returns the Saka year on this instance of SakaDate (same as the underlyiung `year` property from the {@link BaseDate} class).
	SakaYear() *float64
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
	// The hour part from the sunrise time for this date.
	//
	// **Set separately after an instance is created**.
	SunriseHour() *float64
	SetSunriseHour(val *float64)
	// The minute part from the sunrise time for this date.
	//
	// **Set separately after an instance is created**.
	SunriseMinute() *float64
	SetSunriseMinute(val *float64)
	// Returns the Tithi on this instance of SakaDate (same as the underlyiung `date` property from the {@link BaseDate} class).
	//
	// In Vedic timekeeping, a tithi (also spelled thithi) is a lunar day, or the time it takes for the longitudinal angle between the Moon and the Sun to increase by 12°.
	// Tithis begin at varying times of day and vary in duration from approximately 19 to approximately 26 hours.
	//
	// _Source_: https://en.wikipedia.org/wiki/Tithi
	Tithi() *float64
	// Returns the Vikrama year corresponding to the Saka year of this instance.
	VikramaYear() *float64
	// Returns the weekday (in English) for the current instance of date.
	WeekdayName() *string
	// The year corresponding to this instance of the date.
	Year() *float64
	SetYear(val *float64)
	// Generates an instance of {@link KollavarshamDate} from this instance of SakaDate.
	GenerateKollavarshamDate() KollavarshamDate
	// Converts the Date to a nicely formatted string with year, month and date.
	ToString() *string
}

// The jsii proxy struct for SakaDate
type jsiiProxy_SakaDate struct {
	jsiiProxy_BaseDate
}

func (j *jsiiProxy_SakaDate) Adhimasa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adhimasa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) Ahargana() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ahargana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) Date() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"date",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) FractionalTithi() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fractionalTithi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) GregorianDate() *time.Time {
	var returns *time.Time
	_jsii_.Get(
		j,
		"gregorianDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) JulianDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"julianDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) KaliYear() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kaliYear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) MasaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) MlWeekdayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlWeekdayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) Month() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"month",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) Naksatra() *Naksatra {
	var returns *Naksatra
	_jsii_.Get(
		j,
		"naksatra",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) NaksatraName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"naksatraName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) OriginalAhargana() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originalAhargana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) Paksa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paksa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) SakaYear() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sakaYear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) SauraDivasa() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sauraDivasa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) SauraMasa() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sauraMasa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) SauraMasaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sauraMasaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) SunriseHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sunriseHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) SunriseMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sunriseMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) Tithi() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tithi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) VikramaYear() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vikramaYear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) WeekdayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weekdayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SakaDate) Year() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"year",
		&returns,
	)
	return returns
}


func NewSakaDate(year *float64, month *float64, tithi *float64, paksa *string) SakaDate {
	_init_.Initialize()

	j := jsiiProxy_SakaDate{}

	_jsii_.Create(
		"kollavarsham.SakaDate",
		[]interface{}{year, month, tithi, paksa},
		&j,
	)

	return &j
}

func NewSakaDate_Override(s SakaDate, year *float64, month *float64, tithi *float64, paksa *string) {
	_init_.Initialize()

	_jsii_.Create(
		"kollavarsham.SakaDate",
		[]interface{}{year, month, tithi, paksa},
		s,
	)
}

func (j *jsiiProxy_SakaDate)SetAdhimasa(val *string) {
	if err := j.validateSetAdhimasaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adhimasa",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetAhargana(val *float64) {
	if err := j.validateSetAharganaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ahargana",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetDate(val *float64) {
	if err := j.validateSetDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"date",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetFractionalTithi(val *float64) {
	if err := j.validateSetFractionalTithiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fractionalTithi",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetGregorianDate(val *time.Time) {
	if err := j.validateSetGregorianDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gregorianDate",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetJulianDay(val *float64) {
	if err := j.validateSetJulianDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"julianDay",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetKaliYear(val *float64) {
	if err := j.validateSetKaliYearParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kaliYear",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetMonth(val *float64) {
	if err := j.validateSetMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"month",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetNaksatra(val *Naksatra) {
	if err := j.validateSetNaksatraParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"naksatra",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetOriginalAhargana(val *float64) {
	if err := j.validateSetOriginalAharganaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originalAhargana",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetPaksa(val *string) {
	if err := j.validateSetPaksaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paksa",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetSauraDivasa(val *float64) {
	if err := j.validateSetSauraDivasaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sauraDivasa",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetSauraMasa(val *float64) {
	if err := j.validateSetSauraMasaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sauraMasa",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetSunriseHour(val *float64) {
	if err := j.validateSetSunriseHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sunriseHour",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetSunriseMinute(val *float64) {
	if err := j.validateSetSunriseMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sunriseMinute",
		val,
	)
}

func (j *jsiiProxy_SakaDate)SetYear(val *float64) {
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
func SakaDate_GetMasaName(masaNumber *float64) *MasaName {
	_init_.Initialize()

	if err := validateSakaDate_GetMasaNameParameters(masaNumber); err != nil {
		panic(err)
	}
	var returns *MasaName

	_jsii_.StaticInvoke(
		"kollavarsham.SakaDate",
		"getMasaName",
		[]interface{}{masaNumber},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SakaDate) GenerateKollavarshamDate() KollavarshamDate {
	var returns KollavarshamDate

	_jsii_.Invoke(
		s,
		"generateKollavarshamDate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SakaDate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

