package location

import (
	"log"
	"net/http"

	"github.com/tkrajina/go-elevations/geoelevations"
)

var (
	srtmInitialized bool = false
	srtm            *geoelevations.Srtm
)

func initSrtm() {
	log.Printf("Initialiting Srtm")

	var err error
	srtm, err = geoelevations.NewSrtm(http.DefaultClient)
	if err != nil {
		panic(err)
	}

	srtmInitialized = true
}

// GetElevation returns the elevation of the given location
func (a Location) GetElevation() float64 {
	if !srtmInitialized {
		initSrtm()
	}

	ele, err := srtm.GetElevation(http.DefaultClient, a.Latitude, a.Longitude)
	if err != nil {
		panic(err)
	}
	return float64(ele)
}
