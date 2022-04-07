package location

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"

	badger "github.com/dgraph-io/badger/v3"
)

const bytesIn64Bits int = 8

var (
	dbInitialized bool = false
	db            *badger.DB
)

func (place *Location) LoadHorizonOfLocation() {
	ok := place.tryLoadHorizon()

	if ok {
		log.Printf("Found horizon for location in DB for %s", string(place.getHorizonStoreKey()))
	} else {
		log.Printf("Could not find horizon for location in DB for %s", string(place.getHorizonStoreKey()))
		place.RecomputeHorizon()
	}
}

func (place *Location) RecomputeHorizon() {
	place.computeHorizon()

	place.storeHorizon()
}

func (place Location) getHorizonStoreKey() []byte {
	key := fmt.Sprintf("%02.4f,%02.4f", place.Latitude, place.Longitude)
	return []byte(key)
}

func initializeDatabase() {
	if !dbInitialized {
		var err error
		db, err = badger.Open(badger.DefaultOptions("/tmp/badger"))
		if err != nil {
			panic(err)
		}

		dbInitialized = true
	}

}

func (place *Location) tryLoadHorizon() bool {
	initializeDatabase()

	err := db.View(func(txn *badger.Txn) error {
		item, rerr := txn.Get(place.getHorizonStoreKey())
		if rerr != nil {
			return rerr
		}

		item.Value(func(val []byte) error {
			place.Horizon = horizonArrayFromBytes(val)
			return nil
		})
		return nil
	})

	return err == nil
}

func (place Location) storeHorizon() {
	initializeDatabase()

	err := db.Update(func(txn *badger.Txn) error {
		return txn.Set(place.getHorizonStoreKey(), bytesFromHorizonArray(place.Horizon))
	})

	if err != nil {
		panic(err)
	}
}

func bytesFromHorizonArray(horizon horizonArray) []byte {
	val := make([]byte, horizonAngleResolution*bytesIn64Bits)

	for i := 0; i < horizonAngleResolution; i++ {
		position := i * bytesIn64Bits
		bytes := make([]byte, bytesIn64Bits)
		binary.LittleEndian.PutUint64(bytes, math.Float64bits(horizon[i]))

		for j := 0; j < bytesIn64Bits; j++ {
			val[position+j] = bytes[j]
		}
	}

	return val
}

func horizonArrayFromBytes(val []byte) horizonArray {
	horizon := horizonArray{}

	for i := 0; i < horizonAngleResolution; i++ {
		var bytes = make([]byte, bytesIn64Bits)

		position := i * bytesIn64Bits
		for j := 0; j < bytesIn64Bits; j++ {
			bytes[j] = val[position+j]
		}
		horizon[i] = math.Float64frombits(binary.LittleEndian.Uint64(bytes))
	}

	return horizon
}
