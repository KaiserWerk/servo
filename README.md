# servo

A pure Go library for handling system services/daemons. This currently works for
Windows, Linux (systemd) and Darwin (macOS).

## API

``servo.StartService(name string) error`` starts the service with the supplied name.
This call is platform independent.

``servo.StopService(name string) error`` stops the service with the supplied name.
This call is platform independent.

``servo.InstallService(exePath, name string) error`` installs the exePath as a service under the 
supplied name. This call is platform independent.

``servo.UninstallService(name string) error`` uninstalls the service with the
supplied name. This call is platform independent.

For more control, there are distinct calls for installing a service:

Windows:
``servo.InstallWindowsService(cfg WindowsServiceConfig) error``

Linux:
``servo.InstallLinuxService(cfg LinuxServiceConfig, servo.TypeSystemD) error``

MacOS:
``servo.InstallDarwinService(cfg DarwinServiceConfig) error``