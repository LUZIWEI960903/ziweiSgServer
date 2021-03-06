package constant

const (
	ProxyConnectError      = -4 + iota //代理连接失败
	ProxyNotInConnect                  //代理错误
	UserNotInConnect2                  //链接没有找到用户
	RoleNotInConnect1                  //链接没有找到角色
	OK                                 //ok
	InvalidParam                       //参数有误
	DBError                            //数据库异常
	UserExist                          //用户已存在
	PwdIncorrect                       //密码不正确
	UserNotExist                       //用户不存在
	SessionInvalid                     //session无效
	HardwareIncorrect                  //Hardware错误
	RoleAlreadyCreate                  //已经创建过角色了
	RoleNotExist                       //角色不存在
	CityNotExist                       //城市不存在
	CityNotMe                          //城市不是自己的
	UpError                            //升级失败
	GeneralNotFound                    //武将不存在
	GeneralNotMe                       //武将不是自己的
	ArmyNotFound                       //军队不存在
	ArmyNotMe                          //军队不是自己的
	ResNotEnough                       //资源不足
	OutArmyLimit                       //超过带兵限制
	ArmyBusy                           //军队再忙
	GeneralBusy                        //将领再忙
	CannotGiveUp                       //不能放弃
	BuildNotMe                         //领地不是自己的
	ArmyNotMain                        //军队没有主将
	UnReachable                        //不可到达
	PhysicalPowerNotEnough             //体力不足
	DecreeNotEnough                    //政令不足
	GoldNotEnough                      //金币不足
	GeneralRepeat                      //重复上阵
	CostNotEnough                      //cost不足
	GeneralNoHas                       //没有该合成武将
	GeneralNoSame                      //合成武将非同名
	ArmyNotEnough                      //队伍数不足
	TongShuaiNotEnough                 //统帅不足
	GeneralStarMax                     //升级到最大星级
	UnionCreateError                   //联盟创建失败
	UnionNotFound                      //联盟不存在
	PermissionDenied                   //权限不足
	UnionAlreadyHas                    //已经有联盟
	UnionNotAllowExit                  //不允许退出
	ContentTooLong                     //内容太长
	NotBelongUnion                     //不属于该联盟
	PeopleIsFull                       //用户已满
	HasApply                           //已经申请过了
	BuildCanNotDefend                  //不能驻守
	BuildCanNotAttack                  //不能占领
	BuildMBSNotFound                   //没有军营
	BuildWarFree                       //免战中
	ArmyConscript                      //征兵中
	BuildGiveUpAlready                 //领地已经在放弃了
	CanNotBuildNew                     //不能再新建建筑在领地上
	CanNotTransfer                     //不能调兵
	HoldIsFull                         //坑位已满
	ArmyIsOutside                      //队伍在城外
	CanNotUpBuild                      //不能升级建筑
	CanNotDestroy                      //不能拆除建筑
	OutCollectTimesLimit               //超过征收次数
	InCdCanNotOperate                  //cd内不能操作
	OutGeneralLimit                    //武将超过上限了
	NotHasJiShi                        //没有集市
	OutPosTagLimit                     //超过了收藏上限
	OutSkillLimit                      //超过了技能上限
	UpSkillError                       //装备技能失败
	DownSkillError                     //取下技能失败
	OutArmNotMatch                     //兵种不符
	PosNotSkill                        //该位置没有技能
	SkillLevelFull                     //技能等级已满
)
