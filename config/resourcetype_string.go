package config

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Firewall-0]
	_ = x[LoadBalancer-1]
	_ = x[ServicePool-2]
	_ = x[ConsulCluster-3]
	_ = x[ConsulServer-4]
	_ = x[ConsulClient-5]
}

const _ResourceType_name = "firewallload-balancerservice-poolconsul-clusterconsul-serverconsul-client"

var _ResourceType_index = [...]uint8{0, 9, 21, 32, 45, 57, 69}

func (i ResourceType) String() string {
	if i < 0 || i >= ResourceType(len(_ResourceType_index)-1) {
		return "ResourceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ResourceType_name[_ResourceType_index[i]:_ResourceType_index[i+1]]
}
