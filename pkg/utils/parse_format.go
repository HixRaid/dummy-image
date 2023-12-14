package utils

import (
	"errors"

	"github.com/hixraid/dummy-image/pkg/data"
)

var formats = map[string]data.ImageFormat{
	"png":  data.PNG,
	"jpeg": data.JPEG,
	"svg":  data.SVG,
}

func ParseFormat(s string) (data.ImageFormat, error) {
	if format, ok := formats[s]; ok {
		return format, nil
	}
	return 0, errors.New("invalid image format")
}