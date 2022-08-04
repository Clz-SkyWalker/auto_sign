package service

import (
	"fmt"
	"testing"
)

func TestJueJinProcess(t *testing.T) {
	result := NewJueJinSign([]string{"_ga=GA1.2.836014813.1651022651; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227091088229160650273%2522%252C%2522user_unique_id%2522%253A%25227091088229160650273%2522%252C%2522timestamp%2522%253A1651022651213%257D; n_mh=5H3rzt08HLFhBkeSpwdBThFcdeV90_GNkNcSp76qsrU; sid_guard=f78e6a03a743587dab720ce586619c58%7C1652085905%7C31535999%7CTue%2C+09-May-2023+08%3A45%3A04+GMT; uid_tt=dafb1d9673b5e6b2563abaf59075d250; uid_tt_ss=dafb1d9673b5e6b2563abaf59075d250; sid_tt=f78e6a03a743587dab720ce586619c58; sessionid=f78e6a03a743587dab720ce586619c58; sessionid_ss=f78e6a03a743587dab720ce586619c58; sid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; ssid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; _tea_utm_cache_2608={%22utm_source%22:%22xtxx0725%22%2C%22utm_medium%22:%22push%22%2C%22utm_campaign%22:%22202207vip%22}; _gid=GA1.2.589179463.1659315767; MONITOR_WEB_ID=30d31f9d-f8b7-434e-977b-5a62ed65b294"})
	result.Process()
	fmt.Println(result.ResultInfo)
}

// 测试获取名字
func TestJueJinGetName(t *testing.T) {
	result := NewJueJinSign([]string{"_ga=GA1.2.836014813.1651022651; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227091088229160650273%2522%252C%2522user_unique_id%2522%253A%25227091088229160650273%2522%252C%2522timestamp%2522%253A1651022651213%257D; n_mh=5H3rzt08HLFhBkeSpwdBThFcdeV90_GNkNcSp76qsrU; sid_guard=f78e6a03a743587dab720ce586619c58%7C1652085905%7C31535999%7CTue%2C+09-May-2023+08%3A45%3A04+GMT; uid_tt=dafb1d9673b5e6b2563abaf59075d250; uid_tt_ss=dafb1d9673b5e6b2563abaf59075d250; sid_tt=f78e6a03a743587dab720ce586619c58; sessionid=f78e6a03a743587dab720ce586619c58; sessionid_ss=f78e6a03a743587dab720ce586619c58; sid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; ssid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; _tea_utm_cache_2608={%22utm_source%22:%22xtxx0725%22%2C%22utm_medium%22:%22push%22%2C%22utm_campaign%22:%22202207vip%22}; _gid=GA1.2.589179463.1659315767; MONITOR_WEB_ID=30d31f9d-f8b7-434e-977b-5a62ed65b294"})
	for _, item := range result.signList {
		result.getName(item)
		if result.request.Err != nil || item.err!=nil {
			fmt.Println(result.request.Err)
			return
		}
	}
}

// 测试检测签到状态
func TestJueJinCheckSign(t *testing.T) {
	result := NewJueJinSign([]string{"_ga=GA1.2.836014813.1651022651; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227091088229160650273%2522%252C%2522user_unique_id%2522%253A%25227091088229160650273%2522%252C%2522timestamp%2522%253A1651022651213%257D; n_mh=5H3rzt08HLFhBkeSpwdBThFcdeV90_GNkNcSp76qsrU; sid_guard=f78e6a03a743587dab720ce586619c58%7C1652085905%7C31535999%7CTue%2C+09-May-2023+08%3A45%3A04+GMT; uid_tt=dafb1d9673b5e6b2563abaf59075d250; uid_tt_ss=dafb1d9673b5e6b2563abaf59075d250; sid_tt=f78e6a03a743587dab720ce586619c58; sessionid=f78e6a03a743587dab720ce586619c58; sessionid_ss=f78e6a03a743587dab720ce586619c58; sid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; ssid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; _tea_utm_cache_2608={%22utm_source%22:%22xtxx0725%22%2C%22utm_medium%22:%22push%22%2C%22utm_campaign%22:%22202207vip%22}; _gid=GA1.2.589179463.1659315767; MONITOR_WEB_ID=30d31f9d-f8b7-434e-977b-5a62ed65b294"})
	for _, item := range result.signList {
		result.checkSignStatus(item)
		if result.request.Err != nil {
			fmt.Println(result.request.Err)
			return
		}
	}
}

// 检测免费抽奖次数
func TestJueJinCheckFreeLuckyDraw(t *testing.T) {
	result := NewJueJinSign([]string{"_ga=GA1.2.836014813.1651022651; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227091088229160650273%2522%252C%2522user_unique_id%2522%253A%25227091088229160650273%2522%252C%2522timestamp%2522%253A1651022651213%257D; n_mh=5H3rzt08HLFhBkeSpwdBThFcdeV90_GNkNcSp76qsrU; sid_guard=f78e6a03a743587dab720ce586619c58%7C1652085905%7C31535999%7CTue%2C+09-May-2023+08%3A45%3A04+GMT; uid_tt=dafb1d9673b5e6b2563abaf59075d250; uid_tt_ss=dafb1d9673b5e6b2563abaf59075d250; sid_tt=f78e6a03a743587dab720ce586619c58; sessionid=f78e6a03a743587dab720ce586619c58; sessionid_ss=f78e6a03a743587dab720ce586619c58; sid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; ssid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; _tea_utm_cache_2608={%22utm_source%22:%22xtxx0725%22%2C%22utm_medium%22:%22push%22%2C%22utm_campaign%22:%22202207vip%22}; _gid=GA1.2.589179463.1659315767; MONITOR_WEB_ID=30d31f9d-f8b7-434e-977b-5a62ed65b294"})
	for _, item := range result.signList {
		result.checkLuckyDraw(item)
		if result.request.Err != nil {
			fmt.Println(result.request.Err)
			return
		}
	}
}

// 获取矿石总数
func TestJueJinGetPoint(t *testing.T) {
	result := NewJueJinSign([]string{"_ga=GA1.2.836014813.1651022651; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227091088229160650273%2522%252C%2522user_unique_id%2522%253A%25227091088229160650273%2522%252C%2522timestamp%2522%253A1651022651213%257D; n_mh=5H3rzt08HLFhBkeSpwdBThFcdeV90_GNkNcSp76qsrU; sid_guard=f78e6a03a743587dab720ce586619c58%7C1652085905%7C31535999%7CTue%2C+09-May-2023+08%3A45%3A04+GMT; uid_tt=dafb1d9673b5e6b2563abaf59075d250; uid_tt_ss=dafb1d9673b5e6b2563abaf59075d250; sid_tt=f78e6a03a743587dab720ce586619c58; sessionid=f78e6a03a743587dab720ce586619c58; sessionid_ss=f78e6a03a743587dab720ce586619c58; sid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; ssid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; _tea_utm_cache_2608={%22utm_source%22:%22xtxx0725%22%2C%22utm_medium%22:%22push%22%2C%22utm_campaign%22:%22202207vip%22}; _gid=GA1.2.589179463.1659315767; MONITOR_WEB_ID=30d31f9d-f8b7-434e-977b-5a62ed65b294"})
	for _, item := range result.signList {
		result.getTotalPoint(item)
		if result.request.Err != nil {
			fmt.Println(result.request.Err)
			return
		}
	}
}

// 获取签到天数
func TestJueJinSignDay(t *testing.T) {
	result := NewJueJinSign([]string{"_ga=GA1.2.836014813.1651022651; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227091088229160650273%2522%252C%2522user_unique_id%2522%253A%25227091088229160650273%2522%252C%2522timestamp%2522%253A1651022651213%257D; n_mh=5H3rzt08HLFhBkeSpwdBThFcdeV90_GNkNcSp76qsrU; sid_guard=f78e6a03a743587dab720ce586619c58%7C1652085905%7C31535999%7CTue%2C+09-May-2023+08%3A45%3A04+GMT; uid_tt=dafb1d9673b5e6b2563abaf59075d250; uid_tt_ss=dafb1d9673b5e6b2563abaf59075d250; sid_tt=f78e6a03a743587dab720ce586619c58; sessionid=f78e6a03a743587dab720ce586619c58; sessionid_ss=f78e6a03a743587dab720ce586619c58; sid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; ssid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; _tea_utm_cache_2608={%22utm_source%22:%22xtxx0725%22%2C%22utm_medium%22:%22push%22%2C%22utm_campaign%22:%22202207vip%22}; _gid=GA1.2.589179463.1659315767; MONITOR_WEB_ID=30d31f9d-f8b7-434e-977b-5a62ed65b294"})
	for _, item := range result.signList {
		result.getTotalSignDay(item)
		if result.request.Err != nil {
			fmt.Println(result.request.Err)
			return
		}
	}
}

// 签到
func TestJueJinSign(t *testing.T) {
	result := NewJueJinSign([]string{"_ga=GA1.2.836014813.1651022651; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227091088229160650273%2522%252C%2522user_unique_id%2522%253A%25227091088229160650273%2522%252C%2522timestamp%2522%253A1651022651213%257D; n_mh=5H3rzt08HLFhBkeSpwdBThFcdeV90_GNkNcSp76qsrU; sid_guard=f78e6a03a743587dab720ce586619c58%7C1652085905%7C31535999%7CTue%2C+09-May-2023+08%3A45%3A04+GMT; uid_tt=dafb1d9673b5e6b2563abaf59075d250; uid_tt_ss=dafb1d9673b5e6b2563abaf59075d250; sid_tt=f78e6a03a743587dab720ce586619c58; sessionid=f78e6a03a743587dab720ce586619c58; sessionid_ss=f78e6a03a743587dab720ce586619c58; sid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; ssid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; _tea_utm_cache_2608={%22utm_source%22:%22xtxx0725%22%2C%22utm_medium%22:%22push%22%2C%22utm_campaign%22:%22202207vip%22}; _gid=GA1.2.589179463.1659315767; MONITOR_WEB_ID=30d31f9d-f8b7-434e-977b-5a62ed65b294"})
	for _, item := range result.signList {
		result.sign(item)
		if result.request.Err != nil {
			fmt.Println(result.request.Err)
			return
		}
	}
}

// 测试抽奖
func TestLuckyDraw(t *testing.T) {
	result := NewJueJinSign([]string{"_ga=GA1.2.836014813.1651022651; __tea_cookie_tokens_2608=%257B%2522web_id%2522%253A%25227091088229160650273%2522%252C%2522user_unique_id%2522%253A%25227091088229160650273%2522%252C%2522timestamp%2522%253A1651022651213%257D; n_mh=5H3rzt08HLFhBkeSpwdBThFcdeV90_GNkNcSp76qsrU; sid_guard=f78e6a03a743587dab720ce586619c58%7C1652085905%7C31535999%7CTue%2C+09-May-2023+08%3A45%3A04+GMT; uid_tt=dafb1d9673b5e6b2563abaf59075d250; uid_tt_ss=dafb1d9673b5e6b2563abaf59075d250; sid_tt=f78e6a03a743587dab720ce586619c58; sessionid=f78e6a03a743587dab720ce586619c58; sessionid_ss=f78e6a03a743587dab720ce586619c58; sid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; ssid_ucp_v1=1.0.0-KGI3MmM1OTI4MGU2ODkwMDNiYWFmMmM5MzQ5ZDIwNmM5ZWRhNzQyYWEKFgitwpDA_fW9AhCRqeOTBhiwFDgIQDgaAmxmIiBmNzhlNmEwM2E3NDM1ODdkYWI3MjBjZTU4NjYxOWM1OA; _tea_utm_cache_2608={%22utm_source%22:%22xtxx0725%22%2C%22utm_medium%22:%22push%22%2C%22utm_campaign%22:%22202207vip%22}; _gid=GA1.2.589179463.1659315767; MONITOR_WEB_ID=30d31f9d-f8b7-434e-977b-5a62ed65b294"})
	for _, item := range result.signList {
		result.luckyDraw(item)
		if result.request.Err != nil {
			fmt.Println(result.request.Err)
			return
		}
	}
}
