// +build windows

package goserialproxy


import (
	serial "github.com/TheDistractor/goserial"
	"io"
)

func init() {
}

type proxyGoSerial struct {
	ProxyPort io.ReadWriteCloser
}


//proxy for Windows - must refactor from rs232 options as both librarys are different
func newSerialProxy(devicename string, options Options) (interface{}, error) {

	var pm serial.ParityMode
	switch options.Parity {
		case PARITY_NONE:
			pm = serial.ParityNone
		case PARITY_ODD:
			pm = serial.ParityOdd
		case PARITY_EVEN:
			pm = serial.ParityEven
	}

	var sb serial.StopBits
	switch options.StopBits {
	case 1:
		sb = serial.StopBits1
	case 2:
		sb = serial.StopBits2
	}


	var bs serial.ByteSize
	switch options.DataBits {
	case 5:
		bs = serial.Byte5
	case 6:
		bs = serial.Byte6
	case 7:
		bs = serial.Byte7
	case 8:
		bs = serial.Byte8
	}

	//nb: units are different
	var rto uint32
	rto = uint32(options.Timeout) * uint32(100) //0.1s == 001 = 100 ms


	opt := &serial.Config{Name: devicename, Baud: int(options.BitRate), Size: bs, StopBits: sb, Parity: pm , ReadTimeout: rto }
	dev, err := serial.OpenPort(opt)

	pxy := &proxyGoSerial{ProxyPort:dev}


return pxy, err

}


func (p *proxyGoSerial) Close() (error) {
	return p.ProxyPort.Close()
}

func (p *proxyGoSerial) Read(b []byte) (n int, err error)  {
	return p.ProxyPort.Read(b)
}

func (p *proxyGoSerial) Write(b []byte) (n int, err error)  {
	return p.ProxyPort.Write(b)
}

func (p *proxyGoSerial) SetDTR(level bool) error {
	return SerialPort(p.ProxyPort.(SerialPort)).SetDTR(level)
}

func (p *proxyGoSerial) SetRTS(level bool) error {
	return SerialPort(p.ProxyPort.(SerialPort)).SetRTS(level)
}
