package json

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var str = "\"[{\\\"appid\\\":\\\"5032\\\",\\\"appn\\\":\\\"aicard\\\",\\\"stuid\\\":\\\"1534714427629133824\\\",\\\"tn\\\":\\\"gzminjieadmin_test\\\",\\\"at\\\":\\\"微信小程序\\\",\\\"osn\\\":\\\"Android\\\",\\\"osv\\\":\\\"12\\\",\\\"wv\\\":\\\"8.0.28\\\",\\\"dmf\\\":\\\"2206123SC\\\",\\\"nt\\\":\\\"4G\\\",\\\"pg\\\":\\\"page/mainPage/pages/allHouseList/main\\\",\\\"ps\\\":\\\"\\\",\\\"pgt\\\":\\\"首页\\\",\\\"r\\\":\\\"\\\",\\\"pgid\\\":1,\\\"sid\\\":\\\"1gin3v6e-0fb5ec35-265c4d6f\\\",\\\"eid\\\":\\\"1gin3v63-0a816986-265c4d6f\\\",\\\"t\\\":\\\"2022/11/25 18:05:14\\\",\\\"at_appid\\\":\\\"5032\\\",\\\"at_appn\\\":\\\"aicard\\\",\\\"u\\\":\\\"1585560930459361280\\\",\\\"unid\\\":\\\"olgEc028MCblhG3EJ6pEyHM9rcxA\\\",\\\"oid\\\":\\\"osgwc5DGJmaaNQd1PywFaR4paeCE\\\",\\\"u_type\\\":\\\"5\\\",\\\"at_u\\\":\\\"1585560930459361280\\\",\\\"long\\\":\\\"\\\",\\\"lat\\\":\\\"\\\",\\\"scid\\\":1089,\\\"plt\\\":390,\\\"from_event_id\\\":\\\"\\\",\\\"f_content_type\\\":\\\"集团首页\\\",\\\"f_content_name\\\":\\\"集团首页\\\",\\\"is_carrier\\\":0,\\\"is_first_ss\\\":1,\\\"ecid\\\":50321041,\\\"iskey\\\":\\\"1\\\",\\\"st\\\":0,\\\"pt\\\":\\\"allHouseList\\\",\\\"url\\\":\\\"page/mainPage/pages/allHouseList/main\\\",\\\"ebn\\\":\\\"打开小程序\\\",\\\"e\\\":\\\"打开小程序\\\",\\\"tc_id\\\":\\\"1596082574474653696\\\",\\\"etp\\\":\\\"浏览\\\",\\\"l\\\":{\\\"st\\\":0,\\\"channel_id\\\":\\\"0\\\",\\\"mid\\\":\\\"wxb988caa4c33af7b6\\\",\\\"scene_id\\\":\\\"1089\\\",\\\"qk_share_id\\\":\\\"\\\",\\\"from_share_hash\\\":\\\"\\\",\\\"from_user_id\\\":\\\"0\\\",\\\"from_seller_id\\\":\\\"0\\\",\\\"from_seller_type\\\":\\\"0\\\",\\\"user_name\\\":\\\"朱老师\\\",\\\"user_phone\\\":\\\"\\\",\\\"user_gender\\\":\\\"3\\\",\\\"ai_city_id\\\":\\\"\\\",\\\"area_id\\\":\\\"0\\\",\\\"area_name\\\":\\\"全国\\\",\\\"scene\\\":\\\"1089\\\",\\\"referrer_info\\\":\\\"{}\\\",\\\"pg\\\":\\\"page/mainPage/pages/allHouseList/main\\\"}},{\\\"appid\\\":\\\"5032\\\",\\\"appn\\\":\\\"aicard\\\",\\\"stuid\\\":\\\"1534714427629133824\\\",\\\"tn\\\":\\\"gzminjieadmin_test\\\",\\\"at\\\":\\\"微信小程序\\\",\\\"osn\\\":\\\"Android\\\",\\\"osv\\\":\\\"12\\\",\\\"wv\\\":\\\"8.0.28\\\",\\\"dmf\\\":\\\"2206123SC\\\",\\\"nt\\\":\\\"4G\\\",\\\"pg\\\":\\\"page/mainPage/pages/allHouseList/main\\\",\\\"ps\\\":\\\"\\\",\\\"pgt\\\":\\\"首页\\\",\\\"r\\\":\\\"\\\",\\\"pgid\\\":1,\\\"sid\\\":\\\"1gin3v7c-05a108e6-265c4d6f\\\",\\\"eid\\\":\\\"1gin3v7b-020ba2f7-265c4d6f\\\",\\\"t\\\":\\\"2022/11/25 18:05:15\\\",\\\"at_appid\\\":\\\"5032\\\",\\\"at_appn\\\":\\\"aicard\\\",\\\"u\\\":\\\"1585560930459361280\\\",\\\"unid\\\":\\\"olgEc028MCblhG3EJ6pEyHM9rcxA\\\",\\\"oid\\\":\\\"osgwc5DGJmaaNQd1PywFaR4paeCE\\\",\\\"u_type\\\":\\\"5\\\",\\\"at_u\\\":\\\"1585560930459361280\\\",\\\"long\\\":\\\"113.95206353081598\\\",\\\"lat\\\":\\\"22.527163628472223\\\",\\\"scid\\\":1089,\\\"plt\\\":390,\\\"from_event_id\\\":\\\"1gin3v7b-0891c2b3-265c4d6f\\\",\\\"ecid\\\":1000,\\\"iskey\\\":\\\"1\\\",\\\"st\\\":0,\\\"pt\\\":\\\"allHouseList\\\",\\\"url\\\":\\\"page/mainPage/pages/allHouseList/main\\\",\\\"ebn\\\":\\\"授权地理位置\\\",\\\"e\\\":\\\"完成授权地理位置：深圳市·南山区·高新南十道\\\",\\\"tc_id\\\":\\\"1596082574474653696\\\",\\\"etp\\\":\\\"授权\\\",\\\"l\\\":{\\\"st\\\":0,\\\"channel_id\\\":\\\"0\\\",\\\"mid\\\":\\\"wxb988caa4c33af7b6\\\",\\\"scene_id\\\":\\\"1089\\\",\\\"qk_share_id\\\":\\\"\\\",\\\"from_share_hash\\\":\\\"\\\",\\\"from_user_id\\\":\\\"0\\\",\\\"from_seller_id\\\":\\\"0\\\",\\\"from_seller_type\\\":\\\"0\\\",\\\"user_name\\\":\\\"朱老师\\\",\\\"user_phone\\\":\\\"\\\",\\\"user_gender\\\":\\\"3\\\",\\\"ai_city_id\\\":\\\"4fd71566-904b-11e4-a2d4-d89d672b5244\\\",\\\"area_id\\\":\\\"0\\\",\\\"area_name\\\":\\\"全国\\\",\\\"wx_auth_status\\\":1,\\\"lnt_lat_location\\\":\\\"中国,广东省,深圳市,南山区,高新南十道,\\\",\\\"pg\\\":\\\"page/mainPage/pages/allHouseList/main\\\"}},{\\\"appid\\\":\\\"5032\\\",\\\"appn\\\":\\\"aicard\\\",\\\"stuid\\\":\\\"1534714427629133824\\\",\\\"tn\\\":\\\"gzminjieadmin_test\\\",\\\"at\\\":\\\"微信小程序\\\",\\\"osn\\\":\\\"Android\\\",\\\"osv\\\":\\\"12\\\",\\\"wv\\\":\\\"8.0.28\\\",\\\"dmf\\\":\\\"2206123SC\\\",\\\"nt\\\":\\\"4G\\\",\\\"pg\\\":\\\"page/mainPage/pages/allHouseList/main\\\",\\\"ps\\\":\\\"\\\",\\\"pgt\\\":\\\"首页\\\",\\\"r\\\":\\\"\\\",\\\"pgid\\\":1,\\\"sid\\\":\\\"1gin3v7c-01a27f4a-265c4d6f\\\",\\\"eid\\\":\\\"1gin3v7b-0891c2b3-265c4d6f\\\",\\\"t\\\":\\\"2022/11/25 18:05:15\\\",\\\"at_appid\\\":\\\"5032\\\",\\\"at_appn\\\":\\\"aicard\\\",\\\"u\\\":\\\"1585560930459361280\\\",\\\"unid\\\":\\\"olgEc028MCblhG3EJ6pEyHM9rcxA\\\",\\\"oid\\\":\\\"osgwc5DGJmaaNQd1PywFaR4paeCE\\\",\\\"u_type\\\":\\\"5\\\",\\\"at_u\\\":\\\"1585560930459361280\\\",\\\"long\\\":\\\"113.95206353081598\\\",\\\"lat\\\":\\\"22.527163628472223\\\",\\\"scid\\\":1089,\\\"plt\\\":390,\\\"from_event_id\\\":\\\"\\\",\\\"ecid\\\":1001,\\\"iskey\\\":\\\"1\\\",\\\"st\\\":0,\\\"pt\\\":\\\"allHouseList\\\",\\\"url\\\":\\\"page/mainPage/pages/allHouseList/main\\\",\\\"ebn\\\":\\\"上报地理位置（静默）\\\",\\\"e\\\":\\\"完成授权地理位置：深圳市·南山区·高新南十道\\\",\\\"tc_id\\\":\\\"1596082574474653696\\\",\\\"etp\\\":\\\"授权\\\",\\\"l\\\":{\\\"st\\\":0,\\\"channel_id\\\":\\\"0\\\",\\\"mid\\\":\\\"wxb988caa4c33af7b6\\\",\\\"scene_id\\\":\\\"1089\\\",\\\"qk_share_id\\\":\\\"\\\",\\\"from_share_hash\\\":\\\"\\\",\\\"from_user_id\\\":\\\"0\\\",\\\"from_seller_id\\\":\\\"0\\\",\\\"from_seller_type\\\":\\\"0\\\",\\\"user_name\\\":\\\"朱老师\\\",\\\"user_phone\\\":\\\"\\\",\\\"user_gender\\\":\\\"3\\\",\\\"ai_city_id\\\":\\\"4fd71566-904b-11e4-a2d4-d89d672b5244\\\",\\\"area_id\\\":\\\"0\\\",\\\"area_name\\\":\\\"全国\\\",\\\"wx_auth_status\\\":1,\\\"lnt_lat_location\\\":\\\"中国,广东省,深圳市,南山区,高新南十道,\\\",\\\"pg\\\":\\\"page/mainPage/pages/allHouseList/main\\\"}}]\""

type CustomerTrendMsg struct {
	Ecid  int64  `json:"ecid"`
	Tn    string `json:"tn"`
	OneId string `json:"one_id"`
	Stuid string `json:"stuid"`
	E     string `json:"e"`
}

var str2 = "{a:123,b:321}"

func parseToCustomerTrendMsg(v interface{}) []*CustomerTrendMsg {
	unquoted, err := strconv.Unquote(str)
	if err != nil {
		panic(err)
	}
	// result, err := json.Marshal(v)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s", result)
	// // if err := json.Unmarshal([]byte(str), res); err != nil {
	// // 	panic(err)
	// // }
	// // fmt.Printf("%+v", res)
	// return nil
	var result []*CustomerTrendMsg
	if err := json.Unmarshal([]byte(unquoted), &result); err != nil {
		panic(err)
	}
	rec, _ := json.Marshal(result)
	fmt.Printf("%s", rec)
	return result

}
