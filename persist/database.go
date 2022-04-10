package persist

import (
	"encoding/binary"
	"math"

	"github.com/cloudsftp/Sunangel/dir"
	badger "github.com/dgraph-io/badger/v3"
)

const bytesIn64Bits int = 8

var (
	dbInitialized bool = false
	db            *badger.DB
)

func initializeDatabase() {
	if !dbInitialized {
		var err error
		db, err = badger.Open(badger.DefaultOptions(dir.GetStoreDir("badger")))
		if err != nil {
			panic(err)
		}

		dbInitialized = true
	}
}

func float64FromBytes(bytes []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(bytes))
}

func bytesFromFloat64(val float64, bytes []byte) {
	binary.LittleEndian.PutUint64(bytes, math.Float64bits(val))
}
