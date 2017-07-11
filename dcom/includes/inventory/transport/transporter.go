package transport

import (
	"inventory"
	"net"
)

type InventoryTransporter interface {
	inventory.Service
	Serve(net.Listener) error
}
