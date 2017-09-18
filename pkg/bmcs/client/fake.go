package client

import (
	api "k8s.io/api/core/v1"

	baremetal "github.com/oracle/bmcs-go-sdk"
)

// FakeClient that should be used for testing
type FakeClient struct {
	Err                  error
	Calls                []string
	DefaultSecurityLists map[string]*baremetal.SecurityList
}

// NewFakeClient creates a new fake client for testing
func NewFakeClient() *FakeClient {
	return &FakeClient{
		Calls:                []string{},
		DefaultSecurityLists: make(map[string]*baremetal.SecurityList),
	}
}

func (f *FakeClient) call(s string) {
	f.Calls = append(f.Calls, s)
}

// Compartment fake
func (f *FakeClient) Compartment(id string) Interface {
	f.call("compartment")
	return nil
}

// GetInstanceByNodeName fake
func (f *FakeClient) GetInstanceByNodeName(name string) (*baremetal.Instance, error) {
	f.call("get-instance-by-node-name")
	return nil, f.Err
}

// GetNodeAddressesForInstance fake
func (f *FakeClient) GetNodeAddressesForInstance(id string) ([]api.NodeAddress, error) {
	f.call("get-node-addresses-for-instance")
	return nil, f.Err
}

// GetAttachedVnicsForInstance fake
func (f *FakeClient) GetAttachedVnicsForInstance(id string) ([]*baremetal.Vnic, error) {
	f.call("get-attached-vnics-for-instance")
	return nil, f.Err
}

// CreateAndAwaitLoadBalancer fake
func (f *FakeClient) CreateAndAwaitLoadBalancer(name, shape string, subnets []string) (*baremetal.LoadBalancer, error) {
	f.call("create-and-await-load-balancer")
	return nil, f.Err
}

// GetLoadBalancerByName fake
func (f *FakeClient) GetLoadBalancerByName(name string) (*baremetal.LoadBalancer, error) {
	f.call("get-load-balancer-by-name")
	return nil, f.Err
}

// GetCertificateByName fake
func (f *FakeClient) GetCertificateByName(loadBalancerID string, name string) (*baremetal.Certificate, error) {
	f.call("get-certificate-by-name")
	return nil, f.Err
}

// CreateAndAwaitBackendSet fake
func (f *FakeClient) CreateAndAwaitBackendSet(lb *baremetal.LoadBalancer, bs baremetal.BackendSet) (*baremetal.BackendSet, error) {
	f.call("create-and-await-backend-set")
	return nil, f.Err
}

// CreateAndAwaitListener fake
func (f *FakeClient) CreateAndAwaitListener(lb *baremetal.LoadBalancer, listener baremetal.Listener) error {
	f.call("create-and-await-listener")
	return f.Err
}

// CreateAndAwaitCertificate fake
func (f *FakeClient) CreateAndAwaitCertificate(lb *baremetal.LoadBalancer, name string, certificate string, key string) error {
	f.call("create-and-await-certificate")
	return f.Err
}

// AwaitWorkRequest fake
func (f *FakeClient) AwaitWorkRequest(id string) (*baremetal.WorkRequest, error) {
	f.call("await-work-request")
	return nil, f.Err
}

// GetSubnetsForInternalIPs fake
func (f *FakeClient) GetSubnetsForInternalIPs(ips []string) ([]*baremetal.Subnet, error) {
	f.call("get-subnets-for-internal-ips")
	return nil, f.Err
}

// GetDefaultSecurityList fake
func (f *FakeClient) GetDefaultSecurityList(subnet *baremetal.Subnet) (*baremetal.SecurityList, error) {
	f.call("get-default-security-list")
	for _, id := range subnet.SecurityListIDs {
		secList, ok := f.DefaultSecurityLists[id]
		if ok {
			return secList, nil
		}
	}
	return nil, f.Err
}

// Validate returns nil
func (f *FakeClient) Validate() error {
	f.call("validate")
	return f.Err
}

// GetInstance fake
func (f *FakeClient) GetInstance(id string) (*baremetal.Instance, error) {
	f.call("get-instance")
	return nil, f.Err
}

// GetSubnet fake
func (f *FakeClient) GetSubnet(ocid string) (*baremetal.Subnet, error) {
	f.call("get-subnet")
	return nil, f.Err
}

// UpdateSecurityList fake
func (f *FakeClient) UpdateSecurityList(id string, opts *baremetal.UpdateSecurityListOptions) (*baremetal.SecurityList, error) {
	f.call("update-security-list")
	return nil, f.Err
}

//CreateBackendSet fake
func (f *FakeClient) CreateBackendSet(
	loadBalancerID string,
	name string,
	policy string,
	backends []baremetal.Backend,
	healthChecker *baremetal.HealthChecker,
	sslConfig *baremetal.SSLConfiguration,
	sessionPersistenceConfig *baremetal.SessionPersistenceConfiguration,
	opts *baremetal.LoadBalancerOptions,
) (workRequestID string, e error) {
	f.call("create-backend-set")
	return "", f.Err
}

// UpdateBackendSet fake
func (f *FakeClient) UpdateBackendSet(loadBalancerID string, backendSetName string, opts *baremetal.UpdateLoadBalancerBackendSetOptions) (workRequestID string, e error) {
	f.call("update-backend-set")
	return "", f.Err
}

// DeleteBackendSet fake
func (f *FakeClient) DeleteBackendSet(loadBalancerID string, backendSetName string, opts *baremetal.ClientRequestOptions) (string, error) {
	f.call("delete-backend-set")
	return "", f.Err
}

// CreateListener fake
func (f *FakeClient) CreateListener(
	loadBalancerID string,
	name string,
	defaultBackendSetName string,
	protocol string,
	port int,
	sslConfig *baremetal.SSLConfiguration,
	opts *baremetal.LoadBalancerOptions,
) (workRequestID string, e error) {
	f.call("create-listener")
	return "", f.Err
}

// UpdateListener fake
func (f *FakeClient) UpdateListener(loadBalancerID string, listenerName string, opts *baremetal.UpdateLoadBalancerListenerOptions) (workRequestID string, e error) {
	f.call("update-listener")
	return "", f.Err
}

// DeleteListener fake
func (f *FakeClient) DeleteListener(loadBalancerID string, listenerName string, opts *baremetal.ClientRequestOptions) (workRequestID string, e error) {
	f.call("delete-listener")
	return "", f.Err
}

// DeleteLoadBalancer fake
func (f *FakeClient) DeleteLoadBalancer(id string, opts *baremetal.ClientRequestOptions) (string, error) {
	f.call("delete-load-balancer")
	return "", f.Err
}
