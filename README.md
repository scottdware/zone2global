## zone2global
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/scottdware/zone2global/master/LICENSE)

This tool, when run against a Juniper SRX firewall, will convert all of the zone-based address books to a single global one.
> You MUST be running JUNOS 11.2 or above

### Download

Head over to the [releases][releases] page to download the latest version.

### Example Usage

* The script takes three (3) arguments:
	* `-srx`
		* One or more SRX firewalls. If multiple, enclose in quotes like so: `-srx "firewall-1, firewall-2"`
	* `-u` Username
	* `-p` Password
* You can optionally choose to automatically apply the new configuration by specifying the `-commit` option.

`zone2global -srx somefirewall -u admin -p password -commit`

If you omit the `-commit` option, then the configuration will be written to a file for each SRX you specified. For example, running the above command without the `-commit` option will create a file called `somefirewall_globaladdrbook.txt` in the same location where the script was run.

[license]: https://github.com/scottdware/zone2global/blob/master/LICENSE
[releases]: https://github.com/scottdware/zone2global/releases