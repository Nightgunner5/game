package networking

import (
	"github.com/Nightgunner5/game/networking/packet"
	"net"
	"testing"
)

func TestClientServer(t *testing.T) {
	ln, err := net.ListenPacket("udp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	s := NewServer(ln)

	go s.Recieve()

	conn, err := net.Dial("udp", ln.LocalAddr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	c := NewClient(conn)

	c.Write(&packet.Packet{
		TestingPacket: &packet.Packet_TestingPacket{
			Type: packet.Packet_TestingPacket_Request.Enum(),
		},
	}, false)

	p, err := c.Read()
	if err != nil {
		t.Error(err)
	}
	if p.GetTestingPacket().GetType() != packet.Packet_TestingPacket_Response {
		t.Error("Incorrect response:", p)
	}
}
