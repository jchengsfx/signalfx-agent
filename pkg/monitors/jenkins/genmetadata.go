// Code generated by monitor-code-gen. DO NOT EDIT.

package jenkins

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "jenkins"

const (
	groupHealth      = "health"
	groupJob         = "job"
	groupNode        = "node"
	groupNodeStatus  = "node/status"
	groupSlaveStatus = "slave/status"
)

var groupSet = map[string]bool{
	groupHealth:      true,
	groupJob:         true,
	groupNode:        true,
	groupNodeStatus:  true,
	groupSlaveStatus: true,
}

const (
	jenkinsJobBuildCount            = "jenkins.job.build.count"
	jenkinsJobTotalTime             = "jenkins.job.total.time"
	jenkinsNodeExecutorCount        = "jenkins.node.executor.count"
	jenkinsNodeExecutorInUse        = "jenkins.node.executor.in-use"
	jenkinsNodeHealthCheckScore     = "jenkins.node.health-check.score"
	jenkinsNodeHealthDiskSpace      = "jenkins.node.health.disk.space"
	jenkinsNodeHealthPlugins        = "jenkins.node.health.plugins"
	jenkinsNodeHealthTemporarySpace = "jenkins.node.health.temporary.space"
	jenkinsNodeHealthThreadDeadlock = "jenkins.node.health.thread-deadlock"
	jenkinsNodeOnlineStatus         = "jenkins.node.online.status"
	jenkinsNodeQueueSize            = "jenkins.node.queue.size"
	jenkinsNodeSlaveOnlineStatus    = "jenkins.node.slave.online.status"
	jenkinsNodeVMMemoryHeapUsage    = "jenkins.node.vm.memory.heap.usage"
	jenkinsNodeVMMemoryNonHeapUsed  = "jenkins.node.vm.memory.non-heap.used"
	jenkinsNodeVMMemoryTotalUsed    = "jenkins.node.vm.memory.total.used"
)

var metricSet = map[string]monitors.MetricInfo{
	jenkinsJobBuildCount:            {Type: datapoint.Counter, Group: groupJob},
	jenkinsJobTotalTime:             {Type: datapoint.Counter, Group: groupJob},
	jenkinsNodeExecutorCount:        {Type: datapoint.Gauge, Group: groupNode},
	jenkinsNodeExecutorInUse:        {Type: datapoint.Gauge, Group: groupNode},
	jenkinsNodeHealthCheckScore:     {Type: datapoint.Gauge, Group: groupNode},
	jenkinsNodeHealthDiskSpace:      {Type: datapoint.Gauge, Group: groupHealth},
	jenkinsNodeHealthPlugins:        {Type: datapoint.Gauge, Group: groupHealth},
	jenkinsNodeHealthTemporarySpace: {Type: datapoint.Gauge, Group: groupHealth},
	jenkinsNodeHealthThreadDeadlock: {Type: datapoint.Gauge, Group: groupHealth},
	jenkinsNodeOnlineStatus:         {Type: datapoint.Gauge, Group: groupNodeStatus},
	jenkinsNodeQueueSize:            {Type: datapoint.Gauge, Group: groupNode},
	jenkinsNodeSlaveOnlineStatus:    {Type: datapoint.Gauge, Group: groupSlaveStatus},
	jenkinsNodeVMMemoryHeapUsage:    {Type: datapoint.Gauge, Group: groupNode},
	jenkinsNodeVMMemoryNonHeapUsed:  {Type: datapoint.Gauge, Group: groupNode},
	jenkinsNodeVMMemoryTotalUsed:    {Type: datapoint.Gauge, Group: groupNode},
}

var defaultMetrics = map[string]bool{
	jenkinsJobBuildCount:            true,
	jenkinsJobTotalTime:             true,
	jenkinsNodeExecutorCount:        true,
	jenkinsNodeExecutorInUse:        true,
	jenkinsNodeHealthCheckScore:     true,
	jenkinsNodeHealthDiskSpace:      true,
	jenkinsNodeHealthPlugins:        true,
	jenkinsNodeHealthTemporarySpace: true,
	jenkinsNodeHealthThreadDeadlock: true,
	jenkinsNodeOnlineStatus:         true,
	jenkinsNodeQueueSize:            true,
	jenkinsNodeSlaveOnlineStatus:    true,
	jenkinsNodeVMMemoryHeapUsage:    true,
	jenkinsNodeVMMemoryNonHeapUsed:  true,
	jenkinsNodeVMMemoryTotalUsed:    true,
}

var groupMetricsMap = map[string][]string{
	groupHealth: []string{
		jenkinsNodeHealthDiskSpace,
		jenkinsNodeHealthPlugins,
		jenkinsNodeHealthTemporarySpace,
		jenkinsNodeHealthThreadDeadlock,
	},
	groupJob: []string{
		jenkinsJobBuildCount,
		jenkinsJobTotalTime,
	},
	groupNode: []string{
		jenkinsNodeExecutorCount,
		jenkinsNodeExecutorInUse,
		jenkinsNodeHealthCheckScore,
		jenkinsNodeQueueSize,
		jenkinsNodeVMMemoryHeapUsage,
		jenkinsNodeVMMemoryNonHeapUsed,
		jenkinsNodeVMMemoryTotalUsed,
	},
	groupNodeStatus: []string{
		jenkinsNodeOnlineStatus,
	},
	groupSlaveStatus: []string{
		jenkinsNodeSlaveOnlineStatus,
	},
}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "jenkins",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           false,
}
