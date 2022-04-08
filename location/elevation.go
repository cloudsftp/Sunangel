package location

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/tkrajina/go-elevations/geoelevations"
)

var (
	srtmInitialized bool = false
	srtm            *geoelevations.Srtm
)

func initSrtm() {
	log.Printf("Initialiting Srtm")

	storage, err := geoelevations.NewLocalFileSrtmStorage(path.Join(os.Getenv("HOME"), ".geoelevations"))
	if err != nil {
		panic(err)
	}

	srtm, err = geoelevations.NewSrtmWithCustomStorage(http.DefaultClient, storage)
	if err != nil {
		panic(err)
	}

	srtmInitialized = true
}

// GetElevation returns the elevation of the given location
func (loc Location) GetElevation() float64 {
	if !srtmInitialized {
		initSrtm()
	}

	ele, err := srtm.GetElevation(http.DefaultClient, loc.Latitude, loc.Longitude)
	if err != nil {
		panic(err)
	}

	return float64(ele)
}
