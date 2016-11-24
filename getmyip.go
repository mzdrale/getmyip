// Package getmyip contains functions for getting client IP address

package getmyip

import (
    "fmt"
    "github.com/bogdanovich/dns_resolver"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
    "strings"
)

var opendns_servers = []string {
    "resolver1.opendns.com",
    "resolver2.opendns.com",
    "resolver3.opendns.com",
    "resolver4.opendns.com",
}
var opendns_myhost = "myip.opendns.com"

// Validate IPv4 address
func ValidateIPv4(ip string) bool {
    ip = strings.TrimSpace(ip)
    re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
    if re.MatchString(ip) {
        return true
    }
    return false
}

// Get client remote IP using HTTP request to custom web page returning client IP only
func ViaHTTP(url string) string {
    myip := "unknown"
    res, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    lines := strings.Split(string(body), "\n")
    firstline := string(strings.TrimSpace(lines[0]))

    if ValidateIPv4(firstline) {
        myip = firstline
    }
    return string(myip)
}

// Get client remote IP using OpenDNS DNS lookup
func ViaOpenDNS() string {
    myip := "unknown"
    resolver := dns_resolver.New(opendns_servers)
    resolver.RetryTimes = 5
    ip, err := resolver.LookupHost(opendns_myhost)
    if err == nil {
        firstip := ip[0].String()
        if ValidateIPv4(firstip) {
            myip = firstip
        }
    }
    return string(myip)
}
