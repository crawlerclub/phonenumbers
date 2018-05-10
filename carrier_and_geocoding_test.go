package phonenumbers

import (
	"testing"
)

func TestGetCarrierForNumber(t *testing.T) {
	tests := []prefixMapTestCases{
		{num: "+8613702032331", lang: "en", expected: "China Mobile"},
		{num: "+8613702032331", lang: "zh", expected: "中国移动"},
		{num: "+6281377468527", lang: "en", expected: "Telkomsel"},
		{num: "+8613323241342", lang: "en", expected: "China Telecom"},
		{num: "+61491570156", lang: "en", expected: "Telstra"},
		{num: "+917999999543", lang: "en", expected: "Reliance Jio"},
	}
	for _, test := range tests {
		number, err := Parse(test.num, "ZZ")
		if err != nil {
			t.Errorf("Failed to parse number %s: %s", test.num, err)
		}
		carrier, err := GetCarrierForNumber(number, test.lang)
		if err != nil {
			t.Errorf("Failed to getCarrier for the number %s: %s", test.num, err)
		}
		if test.expected != carrier {
			t.Errorf("Expected '%s', got '%s' for '%s'", test.expected, carrier, test.num)
		}
	}
}

func TestGetGeocodingForNumber(t *testing.T) {
	tests := []prefixMapTestCases{
		{num: "+8613702032331", lang: "en", expected: "Tianjin"},
		{num: "+8613702032331", lang: "zh", expected: "天津市"},
		{num: "+863197785050", lang: "zh", expected: "河北省邢台市"},
		{num: "+8613323241342", lang: "en", expected: "Baoding, Hebei"},
		{num: "+917999999543", lang: "en", expected: "Ahmedabad Local, Gujarat"},
		{num: "+17047181840", lang: "en", expected: "North Carolina"},
		{num: "+12542462158", lang: "en", expected: "Texas"},
		{num: "+16193165996", lang: "en", expected: "California"},
	}
	for _, test := range tests {
		number, err := Parse(test.num, "ZZ")
		if err != nil {
			t.Errorf("Failed to parse number %s: %s", test.num, err)
		}
		geocoding, err := GetGeocodingForNumber(number, test.lang)
		if err != nil {
			t.Errorf("Failed to getGeocoding for the number %s: %s", test.num, err)
		}
		if test.expected != geocoding {
			t.Errorf("Expected '%s', got '%s' for '%s'", test.expected, geocoding, test.num)
		}
	}
}
