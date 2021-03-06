package packets

import (
	"fmt"
	"io"
	"sync"
)

var _disconnectPacketPool = sync.Pool{
	New: func() interface{} {
		return &DisconnectPacket{FixedHeader: &FixedHeader{MessageType: Disconnect}}
	},
}

//DisconnectPacket is an internal representation of the fields of the
//Disconnect MQTT packet
type DisconnectPacket struct {
	*FixedHeader
}

// NewDisconnectPacket return the disconnect packet
func NewDisconnectPacket() *DisconnectPacket {
	return _disconnectPacketPool.Get().(*DisconnectPacket)
}

//Reset will initialize the fields in control packet
func (d *DisconnectPacket) Reset() {
	d.FixedHeader.Dup = false
	d.FixedHeader.QoS = byte(0)
	d.FixedHeader.RemainingLength = 0
	d.FixedHeader.Retain = false
}

//Close reset the packet field put the control packet back to pool
func (d *DisconnectPacket) Close() {
	d.Reset()
	_disconnectPacketPool.Put(d)
}

//SetFixedHeader will set fh for our header
func (d *DisconnectPacket) SetFixedHeader(fh *FixedHeader) {
	d.FixedHeader = fh
}

//Type return the packet type
func (d *DisconnectPacket) Type() byte {
	return d.FixedHeader.MessageType
}

func (d *DisconnectPacket) String() string {
	return fmt.Sprintf("%s", d.FixedHeader)
}

func (d *DisconnectPacket) Write(w io.Writer) error {
	b := Getbuf()
	d.FixedHeader.pack(b[:5])
	_, err := w.Write(b[3:5])
	Putbuf(b)
	return err
}

//Unpack decodes the details of a ControlPacket after the fixed
//header has been read
func (d *DisconnectPacket) Unpack([]byte) error {
	return nil
}

//Details returns a Details struct containing the QoS and
//MessageID of this ControlPacket
func (d *DisconnectPacket) Details() Details {
	return Details{QoS: 0, MessageID: 0}
}
