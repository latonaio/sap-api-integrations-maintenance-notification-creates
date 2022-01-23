package responses

type Item struct {
	D struct {
	MaintenanceNotification         string  `json:"MaintenanceNotification"`
	MaintenanceNotificationItem     string  `json:"MaintenanceNotificationItem"`
	MaintNotifItemText              string  `json:"MaintNotifItemText"`
	MaintNotifDamageCodeGroup       string  `json:"MaintNotifDamageCodeGroup"`
	MaintNotifDamageCodeGroupName   string  `json:"MaintNotifDamageCodeGroupName"`
	MaintNotificationDamageCode     string  `json:"MaintNotificationDamageCode"`
	MaintNotifDamageCodeName        string  `json:"MaintNotifDamageCodeName"`
	MaintNotifObjPrtCodeGroup       string  `json:"MaintNotifObjPrtCodeGroup"`
	MaintNotifObjPrtCodeGroupName   string  `json:"MaintNotifObjPrtCodeGroupName"`
	MaintNotifObjPrtCode            string  `json:"MaintNotifObjPrtCode"`
	MaintNotifObjPrtCodeName        string  `json:"MaintNotifObjPrtCodeName"`
	IsDeleted                       bool    `json:"IsDeleted"`
	} `json:"d"`
}
