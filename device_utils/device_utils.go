package device_utils

import (
	"github.com/google/uuid"
	"github.com/yumaojun03/dmidecode"
	"runtime"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
	//"github.com/shirou/gopsutil/disk"
	//"github.com/shirou/gopsutil/host"
	//"github.com/shirou/gopsutil/mem"
	//"github.com/shirou/gopsutil/net"
)

func GetSystem() string {
	return runtime.GOOS
}

func GetDeviceUuid() string {
	return uuid.NewSHA1(uuid.NameSpaceDNS, uuid.NodeID()).String()
}

func GetDeviceInfo() map[string]interface{} {

	deviceInfo := make(map[string]interface{})

	dmi, err := dmidecode.New()

	if err == nil {

		//biosInfo, err := dmi.BIOS()
		//baseBoardInfo, err := dmi.BaseBoard()
		//dmi.Chassis()
		//dmi.MemoryArray()
		//dmi.MemoryDevice()
		//dmi.Onboard()
		//dmi.PortConnector()
		//dmi.Processor()
		//dmi.ProcessorCache()
		//dmi.Slot()
		//dmi.System()

		processorInfoArray, err := dmi.Processor()

		if err == nil {

			processorInfoMap := make(map[int]map[string]string)

			deviceInfo["processor"] = processorInfoMap

			for index, processorInfo := range processorInfoArray {

				currentProcessorInfoMap := make(map[string]string)

				//smbios.Header
				//SocketDesignation string                   `json:"socket_designation,omitempty"`
				//ProcessorType     ProcessorType            `json:"processor_type,omitempty"`
				//Family            ProcessorFamily          `json:"family,omitempty"`
				//Manufacturer      string                   `json:"manufacturer,omitempty"`
				//ID                ProcessorID              `json:"id,omitempty"`
				//Version           string                   `json:"version,omitempty"`
				//Voltage           ProcessorVoltage         `json:"voltage,omitempty"`
				//ExternalClock     uint16                   `json:"external_clock,omitempty"`
				//MaxSpeed          uint16                   `json:"max_speed,omitempty"`
				//CurrentSpeed      uint16                   `json:"current_speed,omitempty"`
				//Status            ProcessorStatus          `json:"status,omitempty"`
				//Upgrade           ProcessorUpgrade         `json:"upgrade,omitempty"`
				//L1CacheHandle     uint16                   `json:"l_1_cache_handle,omitempty"`
				//L2CacheHandle     uint16                   `json:"l_2_cache_handle,omitempty"`
				//L3CacheHandle     uint16                   `json:"l_3_cache_handle,omitempty"`
				//SerialNumber      string                   `json:"serial_number,omitempty"`
				//AssetTag          string                   `json:"asset_tag,omitempty"`
				//PartNumber        string                   `json:"part_number,omitempty"`
				//CoreCount         byte                     `json:"core_count,omitempty"`
				//CoreEnabled       byte                     `json:"core_enabled,omitempty"`
				//ThreadCount       byte                     `json:"thread_count,omitempty"`
				//Characteristics   ProcessorCharacteristics `json:"characteristics,omitempty"`
				//Family2           ProcessorFamily          `json:"family_2,omitempty"`

				currentProcessorInfoMap["manufacturer"] = processorInfo.Manufacturer
				currentProcessorInfoMap["version"] = processorInfo.Version
				currentProcessorInfoMap["maxSpeed"] = strconv.Itoa(int(processorInfo.MaxSpeed))
				currentProcessorInfoMap["currentSpeed"] = strconv.Itoa(int(processorInfo.CurrentSpeed))

				processorInfoMap[index] = currentProcessorInfoMap
			}

		}

	}
	return deviceInfo
}

func GetDeviceRunningInfo() map[string]interface{} {

	deviceRunningInfo := make(map[string]interface{})

	//cpuStatArray, err := cpu.Info()
	cpuPercentArray, err := cpu.Percent(time.Second, true)

	if err == nil {

		cpuRunningInfo := make(map[string]interface{})

		deviceRunningInfo["cpuRunningInfo"] = cpuRunningInfo

		var cpuPercentStrArray []string
		var cpuAllPercent float64

		for _, percent := range cpuPercentArray {
			cpuAllPercent = cpuAllPercent + percent
			cpuPercentStrArray = append(cpuPercentStrArray, strconv.FormatFloat(percent, 'f', 4, 64))
		}

		cpuRunningInfo["avgPercent"] = strconv.FormatFloat(cpuAllPercent/float64(len(cpuPercentStrArray)), 'f', 4, 64)
		cpuRunningInfo["corePercent"] = cpuPercentStrArray

	}

	//v, _ := mem.VirtualMemory()
	//d, _ := disk.Usage("/")
	//n, _ := host.Info()
	//nv, _ := net.IOCounters(true)
	//boottime, _ := host.BootTime()

	return deviceRunningInfo
}
