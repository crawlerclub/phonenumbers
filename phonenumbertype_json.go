package phonenumbers

import "encoding/json"

func (i PhoneNumberType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i PhoneNumberType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}
