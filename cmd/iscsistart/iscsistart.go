package main

import (
	"flag"
	"log"

	"github.com/u-root/iscsinl"
)

var (
	targetAddr      = flag.String("addr", "instance-1:3260", "target addr")
	volumeName      = flag.String("volume", "FOO", "volume name")
	monitorNetlink  = flag.Bool("monitorNetlink", false, "Set to true to monitor netlink socket until killed")
	teardownSession = flag.Int("teardownSid", -1, "Set to teardown a session")
)

func main() {
	flag.Parse()

	if *teardownSession != -1 {
		if err := iscsinl.TearDownIscsi((uint32)(*teardownSession), 0); err != nil {
			log.Fatal(err)
		}
		return
	}

	device, err := iscsinl.MountIscsi(*targetAddr, *volumeName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Mounted at dev %v", device)
}