package persist

import (
	"log"
	"strings"

	"github.com/cloudsftp/Sunangel/dir"
	badger "github.com/dgraph-io/badger/v3"
)

const bytesIn32Bits int = 4
const bytesIn64Bits int = 2 * bytesIn32Bits

var (
	dbInitialized bool = false
	db            *badger.DB
)

func initializeDatabase() {
	if !dbInitialized {
		log.Printf("Initialiting Badger")

		var err error
		db, err = badger.Open(badger.DefaultOptions(dir.GetStoreDir("badger")))
		if err != nil {
			panic(err)
		}

		dbInitialized = true
	}
}

func keyHasPrefix(key []byte, prefix string) bool {
	postfix := keyRemovePrefix(key, prefix)

	return string(key) == prefix+postfix
}

func keyRemovePrefix(key []byte, prefix string) string {
	keyStr := string(key)
	return strings.Replace(keyStr, prefix, "", 1)
}
