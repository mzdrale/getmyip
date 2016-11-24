# Simple tools to find your IP address

## Installation

Run `go get github.com/mzdrale/getmyip` from command line.

## Example

```
package main

import (
    "fmt"
    "github.com/mzdrale/getmyip"
)

func main() {
    // Get client remote IP address
    // myip := getmyip.ViaHTTP("http://somehost.com/ip/") // It's up to you to set some web page which returns client IP
    myip := getmyip.ViaOpenDNS()
    fmt.Println(myip)
}
```
