package main

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/frame/g"
)

// Gat1400 GA/T1400机动车对象
type Gat1400MotorVehicle struct {
	MotorVehicleListObject MotorVehicleListObject `json:"MotorVehicleListObject"`
}

// MotorVehicleListObject 机动车对象列表
type MotorVehicleListObject struct {
	MotorVehicleObject []MotorVehicleObject `json:"MotorVehicleObject"`
}

// MotorVehicleObject 机动车对象
type MotorVehicleObject struct {
	MotorVehicleID   string       `json:"MotorVehicleID"`
	InfoKind         string       `json:"InfoKind"`
	SourceID         string       `json:"SourceID"`
	HasPlate         string       `json:"HasPlate"`
	VehicleColor     string       `json:"VehicleColor"`
	VehicleClass     string       `json:"VehicleClass"`
	VehicleBrand     string       `json:"VehicleBrand"`
	PlateClass       string       `json:"PlateClass"`
	PlateColor       string       `json:"PlateColor"`
	PlateNo          string       `json:"PlateNo"`
	DeviceID         string       `json:"DeviceID"`
	PassTime         string       `json:"PassTime"`
	ParkingInOutType string       `json:"ParkingInOutType"`
	SubImageList     string       `json:"SubImageList"`
	SubImageListObj  SubImageList `json:"SubImageListObj"`
}

// SubImageList 图像列表
type SubImageList struct {
	SubImageInfoObject []SubImageInfoObject `json:"SubImageInfoObject"`
}

// SubImageInfoObject 图像详情
type SubImageInfoObject struct {
	ImageID     string `json:"ImageID"`
	EventSort   string `json:"EventSort"`
	DeviceID    string `json:"DeviceID"`
	StoragePath string `json:"StoragePath"`
	Type        string `json:"Type"`
	FileFormat  string `json:"FileFormat"`
	ShotTime    string `json:"ShotTime"`
	Width       int64  `json:"Width"`
	Height      int64  `json:"Height"`
	Data        string `json:"Data"`
}

func main() {
	j := "{\"AppearTime\":\"\",\"BrandReliability\":\"\",\"Calling\":\"\",\"CarOfVehicle\":\"\",\"DescOfFrontItem\":\"\",\"DescOfRearItem\":\"\",\"DeviceID\":\"11011900011320000116\",\"Direction\":\"\",\"DisappearTime\":\"\",\"DrivingStatusCode\":\"\",\"FilmColor\":\"\",\"HasPlate\":\"1\",\"HitMarkInfo\":\"\",\"Id\":\"5355953550d048fb9638dbfe3db06646\",\"InfoKind\":\"1\",\"InsertMongoTime\":\"2021-01-09 18:57:16.795719624 +0800 CST m=+89582.919806769\",\"IsAltered\":\"\",\"IsCovered\":\"\",\"IsDecked\":\"\",\"IsModified\":\"\",\"IsSuspicious\":\"\",\"LaneNo\":\"\",\"LeftTopX\":\"\",\"LeftTopY\":\"\",\"MarkTime\":\"\",\"MotorVehicleID\":\"1a7ad108385e4dcba6e168458af23a591610189836K1X2pR\",\"NameOfPassedRoad\":\"\",\"NumOfPassenger\":\"\",\"ParkingInOutType\":\"1\",\"PassTime\":\"20210109185716\",\"PlateCharReliability\":\"\",\"PlateClass\":\"99\",\"PlateColor\":\"6\",\"PlateDescribe\":\"\",\"PlateNo\":\"京P0MG30\",\"PlateNoAttach\":\"\",\"PlateReliability\":\"\",\"RearviewMirror\":\"\",\"RightBtmX\":\"\",\"RightBtmY\":\"\",\"SafetyBelt\":\"\",\"SideOfVehicle\":\"\",\"SourceID\":\"51d1ccac224d43f08daeb996dc8c955a8ZTdY7Gm1\",\"Speed\":\"\",\"StorageUrl1\":\"\",\"StorageUrl2\":\"\",\"StorageUrl3\":\"\",\"StorageUrl4\":\"\",\"StorageUrl5\":\"\",\"SubImageList\":\"{\\\"SubImageInfoObject\\\":[{\\\"Data\\\":\\\"f9cfaee690f94369ac495cf81112a43b.png\\\",\\\"DeviceID\\\":\\\"11011900011320000116\\\",\\\"EventSort\\\":\\\"12\\\",\\\"FileFormat\\\":\\\"Jpeg\\\",\\\"Height\\\":192,\\\"ImageID\\\":\\\"51d1ccac224d43f08daeb996dc8c955artA48Av01\\\",\\\"ShotTime\\\":\\\"20210109185716\\\",\\\"StoragePath\\\":\\\"\\\",\\\"Type\\\":\\\"01\\\",\\\"Width\\\":192},{\\\"Data\\\":\\\"a781132f3dbe4683be80daa0fe733629.png\\\",\\\"DeviceID\\\":\\\"11011900011320000116\\\",\\\"EventSort\\\":\\\"16\\\",\\\"FileFormat\\\":\\\"Jpeg\\\",\\\"Height\\\":192,\\\"ImageID\\\":\\\"51d1ccac224d43f08daeb996dc8c955a1Z2bHxJ02\\\",\\\"ShotTime\\\":\\\"20210109185716\\\",\\\"StoragePath\\\":\\\"\\\",\\\"Type\\\":\\\"02\\\",\\\"Width\\\":192}]}\",\"Sunvisor\":\"\",\"TollgateID\":\"\",\"UsingPropertiesCode\":\"\",\"VehicleBodyDesc\":\"\",\"VehicleBrand\":\"0\",\"VehicleChassis\":\"\",\"VehicleClass\":\"999\",\"VehicleColor\":\"6\",\"VehicleColorDepth\":\"\",\"VehicleDoor\":\"\",\"VehicleFrontItem\":\"\",\"VehicleHeight\":\"\",\"VehicleHood\":\"\",\"VehicleLength\":\"\",\"VehicleModel\":\"\",\"VehicleRearItem\":\"\",\"VehicleRoof\":\"\",\"VehicleShielding\":\"\",\"VehicleStyles\":\"\",\"VehicleTrunk\":\"\",\"VehicleWheel\":\"\",\"VehicleWidth\":\"\",\"VehicleWindow\":\"\",\"WheelPrintedPattern\":\"\"}"
	motorVehicle := &MotorVehicleObject{}
	err := json.Unmarshal([]byte(j), motorVehicle)
	if err != nil {
		fmt.Println(err)
		return
	}
	if motorVehicle.SubImageList != "" {
		subImageList := &SubImageList{}
		err := json.Unmarshal([]byte(motorVehicle.SubImageList), subImageList)
		if err != nil {
			g.Log().Debug("+++ShiYuanHui-subImageList-json.Unmarshal+++", err)
			return
		}
		g.Log().Debug("+++ShiYuanHui-motorVehicle.subImageListObj+++", subImageList)
		motorVehicle.SubImageListObj = *subImageList
		g.Log().Debug("+++ShiYuanHui-motorVehicle.SubImageListObj+++", motorVehicle.SubImageList)
		for _, SubImage := range motorVehicle.SubImageListObj.SubImageInfoObject {
			if SubImage.Data != "" {
				imageBase64 := "acd"
				SubImage.Data = imageBase64
			} else {
				g.Log().Debug("+++ShiYuanHui+++SubImage.Data is empty", err)
				return
			}
		}
	}
}
