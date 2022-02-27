package location

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"

	badger "github.com/dgraph-io/badger/v3"
)

const bytesIn64Bits int = 8

func (place *Location) LoadHorizonOfLocation() {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	key := []byte(fmt.Sprintf("%02.4f,%02.4f", place.Latitude, place.Longitude))

	err = db.View(func(txn *badger.Txn) error {
		item, rerr := txn.Get(key)
		if rerr != nil {
			return rerr
		}

		log.Printf("Found horizon for location in DB for %s", string(key))

		item.Value(func(val []byte) error {
			for i := 0; i < horizonAngleResolution; i++ {
				position := i * bytesIn64Bits
				var bytes = make([]byte, bytesIn64Bits)
				for j := 0; j < bytesIn64Bits; j++ {
					bytes[j] = val[position+j]
				}
				place.Horizon[i] = math.Float64frombits(binary.LittleEndian.Uint64(bytes))
			}

			return nil
		})
		return nil
	})

	if err != nil {
		log.Printf("Could not find horizon for location in DB for %s", string(key))

		place.computeHorizon()

		err = db.Update(func(txn *badger.Txn) error {
			val := make([]byte, horizonAngleResolution*bytesIn64Bits)
			for i := 0; i < horizonAngleResolution; i++ {
				position := i * bytesIn64Bits
				bytes := make([]byte, bytesIn64Bits)
				binary.LittleEndian.PutUint64(bytes, math.Float64bits(place.Horizon[i]))
				for j := 0; j < bytesIn64Bits; j++ {
					val[position+j] = bytes[j]
				}
			}
			err = txn.Set(key, val)
			return err
		})

		if err != nil {
			panic(err)
		}
	}
}
