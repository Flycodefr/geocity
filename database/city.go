package database

type City struct {
	ID uint64 `sql:"type:bigint PRIMARY KEY`

	FullName   string
	PostalCode string

	Lat  float64
	Long float64
}

func (c City) String() string {
	return c.PostalCode + " " + c.FullName
}

func CityNew(FullName string, PostalCode string) *City {
	return &City{
		ID:         0,
		FullName:   FullName,
		PostalCode: PostalCode,
	}
}

func CityCount() int {
	db, _ := GetDataBase()
	defer db.Close()

	var count int
	db.Table("cities").Count(&count)

	return count
}
