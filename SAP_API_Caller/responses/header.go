package responses

type Header struct {
	D struct {
	MaintenanceNotification          string `json:"MaintenanceNotification"`
	MaintNotifInternalID             string `json:"MaintNotifInternalID"`
	NotificationText                 string `json:"NotificationText"`
	MaintPriority                    string `json:"MaintPriority"`
	NotificationType                 string `json:"NotificationType"`
	NotifProcessingPhase             string `json:"NotifProcessingPhase"`
	NotifProcessingPhaseDesc         string `json:"NotifProcessingPhaseDesc"`
	MaintPriorityDesc                string `json:"MaintPriorityDesc"`
	CreationDate                     string `json:"CreationDate"`
	LastChangeTime                   string `json:"LastChangeTime"`
	LastChangeDate                   string `json:"LastChangeDate"`
	LastChangeDateTime               string `json:"LastChangeDateTime"`
	CreationTime                     string `json:"CreationTime"`
	ReportedByUser                   string `json:"ReportedByUser"`
	ReporterFullName                 string `json:"ReporterFullName"`
	PersonResponsible                string `json:"PersonResponsible"`
	MalfunctionEffect                string `json:"MalfunctionEffect"`
	MalfunctionEffectText            string `json:"MalfunctionEffectText"`
	MalfunctionStartDate             string `json:"MalfunctionStartDate"`
	MalfunctionStartTime             string `json:"MalfunctionStartTime"`
	MalfunctionEndDate               string `json:"MalfunctionEndDate"`
	MalfunctionEndTime               string `json:"MalfunctionEndTime"`
	MaintNotificationCatalog         string `json:"MaintNotificationCatalog"`
	MaintNotificationCode            string `json:"MaintNotificationCode"`
	MaintNotificationCodeGroup       string `json:"MaintNotificationCodeGroup"`
	CatalogProfile                   string `json:"CatalogProfile"`
	NotificationCreationDate         string `json:"NotificationCreationDate"`
	NotificationCreationTime         string `json:"NotificationCreationTime"`
	NotificationTimeZone             string `json:"NotificationTimeZone"`
	RequiredStartDate                string `json:"RequiredStartDate"`
	RequiredStartTime                string `json:"RequiredStartTime"`
	RequiredEndDate                  string `json:"RequiredEndDate"`
	RequiredEndTime                  string `json:"RequiredEndTime"`
	LatestAcceptableCompletionDate   string `json:"LatestAcceptableCompletionDate"`
	MaintenanceObjectIsDown          bool   `json:"MaintenanceObjectIsDown"`
	MaintNotificationLongText        string `json:"MaintNotificationLongText"`
	MaintNotifLongTextForEdit        string `json:"MaintNotifLongTextForEdit"`
	TechnicalObject                  string `json:"TechnicalObject"`
	TechObjIsEquipOrFuncnlLoc        string `json:"TechObjIsEquipOrFuncnlLoc"`
	TechnicalObjectLabel             string `json:"TechnicalObjectLabel"`
	MaintenancePlanningPlant         string `json:"MaintenancePlanningPlant"`
	MaintenancePlannerGroup          string `json:"MaintenancePlannerGroup"`
	SuperiorTechnicalObject          string `json:"SuperiorTechnicalObject"`
	SuperiorTechnicalObjectName      string `json:"SuperiorTechnicalObjectName"`
	SuperiorObjIsEquipOrFuncnlLoc    string `json:"SuperiorObjIsEquipOrFuncnlLoc"`
	SuperiorTechnicalObjectLabel     string `json:"SuperiorTechnicalObjectLabel"`
	ManufacturerPartTypeName         string `json:"ManufacturerPartTypeName"`
	TechObjIsEquipOrFuncnlLocDesc    string `json:"TechObjIsEquipOrFuncnlLocDesc"`
	FunctionalLocation               string `json:"FunctionalLocation"`
	FunctionalLocationLabelName      string `json:"FunctionalLocationLabelName"`
	TechnicalObjectDescription       string `json:"TechnicalObjectDescription"`
	AssetLocation                    string `json:"AssetLocation"`
	LocationName                     string `json:"LocationName"`
	BusinessArea                     string `json:"BusinessArea"`
	CompanyCode                      string `json:"CompanyCode"`
	TechnicalObjectCategory          string `json:"TechnicalObjectCategory"`
	TechnicalObjectType              string `json:"TechnicalObjectType"`
	MainWorkCenterPlant              string `json:"MainWorkCenterPlant"`
	MainWorkCenter                   string `json:"MainWorkCenter"`
	PlantName                        string `json:"PlantName"`
	MaintenancePlannerGroupName      string `json:"MaintenancePlannerGroupName"`
	MaintenancePlant                 string `json:"MaintenancePlant"`
	LocationDescription              string `json:"LocationDescription"`
	MainWorkCenterText               string `json:"MainWorkCenterText"`
	MainWorkCenterPlantName          string `json:"MainWorkCenterPlantName"`
	MaintenancePlantName             string `json:"MaintenancePlantName"`
	PlantSectionPersonRespName       string `json:"PlantSectionPersonRespName"`
	PersonResponsibleName            string `json:"PersonResponsibleName"`
	MaintenanceOrder                 string `json:"MaintenanceOrder"`
	MaintenanceOrderType             string `json:"MaintenanceOrderType"`
	MaintenanceActivityType          string `json:"MaintenanceActivityType"`
	MaintObjDowntimeDurationUnit     string `json:"MaintObjDowntimeDurationUnit"`
	MaintObjectDowntimeDuration      string `json:"MaintObjectDowntimeDuration"`
	MaintenancePlan                  string `json:"MaintenancePlan"`
	MaintenanceItem                  string `json:"MaintenanceItem"`
	TaskListGroup                    string `json:"TaskListGroup"`
	TaskListGroupCounter             string `json:"TaskListGroupCounter"`
	MaintenancePlanCallNumber        int    `json:"MaintenancePlanCallNumber"`
	MaintenanceTaskListType          string `json:"MaintenanceTaskListType"`
	NotificationReferenceDate        string `json:"NotificationReferenceDate"`
	NotificationReferenceTime        string `json:"NotificationReferenceTime"`
	NotificationCompletionDate       string `json:"NotificationCompletionDate"`
	CompletionTime                   string `json:"CompletionTime"`
	AssetRoom                        string `json:"AssetRoom"`
	MaintNotifExtReferenceNumber     string `json:"MaintNotifExtReferenceNumber"`
	MaintNotifRejectionReasonCode    string `json:"MaintNotifRejectionReasonCode"`
	MaintNotifRejectionRsnCodeTxt    string `json:"MaintNotifRejectionRsnCodeTxt"`
	MaintNotifDetectionCatalog       string `json:"MaintNotifDetectionCatalog"`
	MaintNotifDetectionCode          string `json:"MaintNotifDetectionCode"`
	MaintNotifDetectionCodeText      string `json:"MaintNotifDetectionCodeText"`
	MaintNotifDetectionCodeGroup     string `json:"MaintNotifDetectionCodeGroup"`
	MaintNotifDetectionCodeGrpTxt    string `json:"MaintNotifDetectionCodeGrpTxt"`
	MaintNotifProcessPhaseCode       string `json:"MaintNotifProcessPhaseCode"`
	MaintNotifProcessSubPhaseCode    string `json:"MaintNotifProcessSubPhaseCode"`
		ToItem              struct {
			ToItemResults []Item `json:"results"`
		} `json:"to_Item"`
	} `json:"d"`
}
