package servo

func StartService(name string) error {
	return startService(name)
}

func StopService(name string) error {
	return stopService(name)
}

func InstallService(executable, name string) error {
	return installService(executable, name)
}

func UninstallService(name string) error {
	return uninstallService(name)
}

func InstallWindowsService(cfg WindowsServiceConfig) error {
	return installWindowsService(cfg)
}

func InstallLinuxService(cfg WindowsServiceConfig, serviceType LinuxServiceType) error {
	return installLinuxService(cfg, serviceType)
}

func InstallDarwinService(cfg DarwinServiceConfig) error {
	return installDarwinService(cfg)
}
