package networking

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/Nightgunner5/game/networking/packet"
	"log"
	"net"
	"sync"
	"time"
)

// TODO: Documentation
func NewClient(conn net.Conn) *Client {
	c := &Client{
		conn:    conn,
		ensured: make(map[uint64]*clientEnsure),
	}
	go c.ensure()

	return c
}

// TODO: Documentation
type Client struct {
	conn       net.Conn
	lock       sync.Mutex
	ensured    map[uint64]*clientEnsure
	lastEnsure uint64
}

type clientEnsure struct {
	b []byte
	t time.Time
	r int
}

func (c *Client) ensure() {
	for {
		time.Sleep(EnsureInterval)
		c.lock.Lock()
		for i, e := range c.ensured {
			if time.Since(e.t) < EnsureInterval {
				continue
			}

			if e.r <= 0 {
				log.Println("Dropping ensured packet", i, "after", MaxRetries, "retries")
				delete(c.ensured, i)
				continue
			}

			e.t = time.Now()
			e.r--
			if _, err := c.conn.Write(e.b); err != nil {
				log.Println("Error in ensured packet", i, "-", err)
				delete(c.ensured, i)
				continue
			}
		}
		c.lock.Unlock()
	}

	panic("unreachable")
}

// TODO: Documentation
func (c *Client) Read() (*packet.Packet, error) {
	var buf [MaxPacketSize]byte
	parsed := new(packet.Packet)

	for {
		n, err := c.conn.Read(buf[:])
		if err != nil {
			return nil, err
		}

		if err = proto.Unmarshal(buf[:n], parsed); err != nil {
			return nil, err
		}

		if an := parsed.GetArrivalNotice().GetPacketID(); an != 0 {
			c.lock.Lock()
			delete(c.ensured, an)
			c.lock.Unlock()
			continue
		}

		if ensure := parsed.GetEnsureArrival(); ensure != 0 {
			c.Write(&packet.Packet{
				ArrivalNotice: &packet.Packet_ArrivalNotice{
					PacketID: proto.Uint64(ensure),
				},
			}, false)
		}

		return parsed, nil
	}

	panic("unreachable")
}

// TODO: Documentation
func (c *Client) Write(p *packet.Packet, ensure bool) error {
	if ensure {
		c.lock.Lock()
		c.lastEnsure++
		p.EnsureArrival = proto.Uint64(c.lastEnsure)
		c.lock.Unlock()
	} else {
		p.EnsureArrival = nil
	}

	b, err := proto.Marshal(p)
	if err != nil {
		return err
	}

	if ensure {
		c.lock.Lock()
		c.ensured[p.GetEnsureArrival()] = &clientEnsure{b, time.Now(), MaxRetries}
		c.lock.Unlock()
	}

	_, err = c.conn.Write(b)
	return err
}
