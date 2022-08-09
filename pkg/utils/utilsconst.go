package utils

const (
	Key         = "[key]"
	Title       = "title"
	Desp        = "desp"
	ServerJBase = "https://sctapi.ftqq.com/" + Key + ".send"

	// 掘金
	JUEJIN_API           = "https://api.juejin.cn"
	JUEJIN_GetName       = JUEJIN_API + "/user_api/v1/user/get"             // 获取昵称
	JUEJIN_CheckSign     = JUEJIN_API + "/growth_api/v1/get_today_status"   // 检测今天是否签到
	JUEJIN_Sign          = JUEJIN_API + "/growth_api/v1/check_in"           // 签到
	JUEJIN_Draw          = JUEJIN_API + "/growth_api/v1/lottery/draw"       // 抽奖
	JUEJIN_CheckFreeDraw = JUEJIN_API + "/growth_api/v1/lottery_config/get" // 获取免费抽奖次数
	JUEJIN_Total         = JUEJIN_API + "/growth_api/v1/get_cur_point"      // 获取矿石数目
	JUEJIN_TotalSignDay  = JUEJIN_API + "/growth_api/v1/get_counts"         // 统计签到天数

	// ireader
	IREADER_API  = "http://ah2.zhangyue.com/zyam/app/app.php"
	IREADER_Seed = IREADER_API + "?ca=Sign.Seed&pca=Sign.Index&usr="
	IREADER_Sign = IREADER_API + "?pk=BEqiandao&rgt=7&from=1&type=0&usr="
)
