## hioctl export

export cluster configuration

```
hioctl export [file] [flags]
```

### Options

```
  -h, --help              help for export
      --include strings   Data to include in the export file (default [broker,clusters,guests,hosts,pools,profiles,realms,storagePools,templates,users])
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

###### Auto generated by spf13/cobra on 9-Jul-2025
