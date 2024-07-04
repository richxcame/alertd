package server

type ChassisStatus struct {
	ODataContext string `json:"@odata.context"`
	ODataID      string `json:"@odata.id"`
	ODataType    string `json:"@odata.type"`
	ChassisType  string `json:"ChassisType"`
	ID           string `json:"Id"`
	Manufacturer string `json:"Manufacturer"`
	Model        string `json:"Model"`
	Name         string `json:"Name"`
	Oem          struct {
		Hp struct {
			ODataType string `json:"@odata.type"`
			Firmware  struct {
				PlatformDefinitionTable struct {
					Current struct {
						VersionString string `json:"VersionString"`
					} `json:"Current"`
				} `json:"PlatformDefinitionTable"`
				PowerManagementController struct {
					Current struct {
						VersionString string `json:"VersionString"`
					} `json:"Current"`
				} `json:"PowerManagementController"`
				PowerManagementControllerBootloader struct {
					Current struct {
						Family        string `json:"Family"`
						VersionString string `json:"VersionString"`
					} `json:"Current"`
				} `json:"PowerManagementControllerBootloader"`
				SASProgrammableLogicDevice struct {
					Current struct {
						VersionString string `json:"VersionString"`
					} `json:"Current"`
				} `json:"SASProgrammableLogicDevice"`
				SPSFirmwareVersionData struct {
					Current struct {
						VersionString string `json:"VersionString"`
					} `json:"Current"`
				} `json:"SPSFirmwareVersionData"`
				SystemProgrammableLogicDevice struct {
					Current struct {
						VersionString string `json:"VersionString"`
					} `json:"Current"`
				} `json:"SystemProgrammableLogicDevice"`
			} `json:"Firmware"`
			Location struct {
				LocationInRack struct {
					RackLdsPartNumber         string `json:"RackLdsPartNumber"`
					RackLdsProductDescription string `json:"RackLdsProductDescription"`
					RackUHeight               int    `json:"RackUHeight"`
					TagVersion                int    `json:"TagVersion"`
					ULocation                 string `json:"ULocation"`
				} `json:"LocationInRack"`
				LocationOfChassis struct {
					UUID string `json:"UUID"`
				} `json:"LocationOfChassis"`
			} `json:"Location"`
			Type string `json:"Type"`
		} `json:"Hp"`
	} `json:"Oem"`
	Power struct {
		ODataID string `json:"@odata.id"`
	} `json:"Power"`
	SKU          string `json:"SKU"`
	SerialNumber string `json:"SerialNumber"`
	Status       struct {
		Health string `json:"Health"`
		State  string `json:"State"`
	} `json:"Status"`
	Thermal struct {
		ODataID string `json:"@odata.id"`
	} `json:"Thermal"`
	Type  string `json:"Type"`
	Links struct {
		ComputerSystems []struct {
			Href string `json:"href"`
		} `json:"ComputerSystems"`
		ManagedBy struct {
			Href string `json:"href"`
		} `json:"ManagedBy"`
		PowerMetrics struct {
			Href string `json:"href"`
		} `json:"PowerMetrics"`
		ThermalMetrics struct {
			Href string `json:"href"`
		} `json:"ThermalMetrics"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}
