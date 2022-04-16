package persist

import (
	"crypto/sha512"

	"github.com/cloudsftp/Sunangel/horizon"
	"github.com/cloudsftp/Sunangel/location"
	"github.com/cloudsftp/Sunangel/util"
	badger "github.com/dgraph-io/badger/v3"
)

const horizonPrefix = "Horizon: "

func getHorizonStoreKey(place *location.Location, startRadius int) []byte {
	preKey := make([]byte, 3*bytesIn64Bits)

	util.BytesFromFloat64(place.Latitude, preKey[:bytesIn64Bits])
	util.BytesFromFloat64(place.Longitude, preKey[bytesIn64Bits:2*bytesIn64Bits])
	util.BytesFromFloat64(float64(startRadius), preKey[2*bytesIn64Bits:])

	keyHash := sha512.Sum512(preKey)
	return append([]byte(horizonPrefix), keyHash[:]...)
}

func AddHorizon(horizon *horizon.Horizon) {
	initializeDatabase()

	err := db.Update(func(txn *badger.Txn) error {
		return txn.Set(
			getHorizonStoreKey(horizon.Place, horizon.GetStartRadius()),
			horizon.AltitudeToBytes(),
		)
	})

	if err != nil {
		panic(err)
	}
}

func GetHorizon(place *location.Location, startRadius int) (*horizon.Horizon, error) {
	initializeDatabase()

	var hor *horizon.Horizon
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(getHorizonStoreKey(place, startRadius))
		if err != nil {
			return err
		}

		item.Value(func(val []byte) error {
			altitude, err := horizon.AltitudeFromBytes(val)
			if err != nil {
				return err
			}

			hor = horizon.NewHorizonWithAltitude(place, startRadius, altitude)
			return nil
		})
		return nil
	})

	return hor, err
}

func DeleteHorizonAll() {
	initializeDatabase()

	err := db.Update(func(txn *badger.Txn) error {
		iter := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iter.Close()
		for iter.Rewind(); iter.Valid(); iter.Next() {
			key := iter.Item().Key()

			if keyHasPrefix(key, horizonPrefix) {
				err := txn.Delete(key)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
