// Package open is a geo-golang based MapRequest Open geocode/reverse geocode client
package open

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"strings"
)

type baseURL string

type geocodeResponse struct {
	Results []struct {
		Locations []struct {
			LatLng                                     geo.Location
			Street, AdminArea5, AdminArea3, AdminArea1 string
		}
	}
}

// Geocoder constructs MapRequest Open geocoder
func Geocoder(key string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL("http://open.mapquestapi.com/geocoding/v1/*?key=" + key + "&location="),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func (b baseURL) GeocodeURL(address string) string {
	return strings.Replace(string(b), "*", "address", 1) + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return strings.Replace(string(b), "*", "reverse", 1) + fmt.Sprintf("%f,%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() geo.Location {
	return r.Results[0].Locations[0].LatLng
}

func (r *geocodeResponse) Address() (address string) {
	p := r.Results[0].Locations[0]
	if p.AdminArea1 != "" {
		address = p.Street + ", " + p.AdminArea5 + ", " + p.AdminArea3 + ", " + p.AdminArea1
	}
	return
}
