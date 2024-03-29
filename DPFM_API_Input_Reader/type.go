package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey                   string                          `json:"connection_key"`
	Result                          bool                            `json:"result"`
	RedisKey                        string                          `json:"redis_key"`
	Filepath                        string                          `json:"filepath"`
	APIStatusCode                   int                             `json:"api_status_code"`
	RuntimeSessionID                string                          `json:"runtime_session_id"`
	BusinessPartnerID               *int                            `json:"business_partner"`
	ServiceLabel                    string                          `json:"service_label"`
	APIType                         string                          `json:"APIType"`
	Header                          Header                          `json:"PlannedFreight"`
	APISchema                       string                          `json:"api_schema"`
	Accepter                        []string                        `json:"accepter"`
	Deleted                         bool                            `json:"deleted"`
}

type Header struct {
	PlannedFreight							int      `json:"PlannedFreight"`
	PlannedFreightType						*string  `json:"PlannedFreightType"`
	FreightAgreement                        *int     `json:"FreightAgreement"`
	FreightAgreementItem                    *int     `json:"FreightAgreementItem"`
	FreightAgreementItemAvailableFreight	*int     `json:"FreightAgreementItemAvailableFreight"`
	FreightType								*string  `json:"FreightType"`
	FreightSpec								*string	 `json:"FreightSpec"`
	FreightCalendar							*string	 `json:"FreightCalendar"`
	PlannedFreightDepartureDate				*string	 `json:"PlannedFreightDepartureDate"`
	PlannedFreightDepartureTime				*string	 `json:"PlannedFreightDepartureTime"`
	PlannedFreightArrivalDate				*string	 `json:"PlannedFreightArrivalDate"`
	PlannedFreightArrivalTime				*string	 `json:"PlannedFreightArrivalTime"`
	SupplyChainRelationshipID             	*int	 `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID     	*int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID	*int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipFreightID		*int     `json:"SupplyChainRelationshipFreightID"`
	FreightPartner							*int     `json:"FreightPartner"`
	Buyer									*int     `json:"Buyer"`
	Seller									*int     `json:"Seller"`
	DeliverToParty							*int     `json:"DeliverToParty"`
	DeliverToPlant							*string  `json:"DeliverToPlant"`
	DeliverFromParty						*int     `json:"DeliverFromParty"`
	DeliverFromPlant						*string  `json:"DeliverFromPlant"`
	PlannedFreightNumberInCharacter			*string  `json:"PlannedFreightNumberInCharacter"`
	PlannedFreightNumberDescription			*string  `json:"PlannedFreightNumberDescription"`
	FRPArea									*string  `json:"FRPArea"`
	FRPController							*string  `json:"FRPController"`
	FreightCapacityWeight					*float32 `json:"FreightCapacityWeight"`
	FreightCapacityWeightUnit				*string  `json:"FreightCapacityWeightUnit"`
	PlannedFreightLongText					*string  `json:"PlannedFreightLongText"`
	CreationDate							*string  `json:"CreationDate"`
	CreationTime							*string  `json:"CreationTime"`
	LastChangeDate							*string  `json:"LastChangeDate"`
	LastChangeTime							*string  `json:"LastChangeTime"`
	IsReleased								*bool    `json:"IsReleased"`
	IsMarkedForDeletion                     *bool    `json:"IsMarkedForDeletion"`
}
