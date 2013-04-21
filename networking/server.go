package networking

import (
	"log"
	"net"
	"sync"
	"time"
)

// NewServer allocates a new *Server and starts a goroutine to re-send dropped
// "ensured" packets.
func NewServer(listener net.PacketConn) *Server {
	s := &Server{
		ln:      listener,
		client:  make(map[string]*ServerClient),
		ensured: make(map[uint64]*ensuredPacket),
	}
	go s.ensure()

	return s
}

// TODO: Documentation
type Server struct {
	ln     net.PacketConn
	lock   sync.Mutex
	client map[string]*ServerClient

	ensuredLock sync.Mutex // This uses a different lock to avoid contention
	ensured     map[uint64]*ensuredPacket
	lastEnsured uint64
}

func (s *Server) ensure() {
	for {
		time.Sleep(EnsureInterval)
		s.ensuredLock.Lock()
		for i, e := range s.ensured {
			if time.Since(e.t) < EnsureInterval {
				continue
			}

			if e.r <= 0 {
				log.Println("Dropping ensured packet", i, "to", e.a, "after", MaxRetries, "retries")
				delete(s.ensured, i)
				continue
			}

			e.t = time.Now()
			e.r--
			if _, err := s.ln.WriteTo(e.b, e.a); err != nil {
				log.Println("Error in ensured packet", i, "to", e.a, "-", err)
				delete(s.ensured, i)
				continue
			}
		}
		s.ensuredLock.Unlock()
	}

	panic("unreachable")
}

// TODO: Documentation
func (s *Server) Send(addr net.Addr, b []byte, ensureArrivalID uint64) {
	if ensureArrivalID != 0 {
		s.ensuredLock.Lock()
		s.ensured[ensureArrivalID] = &ensuredPacket{addr, b, time.Now(), MaxRetries}
		s.ensuredLock.Unlock()
	}
	_, err := s.ln.WriteTo(b, addr)
	if err != nil {
		log.Println("Error sending to", addr, "-", err)
	}
}

type ensuredPacket struct {
	a net.Addr
	b []byte
	t time.Time
	r int
}

func (s *Server) ensureID() *uint64 {
	id := new(uint64)

	s.ensuredLock.Lock()
	s.lastEnsured++
	*id = s.lastEnsured
	s.ensuredLock.Unlock()

	return id
}

func (s *Server) arrivalNotice(addr net.Addr, id uint64) {
	s.ensuredLock.Lock()
	if e := s.ensured[id]; e != nil {
		if e.a == addr {
			delete(s.ensured, id)
		}
	}
	s.ensuredLock.Unlock()
}

// TODO: Documentation
func (s *Server) Prune(maxIdle time.Duration) {
	for {
		time.Sleep(maxIdle)
		s.lock.Lock()
		for addr, c := range s.client {
			if time.Since(c.lastSeen) > maxIdle {
				delete(s.client, addr)
				close(c.in)
				log.Println("Dropped client", addr, "due to inactivity.")
			}
		}
		s.lock.Unlock()
	}

	panic("unreachable")
}

// Recieve accepts and dispatches packets sent to the server's listener. If this
// method returns, the error will always be non-nil. If the error is a
// *net.OpError that is marked as Temporary, this method will continue execution
// without returning. Packets with more than 8192 bytes of data are truncated.
func (s *Server) Recieve() error {
	for {
		var buf [MaxPacketSize]byte
		n, addr, err := s.ln.ReadFrom(buf[:])
		if err != nil {
			//log.Println(addr, err)

			if e, ok := err.(*net.OpError); ok && e.Temporary() {
				continue
			}

			return err
		}
		b := buf[:n]

		s.lock.Lock()
		if c, ok := s.client[addr.String()]; ok {
			c.Accept(b)
		} else {
			c = NewServerClient(addr, s)
			s.client[addr.String()] = c
			c.Accept(b)
		}
		s.lock.Unlock()
	}

	panic("unreachable")
}
