package alicloud

import (
	"context"
	"fmt"
	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/ecs"
	"github.com/denverdino/aliyungo/pvtz"
	"github.com/denverdino/aliyungo/slb"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cloud-provider-alibaba-cloud/cloud-controller-manager/utils"
	"k8s.io/klog"
	"strconv"
)

func NewDryRunClientSLB(key, secret, region string) *DryRunClientSLB {
	return &DryRunClientSLB{
		BaseClient: BaseClient{},
		slb:        slb.NewSLBClientWithSecurityToken4RegionalDomain(key, secret, "", common.Region(region)),
	}
}

type DryRunClientSLB struct {
	BaseClient
	// base slb client
	slb *slb.Client
}

func (c *DryRunClientSLB) GetCloudClient() interface{} { return c.slb }
func (c *DryRunClientSLB) SetCloudClient(cloud interface{}) {
	client, ok := cloud.(*slb.Client)
	if !ok {
		panic("cloud does not implement slb client")
	}
	c.slb = client
}

func (c *DryRunClientSLB) DescribeLoadBalancers(
	ctx context.Context,
	args *slb.DescribeLoadBalancersArgs,
) (loadBalancers []slb.LoadBalancerType, err error) {
	return c.slb.DescribeLoadBalancers(args)
}

func (c *DryRunClientSLB) DescribeLoadBalancerAttribute(
	ctx context.Context,
	loadBalancerId string,
) (loadBalancer *slb.LoadBalancerType, err error) {
	return c.slb.DescribeLoadBalancerAttribute(loadBalancerId)
}

func (c *DryRunClientSLB) DescribeLoadBalancerHTTPSListenerAttribute(
	ctx context.Context,
	loadBalancerId string,
	port int,
) (response *slb.DescribeLoadBalancerHTTPSListenerAttributeResponse, err error) {
	return c.slb.DescribeLoadBalancerHTTPSListenerAttribute(loadBalancerId, port)
}

func (c *DryRunClientSLB) DescribeLoadBalancerTCPListenerAttribute(
	ctx context.Context,
	loadBalancerId string,
	port int,
) (response *slb.DescribeLoadBalancerTCPListenerAttributeResponse, err error) {
	return c.slb.DescribeLoadBalancerTCPListenerAttribute(loadBalancerId, port)
}

func (c *DryRunClientSLB) DescribeLoadBalancerUDPListenerAttribute(
	ctx context.Context,
	loadBalancerId string,
	port int,
) (response *slb.DescribeLoadBalancerUDPListenerAttributeResponse, err error) {
	return c.slb.DescribeLoadBalancerUDPListenerAttribute(loadBalancerId, port)
}

func (c *DryRunClientSLB) DescribeLoadBalancerHTTPListenerAttribute(
	ctx context.Context,
	loadBalancerId string,
	port int,
) (response *slb.DescribeLoadBalancerHTTPListenerAttributeResponse, err error) {
	return c.slb.DescribeLoadBalancerHTTPListenerAttribute(loadBalancerId, port)
}

func (c *DryRunClientSLB) DescribeTags(ctx context.Context, args *slb.DescribeTagsArgs) (tags []slb.TagItemType, pagination *common.PaginationResult, err error) {
	return c.slb.DescribeTags(args)
}

func (c *DryRunClientSLB) DescribeVServerGroups(
	ctx context.Context,
	args *slb.DescribeVServerGroupsArgs,
) (response *slb.DescribeVServerGroupsResponse, err error) {
	return c.slb.DescribeVServerGroups(args)
}

func (c *DryRunClientSLB) DescribeVServerGroupAttribute(
	ctx context.Context,
	args *slb.DescribeVServerGroupAttributeArgs,
) (response *slb.DescribeVServerGroupAttributeResponse, err error) {
	return c.slb.DescribeVServerGroupAttribute(args)
}

func skey(svc *v1.Service) string {
	return fmt.Sprintf("%s/%s", svc.Namespace, svc.Name)
}

func unknown() *v1.Service {
	return &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "unknown",
			Namespace: "unkonwn",
		},
	}
}

func getService(ctx context.Context) *v1.Service {
	isvc := ctx.Value(utils.ContextService)
	if isvc == nil {
		return unknown()
	}
	svc, ok := isvc.(*v1.Service)
	if !ok {
		return unknown()
	}
	return svc
}

func getSlb(ctx context.Context) *slb.LoadBalancerType {
	islb := ctx.Value(utils.ContextSLB)
	if islb == nil {
		return &slb.LoadBalancerType{}
	}
	mslb, ok := islb.(*slb.LoadBalancerType)
	if !ok {
		return &slb.LoadBalancerType{}
	}
	return mslb
}

func getDryRunMsg(ctx context.Context) string {
	isMsg := ctx.Value(utils.DryRunMsg)
	if isMsg == nil {
		return ""
	}
	msg, ok := isMsg.(string)
	if !ok {
		return ""
	}
	return msg
}

func (c *DryRunClientSLB) CreateLoadBalancer(
	ctx context.Context,
	args *slb.CreateLoadBalancerArgs,
) (response *slb.CreateLoadBalancerResponse, err error) {
	mtype := "CreateLoadBalancer"
	svc := getService(ctx)
	utils.AddEvent(utils.SLB, skey(svc), "", "CreateSLB", utils.ERROR, "")
	return nil, fmt.Errorf("api %s should not be called", mtype)
}

func (c *DryRunClientSLB) SetLoadBalancerName(ctx context.Context, loadBalancerId string, loadBalancerName string) (err error) {
	mtype := "SetLoadBalancerName"
	svc := getService(ctx)
	lb := getSlb(ctx)
	utils.AddEvent(utils.SLB, skey(svc), lb.LoadBalancerId, "SetSLBName", utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) DeleteLoadBalancer(ctx context.Context, loadBalancerId string) (err error) {
	mtype := "DeleteLoadBalancer"
	svc := getService(ctx)
	lb := getSlb(ctx)
	utils.AddEvent(utils.SLB, skey(svc), lb.LoadBalancerId, "DeleteSLB", utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) SetLoadBalancerDeleteProtection(ctx context.Context, args *slb.SetLoadBalancerDeleteProtectionArgs) (err error) {
	mtype := "SetLoadBalancerDeleteProtection"
	svc := getService(ctx)
	lb := getSlb(ctx)
	utils.AddEvent(utils.SLB, skey(svc), lb.LoadBalancerId, "SetSLBDeleteProtection", utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) SetLoadBalancerModificationProtection(
	ctx context.Context,
	args *slb.SetLoadBalancerModificationProtectionArgs,
) (err error) {
	mtype := "SetLoadBalancerModificationProtection"
	svc := getService(ctx)
	lb := getSlb(ctx)
	utils.AddEvent(utils.SLB, skey(svc), lb.LoadBalancerId, "SetSLBModificationProtection", utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) ModifyLoadBalancerInstanceSpec(ctx context.Context, args *slb.ModifyLoadBalancerInstanceSpecArgs) (err error) {
	mtype := "ModifyLoadBalancerInstanceSpec"
	svc := getService(ctx)
	lb := getSlb(ctx)
	utils.AddEvent(utils.SLB, skey(svc), lb.LoadBalancerId, "ModifySLBSpec", utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}
func (c *DryRunClientSLB) ModifyLoadBalancerInternetSpec(ctx context.Context, args *slb.ModifyLoadBalancerInternetSpecArgs) (err error) {
	mtype := "ModifyLoadBalancerInternetSpec"
	svc := getService(ctx)
	lb := getSlb(ctx)
	utils.AddEvent(utils.SLB, skey(svc), lb.LoadBalancerId, "ModifyInternetSpec", utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) RemoveBackendServers(ctx context.Context, loadBalancerId string, backendServers []slb.BackendServerType) (result []slb.BackendServerType, err error) {
	mtype := "RemoveBackendServers"
	svc := getService(ctx)
	lb := getSlb(ctx)
	utils.AddEvent(utils.SLB, skey(svc), lb.LoadBalancerId, "RemoveBackendServer", utils.ERROR, "")
	return nil, fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) AddBackendServers(ctx context.Context, loadBalancerId string, backendServers []slb.BackendServerType) (result []slb.BackendServerType, err error) {
	mtype := "AddBackendServers"
	svc := getService(ctx)
	utils.AddEvent(utils.SLB, skey(svc), loadBalancerId, "AddBackendServer", utils.ERROR, "")
	return nil, fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) StopLoadBalancerListener(ctx context.Context, loadBalancerId string, port int) (err error) {
	mtype := "StopLoadBalancerListener"
	svc := getService(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/%d", skey(svc), port), loadBalancerId, "StopListener",
		utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) StartLoadBalancerListener(ctx context.Context, loadBalancerId string, port int) (err error) {
	mtype := "StartLoadBalancerListener"
	svc := getService(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/%d", skey(svc), port), loadBalancerId, "StartListener",
		utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) CreateLoadBalancerTCPListener(ctx context.Context, args *slb.CreateLoadBalancerTCPListenerArgs) (err error) {
	mtype := "CreateLoadBalancerTCPListener"
	svc := getService(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/%d", skey(svc), args.ListenerPort), args.LoadBalancerId,
		"CreateListener", utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) CreateLoadBalancerUDPListener(ctx context.Context, args *slb.CreateLoadBalancerUDPListenerArgs) (err error) {
	mtype := "CreateLoadBalancerUDPListener"
	svc := getService(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/%d", skey(svc), args.ListenerPort), args.LoadBalancerId,
		"CreateListener", utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) CreateLoadBalancerHTTPListener(ctx context.Context, args *slb.CreateLoadBalancerHTTPListenerArgs) (err error) {
	mtype := "CreateLoadBalancerHTTPListener"
	svc := getService(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/%d", skey(svc), args.ListenerPort), args.LoadBalancerId,
		"CreateListener", utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) CreateLoadBalancerHTTPSListener(ctx context.Context, args *slb.CreateLoadBalancerHTTPSListenerArgs) (err error) {
	mtype := "CreateLoadBalancerHTTPSListener"
	svc := getService(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/%d", skey(svc), args.ListenerPort), args.LoadBalancerId, "CreateListener",
		utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) DeleteLoadBalancerListener(ctx context.Context, loadBalancerId string, port int) (err error) {
	mtype := "DeleteLoadBalancerListener"
	svc := getService(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/%d", skey(svc), port), loadBalancerId, "DeleteListener", utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) SetLoadBalancerHTTPListenerAttribute(ctx context.Context, args *slb.SetLoadBalancerHTTPListenerAttributeArgs) (err error) {
	mtype := "SetLoadBalancerHTTPListenerAttribute"
	svc := getService(ctx)
	reason := getDryRunMsg(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/%d", skey(svc), args.ListenerPort), args.LoadBalancerId,
		"UpdateListener", utils.ERROR, reason)
	return fmt.Errorf("function should not be called on start %s, called reason: %s", mtype, reason)
}

func (c *DryRunClientSLB) SetLoadBalancerHTTPSListenerAttribute(ctx context.Context, args *slb.SetLoadBalancerHTTPSListenerAttributeArgs) (err error) {
	mtype := "SetLoadBalancerHTTPSListenerAttribute"
	svc := getService(ctx)
	reason := getDryRunMsg(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/%d", skey(svc), args.ListenerPort), args.LoadBalancerId,
		"UpdateListener", utils.ERROR, reason)
	return fmt.Errorf("function should not be called on start %s, called reason: %s", mtype, reason)
}

func (c *DryRunClientSLB) SetLoadBalancerTCPListenerAttribute(ctx context.Context, args *slb.SetLoadBalancerTCPListenerAttributeArgs) (err error) {
	mtype := "SetLoadBalancerTCPListenerAttribute"
	svc := getService(ctx)
	reason := getDryRunMsg(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/%d", skey(svc), args.ListenerPort), args.LoadBalancerId,
		"UpdateListener", utils.ERROR, reason)
	return fmt.Errorf("function should not be called on start %s, called reason: %s", mtype, reason)
}

func (c *DryRunClientSLB) SetLoadBalancerUDPListenerAttribute(ctx context.Context, args *slb.SetLoadBalancerUDPListenerAttributeArgs) (err error) {
	mtype := "SetLoadBalancerUDPListenerAttribute"
	svc := getService(ctx)
	reason := getDryRunMsg(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/%d", skey(svc), args.ListenerPort), args.LoadBalancerId,
		"UpdateListener", utils.ERROR, reason)
	return fmt.Errorf("function should not be called on start %s, called reason: %s", mtype, reason)
}

func (c *DryRunClientSLB) RemoveTags(ctx context.Context, args *slb.RemoveTagsArgs) error {
	mtype := "RemoveTags"
	svc := getService(ctx)
	utils.AddEvent(utils.SLB, skey(svc), args.LoadBalancerID,
		fmt.Sprintf("should not remove tags of slb %s ", args.LoadBalancerID), utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) AddTags(ctx context.Context, args *slb.AddTagsArgs) error {
	return c.slb.AddTags(args)
}

/*
From v1.9.3.313-g748f81e-aliyun, ccm sets backend type to eni in newly created Terway clusters.
*/
func (c *DryRunClientSLB) CreateVServerGroup(ctx context.Context, args *slb.CreateVServerGroupArgs) (response *slb.CreateVServerGroupResponse, err error) {
	mtype := "CreateVServerGroup"
	utils.AddEvent(utils.SLB, args.VServerGroupName, args.LoadBalancerId, "CreateVgroup", utils.NORMAL, "")
	klog.Errorf("function should not be called on start %s, CreateVServerGroupArgs %++v", mtype, args)
	return nil, fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) DeleteVServerGroup(ctx context.Context, args *slb.DeleteVServerGroupArgs) (response *slb.DeleteVServerGroupResponse, err error) {
	mtype := "DeleteVServerGroup"
	svc := getService(ctx)
	lb := getSlb(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/VGroupID/%s", skey(svc), args.VServerGroupId), lb.LoadBalancerId,
		"DeleteVgroup", utils.ERROR, "")
	return nil, fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) SetVServerGroupAttribute(ctx context.Context, args *slb.SetVServerGroupAttributeArgs) (response *slb.SetVServerGroupAttributeResponse, err error) {
	svc := getService(ctx)
	lb := getSlb(ctx)
	klog.Infof("update vgroup %s, svc %s, lb %s", args.VServerGroupId, svc, lb.LoadBalancerId)
	return c.slb.SetVServerGroupAttribute(args)
}

func (c *DryRunClientSLB) ModifyVServerGroupBackendServers(ctx context.Context, args *slb.ModifyVServerGroupBackendServersArgs) (response *slb.ModifyVServerGroupBackendServersResponse, err error) {
	mtype := "ModifyVServerGroupBackendServers"
	svc := getService(ctx)
	lb := getSlb(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/VGroupID/%s", skey(svc), args.VServerGroupId), lb.LoadBalancerId,
		"ModifyVgroup", utils.ERROR, "")
	return nil, fmt.Errorf("function should not be called on start %s", mtype)
}

/*
 From v1.9.3.239-g40d97e1-aliyun, ccm support ecs and eni together.
 If a svc who has ecs and eni backends together, it's normal to call the AddVServerGroupBackendServers api to add eci backend.
*/
func (c *DryRunClientSLB) AddVServerGroupBackendServers(ctx context.Context, args *slb.AddVServerGroupBackendServersArgs) (response *slb.AddVServerGroupBackendServersResponse, err error) {
	mtype := "AddVServerGroupBackendServers"
	svc := getService(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/VGroupID/%s", skey(svc), args.VServerGroupId), args.LoadBalancerId,
		"AddVgroup", utils.NORMAL, "")
	return nil, fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientSLB) RemoveVServerGroupBackendServers(ctx context.Context, args *slb.RemoveVServerGroupBackendServersArgs) (response *slb.RemoveVServerGroupBackendServersResponse, err error) {
	mtype := "RemoveVServerGroupBackendServers"
	svc := getService(ctx)
	lb := getSlb(ctx)
	utils.AddEvent(utils.SLB, fmt.Sprintf("%s/VGroupID/%s", skey(svc), args.VServerGroupId), lb.LoadBalancerId,
		"RemoveVgroup", utils.ERROR, "")
	return nil, fmt.Errorf("function should not be called on start %s", mtype)
}

// =====================================================================================================================

func NewDryRunClientINS(key, secret, region string) *DryRunClientINS {
	return &DryRunClientINS{
		BaseClient: BaseClient{},
		ecs:        ecs.NewECSClientWithSecurityToken(key, secret, "", common.Region(region)),
	}
}

type DryRunClientINS struct {
	BaseClient
	// base slb client
	ecs *ecs.Client
}

func (c *DryRunClientINS) GetCloudClient() interface{} { return c.ecs }
func (c *DryRunClientINS) SetCloudClient(cloud interface{}) {
	client, ok := cloud.(*ecs.Client)
	if !ok {
		panic("cloud does not implement slb client")
	}
	c.ecs = client
}

func (c *DryRunClientINS) AddTags(ctx context.Context, args *ecs.AddTagsArgs) error {
	return c.ecs.AddTags(args)
}

func (c *DryRunClientINS) DescribeInstances(
	ctx context.Context,
	args *ecs.DescribeInstancesArgs,
) (instances []ecs.InstanceAttributesType, pagination *common.PaginationResult, err error) {
	return c.ecs.DescribeInstances(args)
}
func (c *DryRunClientINS) DescribeNetworkInterfaces(
	ctx context.Context,
	args *ecs.DescribeNetworkInterfacesArgs,
) (resp *ecs.DescribeNetworkInterfacesResponse, err error) {
	return c.ecs.DescribeNetworkInterfaces(args)
}

func (c *DryRunClientINS) DescribeEipAddresses(
	ctx context.Context,
	args *ecs.DescribeEipAddressesArgs,
) (eipAddresses []ecs.EipAddressSetType, pagination *common.PaginationResult, err error) {
	return c.ecs.DescribeEipAddresses(args)
}

// =====================================================================================================================
func NewDryRunClientPVTZ(key, secret, region string) *DryRunClientPVTZ {
	return &DryRunClientPVTZ{
		BaseClient: BaseClient{},
		pvtz:       pvtz.NewPVTZClientWithSecurityToken4RegionalDomain(key, secret, "", common.Region(region)),
	}
}

type DryRunClientPVTZ struct {
	BaseClient
	// base slb client
	pvtz *pvtz.Client
}

func (c *DryRunClientPVTZ) GetCloudClient() interface{} { return c.pvtz }
func (c *DryRunClientPVTZ) SetCloudClient(cloud interface{}) {
	client, ok := cloud.(*pvtz.Client)
	if !ok {
		panic("cloud does not implement slb client")
	}
	c.pvtz = client
}

func (c *DryRunClientPVTZ) DescribeZones(ctx context.Context, args *pvtz.DescribeZonesArgs) (zones []pvtz.ZoneType, err error) {
	return c.pvtz.DescribeZones(args)
}

func (c *DryRunClientPVTZ) CheckZoneName(ctx context.Context, args *pvtz.CheckZoneNameArgs) (bool, error) {
	return c.pvtz.CheckZoneName(args)
}

func (c *DryRunClientPVTZ) DescribeZoneInfo(ctx context.Context, args *pvtz.DescribeZoneInfoArgs) (response *pvtz.DescribeZoneInfoResponse, err error) {
	return c.pvtz.DescribeZoneInfo(args)
}

func (c *DryRunClientPVTZ) DescribeRegions(ctx context.Context) (regions []pvtz.RegionType, err error) {
	return c.pvtz.DescribeRegions()
}

func (c *DryRunClientPVTZ) DescribeZoneRecords(ctx context.Context, args *pvtz.DescribeZoneRecordsArgs) (records []pvtz.ZoneRecordType, err error) {
	return c.pvtz.DescribeZoneRecords(args)
}

func (c *DryRunClientPVTZ) DescribeZoneRecordsByRR(ctx context.Context, zoneId string, rr string) (records []pvtz.ZoneRecordType, err error) {
	return c.pvtz.DescribeZoneRecordsByRR(zoneId, rr)
}

func (c *DryRunClientPVTZ) AddZone(ctx context.Context, args *pvtz.AddZoneArgs) (response *pvtz.AddZoneResponse, err error) {
	mtype := "pvtz.AddZone"
	svc := getService(ctx)
	utils.AddEvent(utils.PVTZ, skey(svc), args.ZoneName,
		fmt.Sprintf("should not add zone %s", args.ZoneName), utils.ERROR, "")
	return nil, fmt.Errorf("function should not be called on start %s", mtype)
}
func (c *DryRunClientPVTZ) DeleteZone(ctx context.Context, args *pvtz.DeleteZoneArgs) (err error) {
	mtype := "pvtz.DeleteZone"
	svc := getService(ctx)
	utils.AddEvent(utils.PVTZ, skey(svc), args.ZoneId,
		fmt.Sprintf("should not delete zone %s", args.ZoneId), utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}
func (c *DryRunClientPVTZ) UpdateZoneRemark(ctx context.Context, args *pvtz.UpdateZoneRemarkArgs) error {
	mtype := "pvtz.UpdateZoneRemark"
	svc := getService(ctx)
	utils.AddEvent(utils.PVTZ, skey(svc), args.ZoneId,
		fmt.Sprintf("should not update zone %s remark %s", args.ZoneId, args.Remark), utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}
func (c *DryRunClientPVTZ) BindZoneVpc(ctx context.Context, args *pvtz.BindZoneVpcArgs) (err error) {
	mtype := "pvtz.BindZoneVpc"
	svc := getService(ctx)
	utils.AddEvent(utils.PVTZ, skey(svc), args.ZoneId,
		fmt.Sprintf("should not bind zone %s to vpc", args.ZoneId), utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}
func (c *DryRunClientPVTZ) DeleteZoneRecordsByRR(ctx context.Context, zoneId string, rr string) error {
	mtype := "pvtz.DeleteZoneRecordsByRR"
	svc := getService(ctx)
	utils.AddEvent(utils.PVTZ, skey(svc), zoneId,
		fmt.Sprintf("should not delete zone %s records by rr %s", zoneId, rr), utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}
func (c *DryRunClientPVTZ) AddZoneRecord(ctx context.Context, args *pvtz.AddZoneRecordArgs) (response *pvtz.AddZoneRecordResponse, err error) {
	mtype := "pvtz.AddZoneRecord"
	svc := getService(ctx)
	utils.AddEvent(utils.PVTZ, skey(svc), args.ZoneId,
		fmt.Sprintf("should not add records to zone %s", args.ZoneId), utils.ERROR, "")
	return nil, fmt.Errorf("function should not be called on start %s", mtype)
}
func (c *DryRunClientPVTZ) UpdateZoneRecord(ctx context.Context, args *pvtz.UpdateZoneRecordArgs) (err error) {
	mtype := "pvtz.UpdateZoneRecord"
	svc := getService(ctx)
	utils.AddEvent(utils.PVTZ, skey(svc), strconv.Itoa(int(args.RecordId)),
		fmt.Sprintf("should not update record %d", args.RecordId), utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}
func (c *DryRunClientPVTZ) DeleteZoneRecord(ctx context.Context, args *pvtz.DeleteZoneRecordArgs) (err error) {
	mtype := "pvtz.DeleteZoneRecord"
	svc := getService(ctx)
	utils.AddEvent(utils.PVTZ, skey(svc), strconv.Itoa(int(args.RecordId)),
		fmt.Sprintf("should not delete record %s", args.RecordId), utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}
func (c *DryRunClientPVTZ) SetZoneRecordStatus(ctx context.Context, args *pvtz.SetZoneRecordStatusArgs) (err error) {
	mtype := "pvtz.SetZoneRecordStatus"
	svc := getService(ctx)
	utils.AddEvent(utils.PVTZ, skey(svc), strconv.Itoa(int(args.RecordId)),
		fmt.Sprintf("should not update record %s", args.RecordId), utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

// =====================================================================================================================

func NewDryRunClientRoute(key, secret, region string) *DryRunClientRoute {
	return &DryRunClientRoute{
		BaseClient: BaseClient{},
		ecs:        ecs.NewVPCClientWithSecurityToken4RegionalDomain(key, secret, "", common.Region(region)),
	}
}

type DryRunClientRoute struct {
	BaseClient
	// base slb client
	ecs *ecs.Client
}

func (c *DryRunClientRoute) GetCloudClient() interface{} { return c.ecs }
func (c *DryRunClientRoute) SetCloudClient(cloud interface{}) {
	client, ok := cloud.(*ecs.Client)
	if !ok {
		panic("cloud does not implement slb client")
	}
	c.ecs = client
}

func (c *DryRunClientRoute) DescribeVpcs(ctx context.Context, args *ecs.DescribeVpcsArgs) (vpcs []ecs.VpcSetType, pagination *common.PaginationResult, err error) {
	return c.ecs.DescribeVpcs(args)
}

func (c *DryRunClientRoute) DescribeVRouters(ctx context.Context, args *ecs.DescribeVRoutersArgs) (vrouters []ecs.VRouterSetType, pagination *common.PaginationResult, err error) {
	return c.ecs.DescribeVRouters(args)
}

func (c *DryRunClientRoute) DescribeRouteTables(ctx context.Context, args *ecs.DescribeRouteTablesArgs) (routeTables []ecs.RouteTableSetType, pagination *common.PaginationResult, err error) {
	return c.ecs.DescribeRouteTables(args)
}

func (c *DryRunClientRoute) DescribeRouteEntryList(ctx context.Context, args *ecs.DescribeRouteEntryListArgs) (response *ecs.DescribeRouteEntryListResponse, err error) {
	return c.ecs.DescribeRouteEntryList(args)
}

func (c *DryRunClientRoute) DeleteRouteEntry(ctx context.Context, args *ecs.DeleteRouteEntryArgs) error {
	mtype := "route.DeleteRouteEntry"
	utils.AddEvent(utils.VPC, args.RouteTableId, args.RouteTableId,
		fmt.Sprintf("should not delete route entry, table id %s, ecs %s", args.RouteTableId, args.NextHopId),
		utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientRoute) CreateRouteEntry(ctx context.Context, args *ecs.CreateRouteEntryArgs) error {
	mtype := "route.CreateRouteEntry"
	utils.AddEvent(utils.VPC, args.RouteTableId, args.RouteTableId,
		fmt.Sprintf("should not create route entry, table id %s, ecs %s", args.RouteTableId, args.NextHopId),
		utils.ERROR, "")
	return fmt.Errorf("function should not be called on start %s", mtype)
}

func (c *DryRunClientRoute) WaitForAllRouteEntriesAvailable(ctx context.Context, vrouterId string, routeTableId string, timeout int) error {
	return c.ecs.WaitForAllRouteEntriesAvailable(vrouterId, routeTableId, timeout)
}
