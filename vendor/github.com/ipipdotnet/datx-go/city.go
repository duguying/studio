package datx

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

var ErrIPv4Format = errors.New("ipv4 format error")
var ErrNotFound = errors.New("not found")

// NewCity ...
func NewCity(name string) (*City, error) {
	db := &City{}

	if err := db.load(name); err != nil {
		return nil, err
	}

	return db, nil
}

// City ...
type City struct {
	file  *os.File
	index []byte
	data  []byte
}

func (db *City) load(fn string) error {
	var err error
	db.file, err = os.Open(fn)
	if err != nil {
		return err
	}
	b4 := make([]byte, 4)
	_, err = db.file.Read(b4)
	if err != nil {
		return err
	}
	off := int(binary.BigEndian.Uint32(b4))

	_, err = db.file.Seek(262148, 0)
	if err != nil {
		return err
	}

	l := off - 262148 - 262144
	db.index = make([]byte, l)
	_, err = db.file.Read(db.index)
	if err != nil {
		return err
	}

	db.data, err = ioutil.ReadAll(db.file)
	if err != nil {
		return err
	}
	return nil
}

// Find ...
func (db *City) Find(s string) ([]string, error) {
	ipv := net.ParseIP(s)
	if ipv == nil || ipv.To4() == nil {
		return nil, ErrIPv4Format
	}

	low := 0
	mid := 0
	high := int(len(db.index)/9) - 1

	val := binary.BigEndian.Uint32(ipv.To4())

	for low <= high {
		mid = int((low + high) / 2)
		pos := mid * 9

		var start uint32
		if mid > 0 {
			pos1 := (mid - 1) * 9
			start = binary.BigEndian.Uint32(db.index[pos1:pos1+4]) + 1
		} else {
			start = 0
		}

		end := binary.BigEndian.Uint32(db.index[pos : pos+4])

		if val < start {
			high = mid - 1
		} else if val > end {
			low = mid + 1
		} else {
			off := int(binary.LittleEndian.Uint32([]byte{
				db.index[pos+4],
				db.index[pos+5],
				db.index[pos+6],
				0,
			}))
			l := int(db.index[pos+7])*256 + int(db.index[pos+8])

			return strings.Split(string(db.data[off:off+l]), "\t"), nil
		}
	}

	return nil, ErrNotFound
}

func (db *City) FindLocation(s string) (Location, error) {
	var loc Location

	a, e := db.Find(s)
	if e != nil {
		return loc, e
	}

	loc.Country = a[0]
	loc.Province = a[1]
	loc.City = a[2]
	if len(a) < 10 {
		return loc, nil
	}

	loc.Organization = a[3]
	loc.ISP = a[4]
	loc.Latitude = a[5]
	loc.Longitude = a[6]
	loc.TimeZone = a[7]
	loc.TimeZone2 = a[8]
	loc.CityCode = a[9]
	loc.PhonePrefix = a[10]
	loc.CountryCode = a[11]
	loc.ContinentCode = a[12]

	if len(a) == 17 {
		loc.IDC = a[13]
		loc.BaseStation = a[14]
		loc.CountryCode3 = a[15]
		if a[15] == "1" {
			loc.EuropeanUnion = true
		} else {
			loc.EuropeanUnion = false
		}
	} else if len(a) == 20 {
		loc.IDC = a[13]
		loc.BaseStation = a[14]
		loc.CountryCode3 = a[15]
		if a[15] == "1" {
			loc.EuropeanUnion = true
		} else {
			loc.EuropeanUnion = false
		}
		if a[19] == "ANYCAST" {
			loc.Anycast = true
		}
	}

	return loc, nil
}

type Location struct {
	Country       string
	Province      string
	City          string
	Organization  string
	ISP           string
	Latitude      string
	Longitude     string
	TimeZone      string
	TimeZone2     string
	CityCode      string
	PhonePrefix   string
	CountryCode   string
	ContinentCode string
	IDC           string // IDC | VPN
	BaseStation   string // WIFI | BS (Base Station)
	CountryCode3  string
	EuropeanUnion bool
	CurrencyCode  string
	CurrencyName  string
	Anycast       bool
}

func (l Location) ToJSON() []byte {
	all, err := json.Marshal(l)
	if err == nil {
		return all
	}

	return nil
}
