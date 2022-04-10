package persist

import (
	"github.com/cloudsftp/Sunangel/horizon"
	badger "github.com/dgraph-io/badger/v3"
)

func getHorizonStoreKey(horizon *horizon.Horizon) []byte {
	// TODO: reimplement
	// key := fmt.Sprintf("%02.4f,%02.4f", place.Latitude, place.Longitude)
	return []byte("nil")
}

func tryLoadHorizon(horizon *horizon.Horizon) bool {
	initializeDatabase()

	err := db.View(func(txn *badger.Txn) error {
		item, rerr := txn.Get(getHorizonStoreKey(horizon))
		if rerr != nil {
			return rerr
		}

		item.Value(func(val []byte) error {
			horizonFromBytes(val)
			return nil
		})
		return nil
	})

	return err == nil
}

func storeHorizon(horizon *horizon.Horizon) {
	initializeDatabase()

	err := db.Update(func(txn *badger.Txn) error {
		return txn.Set(getHorizonStoreKey(horizon), bytesFromHorizon(horizon))
	})

	if err != nil {
		panic(err)
	}
}

func bytesFromHorizon(horizon *horizon.Horizon) []byte {
	val := make([]byte, 0*bytesIn64Bits)

	// TODO: reimplement
	/*
		for i := 0; i < horizonAngleResolution; i++ {
			position := i * bytesIn64Bits
			bytes := make([]byte, bytesIn64Bits)
			binary.LittleEndian.PutUint64(bytes, math.Float64bits(horizon[i]))

			for j := 0; j < bytesIn64Bits; j++ {
				val[position+j] = bytes[j]
			}
		}
	*/

	return val
}

func horizonFromBytes(val []byte) *horizon.Horizon {
	horizon := &horizon.Horizon{}

	// TODO: reimplement
	/*
		for i := 0; i < horizonAngleResolution; i++ {
			var bytes = make([]byte, bytesIn64Bits)

			position := i * bytesIn64Bits
			for j := 0; j < bytesIn64Bits; j++ {
				bytes[j] = val[position+j]
			}
			horizon[i] = math.Float64frombits(binary.LittleEndian.Uint64(bytes))
		}
	*/

	return horizon
}
