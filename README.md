## zone2global
[![Travis-CI](https://travis-ci.org/scottdware/zone2global.svg?branch=master)](https://travis-ci.org/scottdware/zone2global)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/scottdware/zone2global/master/LICENSE)

This tool, when run against a Juniper SRX firewall, will convert all of the zone-based address books to a single global one.
> You MUST be running JUNOS 11.2 or above

This script is used to build the standalone binaries. If you wish to use the same functionality in your own scripts, you can 
view the function(s) code in the main [go-junos][go-junos] library [here][convert-code].

> The underlying communication is over Netconf/SSH. Please make sure your devices allow this by using one or both of the following commands:
```
set system services netconf ssh
set security zones security-zone <xxx> interfaces <xxx> host-inbound-traffic system-services netconf
```

### Download

Head over to the [releases][releases] page to download the latest version.

### Example Usage

```
zone2global - Convert an SRX from a zone-based address book to a global one.

Usage: zone2global [OPTIONS]
  -commit
        Choose to apply the configuration directly instead of creating a file.
  -p string
        Password
  -srx string
        SRX to run the conversion against. If specifying multiple, enclose in quotes, i.e. "srx240-1 srx1400-2"
  -u string
        Username
```

```
zone2global -srx somefirewall -u admin -p password -commit
```

If you omit the `-commit` option, then the configuration will be written to a file for each SRX you specified. For example, running the above command without the `-commit` option will create a file called `somefirewall_globaladdrbook.txt` in the same location where the script was run.

[license]: https://github.com/scottdware/zone2global/blob/master/LICENSE
[releases]: https://github.com/scottdware/zone2global/releases
[go-junos]: https://github.com/scottdware/go-junos
[convert-code]: https://github.com/scottdware/go-junos/blob/master/srx.go#L410