// Convert Gregorian date to Kollavarsham date and vice versa
package converter


type Naksatra struct {
	EnMalayalam *string `field:"required" json:"enMalayalam" yaml:"enMalayalam"`
	MlMalayalam *string `field:"required" json:"mlMalayalam" yaml:"mlMalayalam"`
	Saka *string `field:"required" json:"saka" yaml:"saka"`
}

