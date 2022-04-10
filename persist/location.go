package persist

import (
	"strings"

	"github.com/cloudsftp/Sunangel/location"
	badger "github.com/dgraph-io/badger/v3"
)

const locationPrefix = "Location: "

func GetLocations() map[string]*location.Location {
	initializeDatabase()

	locations := map[string]*location.Location{}

	err := db.View(func(txn *badger.Txn) error {
		iter := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iter.Close()
		for iter.Rewind(); iter.Valid(); iter.Next() {
			item := iter.Item()

			key := string(item.Key())
			name := strings.Replace(key, locationPrefix, "", 1)

			if key == locationPrefix+name { // Only if key starts with locationPrefix
				var location *location.Location
				item.Value(func(val []byte) error {
					location = locationFromBytes(val)
					return nil
				})

				locations[name] = location
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return locations
}

func GetLocation(name string) (*location.Location, error) {
	initializeDatabase()

	var location *location.Location
	err := db.View(func(txn *badger.Txn) error {
		key := []byte(locationPrefix + name)
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			location = locationFromBytes(val)
			return nil
		})
		return err
	})

	if err != nil {
		return nil, err
	}
	return location, nil
}

func AddLocation(name string, place *location.Location) {
	initializeDatabase()

	err := db.Update(func(txn *badger.Txn) error {
		key := []byte(locationPrefix + name)
		val := bytesFromLocation(place)
		return txn.Set(key, val)
	})
	if err != nil {
		panic(err)
	}
}

func locationFromBytes(val []byte) *location.Location {
	latitude := float64FromBytes(val[:bytesIn64Bits])
	longitude := float64FromBytes(val[bytesIn64Bits:])

	return location.NewLocation(latitude, longitude)
}

func bytesFromLocation(place *location.Location) []byte {
	bytes := make([]byte, 2*bytesIn64Bits)

	bytesFromFloat64(place.Latitude, bytes[:bytesIn64Bits])
	bytesFromFloat64(place.Longitude, bytes[bytesIn64Bits:])

	return bytes
}
