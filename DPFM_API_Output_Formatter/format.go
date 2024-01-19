package dpfm_api_output_formatter

import (
	"data-platform-api-planned-freight-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.PlannedFreight,
			&pm.PlannedFreightType,
			&pm.FreightAgreement,
			&pm.FreightAgreementItem,
			&pm.FreightAgreementItemAvailableFreight,
			&pm.FreightType,
			&pm.FreightSpec,
			&pm.FreightCalendar,
			&pm.PlannedFreightDepartureDate,
			&pm.PlannedFreightDepartureTime,
			&pm.PlannedFreightArrivalDate,
			&pm.PlannedFreightArrivalTime,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,			
			&pm.SupplyChainRelationshipFreightID,
			&pm.FreightPartner,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.PlannedFreightNumberInCharacter,
			&pm.PlannedFreightNumberDescription,
			&pm.FRPArea,
			&pm.FRPController,
			&pm.FreightCapacityWeight,
			&pm.FreightCapacityWeightUnit,
			&pm.PlannedFreightLongText,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			PlannedFreight:							data.PlannedFreight,
			PlannedFreightType:						data.PlannedFreightType,
			FreightAgreement:						data.FreightAgreement,
			FreightAgreementItem:					data.FreightAgreementItem,
			FreightAgreementItemAvailableFreight:	data.FreightAgreementItemAvailableFreight,
			FreightType:							data.FreightType,
			FreightSpec:							data.FreightSpec,
			FreightCalendar:						data.FreightCalendar,
			PlannedFreightDepartureDate:			data.PlannedFreightDepartureDate,
			PlannedFreightDepartureTime:			data.PlannedFreightDepartureTime,
			PlannedFreightArrivalDate:				data.PlannedFreightArrivalDate,
			PlannedFreightArrivalTime:				data.PlannedFreightArrivalTime,
			SupplyChainRelationshipID:            	data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:		data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:	data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipFreightID:		data.SupplyChainRelationshipFreightID,
			FreightPartner:							data.FreightPartner,
			Buyer:									data.Buyer,
			Seller:									data.Seller,
			DeliverToParty:							data.DeliverToParty,
			DeliverToPlant:							data.DeliverToPlant,
			DeliverFromParty:						data.DeliverFromParty,
			DeliverFromPlant:						data.DeliverFromPlant,
			PlannedFreightNumberInCharacter:		data.PlannedFreightNumberInCharacter,
			PlannedFreightNumberDescription:		data.PlannedFreightNumberDescription,
			FRPArea:								data.FRPArea,
			FRPController:							data.FRPController,
			FreightCapacityWeight:					data.FreightCapacityWeight,
			FreightCapacityWeightUnit:				data.FreightCapacityWeightUnit,
			PlannedFreightLongText:					data.PlannedFreightLongText,
			CreationDate:							data.CreationDate,
			CreationTime:							data.CreationTime,
			LastChangeDate:							data.LastChangeDate,
			LastChangeTime:							data.LastChangeTime,
			IsReleased:								data.IsReleased,
			IsMarkedForDeletion:					data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}
