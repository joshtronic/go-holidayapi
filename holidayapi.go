package holidayapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type V1 struct {
	Url string
	Key string
}

func NewV1(key string) *V1 {
	v1 := &V1{
		Key: key,
		Url: "https://holidayapi.com/v1/holidays?",
	}

	return v1
}

func (v1 *V1) Holidays(args map[string]interface{}) (ResponseHoliday, error) {
	var data ResponseHoliday

	if _, ok := args["key"]; !ok {
		args["key"] = v1.Key
	}

	params := url.Values{}

	for k, v := range args {
		params.Add(k, v.(string))
	}

	resp, err := http.Get(v1.Url + params.Encode())

	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return data, err
	}

	json.Unmarshal([]byte(string(body)), &data)

	if resp.StatusCode != 200 {
		if data.Error == "" {
			data.Error = "Unknown error."
		}
	}

	return data, nil
}
