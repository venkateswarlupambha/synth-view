package capture

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
)

// PacketEvent represents a normalized packet extracted from PCAP
type PacketEvent struct {
	Timestamp time.Time
	SrcIP     string
	DstIP     string
	SrcPort   string
	DstPort   string
	Protocol  string
	Length    int
}

// ReadPCAP reads packets from a PCAP file using pcapgo (no libpcap needed)
func ReadPCAP(pcapPath string, output chan<- PacketEvent) {

	file, err := os.Open(pcapPath)
	if err != nil {
		log.Fatalf("Error opening PCAP file: %v", err)
	}
	defer file.Close()

	reader, err := pcapgo.NewReader(file)
	if err != nil {
		log.Fatalf("Error creating PCAP reader: %v", err)
	}

	packetSource := gopacket.NewPacketSource(reader, reader.LinkType())

	for packet := range packetSource.Packets() {

		event := PacketEvent{
			Timestamp: packet.Metadata().Timestamp,
			Length:    len(packet.Data()),
		}

		// Network layer
		netLayer := packet.NetworkLayer()
		if netLayer == nil {
			continue // drop non-IP packets
		}

		switch ip := netLayer.(type) {
		case *layers.IPv4:
			event.SrcIP = ip.SrcIP.String()
			event.DstIP = ip.DstIP.String()
			event.Protocol = ip.Protocol.String()
		case *layers.IPv6:
			event.SrcIP = ip.SrcIP.String()
			event.DstIP = ip.DstIP.String()
			event.Protocol = ip.NextHeader.String()
		default:
			continue 
			// not IP traffic
		}


		// Transport layer
		transport := packet.TransportLayer()
if transport == nil {
	continue
}

switch t := transport.(type) {
case *layers.TCP:
	event.SrcPort = t.SrcPort.String()
	event.DstPort = t.DstPort.String()
case *layers.UDP:
	event.SrcPort = t.SrcPort.String()
	event.DstPort = t.DstPort.String()
default:
	continue
}


		output <- event
	}

	close(output)
	fmt.Println("[✔] Finished reading PCAP file (pcapgo)")
}

