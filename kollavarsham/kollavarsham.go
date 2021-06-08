// Convert Gregorian date to Kollavarsham date and vice versa
package kollavarsham

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kollavarsham/kollavarsham-go/kollavarsham/v2/jsii"
)

// Serves as the base class for both the {@link JulianDate} and {@link KollavarshamDate} classes.
//
// **INTERNAL/PRIVATE**
type BaseDate interface {
	Ahargana() *float64
	SetAhargana(val *float64)
	Date() *float64
	SetDate(val *float64)
	GregorianDate() *time.Time
	SetGregorianDate(val *time.Time)
	JulianDay() *float64
	SetJulianDay(val *float64)
	MlWeekdayName() *string
	Month() *float64
	SetMonth(val *float64)
	Naksatra() *Naksatra
	SetNaksatra(val *Naksatra)
	SauraDivasa() *float64
	SetSauraDivasa(val *float64)
	SauraMasa() *float64
	SetSauraMasa(val *float64)
	SauraMasaName() *string
	WeekdayName() *string
	Year() *float64
	SetYear(val *float64)
	ToString() *string
}

// The jsii proxy struct for BaseDate
type jsiiProxy_BaseDate struct {
	_ byte // padding
}

func (j *jsiiProxy_BaseDate) Ahargana() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ahargana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDate) Date() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"date",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDate) GregorianDate() *time.Time {
	var returns *time.Time
	_jsii_.Get(
		j,
		"gregorianDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDate) JulianDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"julianDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDate) MlWeekdayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlWeekdayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDate) Month() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"month",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDate) Naksatra() *Naksatra {
	var returns *Naksatra
	_jsii_.Get(
		j,
		"naksatra",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDate) SauraDivasa() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sauraDivasa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDate) SauraMasa() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sauraMasa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDate) SauraMasaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sauraMasaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDate) WeekdayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weekdayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDate) Year() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"year",
		&returns,
	)
	return returns
}


func NewBaseDate_Override(b BaseDate, year *float64, month *float64, date *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"kollavarsham.BaseDate",
		[]interface{}{year, month, date},
		b,
	)
}

func (j *jsiiProxy_BaseDate) SetAhargana(val *float64) {
	_jsii_.Set(
		j,
		"ahargana",
		val,
	)
}

func (j *jsiiProxy_BaseDate) SetDate(val *float64) {
	_jsii_.Set(
		j,
		"date",
		val,
	)
}

func (j *jsiiProxy_BaseDate) SetGregorianDate(val *time.Time) {
	_jsii_.Set(
		j,
		"gregorianDate",
		val,
	)
}

func (j *jsiiProxy_BaseDate) SetJulianDay(val *float64) {
	_jsii_.Set(
		j,
		"julianDay",
		val,
	)
}

func (j *jsiiProxy_BaseDate) SetMonth(val *float64) {
	_jsii_.Set(
		j,
		"month",
		val,
	)
}

func (j *jsiiProxy_BaseDate) SetNaksatra(val *Naksatra) {
	_jsii_.Set(
		j,
		"naksatra",
		val,
	)
}

func (j *jsiiProxy_BaseDate) SetSauraDivasa(val *float64) {
	_jsii_.Set(
		j,
		"sauraDivasa",
		val,
	)
}

func (j *jsiiProxy_BaseDate) SetSauraMasa(val *float64) {
	_jsii_.Set(
		j,
		"sauraMasa",
		val,
	)
}

func (j *jsiiProxy_BaseDate) SetYear(val *float64) {
	_jsii_.Set(
		j,
		"year",
		val,
	)
}

// Returns the month names object that has Saka, Saura and Kollavarsham (English & Malayalam) month names for the specified index `masaNumber`.
func BaseDate_GetMasaName(masaNumber *float64) *MasaName {
	_init_.Initialize()

	var returns *MasaName

	_jsii_.StaticInvoke(
		"kollavarsham.BaseDate",
		"getMasaName",
		[]interface{}{masaNumber},
		&returns,
	)

	return returns
}

// Converts the Date to a nicely formatted string with year, month and date.
func (b *jsiiProxy_BaseDate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a Julian date's year, month and day `toGregorianDateFromSaka` method of the {@link Kollavarsham} class returns an instance of this type for dates older than `1st January 1583 AD`.
//
// **INTERNAL/PRIVATE**
type JulianDate interface {
	BaseDate
	Ahargana() *float64
	SetAhargana(val *float64)
	Date() *float64
	SetDate(val *float64)
	GregorianDate() *time.Time
	SetGregorianDate(val *time.Time)
	JulianDay() *float64
	SetJulianDay(val *float64)
	MlWeekdayName() *string
	Month() *float64
	SetMonth(val *float64)
	Naksatra() *Naksatra
	SetNaksatra(val *Naksatra)
	SauraDivasa() *float64
	SetSauraDivasa(val *float64)
	SauraMasa() *float64
	SetSauraMasa(val *float64)
	SauraMasaName() *string
	WeekdayName() *string
	Year() *float64
	SetYear(val *float64)
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

func (j *jsiiProxy_JulianDate) SetAhargana(val *float64) {
	_jsii_.Set(
		j,
		"ahargana",
		val,
	)
}

func (j *jsiiProxy_JulianDate) SetDate(val *float64) {
	_jsii_.Set(
		j,
		"date",
		val,
	)
}

func (j *jsiiProxy_JulianDate) SetGregorianDate(val *time.Time) {
	_jsii_.Set(
		j,
		"gregorianDate",
		val,
	)
}

func (j *jsiiProxy_JulianDate) SetJulianDay(val *float64) {
	_jsii_.Set(
		j,
		"julianDay",
		val,
	)
}

func (j *jsiiProxy_JulianDate) SetMonth(val *float64) {
	_jsii_.Set(
		j,
		"month",
		val,
	)
}

func (j *jsiiProxy_JulianDate) SetNaksatra(val *Naksatra) {
	_jsii_.Set(
		j,
		"naksatra",
		val,
	)
}

func (j *jsiiProxy_JulianDate) SetSauraDivasa(val *float64) {
	_jsii_.Set(
		j,
		"sauraDivasa",
		val,
	)
}

func (j *jsiiProxy_JulianDate) SetSauraMasa(val *float64) {
	_jsii_.Set(
		j,
		"sauraMasa",
		val,
	)
}

func (j *jsiiProxy_JulianDate) SetYear(val *float64) {
	_jsii_.Set(
		j,
		"year",
		val,
	)
}

// Returns the month names object that has Saka, Saura and Kollavarsham (English & Malayalam) month names for the specified index `masaNumber`.
func JulianDate_GetMasaName(masaNumber *float64) *MasaName {
	_init_.Initialize()

	var returns *MasaName

	_jsii_.StaticInvoke(
		"kollavarsham.JulianDate",
		"getMasaName",
		[]interface{}{masaNumber},
		&returns,
	)

	return returns
}

// Converts the Date to a nicely formatted string with year, month and date.
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

// The Kollavarsham class implements all the public APIs of the library.
//
// Create a new instance of this class, passing in the relevant options and call methods on the instance.
//
// TODO: EXAMPLE
//
type Kollavarsham interface {
	Settings() *Settings
	SetSettings(val *Settings)
	FromGregorianDate(date *time.Time) KollavarshamDate
	ToGregorianDate(date KollavarshamDate) *time.Time
	ToGregorianDateFromSaka(sakaDate SakaDate) KollavarshamDate
}

// The jsii proxy struct for Kollavarsham
type jsiiProxy_Kollavarsham struct {
	_ byte // padding
}

func (j *jsiiProxy_Kollavarsham) Settings() *Settings {
	var returns *Settings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


func NewKollavarsham(options *Settings) Kollavarsham {
	_init_.Initialize()

	j := jsiiProxy_Kollavarsham{}

	_jsii_.Create(
		"kollavarsham.Kollavarsham",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewKollavarsham_Override(k Kollavarsham, options *Settings) {
	_init_.Initialize()

	_jsii_.Create(
		"kollavarsham.Kollavarsham",
		[]interface{}{options},
		k,
	)
}

func (j *jsiiProxy_Kollavarsham) SetSettings(val *Settings) {
	_jsii_.Set(
		j,
		"settings",
		val,
	)
}

// Converts a Gregorian date to the equivalent Kollavarsham date, respecting the current configuration.
//
// Returns: Converted date
//
// TODO: EXAMPLE
//
func (k *jsiiProxy_Kollavarsham) FromGregorianDate(date *time.Time) KollavarshamDate {
	var returns KollavarshamDate

	_jsii_.Invoke(
		k,
		"fromGregorianDate",
		[]interface{}{date},
		&returns,
	)

	return returns
}

// Converts a Kollavarsham date (an instance of {@link kollavarshamDate}) to its equivalent Gregorian date, respecting the current configuration.
//
// This method Will return {@link JulianDate} object for any date before 1st January 1583 AD and
// [Date](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date) objects for later dates.
//
//   **This API has not been implemented yet**
//
// Returns: Converted date
func (k *jsiiProxy_Kollavarsham) ToGregorianDate(date KollavarshamDate) *time.Time {
	var returns *time.Time

	_jsii_.Invoke(
		k,
		"toGregorianDate",
		[]interface{}{date},
		&returns,
	)

	return returns
}

// Converts a Saka date (an instance of {@link sakaDate}) to its equivalent Gregorian date, respecting the current configuration.
//
// This method Will return {@link JulianDate} object for any date before 1st January 1583 AD and
// [Date](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date) objects for later dates.
//
// Returns: Converted date
func (k *jsiiProxy_Kollavarsham) ToGregorianDateFromSaka(sakaDate SakaDate) KollavarshamDate {
	var returns KollavarshamDate

	_jsii_.Invoke(
		k,
		"toGregorianDateFromSaka",
		[]interface{}{sakaDate},
		&returns,
	)

	return returns
}

// Represents a Kollavarsham date's year, month and date.
type KollavarshamDate interface {
	BaseDate
	Ahargana() *float64
	SetAhargana(val *float64)
	Date() *float64
	SetDate(val *float64)
	GregorianDate() *time.Time
	SetGregorianDate(val *time.Time)
	JulianDay() *float64
	SetJulianDay(val *float64)
	MasaName() *string
	MlMasaName() *string
	MlNaksatraName() *string
	MlWeekdayName() *string
	Month() *float64
	SetMonth(val *float64)
	Naksatra() *Naksatra
	SetNaksatra(val *Naksatra)
	NaksatraName() *string
	SakaDate() SakaDate
	SetSakaDate(val SakaDate)
	SauraDivasa() *float64
	SetSauraDivasa(val *float64)
	SauraMasa() *float64
	SetSauraMasa(val *float64)
	SauraMasaName() *string
	WeekdayName() *string
	Year() *float64
	SetYear(val *float64)
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

func (j *jsiiProxy_KollavarshamDate) SetAhargana(val *float64) {
	_jsii_.Set(
		j,
		"ahargana",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate) SetDate(val *float64) {
	_jsii_.Set(
		j,
		"date",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate) SetGregorianDate(val *time.Time) {
	_jsii_.Set(
		j,
		"gregorianDate",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate) SetJulianDay(val *float64) {
	_jsii_.Set(
		j,
		"julianDay",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate) SetMonth(val *float64) {
	_jsii_.Set(
		j,
		"month",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate) SetNaksatra(val *Naksatra) {
	_jsii_.Set(
		j,
		"naksatra",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate) SetSakaDate(val SakaDate) {
	_jsii_.Set(
		j,
		"sakaDate",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate) SetSauraDivasa(val *float64) {
	_jsii_.Set(
		j,
		"sauraDivasa",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate) SetSauraMasa(val *float64) {
	_jsii_.Set(
		j,
		"sauraMasa",
		val,
	)
}

func (j *jsiiProxy_KollavarshamDate) SetYear(val *float64) {
	_jsii_.Set(
		j,
		"year",
		val,
	)
}

// Returns the month names object that has Saka, Saura and Kollavarsham (English & Malayalam) month names for the specified index `masaNumber`.
func KollavarshamDate_GetMasaName(masaNumber *float64) *MasaName {
	_init_.Initialize()

	var returns *MasaName

	_jsii_.StaticInvoke(
		"kollavarsham.KollavarshamDate",
		"getMasaName",
		[]interface{}{masaNumber},
		&returns,
	)

	return returns
}

// Converts the Date to a nicely formatted string with year, month and date.
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

type MasaName struct {
	EnMalayalam *string `json:"enMalayalam"`
	MlMalayalam *string `json:"mlMalayalam"`
	Saka *string `json:"saka"`
	Saura *string `json:"saura"`
}

type Naksatra struct {
	EnMalayalam *string `json:"enMalayalam"`
	MlMalayalam *string `json:"mlMalayalam"`
	Saka *string `json:"saka"`
}

// Represents an Saka date's year, month and paksa and tithi.
type SakaDate interface {
	BaseDate
	Adhimasa() *string
	SetAdhimasa(val *string)
	Ahargana() *float64
	SetAhargana(val *float64)
	Date() *float64
	SetDate(val *float64)
	FractionalTithi() *float64
	SetFractionalTithi(val *float64)
	GregorianDate() *time.Time
	SetGregorianDate(val *time.Time)
	JulianDay() *float64
	SetJulianDay(val *float64)
	KaliYear() *float64
	SetKaliYear(val *float64)
	MasaName() *string
	MlWeekdayName() *string
	Month() *float64
	SetMonth(val *float64)
	Naksatra() *Naksatra
	SetNaksatra(val *Naksatra)
	NaksatraName() *string
	OriginalAhargana() *float64
	SetOriginalAhargana(val *float64)
	Paksa() *string
	SetPaksa(val *string)
	SakaYear() *float64
	SauraDivasa() *float64
	SetSauraDivasa(val *float64)
	SauraMasa() *float64
	SetSauraMasa(val *float64)
	SauraMasaName() *string
	SunriseHour() *float64
	SetSunriseHour(val *float64)
	SunriseMinute() *float64
	SetSunriseMinute(val *float64)
	Tithi() *float64
	VikramaYear() *float64
	WeekdayName() *string
	Year() *float64
	SetYear(val *float64)
	GenerateKollavarshamDate() KollavarshamDate
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

func (j *jsiiProxy_SakaDate) SetAdhimasa(val *string) {
	_jsii_.Set(
		j,
		"adhimasa",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetAhargana(val *float64) {
	_jsii_.Set(
		j,
		"ahargana",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetDate(val *float64) {
	_jsii_.Set(
		j,
		"date",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetFractionalTithi(val *float64) {
	_jsii_.Set(
		j,
		"fractionalTithi",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetGregorianDate(val *time.Time) {
	_jsii_.Set(
		j,
		"gregorianDate",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetJulianDay(val *float64) {
	_jsii_.Set(
		j,
		"julianDay",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetKaliYear(val *float64) {
	_jsii_.Set(
		j,
		"kaliYear",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetMonth(val *float64) {
	_jsii_.Set(
		j,
		"month",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetNaksatra(val *Naksatra) {
	_jsii_.Set(
		j,
		"naksatra",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetOriginalAhargana(val *float64) {
	_jsii_.Set(
		j,
		"originalAhargana",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetPaksa(val *string) {
	_jsii_.Set(
		j,
		"paksa",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetSauraDivasa(val *float64) {
	_jsii_.Set(
		j,
		"sauraDivasa",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetSauraMasa(val *float64) {
	_jsii_.Set(
		j,
		"sauraMasa",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetSunriseHour(val *float64) {
	_jsii_.Set(
		j,
		"sunriseHour",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetSunriseMinute(val *float64) {
	_jsii_.Set(
		j,
		"sunriseMinute",
		val,
	)
}

func (j *jsiiProxy_SakaDate) SetYear(val *float64) {
	_jsii_.Set(
		j,
		"year",
		val,
	)
}

// Returns the month names object that has Saka, Saura and Kollavarsham (English & Malayalam) month names for the specified index `masaNumber`.
func SakaDate_GetMasaName(masaNumber *float64) *MasaName {
	_init_.Initialize()

	var returns *MasaName

	_jsii_.StaticInvoke(
		"kollavarsham.SakaDate",
		"getMasaName",
		[]interface{}{masaNumber},
		&returns,
	)

	return returns
}

// Generates an instance of {@link KollavarshamDate} from this instance of SakaDate.
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

// Converts the Date to a nicely formatted string with year, month and date.
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

type Settings struct {
	Latitude *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	System *string `json:"system"`
}

