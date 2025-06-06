// SPDX-FileCopyrightText: 2023 Free Mobile
// SPDX-License-Identifier: AGPL-3.0-only

package decoder

import (
	"net/netip"
	"path/filepath"
	"testing"

	"akvorado/common/helpers"
	"akvorado/common/schema"
)

func TestDecodeMPLSAndIPv4(t *testing.T) {
	sch := schema.NewMock(t).EnableAllColumns()
	pcap := helpers.ReadPcapL2(t, filepath.Join("testdata", "mpls-ipv4.pcap"))
	bf := &schema.FlowMessage{}
	l := ParseEthernet(sch, bf, pcap)
	if l != 40 {
		t.Errorf("ParseEthernet() returned %d, expected 40", l)
	}
	expected := schema.FlowMessage{
		SrcAddr: netip.MustParseAddr("::ffff:10.31.0.1"),
		DstAddr: netip.MustParseAddr("::ffff:10.34.0.1"),
		ProtobufDebug: map[schema.ColumnKey]interface{}{
			schema.ColumnEType:        helpers.ETypeIPv4,
			schema.ColumnProto:        6,
			schema.ColumnSrcPort:      11001,
			schema.ColumnDstPort:      23,
			schema.ColumnTCPFlags:     16,
			schema.ColumnMPLSLabels:   []uint64{18, 16},
			schema.ColumnIPTTL:        255,
			schema.ColumnIPTos:        0xb0,
			schema.ColumnIPFragmentID: 8,
			schema.ColumnSrcMAC:       0x003096052838,
			schema.ColumnDstMAC:       0x003096e6fc39,
		},
	}
	if diff := helpers.Diff(bf, expected); diff != "" {
		t.Fatalf("ParseEthernet() (-got, +want):\n%s", diff)
	}
}

func TestDecodeVLANAndIPv6(t *testing.T) {
	sch := schema.NewMock(t).EnableAllColumns()
	pcap := helpers.ReadPcapL2(t, filepath.Join("testdata", "vlan-ipv6.pcap"))
	bf := &schema.FlowMessage{}
	l := ParseEthernet(sch, bf, pcap)
	if l != 179 {
		t.Errorf("ParseEthernet() returned %d, expected 179", l)
	}
	expected := schema.FlowMessage{
		SrcVlan: 100,
		SrcAddr: netip.MustParseAddr("2402:f000:1:8e01::5555"),
		DstAddr: netip.MustParseAddr("2607:fcd0:100:2300::b108:2a6b"),
		ProtobufDebug: map[schema.ColumnKey]interface{}{
			schema.ColumnEType:  helpers.ETypeIPv6,
			schema.ColumnProto:  4,
			schema.ColumnIPTTL:  246,
			schema.ColumnSrcMAC: 0x00121ef2613d,
			schema.ColumnDstMAC: 0xc500000082c4,
		},
	}
	if diff := helpers.Diff(bf, expected); diff != "" {
		t.Fatalf("ParseEthernet() (-got, +want):\n%s", diff)
	}
}

// ADD TESTING FOR HOSTNAME FUNCTION TO WELL KNOWN DNS ADDRESS
func TestRetrieveHost(t *testing.T) {
	// cases here
	var tests = []struct {
		name  string
		input netip.Addr
		want  string
	}{
		// the table itself
		{"8.8.4.4 should be dns.google.", netip.MustParseAddr("8.8.4.4"), "dns.google."},                           // Ipv4 Google public DNS
		{"2001:4860:4860::8888 should be dns.google.", netip.MustParseAddr("2001:4860:4860::8888"), "dns.google."}, // Ipv6 Google public DNS.
		{"192.168.40.56 should be void", netip.MustParseAddr("192.168.40.56"), ""},                                 // Try to use an unused ip w/ no host resolution.
	}
	// execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Hostname(tt.input)
			// fmt.Println("ans", ans, "tt.want", tt.want)
			if ans != tt.want {
				t.Errorf("Hostname() returned %s, expected ", tt.want)
			}
		})
	}
}