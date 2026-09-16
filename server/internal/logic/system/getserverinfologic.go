package system

import (
	"context"
	"runtime"
	"time"

	"zero/internal/svc"
	"zero/internal/types"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	unitMB = 1024 * 1024
	unitGB = 1024 * unitMB
)

type GetServerInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServerInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServerInfoLogic {
	return &GetServerInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServerInfoLogic) GetServerInfo() (resp *types.ServerInfo, err error) {
	var info types.ServerInfo

	info.Os = l.getOsInfo()

	cpuInfo, err := l.getCpuInfo()
	if err != nil {
		return nil, errors.Wrap(err, "获取CPU信息失败")
	}
	info.Host = cpuInfo

	memInfo, err := l.getMemInfo()
	if err != nil {
		return nil, errors.Wrap(err, "获取内存信息失败")
	}
	info.Mem = memInfo

	diskInfo, err := l.getDiskInfo()
	if err != nil {
		return nil, errors.Wrap(err, "获取磁盘信息失败")
	}
	info.Disk = diskInfo

	resp = &info
	return resp, nil
}

func (l *GetServerInfoLogic) getOsInfo() types.Os {
	var o types.Os
	o.Arch = runtime.GOARCH
	o.System = runtime.GOOS
	o.Kernel = runtime.Version()

	if h, err := host.Info(); err == nil {
		o.Hostname = h.Hostname
		o.Platform = h.Platform
		o.Uptime = h.Uptime
		o.HostID = h.HostID
	}
	return o
}

func (l *GetServerInfoLogic) getCpuInfo() (types.Host, error) {
	var h types.Host

	infos, err := cpu.Info()
	if err != nil {
		return h, err
	}
	if len(infos) > 0 {
		h.Name = infos[0].ModelName
		h.Core = len(infos)
		h.Mhz = infos[0].Mhz
	}

	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return h, err
	}
	if len(percent) > 0 {
		h.IdleLoadPerc = percent[0]
		h.UserLoadPerc = 100 - percent[0]
	}
	return h, nil
}

func (l *GetServerInfoLogic) getMemInfo() (types.Mem, error) {
	var m types.Mem

	v, err := mem.VirtualMemory()
	if err != nil {
		return m, err
	}

	m.Total = v.Total / unitMB
	m.Used = v.Used / unitMB
	m.Free = v.Free / unitMB
	m.Perc = v.UsedPercent
	return m, nil
}

func (l *GetServerInfoLogic) getDiskInfo() ([]types.Disk, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	var disks []types.Disk
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}
		disks = append(disks, types.Disk{
			Path:  p.Mountpoint,
			Total: usage.Total / unitGB,
			Free:  usage.Free / unitGB,
			Used:  usage.Used / unitGB,
			Perc:  usage.UsedPercent,
		})
	}

	// 获取网络连接数作为补充信息
	connections, err := net.Connections("tcp")
	if err == nil {
		logx.WithContext(l.ctx).Infow("网络连接信息",
			logx.Field("module", "system"),
			logx.Field("action", "get_server_info"),
			logx.Field("tcp_connections", len(connections)),
		)
	}

	return disks, nil
}
