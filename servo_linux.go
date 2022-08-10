package servo

type LinuxServiceType uint8

const (
	TypeSystemD LinuxServiceType = 1 << iota
)

type LinuxServiceConfig struct{}

func startService(name string) error {

}

func stopService(name string) error {

}

func installService(executable string, name string) error {

}

func uninstallService(name string) error {

}

func installLinuxService(cfg WindowsServiceConfig, serviceType LinuxServiceType) error {

}
