package request

type PlayAlarmParams struct {
	Duration int         `json:"alarm_duration"`
	Volume   AlarmVolume `json:"alarm_volume"`
	Type     string      `json:"alarm_type"`
}

type ControlChildParams struct {
	DeviceId    string             `json:"device_id"`
	RequestData ChildRequestParams `json:"requestData"`
}

type ChildRequestParams struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type AlarmVolume string

const (
	AlarmVolumeLow    AlarmVolume = "low"
	AlarmVolumeMedium AlarmVolume = "normal"
	AlarmVolumeHigh   AlarmVolume = "high"
)
