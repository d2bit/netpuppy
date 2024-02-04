package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	// NetPuppy modules:
	"netpuppy/utils"
)

//func worker(done chan bool) {
//	fmt.Print("working...")
//	time.Sleep(time.Second * 5)
//	fmt.Println("done")
//
//	done <- true
//	done <- true
//	done <- true
//	fmt.Println("after done")
//}

//func main() {
//	done := make(chan bool, 1)
//	done <- true
//	go worker(done)
//	//	<-done
//	fmt.Printf("tiddies\n")
//	time.Sleep(time.Second * 3)
//	fmt.Printf("tiddies 2\n")
//}

func main() {
	// Set flag values based on input:
	listenFlag := flag.Bool("l", false, "put NetPuppy in listen mode")
	hostFlag := flag.String("H", "0.0.0.0", "target host IP address to connect to")
	turdnuggies := flag.Int("p", 40404, "target port") // portFlag @Trauma_x_Sella

	// Parse command line arguments:
	//                                            error?
	flag.Parse()

	// Print banner:
	fmt.Printf("%s", utils.Banner())

	// Depending on input, create this peer's type:
	type peer struct {
		connection_type string
		rPort           int
		lPort           string
		address         string
		connection      net.Conn
	}

	// Initiate peer struct:
	thisPeer := peer{rPort: *turdnuggies, address: *hostFlag}

	// If -l was given, create an 'offense' peer:
	if *listenFlag {
		thisPeer.connection_type = "offense"
		thisPeer.address = "0.0.0.0"
	} else {
		thisPeer.connection_type = "connect_back"
	}

	// Now that we have our peer: try to make connection
	var asyncio_rocks net.Conn // connection @0xtib3rius
	var err error

	if thisPeer.connection_type == "offense" {
		listener, err1 := net.Listen("tcp", fmt.Sprintf(":%v", thisPeer.rPort))
		if err1 != nil {
			fmt.Printf("Error when creating listener: %v\n", err1)
			os.Stderr.WriteString(err1.Error())
			os.Exit(1)
		}

		asyncio_rocks, err = listener.Accept()
		if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
			//  log.Fatal(err1.Error()
		}
	} else {
		remoteHost := fmt.Sprintf("%v:%v", thisPeer.address, thisPeer.rPort)
		asyncio_rocks, err = net.Dial("tcp", remoteHost)

		// If there is an err, try the host address as ipv6 (need to add [] around string):
		if err != nil {
			remoteHost := fmt.Sprintf("[%v]:%v", thisPeer.address, thisPeer.rPort)
			asyncio_rocks, err = net.Dial("tcp", remoteHost)

			if err != nil {
				os.Stderr.WriteString(err.Error())
				os.Exit(1)
			}
		}
	}

	// Attach connection to peer struct:
	thisPeer.connection = asyncio_rocks
	localPortArr := strings.Split(thisPeer.connection.LocalAddr().String(), ":")
	localPort := localPortArr[len(localPortArr)-1]
	thisPeer.lPort = localPort

	// Update user:
	var updateUserBanner string = utils.UserSelectionBanner(thisPeer.connection_type, thisPeer.address, thisPeer.rPort, thisPeer.lPort)
	fmt.Println(updateUserBanner)

	/*
		func readstream()
			for())))))
				data = connection.readstram

	*/

	/*
		if -l is on,
			net.Listen('tcp', PORT)
			set connection address for socket to any
		if not
			connection address = host flag


		struct/ objsect thing (this peer)
			- connect back (executed on the target)
				- start the subprocess
			- offense (exe on hacker machine)
				- keeep taking user input


			- method:
				func make connection(){
					if this.type = offense:
						connection = net.Listener
						(needs Accept() to actually become a connection)
					else:
						connection = net.Dial
				}
	*/

	// Try to create connection:
	return
}
