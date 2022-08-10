package servo

import (
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func (wsc *WindowsServiceConfig) toMgrConfig() mgr.Config {
	return mgr.Config{
		// TODO
	}
}

func startService(name string, args ...string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}

	s, err := m.OpenService(name)
	if err != nil {
		return err
	}

	if err = s.Start(args...); err != nil {
		return err
	}

	return m.Disconnect()
}

func stopService(name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}

	s, err := m.OpenService(name)
	if err != nil {
		return err
	}

	if _, err = s.Control(svc.Stop); err != nil {
		return err
	}

	return m.Disconnect()
}

func installService(executable string, name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}

	_, err = m.CreateService(name, executable, mgr.Config{})
	if err != nil {
		return err
	}

	return m.Disconnect()
}

func uninstallService(name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}

	s, err := m.OpenService(name)
	if err != nil {
		return err
	}

	if err = s.Delete(); err != nil {
		return err
	}

	return m.Disconnect()
}

func installWindowsService(cfg WindowsServiceConfig) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}

	// TODO

	return m.Disconnect()
}
