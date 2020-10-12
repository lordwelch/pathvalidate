# pathvalidate
[![PkgGoDev](https://pkg.go.dev/badge/github.com/lordwelch/pathvalidate)](https://pkg.go.dev/github.com/lordwelch/pathvalidate)
[![Go Report Card](https://goreportcard.com/badge/github.com/lordwelch/pathvalidate)](https://goreportcard.com/report/github.com/lordwelch/pathvalidate)

Path santization based on pathvalidate from Python https://pypi.org/project/pathvalidate/

import path: `github.com/lordwelch/pathvalidate`

Example:
```Go
# Validate Path
err := pathvalidate.ValidateFilepath("Simple/Name", '_')
sanitized, err := pathvalidate.SanitizeFilepath("Simple/Name", '_')

# Validate Filename
err := pathvalidate.ValidateFilename("Simple/Name")
sanitized, err := pathvalidate.SanitizeFilename("Simple/Name")
```
Output:
```
# Validate Path
err: <nil>
sanitized: Simple/Name err: <nil>

# Validate Filename
err: pathvalidate: invalid character: '/' (0x2f)
sanitized: Simple_Name err: <nil>
```
## defaults
### Windows
Invalid Path: Unicode categories: Cc, Cf, Z excluding space + `:*?"<>|`

Invalid Filename: Invalid Path + `/` + `\`

Max Path Length: 260
#### Reserved words

NTFS Reserved Names: $MFT, $MFTMIRR, $LOGFILE $VOLUME, $ATTRDEF, $BITMAP, $BOOT, $BADCLUS, $SECURE, $UPCASE, $EXTEND, $QUOTA, $OBJID, $REPARSE

Windows Reserved Names: CON, PRN, AUX, CLOCK$, NUL, COM1, COM2, COM3, COM4, COM5, COM6, COM7, COM8, COM9, COM10, LPT1, LPT2, LPT3, LPT4, LPT5, LPT6, LPT7, LPT8, LPT9, LPT10

### Linux
Invalid Path: Unicode categories: Cc, Cf, Z excluding space +

Invalid Filename: Invalid Path + `/`

Max Path Length: 4096
#### Reserved words

None

### Darwin
Invalid Path: Unicode categories: Cc, Cf, Z excluding space +

Invalid Filename: Invalid Path + `/`

Max Path Length: 4096
#### Reserved words

`:`
