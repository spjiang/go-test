package main

import (
	"errors"
	"regexp"
	"strings"
)

func main() {

	msg := `
		ZCZC GSJ0580 150613     
QU CIHZPCA CIHTZ8X
.     
QU CIHZPCA CIHTZ8X
.     
QU TAOZPCA PEKFP8X SHAFP8X SHAZPCA PEKZBCA CGQZPCA PVGZPCA NNGZPCA
CANZPCA CKGZPCA SYXZPCA CIHZPCA CIHTZ8X PEKZBCA PEKZBCA YYAZP8X
HUZUD1E TYNZPCA PEKZBCA HFEZPCA
.SHAUOHO 150611
PLN 16NOV
76 HO1076 B30CT CGQ1215(16NOV) TAO
77 HO1076 B30CT TAO1515(16NOV) PVG
78 HO1176 B30EP NNG0420(16NOV) PVG
79 HO1245 B30EQ SHA2320(15NOV) CKG
80 HO1246 B30EQ CKG0330(16NOV) SHA
81 HO1225 B30EQ SHA0705(16NOV) SYX
82 HO1226 B30EQ SYX1125(16NOV) SHA
83 HO1139 B30FC PVG2340(15NOV) CIH
84 HO1140 B30FC CIH0250(16NOV) PVG
85 HO1925 B320Z HUZ2310(15NOV) YYA
86 HO1925 B320Z YYA0140(16NOV) TYN
87 HO1926 B320Z TYN0430(16NOV) YYA
88 HO1926 B320Z YYA0740(16NOV) HUZ
89 HO1909 B320Z HUZ1025(16NOV) HFE
SI: TO BE CONTINUE
=

NNNN
`
	msgType := "PLN"
	isSita := IsSita(msg, msgType)
	if isSita {
		GetSitaBody(msg, msgType)
	}

}

func IsSita(msg string, msgType string) bool {
	// 检查是否包含报文类型
	if find := strings.Contains(msg, msgType); find {
		// 检查是否为SITA报文
		if find := strings.Contains(msg, "="); find {
			rex := regexp.MustCompile(`\(([^)]+)\)`)
			out := rex.FindAllStringSubmatch(msg, -1)
			if len(out) > 0 {
				// 判断包含括号内字符串长度
				if len(out[0][1]) < 10 {
					return true
				}
			}
		}
	}
	return false
}

func GetSitaBody(msg string, msgType string) (data [][]string, err error) {
	// 获取body内容
	regRex := regexp.MustCompile("(?s)" + msgType + ".*?NNNN")
	bodyOut := regRex.FindStringSubmatch(msg)
	if len(bodyOut) != 1 {
		err = errors.New("not match sita body")
		return
	}
	msgReg := regexp.MustCompile("((?m).*?$)")
	msgOut := msgReg.FindAllStringSubmatch(bodyOut[0], -1)
	if len(msgOut) <= 0 {
		err = errors.New("not match sita message")
	}

	for index := range msgOut {
		matchStr := msgOut[index][1]
		msgFields := strings.Split(matchStr, " ")
		if len(msgFields) > 4 {
			data = append(data, msgFields)
		} else {
			// 过滤掉...不符号定制规范.
		}
	}
	return
}
