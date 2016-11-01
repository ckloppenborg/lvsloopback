/* 
Run webserver on port 8888, if uri specifies an active IP address of the realserver it reports back Online
Carl Kloppenborg
Version 0.1
*/

package main

import (
    "fmt"
    "net/http"
    "net"
    "bytes"
)

func handler(w http.ResponseWriter, r *http.Request) {
    uri := r.URL.Path[1:]
    testip := net.ParseIP(uri)
    // verify uri is an IP address, return error if not
    if testip.To4() == nil {
        fmt.Fprintf(w, "Not a valid address")
    } else {
        // Get list of all active IP addresses
        ipaddrlist, err := net.InterfaceAddrs()
        if err == nil {
            i := 0
            // Go through list of IP address and check to see if it matches the IP from the uri
            for _, value := range ipaddrlist {
                // Parse output of net.InterfaceAddrs() to remove mask 
                if ipnet, ok := value.(*net.IPNet); ok {
                    if bytes.Compare(ipnet.IP, testip) == 0 {
                        // We found a match, incriment i so we do not send 'Offline'
                        fmt.Fprintf(w, "Online")
                        i = i + 1
                    }
                }
            }
            // If there was no match i should still be 0
            if i == 0 {
                fmt.Fprint(w, "Offline")
            }
        // Somthing went wrong getting list of system IP's
        } else {
            fmt.Fprintf(w, "Something went wrong")
        }
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8888", nil)
}


