// Convert Gregorian date to Kollavarsham date and vice versa
package converter

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kollavarsham/kollavarsham-go/converter/v2/jsii"
)

// The Kollavarsham class implements all the public APIs of the library.
//
// Create a new instance of this class, passing in the relevant options and call methods on the instance.
//
// Example:
//   const Kollavarsham = require('kollavarsham');
//
//   const options = {
//    system: 'SuryaSiddhanta',
//    latitude: 10,
//    longitude: 76.2
//   };
//
//   const kollavarsham = new Kollavarsham(options);
//
//   let todayInMalayalamEra = kollavarsham.fromGregorianDate(new Date());
//
//   let today = kollavarsham.toGregorianDate(todayInMalayalamEra);  // Not implemented yet
//
type Kollavarsham interface {
	Settings() *Settings
	SetSettings(val *Settings)
	// Converts a Gregorian date to the equivalent Kollavarsham date, respecting the current configuration.
	//
	// Returns: Converted date.
	//
	// Example:
	//   const Kollavarsham = require('Kollavarsham');
	//   const kollavarsham = new Kollavarsham();
	//   let today = kollavarsham.fromGregorianDate(new Date(1979, 4, 22));
	//
	FromGregorianDate(date *time.Time) KollavarshamDate
	// Converts a Kollavarsham date (an instance of {@link kollavarshamDate}) to its equivalent Gregorian date, respecting the current configuration.
	//
	// This method Will return {@link JulianDate} object for any date before 1st January 1583 AD and
	// [Date](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date) objects for later dates.
	//
	// **This API has not been implemented yet**.
	//
	// Returns: Converted date.
	ToGregorianDate(date KollavarshamDate) *time.Time
	// Converts a Saka date (an instance of {@link sakaDate}) to its equivalent Gregorian date, respecting the current configuration.
	//
	// This method Will return {@link JulianDate} object for any date before 1st January 1583 AD and
	// [Date](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date) objects for later dates.
	//
	// Returns: Converted date.
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

	if err := validateNewKollavarshamParameters(options); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_Kollavarsham)SetSettings(val *Settings) {
	if err := j.validateSetSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"settings",
		val,
	)
}

func (k *jsiiProxy_Kollavarsham) FromGregorianDate(date *time.Time) KollavarshamDate {
	if err := k.validateFromGregorianDateParameters(date); err != nil {
		panic(err)
	}
	var returns KollavarshamDate

	_jsii_.Invoke(
		k,
		"fromGregorianDate",
		[]interface{}{date},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kollavarsham) ToGregorianDate(date KollavarshamDate) *time.Time {
	if err := k.validateToGregorianDateParameters(date); err != nil {
		panic(err)
	}
	var returns *time.Time

	_jsii_.Invoke(
		k,
		"toGregorianDate",
		[]interface{}{date},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kollavarsham) ToGregorianDateFromSaka(sakaDate SakaDate) KollavarshamDate {
	if err := k.validateToGregorianDateFromSakaParameters(sakaDate); err != nil {
		panic(err)
	}
	var returns KollavarshamDate

	_jsii_.Invoke(
		k,
		"toGregorianDateFromSaka",
		[]interface{}{sakaDate},
		&returns,
	)

	return returns
}

