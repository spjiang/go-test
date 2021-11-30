package example

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestString01(t *testing.T)  {
	str :=`
FPL-CCA1286-IS                                                   
   
  
 -B737/M-SDE3FGHIRWYZ/EB1                                          
   
  
 -ZBCZ0735                                                         
   
  
 -K0747S0690 SQ H156 P99 W56 DUGEB                                 
   
  
 -ZBAA0107 ZBHH                                                    
   
  
 -PBN/A1B1C1D1L1O2S2T1 NAV/ABAS RNP2 DOF/211115 REG/B5045
  SEL/MQBR CODE/7BB0DE
  RMK/ACAS II
`
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	fmt.Println(str)
	reg := regexp.MustCompile("[\\s]{1,}-")
	str = reg.ReplaceAllString(str, `${1}-`)
	fmt.Println(str)


	fmt.Println(GetDBTableSql("SPJIANG"))
}

// getDBTableSql 获取表结构
func GetDBTableSql(table string) string {
	return "CREATE TABLE " + table + " ( " + ` 
  Id int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  MsgId char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息ID',
  MsgType char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息类型',
  MsgClient varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息客户端',
  DeliverStatus tinyint(4) NOT NULL DEFAULT 0 COMMENT '投递状态：0-未投递；1-投递成功；2-投递失败',
  Msg text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  MsgOriginal text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  CreateTime bigint(10) NOT NULL,
  UpdateTime bigint(10) NOT NULL,
  RowTime timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY ("Id") USING BTREE,
  UNIQUE INDEX MsgId_unique(MsgId) USING BTREE,
  INDEX "CreateTime_index"(CreateTime) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
`
}

func TestString02(t *testing.T)  {
	str := "XHelloWorldX"
	content := str[0:2]
	fmt.Println(content)
}