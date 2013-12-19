package main

type StringSlice []string

func (slice StringSlice) Head() string {
	if len(slice) > 0 {
		return slice[0]
	}

	return ""
}

func (slice StringSlice) Map(fn func(string) string) StringSlice {
	var newSlice StringSlice

	for _, s := range slice {
		newSlice = append(newSlice, fn(s))
	}

	return newSlice
}
