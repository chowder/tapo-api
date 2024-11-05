package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response/child_devices"
)

type DeviceS210 struct {
	h    *TapoHub
	Info childdevices.DeviceInfoS210
}

func NewDeviceS210(h *TapoHub, info childdevices.DeviceInfoS210) *DeviceS210 {
	return &DeviceS210{h: h, Info: info}
}

func (d *DeviceS210) On() error {
	return d.h.ControlChild(d.Info.DeviceId, request.RequestSetDeviceInfo, request.PlugDeviceInfoParams{On: true})
}

func (d *DeviceS210) Off() error {
	return d.h.ControlChild(d.Info.DeviceId, request.RequestSetDeviceInfo, request.PlugDeviceInfoParams{On: false})
}
