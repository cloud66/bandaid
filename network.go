package bandaid

import (
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

type IPType int
type IPSource string

type IPInfo struct {
	Address    string `json:"address"`
	Source     string `json:"source"`
	IsFallback bool   `json:"is_fallback"`

	HasCity bool `json:"has_city"`
	City    struct {
		GeoNameID uint              `json:"geoname_id"`
		Names     map[string]string `json:"names"`
	} `json:"city"`
	Continent struct {
		Code      string            `json:"code"`
		GeoNameID uint              `json:"geoname_id"`
		Names     map[string]string `json:"names"`
	} `json:"continent"`
	Country struct {
		GeoNameID         uint              `json:"geoname_id"`
		IsInEuropeanUnion bool              `json:"is_in_european_union"`
		IsoCode           string            `json:"iso_code"`
		Names             map[string]string `json:"names"`
	} `json:"country"`
	Location struct {
		AccuracyRadius uint16  `json:"accuracy_radius"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		MetroCode      uint    `json:"metro_code"`
		TimeZone       string  `json:"time_zone"`
	} `json:"location"`
	Postal struct {
		Code string `json:"code"`
	} `json:"postal"`
	RegisteredCountry struct {
		GeoNameID         uint              `json:"geoname_id"`
		IsInEuropeanUnion bool              `json:"is_in_european_union"`
		IsoCode           string            `json:"iso_code"`
		Names             map[string]string `json:"names"`
	} `json:"registered_country"`
	RepresentedCountry struct {
		GeoNameID         uint              `json:"geoname_id"`
		IsInEuropeanUnion bool              `json:"is_in_european_union"`
		IsoCode           string            `json:"iso_code"`
		Names             map[string]string `json:"names"`
		Type              string            `json:"type"`
	} `json:"represented_country"`
	Subdivisions []struct {
		GeoNameID uint              `json:"geoname_id"`
		IsoCode   string            `json:"iso_code"`
		Names     map[string]string `json:"names"`
	} `json:"subdivisions"`
	Traits struct {
		IsAnonymousProxy    bool `json:"is_anonymous_proxy"`
		IsSatelliteProvider bool `json:"is_satellite_provider"`
	} `json:"traits"`

	HasASN bool `json:"has_asn"`
	ASN    struct {
		AutonomousSystemNumber       uint   `json:"autonomous_system_number"`
		AutonomousSystemOrganization string `json:"autonomous_system_organization"`
	} `json:"asn"`

	HasAnonymousIP bool `json:"has_anonymous_ip"`
	AnonymousIP    struct {
		IsAnonymous       bool `json:"is_anonymous"`
		IsAnonymousVPN    bool `json:"is_anonymous_vpn"`
		IsHostingProvider bool `json:"is_hosting_provider"`
		IsPublicProxy     bool `json:"is_public_proxy"`
		IsTorExitNode     bool `json:"is_tor_exit_node"`
	} `json:"anonymous_ip"`
}

func StripPort(hostport string) string {
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return hostport
	}
	if i := strings.IndexByte(hostport, ']'); i != -1 {
		return strings.TrimPrefix(hostport[:i], "[")
	}
	return hostport[:colon]
}

func GetMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
