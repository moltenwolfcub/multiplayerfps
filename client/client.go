package client

import (
	"fmt"
	"log"
	"net"

	"github.com/moltenwolfcub/multiplayerfps/common"
)

type Client struct {
	listenAddr string
	connection common.Connection
}

func NewClient(listenAddr string) *Client {
	return &Client{
		listenAddr: listenAddr,
	}
}

/*
Connects to the server and starts running the loops
which handle the rest of the logic
*/
func (c *Client) Start() error {
	conn, err := net.Dial("tcp", c.listenAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	c.connection = common.NewConnection(conn)

	go c.readLoop()
	return c.mainLoop()
}

/*
A loop to manage clientbound traffic and send recieved packets
to the handlepacket method for processing.
*/
func (c *Client) readLoop() error {
	for {
		rawPacket := c.connection.MustRecieve()
		err := c.handlePacket(rawPacket)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

// ONLY EDIT BELOW THIS LINE! The above code handles the client setup and manages the network connection

/*
Main loop that'll handle the clientside logic and state.
*/
func (c *Client) mainLoop() error {
	for {
	}
}

/*
Will figure out what kind of packet has been recieved
and correctly handle how it should behave.
*/
func (c *Client) handlePacket(rawPacket common.Packet) error {
	switch packet := rawPacket.(type) {
	default:
		return fmt.Errorf("unkown packet: %s", packet)
	}
	return nil
}
