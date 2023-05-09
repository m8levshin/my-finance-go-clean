package utils

import "errors"

func CastAnyToString(reference *any) (*string, error) {
	value := *reference
	stringValue, ok := value.(string)
	if !ok {
		return nil, errors.New("//todo")
	}
	return &stringValue, nil
}
