package server

import (
	"errors"
	"io"
	"net"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/moltenwolfcub/multiplayerfps/common"
)

type Server struct {
	listenAddr string
	listener   net.Listener
	quitCh     chan struct{}
	inMsgCh    chan common.RecievedPacket
	peers      map[net.Addr]common.Connection

	lightColor mgl32.Vec3
}

func NewServer(listenAddr string) *Server {

	return &Server{
		listenAddr: listenAddr,
		quitCh:     make(chan struct{}),
		inMsgCh:    make(chan common.RecievedPacket, 10),
		peers:      make(map[net.Addr]common.Connection),
	}
}

/*
Sets up the connection the network and starts running all
the loops to handle the connection
*/
func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		common.ErrorLogger.Fatal(err)
	}
	defer listener.Close()
	s.listener = listener

	addr, ok := s.listener.Addr().(*net.TCPAddr)
	if !ok {
		common.ErrorLogger.Fatal("couldn't convert listener's address to a TCP address")
	}
	common.InfoLogger.Printf("Local server hosted on port %d\n", addr.Port)

	s.initialise()

	go s.mainLoop()
	go s.packetLoop()
	go s.acceptLoop()

	<-s.quitCh
	close(s.inMsgCh)
}

/*
A loop that checks the net.listener for new connections,
adds them to the server and starts a new readloop for them.
*/
func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			common.WarningLogger.Println("accept error: ", err)
			continue
		}

		common.InfoLogger.Println("New connection to the server: ", conn.RemoteAddr())
		s.peers[conn.RemoteAddr()] = common.NewConnection(conn)

		go s.readLoop(conn)
	}
}

/*
A loop for each connection to manage serverbound traffic
and copy recieved packets into the server inMsgCh for future
processing.

Also manages disconnection of the clients.
*/
func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	for {
		rawPacket, err := s.peers[conn.RemoteAddr()].Recieve()

		if errors.Is(err, io.EOF) {
			common.InfoLogger.Println("Lost connection to peer: ", conn.RemoteAddr())
			delete(s.peers, conn.RemoteAddr())
			return
		}

		if err != nil {
			common.WarningLogger.Println("read error: ", err.Error())
			continue
		}

		s.inMsgCh <- common.RecievedPacket{
			Packet: rawPacket,
			Sender: conn.RemoteAddr(),
		}
	}
}

/*
Runs each new packet on the inMsgCh through
the handlePacket() function
*/
func (s *Server) packetLoop() {
	for rawPacket := range s.inMsgCh {
		s.handlePacket(rawPacket)
	}
}

// ONLY EDIT BELOW THIS LINE! The above code handles the server setup and network connections

func (s *Server) initialise() {
	s.lightColor = mgl32.Vec3{1, 0, 1}
}

/*
Main loop that'll handle the serverside logic and state.
*/
func (s *Server) mainLoop() {
	for {
	}
}

/*
Will figure out what kind of packet has been recieved
and correctly handle how it should behave.
*/
func (s *Server) handlePacket(recieved common.RecievedPacket) {
	switch packet := recieved.Packet.(type) {
	case common.ServerBoundLightingRequest:
		err := s.peers[recieved.Sender].Send(common.ClientBoundLightingUpdate{Color: s.lightColor})
		if err != nil {
			common.WarningLogger.Println("unable to send lighting update to client", err)
		}
	default:
		common.ErrorLogger.Fatalf("unknown packet: %s", packet)
	}
}
