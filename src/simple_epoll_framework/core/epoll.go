package core

import (
	"fmt"
	"net"
	"simple_epoll_framework/conf"
	"syscall"
)

type Epoll struct {
	epollfd int
}

// 创建epoll
func NewEpoll() *Epoll {
	return &Epoll{}
}

// 开始工作
func (e *Epoll) Work() error {

	if len(conf.Gconf.Epoll.Ports) == 0 {
		return fmt.Errorf("no port need listen")
	}

	// todo 创建epoll


	// 创建sever socket并添加到epoll里
	for _, port := range conf.Gconf.Epoll.Ports {
		fd, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)
		if err != nil {
			return fmt.Errorf("create socket failed, err:%v", err)
		}

		if err = syscall.SetNonblock(fd, true); err != nil {
			_ = syscall.Close(fd)
			return fmt.Errorf("set socket nonblock failed, err:%v", err)
		}

		addr := syscall.SockaddrInet4{Port: port}
		copy(addr.Addr[:], net.ParseIP("0.0.0.0").To4())

		if err = syscall.Bind(fd, &addr); err != nil {
			_ = syscall.Close(fd)
			return fmt.Errorf("fd bind addr failed, port:%v, err:%v", port, err)
		}

		if err = syscall.Listen(fd, 10); err != nil {
			_ = syscall.Close(fd)
			return fmt.Errorf("sycall.Listen fd failed, err:%v", err)
		}

		// 添加到epoll


	}

}
