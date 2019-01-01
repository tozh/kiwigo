package client

import (
	"net"
	"fmt"
	"io"
	"bufio"
	"time"
	"sync"
)

type Client struct {
	conn           net.Conn
	passiveCloseCh chan struct{}
	closeCh        chan struct{}
	readBuf        *LargeBuffer
	writer         *bufio.Writer
	readCount      int
	timeOut        time.Duration
	wg             sync.WaitGroup
	closed         bool
}

func CreateClient(conn net.Conn) *Client {
	client := Client{
		conn:           conn,
		passiveCloseCh: make(chan struct{}, 1),
		closeCh:        make(chan struct{}, 1),
		readBuf:        &LargeBuffer{},
		writer:         bufio.NewWriterSize(conn, PROTO_IOBUF_LEN),
		readCount:      0,
		timeOut:        CONFIG_DEFAULE_TIMEOUT_MILLISECOND * time.Millisecond,
		wg:             sync.WaitGroup{},
		closed:         false,
	}
	return &client
}

func TcpAddress(ip string, port int) string {
	return fmt.Sprintf("%s:%d", ip, port)
}

func TcpClient(ip string, port int) *Client {
	conn, err := net.Dial("tcp", TcpAddress(ip, port))
	if err != nil {
		panic(err)
		conn.Close()
		return nil
	}
	return CreateClient(conn)
}

func (c *Client) ChangeTimeOut(dur time.Duration) {
	c.timeOut = dur
}

func UnixClient(address string) *Client {
	conn, err := net.Dial("unix", address)
	if err != nil {
		panic("err")
		conn.Close()
		return nil
	}
	return CreateClient(conn)
}

func (c *Client) readFromServer(readCh chan error) {
	// fmt.Println("readFromServer")
	reader := bufio.NewReaderSize(c.conn, PROTO_IOBUF_LEN)
	for {
		recieved, err := reader.ReadBytes(0)
		// fmt.Println(len(recieved), len(string(recieved)))
		if err != nil {
			// fmt.Println(err)
			if err == io.EOF {
				// fmt.Println("err is EOF")
				c.broadcastPassiveClose()
				return
			} else {
				readCh <- err
			}
		}
		if len(recieved) > 0 {
			//// fmt.Println("recieved----->", string(recieved))
			c.readBuf.Write(recieved)
		}
		if err == nil {
			break
		}
	}
	c.readCount++
	readCh <- nil
}

func (c *Client) writeToServer() {
	c.writer.WriteByte(0)
	err := c.writer.Flush()
	if err != nil {
		return
	}
}

func (c *Client) reset() {
	c.readBuf.Reset()
	c.writer.Reset(c.conn)
}

func (c *Client) broadcastPassiveClose() {
	// fmt.Println("broadcastPassiveClose")

	close(c.passiveCloseCh)
}

func (c *Client) passiveClose() {
	// fmt.Println("passiveClose")
	if !c.closed {
		c.writer = nil
		c.readBuf = nil
		c.conn.Close()
		c.conn = nil
		close(c.closeCh)
	}
	c.closed = true
}

func (c *Client) Close() {
	// fmt.Println("Close")
	if !c.closed {
		c.writer = nil
		c.readBuf = nil
		c.conn.Close()
		c.conn = nil
		close(c.closeCh)
	}
	c.closed = true

}

func (c *Client) closeListener() {
	select {
	case <-c.passiveCloseCh:
		go c.passiveClose()
	case <-c.closeCh:
		close(c.passiveCloseCh)
	}
}
