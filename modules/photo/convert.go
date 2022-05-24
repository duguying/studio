package photo

import (
	"io"
)

// Convert an interface to convert any photo as jpeg/webp
type Convert interface{
	io.Reader

	AsJPEG()
	AsWebp()
}