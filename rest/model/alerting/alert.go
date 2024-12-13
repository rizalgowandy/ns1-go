package alerting

import "encoding/json"

type Alert struct {
	ID              *string         `json:"id,omitempty"`
	CreatedAt       *int64          `json:"created_at,omitempty"`
	CreatedBy       *string         `json:"created_by,omitempty"`
	Data            json.RawMessage `json:"data,omitempty"`
	Name            *string         `json:"name,omitempty"`
	NotifierListIds []string        `json:"notifier_list_ids"`
	RecordIds       []string        `json:"record_ids"`
	Subtype         *string         `json:"subtype,omitempty"`
	Type            *string         `json:"type,omitempty"`
	UpdatedAt       *int64          `json:"updated_at,omitempty"`
	UpdatedBy       *string         `json:"updated_by,omitempty"`
	ZoneNames       []string        `json:"zone_names"`
}

var (
	zoneAlertType   string = "zone"
	recordAlertType string = "record"
)

func NewZoneAlert(alertName string, subtype string, notifierListIds []string, zoneNames []string) *Alert {
	return &Alert{
		Name:            &alertName,
		Type:            &zoneAlertType,
		Subtype:         &subtype,
		Data:            nil,
		NotifierListIds: notifierListIds,
		ZoneNames:       zoneNames,
	}
}

func NewRecordAlert(notifierListIds []string, recordIds []string, subtype string) *Alert {
	return &Alert{
		Type:            &recordAlertType,
		Subtype:         &subtype,
		Data:            nil,
		NotifierListIds: notifierListIds,
		RecordIds:       recordIds,
	}
}
