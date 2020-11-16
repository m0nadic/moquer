package util

import "github.com/brianvoe/gofakeit/v5"

// Faker methods ...
func (p Payload) UUID() string {
	return gofakeit.UUID()
}

func (p Payload) Name() string {
	return gofakeit.Name()
}

func (p Payload) Noun() string {
	return gofakeit.Noun()
}