package goserialproxy


func init() {

}

//TODO: consolidate baud between both librarys

//TODO: bring in Posix


// Base Options taken from github.com/chimera/rs232

// Serial port options.
//
// Valid bit rates (linux):
// 200, 300, 600, 1200, 1800, 2400, 4800, 9600, 19200, 38400, 57600,
// 115200, 230400, 460800, 500000, 576000, 921600, 1000000, 1152000,
// 1500000, 2000000, 2500000, 3000000, 3500000, 4000000.
//
// Valid bit rates (OS X):
// 200, 300, 600, 1200, 1800, 2400, 4800, 7200, 9600, 14400, 19200, 28800,
// 38400, 57600, 78600, 115200, 230400
type Options struct {
	BitRate  uint32     // number of bits per second (baud)
	DataBits uint8      // number of data bits (5, 6, 7, 8)
	StopBits uint8      // number of stop bits (1, 2)
	Parity   ParityMode // none, odd, even
	Timeout  uint8      // read timeout in units of 0.1 seconds (0 = disable)
}

type ParityMode uint8

// Parity modes
const (
	PARITY_NONE = ParityMode(0)
	PARITY_ODD  = ParityMode(1)
	PARITY_EVEN = ParityMode(2)
)



//TODO: open the interfaces up as they are restrospectively added
//DONE: minimal interface for my needs now are Close,Read,Write,SetRTS,SetDTR
type SerialPort interface {
//	Open(devicename string, opt Options) (port *SerialPortProxy, err error)
//	String() string							//TODO:
	Close() error							//rs323, goserial
	Read(b []byte) (n int, err error)       //rs232, goserial
	Write(b []byte) (n int, err error)      //rs232, goserial
//	GetRTS() (bool, error)					//TODO:
//	GetCTS() (bool, error)					//TODO:
//	GetDTR() (bool, error)					//TODO:
//	GetDSR() (bool, error)					//TODO:
	SetDTR(level bool) error
	SetRTS(level bool) error
//	BytesAvailable() (int, error)			//TODO:

}




//======== X-platform stub for interim inclusion of both rs232 and goserial ========
//TODO: either push windows stuff to rs232 after refactor
//      or pull and the needed stuff from rs232 into goserial
func NewSerialProxy(devicename string, options Options) (SerialPort, error) {

	//platform specific stub
	pxy, err := newSerialProxy(devicename, options)

	//convert to proxy interface
	return pxy.(SerialPort), err

}



