package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("not found")

func (d Dictionary) Search(text string) (string, error) {
	result, ok := d[text]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

func (d Dictionary) Add(key string, value string) {
	d[key] = value
}
