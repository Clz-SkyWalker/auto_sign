package utils

const (
	JUEJIN_API           = "https://api.juejin.cn"
	JUEJIN_GetName       = JUEJIN_API + "/user_api/v1/user/get"             // 获取昵称
	JUEJIN_CheckSign     = JUEJIN_API + "/growth_api/v1/get_today_status"   // 检测今天是否签到
	JUEJIN_Sign          = JUEJIN_API + "/growth_api/v1/check_in"           // 签到
	JUEJIN_Draw          = JUEJIN_API + "/growth_api/v1/lottery/draw"       // 抽奖
	JUEJIN_CheckFreeDraw = JUEJIN_API + "/growth_api/v1/lottery_config/get" // 获取免费抽奖次数
	JUEJIN_Total         = JUEJIN_API + "/growth_api/v1/get_cur_point"      // 获取矿石数目
	JUEJIN_TotalSignDay  = JUEJIN_API + "/growth_api/v1/get_counts"         // 统计签到天数

	Key         = "[key]"
	Title       = "title"
	Desp        = "desp"
	ServerJBase = "https://sctapi.ftqq.com/" + Key + ".send"
)
