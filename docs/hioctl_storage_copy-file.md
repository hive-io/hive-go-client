## hioctl storage copy-file

copy a storage pool file

### Synopsis

copy a storage pool file

```
hioctl storage copy-file [flags]
```

### Options

```
      --destFilePath string    path to file in the destination storage pool
      --destStorageId string   Destination storage pool id
  -h, --help                   help for copy-file
      --progress-bar           show a progress bar with --wait
      --raw-progress           print progress as a number with --wait
      --srcFilePath string     path to file in the source storage pool
      --srcStorageId string    Source storage pool id
      --wait                   wait for task to complete
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

* [hioctl storage](hioctl_storage.md)	 - storage operations

###### Auto generated by spf13/cobra on 20-Jan-2020