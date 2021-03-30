## hioctl pool delete-assignment

delete the assignment for a standalone pool

### Synopsis

delete the assignment for a standalone pool

```
hioctl pool delete-assignment [flags]
```

### Options

```
  -h, --help          help for delete-assignment
  -i, --id string     pool Id
  -n, --name string   pool Name
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

* [hioctl pool](hioctl_pool.md)	 - pool operations

###### Auto generated by spf13/cobra on 10-May-2021