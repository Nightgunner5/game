package networking

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/Nightgunner5/game/networking/packet"
	"log"
	"net"
	"time"
)

// TODO: Documentation
func NewServerClient(addr net.Addr, s *Server) *ServerClient {
	sc := &ServerClient{
		addr:   addr,
		server: s,
		in:     make(chan *packet.Packet, 16),
	}
	go sc.dispatch()

	return sc
}

// TODO: Documentation
type ServerClient struct {
	addr     net.Addr
	server   *Server
	lastSeen time.Time
	in       chan *packet.Packet
}

// TODO: Documentation
func (sc *ServerClient) Send(p *packet.Packet, ensure bool) {
	if ensure {
		p.EnsureArrival = sc.server.ensureID()
	} else {
		p.EnsureArrival = nil
	}
	b, err := proto.Marshal(p)
	if err != nil {
		log.Println("Error encoding packet to", sc.addr, "-", err)
	} else {
		sc.server.Send(sc.addr, b, p.GetEnsureArrival())
	}
}

func (sc *ServerClient) dispatch() {
	for p := range sc.in {
		if p.GetEnsureArrival() != 0 {
			sc.Send(&packet.Packet{
				ArrivalNotice: &packet.Packet_ArrivalNotice{
					PacketID: proto.Uint64(p.GetEnsureArrival()),
				},
			}, false)
		} else if an := p.GetArrivalNotice().GetPacketID(); an != 0 {
			sc.server.arrivalNotice(sc.addr, an)
			return
		}

		log.Println(sc.addr, p)
	}
}

// TODO: Documentation
func (sc *ServerClient) Accept(b []byte) {
	sc.lastSeen = time.Now()

	parsed := new(packet.Packet)
	if err := proto.Unmarshal(b, parsed); err != nil {
		log.Println("Error decoding packet from", sc.addr, "-", err)
		return
	}

	select {
	case sc.in <- parsed:
	default:
		log.Println("Dropped packet from", sc.addr, "-", parsed)
	}
}
