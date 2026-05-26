package cpu

const cgroupRootDir = "/sys/fs/cgroup"

// cgroup Linux cgroup
type cgroup struct {
	cgroupSet map[string]string
}

// CPUCFSQuotaUs cpu.cfs_quota_us
func (c *cgroup) CPUCFSQuotaUs() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// CPUCFSPeriodUs cpu.cfs_period_us
func (c *cgroup) CPUCFSPeriodUs() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// CPUAcctUsage cpuacct.usage
func (c *cgroup) CPUAcctUsage() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// CPUAcctUsagePerCPU cpuacct.usage_percpu
func (c *cgroup) CPUAcctUsagePerCPU() ([]uint64, error) { _ = "STUB: not implemented"; return nil, nil }

// fix possible_cpu:https://www.ibm.com/support/knowledgecenter/en/linuxonibm/com.ibm.linux.z.lgdd/lgdd_r_posscpusparm.html

// CPUSetCPUs cpuset.cpus
func (c *cgroup) CPUSetCPUs() ([]uint64, error) { _ = "STUB: not implemented"; return nil, nil }

// CurrentcGroup get current process cgroup
func currentcGroup() (*cgroup, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint

// When dir is not equal to /, it must be in docker
