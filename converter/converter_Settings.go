// Convert Gregorian date to Kollavarsham date and vice versa
package converter


type Settings struct {
	Latitude *float64 `field:"required" json:"latitude" yaml:"latitude"`
	Longitude *float64 `field:"required" json:"longitude" yaml:"longitude"`
	System *string `field:"required" json:"system" yaml:"system"`
}

