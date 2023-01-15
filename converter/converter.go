package converter

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"kollavarsham.BaseDate",
		reflect.TypeOf((*BaseDate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ahargana", GoGetter: "Ahargana"},
			_jsii_.MemberProperty{JsiiProperty: "date", GoGetter: "Date"},
			_jsii_.MemberProperty{JsiiProperty: "gregorianDate", GoGetter: "GregorianDate"},
			_jsii_.MemberProperty{JsiiProperty: "julianDay", GoGetter: "JulianDay"},
			_jsii_.MemberProperty{JsiiProperty: "mlWeekdayName", GoGetter: "MlWeekdayName"},
			_jsii_.MemberProperty{JsiiProperty: "month", GoGetter: "Month"},
			_jsii_.MemberProperty{JsiiProperty: "naksatra", GoGetter: "Naksatra"},
			_jsii_.MemberProperty{JsiiProperty: "sauraDivasa", GoGetter: "SauraDivasa"},
			_jsii_.MemberProperty{JsiiProperty: "sauraMasa", GoGetter: "SauraMasa"},
			_jsii_.MemberProperty{JsiiProperty: "sauraMasaName", GoGetter: "SauraMasaName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "weekdayName", GoGetter: "WeekdayName"},
			_jsii_.MemberProperty{JsiiProperty: "year", GoGetter: "Year"},
		},
		func() interface{} {
			return &jsiiProxy_BaseDate{}
		},
	)
	_jsii_.RegisterClass(
		"kollavarsham.JulianDate",
		reflect.TypeOf((*JulianDate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ahargana", GoGetter: "Ahargana"},
			_jsii_.MemberProperty{JsiiProperty: "date", GoGetter: "Date"},
			_jsii_.MemberProperty{JsiiProperty: "gregorianDate", GoGetter: "GregorianDate"},
			_jsii_.MemberProperty{JsiiProperty: "julianDay", GoGetter: "JulianDay"},
			_jsii_.MemberProperty{JsiiProperty: "mlWeekdayName", GoGetter: "MlWeekdayName"},
			_jsii_.MemberProperty{JsiiProperty: "month", GoGetter: "Month"},
			_jsii_.MemberProperty{JsiiProperty: "naksatra", GoGetter: "Naksatra"},
			_jsii_.MemberProperty{JsiiProperty: "sauraDivasa", GoGetter: "SauraDivasa"},
			_jsii_.MemberProperty{JsiiProperty: "sauraMasa", GoGetter: "SauraMasa"},
			_jsii_.MemberProperty{JsiiProperty: "sauraMasaName", GoGetter: "SauraMasaName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "weekdayName", GoGetter: "WeekdayName"},
			_jsii_.MemberProperty{JsiiProperty: "year", GoGetter: "Year"},
		},
		func() interface{} {
			j := jsiiProxy_JulianDate{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseDate)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"kollavarsham.Kollavarsham",
		reflect.TypeOf((*Kollavarsham)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "fromGregorianDate", GoMethod: "FromGregorianDate"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberMethod{JsiiMethod: "toGregorianDate", GoMethod: "ToGregorianDate"},
			_jsii_.MemberMethod{JsiiMethod: "toGregorianDateFromSaka", GoMethod: "ToGregorianDateFromSaka"},
		},
		func() interface{} {
			return &jsiiProxy_Kollavarsham{}
		},
	)
	_jsii_.RegisterClass(
		"kollavarsham.KollavarshamDate",
		reflect.TypeOf((*KollavarshamDate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ahargana", GoGetter: "Ahargana"},
			_jsii_.MemberProperty{JsiiProperty: "date", GoGetter: "Date"},
			_jsii_.MemberProperty{JsiiProperty: "gregorianDate", GoGetter: "GregorianDate"},
			_jsii_.MemberProperty{JsiiProperty: "julianDay", GoGetter: "JulianDay"},
			_jsii_.MemberProperty{JsiiProperty: "masaName", GoGetter: "MasaName"},
			_jsii_.MemberProperty{JsiiProperty: "mlMasaName", GoGetter: "MlMasaName"},
			_jsii_.MemberProperty{JsiiProperty: "mlNaksatraName", GoGetter: "MlNaksatraName"},
			_jsii_.MemberProperty{JsiiProperty: "mlWeekdayName", GoGetter: "MlWeekdayName"},
			_jsii_.MemberProperty{JsiiProperty: "month", GoGetter: "Month"},
			_jsii_.MemberProperty{JsiiProperty: "naksatra", GoGetter: "Naksatra"},
			_jsii_.MemberProperty{JsiiProperty: "naksatraName", GoGetter: "NaksatraName"},
			_jsii_.MemberProperty{JsiiProperty: "sakaDate", GoGetter: "SakaDate"},
			_jsii_.MemberProperty{JsiiProperty: "sauraDivasa", GoGetter: "SauraDivasa"},
			_jsii_.MemberProperty{JsiiProperty: "sauraMasa", GoGetter: "SauraMasa"},
			_jsii_.MemberProperty{JsiiProperty: "sauraMasaName", GoGetter: "SauraMasaName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "weekdayName", GoGetter: "WeekdayName"},
			_jsii_.MemberProperty{JsiiProperty: "year", GoGetter: "Year"},
		},
		func() interface{} {
			j := jsiiProxy_KollavarshamDate{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseDate)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"kollavarsham.MasaName",
		reflect.TypeOf((*MasaName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"kollavarsham.Naksatra",
		reflect.TypeOf((*Naksatra)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"kollavarsham.SakaDate",
		reflect.TypeOf((*SakaDate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "adhimasa", GoGetter: "Adhimasa"},
			_jsii_.MemberProperty{JsiiProperty: "ahargana", GoGetter: "Ahargana"},
			_jsii_.MemberProperty{JsiiProperty: "date", GoGetter: "Date"},
			_jsii_.MemberProperty{JsiiProperty: "fractionalTithi", GoGetter: "FractionalTithi"},
			_jsii_.MemberMethod{JsiiMethod: "generateKollavarshamDate", GoMethod: "GenerateKollavarshamDate"},
			_jsii_.MemberProperty{JsiiProperty: "gregorianDate", GoGetter: "GregorianDate"},
			_jsii_.MemberProperty{JsiiProperty: "julianDay", GoGetter: "JulianDay"},
			_jsii_.MemberProperty{JsiiProperty: "kaliYear", GoGetter: "KaliYear"},
			_jsii_.MemberProperty{JsiiProperty: "masaName", GoGetter: "MasaName"},
			_jsii_.MemberProperty{JsiiProperty: "mlWeekdayName", GoGetter: "MlWeekdayName"},
			_jsii_.MemberProperty{JsiiProperty: "month", GoGetter: "Month"},
			_jsii_.MemberProperty{JsiiProperty: "naksatra", GoGetter: "Naksatra"},
			_jsii_.MemberProperty{JsiiProperty: "naksatraName", GoGetter: "NaksatraName"},
			_jsii_.MemberProperty{JsiiProperty: "originalAhargana", GoGetter: "OriginalAhargana"},
			_jsii_.MemberProperty{JsiiProperty: "paksa", GoGetter: "Paksa"},
			_jsii_.MemberProperty{JsiiProperty: "sakaYear", GoGetter: "SakaYear"},
			_jsii_.MemberProperty{JsiiProperty: "sauraDivasa", GoGetter: "SauraDivasa"},
			_jsii_.MemberProperty{JsiiProperty: "sauraMasa", GoGetter: "SauraMasa"},
			_jsii_.MemberProperty{JsiiProperty: "sauraMasaName", GoGetter: "SauraMasaName"},
			_jsii_.MemberProperty{JsiiProperty: "sunriseHour", GoGetter: "SunriseHour"},
			_jsii_.MemberProperty{JsiiProperty: "sunriseMinute", GoGetter: "SunriseMinute"},
			_jsii_.MemberProperty{JsiiProperty: "tithi", GoGetter: "Tithi"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vikramaYear", GoGetter: "VikramaYear"},
			_jsii_.MemberProperty{JsiiProperty: "weekdayName", GoGetter: "WeekdayName"},
			_jsii_.MemberProperty{JsiiProperty: "year", GoGetter: "Year"},
		},
		func() interface{} {
			j := jsiiProxy_SakaDate{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseDate)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"kollavarsham.Settings",
		reflect.TypeOf((*Settings)(nil)).Elem(),
	)
}
