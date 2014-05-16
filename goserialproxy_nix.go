// +build linux darwin,!cgo

package goserialproxy


import (
	"github.com/chimera/rs232"
)


func newSerialProxy(devicename string, options Options) (interface{}, error) {

	opt := rs232.Options{BitRate: options.BitRate, DataBits: options.DataBits, StopBits: options.StopBits, Timeout:options.Timeout}
	dev, err := rs232.Open(devicename, opt)

	return dev, err

}

