## hioctl storage copy-url

copy a url to the storage pool

```
hioctl storage copy-url [flags]
```

### Options

```
      --filename string   filename for the disk
  -h, --help              help for copy-url
  -i, --id string         Storage Pool Id
  -n, --name string       Storage Pool Name
      --progress-bar      show a progress bar with --wait
      --raw-progress      print progress as a number with --wait
      --url string        url to download
      --wait              wait for task to complete
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

* [hioctl storage](hioctl_storage.md)	 - storage operations

###### Auto generated by spf13/cobra on 9-Jul-2025
