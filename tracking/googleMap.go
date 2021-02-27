package tracking

import (
	"context"
	"fmt"
	"googlemaps.github.io/maps"
)

//CM is a google cloud maps
type CM struct {
	Client *maps.Client
}

// Location is a gps
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// GeocodeAdress provided Location data from gcp maps geocoder api
func (cm *CM) GeocodeAdress(address string) (Location, error) {

	var loc Location

	r := &maps.GeocodingRequest{
		Address: address,
	}

	res, err := cm.Client.Geocode(context.Background(), r)
	if err != nil || len(res) == 0 {
		return loc, fmt.Errorf("res Geocode err: %v", err)
	}

	loc.Lat = res[0].Geometry.Location.Lat
	loc.Lng = res[0].Geometry.Location.Lng

	return loc, nil
}


