package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"strings"

	"github.com/apparentlymart/go-cidr/cidr"
)

const (
	maxMask = 16
)

func main() {
	ip, ipnet, err := parseCIDR()
	if err != nil {
		printErrlnExit("ERROR: Failed parsing CIDR: " + err.Error())
	}

	maskSize, _ := ipnet.Mask.Size()
	if maskSize >= maxMask {
		fmt.Println(ipnet)
		os.Exit(0)
	}

	maskDiff := maxMask - maskSize
	newSubnetCount := math.Pow(2, float64(maskDiff))

	printNextIPNet(ip.String()+"/16", newSubnetCount)
}

func printNextIPNet(ipnet string, count float64) {
	fmt.Println(ipnet)
	count--

	if count == 0 {
		return
	}

	ip := strings.Split(ipnet, "/")[0]

	_, tmp, _ := net.ParseCIDR(ip + "/16")
	nextIPNet, _ := cidr.NextSubnet(tmp, 16)

	printNextIPNet(nextIPNet.String(), count)
}

func parseCIDR() (ip net.IP, ipnet *net.IPNet, err error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		reader := bufio.NewReader(os.Stdin)
		stdin, err := reader.ReadString('\n')
		if err != nil {
			return ip, ipnet, err
		}
		stdin = strings.TrimSpace(strings.Replace(stdin, "\n", "", -1))

		if len(stdin) != 0 {
			return net.ParseCIDR(stdin)
		}
	}

	if len(os.Args) > 1 {
		return net.ParseCIDR(os.Args[1])
	}

	err = errors.New("No stdin & No arg provided")
	return
}

func printErrlnExit(in string) {
	printErrln(in)
	os.Exit(1)
}

func printErrln(in string) {
	io.Copy(os.Stderr, bytes.NewBufferString(in+"\n"))
}
