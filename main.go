package main

import (
	"fmt"
	"log"
	"os/exec"
)

func nmapArpScan(ipRange string) {
	out, err := exec.Command("nmap", "-sP", "-PR", ipRange).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("nmap output %s\n", out)
}

func main() {
	nmapArpScan("192.168.217.0/24")
}
