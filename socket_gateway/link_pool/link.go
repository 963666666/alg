package link_pool

import "sync"

type Link struct {
	lock sync.RWMutex

	socket *socket
}