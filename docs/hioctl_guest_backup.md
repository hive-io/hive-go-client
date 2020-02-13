## hioctl guest backup

start guest backup

### Synopsis

start guest backup

```
hioctl guest backup [Name] [flags]
```

### Options

```
  -h, --help           help for backup
      --progress-bar   show a progress bar with --wait
      --raw-progress   print progress as a number with --wait
      --wait           wait for task to complete
```

### Options inherited from parent commands

```
      --config string     config file
      --format string     format (json/yaml) (default "json")
      --host string       Hostname or ip address
  -k, --insecure          ignore certificate errors
  -p, --password string   Admin user password
      --port uint         port (default 8443)
  -r, --realm string      Admin user realm (default "local")
  -u, --user string       Admin username (default "admin")
```

### SEE ALSO

* [hioctl guest](hioctl_guest.md)	 - guest operations

###### Auto generated by spf13/cobra on 20-Jan-2020