package main

import (
	"fmt"
	"os"
	"strings"
)

// 核酸接龙
var hesuanList map[string]interface{}

// 报备列表（家里有人）
var baobeiList map[string]interface{}

// 无人列表（家里有无人）
var notHomeList map[string]interface{}

func init() {
	hesuanList = make(map[string]interface{})

	parseWechatRecord(hesuanList)

	baobeiList = make(map[string]interface{})

	baobeiList["室号"] = nil     //-人数
	baobeiList["1808"] = nil   //--1
	baobeiList["1002"] = nil   // 1人
	baobeiList["102"] = nil    //-1
	baobeiList["2108"] = nil   // -2人
	baobeiList["2203"] = nil   // -1
	baobeiList["1407"] = nil   // -2
	baobeiList["205"] = nil    // 1
	baobeiList["2205"] = nil   // 2205-2
	baobeiList["1508"] = nil   // 2
	baobeiList["1503"] = nil   // -2
	baobeiList["2005"] = nil   // ～1
	baobeiList["2405"] = nil   // -1
	baobeiList["705"] = nil    // -2
	baobeiList["1301"] = nil   // ，2
	baobeiList["2103"] = nil   // 2
	baobeiList["606"] = nil    // 2
	baobeiList["2101"] = nil   // 1
	baobeiList["2607"] = nil   // –2
	baobeiList["906"] = nil    // -2人
	baobeiList["2507"] = nil   // -2
	baobeiList["2606"] = nil   //-1
	baobeiList["1905"] = nil   // 1
	baobeiList["901"] = nil    // -1
	baobeiList["103"] = nil    //  2
	baobeiList["2106"] = nil   // -1
	baobeiList["604"] = nil    // -1
	baobeiList["607"] = nil    // 1
	baobeiList["905"] = nil    // - 2
	baobeiList["2305"] = nil   // -1
	baobeiList["1602"] = nil   // -1
	baobeiList["1604"] = nil   // -1
	baobeiList["2503"] = nil   // -1
	baobeiList["1405"] = nil   // -1
	baobeiList["2801"] = nil   // 1
	baobeiList["1804"] = nil   // 1
	baobeiList["104"] = nil    // -2
	baobeiList["2407"] = nil   // -2
	baobeiList["406"] = nil    // -1
	baobeiList["702"] = nil    // 1
	baobeiList["1806"] = nil   // -1
	baobeiList["1205"] = nil   // -1
	baobeiList["2608"] = nil   // -4
	baobeiList["1703"] = nil   // -1
	baobeiList["2306"] = nil   // -2
	baobeiList["1106"] = nil   // -2
	baobeiList["605"] = nil    // -1
	baobeiList["601"] = nil    // -4
	baobeiList["1208"] = nil   // -4
	baobeiList["1202"] = nil   // -1
	baobeiList["1708"] = nil   // -2
	baobeiList["2302"] = nil   // -2
	baobeiList["902"] = nil    // - 1
	baobeiList["804"] = nil    // -1
	baobeiList["1402"] = nil   // -2
	baobeiList["1303"] = nil   // -1
	baobeiList["1908"] = nil   //-1
	baobeiList["1701"] = nil   // -1
	baobeiList["2506"] = nil   // -1
	baobeiList["1507"] = nil   // -2
	baobeiList["1502"] = nil   // -1
	baobeiList["808"] = nil    //-2
	baobeiList["402"] = nil    // -2
	baobeiList["206"] = nil    // -2
	baobeiList["608"] = nil    // -2
	baobeiList["1307"] = nil   //-1
	baobeiList["2404"] = nil   //-1
	baobeiList["2602—2"] = nil //
	baobeiList["2105"] = nil   // -1
	baobeiList["1504"] = nil   // -1
	baobeiList["1605"] = nil   // 3
	baobeiList["802–1"] = nil  //
	baobeiList["203"] = nil    // -1
	baobeiList["1601"] = nil   // -2
	baobeiList["1608"] = nil   // -2
	baobeiList["1506"] = nil   // -2
	baobeiList["2508"] = nil   // -3
	baobeiList["503"] = nil    // - 1
	baobeiList["2802"] = nil   // -2
	baobeiList["2201"] = nil   // 4
	baobeiList["1206"] = nil   //-1
	baobeiList["201"] = nil    // -2
	baobeiList["301"] = nil    // -1
	baobeiList["204"] = nil    //  –1
	baobeiList["1001"] = nil   //-1
	baobeiList["2104"] = nil   //  2人
	baobeiList["1707"] = nil   // 1
	baobeiList["806"] = nil    // -1
	baobeiList["1008"] = nil   // -2
	baobeiList["1408"] = nil   // - 4
	baobeiList["1102"] = nil   // -2
	baobeiList["1006"] = nil   // -1
	baobeiList["2703"] = nil   //-1
	baobeiList["2708"] = nil   // -1
	baobeiList["603"] = nil    //-2
	baobeiList["507"] = nil    // -1
	baobeiList["602"] = nil    // -2
	baobeiList["2207"] = nil   // -1
	baobeiList["1305"] = nil   //-1
	baobeiList["1302"] = nil   // -1
	baobeiList["903"] = nil    //-2
	baobeiList["1004"] = nil   // -2
	baobeiList["105"] = nil    // -1
	baobeiList["2008"] = nil   // -2
	baobeiList["405"] = nil    // -1
	baobeiList["801"] = nil    //  3
	baobeiList["502"] = nil    // 1
	baobeiList["107"] = nil    // -1
	baobeiList["2202"] = nil   // -1
	baobeiList["1207"] = nil   // 2
	baobeiList["1406"] = nil   //-2
	baobeiList["303"] = nil    // -1
	baobeiList["2603"] = nil   // -2
	baobeiList["304"] = nil    // - 1
	baobeiList["2004"] = nil   // 2
	baobeiList["505"] = nil    // -1
	baobeiList["404"] = nil    // -1
	baobeiList["1607"] = nil   // -1
	baobeiList["1603"] = nil   // -1
	baobeiList["1706"] = nil   // -2
	baobeiList["2006"] = nil   // -1
	baobeiList["2206"] = nil   // -1
	baobeiList["1203"] = nil   //-2
	baobeiList["2204"] = nil   // -1
	baobeiList["2304"] = nil   // 2
	baobeiList["208"] = nil    // -1
	baobeiList["706"] = nil    // -1
	baobeiList["1501"] = nil   // -2人
	baobeiList["2702"] = nil   // -1
	baobeiList["1007"] = nil   //-2人
	baobeiList["1103"] = nil   // -2
	baobeiList["708"] = nil    // -1
	baobeiList["1903"] = nil   // 2人
	baobeiList["207"] = nil    // -1
	baobeiList["704"] = nil    // 2人
	baobeiList["803"] = nil    // 2人
	baobeiList["2307"] = nil   // 2人
	baobeiList["701"] = nil    // 2人
	baobeiList["908"] = nil    // 2人
	baobeiList["2707"] = nil   // 1人
	baobeiList["202"] = nil    // 1人
	baobeiList["805"] = nil    // 1人
	baobeiList["1505"] = nil   // 1人
	baobeiList["1403"] = nil   //  1人
	baobeiList["1306"] = nil   // 1人
	baobeiList["2706"] = nil   // 1
	baobeiList["902"] = nil    //-2人
	baobeiList["508"] = nil    // 3人
	baobeiList["307"] = nil    // 1人
	baobeiList["1901"] = nil   // -2人
	baobeiList["106"] = nil    // 2
	baobeiList["1105"] = nil   // -1人
	baobeiList["305"] = nil    // -1
	baobeiList["306"] = nil    // -1
	baobeiList["1005"] = nil   // -1人
	baobeiList["1704"] = nil   //-2
	baobeiList["101"] = nil    // -3
	baobeiList["1906"] = nil   // -1
	baobeiList["1702"] = nil   // -1
	baobeiList["2701"] = nil   // -1
	baobeiList["407"] = nil    // -2
	baobeiList["1807"] = nil   //-2
	baobeiList["1803"] = nil   // 两人
	baobeiList["308"] = nil    // 2人
	baobeiList["707"] = nil    // - 1
	baobeiList["1801"] = nil   // _3人
	baobeiList["1003"] = nil   // - 1
	baobeiList["2408"] = nil   //  - 2
	baobeiList["2704"] = nil   // -2
	baobeiList["1101"] = nil   // -2
	baobeiList["2001"] = nil   // -2
	baobeiList["2605"] = nil   // -3
	baobeiList["1705"] = nil   // -1
	baobeiList["2502"] = nil   // - 1
	baobeiList["2102"] = nil   // -1
	baobeiList["2401"] = nil   // -2
	baobeiList["108"] = nil    //-1
	baobeiList["2504"] = nil   // -1
	baobeiList["1902"] = nil   // -2
	baobeiList["2402"] = nil   // -2
	baobeiList["1904"] = nil   // -2
	baobeiList["907"] = nil    // 2人
	baobeiList["408"] = nil    //  -4
	baobeiList["1308"] = nil   //-3
	baobeiList["2303"] = nil   // -1
	baobeiList["2705"] = nil   // -2
	baobeiList["2308"] = nil   // -2
	baobeiList["807"] = nil    // -1
	baobeiList["2007"] = nil   // -1
	baobeiList["1108"] = nil   // -1
	baobeiList["2208"] = nil   // -3
	baobeiList["2601"] = nil   // 2
	baobeiList["1606"] = nil   // -1
	baobeiList["1201"] = nil   // -1
	baobeiList["1104"] = nil   // -1
	baobeiList["1204"] = nil   // -1
	baobeiList["501"] = nil    //-1
	baobeiList["2301"] = nil   // -1

	notHomeList = make(map[string]interface{})
	// 确定的不在家
	notHomeList["504"] = 0
	notHomeList["703"] = 0
	notHomeList["1107"] = 0
	notHomeList["1907"] = 0
	notHomeList["2002"] = 0
	notHomeList["2003"] = 0
	notHomeList["2602"] = 0
	notHomeList["2604"] = 0
	notHomeList["2403"] = 0
	notHomeList["1401"] = 0
	notHomeList["1602"] = 0
}

func parseWechatRecord(m map[string]interface{}) {
	b, err := os.ReadFile("wechatRecord.txt")
	if err != nil {
		panic(err)
	}
	bTxt := string(b)
	rows := strings.Split(bTxt, "\n")
	for i, row := range rows {
		splits := strings.Split(row, ". ")
		if len(splits) != 2 {
			fmt.Println("解析内容=>", i, row, "<=失败")
			continue
		}
		row = splits[1]
		row = strings.TrimSpace(row)

		// 解析头上的数字
		for j := range row {
			if row[j] >= '0' && row[j] <= '9' {
				continue
			} else {
				if j == 0 {
					fmt.Println("解析内容=>", i, row, "<=失败")
				}
				row = row[:j]
				break
			}
		}
		m[row] = nil
	}
}
