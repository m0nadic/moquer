package util

import (
	"github.com/brianvoe/gofakeit/v5"
	"text/template"
)




var FakerFuncs = template.FuncMap{
	"UUID":         gofakeit.UUID,
	"name":         gofakeit.Name,
	"firstName":    gofakeit.FirstName,
	"lastName":     gofakeit.LastName,
	"gender":       gofakeit.Gender,
	"SSN":          gofakeit.SSN,
	"email":        gofakeit.Email,
	"phone":        gofakeit.Phone,
	"phoneF":       gofakeit.PhoneFormatted,
	"namePrefix":   gofakeit.NamePrefix,
	"nameSuffix":   gofakeit.NameSuffix,
	"city":         gofakeit.City,
	"country":      gofakeit.Country,
	"countryAbr":   gofakeit.CountryAbr,
	"state":        gofakeit.State,
	"stateAbr":     gofakeit.StateAbr,
	"street":       gofakeit.Street,
	"streetName":   gofakeit.StreetName,
	"streetNumber": gofakeit.StreetNumber,
	"streetPrefix": gofakeit.StreetPrefix,
	"streetSuffix": gofakeit.StreetSuffix,
	"zip":          gofakeit.Zip,
	"latitude":     gofakeit.Latitude,
	"longitude":    gofakeit.Longitude,
}
