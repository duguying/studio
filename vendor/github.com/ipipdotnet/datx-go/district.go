package datx

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

// NewDistrict ...
func NewDistrict(fn string) (*District, error) {

	db := &District{}

	if err := db.load(fn); err != nil {
		return nil, err
	}

	return db, nil
}

// District ...
type District struct {
	file *os.File

	index []byte
	data  []byte
}

func (db *District) load(fn string) error {
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
	db.index = make([]byte, off-262148-262144)
	_, err = db.file.Read(db.index)
	if err != nil {
		return err
	}
	db.data, err = ioutil.ReadAll(db.file)
	if err != nil {
		return err
	}
	//	fmt.Println(len(db.data))
	return nil
}

// Find ...
func (db *District) Find(s string) ([]string, error) {

	ipv := net.ParseIP(s)
	if ipv == nil {
		return nil, fmt.Errorf("%s", "ip format error.")
	}

	low := 0
	high := int(len(db.index)/13) - 1
	mid := 0

	val := binary.BigEndian.Uint32(ipv.To4())

	for low <= high {
		mid = int((low + high) / 2)
		pos := mid * 13

		start := binary.BigEndian.Uint32(db.index[pos : pos+4])
		end := binary.BigEndian.Uint32(db.index[pos+4 : pos+8])

		if val < start {
			high = mid - 1
		} else if val > end {
			low = mid + 1
		} else {

			off := int(binary.LittleEndian.Uint32(db.index[pos+8 : pos+12]))

			return strings.Split(string(db.data[off:off+int(db.index[pos+12])]), "\t"), nil
		}
	}
	return nil, fmt.Errorf("%s", "not found")
}
