
 -- ParkingBindCars 新增VillageId防止相同车牌号出现在不同小区
ALTER TABLE `SmartCity_Main`.ParkingBindCars add `VillageId` char(32) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '小区ID';

-- 更改时间长度
ALTER TABLE `SmartCity_Main`.`ParkingBindCars`
MODIFY COLUMN `StartTime` bigint(11) NOT NULL DEFAULT 0 COMMENT '开始时间' AFTER `PlateNumber`,
MODIFY COLUMN `EndTime` bigint(11) NOT NULL COMMENT '结束时间' AFTER `StartTime`;

-- car主表增加唯一索引，防止CarId重复（错误导入等情况）
ALTER TABLE `SmartCity_Main`.`P_IdentityCar` ADD UNIQUE INDEX `CarId`(`CarId`) USING BTREE;