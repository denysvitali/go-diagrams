package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azurenetworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureNetwork = &azurenetworkContainer{path: "assets/azure/network"}

func (c *azurenetworkContainer) DdosProtectionPlans(opts ...attr.Attribute) *node.Node {
	return node.New("ddos-protection-plans", attr.AssetImage("assets/azure/network/ddos-protection-plans.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) DnsPrivateZones(opts ...attr.Attribute) *node.Node {
	return node.New("dns-private-zones", attr.AssetImage("assets/azure/network/dns-private-zones.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) ExpressrouteCircuits(opts ...attr.Attribute) *node.Node {
	return node.New("expressroute-circuits", attr.AssetImage("assets/azure/network/expressroute-circuits.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) LoadBalancers(opts ...attr.Attribute) *node.Node {
	return node.New("load-balancers", attr.AssetImage("assets/azure/network/load-balancers.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) NetworkInterfaces(opts ...attr.Attribute) *node.Node {
	return node.New("network-interfaces", attr.AssetImage("assets/azure/network/network-interfaces.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) TrafficManagerProfiles(opts ...attr.Attribute) *node.Node {
	return node.New("traffic-manager-profiles", attr.AssetImage("assets/azure/network/traffic-manager-profiles.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) VirtualNetworkGateways(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-network-gateways", attr.AssetImage("assets/azure/network/virtual-network-gateways.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) VirtualNetworks(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-networks", attr.AssetImage("assets/azure/network/virtual-networks.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) VirtualWans(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-wans", attr.AssetImage("assets/azure/network/virtual-wans.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) ApplicationSecurityGroups(opts ...attr.Attribute) *node.Node {
	return node.New("application-security-groups", attr.AssetImage("assets/azure/network/application-security-groups.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) CdnProfiles(opts ...attr.Attribute) *node.Node {
	return node.New("cdn-profiles", attr.AssetImage("assets/azure/network/cdn-profiles.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) DnsZones(opts ...attr.Attribute) *node.Node {
	return node.New("dns-zones", attr.AssetImage("assets/azure/network/dns-zones.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) LocalNetworkGateways(opts ...attr.Attribute) *node.Node {
	return node.New("local-network-gateways", attr.AssetImage("assets/azure/network/local-network-gateways.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) NetworkSecurityGroupsClassic(opts ...attr.Attribute) *node.Node {
	return node.New("network-security-groups-classic", attr.AssetImage("assets/azure/network/network-security-groups-classic.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) NetworkWatcher(opts ...attr.Attribute) *node.Node {
	return node.New("network-watcher", attr.AssetImage("assets/azure/network/network-watcher.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) OnPremisesDataGateways(opts ...attr.Attribute) *node.Node {
	return node.New("on-premises-data-gateways", attr.AssetImage("assets/azure/network/on-premises-data-gateways.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) ServiceEndpointPolicies(opts ...attr.Attribute) *node.Node {
	return node.New("service-endpoint-policies", attr.AssetImage("assets/azure/network/service-endpoint-policies.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) ApplicationGateway(opts ...attr.Attribute) *node.Node {
	return node.New("application-gateway", attr.AssetImage("assets/azure/network/application-gateway.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) Connections(opts ...attr.Attribute) *node.Node {
	return node.New("connections", attr.AssetImage("assets/azure/network/connections.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) Firewall(opts ...attr.Attribute) *node.Node {
	return node.New("firewall", attr.AssetImage("assets/azure/network/firewall.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) FrontDoors(opts ...attr.Attribute) *node.Node {
	return node.New("front-doors", attr.AssetImage("assets/azure/network/front-doors.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) PublicIpAddresses(opts ...attr.Attribute) *node.Node {
	return node.New("public-ip-addresses", attr.AssetImage("assets/azure/network/public-ip-addresses.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) RouteFilters(opts ...attr.Attribute) *node.Node {
	return node.New("route-filters", attr.AssetImage("assets/azure/network/route-filters.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) ReservedIpAddressesClassic(opts ...attr.Attribute) *node.Node {
	return node.New("reserved-ip-addresses-classic", attr.AssetImage("assets/azure/network/reserved-ip-addresses-classic.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) RouteTables(opts ...attr.Attribute) *node.Node {
	return node.New("route-tables", attr.AssetImage("assets/azure/network/route-tables.png"), attr.Shape(attr.None))
}

func (c *azurenetworkContainer) VirtualNetworkClassic(opts ...attr.Attribute) *node.Node {
	return node.New("virtual-network-classic", attr.AssetImage("assets/azure/network/virtual-network-classic.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "azure.network"
  Register(namespace, "DdosProtectionPlans", azureNetwork.DdosProtectionPlans)
  Register(namespace, "DnsPrivateZones", azureNetwork.DnsPrivateZones)
  Register(namespace, "ExpressrouteCircuits", azureNetwork.ExpressrouteCircuits)
  Register(namespace, "LoadBalancers", azureNetwork.LoadBalancers)
  Register(namespace, "NetworkInterfaces", azureNetwork.NetworkInterfaces)
  Register(namespace, "TrafficManagerProfiles", azureNetwork.TrafficManagerProfiles)
  Register(namespace, "VirtualNetworkGateways", azureNetwork.VirtualNetworkGateways)
  Register(namespace, "VirtualNetworks", azureNetwork.VirtualNetworks)
  Register(namespace, "VirtualWans", azureNetwork.VirtualWans)
  Register(namespace, "ApplicationSecurityGroups", azureNetwork.ApplicationSecurityGroups)
  Register(namespace, "CdnProfiles", azureNetwork.CdnProfiles)
  Register(namespace, "DnsZones", azureNetwork.DnsZones)
  Register(namespace, "LocalNetworkGateways", azureNetwork.LocalNetworkGateways)
  Register(namespace, "NetworkSecurityGroupsClassic", azureNetwork.NetworkSecurityGroupsClassic)
  Register(namespace, "NetworkWatcher", azureNetwork.NetworkWatcher)
  Register(namespace, "OnPremisesDataGateways", azureNetwork.OnPremisesDataGateways)
  Register(namespace, "ServiceEndpointPolicies", azureNetwork.ServiceEndpointPolicies)
  Register(namespace, "ApplicationGateway", azureNetwork.ApplicationGateway)
  Register(namespace, "Connections", azureNetwork.Connections)
  Register(namespace, "Firewall", azureNetwork.Firewall)
  Register(namespace, "FrontDoors", azureNetwork.FrontDoors)
  Register(namespace, "PublicIpAddresses", azureNetwork.PublicIpAddresses)
  Register(namespace, "RouteFilters", azureNetwork.RouteFilters)
  Register(namespace, "ReservedIpAddressesClassic", azureNetwork.ReservedIpAddressesClassic)
  Register(namespace, "RouteTables", azureNetwork.RouteTables)
  Register(namespace, "VirtualNetworkClassic", azureNetwork.VirtualNetworkClassic)
}
