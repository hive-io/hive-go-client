## hioctl host

host operations

```
hioctl host [flags]
```

### Options

```
  -h, --help   help for host
```

### Options inherited from parent commands

```
      --config string     config file
      --format string     format (json/yaml) (default "json")
      --host string       Hostname or ip address
  -k, --insecure          ignore certificate errors
  -p, --password string   Admin user password
      --port uint         port (default 8443)
      --profile string    Load a profile from the config file
  -r, --realm string      Admin user realm (default "local")
  -u, --user string       Admin username (default "admin")
```

### SEE ALSO

* [hioctl](hioctl.md)	 - hive fabric rest api client
* [hioctl host delete-software](hioctl_host_delete-software.md)	 - delete a software package
* [hioctl host disable-crs](hioctl_host_disable-crs.md)	 - disable crs on a host
* [hioctl host disable-gateway-mode](hioctl_host_disable-gateway-mode.md)	 - Convert the host from a gateway appliance to a regular fabric host
* [hioctl host enable-crs](hioctl_host_enable-crs.md)	 - enable crs on a host
* [hioctl host enable-gateway-mode](hioctl_host_enable-gateway-mode.md)	 - Convert the host into a gateway appliance
* [hioctl host get](hioctl_host_get.md)	 - get host details
* [hioctl host get-id](hioctl_host_get-id.md)	 - get hostid from hostname
* [hioctl host info](hioctl_host_info.md)	 - hostid and version
* [hioctl host iscsi-discover](hioctl_host_iscsi-discover.md)	 - discover iscsi targets
* [hioctl host iscsi-login](hioctl_host_iscsi-login.md)	 - login to an iscsi target
* [hioctl host iscsi-logout](hioctl_host_iscsi-logout.md)	 - logout from an iscsi target
* [hioctl host iscsi-sessions](hioctl_host_iscsi-sessions.md)	 - list iscsi sessions
* [hioctl host list](hioctl_host_list.md)	 - list hosts
* [hioctl host list-software](hioctl_host_list-software.md)	 - list available software packages on a host
* [hioctl host log-level](hioctl_host_log-level.md)	 - get or set host log level
* [hioctl host network](hioctl_host_network.md)	 - host network operations
* [hioctl host reboot](hioctl_host_reboot.md)	 - reboot a host
* [hioctl host restart-services](hioctl_host_restart-services.md)	 - restart hive servies
* [hioctl host shutdown](hioctl_host_shutdown.md)	 - shutdown a host
* [hioctl host state](hioctl_host_state.md)	 - get or set host state
* [hioctl host unjoin](hioctl_host_unjoin.md)	 - remove host from cluster
* [hioctl host update-gpu](hioctl_host_update-gpu.md)	 - Update GPU settings for a host
* [hioctl host update-sriov](hioctl_host_update-sriov.md)	 - Update settings for sriov devices on a host
* [hioctl host upload-software](hioctl_host_upload-software.md)	 - upload a software pkg file to a host

###### Auto generated by spf13/cobra on 9-Jul-2025
