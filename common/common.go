/*
 * Copyright (c) 2022 NetLOX Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"net"
)

// This file defines the go interface implementation needed to interact with loxinet go module

const (
	// CIStateMaster - HA Master state
	CIStateMaster = 1 + iota
	// CIStateBackup - HA Backup/Slave state
	CIStateBackup
	// CIStateConflict - HA Fault/Conflict State
	CIStateConflict
	// CIStateNotDefined - HA State not enabled or stopped
	CIStateNotDefined
)

const (
	// CIDefault - Default CI Instance name
	CIDefault = "default"
)

const (
	// HighLocalPref - High local preference for advertising BGP route(Default or Master)
	HighLocalPref = 101
	// LowLocalPref - Low local preference for advertising BGP route(Backup)
	LowLocalPref = 100
)

const (
	// CertPath - SSL certificates path
	CertPath = "/opt/loxilb/cert/"

	// CACertFileName - loxilb CA cert file
	CACertFileName = "rootCA.crt"

	// PrivateCertName - loxilb private certificate name
	PrivateCertName = "server.crt"

	// PrivateKeyName - loxilb private key name
	PrivateKeyName = "server.key"
)

const (
	// AuWorkqLen - Address worker channel depth
	AuWorkqLen = 1024
	// LuWorkQLen - Link worker channel depth
	LuWorkQLen = 1024
	// NuWorkQLen - Neigh worker channel depth
	NuWorkQLen = 1024
	// RuWorkQLen - Route worker channel depth
	RuWorkQLen = 40827
)

const (
	// PortReal - Base port type
	PortReal = 1 << iota
	// PortBondSif - Bond slave port type
	PortBondSif
	// PortBond - Bond port type
	PortBond
	// PortVlanSif - Vlan slave port type
	PortVlanSif
	// PortVlanBr - Vlan Br port type
	PortVlanBr
	// PortVxlanSif - Vxlan slave port type
	PortVxlanSif
	// PortVxlanBr - Vxlan br port type
	PortVxlanBr
	// PortWg - Wireguard port type
	PortWg
	// PortVti - Vti port type
	PortVti
	// PortIPTun - IPInIP port type
	PortIPTun
	// PortGre - GRE port type
	PortGre
)

// PortProp - Defines auxiliary port properties
type PortProp uint8

const (
	// PortPropUpp - User-plane processing enabled
	PortPropUpp PortProp = 1 << iota
	// PortPropSpan - SPAN is enabled
	PortPropSpan
	// PortPropPol - Policer is active
	PortPropPol
)

// DpStatusT - Generic status of operation
type DpStatusT uint8

// PortDump - Generic dump info of a port
type PortDump struct {
	// Name - name of the port
	Name string `json:"portName"`
	// PortNo - port number
	PortNo int `json:"portNo"`
	// Zone - security zone info
	Zone string `json:"zone"`
	// SInfo - software specific port information
	SInfo PortSwInfo `json:"portSoftwareInformation"`
	// HInfo - hardware specific port information
	HInfo PortHwInfo `json:"portHardwareInformation"`
	// Stats - port statistics related information
	Stats PortStatsInfo `json:"portStatisticInformation"`
	// L3 - layer3 info related to port
	L3 PortLayer3Info `json:"portL3Information"`
	// L2 - layer2 info related to port
	L2 PortLayer2Info `json:"portL2Information"`
	// Sync - sync state
	Sync DpStatusT `json:"DataplaneSync"`
}

// PortStatsInfo - stats information of port
type PortStatsInfo struct {
	// RxBytes - rx Byte count
	RxBytes uint64 `json:"rxBytes"`
	// TxBytes - tx Byte count
	TxBytes uint64 `json:"txBytes"`
	// RxPackets - tx Packets count
	RxPackets uint64 `json:"rxPackets"`
	// TxPackets - tx Packets count
	TxPackets uint64 `json:"txPackets"`
	// RxError - rx error count
	RxError uint64 `json:"rxErrors"`
	// TxError - tx error count
	TxError uint64 `json:"txErrors"`
}

// PortHwInfo - hw info of a port
type PortHwInfo struct {
	// MacAddr - mac address as byte array
	MacAddr [6]byte `json:"rawMacAddress"`
	// MacAddrStr - mac address in string format
	MacAddrStr string `json:"macAddress"`
	// Link - lowerlayer state
	Link bool `json:"link"`
	// State - administrative state
	State bool `json:"state"`
	// Mtu - maximum transfer unit
	Mtu int `json:"mtu"`
	// Master - master of this port if any
	Master string `json:"master"`
	// Real - underlying real dev info if any
	Real string `json:"real"`
	// TunID - tunnel info if any
	TunID uint32 `json:"tunnelId"`
}

// PortLayer3Info - layer3 info of a port
type PortLayer3Info struct {
	// Routed - routed mode or not
	Routed bool `json:"routed"`
	// Ipv4Addrs - ipv4 address set
	Ipv4Addrs []string `json:"IPv4Address"`
	// Ipv6Addrs - ipv6 address set
	Ipv6Addrs []string `json:"IPv6Address"`
}

// PortSwInfo - software specific info of a port
type PortSwInfo struct {
	// OsID - interface id of an OS
	OsID int `json:"osId"`
	// PortType - type of port
	PortType int `json:"portType"`
	// PortProp - port property
	PortProp PortProp `json:"portProp"`
	// PortActive - port enabled/disabled
	PortActive bool `json:"portActive"`
	// PortReal - pointer to real port if any
	PortReal *PortDump `json:"portReal"`
	// PortOvl - pointer to ovl port if any
	PortOvl *PortDump `json:"portOvl"`
	// BpfLoaded - eBPF loaded or not flag
	BpfLoaded bool `json:"bpfLoaded"`
}

// PortLayer2Info - layer2 info of a port
type PortLayer2Info struct {
	// IsPvid - this vid is Pvid or not
	IsPvid bool `json:"isPvid"`
	// Vid - vid related to prot
	Vid int `json:"vid"`
}

// PortMod - port modification info
type PortMod struct {
	// Dev - name of port
	Dev string
	// LinkIndex - OS allocated index
	LinkIndex int
	// Ptype - port type
	Ptype int
	// MacAddr - mac address
	MacAddr [6]byte
	// Link - lowerlayer state
	Link bool
	// State - administrative state
	State bool
	// Mtu - maximum transfer unit
	Mtu int
	// Master - master of this port if any
	Master string
	// Real - underlying real dev info if any
	Real string
	// TunID - tunnel info if any
	TunID int
	// TunSrc - tunnel source
	TunSrc net.IP
	// TunDst - tunnel dest
	TunDst net.IP
}

// VlanMod - Info about a vlan
type VlanMod struct {
	// Vid - vlan identifier
	Vid int `json:"vid"`
	// Dev - name of the related device
	Dev string `json:"dev"`
	// LinkIndex - OS allocated index
	LinkIndex int
	// MacAddr - mac address
	MacAddr [6]byte
	// Link - lowerlayer state
	Link bool
	// State - administrative state
	State bool
	// Mtu - maximum transfer unit
	Mtu int
	// TunID - tunnel info if any
	TunID uint32
}

// VlanPortMod - Info about a port attached to a vlan
type VlanPortMod struct {
	// Vid - vlan identifier
	Vid int `json:"vid"`
	// Dev - name of the related device
	Dev string `json:"dev"`
	// Tagged - tagged or not
	Tagged bool `json:"tagged"`
}

// vlanStat - statistics for vlan interface
type VlanStat struct {
	InBytes    uint64
	InPackets  uint64
	OutBytes   uint64
	OutPackets uint64
}

// VlanGet - Info for vlan interface to get
type VlanGet struct {
	// Vid - vlan identifier
	Vid int `json:"vid"`
	// Dev - name of port
	Dev string `json:"dev"`
	// Slaves - name of slave ports
	Member []VlanPortMod `json:"member"`
	// Stat Vlan traffic statistics
	Stat VlanStat `json:"vlanStatistic"`
}

const (
	// FdbPhy - fdb of a real dev
	FdbPhy = 0
	// FdbTun - fdb of a tun dev
	FdbTun = 1
	// FdbVlan - fdb of a vlan dev
	FdbVlan = 2
)

// FdbMod - Info about a forwarding data-base
type FdbMod struct {
	// MacAddr - mac address
	MacAddr [6]byte
	// BridgeID - bridge domain-id
	BridgeID int
	// Dev - name of the related device
	Dev string
	// Dst - ip addr related to fdb
	Dst net.IP
	// Type - One of FdbPhy/FdbTun/FdbVlan
	Type int
}

// IpAddrMod - Info about an ip address
type IpAddrMod struct {
	// Dev - name of the related device
	Dev string
	// IP - Actual IP address
	IP string
}

// NeighMod - Info about an neighbor
type NeighMod struct {
	// IP - The IP address
	IP net.IP
	// LinkIndex - OS allocated index
	LinkIndex int
	// State - active or inactive
	State int
	// HardwareAddr - resolved hardware address if any
	HardwareAddr net.HardwareAddr
}

// IpAddrGet - Info about an ip addresses
type IpAddrGet struct {
	// Dev - name of the related device
	Dev string
	// IP - Actual IP address
	IP []string
	// Sync - sync state
	Sync DpStatusT
}

// RouteGetEntryStatistic - Info about an route statistic
type RouteGetEntryStatistic struct {
	// Statistic of the ingress port bytes.
	Bytes int
	// Statistic of the egress port bytes.
	Packets int
}

// RouteGet - Info about an route
type RouteGet struct {
	// Protocol - Protocol type
	Protocol int
	// Flags - flag type
	Flags string
	// Gw - gateway information if any
	Gw string
	// LinkIndex - OS allocated index
	LinkIndex int
	// Dst - ip addr
	Dst string
	// index of the route
	HardwareMark int
	// statistic
	Statistic RouteGetEntryStatistic
	// sync
	Sync DpStatusT
}

// RouteMod - Info about a route
type RouteMod struct {
	// Protocol - Protocol type
	Protocol int
	// Flags - flag type
	Flags int
	// Gw - gateway information if any
	Gw net.IP
	// LinkIndex - OS allocated index
	LinkIndex int
	// Dst - ip addr
	Dst net.IPNet
}

// FwRuleOpts - Information related to Firewall options
type FwOptArg struct {
	// Drop - Drop any matching rule
	Drop bool `json:"drop"`
	// Trap - Trap anything matching rule
	Trap bool `json:"trap"`
	// Record - Record packets matching rule
	Record bool `json:"record"`
	// Redirect - Redirect any matching rule
	Rdr     bool   `json:"redirect"`
	RdrPort string `json:"redirectPortName"`
	// Allow - Allow any matching rule
	Allow bool `json:"allow"`
	// Mark - Mark the matching rule
	Mark uint32 `json:"fwMark"`
}

// FwRuleArg - Information related to firewall rule
type FwRuleArg struct {
	// SrcIP - Source IP in CIDR notation
	SrcIP string `json:"sourceIP"`
	// DstIP - Destination IP in CIDR notation
	DstIP string `json:"destinationIP"`
	// SrcPortMin - Minimum source port range
	SrcPortMin uint16 `json:"minSourcePort"`
	// SrcPortMax - Maximum source port range
	SrcPortMax uint16 `json:"maxSourcePort"`
	// DstPortMin - Minimum destination port range
	DstPortMin uint16 `json:"minDestinationPort"`
	// SrcPortMax - Maximum source port range
	DstPortMax uint16 `json:"maxDestinationPort"`
	// Proto - the protocol
	Proto uint8 `json:"protocol"`
	// InPort - the incoming port
	InPort string `json:"portName"`
	// Pref - User preference for ordering
	Pref uint16 `json:"preference"`
}

// FwRuleMod - Info related to a firewall entry
type FwRuleMod struct {
	// Rule - service argument of type FwRuleArg
	Rule FwRuleArg `json:"ruleArguments"`
	// Opts - firewall options
	Opts FwOptArg `json:"opts"`
}

// EndPointMod - Info related to a end-point entry
type EndPointMod struct {
	// HostName - hostname in CIDR
	HostName string `json:"hostName"`
	//  Name - Endpoint Identifier
	Name string `json:"name"`
	// InActTries - No. of inactive probes to mark
	// an end-point inactive
	InActTries int `json:"inactiveReTries"`
	// ProbeType - Type of probe : "icmp","connect-tcp", "connect-udp", "connect-sctp", "http", "https"
	ProbeType string `json:"probeType"`
	// ProbeReq - Request string in case of http probe
	ProbeReq string `json:"probeReq"`
	// ProbeResp - Response string in case of http probe
	ProbeResp string `json:"probeResp"`
	// ProbeDuration - How frequently (in seconds) to check activity
	ProbeDuration uint32 `json:"probeDuration"`
	// ProbePort - Port to probe for connect type
	ProbePort uint16 `json:"probePort"`
	// MinDelay - Minimum delay in this end-point
	MinDelay string `json:"minDelay"`
	// AvgDelay - Average delay in this end-point
	AvgDelay string `json:"avgDelay"`
	// MaxDelay - Max delay in this end-point
	MaxDelay string `json:"maxDelay"`
	// CurrState - Current state of this end-point
	CurrState string `json:"currState"`
}

// EpSelect - Selection method of load-balancer end-point
type EpSelect uint

const (
	// LbSelRr - select the lb end-points based on round-robin
	LbSelRr EpSelect = iota
	// LbSelHash - select the lb end-points based on hashing
	LbSelHash
	// LbSelPrio - select the lb based on weighted round-robin
	LbSelPrio
)

type LBMode int32

const (
	// LBModeDefault - Default Mode(DNAT)
	LBModeDefault LBMode = iota
	// LBModeOneArm - One Arm Mode
	LBModeOneArm
	// LBModeOneArm - Full NAT Mode
	LBModeFullNAT
	// LBModeDSR - DSR Mode
	LBModeDSR
)

// LbServiceArg - Information related to load-balancer service
type LbServiceArg struct {
	// ServIP - the service ip or vip  of the load-balancer rule
	ServIP string `json:"externalIP"`
	// ServPort - the service port of the load-balancer rule
	ServPort uint16 `json:"port"`
	// Proto - the service protocol of the load-balancer rule
	Proto string `json:"protocol"`
	// BlockNum - An arbitrary block num to further segregate a service
	BlockNum uint16 `json:"block"`
	// Sel - one of LbSelRr,LbSelHash, or LbSelHash
	Sel EpSelect `json:"sel"`
	// Bgp - export this rule with goBGP
	Bgp bool `json:"bgp"`
	// Monitor - monitor end-points of this rule
	Monitor bool `json:"monitor"`
	// Mode - NAT mode
	Mode LBMode `json:"mode"`
	// InactiveTimeout - Forced session reset after inactive timeout
	InactiveTimeout uint32 `json:"inactiveTimeout"`
}

// LbEndPointArg - Information related to load-balancer end-point
type LbEndPointArg struct {
	// EpIP - endpoint IP address
	EpIP string `json:"endpointIP"`
	// EpPort - endpoint Port
	EpPort uint16 `json:"targetPort"`
	// Weight - weight associated with end-point
	// Only valid for weighted round-robin selection
	Weight uint8 `json:"weight"`
	// State - current state of the end-point
	State string `json:"state"`
}

// LbSecIpArg - Secondary IP
type LbSecIpArg struct {
	// SecIP - Secondary IP address
	SecIP string `json:"secondaryIP"`
}

// LbRuleMod - Info related to a load-balancer entry
type LbRuleMod struct {
	// Serv - service argument of type LbServiceArg
	Serv LbServiceArg `json:"serviceArguments"`
	// SecIPs - Secondary IPs for SCTP multi-homed service
	SecIPs []LbSecIpArg `json:"secondaryIPs"`
	// Eps - slice containing LbEndPointArg
	Eps []LbEndPointArg `json:"endpoints"`
}

// CtInfo - Conntrack Information
type CtInfo struct {
	// Dip - destination ip address
	Dip net.IP `json:"destinationIP"`
	// Sip - source ip address
	Sip net.IP `json:"sourceIP"`
	// Dport - destination port information
	Dport uint16 `json:"destinationPort"`
	// Sport - source port information
	Sport uint16 `json:"sourcePort"`
	// Proto - IP protocol information
	Proto string `json:"protocol"`
	// CState - current state of conntrack
	CState string `json:"conntrackState"`
	// CAct - any related action
	CAct string `json:"conntrackAct"`
	// Pkts - packets tracked by ct entry
	Pkts uint64 `json:"packets"`
	// Bytes - bytes tracked by ct entry
	Bytes uint64 `json:"bytes"`
}

// UlClArg - ulcl argument information
type UlClArg struct {
	// Addr - filter ip addr
	Addr net.IP `json:"ulclIP"`
	// Qfi - qfi id related to this filter
	Qfi uint8 `json:"qfi"`
}

// SessTun - session tunnel(l3) information
type SessTun struct {
	// TeID - tunnel-id
	TeID uint32 `json:"TeID"`
	// Addr - tunnel ip addr of remote-end
	Addr net.IP `json:"tunnelIP"`
}

// ParamMod - Info related to a operational parameters
type ParamMod struct {
	// LogLevel - log level of loxilb
	LogLevel string `json:"logLevel"`
}

// Equal - check if two session tunnel entries are equal
func (ut *SessTun) Equal(ut1 *SessTun) bool {
	if ut.TeID == ut1.TeID && ut.Addr.Equal(ut1.Addr) {
		return true
	}
	return false
}

// SessionMod - information related to a user-session
type SessionMod struct {
	// Ident - unique identifier for this session
	Ident string `json:"ident"`
	// IP - ip address of the end-user of this session
	IP net.IP `json:"sessionIP"`
	// AnTun - access tunnel network information
	AnTun SessTun `json:"accessNetworkTunnel"`
	// CnTun - core tunnel network information
	CnTun SessTun `json:"coreNetworkTunnel"`
}

// SessionUlClMod - information related to a ulcl filter
type SessionUlClMod struct {
	// Ident - identifier of the session for this filter
	Ident string `json:"ulclIdent"`
	// Args - ulcl filter information
	Args UlClArg `json:"ulclArgument"`
}

// HASMod - information related to a cluster HA instance
type HASMod struct {
	// Instance - Cluster Instance
	Instance string `json:"instance"`
	// State - current HA state
	State string `json:"haState"`
	// Vip - Instance virtual IP address
	Vip net.IP `json:"Addr"`
}

// ClusterNodeMod - information related to a cluster node instance
type CluserNodeMod struct {
	// Instance - Cluster Instance
	Addr net.IP `json:"Addr"`
}

const (
	// PolTypeTrtcm - Policer type trtcm
	PolTypeTrtcm = 0 // Default
	// PolTypeSrtcm - Policer type srtcm
	PolTypeSrtcm = 1
)

// PolInfo - information related to a policer
type PolInfo struct {
	// PolType - one of PolTypeTrtcm or PolTypeSrtcm
	PolType int
	// ColorAware - color aware or not
	ColorAware bool
	// CommittedInfoRate - CIR in Mbps
	CommittedInfoRate uint64
	// PeakInfoRate - PIR in Mbps
	PeakInfoRate uint64
	// CommittedBlkSize -  CBS in bytes
	// 0 for default selection
	CommittedBlkSize uint64
	// ExcessBlkSize - EBS in bytes
	// 0 for default selection
	ExcessBlkSize uint64
}

// PolObjType - type  of a policer attachment object
type PolObjType uint

const (
	// PolAttachPort - attach policer to port
	PolAttachPort PolObjType = 1 << iota
	// PolAttachLbRule - attach policer to a rule
	PolAttachLbRule
)

// PolObj - Information related to policer attachment point
type PolObj struct {
	// PolObjName - name of the object
	PolObjName string
	// AttachMent - attach point type of the object
	AttachMent PolObjType
}

// PolMod - Information related to policer entry
type PolMod struct {
	// Ident - identifier
	Ident string
	// Info - policer info of type PolInfo
	Info PolInfo
	// Target - target object information
	Target PolObj
}

const (
	// MirrTypeSpan - simple SPAN
	MirrTypeSpan = 0 // Default
	// MirrTypeRspan - type RSPAN
	MirrTypeRspan = 1
	// MirrTypeErspan - type ERSPAN
	MirrTypeErspan = 2
)

// MirrInfo - information related to a mirror entry
type MirrInfo struct {
	// MirrType - one of MirrTypeSpan, MirrTypeRspan or MirrTypeErspan
	MirrType int
	// MirrPort - port where mirrored traffic needs to be sent
	MirrPort string
	// MirrVlan - for RSPAN we may need to send tagged mirror traffic
	MirrVlan int
	// MirrRip - RemoteIP. For ERSPAN we may need to send tunnelled mirror traffic
	MirrRip net.IP
	// MirrRip - SourceIP. For ERSPAN we may need to send tunnelled mirror traffic
	MirrSip net.IP
	// MirrTid - mirror tunnel-id. For ERSPAN we may need to send tunnelled mirror traffic
	MirrTid uint32
}

// MirrObjType - type of mirror attachment
type MirrObjType uint

const (
	// MirrAttachPort - mirror attachment to a port
	MirrAttachPort MirrObjType = 1 << iota
	// MirrAttachRule - mirror attachment to a lb rule
	MirrAttachRule
)

// MirrObj - information of object attached to mirror
type MirrObj struct {
	// MirrObjName - object name to be attached to mirror
	MirrObjName string
	// AttachMent - one of MirrAttachPort or MirrAttachRule
	AttachMent MirrObjType
}

// MirrMod - information related to a  mirror entry
type MirrMod struct {
	// Ident - unique identifier for the mirror
	Ident string
	// Info - information about the mirror
	Info MirrInfo
	// Target - information about object to which mirror needs to be attached
	Target MirrObj
}

// MirrGetMod - information related to Get a mirror entry
type MirrGetMod struct {
	// Ident - unique identifier for the mirror
	Ident string
	// Info - information about the mirror
	Info MirrInfo
	// Target - information about object to which mirror needs to be attached
	Target MirrObj
	// Sync - sync state
	Sync DpStatusT
}

// NetHookInterface - Go interface which needs to be implemented to talk to loxinet module
type NetHookInterface interface {
	NetMirrorGet() ([]MirrGetMod, error)
	NetMirrorAdd(*MirrMod) (int, error)
	NetMirrorDel(*MirrMod) (int, error)
	NetPortGet() ([]PortDump, error)
	NetPortAdd(*PortMod) (int, error)
	NetPortDel(*PortMod) (int, error)
	NetVlanGet() ([]VlanGet, error)
	NetVlanAdd(*VlanMod) (int, error)
	NetVlanDel(*VlanMod) (int, error)
	NetVlanPortAdd(*VlanPortMod) (int, error)
	NetVlanPortDel(*VlanPortMod) (int, error)
	NetFdbAdd(*FdbMod) (int, error)
	NetFdbDel(*FdbMod) (int, error)
	NetAddrGet() ([]IpAddrGet, error)
	NetAddrAdd(*IpAddrMod) (int, error)
	NetAddrDel(*IpAddrMod) (int, error)
	NetNeighGet() ([]NeighMod, error)
	NetNeighAdd(*NeighMod) (int, error)
	NetNeighDel(*NeighMod) (int, error)
	NetRouteGet() ([]RouteGet, error)
	NetRouteAdd(*RouteMod) (int, error)
	NetRouteDel(*RouteMod) (int, error)
	NetLbRuleAdd(*LbRuleMod) (int, error)
	NetLbRuleDel(*LbRuleMod) (int, error)
	NetLbRuleGet() ([]LbRuleMod, error)
	NetCtInfoGet() ([]CtInfo, error)
	NetSessionGet() ([]SessionMod, error)
	NetSessionUlClGet() ([]SessionUlClMod, error)
	NetSessionAdd(*SessionMod) (int, error)
	NetSessionDel(*SessionMod) (int, error)
	NetSessionUlClAdd(*SessionUlClMod) (int, error)
	NetSessionUlClDel(*SessionUlClMod) (int, error)
	NetPolicerGet() ([]PolMod, error)
	NetPolicerAdd(*PolMod) (int, error)
	NetPolicerDel(*PolMod) (int, error)
	NetCIStateMod(*HASMod) (int, error)
	NetCIStateGet() ([]HASMod, error)
	NetFwRuleAdd(*FwRuleMod) (int, error)
	NetFwRuleDel(*FwRuleMod) (int, error)
	NetFwRuleGet() ([]FwRuleMod, error)
	NetEpHostAdd(fm *EndPointMod) (int, error)
	NetEpHostDel(fm *EndPointMod) (int, error)
	NetEpHostGet() ([]EndPointMod, error)
	NetParamSet(param ParamMod) (int, error)
	NetParamGet(param *ParamMod) (int, error)
}
