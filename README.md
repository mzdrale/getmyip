# Simple tools to find your IP address

## Installation

Run `go get github.com/mzdrale/getmyip` from command line.

## Example

```go
package main

import (
    "fmt"
    "github.com/mzdrale/getmyip"
)

func main() {
    // Get client remote IP address
    // myip := getmyip.ViaHTTP("http://yourwebserver.com/ip.php")
    // OR
    myip := getmyip.ViaOpenDNS()
    fmt.Println(myip)
}
```

If you are using `ViaHTTP` function, then you have to provide some web page URL which prints client IP address only. You can set `ip.php` for example in web root on your publicly available web server:

```php
<?php

print $_SERVER["REMOTE_ADDR"];

?>
```

You could also use [ifconfig.me service](http://ifconfig.me/ip), but from my experience it often responds very slow.

Instead of `ViaHTTP` function, you can use `ViaOpenDNS` function, which uses OpenDNS DNS servers.
