## zone2global
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/scottdware/zone2global/master/LICENSE)

This tool, when run against a Juniper SRX firewall, will convert all of the zone-based address books to a single global one.
> You MUST be running JUNOS 11.2 or above

### Example Usage

* The script takes three (3) arguments: `-srx (firewall device)`, `-u (username)`, `-p (password)`
* You can optionally choose to automatically apply the new configuration by specifying the `-commit` option.

`zone2global.exe -srx somefirewall -u admin -p password -commit`

If you omit the `-commit` option, then the configuration will be displayed on the console for you to copy/paste.

[license]: https://github.com/scottdware/zone2global/blob/master/LICENSE