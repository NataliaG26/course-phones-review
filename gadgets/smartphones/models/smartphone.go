package models

//smarthphones model structure for smartphone
type Smartphone struct {
	Id int64 //similar to a long in java
	Name string
	Price int
	CountryOrigin string
	Os string
}


// CMD + alt + L // gofmt -w .