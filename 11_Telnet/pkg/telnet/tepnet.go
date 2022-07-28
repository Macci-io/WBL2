package telnet

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"wget/pkg/telnet/config"
	"wget/pkg/utils"
)

//Connect хранение информации о соединении
type Connect struct {
	conn   net.Conn
	in     io.Reader
	out    io.Writer
	statMT sync.Mutex
	status bool
}

//Run запуск основной программы
func (c *Connect) Run() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ok := make(chan error, 1)
	go c.Send(ok)
	go c.Recv(ok)

	select {
	case <-sig:
		break
	case err := <-ok:
		if err != nil {
			fmt.Println(err)
		}
	}
	c.Disconnect()
}

// Recv получение данных
func (c *Connect) Recv(okChan chan error) {
	var ok error
	_, ok = io.Copy(c.out, c.conn)
	okChan <- ok
}

// Send отправка данных
func (c *Connect) Send(okChan chan error) {
	var ok error
	_, ok = io.Copy(c.conn, c.in)
	okChan <- ok
}

// Disconnect отключиться от хоста
func (c *Connect) Disconnect() {
	c.statMT.Lock()
	if c.status {
		c.status = false
		c.statMT.Unlock()

		_ = c.conn.Close()
		utils.ServiceMessage("disconnected")
	} else {
		c.statMT.Unlock()
	}
}

func connect(conf *config.Config) (conn net.Conn, ok error) {
	ok = utils.TryItUntilTimeOut(conf.GetTimeOut(), time.Millisecond*50, func() error {
		conn, ok = net.DialTimeout("tcp", conf.GetConnectionInfo(), conf.GetTimeOut())
		return ok
	})
	if ok != nil {
		return nil, ok
	}
	return
}

//NewConnection создание нового подключения
func NewConnection(in io.Reader, out io.Writer) (conn *Connect, ok error) {
	var conf *config.Config
	conn = &Connect{nil, in, out, sync.Mutex{}, false}
	if conf, ok = config.NewConfig(); ok != nil {
		return nil, ok
	}
	if conn.conn, ok = connect(conf); ok != nil {
		return nil, ok
	}
	conn.status = true
	utils.ServiceMessage("connected")
	return conn, nil
}
