package fortnox

import (
	"encoding/json"
	"strconv"
)

type Number float64

func (i *Number) UnmarshalJSON(data []byte) error {
	realNumber := 0.0
	err := json.Unmarshal(data, &realNumber)
	if err == nil {
		*i = Number(realNumber)
		return nil
	}

	// error, so maybe it isn't an int
	str := ""
	err = json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	if str == "" {
		*i = 0
		return nil
	}

	realNumber, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	i2 := Number(realNumber)
	*i = i2
	return nil
}

func (i Number) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(i))
}

type StringFloat float64

func (f *StringFloat) UnmarshalJSON(text []byte) (err error) {
	var flt float64
	err = json.Unmarshal(text, &flt)
	if err == nil {
		*f = StringFloat(flt)
		return err
	}

	// error, so try string
	var s string
	err = json.Unmarshal(text, &s)
	if err != nil {
		return err
	}

	flt, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*f = StringFloat(flt)
	return nil
}
