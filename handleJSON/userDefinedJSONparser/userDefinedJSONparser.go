package userDefinedJSONparser

// When you have to process another time format other than RFC,
// You can deal with it implementing json.Marshaler and Unmarshaler.

import "time"

type RFC822ZTime struct {
	time.Time
}

func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := rt.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}

func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
	if err != nil {
		return nil
	}
	*rt = RFC822ZTime{t}
	return nil
}

type Item int
type Order struct {
	ID          string      `json:"id"`
	DateOrdered RFC822ZTime `json:"date_ordered"`
	CustomerID  string      `json:"customer_id"`
	Items       []Item      `json:"items"`
}

func UsingUserDefinedJSONParser() {

}
