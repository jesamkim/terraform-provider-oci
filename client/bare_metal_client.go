// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package client

import "github.com/MustWin/baremetal-sdk-go"

type BareMetalClient interface {
	AddUserToGroup(userID, groupID string, opts *baremetal.RetryTokenOptions) (res *baremetal.UserGroupMembership, e error)

	AttachVolume(attachmentType, instanceID, volumeID string, opts *baremetal.CreateOptions) (res *baremetal.VolumeAttachment, e error)

	CaptureConsoleHistory(instanceID string, opts *baremetal.RetryTokenOptions) (icHistory *baremetal.ConsoleHistoryMetadata, e error)

	CreateBackend(loadBalancerID string, backendSetName string, ipAddr string, port int, opts *baremetal.CreateLoadBalancerBackendOptions) (workRequestID string, e error)
	CreateBackendSet(loadBalancerID string, name string, policy string, backends []baremetal.Backend, healthChecker *baremetal.HealthChecker, sslConfig *baremetal.SSLConfiguration, opts *baremetal.LoadBalancerOptions) (workRequestID string, e error)
	CreateBucket(compartmentID string, name string, namespaceName baremetal.Namespace, opts *baremetal.CreateBucketOptions) (bckt *baremetal.Bucket, e error)
	CreateCertificate(loadBalancerID string, certificateName string, caCertificate string, privateKey string, passphrase string, publicCertificate string, opts *baremetal.LoadBalancerOptions) (workRequestID string, e error)
	CreateCompartment(name, desc string, opts *baremetal.RetryTokenOptions) (res *baremetal.Compartment, e error)
	CreateCpe(compartmentID, ipAddress string, opts *baremetal.CreateOptions) (cpe *baremetal.Cpe, e error)
	CreateDHCPOptions(compartmentID, vcnID string, dhcpOptions []baremetal.DHCPDNSOption, opts *baremetal.CreateOptions) (res *baremetal.DHCPOptions, e error)
	CreateDrg(compartmentID string, opts *baremetal.CreateOptions) (res *baremetal.Drg, e error)
	CreateDrgAttachment(drgID, vcnID string, opts *baremetal.CreateOptions) (res *baremetal.DrgAttachment, e error)
	CreateGroup(name, desc string, opts *baremetal.RetryTokenOptions) (res *baremetal.Group, e error)
	CreateIPSecConnection(compartmentID, cpeID, drgID string, staticRoutes []string, opts *baremetal.CreateOptions) (conn *baremetal.IPSecConnection, e error)
	CreateImage(compartmentID, instanceID string, opts *baremetal.CreateOptions) (res *baremetal.Image, e error)
	CreateInternetGateway(compartmentID, vcnID string, isEnabled bool, opts *baremetal.CreateOptions) (gw *baremetal.InternetGateway, e error)
	CreateListener(loadBalancerID string, name string, defaultBackendSetName string, protocol string, port int, sslConfig *baremetal.SSLConfiguration, opts *baremetal.LoadBalancerOptions) (workRequestID string, e error)
	CreateLoadBalancer(backendSets *baremetal.BackendSet, certificates *baremetal.Certificate, compartmentID string, listeners *baremetal.Listener, shape string, subnetIDs []string, opts *baremetal.CreateOptions) (workRequestID string, e error)
	CreateOrResetUIPassword(userID string, opts *baremetal.RetryTokenOptions) (resource *baremetal.UIPassword, e error)
	CreatePolicy(name, desc, compartmentID string, statements []string, opts *baremetal.CreatePolicyOptions) (res *baremetal.Policy, e error)
	CreateRouteTable(compartmentID, vcnID string, routeRules []baremetal.RouteRule, opts *baremetal.CreateOptions) (res *baremetal.RouteTable, e error)
	CreateSecurityList(compartmentID, vcnID string, egressRules []baremetal.EgressSecurityRule, ingressRules []baremetal.IngressSecurityRule, opts *baremetal.CreateOptions) (res *baremetal.SecurityList, e error)
	CreateSubnet(availabilityDomain, cidrBlock, compartmentID, vcnID string, opts *baremetal.CreateSubnetOptions) (sn *baremetal.Subnet, e error)
	CreateSwiftPassword(userID, desc string, opts *baremetal.RetryTokenOptions) (res *baremetal.SwiftPassword, e error)
	CreateUser(name, desc string, opts *baremetal.RetryTokenOptions) (res *baremetal.User, e error)
	CreateVirtualNetwork(cidrBlock, compartmentID string, opts *baremetal.CreateVcnOptions) (vcn *baremetal.VirtualNetwork, e error)
	CreateVolume(availabilityDomain, compartmentID string, opts *baremetal.CreateVolumeOptions) (res *baremetal.Volume, e error)
	CreateVolumeBackup(volumeID string, opts *baremetal.CreateOptions) (vol *baremetal.VolumeBackup, e error)

	DeleteAPIKey(userID, fingerprint string, opts *baremetal.IfMatchOptions) (e error)
	DeleteBackend(loadBalancerID string, backendSetName string, backendName string, opts *baremetal.ClientRequestOptions) (workRequestID string, e error)
	DeleteBackendSet(loadBalancerID string, backendSetName string, opts *baremetal.ClientRequestOptions) (workRequestID string, e error)
	DeleteBucket(name string, namespaceName baremetal.Namespace, opts *baremetal.IfMatchOptions) (e error)
	DeleteCertificate(loadBalancerID string, certificateName string, opts *baremetal.ClientRequestOptions) (workRequestID string, e error)
	DeleteCpe(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteDHCPOptions(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteDrg(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteDrgAttachment(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteGroup(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteIPSecConnection(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteImage(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteInternetGateway(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteListener(loadBalancerID string, listenerName string, opts *baremetal.ClientRequestOptions) (workRequestID string, e error)
	DeleteLoadBalancer(id string, opts *baremetal.ClientRequestOptions) (workRequestID string, e error)
	DeleteObject(namespace baremetal.Namespace, bucketName string, objectName string, opts *baremetal.DeleteObjectOptions) (object *baremetal.DeleteObject, e error)
	DeletePolicy(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteRouteTable(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteSecurityList(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteSubnet(id string, opts *baremetal.IfMatchOptions) error
	DeleteSwiftPassword(id, userID string, opts *baremetal.IfMatchOptions) (e error)
	DeleteUser(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteUserGroupMembership(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteVirtualNetwork(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteVolume(id string, opts *baremetal.IfMatchOptions) (e error)
	DeleteVolumeBackup(id string, opts *baremetal.IfMatchOptions) (e error)

	DetachVolume(id string, opts *baremetal.IfMatchOptions) (e error)

	GetBackend(loadBalancerID string, backendSetName string, backendName string, opts *baremetal.ClientRequestOptions) (backend *baremetal.Backend, e error)
	GetBackendSet(loadBalancerID string, backendSetName string, opts *baremetal.ClientRequestOptions) (backendset *baremetal.BackendSet, e error)
	GetBucket(bucketName string, namespaceName baremetal.Namespace) (bckt *baremetal.Bucket, e error)
	GetCompartment(id string) (res *baremetal.Compartment, e error)
	GetConsoleHistory(instanceID string) (consoleHistoryMetadata *baremetal.ConsoleHistoryMetadata, e error)
	GetCpe(id string) (cpe *baremetal.Cpe, e error)
	GetDBHome(id string) (res *baremetal.DBHome, e error)
	GetDBNode(id string) (res *baremetal.DBNode, e error)
	GetDBSystem(id string) (res *baremetal.DBSystem, e error)
	GetDHCPOptions(id string) (res *baremetal.DHCPOptions, e error)
	GetDatabase(id string) (res *baremetal.Database, e error)
	GetDrg(id string) (res *baremetal.Drg, e error)
	GetDrgAttachment(id string) (res *baremetal.DrgAttachment, e error)
	GetGroup(id string) (res *baremetal.Group, e error)
	GetHealthChecker(loadBalancerID string, backendSetName string, opts *baremetal.ClientRequestOptions) (healthChecker *baremetal.HealthChecker, e error)
	GetIPSecConnection(id string) (conn *baremetal.IPSecConnection, e error)
	GetIPSecConnectionDeviceConfig(id string) (config *baremetal.IPSecConnectionDeviceConfig, e error)
	GetIPSecConnectionDeviceStatus(id string) (status *baremetal.IPSecConnectionDeviceStatus, e error)
	GetImage(id string) (res *baremetal.Image, e error)
	GetInstance(id string) (inst *baremetal.Instance, e error)
	GetInternetGateway(id string) (gw *baremetal.InternetGateway, e error)
	GetLoadBalancer(id string, opts *baremetal.ClientRequestOptions) (loadbalancer *baremetal.LoadBalancer, e error)
	GetNamespace() (*baremetal.Namespace, error)
	GetObject(namespace baremetal.Namespace, bucketName string, objectName string, opts *baremetal.GetObjectOptions) (object *baremetal.Object, e error)
	GetPolicy(id string) (res *baremetal.Policy, e error)
	GetRouteTable(id string) (res *baremetal.RouteTable, e error)
	GetSecurityList(id string) (res *baremetal.SecurityList, e error)
	GetSubnet(id string) (subnet *baremetal.Subnet, e error)
	GetUser(id string) (res *baremetal.User, e error)
	GetUserGroupMembership(id string) (res *baremetal.UserGroupMembership, e error)
	GetVirtualNetwork(id string) (vcn *baremetal.VirtualNetwork, e error)
	GetVnic(id string) (vnic *baremetal.Vnic, e error)
	GetVolume(id string) (res *baremetal.Volume, e error)
	GetVolumeAttachment(id string) (res *baremetal.VolumeAttachment, e error)
	GetVolumeBackup(id string) (vol *baremetal.VolumeBackup, e error)
	GetWorkRequest(workRequestID string, opts *baremetal.ClientRequestOptions) (workRequest *baremetal.WorkRequest, e error)

	HeadObject(namespace baremetal.Namespace, bucketName string, objectName string, opts *baremetal.HeadObjectOptions) (headObject *baremetal.HeadObject, e error)

	InstanceAction(id string, action baremetal.InstanceActions, opts *baremetal.HeaderOptions) (inst *baremetal.Instance, e error)

	LaunchDBSystem(availabilityDomain, compartmentID, shape, subnetID string, sshPublicKeys []string, cpuCoreCount uint64, opts *baremetal.LaunchDBSystemOptions) (res *baremetal.DBSystem, e error)
	LaunchInstance(availabilityDomain, compartmentID, image, shape, subnetID string, opts *baremetal.LaunchInstanceOptions) (inst *baremetal.Instance, e error)

	ListAPIKeys(userID string) (response *baremetal.ListAPIKeyResponses, e error)
	ListAvailabilityDomains(compartmentID string) (ads *baremetal.ListAvailabilityDomains, e error)
	ListBackends(loadBalancerID string, backendSetName string) (backends *baremetal.ListBackends, e error)
	ListBackendSets(loadBalancerID string, opts *baremetal.ClientRequestOptions) (backends *baremetal.ListBackendSets, e error)
	ListBuckets(compartmentID string, namespaceName baremetal.Namespace, opts *baremetal.ListBucketsOptions) (buckets *baremetal.ListBuckets, e error)
	ListCertificates(loadBalancerID string, opts *baremetal.ClientRequestOptions) (certs *baremetal.ListCertificates, e error)
	ListCompartments(opts *baremetal.ListOptions) (resources *baremetal.ListCompartments, e error)
	ListConsoleHistories(compartmentID string, opts *baremetal.ListConsoleHistoriesOptions) (icHistories *baremetal.ListConsoleHistories, e error)
	ListCpes(compartmentID string, opts *baremetal.ListOptions) (cpes *baremetal.ListCpes, e error)
	ListDBHomes(compartmentID, dbSystemID string, limit uint64, opts *baremetal.PageListOptions) (res *baremetal.ListDBHomes, e error)
	ListDBNodes(compartmentID, dbSystemID string, limit uint64, opts *baremetal.PageListOptions) (resources *baremetal.ListDBNodes, e error)
	ListDBSystemShapes(availabilityDomain, compartmentID string, limit uint64, opts *baremetal.PageListOptions) (resources *baremetal.ListDBSystemShapes, e error)
	ListDBSystems(compartmentID string, limit uint64, opts *baremetal.PageListOptions) (res *baremetal.ListDBSystems, e error)
	ListDBVersions(compartmentID string, limit uint64, opts *baremetal.PageListOptions) (resources *baremetal.ListDBVersions, e error)
	ListDHCPOptions(compartmentID, vcnID string, opts *baremetal.ListOptions) (res *baremetal.ListDHCPOptions, e error)
	ListDatabases(compartmentID, dbHomeID string, limit uint64, opts *baremetal.PageListOptions) (resources *baremetal.ListDatabases, e error)
	ListDrgAttachments(compartmentID string, opts *baremetal.ListDrgAttachmentsOptions) (res *baremetal.ListDrgAttachments, e error)
	ListDrgs(compartmentID string, opts *baremetal.ListOptions) (res *baremetal.ListDrgs, e error)
	ListGroups(opts *baremetal.ListOptions) (resources *baremetal.ListGroups, e error)
	ListIPSecConnections(compartmentID string, opts *baremetal.ListIPSecConnsOptions) (conns *baremetal.ListIPSecConnections, e error)
	ListImages(compartmentID string, opts *baremetal.ListImagesOptions) (res *baremetal.ListImages, e error)
	ListInstances(compartmentID string, opts *baremetal.ListInstancesOptions) (insts *baremetal.ListInstances, e error)
	ListInternetGateways(compartmentID, vcnID string, opts *baremetal.ListOptions) (list *baremetal.ListInternetGateways, e error)
	ListLoadBalancers(compartmentID string, opts *baremetal.ListOptions) (loadbalancers *baremetal.ListLoadBalancers, e error)
	ListLoadBalancerPolicies(compartmentID string, opts *baremetal.ListLoadBalancerPolicyOptions) (loadbalancerPolicies *baremetal.ListLoadBalancerPolicies, e error)
	ListLoadBalancerProtocols(compartmentID string, opts *baremetal.ListLoadBalancerPolicyOptions) (loadbalancerProtocols *baremetal.ListLoadBalancerProtocols, e error)
	ListLoadBalancerShapes(compartmentID string, opts *baremetal.ListLoadBalancerPolicyOptions) (loadbalancerShapes *baremetal.ListLoadBalancerShapes, e error)
	ListObjects(namespace baremetal.Namespace, bucket string, opts *baremetal.ListObjectsOptions) (objects *baremetal.ListObjects, e error)
	ListPolicies(compartmentID string, opts *baremetal.ListOptions) (resources *baremetal.ListPolicies, e error)
	ListRouteTables(compartmentID, vcnID string, opts *baremetal.ListOptions) (res *baremetal.ListRouteTables, e error)
	ListSecurityLists(compartmentID, vcnID string, opts *baremetal.ListOptions) (res *baremetal.ListSecurityLists, e error)
	ListShapes(compartmentID string, opts *baremetal.ListShapesOptions) (shapes *baremetal.ListShapes, e error)
	ListSubnets(compartmentID, vcnID string, opts *baremetal.ListOptions) (subnets *baremetal.ListSubnets, e error)
	ListSupportedOperations() (*baremetal.ListSupportedOperations, error)
	ListSwiftPasswords(userID string) (resources *baremetal.ListSwiftPasswords, e error)
	ListUserGroupMemberships(opts *baremetal.ListMembershipsOptions) (resources *baremetal.ListUserGroupMemberships, e error)
	ListUsers(opts *baremetal.ListOptions) (resources *baremetal.ListUsers, e error)
	ListVirtualNetworks(compartmentID string, opts *baremetal.ListOptions) (vcns *baremetal.ListVirtualNetworks, e error)
	ListVnicAttachments(compartmentID string, opts *baremetal.ListVnicAttachmentsOptions) (res *baremetal.ListVnicAttachments, e error)
	ListVolumeAttachments(compartmentID string, opts *baremetal.ListVolumeAttachmentsOptions) (res *baremetal.ListVolumeAttachments, e error)
	ListVolumeBackups(compartmentID string, opts *baremetal.ListBackupsOptions) (vols *baremetal.ListVolumeBackups, e error)
	ListVolumes(compartmentID string, opts *baremetal.ListVolumesOptions) (res *baremetal.ListVolumes, e error)
	ListWorkRequests(loadBalancerID string, opts *baremetal.ListLoadBalancerPolicyOptions) (workRequests *baremetal.ListWorkRequests, e error)

	PutObject(namespace baremetal.Namespace, bucketName string, objectName string, content []byte, opts *baremetal.PutObjectOptions) (object *baremetal.Object, e error)

	ShowConsoleHistoryData(instanceConsoleHistoryID string, opts *baremetal.ConsoleHistoryDataOptions) (response *baremetal.ConsoleHistoryData, e error)

	TerminateDBSystem(id string, opts *baremetal.IfMatchOptions) (e error)
	TerminateInstance(id string, opts *baremetal.IfMatchOptions) (e error)

	UpdateBackend(loadBalancerID string, backendSetName string, backendName string, opts *baremetal.CreateLoadBalancerBackendOptions) (workRequestID string, e error)
	UpdateBackendSet(loadBalancerID string, backendSetName string, opts *baremetal.UpdateLoadBalancerBackendSetOptions) (workRequestID string, e error)
	UpdateBucket(compartmentID string, name string, namespaceName baremetal.Namespace, opts *baremetal.UpdateBucketOptions) (bckt *baremetal.Bucket, e error)
	UpdateCompartment(id string, opts *baremetal.UpdateIdentityOptions) (res *baremetal.Compartment, e error)
	UpdateCpe(id string, opts *baremetal.IfMatchDisplayNameOptions) (cpe *baremetal.Cpe, e error)
	UpdateDHCPOptions(id string, opts *baremetal.UpdateDHCPDNSOptions) (res *baremetal.DHCPOptions, e error)
	UpdateDrg(id string, opts *baremetal.IfMatchDisplayNameOptions) (drg *baremetal.Drg, e error)
	UpdateDrgAttachment(id string, opts *baremetal.IfMatchDisplayNameOptions) (drg *baremetal.DrgAttachment, e error)
	UpdateGroup(id string, opts *baremetal.UpdateIdentityOptions) (res *baremetal.Group, e error)
	UpdateHealthChecker(loadBalancerID string, backendSetName string, healthCheckerOptions baremetal.HealthChecker, opts *baremetal.LoadBalancerOptions) (workRequestID string, e error)
	UpdateIPSecConnection(id string, opts *baremetal.IfMatchDisplayNameOptions) (conn *baremetal.IPSecConnection, e error)
	UpdateImage(id string, opts *baremetal.UpdateOptions) (res *baremetal.Image, e error)
	UpdateInstance(id string, opts *baremetal.UpdateOptions) (inst *baremetal.Instance, e error)
	UpdateInternetGateway(id string, opts *baremetal.UpdateGatewayOptions) (gw *baremetal.InternetGateway, e error)
	UpdateListener(loadBalancerID string, listenerName string, opts *baremetal.UpdateLoadBalancerListenerOptions) (workRequestID string, e error)
	UpdateLoadBalancer(id string, opts *baremetal.UpdateOptions) (workRequestID string, e error)
	UpdatePolicy(id string, opts *baremetal.UpdatePolicyOptions) (res *baremetal.Policy, e error)
	UpdateRouteTable(id string, opts *baremetal.UpdateRouteTableOptions) (res *baremetal.RouteTable, e error)
	UpdateSecurityList(id string, opts *baremetal.UpdateSecurityListOptions) (res *baremetal.SecurityList, e error)
	UpdateSubnet(id string, opts *baremetal.IfMatchDisplayNameOptions) (subnet *baremetal.Subnet, e error)
	UpdateSwiftPassword(id, userID string, opts *baremetal.UpdateIdentityOptions) (res *baremetal.SwiftPassword, e error)
	UpdateUser(id string, opts *baremetal.UpdateIdentityOptions) (res *baremetal.User, e error)
	UpdateVirtualNetwork(id string, opts *baremetal.IfMatchDisplayNameOptions) (vcn *baremetal.VirtualNetwork, e error)
	UpdateVolume(id string, opts *baremetal.UpdateOptions) (res *baremetal.Volume, e error)
	UpdateVolumeBackup(id string, opts *baremetal.IfMatchDisplayNameOptions) (vol *baremetal.VolumeBackup, e error)

	UploadAPIKey(userID, key string, opts *baremetal.RetryTokenOptions) (apiKey *baremetal.APIKey, e error)
}
