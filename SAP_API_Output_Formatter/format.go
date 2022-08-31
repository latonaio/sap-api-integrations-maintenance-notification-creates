package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-maintenance-notification-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
		MaintenanceNotification:        data.MaintenanceNotification,
		MaintNotifInternalID:           data.MaintNotifInternalID,
		NotificationText:               data.NotificationText,
		MaintPriority:                  data.MaintPriority,
		NotificationType:               data.NotificationType,
		NotifProcessingPhase:           data.NotifProcessingPhase,
		NotifProcessingPhaseDesc:       data.NotifProcessingPhaseDesc,
		MaintPriorityDesc:              data.MaintPriorityDesc,
		CreationDate:                   data.CreationDate,
		LastChangeTime:                 data.LastChangeTime,
		LastChangeDate:                 data.LastChangeDate,
		LastChangeDateTime:             data.LastChangeDateTime,
		CreationTime:                   data.CreationTime,
		ReportedByUser:                 data.ReportedByUser,
		ReporterFullName:               data.ReporterFullName,
		PersonResponsible:              data.PersonResponsible,
		MalfunctionEffect:              data.MalfunctionEffect,
		MalfunctionEffectText:          data.MalfunctionEffectText,
		MalfunctionStartDate:           data.MalfunctionStartDate,
		MalfunctionStartTime:           data.MalfunctionStartTime,
		MalfunctionEndDate:             data.MalfunctionEndDate,
		MalfunctionEndTime:             data.MalfunctionEndTime,
		MaintNotificationCatalog:       data.MaintNotificationCatalog,
		MaintNotificationCode:          data.MaintNotificationCode,
		MaintNotificationCodeGroup:     data.MaintNotificationCodeGroup,
		CatalogProfile:                 data.CatalogProfile,
		NotificationCreationDate:       data.NotificationCreationDate,
		NotificationCreationTime:       data.NotificationCreationTime,
		NotificationTimeZone:           data.NotificationTimeZone,
		RequiredStartDate:              data.RequiredStartDate,
		RequiredStartTime:              data.RequiredStartTime,
		RequiredEndDate:                data.RequiredEndDate,
		RequiredEndTime:                data.RequiredEndTime,
		LatestAcceptableCompletionDate: data.LatestAcceptableCompletionDate,
		MaintenanceObjectIsDown:        data.MaintenanceObjectIsDown,
		MaintNotificationLongText:      data.MaintNotificationLongText,
		MaintNotifLongTextForEdit:      data.MaintNotifLongTextForEdit,
		TechnicalObject:                data.TechnicalObject,
		TechObjIsEquipOrFuncnlLoc:      data.TechObjIsEquipOrFuncnlLoc,
		TechnicalObjectLabel:           data.TechnicalObjectLabel,
		MaintenancePlanningPlant:       data.MaintenancePlanningPlant,
		MaintenancePlannerGroup:        data.MaintenancePlannerGroup,
		SuperiorTechnicalObject:        data.SuperiorTechnicalObject,
		SuperiorTechnicalObjectName:    data.SuperiorTechnicalObjectName,
		SuperiorObjIsEquipOrFuncnlLoc:  data.SuperiorObjIsEquipOrFuncnlLoc,
		SuperiorTechnicalObjectLabel:   data.SuperiorTechnicalObjectLabel,
		ManufacturerPartTypeName:       data.ManufacturerPartTypeName,
		TechObjIsEquipOrFuncnlLocDesc:  data.TechObjIsEquipOrFuncnlLocDesc,
		FunctionalLocation:             data.FunctionalLocation,
		FunctionalLocationLabelName:    data.FunctionalLocationLabelName,
		TechnicalObjectDescription:     data.TechnicalObjectDescription,
		AssetLocation:                  data.AssetLocation,
		LocationName:                   data.LocationName,
		BusinessArea:                   data.BusinessArea,
		CompanyCode:                    data.CompanyCode,
		TechnicalObjectCategory:        data.TechnicalObjectCategory,
		TechnicalObjectType:            data.TechnicalObjectType,
		MainWorkCenterPlant:            data.MainWorkCenterPlant,
		MainWorkCenter:                 data.MainWorkCenter,
		PlantName:                      data.PlantName,
		MaintenancePlannerGroupName:    data.MaintenancePlannerGroupName,
		MaintenancePlant:               data.MaintenancePlant,
		LocationDescription:            data.LocationDescription,
		MainWorkCenterText:             data.MainWorkCenterText,
		MainWorkCenterPlantName:        data.MainWorkCenterPlantName,
		MaintenancePlantName:           data.MaintenancePlantName,
		PlantSectionPersonRespName:     data.PlantSectionPersonRespName,
		PersonResponsibleName:          data.PersonResponsibleName,
		MaintenanceOrder:               data.MaintenanceOrder,
		MaintenanceOrderType:           data.MaintenanceOrderType,
		MaintenanceActivityType:        data.MaintenanceActivityType,
		MaintObjDowntimeDurationUnit:   data.MaintObjDowntimeDurationUnit,
		MaintObjectDowntimeDuration:    data.MaintObjectDowntimeDuration,
		MaintenancePlan:                data.MaintenancePlan,
		MaintenanceItem:                data.MaintenanceItem,
		TaskListGroup:                  data.TaskListGroup,
		TaskListGroupCounter:           data.TaskListGroupCounter,
		MaintenancePlanCallNumber:      data.MaintenancePlanCallNumber,
		MaintenanceTaskListType:        data.MaintenanceTaskListType,
		NotificationReferenceDate:      data.NotificationReferenceDate,
		NotificationReferenceTime:      data.NotificationReferenceTime,
		NotificationCompletionDate:     data.NotificationCompletionDate,
		CompletionTime:                 data.CompletionTime,
		AssetRoom:                      data.AssetRoom,
		MaintNotifExtReferenceNumber:   data.MaintNotifExtReferenceNumber,
		MaintNotifRejectionReasonCode:  data.MaintNotifRejectionReasonCode,
		MaintNotifRejectionRsnCodeTxt:  data.MaintNotifRejectionRsnCodeTxt,
		MaintNotifDetectionCatalog:     data.MaintNotifDetectionCatalog,
		MaintNotifDetectionCode:        data.MaintNotifDetectionCode,
		MaintNotifDetectionCodeText:    data.MaintNotifDetectionCodeText,
		MaintNotifDetectionCodeGroup:   data.MaintNotifDetectionCodeGroup,
		MaintNotifDetectionCodeGrpTxt:  data.MaintNotifDetectionCodeGrpTxt,
		MaintNotifProcessPhaseCode:     data.MaintNotifProcessPhaseCode,
		MaintNotifProcessSubPhaseCode:  data.MaintNotifProcessSubPhaseCode,
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
		MaintenanceNotification:       data.MaintenanceNotification,
		MaintenanceNotificationItem:   data.MaintenanceNotificationItem,
		MaintNotifItemText:            data.MaintNotifItemText,
		MaintNotifDamageCodeGroup:     data.MaintNotifDamageCodeGroup,
		MaintNotifDamageCodeGroupName: data.MaintNotifDamageCodeGroupName,
		MaintNotificationDamageCode:   data.MaintNotificationDamageCode,
		MaintNotifDamageCodeName:      data.MaintNotifDamageCodeName,
		MaintNotifObjPrtCodeGroup:     data.MaintNotifObjPrtCodeGroup,
		MaintNotifObjPrtCodeGroupName: data.MaintNotifObjPrtCodeGroupName,
		MaintNotifObjPrtCode:          data.MaintNotifObjPrtCode,
		MaintNotifObjPrtCodeName:      data.MaintNotifObjPrtCodeName,
		IsDeleted:                     data.IsDeleted,
	}

	return item, nil
}
