package server

type SystemsStatus struct {
	OdataContext string `json:"@odata.context"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	Actions      struct {
		ComputerSystemReset struct {
			ResetTypeRedfishAllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
			Target                          string   `json:"target"`
		} `json:"#ComputerSystem.Reset"`
	} `json:"Actions"`
	AssetTag         string `json:"AssetTag"`
	AvailableActions []struct {
		Action       string `json:"Action"`
		Capabilities []struct {
			AllowableValues []string `json:"AllowableValues"`
			PropertyName    string   `json:"PropertyName"`
		} `json:"Capabilities"`
	} `json:"AvailableActions"`
	Bios struct {
		Current struct {
			VersionString string `json:"VersionString"`
		} `json:"Current"`
	} `json:"Bios"`
	BiosVersion string `json:"BiosVersion"`
	Boot        struct {
		BootSourceOverrideEnabled             string   `json:"BootSourceOverrideEnabled"`
		BootSourceOverrideSupported           []string `json:"BootSourceOverrideSupported"`
		BootSourceOverrideTarget              string   `json:"BootSourceOverrideTarget"`
		UefiTargetBootSourceOverride          string   `json:"UefiTargetBootSourceOverride"`
		UefiTargetBootSourceOverrideSupported []string `json:"UefiTargetBootSourceOverrideSupported"`
	} `json:"Boot"`
	Description        string `json:"Description"`
	EthernetInterfaces struct {
		OdataID string `json:"@odata.id"`
	} `json:"EthernetInterfaces"`
	HostCorrelation struct {
		HostMACAddress []string `json:"HostMACAddress"`
		HostName       string   `json:"HostName"`
		IPAddress      []string `json:"IPAddress"`
	} `json:"HostCorrelation"`
	HostName     string `json:"HostName"`
	ID           string `json:"Id"`
	IndicatorLED string `json:"IndicatorLED"`
	LogServices  struct {
		OdataID string `json:"@odata.id"`
	} `json:"LogServices"`
	Manufacturer string `json:"Manufacturer"`
	Memory       struct {
		Status struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
		TotalSystemMemoryGB int `json:"TotalSystemMemoryGB"`
	} `json:"Memory"`
	MemorySummary struct {
		Status struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
		TotalSystemMemoryGiB int `json:"TotalSystemMemoryGiB"`
	} `json:"MemorySummary"`
	Model string `json:"Model"`
	Name  string `json:"Name"`
	Oem   struct {
		Hp struct {
			OdataType string `json:"@odata.type"`
			Actions   struct {
				HpComputerSystemExtPowerButton struct {
					PushTypeRedfishAllowableValues []string `json:"PushType@Redfish.AllowableValues"`
					Target                         string   `json:"target"`
				} `json:"#HpComputerSystemExt.PowerButton"`
				HpComputerSystemExtSystemReset struct {
					ResetTypeRedfishAllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
					Target                          string   `json:"target"`
				} `json:"#HpComputerSystemExt.SystemReset"`
			} `json:"Actions"`
			AvailableActions []struct {
				Action       string `json:"Action"`
				Capabilities []struct {
					AllowableValues []string `json:"AllowableValues"`
					PropertyName    string   `json:"PropertyName"`
				} `json:"Capabilities"`
			} `json:"AvailableActions"`
			Battery []struct {
				Condition       string `json:"Condition"`
				ErrorCode       int    `json:"ErrorCode"`
				FirmwareVersion string `json:"FirmwareVersion"`
				Index           int    `json:"Index"`
				MaxCapWatts     int    `json:"MaxCapWatts"`
				Model           string `json:"Model"`
				Present         string `json:"Present"`
				ProductName     string `json:"ProductName"`
				SerialNumber    string `json:"SerialNumber"`
				Spare           string `json:"Spare"`
			} `json:"Battery"`
			Bios struct {
				Backup struct {
					Date          string `json:"Date"`
					Family        string `json:"Family"`
					VersionString string `json:"VersionString"`
				} `json:"Backup"`
				Current struct {
					Date          string `json:"Date"`
					Family        string `json:"Family"`
					VersionString string `json:"VersionString"`
				} `json:"Current"`
				UefiClass int `json:"UefiClass"`
			} `json:"Bios"`
			DeviceDiscoveryComplete struct {
				AMSDeviceDiscovery  string `json:"AMSDeviceDiscovery"`
				DeviceDiscovery     string `json:"DeviceDiscovery"`
				SmartArrayDiscovery string `json:"SmartArrayDiscovery"`
			} `json:"DeviceDiscoveryComplete"`
			IntelligentProvisioningIndex    int      `json:"IntelligentProvisioningIndex"`
			IntelligentProvisioningLocation string   `json:"IntelligentProvisioningLocation"`
			IntelligentProvisioningVersion  string   `json:"IntelligentProvisioningVersion"`
			PostState                       string   `json:"PostState"`
			PowerAllocationLimit            int      `json:"PowerAllocationLimit"`
			PowerAutoOn                     string   `json:"PowerAutoOn"`
			PowerOnDelay                    string   `json:"PowerOnDelay"`
			PowerRegulatorMode              string   `json:"PowerRegulatorMode"`
			PowerRegulatorModesSupported    []string `json:"PowerRegulatorModesSupported"`
			TrustedModules                  []struct {
				Status string `json:"Status"`
			} `json:"TrustedModules"`
			Type           string `json:"Type"`
			VirtualProfile string `json:"VirtualProfile"`
			Links          struct {
				BIOS struct {
					Href string `json:"href"`
				} `json:"BIOS"`
				EthernetInterfaces struct {
					Href string `json:"href"`
				} `json:"EthernetInterfaces"`
				FirmwareInventory struct {
					Href string `json:"href"`
				} `json:"FirmwareInventory"`
				Memory struct {
					Href string `json:"href"`
				} `json:"Memory"`
				NetworkAdapters struct {
					Href string `json:"href"`
				} `json:"NetworkAdapters"`
				PCIDevices struct {
					Href string `json:"href"`
				} `json:"PCIDevices"`
				PCISlots struct {
					Href string `json:"href"`
				} `json:"PCISlots"`
				SUT struct {
					Href string `json:"href"`
				} `json:"SUT"`
				SecureBoot struct {
					Href string `json:"href"`
				} `json:"SecureBoot"`
				SmartStorage struct {
					Href string `json:"href"`
				} `json:"SmartStorage"`
				SoftwareInventory struct {
					Href string `json:"href"`
				} `json:"SoftwareInventory"`
			} `json:"links"`
		} `json:"Hp"`
	} `json:"Oem"`
	Power            string `json:"Power"`
	PowerState       string `json:"PowerState"`
	ProcessorSummary struct {
		Count  int    `json:"Count"`
		Model  string `json:"Model"`
		Status struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
	} `json:"ProcessorSummary"`
	Processors struct {
		Count           int    `json:"Count"`
		ProcessorFamily string `json:"ProcessorFamily"`
		Status          struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
	} `json:"Processors"`
	SKU          string `json:"SKU"`
	SerialNumber string `json:"SerialNumber"`
	Status       struct {
		Health string `json:"Health"`
		State  string `json:"State"`
	} `json:"Status"`
	SystemType string `json:"SystemType"`
	Type       string `json:"Type"`
	UUID       string `json:"UUID"`
	Links      struct {
		Chassis []struct {
			Href string `json:"href"`
		} `json:"Chassis"`
		EthernetInterfaces struct {
			Href string `json:"href"`
		} `json:"EthernetInterfaces"`
		Logs struct {
			Href string `json:"href"`
		} `json:"Logs"`
		ManagedBy []struct {
			Href string `json:"href"`
		} `json:"ManagedBy"`
		Processors struct {
			Href string `json:"href"`
		} `json:"Processors"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

func (systemsStatus *SystemsStatus) IsOK() bool {
	if systemsStatus.ProcessorSummary.Status.HealthRollUp != "OK" {
		return false
	}
	return systemsStatus.Status.Health == "OK"
}
