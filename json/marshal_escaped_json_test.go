package json

import "testing"

func TestParseToCustomerTrendMsg(t *testing.T) {
	parseToCustomerTrendMsg(str2)

}

// 1、 获取当前租户下的所有带 union_id 的外部联系人
// 2、 获取当前外部联系人的云店小程序留电手机号
// 3、 获取当前外部联系人所有的跟进成员信息（从企微上获取）
// 4、 通过是标识 add_way = 2  来判断是否是从手机号搜索来的，如果是，则将第一个手机号当做搜索手机号保存
// 5、 保存手机号，同时上报信自至客户中心
