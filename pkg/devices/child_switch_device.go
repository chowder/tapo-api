package devices

import (
	"encoding/json"
	"fmt"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response/child_devices"
)

type DeviceS210 struct {
	h        *TapoHub
	deviceId string
}

func NewDeviceS210(h *TapoHub, deviceId string) *DeviceS210 {
	return &DeviceS210{h: h, deviceId: deviceId}
}

func (d *DeviceS210) On() error {
	_, err := d.h.ControlChild(d.deviceId, request.RequestSetDeviceInfo, request.PlugDeviceInfoParams{On: true})
	return err
}

func (d *DeviceS210) Off() error {
	_, err := d.h.ControlChild(d.deviceId, request.RequestSetDeviceInfo, request.PlugDeviceInfoParams{On: false})
	return err
}

func (d *DeviceS210) GetDeviceInfo() (childdevices.DeviceInfoS210, error) {
	ok, c, err := d.h.GetChildById(d.deviceId)
	if err != nil {
		return childdevices.DeviceInfoS210{}, err
	}

	if !ok {
		return childdevices.DeviceInfoS210{}, fmt.Errorf("somehow this device '%s' went missing", d.deviceId)
	}

	var data childdevices.DeviceInfoS210
	err = json.Unmarshal(c.raw, &data)
	if err != nil {
		return childdevices.DeviceInfoS210{}, err
	}

	return data, err
}
