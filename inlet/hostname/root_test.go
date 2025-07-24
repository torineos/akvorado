package hostname

// import (
// 	"net"
// 	"path/filepath"
// 	"testing"

// 	"akvorado/common/helpers"
// 	"akvorado/common/reporter"
// 	"akvorado/common/schema"
// 	"akvorado/inlet/flow/decoder"
// 	"akvorado/inlet/flow/decoder/sflow"k
// )

// // func TestGetHostname(t *testing.T) {

// // }

// // func TestDefaultConfiguration(t *testing.T) {

// // }

// func TestIPV4Sflow(t *testing.T) {
// 	r := reporter.NewMock(t)
// 	sdecoder := sflow.New(r, decoder.Dependencies{Schema: schema.NewMock(t).EnableAllColumns()}, decoder.Option{})

// 	// Send data
// 	data := helpers.ReadPcapL4(t, filepath.Join("../../inlet/flow/decoder/sflow/testdata", "data-icmpv4.pcap"))
// 	decodedData := sdecoder.Decode(decoder.RawFlow{Payload: data, Source: net.ParseIP("127.0.0.1")}) // if root_test.go in inlet/flow/decoder/sflow worked this one sould too in theory

// 	// Fetch ip address from the damn flow
// 	var flow *schema.FlowMessage
// 	decoder.DecodeIP(flow.Bytes(flow.SrcAddr))

// }

// func TestIPV6Sflow(t *testing.T) {}

// func TestIPV4Netflow(t *testing.T) {}

// func TestIPV6Netflow(t *testing.T) {}

// // Emulate arrival of flows to test:
// // - Fetch srcaddr and dstaddr of flows
// // - Enabling of the component

// // TODO: Take inspiration from inlet/flow/decoder/net-sflow/root_test.go to get flows running from pcap files,
// // and get the hostnames on the fly if the option is enabled
// // sflow and netflow with ipv4 and ipv6
