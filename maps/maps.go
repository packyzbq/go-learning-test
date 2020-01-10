package maps

import (
	"errors"
)

func Search(dict map[string]string, key string) string {
	return dict[key]
}

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(key string) (string, error) {
	v, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

func (d Dictionary) Update(key, value string) (string, error) {
	_, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	d[key] = value
	return "", nil
}
