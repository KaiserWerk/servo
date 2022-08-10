package servo

type LinuxServiceType uint8

const (
	TypeSystemD LinuxServiceType = 1 << iota
)

type (
	WindowsServiceConfig struct{}
	LinuxServiceConfig   struct{}
	DarwinServiceConfig  struct{}
)
