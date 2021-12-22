package utils

import "github.com/mitchellh/mapstructure"

func Mapstructure(input interface{}, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName: "json",
		Result:  result,
	}
	decoder, err := mapstructure.NewDecoder(config)
	decoder.Decode(input)
	return err
}
