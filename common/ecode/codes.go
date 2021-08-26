package ecode

//定义内部错误码
//小于100000的为内部错误，内部错误会响应http 500
var (
	// ErrCodexxxx 腾讯错误码
	ErrCode6000    = add(6000)    // 数据版本冲突
	ErrCode40001   = add(40001)   // 不合法的secret参数
	ErrCode40003   = add(40003)   // 无效的UserID
	ErrCode40004   = add(40004)   // 不合法的媒体文件类型
	ErrCode40005   = add(40005)   // 不合法的type参数
	ErrCode40006   = add(40006)   // 不合法的文件大小
	ErrCode40007   = add(40007)   // 不合法的media_id参数
	ErrCode40008   = add(40008)   // 不合法的msgtype参数
	ErrCode40009   = add(40009)   // 上传图片大小不是有效值
	ErrCode40011   = add(40011)   // 上传视频大小不是有效值
	ErrCode40013   = add(40013)   // 不合法的CorpID
	ErrCode40014   = add(40014)   // 不合法的access_tokDetail
	ErrCode40016   = add(40016)   // 不合法的按钮个数
	ErrCode40017   = add(40017)   // 不合法的按钮类型
	ErrCode40018   = add(40018)   // 不合法的按钮名字长度
	ErrCode40019   = add(40019)   // 不合法的按钮KEY长度
	ErrCode40020   = add(40020)   // 不合法的按钮URL长度
	ErrCode40022   = add(40022)   // 不合法的子菜单级数
	ErrCode40023   = add(40023)   // 不合法的子菜单按钮个数
	ErrCode40024   = add(40024)   // 不合法的子菜单按钮类型
	ErrCode40025   = add(40025)   // 不合法的子菜单按钮名字长度
	ErrCode40026   = add(40026)   // 不合法的子菜单按钮KEY长度
	ErrCode40027   = add(40027)   // 不合法的子菜单按钮URL长度
	ErrCode40029   = add(40029)   // 不合法的oauth_code
	ErrCode40031   = add(40031)   // 不合法的UserID列表
	ErrCode40032   = add(40032)   // 不合法的UserID列表长度
	ErrCode40033   = add(40033)   // 不合法的请求字符
	ErrCode40035   = add(40035)   // 不合法的参数
	ErrCode40036   = add(40036)   // 不合法的模板id长度
	ErrCode40037   = add(40037)   // 无效的模板id
	ErrCode40039   = add(40039)   // 不合法的url长度
	ErrCode40050   = add(40050)   // chatid不存在
	ErrCode40054   = add(40054)   // 不合法的子菜单url域名
	ErrCode40055   = add(40055)   // 不合法的菜单url域名
	ErrCode40056   = add(40056)   // 不合法的agDetailtid
	ErrCode40057   = add(40057)   // 不合法的callbackurl或者callbackurl验证失败
	ErrCode40058   = add(40058)   // 不合法的参数
	ErrCode40059   = add(40059)   // 不合法的上报地理位置标志位
	ErrCode40063   = add(40063)   // 参数为空
	ErrCode40066   = add(40066)   // 不合法的部门列表
	ErrCode40068   = add(40068)   // 不合法的标签ID
	ErrCode40070   = add(40070)   // 指定的标签范围结点全部无效
	ErrCode40071   = add(40071)   // 不合法的标签名字
	ErrCode40072   = add(40072)   // 不合法的标签名字长度
	ErrCode40073   = add(40073)   // 不合法的opDetailid
	ErrCode40074   = add(40074)   // news消息不支持保密消息类型
	ErrCode40077   = add(40077)   // 不合法的pre_auth_code参数
	ErrCode40078   = add(40078)   // 不合法的auth_code参数
	ErrCode40080   = add(40080)   // 不合法的suite_secret
	ErrCode40082   = add(40082)   // 不合法的suite_tokDetail
	ErrCode40083   = add(40083)   // 不合法的suite_id
	ErrCode40084   = add(40084)   // 不合法的permanDetailt_code参数
	ErrCode40085   = add(40085)   // 不合法的的suite_ticket参数
	ErrCode40086   = add(40086)   // 不合法的第三方应用appid
	ErrCode40088   = add(40088)   // jobid不存在
	ErrCode40089   = add(40089)   // 批量任务的结果已清理
	ErrCode40091   = add(40091)   // secret不合法
	ErrCode40092   = add(40092)   // 导入文件存在不合法的内容
	ErrCode40093   = add(40093)   // jsapi签名错误
	ErrCode40094   = add(40094)   // 不合法的URL
	ErrCode40096   = add(40096)   // 不合法的外部联系人userid
	ErrCode40097   = add(40097)   // 该成员尚未离职
	ErrCode40098   = add(40098)   // 成员尚未实名认证
	ErrCode40099   = add(40099)   // 外部联系人的数量已达上限
	ErrCode40100   = add(40100)   // 此用户的外部联系人已经在转移流程中
	ErrCode40102   = add(40102)   // 域名或IP不可与应用市场上架应用重复
	ErrCode40123   = add(40123)   // 上传临时图片素材，图片格式非法
	ErrCode40124   = add(40124)   // 推广活动里的sn禁止绑定
	ErrCode40125   = add(40125)   // 无效的opDetailuserid参数
	ErrCode40126   = add(40126)   // 企业标签个数达到上限，最多为3000个
	ErrCode40127   = add(40127)   // 不支持的uri schema
	ErrCode40128   = add(40128)   // 客户转接过于频繁（90天内只允许转接一次，同一个客户最多只能转接两次）
	ErrCode40129   = add(40129)   // 当前客户正在转接中
	ErrCode40130   = add(40130)   // 原跟进人与接手人一样，不可继承
	ErrCode40131   = add(40131)   // handover_userid 并不是外部联系人的跟进人
	ErrCode41001   = add(41001)   // 缺少access_tokDetail参数
	ErrCode41002   = add(41002)   // 缺少corpid参数
	ErrCode41004   = add(41004)   // 缺少secret参数
	ErrCode41006   = add(41006)   // 缺少media_id参数
	ErrCode41008   = add(41008)   // 缺少auth code参数
	ErrCode41009   = add(41009)   // 缺少userid参数
	ErrCode41010   = add(41010)   // 缺少url参数
	ErrCode41011   = add(41011)   // 缺少agDetailtid参数
	ErrCode41016   = add(41016)   // 缺少title参数
	ErrCode41019   = add(41019)   // 缺少 departmDetailt 参数
	ErrCode41017   = add(41017)   // 缺少tagid参数
	ErrCode41021   = add(41021)   // 缺少suite_id参数
	ErrCode41022   = add(41022)   // 缺少suite_access_tokDetail参数
	ErrCode41023   = add(41023)   // 缺少suite_ticket参数
	ErrCode41024   = add(41024)   // 缺少secret参数
	ErrCode41025   = add(41025)   // 缺少permanDetailt_code参数
	ErrCode41033   = add(41033)   // 缺少 description 参数
	ErrCode41035   = add(41035)   // 缺少外部联系人userid参数
	ErrCode41036   = add(41036)   // 不合法的企业对外简称
	ErrCode41037   = add(41037)   // 缺少「联系我」type参数
	ErrCode41038   = add(41038)   // 缺少「联系我」scDetaile参数
	ErrCode41039   = add(41039)   // 无效的「联系我」type参数
	ErrCode41040   = add(41040)   // 无效的「联系我」scDetaile参数
	ErrCode41041   = add(41041)   // 「联系我」使用人数超过限制
	ErrCode41042   = add(41042)   // 无效的「联系我」style参数
	ErrCode41043   = add(41043)   // 缺少「联系我」config_id参数
	ErrCode41044   = add(41044)   // 无效的「联系我」config_id参数
	ErrCode41045   = add(41045)   // API添加「联系我」达到数量上限
	ErrCode41046   = add(41046)   // 缺少企业群发消息id
	ErrCode41047   = add(41047)   // 无效的企业群发消息id
	ErrCode41048   = add(41048)   // 无可发送的客户
	ErrCode41049   = add(41049)   // 缺少欢迎语code参数
	ErrCode41050   = add(41050)   // 无效的欢迎语code
	ErrCode41051   = add(41051)   // 客户和服务人员已经开始聊天了
	ErrCode41052   = add(41052)   // 无效的发送时间
	ErrCode41053   = add(41053)   // 客户未同意聊天存档
	ErrCode41054   = add(41054)   // 该用户尚未激活
	ErrCode41055   = add(41055)   // 群欢迎语模板数量达到上限
	ErrCode41056   = add(41056)   // 外部联系人id类型不正确
	ErrCode41057   = add(41057)   // 企业或服务商未绑定微信开发者账号
	ErrCode41059   = add(41059)   // 缺少momDetailt_id参数
	ErrCode41060   = add(41060)   // 不合法的momDetailt_id参数
	ErrCode41061   = add(41061)   // 不合法朋友圈发送成员userid，当前朋友圈并非此用户发表
	ErrCode41062   = add(41062)   // 企业创建的朋友圈尚未被成员userid发表
	ErrCode41063   = add(41063)   // 群发消息正在被派发中，请稍后再试
	ErrCode41064   = add(41064)   // 附件大小超过限制
	ErrCode41065   = add(41065)   // 无效的附件类型
	ErrCode41066   = add(41066)   // 用户视频号名称错误
	ErrCode41102   = add(41102)   // 缺少菜单名
	ErrCode42001   = add(42001)   // access_tokDetail已过期
	ErrCode42007   = add(42007)   // pre_auth_code已过期
	ErrCode42009   = add(42009)   // suite_access_tokDetail已过期
	ErrCode42012   = add(42012)   // jsapi_ticket不可用，一般是没有正确调用接口来创建jsapi_ticket
	ErrCode42013   = add(42013)   // 小程序未登陆或登录态已经过期
	ErrCode42014   = add(42014)   // 任务卡片消息的task_id不合法
	ErrCode42015   = add(42015)   // 更新的消息的应用与发送消息的应用不匹配
	ErrCode42016   = add(42016)   // 更新的task_id不存在
	ErrCode42017   = add(42017)   // 按钮key值不存在
	ErrCode42018   = add(42018)   // 按钮key值不合法
	ErrCode42019   = add(42019)   // 缺少按钮key值不合法
	ErrCode42020   = add(42020)   // 缺少按钮名称
	ErrCode42021   = add(42021)   // device_access_tokDetail 过期
	ErrCode42022   = add(42022)   // code已经被使用过。只能使用一次
	ErrCode43004   = add(43004)   // 指定的userid未绑定微信或未关注微工作台（原企业号）
	ErrCode43009   = add(43009)   // 企业未验证主体
	ErrCode43012   = add(43012)   // 应用需配置回调url
	ErrCode44001   = add(44001)   // 多媒体文件为空
	ErrCode44004   = add(44004)   // 文本消息contDetailt参数为空
	ErrCode45001   = add(45001)   // 多媒体文件大小超过限制
	ErrCode45002   = add(45002)   // 消息内容大小超过限制
	ErrCode45004   = add(45004)   // 应用description参数长度不符合系统限制
	ErrCode45007   = add(45007)   // 语音播放时间超过限制
	ErrCode45008   = add(45008)   // 图文消息的文章数量不符合系统限制
	ErrCode45009   = add(45009)   // 接口调用超过限制
	ErrCode45022   = add(45022)   // 应用name参数长度不符合系统限制
	ErrCode45024   = add(45024)   // 帐号数量超过上限
	ErrCode45026   = add(45026)   // 触发删除用户数的保护
	ErrCode45029   = add(45029)   // 回包大小超过上限
	ErrCode45032   = add(45032)   // 图文消息author参数长度超过限制
	ErrCode45033   = add(45033)   // 接口并发调用超过限制
	ErrCode45034   = add(45034)   // url必须有协议头
	ErrCode46003   = add(46003)   // 菜单未设置
	ErrCode46004   = add(46004)   // 指定的用户不存在
	ErrCode48002   = add(48002)   // API接口无权限调用
	ErrCode48003   = add(48003)   // 不合法的suite_id
	ErrCode48004   = add(48004)   // 授权关系无效
	ErrCode48005   = add(48005)   // API接口已废弃
	ErrCode48006   = add(48006)   // 接口权限被收回
	ErrCode49004   = add(49004)   // 签名不匹配
	ErrCode49008   = add(49008)   // 群已经解散
	ErrCode50001   = add(50001)   // redirect_url未登记可信域名
	ErrCode50002   = add(50002)   // 成员不在权限范围
	ErrCode50003   = add(50003)   // 应用已禁用
	ErrCode50100   = add(50100)   // 分页查询的游标无效
	ErrCode60001   = add(60001)   // 部门长度不符合限制
	ErrCode60003   = add(60003)   // 部门ID不存在
	ErrCode60004   = add(60004)   // 父部门不存在
	ErrCode60005   = add(60005)   // 部门下存在成员
	ErrCode60006   = add(60006)   // 部门下存在子部门
	ErrCode60007   = add(60007)   // 不允许删除根部门
	ErrCode60008   = add(60008)   // 部门已存在
	ErrCode60009   = add(60009)   // 部门名称含有非法字符
	ErrCode60010   = add(60010)   // 部门存在循环关系
	ErrCode60011   = add(60011)   // 指定的成员/部门/标签参数无权限
	ErrCode60012   = add(60012)   // 不允许删除默认应用
	ErrCode60020   = add(60020)   // 访问ip不在白名单之中
	ErrCode60021   = add(60021)   // userid不在应用可见范围内
	ErrCode60028   = add(60028)   // 不允许修改第三方应用的主页 URL
	ErrCode60102   = add(60102)   // UserID已存在
	ErrCode60103   = add(60103)   // 手机号码不合法
	ErrCode60104   = add(60104)   // 手机号码已存在
	ErrCode60105   = add(60105)   // 邮箱不合法
	ErrCode60106   = add(60106)   // 邮箱已存在
	ErrCode60107   = add(60107)   // 微信号不合法
	ErrCode60110   = add(60110)   // 用户所属部门数量超过限制
	ErrCode60111   = add(60111)   // UserID不存在
	ErrCode60112   = add(60112)   // 成员name参数不合法
	ErrCode60123   = add(60123)   // 无效的部门id
	ErrCode60124   = add(60124)   // 无效的父部门id
	ErrCode60125   = add(60125)   // 非法部门名字
	ErrCode60127   = add(60127)   // 缺少departmDetailt参数
	ErrCode60129   = add(60129)   // 成员手机和邮箱都为空
	ErrCode60132   = add(60132)   // is_leader_in_dept和departmDetailt的元素个数不一致
	ErrCode60136   = add(60136)   // 记录不存在
	ErrCode60137   = add(60137)   // 家长手机号重复
	ErrCode60203   = add(60203)   // 不合法的模版ID
	ErrCode60204   = add(60204)   // 模版状态不可用
	ErrCode60205   = add(60205)   // 模版关键词不匹配
	ErrCode60206   = add(60206)   // 该种类型的消息只支持第三方独立应用使用
	ErrCode60207   = add(60207)   // 第三方独立应用只允许发送模板消息
	ErrCode60208   = add(60208)   // 第三方独立应用不支持指定@all，不支持参数toparty和totag
	ErrCode60209   = add(60209)   // 缺少操作者vid
	ErrCode60210   = add(60210)   // 选择成员列表为空
	ErrCode60211   = add(60211)   // SelectedTicket为空
	ErrCode60214   = add(60214)   // 仅支持第三方应用调用
	ErrCode60215   = add(60215)   // 传入SelectedTicket数量超过最大限制（10个）
	ErrCode60217   = add(60217)   // 当前操作者无权限，操作者需要授权或者在可见范围内
	ErrCode60218   = add(60218)   // 仅支持成员授权模式的应用可调用
	ErrCode60219   = add(60219)   // 消费SelectedTicket和创建SelectedTicket的应用appid不匹配
	ErrCode60220   = add(60220)   // 缺少corpappid
	ErrCode60221   = add(60221)   // opDetail_userid对应的服务商不是当前服务商
	ErrCode60222   = add(60222)   // 非法SelectedTicket
	ErrCode60223   = add(60223)   // 非法BundleId
	ErrCode60224   = add(60224)   // 非法PackagDetailame
	ErrCode60225   = add(60225)   // 当前操作者并非SelectedTicket相关人，不能创建群聊
	ErrCode60226   = add(60226)   // 选人数量超过最大限制（2000）
	ErrCode60227   = add(60227)   // 缺少ServiceCorpid
	ErrCode65000   = add(65000)   // 学校已经迁移
	ErrCode65001   = add(65001)   // 无效的关注模式
	ErrCode65002   = add(65002)   // 导入家长信息数量过多
	ErrCode65003   = add(65003)   // 学校尚未迁移
	ErrCode65004   = add(65004)   // 组织架构不存在
	ErrCode65005   = add(65005)   // 无效的同步模式
	ErrCode65006   = add(65006)   // 无效的管理员类型
	ErrCode65007   = add(65007)   // 无效的家校部门类型
	ErrCode65008   = add(65008)   // 无效的入学年份
	ErrCode65009   = add(65009)   // 无效的标准年级类型
	ErrCode65010   = add(65010)   // 此userid并不是学生
	ErrCode65011   = add(65011)   // 家长userid数量超过限制
	ErrCode65012   = add(65012)   // 学生userid数量超过限制
	ErrCode65013   = add(65013)   // 学生已有家长
	ErrCode65014   = add(65014)   // 非学校企业
	ErrCode65015   = add(65015)   // 父部门类型不匹配
	ErrCode65018   = add(65018)   // 家长人数达到上限
	ErrCode660001  = add(660001)  // 无效的商户号
	ErrCode660002  = add(660002)  // 无效的企业收款人id
	ErrCode660003  = add(660003)  // userid不在应用的可见范围
	ErrCode660004  = add(660004)  // partyid不在应用的可见范围
	ErrCode660005  = add(660005)  // tagid不在应用的可见范围
	ErrCode660006  = add(660006)  // 找不到该商户号
	ErrCode660007  = add(660007)  // 申请已经存在
	ErrCode660008  = add(660008)  // 商户号已经绑定
	ErrCode660009  = add(660009)  // 商户号主体和商户主体不一致
	ErrCode660010  = add(660010)  // 超过商户号绑定数量限制
	ErrCode660011  = add(660011)  // 商户号未绑定
	ErrCode670001  = add(670001)  // 应用不在共享范围
	ErrCode72023   = add(72023)   // 发票已被其他公众号锁定
	ErrCode72024   = add(72024)   // 发票状态错误
	ErrCode72037   = add(72037)   // 存在发票不属于该用户
	ErrCode80001   = add(80001)   // 可信域名不正确，或者无ICP备案
	ErrCode81001   = add(81001)   // 部门下的结点数超过限制（3W）
	ErrCode81002   = add(81002)   // 部门最多15层
	ErrCode81003   = add(81003)   // 标签下节点个数超过30000个
	ErrCode81011   = add(81011)   // 无权限操作标签
	ErrCode81012   = add(81012)   // 缺失可见范围
	ErrCode81013   = add(81013)   // UserID、部门ID、标签ID全部非法或无权限
	ErrCode81014   = add(81014)   // 标签添加成员，单次添加user或party过多
	ErrCode81015   = add(81015)   // 邮箱域名需要跟企业邮箱域名一致
	ErrCode81016   = add(81016)   // logined_userid字段缺失
	ErrCode81017   = add(81017)   // 请求个数超过限制
	ErrCode81018   = add(81018)   // 该服务商可获取名字数量配额不足
	ErrCode81019   = add(81019)   // items数组成员缺少id字段
	ErrCode81020   = add(81020)   // items数组成员缺少type字段
	ErrCode81021   = add(81021)   // items数组成员的type字段不合法
	ErrCode82001   = add(82001)   // 指定的成员/部门/标签全部为空
	ErrCode82002   = add(82002)   // 不合法的PartyID列表长度
	ErrCode82003   = add(82003)   // 不合法的TagID列表长度
	ErrCode82004   = add(82004)   // 不合法的消息内容
	ErrCode84014   = add(84014)   // 成员票据过期
	ErrCode84015   = add(84015)   // 成员票据无效
	ErrCode84019   = add(84019)   // 缺少templateid参数
	ErrCode84020   = add(84020)   // templateid不存在
	ErrCode84021   = add(84021)   // 缺少register_code参数
	ErrCode84022   = add(84022)   // 无效的register_code参数
	ErrCode84023   = add(84023)   // 不允许调用设置通讯录同步完成接口
	ErrCode84024   = add(84024)   // 无注册信息
	ErrCode84025   = add(84025)   // 不符合的state参数
	ErrCode84052   = add(84052)   // 缺少caller参数
	ErrCode84053   = add(84053)   // 缺少callee参数
	ErrCode84054   = add(84054)   // 缺少auth_corpid参数
	ErrCode84055   = add(84055)   // 超过拨打公费电话频率
	ErrCode84056   = add(84056)   // 被拨打用户安装应用时未授权拨打公费电话权限
	ErrCode84057   = add(84057)   // 公费电话余额不足
	ErrCode84058   = add(84058)   // caller 呼叫号码不支持
	ErrCode84059   = add(84059)   // 号码非法
	ErrCode84060   = add(84060)   // callee 呼叫号码不支持
	ErrCode84061   = add(84061)   // 不存在外部联系人的关系
	ErrCode84062   = add(84062)   // 未开启公费电话应用
	ErrCode84063   = add(84063)   // caller不存在
	ErrCode84064   = add(84064)   // callee不存在
	ErrCode84065   = add(84065)   // caller跟callee电话号码一致
	ErrCode84066   = add(84066)   // 服务商拨打次数超过限制
	ErrCode84067   = add(84067)   // 管理员收到的服务商公费电话个数超过限制
	ErrCode84069   = add(84069)   // 拨打方被限制拨打公费电话
	ErrCode84070   = add(84070)   // 不支持的电话号码
	ErrCode84071   = add(84071)   // 不合法的外部联系人授权码
	ErrCode84072   = add(84072)   // 应用未配置客服
	ErrCode84073   = add(84073)   // 客服userid不在应用配置的客服列表中
	ErrCode84074   = add(84074)   // 没有外部联系人权限
	ErrCode84075   = add(84075)   // 不合法或过期的authcode
	ErrCode84076   = add(84076)   // 缺失authcode
	ErrCode84077   = add(84077)   // 订单价格过高，无法受理
	ErrCode84078   = add(84078)   // 购买人数不正确
	ErrCode84079   = add(84079)   // 价格策略不存在
	ErrCode84080   = add(84080)   // 订单不存在
	ErrCode84081   = add(84081)   // 存在未支付订单
	ErrCode84082   = add(84082)   // 存在申请退款中的订单
	ErrCode84083   = add(84083)   // 非服务人员
	ErrCode84084   = add(84084)   // 非跟进用户
	ErrCode84085   = add(84085)   // 应用已下架
	ErrCode84086   = add(84086)   // 订单人数超过可购买最大人数
	ErrCode84087   = add(84087)   // 打开订单支付前禁止关闭订单
	ErrCode84088   = add(84088)   // 禁止关闭已支付的订单
	ErrCode84089   = add(84089)   // 订单已支付
	ErrCode84090   = add(84090)   // 缺失user_ticket
	ErrCode84091   = add(84091)   // 订单价格不可低于下限
	ErrCode84092   = add(84092)   // 无法发起代下单操作
	ErrCode84093   = add(84093)   // 代理关系已占用，无法代下单
	ErrCode84094   = add(84094)   // 该应用未配置代理分润规则，请先联系应用服务商处理
	ErrCode84095   = add(84095)   // 免费试用版，无法扩容
	ErrCode84096   = add(84096)   // 免费试用版，无法续期
	ErrCode84097   = add(84097)   // 当前企业有未处理订单
	ErrCode84098   = add(84098)   // 固定总量，无法扩容
	ErrCode84099   = add(84099)   // 非购买状态，无法扩容
	ErrCode84100   = add(84100)   // 未购买过此应用，无法续期
	ErrCode84101   = add(84101)   // 企业已试用付费版本，无法全新购买
	ErrCode84102   = add(84102)   // 企业当前应用状态已过期，无法扩容
	ErrCode84103   = add(84103)   // 仅可修改未支付订单
	ErrCode84104   = add(84104)   // 订单已支付，无法修改
	ErrCode84105   = add(84105)   // 订单已被取消，无法修改
	ErrCode84106   = add(84106)   // 企业含有该应用的待支付订单，无法代下单
	ErrCode84107   = add(84107)   // 企业含有该应用的退款中订单，无法代下单
	ErrCode84108   = add(84108)   // 企业含有该应用的待生效订单，无法代下单
	ErrCode84109   = add(84109)   // 订单定价不能未0
	ErrCode84110   = add(84110)   // 新安装应用不在试用状态，无法升级为付费版
	ErrCode84111   = add(84111)   // 无足够可用优惠券
	ErrCode84112   = add(84112)   // 无法关闭未支付订单
	ErrCode84113   = add(84113)   // 无付费信息
	ErrCode84114   = add(84114)   // 虚拟版本不支持下单
	ErrCode84115   = add(84115)   // 虚拟版本不支持扩容
	ErrCode84116   = add(84116)   // 虚拟版本不支持续期
	ErrCode84117   = add(84117)   // 在虚拟正式版期内不能扩容
	ErrCode84118   = add(84118)   // 虚拟正式版期内不能变更版本
	ErrCode84119   = add(84119)   // 当前企业未报备，无法进行代下单
	ErrCode84120   = add(84120)   // 当前应用版本已删除
	ErrCode84121   = add(84121)   // 应用版本已删除，无法扩容
	ErrCode84122   = add(84122)   // 应用版本已删除，无法续期
	ErrCode84123   = add(84123)   // 非虚拟版本，无法升级
	ErrCode84124   = add(84124)   // 非行业方案订单，不能添加部分应用版本的订单
	ErrCode84125   = add(84125)   // 购买人数不能少于最少购买人数
	ErrCode84126   = add(84126)   // 购买人数不能多于最大购买人数
	ErrCode84127   = add(84127)   // 无应用管理权限
	ErrCode84128   = add(84128)   // 无该行业方案下全部应用的管理权限
	ErrCode84129   = add(84129)   // 付费策略已被删除，无法下单
	ErrCode84130   = add(84130)   // 订单生效时间不合法
	ErrCode84200   = add(84200)   // 文件转译解析错误
	ErrCode85002   = add(85002)   // 包含不合法的词语
	ErrCode85004   = add(85004)   // 每企业每个月设置的可信域名不可超过20个
	ErrCode85005   = add(85005)   // 可信域名未通过所有权校验
	ErrCode86001   = add(86001)   // 参数 chatid 不合法
	ErrCode86003   = add(86003)   // 参数 chatid 不存在
	ErrCode86004   = add(86004)   // 参数 群名不合法
	ErrCode86005   = add(86005)   // 参数 群主不合法
	ErrCode86006   = add(86006)   // 群成员数过多或过少
	ErrCode86007   = add(86007)   // 不合法的群成员
	ErrCode86008   = add(86008)   // 非法操作非自己创建的群
	ErrCode86101   = add(86101)   // 仅群主才有操作权限
	ErrCode86201   = add(86201)   // 参数 需要chatid
	ErrCode86202   = add(86202)   // 参数 需要群名
	ErrCode86203   = add(86203)   // 参数 需要群主
	ErrCode86204   = add(86204)   // 参数 需要群成员
	ErrCode86205   = add(86205)   // 参数 字符串chatid过长
	ErrCode86206   = add(86206)   // 参数 数字chatid过大
	ErrCode86207   = add(86207)   // 群主不在群成员列表
	ErrCode86214   = add(86214)   // 群发类型不合法
	ErrCode86215   = add(86215)   // 会话ID已经存在
	ErrCode86216   = add(86216)   // 存在非法会话成员ID
	ErrCode86217   = add(86217)   // 会话发送者不在会话成员列表中
	ErrCode86220   = add(86220)   // 指定的会话参数不合法
	ErrCode86224   = add(86224)   // 不是受限群，不允许使用该接口
	ErrCode90001   = add(90001)   // 未认证摇一摇周边
	ErrCode90002   = add(90002)   // 缺少摇一摇周边ticket参数
	ErrCode90003   = add(90003)   // 摇一摇周边ticket参数不合法
	ErrCode90100   = add(90100)   // 非法的对外属性类型
	ErrCode90101   = add(90101)   // 对外属性：文本类型长度不合法
	ErrCode90102   = add(90102)   // 对外属性：网页类型标题长度不合法
	ErrCode90103   = add(90103)   // 对外属性：网页url不合法
	ErrCode90104   = add(90104)   // 对外属性：小程序类型标题长度不合法
	ErrCode90105   = add(90105)   // 对外属性：小程序类型pagepath不合法
	ErrCode90106   = add(90106)   // 对外属性：请求参数不合法
	ErrCode90200   = add(90200)   // 缺少小程序appid参数
	ErrCode90201   = add(90201)   // 小程序通知的contDetailt_item个数超过限制
	ErrCode90202   = add(90202)   // 小程序通知中的key长度不合法
	ErrCode90203   = add(90203)   // 小程序通知中的value长度不合法
	ErrCode90204   = add(90204)   // 小程序通知中的page参数不合法
	ErrCode90206   = add(90206)   // 小程序未关联到企业中
	ErrCode90207   = add(90207)   // 不合法的小程序appid
	ErrCode90208   = add(90208)   // 小程序appid不匹配
	ErrCode90300   = add(90300)   // orderid 不合法
	ErrCode90302   = add(90302)   // 付费应用已过期
	ErrCode90303   = add(90303)   // 付费应用超过最大使用人数
	ErrCode90304   = add(90304)   // 订单中心服务异常，请稍后重试
	ErrCode90305   = add(90305)   // 参数错误，errmsg中有提示具体哪个参数有问题
	ErrCode90306   = add(90306)   // 商户设置不合法，详情请见errmsg
	ErrCode90307   = add(90307)   // 登录态过期
	ErrCode90308   = add(90308)   // 在开启IP鉴权的前提下，识别为无效的请求IP
	ErrCode90309   = add(90309)   // 订单已经存在，请勿重复下单
	ErrCode90310   = add(90310)   // 找不到订单
	ErrCode90311   = add(90311)   // 关单失败, 可能原因：该单并没被拉起支付页面; 已经关单；已经支付；渠道失败；单处于保护状态；等等
	ErrCode90312   = add(90312)   // 退款请求失败, 详情请看errmsg
	ErrCode90313   = add(90313)   // 退款调用频率限制，超过规定的阈值
	ErrCode90314   = add(90314)   // 订单状态错误，可能未支付，或者当前状态操作受限
	ErrCode90315   = add(90315)   // 退款请求失败，主键冲突，请核实退款refund_id是否已使用
	ErrCode90316   = add(90316)   // 退款原因编号不对
	ErrCode90317   = add(90317)   // 尚未注册成为供应商
	ErrCode90318   = add(90318)   // 参数nonce_str 为空或者重复，判定为重放攻击
	ErrCode90319   = add(90319)   // 时间戳为空或者与系统时间间隔太大
	ErrCode90320   = add(90320)   // 订单tokDetail无效
	ErrCode90321   = add(90321)   // 订单tokDetail已过有效时间
	ErrCode90322   = add(90322)   // 旧套件（包含多个应用的套件）不支持支付系统
	ErrCode90323   = add(90323)   // 单价超过限额
	ErrCode90324   = add(90324)   // 商品数量超过限额
	ErrCode90325   = add(90325)   // 预支单已经存在
	ErrCode90326   = add(90326)   // 预支单单号非法
	ErrCode90327   = add(90327)   // 该预支单已经结算下单
	ErrCode90328   = add(90328)   // 结算下单失败，详情请看errmsg
	ErrCode90329   = add(90329)   // 该订单号已经被预支单占用
	ErrCode90330   = add(90330)   // 创建供应商失败
	ErrCode90331   = add(90331)   // 更新供应商失败
	ErrCode90332   = add(90332)   // 还没签署合同
	ErrCode90333   = add(90333)   // 创建合同失败
	ErrCode90338   = add(90338)   // 已经过了可退款期限
	ErrCode90339   = add(90339)   // 供应商主体名包含非法字符
	ErrCode90340   = add(90340)   // 创建客户失败，可能信息真实性校验失败
	ErrCode90341   = add(90341)   // 退款金额大于付款金额
	ErrCode90342   = add(90342)   // 退款金额超过账户余额
	ErrCode90343   = add(90343)   // 退款单号已经存在
	ErrCode90344   = add(90344)   // 指定的付款渠道无效
	ErrCode90345   = add(90345)   // 超过5w人民币不可指定微信支付渠道
	ErrCode90346   = add(90346)   // 同一单的退款次数超过限制
	ErrCode90347   = add(90347)   // 退款金额不可为0
	ErrCode90348   = add(90348)   // 管理端没配置支付密钥
	ErrCode90349   = add(90349)   // 记录数量太大
	ErrCode90350   = add(90350)   // 银行信息真实性校验失败
	ErrCode90351   = add(90351)   // 应用状态异常
	ErrCode90352   = add(90352)   // 延迟试用期天数超过限制
	ErrCode90353   = add(90353)   // 预支单列表不可为空
	ErrCode90354   = add(90354)   // 预支单列表数量超过限制
	ErrCode90355   = add(90355)   // 关联有退款预支单，不可删除
	ErrCode90356   = add(90356)   // 不能0金额下单
	ErrCode90357   = add(90357)   // 代下单必须指定支付渠道
	ErrCode90358   = add(90358)   // 预支单或代下单，不支持部分退款
	ErrCode90359   = add(90359)   // 预支单与下单者企业不匹配
	ErrCode90381   = add(90381)   // 参数 refunded_credit_orderid 不合法
	ErrCode90456   = add(90456)   // 必须指定组织者
	ErrCode90457   = add(90457)   // 日历ID异常
	ErrCode90458   = add(90458)   // 日历ID列表不能为空
	ErrCode90459   = add(90459)   // 日历已删除
	ErrCode90460   = add(90460)   // 日程已删除
	ErrCode90461   = add(90461)   // 日程ID异常
	ErrCode90462   = add(90462)   // 日程ID列表不能为空
	ErrCode90463   = add(90463)   // 不能变更组织者
	ErrCode90464   = add(90464)   // 参与者数量超过限制
	ErrCode90465   = add(90465)   // 不支持的重复类型
	ErrCode90466   = add(90466)   // 不能操作别的应用创建的日历/日程
	ErrCode90467   = add(90467)   // 星期参数异常
	ErrCode90468   = add(90468)   // 不能变更组织者
	ErrCode90469   = add(90469)   // 每页大小超过限制
	ErrCode90470   = add(90470)   // 页数异常
	ErrCode90471   = add(90471)   // 提醒时间异常
	ErrCode90472   = add(90472)   // 没有日历/日程操作权限
	ErrCode90473   = add(90473)   // 颜色参数异常
	ErrCode90474   = add(90474)   // 组织者不能与参与者重叠
	ErrCode90475   = add(90475)   // 不是组织者的日历
	ErrCode90479   = add(90479)   // 不允许操作用户创建的日程
	ErrCode90500   = add(90500)   // 群主并未离职
	ErrCode90501   = add(90501)   // 该群不是客户群
	ErrCode90502   = add(90502)   // 群主已经离职
	ErrCode90503   = add(90503)   // 满人 & 99个微信成员，没办法踢，要客户端确认
	ErrCode90504   = add(90504)   // 群主没变
	ErrCode90507   = add(90507)   // 离职群正在继承处理中
	ErrCode90508   = add(90508)   // 离职群已经继承
	ErrCode91040   = add(91040)   // 获取ticket的类型无效
	ErrCode92000   = add(92000)   // 成员不在应用可见范围之内
	ErrCode92001   = add(92001)   // 应用没有敏感信息权限
	ErrCode92002   = add(92002)   // 不允许跨企业调用
	ErrCode92006   = add(92006)   // 该直播已经开始或取消
	ErrCode92007   = add(92007)   // 该直播回放不能被删除
	ErrCode92008   = add(92008)   // 当前应用没权限操作这个直播
	ErrCode93000   = add(93000)   // 机器人webhookurl不合法或者机器人已经被移除出群
	ErrCode93004   = add(93004)   // 机器人被停用
	ErrCode93008   = add(93008)   // 不在群里
	ErrCode94000   = add(94000)   // 应用未开启工作台自定义模式
	ErrCode94001   = add(94001)   // 不合法的type类型
	ErrCode94002   = add(94002)   // 缺少keydata字段
	ErrCode94003   = add(94003)   // keydata的items列表长度超出限制
	ErrCode94005   = add(94005)   // 缺少list字段
	ErrCode94006   = add(94006)   // list的items列表长度超出限制
	ErrCode94007   = add(94007)   // 缺少webview字段
	ErrCode94008   = add(94008)   // 应用未设置自定义工作台模版类型
	ErrCode301002  = add(301002)  // 无权限操作指定的应用
	ErrCode301005  = add(301005)  // 不允许删除创建者
	ErrCode301012  = add(301012)  // 参数 position 不合法
	ErrCode301013  = add(301013)  // 参数 telephone 不合法
	ErrCode301014  = add(301014)  // 参数 Detailglish_name 不合法
	ErrCode301015  = add(301015)  // 参数 mediaid 不合法
	ErrCode301016  = add(301016)  // 上传语音文件不符合系统要求
	ErrCode301017  = add(301017)  // 上传语音文件仅支持AMR格式
	ErrCode301021  = add(301021)  // 参数 userid 无效
	ErrCode301022  = add(301022)  // 获取打卡数据失败
	ErrCode301023  = add(301023)  // useridlist非法或超过限额
	ErrCode301024  = add(301024)  // 获取打卡记录时间间隔超限
	ErrCode301025  = add(301025)  // 审批开放接口参数错误
	ErrCode301036  = add(301036)  // 不允许更新该用户的userid
	ErrCode301039  = add(301039)  // 请求参数错误，请检查输入参数
	ErrCode301042  = add(301042)  // ip白名单限制，请求ip不在设置白名单范围
	ErrCode301048  = add(301048)  // sdkfileid对应的文件不存在或已过期
	ErrCode301052  = add(301052)  // 会话存档服务已过期
	ErrCode301053  = add(301053)  // 会话存档服务未开启
	ErrCode301058  = add(301058)  // 拉取会话数据请求超过大小限制，可减少limit参数
	ErrCode301059  = add(301059)  // 非内部群，不提供数据
	ErrCode301060  = add(301060)  // 拉取同意情况请求量过大，请减少到100个参数以下
	ErrCode301061  = add(301061)  // userid或者exteropDetailid用户不存在
	ErrCode302003  = add(302003)  // 批量导入任务的文件中userid有重复
	ErrCode302004  = add(302004)  // 组织架构不合法（1不是一棵树，2 多个一样的partyid，3 partyid空，4 partyid name 空，5 同一个父节点下有两个子节点 部门名字一样 可能是以上情况，请一一排查）
	ErrCode302005  = add(302005)  // 批量导入系统失败，请重新尝试导入
	ErrCode302006  = add(302006)  // 批量导入任务的文件中partyid有重复
	ErrCode302007  = add(302007)  // 批量导入任务的文件中，同一个部门下有两个子部门名字一样
	ErrCode2000002 = add(2000002) // CorpId参数无效
	ErrCode600001  = add(600001)  // 不合法的sn
	ErrCode600002  = add(600002)  // 设备已注册
	ErrCode600003  = add(600003)  // 不合法的硬件activecode
	ErrCode600004  = add(600004)  // 该硬件尚未授权任何企业
	ErrCode600005  = add(600005)  // 硬件Secret无效
	ErrCode600007  = add(600007)  // 缺少硬件sn
	ErrCode600008  = add(600008)  // 缺少nonce参数
	ErrCode600009  = add(600009)  // 缺少timestamp参数
	ErrCode600010  = add(600010)  // 缺少signature参数
	ErrCode600011  = add(600011)  // 签名校验失败
	ErrCode600012  = add(600012)  // 长连接已经注册过设备
	ErrCode600013  = add(600013)  // 缺少activecode参数
	ErrCode600014  = add(600014)  // 设备未网络注册
	ErrCode600015  = add(600015)  // 缺少secret参数
	ErrCode600016  = add(600016)  // 设备未激活
	ErrCode600018  = add(600018)  // 无效的起始结束时间
	ErrCode600020  = add(600020)  // 设备未登录
	ErrCode600021  = add(600021)  // 设备sn已存在
	ErrCode600023  = add(600023)  // 时间戳已失效
	ErrCode600024  = add(600024)  // 固件大小超过5M
	ErrCode600025  = add(600025)  // 固件名为空或者超过20字节
	ErrCode600026  = add(600026)  // 固件信息不存在
	ErrCode600027  = add(600027)  // 非法的固件参数
	ErrCode600028  = add(600028)  // 固件版本已存在
	ErrCode600029  = add(600029)  // 非法的固件版本
	ErrCode600030  = add(600030)  // 缺少固件版本参数
	ErrCode600031  = add(600031)  // 硬件固件不允许升级
	ErrCode600032  = add(600032)  // 无法解析硬件二维码
	ErrCode600033  = add(600033)  // 设备型号id冲突
	ErrCode600034  = add(600034)  // 指纹数据大小超过限制
	ErrCode600035  = add(600035)  // 人脸数据大小超过限制
	ErrCode600036  = add(600036)  // 设备sn冲突
	ErrCode600037  = add(600037)  // 缺失设备型号id
	ErrCode600038  = add(600038)  // 设备型号不存在
	ErrCode600039  = add(600039)  // 不支持的设备类型
	ErrCode600040  = add(600040)  // 打印任务id不存在
	ErrCode600041  = add(600041)  // 无效的offset或limit参数值
	ErrCode600042  = add(600042)  // 无效的设备型号id
	ErrCode600043  = add(600043)  // 门禁规则未设置
	ErrCode600044  = add(600044)  // 门禁规则不合法
	ErrCode600045  = add(600045)  // 设备已订阅企业信息
	ErrCode600046  = add(600046)  // 操作id和用户userid不匹配
	ErrCode600047  = add(600047)  // secretno的status非法
	ErrCode600048  = add(600048)  // 无效的指纹算法
	ErrCode600049  = add(600049)  // 无效的人脸识别算法
	ErrCode600050  = add(600050)  // 无效的算法长度
	ErrCode600051  = add(600051)  // 设备过期
	ErrCode600052  = add(600052)  // 无效的文件分块
	ErrCode600053  = add(600053)  // 该链接已经激活
	ErrCode600054  = add(600054)  // 该链接已经订阅
	ErrCode600055  = add(600055)  // 无效的用户类型
	ErrCode600056  = add(600056)  // 无效的健康状态
	ErrCode600057  = add(600057)  // 缺少体温参数
	ErrCode610001  = add(610001)  // 永久二维码超过每个员工5000的限制
	ErrCode610003  = add(610003)  // scDetaile参数不合法
	ErrCode610004  = add(610004)  // userid不在客户联系配置的使用范围内
	ErrCode610014  = add(610014)  // 无效的unionid
	ErrCode610015  = add(610015)  // 小程序对应的开放平台账号未认证
	ErrCode610016  = add(610016)  // 企业未认证
	ErrCode610017  = add(610017)  // 小程序和企业主体不一致
	ErrCode640001  = add(640001)  // 微盘不存在当前空间
	ErrCode640002  = add(640002)  // 文件不存在
	ErrCode640003  = add(640003)  // 文件已删除
	ErrCode640004  = add(640004)  // 无权限访问
	ErrCode640005  = add(640005)  // 成员不在空间内
	ErrCode640006  = add(640006)  // 超出当前成员拥有的容量
	ErrCode640007  = add(640007)  // 超出微盘的容量
	ErrCode640008  = add(640008)  // 没有空间权限
	ErrCode640009  = add(640009)  // 非法文件名
	ErrCode640010  = add(640010)  // 超出空间的最大成员数
	ErrCode640011  = add(640011)  // json格式不匹配
	ErrCode640012  = add(640012)  // 非法的userid
	ErrCode640013  = add(640013)  // 非法的departmDetailtid
	ErrCode640014  = add(640014)  // 空间没有有效的管理员
	ErrCode640015  = add(640015)  // 不支持设置预览权限
	ErrCode640016  = add(640016)  // 不支持设置文件水印
	ErrCode640017  = add(640017)  // 微盘管理端未开通API 权限
	ErrCode640018  = add(640018)  // 微盘管理端未设置编辑权限
	ErrCode640019  = add(640019)  // API 调用次数超出限制
	ErrCode640020  = add(640020)  // 非法的权限类型
	ErrCode640021  = add(640021)  // 非法的fatherid
	ErrCode640022  = add(640022)  // 非法的文件内容的base64
	ErrCode640023  = add(640023)  // 非法的权限范围
	ErrCode640024  = add(640024)  // 非法的fileid
	ErrCode640025  = add(640025)  // 非法的space_name
	ErrCode640026  = add(640026)  // 非法的spaceid
	ErrCode640027  = add(640027)  // 参数错误
	ErrCode640028  = add(640028)  // 空间设置了关闭成员邀请链接
	ErrCode640029  = add(640029)  // 只支持下载普通文件，不支持下载文件夹等其他非文件实体类型
	ErrCode844001  = add(844001)  // 非法的output_file_format

)

var _messages = map[int]Message{
	ErrCode6000.Code(): {
		Msg:    `数据版本冲突`,
		Detail: `排查方法: 可能有多个调用端同时修改数据，稍后重试`,
	},
	ErrCode40001.Code(): {
		Msg:    `不合法的secret参数`,
		Detail: `排查方法: secret在应用详情/通讯录管理助手可查看`,
	},
	ErrCode40003.Code(): {
		Msg:    `无效的UserID`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40003);不合法的UserID。确认：;1）有效的UserID需要满足：长度1~64字符，由英文字母、数字、中划线、下划线以及点号构成。;2）除了创建用户，其余使用UserID的接口，还要保证UserID必须在通讯录中存在。`,
	},
	ErrCode40004.Code(): {
		Msg:    `不合法的媒体文件类型`,
		Detail: `排查方法: 不满足系统文件要求。参考：[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112)`,
	},
	ErrCode40005.Code(): {
		Msg:    `不合法的type参数`,
		Detail: `排查方法: 合法的type取值，参考：[上传临时素材](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112)`,
	},
	ErrCode40006.Code(): {
		Msg:    `不合法的文件大小`,
		Detail: `排查方法: 系统文件要求，参考：[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112)`,
	},
	ErrCode40007.Code(): {
		Msg:    `不合法的media_id参数`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40007);不合法的媒体文件。确认：;1）媒体文件ID的获取方式，是否存在。注：上传临时素材生成的medida_id，有效期是3天。;2）媒体文件类型应符合接口要求（比如发送图片消息，此时不能用音频文件的media_id）。`,
	},
	ErrCode40008.Code(): {
		Msg:    `不合法的msgtype参数`,
		Detail: `排查方法: 合法的msgtype取值，参考：[消息类型](https://work.weixin.qq.com/api/doc/90000/90139/90313#10167)`,
	},
	ErrCode40009.Code(): {
		Msg:    `上传图片大小不是有效值`,
		Detail: `排查方法: 图片大小的系统限制，参考[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112/上传的媒体文件限制)`,
	},
	ErrCode40011.Code(): {
		Msg:    `上传视频大小不是有效值`,
		Detail: `排查方法: 视频大小的系统限制，参考[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112/上传的媒体文件限制)`,
	},
	ErrCode40013.Code(): {
		Msg:    `不合法的CorpID`,
		Detail: `排查方法: 需确认CorpID是否填写正确，在 web管理端-设置 可查看`,
	},
	ErrCode40014.Code(): {
		Msg:    `不合法的access_tokDetail`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40014);access_tokDetail参数错误。确认：;1）access_tokDetail的获取方式;2）access_tokDetail是否已过期;可以重新获取一次access_tokDetail解决`,
	},
	ErrCode40016.Code(): {
		Msg:    `不合法的按钮个数`,
		Detail: `排查方法: 菜单按钮1-3个`,
	},
	ErrCode40017.Code(): {
		Msg:    `不合法的按钮类型`,
		Detail: `排查方法: 支持的类型，参考：[按钮类型](https://work.weixin.qq.com/api/doc/90000/90139/90313#10786)`,
	},
	ErrCode40018.Code(): {
		Msg:    `不合法的按钮名字长度`,
		Detail: `排查方法: 长度应不超过16个字节`,
	},
	ErrCode40019.Code(): {
		Msg:    `不合法的按钮KEY长度`,
		Detail: `排查方法: 长度应不超过128字节`,
	},
	ErrCode40020.Code(): {
		Msg:    `不合法的按钮URL长度`,
		Detail: `排查方法: 长度应不超过1024字节`,
	},
	ErrCode40022.Code(): {
		Msg:    `不合法的子菜单级数`,
		Detail: `排查方法: 只能包含一级菜单和二级菜单`,
	},
	ErrCode40023.Code(): {
		Msg:    `不合法的子菜单按钮个数`,
		Detail: `排查方法: 子菜单按钮1-5个`,
	},
	ErrCode40024.Code(): {
		Msg:    `不合法的子菜单按钮类型`,
		Detail: `排查方法: 支持的类型，参考：[按钮类型](https://work.weixin.qq.com/api/doc/90000/90139/90313#10786)`,
	},
	ErrCode40025.Code(): {
		Msg:    `不合法的子菜单按钮名字长度`,
		Detail: `排查方法: 支持的类型，参考：[按钮类型](https://work.weixin.qq.com/api/doc/90000/90139/90313#10786)`,
	},
	ErrCode40026.Code(): {
		Msg:    `不合法的子菜单按钮KEY长度`,
		Detail: `排查方法: -`,
	},
	ErrCode40027.Code(): {
		Msg:    `不合法的子菜单按钮URL长度`,
		Detail: `排查方法: 长度应不超过1024字节`,
	},
	ErrCode40029.Code(): {
		Msg:    `不合法的oauth_code`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40029);oauth_code参数错误。确认：;1）code只能消费一次，不能重复消费。比如说，是否存在多个服务器同时消费同一code情况。;2）code需要在有效期间消费（5分钟），过期会自动失效。`,
	},
	ErrCode40031.Code(): {
		Msg:    `不合法的UserID列表`,
		Detail: `排查方法: 指定的UserID列表，部分UserID不在通讯录中`,
	},
	ErrCode40032.Code(): {
		Msg:    `不合法的UserID列表长度`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40032);不合法的UserID列表长度。确认：;1）[发消息接口](https://opDetail.work.weixin.qq.com/api/doc#10167)，最多指定1000人。;2）[批量删除成员接口](https://opDetail.work.weixin.qq.com/api/doc#10060)，最多指定200人。`,
	},
	ErrCode40033.Code(): {
		Msg:    `不合法的请求字符`,
		Detail: `排查方法: 不能包含\uxxxx格式的字符`,
	},
	ErrCode40035.Code(): {
		Msg:    `不合法的参数`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40035);不合法的参数。确认：;1）userlist和partylist不能同时为空;2）userlist包含的成员个数不能大于1000;3）partylist包含的部门个数不能大于100;4）指定的userlist和partylist为数组格式，不是字符串格式。比如说， “userlist”:[ “user1”,”user2”]，而不是 “userlist”: “user1|user2”`,
	},
	ErrCode40036.Code(): {
		Msg:    `不合法的模板id长度`,
		Detail: `排查方法: -`,
	},
	ErrCode40037.Code(): {
		Msg:    `无效的模板id`,
		Detail: `排查方法: -`,
	},
	ErrCode40039.Code(): {
		Msg:    `不合法的url长度`,
		Detail: `排查方法: url长度限制1024个字节`,
	},
	ErrCode40050.Code(): {
		Msg:    `chatid不存在`,
		Detail: `排查方法: 会话需要先创建后，才可修改会话详情或者发起聊天`,
	},
	ErrCode40054.Code(): {
		Msg:    `不合法的子菜单url域名`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40054 40055);菜单设置URL不合法。确认：;1）链接需要带上协议头。以 http:// 或者 https:// 开头。比如：https://work.weixin.qq.com;2）微信支付的链接，必须以 weixin://wxpay/bizpayurl 开头`,
	},
	ErrCode40055.Code(): {
		Msg:    `不合法的菜单url域名`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40054 40055);菜单设置URL不合法。确认：;1）链接需要带上协议头。以 http:// 或者 https:// 开头。比如：https://work.weixin.qq.com;2）微信支付的链接，必须以 weixin://wxpay/bizpayurl 开头`,
	},
	ErrCode40056.Code(): {
		Msg:    `不合法的agDetailtid`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40056);agDetailtid不合法。确认：;1）agDetailtid为整型数字;2）在web管理端存在该应用`,
	},
	ErrCode40057.Code(): {
		Msg:    `不合法的callbackurl或者callbackurl验证失败`,
		Detail: `排查方法: 可自助到[开发调试工具](https://work.weixin.qq.com/api/devtools/devtool.php)重现`,
	},
	ErrCode40058.Code(): {
		Msg:    `不合法的参数`,
		Detail: `排查方法: 传递参数不符合系统要求，需要参照具体API接口说明`,
	},
	ErrCode40059.Code(): {
		Msg:    `不合法的上报地理位置标志位`,
		Detail: `排查方法: 开关标志位只能填 0 或者 1`,
	},
	ErrCode40063.Code(): {
		Msg:    `参数为空`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40063);必填的参数缺少，需要参照具体API接口说明。同时确认：;1）Http请求方法，是否正确。比如说接口要求以Post方法，就不能使用Get方式;2）Http请求参数，是否正确。比如说，接口内容要求json结构体，就不能以url参数传递或者form-data方式。`,
	},
	ErrCode40066.Code(): {
		Msg:    `不合法的部门列表`,
		Detail: `排查方法: 部门列表为空，或者至少存在一个部门ID不存在于通讯录中`,
	},
	ErrCode40068.Code(): {
		Msg:    `不合法的标签ID`,
		Detail: `排查方法: 标签ID未指定，或者指定的标签ID不存在`,
	},
	ErrCode40070.Code(): {
		Msg:    `指定的标签范围结点全部无效`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40070);指定的标签范围结点全部无效。确认：;1）指定的参数格式是否正确。比如，”userlist”:[ “user1”]，而不是指定为 “userlist” : “user1”。;2）指定的成员或者部门，是否存在于通讯录中。`,
	},
	ErrCode40071.Code(): {
		Msg:    `不合法的标签名字`,
		Detail: `排查方法: 标签名字已经存在`,
	},
	ErrCode40072.Code(): {
		Msg:    `不合法的标签名字长度`,
		Detail: `排查方法: 不允许为空，最大长度限制为32个字（汉字或英文字母）`,
	},
	ErrCode40073.Code(): {
		Msg:    `不合法的opDetailid`,
		Detail: `排查方法: opDetailid不存在，需确认获取来源`,
	},
	ErrCode40074.Code(): {
		Msg:    `news消息不支持保密消息类型`,
		Detail: `排查方法: 图文消息支持保密类型需改用mpnews`,
	},
	ErrCode40077.Code(): {
		Msg:    `不合法的pre_auth_code参数`,
		Detail: `排查方法: 预授权码不存在，参考：[获取预授权码](https://work.weixin.qq.com/api/doc/90000/90139/90313#10975/获取预授权码)`,
	},
	ErrCode40078.Code(): {
		Msg:    `不合法的auth_code参数`,
		Detail: `排查方法: 需确认获取来源，并且只能消费一次`,
	},
	ErrCode40080.Code(): {
		Msg:    `不合法的suite_secret`,
		Detail: `排查方法: 套件secret可在第三方管理端套件详情查看`,
	},
	ErrCode40082.Code(): {
		Msg:    `不合法的suite_tokDetail`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40082);suite_tokDetail参数错误。确认：;1）suite_tokDetail的获取方式;2）suite_tokDetail是否已过期;可以重新获取一次suite_tokDetail解决`,
	},
	ErrCode40083.Code(): {
		Msg:    `不合法的suite_id`,
		Detail: `排查方法: suite_id不存在`,
	},
	ErrCode40084.Code(): {
		Msg:    `不合法的permanDetailt_code参数`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40084);不合法的永久授权码。确认：;1）是否填写有误;2）企业是否已取消授权该套件;3）永久授权码不能跨服务商使用`,
	},
	ErrCode40085.Code(): {
		Msg:    `不合法的的suite_ticket参数`,
		Detail: `排查方法: suite_ticket不存在或者已失效`,
	},
	ErrCode40086.Code(): {
		Msg:    `不合法的第三方应用appid`,
		Detail: `排查方法: 至少有一个不存在应用id`,
	},
	ErrCode40088.Code(): {
		Msg:    `jobid不存在`,
		Detail: `排查方法: 请检查 jobid 来源`,
	},
	ErrCode40089.Code(): {
		Msg:    `批量任务的结果已清理`,
		Detail: `排查方法: 系统仅保存最近5次批量任务的结果。可在通讯录查看实际导入情况`,
	},
	ErrCode40091.Code(): {
		Msg:    `secret不合法`,
		Detail: `排查方法: 可能用了别的企业的secret`,
	},
	ErrCode40092.Code(): {
		Msg:    `导入文件存在不合法的内容`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40092);导入文件存在不合法的内容。确认：;1）不允许上传空文件;2）文件内容缺少必填字段，比如：手机/邮箱，姓名，UserID或者部门。`,
	},
	ErrCode40093.Code(): {
		Msg:    `jsapi签名错误`,
		Detail: `排查方法: 请检查用于签名的jsapi_ticket是否是正确的，是否过期。可以通过获取相应jsapi_ticket接口获取当前的jsapi_ticket跟用于签名的jsapi_ticket比对是否一致，若jsapi_ticket还在有效期内，当前获取到的jsapi_ticket是一致的。若jsapi_ticket没问题，请检查用于签名的url参数是不是正确的， url（当前网页的URL， 不包含#及其后面部分）。`,
	},
	ErrCode40094.Code(): {
		Msg:    `不合法的URL`,
		Detail: `排查方法: 缺少主页URL参数，或者URL不合法（链接需要带上协议头，以 http:// 或者 https:// 开头）`,
	},
	ErrCode40096.Code(): {
		Msg:    `不合法的外部联系人userid`,
		Detail: `排查方法: -`,
	},
	ErrCode40097.Code(): {
		Msg:    `该成员尚未离职`,
		Detail: `排查方法: 离职成员外部联系人转移接口要求转出用户必须已经离职`,
	},
	ErrCode40098.Code(): {
		Msg:    `成员尚未实名认证`,
		Detail: `排查方法: 确认传入的userid是已经过实名认证成员的`,
	},
	ErrCode40099.Code(): {
		Msg:    `外部联系人的数量已达上限`,
		Detail: `排查方法: -`,
	},
	ErrCode40100.Code(): {
		Msg:    `此用户的外部联系人已经在转移流程中`,
		Detail: `排查方法: -`,
	},
	ErrCode40102.Code(): {
		Msg:    `域名或IP不可与应用市场上架应用重复`,
		Detail: `排查方法: -`,
	},
	ErrCode40123.Code(): {
		Msg:    `上传临时图片素材，图片格式非法`,
		Detail: `排查方法: 请确认上传的内容是否为合法的图片内容`,
	},
	ErrCode40124.Code(): {
		Msg:    `推广活动里的sn禁止绑定`,
		Detail: `排查方法: -`,
	},
	ErrCode40125.Code(): {
		Msg:    `无效的opDetailuserid参数`,
		Detail: `排查方法: -`,
	},
	ErrCode40126.Code(): {
		Msg:    `企业标签个数达到上限，最多为3000个`,
		Detail: `排查方法: -`,
	},
	ErrCode40127.Code(): {
		Msg:    `不支持的uri schema`,
		Detail: `排查方法: 检查uri链接的schema是否符合参数要求`,
	},
	ErrCode40128.Code(): {
		Msg:    `客户转接过于频繁（90天内只允许转接一次，同一个客户最多只能转接两次）`,
		Detail: `排查方法: -`,
	},
	ErrCode40129.Code(): {
		Msg:    `当前客户正在转接中`,
		Detail: `排查方法: -`,
	},
	ErrCode40130.Code(): {
		Msg:    `原跟进人与接手人一样，不可继承`,
		Detail: `排查方法: -`,
	},
	ErrCode40131.Code(): {
		Msg:    `handover_userid 并不是外部联系人的跟进人`,
		Detail: `排查方法: -`,
	},
	ErrCode41001.Code(): {
		Msg:    `缺少access_tokDetail参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41002.Code(): {
		Msg:    `缺少corpid参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41004.Code(): {
		Msg:    `缺少secret参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41006.Code(): {
		Msg:    `缺少media_id参数`,
		Detail: `排查方法: media_id为调用接口必填参数，请确认是否有传递`,
	},
	ErrCode41008.Code(): {
		Msg:    `缺少auth code参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41009.Code(): {
		Msg:    `缺少userid参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41010.Code(): {
		Msg:    `缺少url参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41011.Code(): {
		Msg:    `缺少agDetailtid参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41016.Code(): {
		Msg:    `缺少title参数`,
		Detail: `排查方法: 发送图文消息，标题是必填参数。请确认参数是否有传递。`,
	},
	ErrCode41019.Code(): {
		Msg:    `缺少 departmDetailt 参数`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：41019);缺少 departmDetailt 参数。确认：;1）创建成员接口，成员所属部门是必填信息。;2）所属部门是数字数组格式，不是字符串。如：”departmDetailt: [1, 2]`,
	},
	ErrCode41017.Code(): {
		Msg:    `缺少tagid参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41021.Code(): {
		Msg:    `缺少suite_id参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41022.Code(): {
		Msg:    `缺少suite_access_tokDetail参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41023.Code(): {
		Msg:    `缺少suite_ticket参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41024.Code(): {
		Msg:    `缺少secret参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41025.Code(): {
		Msg:    `缺少permanDetailt_code参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41033.Code(): {
		Msg:    `缺少 description 参数`,
		Detail: `排查方法: [发送文本卡片消息接口](https://work.weixin.qq.com/api/doc/90000/90139/90313#10167/文本卡片消息)，description 是必填字段`,
	},
	ErrCode41035.Code(): {
		Msg:    `缺少外部联系人userid参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41036.Code(): {
		Msg:    `不合法的企业对外简称`,
		Detail: `排查方法: 企业对外简称必须是认证过的，如果要改回默认简称，传空字符串把对外简称清除就可以了`,
	},
	ErrCode41037.Code(): {
		Msg:    `缺少「联系我」type参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41038.Code(): {
		Msg:    `缺少「联系我」scDetaile参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41039.Code(): {
		Msg:    `无效的「联系我」type参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41040.Code(): {
		Msg:    `无效的「联系我」scDetaile参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41041.Code(): {
		Msg:    `「联系我」使用人数超过限制`,
		Detail: `排查方法: 默认限制不超过100人(包括部门展开后的人数)`,
	},
	ErrCode41042.Code(): {
		Msg:    `无效的「联系我」style参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41043.Code(): {
		Msg:    `缺少「联系我」config_id参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41044.Code(): {
		Msg:    `无效的「联系我」config_id参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41045.Code(): {
		Msg:    `API添加「联系我」达到数量上限`,
		Detail: `排查方法: -`,
	},
	ErrCode41046.Code(): {
		Msg:    `缺少企业群发消息id`,
		Detail: `排查方法: -`,
	},
	ErrCode41047.Code(): {
		Msg:    `无效的企业群发消息id`,
		Detail: `排查方法: -`,
	},
	ErrCode41048.Code(): {
		Msg:    `无可发送的客户`,
		Detail: `排查方法: -`,
	},
	ErrCode41049.Code(): {
		Msg:    `缺少欢迎语code参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41050.Code(): {
		Msg:    `无效的欢迎语code`,
		Detail: `排查方法: 欢迎语code(welcome_code)具有时效性，须在添加好友后20秒内使用`,
	},
	ErrCode41051.Code(): {
		Msg:    `客户和服务人员已经开始聊天了`,
		Detail: `排查方法: 已经开始的聊天的客户不能发送欢迎语`,
	},
	ErrCode41052.Code(): {
		Msg:    `无效的发送时间`,
		Detail: `排查方法: -`,
	},
	ErrCode41053.Code(): {
		Msg:    `客户未同意聊天存档`,
		Detail: `排查方法: 须外部联系人同意服务须知后，成员才可发送欢迎语`,
	},
	ErrCode41054.Code(): {
		Msg:    `该用户尚未激活`,
		Detail: `排查方法: -`,
	},
	ErrCode41055.Code(): {
		Msg:    `群欢迎语模板数量达到上限`,
		Detail: `排查方法: -`,
	},
	ErrCode41056.Code(): {
		Msg:    `外部联系人id类型不正确`,
		Detail: `排查方法: -`,
	},
	ErrCode41057.Code(): {
		Msg:    `企业或服务商未绑定微信开发者账号`,
		Detail: `排查方法: -`,
	},
	ErrCode41059.Code(): {
		Msg:    `缺少momDetailt_id参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41060.Code(): {
		Msg:    `不合法的momDetailt_id参数`,
		Detail: `排查方法: -`,
	},
	ErrCode41061.Code(): {
		Msg:    `不合法朋友圈发送成员userid，当前朋友圈并非此用户发表`,
		Detail: `排查方法: -`,
	},
	ErrCode41062.Code(): {
		Msg:    `企业创建的朋友圈尚未被成员userid发表`,
		Detail: `排查方法: -`,
	},
	ErrCode41063.Code(): {
		Msg:    `群发消息正在被派发中，请稍后再试`,
		Detail: `排查方法: [创建企业群发](https://work.weixin.qq.com/api/doc/90000/90139/90313#15836)后，立刻调用[获取企业的全部群发记录](https://work.weixin.qq.com/api/doc/90000/90139/90313#25429)的相关接口，将可能出现该错误`,
	},
	ErrCode41064.Code(): {
		Msg:    `附件大小超过限制`,
		Detail: `排查方法: -`,
	},
	ErrCode41065.Code(): {
		Msg:    `无效的附件类型`,
		Detail: `排查方法: -`,
	},
	ErrCode41066.Code(): {
		Msg:    `用户视频号名称错误`,
		Detail: `排查方法: -`,
	},
	ErrCode41102.Code(): {
		Msg:    `缺少菜单名`,
		Detail: `排查方法: -`,
	},
	ErrCode42001.Code(): {
		Msg:    `access_tokDetail已过期`,
		Detail: `排查方法: access_tokDetail有时效性，需要重新获取一次`,
	},
	ErrCode42007.Code(): {
		Msg:    `pre_auth_code已过期`,
		Detail: `排查方法: pre_auth_code有时效性，需要重新获取一次`,
	},
	ErrCode42009.Code(): {
		Msg:    `suite_access_tokDetail已过期`,
		Detail: `排查方法: suite_access_tokDetail有时效性，需要重新获取一次`,
	},
	ErrCode42012.Code(): {
		Msg:    `jsapi_ticket不可用，一般是没有正确调用接口来创建jsapi_ticket`,
		Detail: `排查方法: 如果是agDetailtConfig使用，请特别注意是否是使用”[获取应用身份的ticket](https://work.weixin.qq.com/api/doc/90000/90139/90313#10029/获取应用的jsapi_ticket)“来获取jsapi_ticket`,
	},
	ErrCode42013.Code(): {
		Msg:    `小程序未登陆或登录态已经过期`,
		Detail: `排查方法: 需要重新走登陆流程`,
	},
	ErrCode42014.Code(): {
		Msg:    `任务卡片消息的task_id不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode42015.Code(): {
		Msg:    `更新的消息的应用与发送消息的应用不匹配`,
		Detail: `排查方法: -`,
	},
	ErrCode42016.Code(): {
		Msg:    `更新的task_id不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode42017.Code(): {
		Msg:    `按钮key值不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode42018.Code(): {
		Msg:    `按钮key值不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode42019.Code(): {
		Msg:    `缺少按钮key值不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode42020.Code(): {
		Msg:    `缺少按钮名称`,
		Detail: `排查方法: -`,
	},
	ErrCode42021.Code(): {
		Msg:    `device_access_tokDetail 过期`,
		Detail: `排查方法: -`,
	},
	ErrCode42022.Code(): {
		Msg:    `code已经被使用过。只能使用一次`,
		Detail: `排查方法: -`,
	},
	ErrCode43004.Code(): {
		Msg:    `指定的userid未绑定微信或未关注微工作台（原企业号）`,
		Detail: `排查方法: 需要成员使用微信登录企业微信或者关注微工作台才能获取opDetailid`,
	},
	ErrCode43009.Code(): {
		Msg:    `企业未验证主体`,
		Detail: `排查方法: -`,
	},
	ErrCode43012.Code(): {
		Msg:    `应用需配置回调url`,
		Detail: `排查方法: -`,
	},
	ErrCode44001.Code(): {
		Msg:    `多媒体文件为空`,
		Detail: `排查方法: 上传格式参考：[上传临时素材](https://work.weixin.qq.com/api/doc#10112)，确认header和body的内容正确。`,
	},
	ErrCode44004.Code(): {
		Msg:    `文本消息contDetailt参数为空`,
		Detail: `排查方法: 发文本消息contDetailt为必填参数，且不能为空`,
	},
	ErrCode45001.Code(): {
		Msg:    `多媒体文件大小超过限制`,
		Detail: `排查方法: 图片不可超过5M；音频不可超过5M；文件不可超过20M`,
	},
	ErrCode45002.Code(): {
		Msg:    `消息内容大小超过限制`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：45002);消息内容大小超过限制。确认：;1）文本消息类型：最长不超过2048个字节。;2）图文消息类型：最长不超过666k个字节`,
	},
	ErrCode45004.Code(): {
		Msg:    `应用description参数长度不符合系统限制`,
		Detail: `排查方法: 设置应用若带有description参数，则长度必须为4至120个字符`,
	},
	ErrCode45007.Code(): {
		Msg:    `语音播放时间超过限制`,
		Detail: `排查方法: 语音播放时长不能超过60秒`,
	},
	ErrCode45008.Code(): {
		Msg:    `图文消息的文章数量不符合系统限制`,
		Detail: `排查方法: 图文消息的文章数量不能超过8条`,
	},
	ErrCode45009.Code(): {
		Msg: `接口调用超过限制`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：45009);接口调用超过限制。;1) 具体频率策略，参考：[主动调用频率限制](https://opDetail.work.weixin.qq.com/api/doc#10785);2) 频率拦截时长一般与调用的限制时长相同，比如说是分钟级别的限制，则在中频率后的1分钟后自动解除。小时、天、以及月份，也是以此类推。;3) 我们对接口调用的频率限制是比较宽松的。对于接口中频率的调用，考虑以下优化：;* 接口实现时，仅系统失败需要重试。其余错误码，应该排查下调用失败原因;* 发消息应该控制合理调用，对于单个成员来说，一天收到大量的推送，体验是不好的;4) 部分频率拦截，可自助解封，访问：[频率自助解封工具](https://opDetail.work.weixin.qq.com/wwopDetail/devtool/checkCorpSpamBlock);5) 发送应用消息的频率拦截，可用api接口查询各个应用消息发送统计，访问：[查询应用消息发送统计](https://opDetail.work.weixin.qq.com/api/doc/90000/90135/92369)
//`,
	},
	ErrCode45022.Code(): {
		Msg:    `应用name参数长度不符合系统限制`,
		Detail: `排查方法: 设置应用若带有name参数，则不允许为空，且不超过32个字符`,
	},
	ErrCode45024.Code(): {
		Msg: `帐号数量超过上限`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：45024);帐号数量超过上限。请确认：;1）通讯录是否有无效或者无用的帐号，可以删除，让出额度;2）提高帐号上限，可以提交重新认证或者申请扩容
//`,
	},
	ErrCode45026.Code(): {
		Msg:    `触发删除用户数的保护`,
		Detail: `排查方法: 限制参考：[全量覆盖成员](https://work.weixin.qq.com/api/doc/90000/90139/90313#10138/全量覆盖成员)`,
	},
	ErrCode45029.Code(): {
		Msg: `回包大小超过上限`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：45029);回包大小超过上限。请确认：;1）/cgi-bin/user/list：由于通讯录组织架构庞大，建议按部门分别拉取，同时不要指定fetch_child=1。
//`,
	},
	ErrCode45032.Code(): {
		Msg:    `图文消息author参数长度超过限制`,
		Detail: `排查方法: 最长64个字节`,
	},
	ErrCode45033.Code(): {
		Msg: `接口并发调用超过限制`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：45033);接口并发调用超过限制。出现这种拦截限制，一般是开发者的程序有bug，导致对同一份资源有过高的并发且持续不断的请求，例如对一个media_id一直持续不断请求“获取临时素材”接口。
//`,
	},
	ErrCode45034.Code(): {
		Msg:    `url必须有协议头`,
		Detail: `排查方法: 在url前面加上协议头 http:// 或 https://`,
	},
	ErrCode46003.Code(): {
		Msg:    `菜单未设置`,
		Detail: `排查方法: 菜单需发布后才能获取到数据`,
	},
	ErrCode46004.Code(): {
		Msg:    `指定的用户不存在`,
		Detail: `排查方法: 需要确认指定的用户存在于通讯录中`,
	},
	ErrCode48002.Code(): {
		Msg:    `API接口无权限调用`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：48002);API接口无权限调用。请确认：;1）写通讯录接口，只能由通讯录同步助手的access_tokDetail来调用。同时需要保证通讯录同步功能是开启的。;2）通讯录同步助手的access_tokDetail，仅用于同步通讯录，不能用于发消息;3）设置应用可见范围，仅支持注册定制化安装情况，详情见：[设置授权应用可见范围](https://opDetail.work.weixin.qq.com/api/doc#14936);4）客户联系相关的接口，只能由系统应用“客户联系”，或配置到“可调用应用”列表中的自建应用的access_tokDetail来调用。;5) 小程序应用仅支持发送[小程序通知消息](https://work.weixin.qq.com/api/doc/90000/90135/90236#%E5%B0%8F%E7%A8%8B%E5%BA%8F%E9%80%9A%E7%9F%A5%E6%B6%88%E6%81%AF)，暂不支持文本、图片、语音、视频、图文等其他类型的消息。`,
	},
	ErrCode48003.Code(): {
		Msg:    `不合法的suite_id`,
		Detail: `排查方法: 确认suite_access_tokDetail由指定的suite_id生成`,
	},
	ErrCode48004.Code(): {
		Msg:    `授权关系无效`,
		Detail: `排查方法: 可能是无授权或授权已被取消`,
	},
	ErrCode48005.Code(): {
		Msg:    `API接口已废弃`,
		Detail: `排查方法: 接口已不再支持，建议改用新接口或者新方案`,
	},
	ErrCode48006.Code(): {
		Msg:    `接口权限被收回`,
		Detail: `排查方法: 由于企业长时间未使用应用，接口权限被收回，需企业管理员重新启用`,
	},
	ErrCode49004.Code(): {
		Msg:    `签名不匹配`,
		Detail: `排查方法: -`,
	},
	ErrCode49008.Code(): {
		Msg:    `群已经解散`,
		Detail: `排查方法: 群主已经解散群聊`,
	},
	ErrCode50001.Code(): {
		Msg: `redirect_url未登记可信域名`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：50001);redirect_url未登记可信域名。请确认：;1）颁发code的场景在哪个应用点击的。消费code使用的access_tokDetail是否有该应用权限。（通过[查询access_tokDetail权限](https://opDetail.work.weixin.qq.com/devtool/query)可确认）;2）secret的获取来源;* 来源于应用：url的域名，需设置到应用可信域名中。;* 来源于通讯录同步助手：仅可同步通讯录，不可用于发消息或者消费code;* 来源于第三方套件授权：套件中至少有一个应用，设置了该url域名为可信域名;* 来源于管理组：管理组配置的应用列表，至少有一个应用设置了该url域名为可信域名;3）url填写的域名，必须与设置的可信域名 **完全匹配**（包括端口号）。比如：填可信域名填qq.com，访问url域名为www.qq.com，就不匹配；或者可信域名填www.qq.com，访问url域名为www.qq.com:8008，也不匹配。
//`,
	},
	ErrCode50002.Code(): {
		Msg:    `成员不在权限范围`,
		Detail: `排查方法: 请检查应用或管理组的权限范围`,
	},
	ErrCode50003.Code(): {
		Msg: `应用已禁用`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：50003);应用禁用之后，将无法再调用api，可在”管理端-应用管理”重新启用该应用。;<img src="https://p.qpic.cn/pic_wework/1283325914/090fad19ebf740ce78b1e49f2cc0f0f8d1791c3262b8e946/0" style="width:800px;"/>
//`,
	},
	ErrCode50100.Code(): {
		Msg:    `分页查询的游标无效`,
		Detail: `排查方法: -`,
	},
	ErrCode60001.Code(): {
		Msg:    `部门长度不符合限制`,
		Detail: `排查方法: 部门名称不能为空且长度不能超过32个字`,
	},
	ErrCode60003.Code(): {
		Msg:    `部门ID不存在`,
		Detail: `排查方法: 需要确认部门ID是否有带，并且存在通讯录中`,
	},
	ErrCode60004.Code(): {
		Msg:    `父部门不存在`,
		Detail: `排查方法: 需要确认父亲部门ID是否有带，并且存在通讯录中`,
	},
	ErrCode60005.Code(): {
		Msg:    `部门下存在成员`,
		Detail: `排查方法: 不允许删除有成员的部门`,
	},
	ErrCode60006.Code(): {
		Msg:    `部门下存在子部门`,
		Detail: `排查方法: 不允许删除有子部门的部门`,
	},
	ErrCode60007.Code(): {
		Msg:    `不允许删除根部门`,
		Detail: `排查方法: -`,
	},
	ErrCode60008.Code(): {
		Msg:    `部门已存在`,
		Detail: `排查方法: 部门ID或者部门名称已存在`,
	},
	ErrCode60009.Code(): {
		Msg:    `部门名称含有非法字符`,
		Detail: `排查方法: 不能含有 \:?*“<>| 等字符`,
	},
	ErrCode60010.Code(): {
		Msg: `部门存在循环关系`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：60010);部门存在循环关系。请确认：;1）创建部门和更新部门时，指定的parDetailtid参数不能是 部门id 或者 子部门id
//`,
	},
	ErrCode60011.Code(): {
		Msg: `指定的成员/部门/标签参数无权限`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：60011);指定的成员/部门/标签参数无权限。请确认：;1) 变更通讯录接口，需要有通讯录编辑权限。;* 普通应用的secret仅有只读权限，可使用通讯录同步助手的secret同步。;2) 其它接口，需要满足配置的通讯录范围。;* 成员：通讯录同步助手access_tokDetail可指定任意成员id；应用access_tokDetail仅能指定可见范围配置的成员，以及部门/标签包含的成员（递归展开）;* 部门：通讯录同步助手access_tokDetail可指定任意部门id；应用access_tokDetail仅能指定可见范围配置的部门id(创建或移动部门，还需要具有父部门的管理权限)，标签包括的部门id，以及上述部门的子部门id;* 标签：通讯录同步助手access_tokDetail可指定超级管理组及通讯录同步助手创建的标签；应用access_tokDetail仅能由应用API创建的标签
//`,
	},
	ErrCode60012.Code(): {
		Msg:    `不允许删除默认应用`,
		Detail: `排查方法: 默认应用的id为0`,
	},
	ErrCode60020.Code(): {
		Msg: `访问ip不在白名单之中`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：60020);访问ip不在白名单之中。请确认：;1）请确认访问ip是否在服务商白名单IP列表。;登录 [服务商管理后台](https://opDetail.work.weixin.qq.com/wwopDetail/login)，在“服务商信息” - “基本信息” - “IP白名单”配置
//`,
	},
	ErrCode60021.Code(): {
		Msg:    `userid不在应用可见范围内`,
		Detail: `排查方法: -`,
	},
	ErrCode60028.Code(): {
		Msg:    `不允许修改第三方应用的主页 URL`,
		Detail: `排查方法: 第三方应用类型，不允许通过接口修改该应用的主页 URL`,
	},
	ErrCode60102.Code(): {
		Msg:    `UserID已存在`,
		Detail: `排查方法: -`,
	},
	ErrCode60103.Code(): {
		Msg:    `手机号码不合法`,
		Detail: `排查方法: 长度不超过32位，字符仅支持数字，加号和减号`,
	},
	ErrCode60104.Code(): {
		Msg:    `手机号码已存在`,
		Detail: `排查方法: 同一个企业内，成员的手机号不能重复。建议更换手机号，或者更新已有的手机记录。`,
	},
	ErrCode60105.Code(): {
		Msg:    `邮箱不合法`,
		Detail: `排查方法: 长度不超过64位，且为有效的email格式`,
	},
	ErrCode60106.Code(): {
		Msg:    `邮箱已存在`,
		Detail: `排查方法: 同一个企业内，成员的邮箱不能重复。建议更换邮箱，或者更新已有的邮箱记录。`,
	},
	ErrCode60107.Code(): {
		Msg:    `微信号不合法`,
		Detail: `排查方法: 微信号格式由字母、数字、”-“、”_“组成，长度为 3-20 字节，首字符必须是字母或”-“或”_“`,
	},
	ErrCode60110.Code(): {
		Msg:    `用户所属部门数量超过限制`,
		Detail: `排查方法: 用户同时归属部门不超过20个`,
	},
	ErrCode60111.Code(): {
		Msg:    `UserID不存在`,
		Detail: `排查方法: UserID参数为空，或者不存在通讯录中`,
	},
	ErrCode60112.Code(): {
		Msg:    `成员name参数不合法`,
		Detail: `排查方法: 不能为空，且不能超过64字符`,
	},
	ErrCode60123.Code(): {
		Msg:    `无效的部门id`,
		Detail: `排查方法: 部门不存在通讯录中`,
	},
	ErrCode60124.Code(): {
		Msg:    `无效的父部门id`,
		Detail: `排查方法: 父部门不存在通讯录中`,
	},
	ErrCode60125.Code(): {
		Msg:    `非法部门名字`,
		Detail: `排查方法: 不能为空，且不能超过64字节，且不能含有\:*?”<>|等字符`,
	},
	ErrCode60127.Code(): {
		Msg:    `缺少departmDetailt参数`,
		Detail: `排查方法: -`,
	},
	ErrCode60129.Code(): {
		Msg:    `成员手机和邮箱都为空`,
		Detail: `排查方法: 成员手机和邮箱至少有个非空`,
	},
	ErrCode60132.Code(): {
		Msg:    `is_leader_in_dept和departmDetailt的元素个数不一致`,
		Detail: `排查方法: -`,
	},
	ErrCode60136.Code(): {
		Msg:    `记录不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode60137.Code(): {
		Msg:    `家长手机号重复`,
		Detail: `排查方法: 同一个家校通讯录中，家长的手机号不能重复。建议更换手机号，或者更新已有的手机记录。`,
	},
	ErrCode60203.Code(): {
		Msg:    `不合法的模版ID`,
		Detail: `排查方法: -`,
	},
	ErrCode60204.Code(): {
		Msg:    `模版状态不可用`,
		Detail: `排查方法: -`,
	},
	ErrCode60205.Code(): {
		Msg:    `模版关键词不匹配`,
		Detail: `排查方法: -`,
	},
	ErrCode60206.Code(): {
		Msg:    `该种类型的消息只支持第三方独立应用使用`,
		Detail: `排查方法: -`,
	},
	ErrCode60207.Code(): {
		Msg:    `第三方独立应用只允许发送模板消息`,
		Detail: `排查方法: -`,
	},
	ErrCode60208.Code(): {
		Msg:    `第三方独立应用不支持指定@all，不支持参数toparty和totag`,
		Detail: `排查方法: -`,
	},
	ErrCode60209.Code(): {
		Msg:    `缺少操作者vid`,
		Detail: `排查方法: -`,
	},
	ErrCode60210.Code(): {
		Msg:    `选择成员列表为空`,
		Detail: `排查方法: -`,
	},
	ErrCode60211.Code(): {
		Msg:    `SelectedTicket为空`,
		Detail: `排查方法: -`,
	},
	ErrCode60214.Code(): {
		Msg:    `仅支持第三方应用调用`,
		Detail: `排查方法: -`,
	},
	ErrCode60215.Code(): {
		Msg:    `传入SelectedTicket数量超过最大限制（10个）`,
		Detail: `排查方法: -`,
	},
	ErrCode60217.Code(): {
		Msg:    `当前操作者无权限，操作者需要授权或者在可见范围内`,
		Detail: `排查方法: -`,
	},
	ErrCode60218.Code(): {
		Msg:    `仅支持成员授权模式的应用可调用`,
		Detail: `排查方法: -`,
	},
	ErrCode60219.Code(): {
		Msg:    `消费SelectedTicket和创建SelectedTicket的应用appid不匹配`,
		Detail: `排查方法: -`,
	},
	ErrCode60220.Code(): {
		Msg:    `缺少corpappid`,
		Detail: `排查方法: -`,
	},
	ErrCode60221.Code(): {
		Msg:    `opDetail_userid对应的服务商不是当前服务商`,
		Detail: `排查方法: -`,
	},
	ErrCode60222.Code(): {
		Msg:    `非法SelectedTicket`,
		Detail: `排查方法: -`,
	},
	ErrCode60223.Code(): {
		Msg:    `非法BundleId`,
		Detail: `排查方法: -`,
	},
	ErrCode60224.Code(): {
		Msg:    `非法PackagDetailame`,
		Detail: `排查方法: -`,
	},
	ErrCode60225.Code(): {
		Msg:    `当前操作者并非SelectedTicket相关人，不能创建群聊`,
		Detail: `排查方法: -`,
	},
	ErrCode60226.Code(): {
		Msg:    `选人数量超过最大限制（2000）`,
		Detail: `排查方法: -`,
	},
	ErrCode60227.Code(): {
		Msg:    `缺少ServiceCorpid`,
		Detail: `排查方法: -`,
	},
	ErrCode65000.Code(): {
		Msg:    `学校已经迁移`,
		Detail: `排查方法: -`,
	},
	ErrCode65001.Code(): {
		Msg:    `无效的关注模式`,
		Detail: `排查方法: -`,
	},
	ErrCode65002.Code(): {
		Msg:    `导入家长信息数量过多`,
		Detail: `排查方法: 批量导入家长每次最多1000个`,
	},
	ErrCode65003.Code(): {
		Msg:    `学校尚未迁移`,
		Detail: `排查方法: -`,
	},
	ErrCode65004.Code(): {
		Msg:    `组织架构不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode65005.Code(): {
		Msg:    `无效的同步模式`,
		Detail: `排查方法: -`,
	},
	ErrCode65006.Code(): {
		Msg:    `无效的管理员类型`,
		Detail: `排查方法: -`,
	},
	ErrCode65007.Code(): {
		Msg:    `无效的家校部门类型`,
		Detail: `排查方法: -`,
	},
	ErrCode65008.Code(): {
		Msg:    `无效的入学年份`,
		Detail: `排查方法: -`,
	},
	ErrCode65009.Code(): {
		Msg:    `无效的标准年级类型`,
		Detail: `排查方法: -`,
	},
	ErrCode65010.Code(): {
		Msg:    `此userid并不是学生`,
		Detail: `排查方法: -`,
	},
	ErrCode65011.Code(): {
		Msg:    `家长userid数量超过限制`,
		Detail: `排查方法: 每次最多批量处理100个家长`,
	},
	ErrCode65012.Code(): {
		Msg:    `学生userid数量超过限制`,
		Detail: `排查方法: 每次最多批量处理10个学生`,
	},
	ErrCode65013.Code(): {
		Msg:    `学生已有家长`,
		Detail: `排查方法: -`,
	},
	ErrCode65014.Code(): {
		Msg:    `非学校企业`,
		Detail: `排查方法: -`,
	},
	ErrCode65015.Code(): {
		Msg:    `父部门类型不匹配`,
		Detail: `排查方法: 添加学校部门，需满足层级关机，班级需要以年级为父部门`,
	},
	ErrCode65018.Code(): {
		Msg:    `家长人数达到上限`,
		Detail: `排查方法: 未验证的学校\企业最多可添加2000名家长，验证过的学校\企业最多可添加20000名家长`,
	},
	ErrCode660001.Code(): {
		Msg:    `无效的商户号`,
		Detail: `排查方法: 请检查商户号是否正确`,
	},
	ErrCode660002.Code(): {
		Msg:    `无效的企业收款人id`,
		Detail: `排查方法: 请检查payee_userid是否正确`,
	},
	ErrCode660003.Code(): {
		Msg:    `userid不在应用的可见范围`,
		Detail: `排查方法: -`,
	},
	ErrCode660004.Code(): {
		Msg:    `partyid不在应用的可见范围`,
		Detail: `排查方法: -`,
	},
	ErrCode660005.Code(): {
		Msg:    `tagid不在应用的可见范围`,
		Detail: `排查方法: -`,
	},
	ErrCode660006.Code(): {
		Msg:    `找不到该商户号`,
		Detail: `排查方法: -`,
	},
	ErrCode660007.Code(): {
		Msg:    `申请已经存在`,
		Detail: `排查方法: 不需要重复申请`,
	},
	ErrCode660008.Code(): {
		Msg:    `商户号已经绑定`,
		Detail: `排查方法: 不需要重新提交申请`,
	},
	ErrCode660009.Code(): {
		Msg:    `商户号主体和商户主体不一致`,
		Detail: `排查方法: -`,
	},
	ErrCode660010.Code(): {
		Msg:    `超过商户号绑定数量限制`,
		Detail: `排查方法: -`,
	},
	ErrCode660011.Code(): {
		Msg:    `商户号未绑定`,
		Detail: `排查方法: -`,
	},
	ErrCode670001.Code(): {
		Msg:    `应用不在共享范围`,
		Detail: `排查方法: -`,
	},
	ErrCode72023.Code(): {
		Msg: `发票已被其他公众号锁定`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：72023);一般为发票已进入后续报销流程，报销企业公众号/企业微信/App锁定了发票。
//`,
	},
	ErrCode72024.Code(): {
		Msg:    `发票状态错误`,
		Detail: `排查方法: reimburse_status状态错误，参考：[更新发票状态](https://work.weixin.qq.com/api/doc/90000/90139/90313#11633)`,
	},
	ErrCode72037.Code(): {
		Msg:    `存在发票不属于该用户`,
		Detail: `排查方法: 只能批量更新该opDetailid的发票，参考：[批量更新发票状态](https://work.weixin.qq.com/api/doc/90000/90139/90313#11634)`,
	},
	ErrCode80001.Code(): {
		Msg:    `可信域名不正确，或者无ICP备案`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：80001);可信域名不正确，未校验域名所有权归属或者可信域名没有ICP备案。请确认：;1）可信域名，只支持全域名匹配，无法通过配置父域来让所有子域都成为可信域名。;2）可信域名，不支持IP地址、端口号及短链域名。;3）如果确认域名已经通过ICP备案，但依然提示这个错误，请尝试重新设置。`,
	},
	ErrCode81001.Code(): {
		Msg:    `部门下的结点数超过限制（3W）`,
		Detail: `排查方法: -`,
	},
	ErrCode81002.Code(): {
		Msg:    `部门最多15层`,
		Detail: `排查方法: -`,
	},
	ErrCode81003.Code(): {
		Msg:    `标签下节点个数超过30000个`,
		Detail: `排查方法: -`,
	},
	ErrCode81011.Code(): {
		Msg: `无权限操作标签`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：81011);无权限操作标签。请确认：;1）除了通讯录同步助手和通讯录应用，其他应用和管理组都只能操作自己创建的标签。;2）通讯录同步助手或者通讯录应用，除了能管理自己的标签，还能操作超级管理组创建的标签。
//`,
	},
	ErrCode81012.Code(): {
		Msg:    `缺失可见范围`,
		Detail: `排查方法: 请求没有填写UserID、部门ID、标签ID`,
	},
	ErrCode81013.Code(): {
		Msg: `UserID、部门ID、标签ID全部非法或无权限`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：81013);UserID、部门ID、标签ID全部非法或无权限。一般有以下两种原因：;1）成员、部门或标签已被删除，此种情况需要调整调用接口的接收人参数。;2）成员、部门或标签被移出应用的可见范围，可在管理端将接收人添加到应用的可见范围内。;<img src="https://p.qpic.cn/pic_wework/1283325914/4e48eb5ea14e98977a1687c6cceab0d61ad57eea9af062f0/0" style="width:800px;"/>
//`,
	},
	ErrCode81014.Code(): {
		Msg:    `标签添加成员，单次添加user或party过多`,
		Detail: `排查方法: -`,
	},
	ErrCode81015.Code(): {
		Msg:    `邮箱域名需要跟企业邮箱域名一致`,
		Detail: `排查方法: -`,
	},
	ErrCode81016.Code(): {
		Msg:    `logined_userid字段缺失`,
		Detail: `排查方法: -`,
	},
	ErrCode81017.Code(): {
		Msg:    `请求个数超过限制`,
		Detail: `排查方法: -`,
	},
	ErrCode81018.Code(): {
		Msg:    `该服务商可获取名字数量配额不足`,
		Detail: `排查方法: -`,
	},
	ErrCode81019.Code(): {
		Msg:    `items数组成员缺少id字段`,
		Detail: `排查方法: -`,
	},
	ErrCode81020.Code(): {
		Msg:    `items数组成员缺少type字段`,
		Detail: `排查方法: -`,
	},
	ErrCode81021.Code(): {
		Msg:    `items数组成员的type字段不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode82001.Code(): {
		Msg: `指定的成员/部门/标签全部为空`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：82001);指定的成员/部门/标签全部为空。请确认：;参数是否有传递，且至少有一个参数非空。
//`,
	},
	ErrCode82002.Code(): {
		Msg:    `不合法的PartyID列表长度`,
		Detail: `排查方法: 发消息，单次不能超过100个部门`,
	},
	ErrCode82003.Code(): {
		Msg:    `不合法的TagID列表长度`,
		Detail: `排查方法: 发消息，单次不能超过100个标签`,
	},
	ErrCode82004.Code(): {
		Msg:    `不合法的消息内容`,
		Detail: `排查方法: 消息内容中可能存在使客户端crash的内容`,
	},
	ErrCode84014.Code(): {
		Msg: `成员票据过期`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：84014);成员票据过期。确认：;1）user_ticket 有时效性，有效时长由expires_in指定。参考接口：[根据code获取成员信息](https://opDetail.work.weixin.qq.com/api/doc#10028/根据code获取成员信息);2）若需再次获取用户详情，需要用户重新点击链接后，根据新的code获取新的user_ticket
//`,
	},
	ErrCode84015.Code(): {
		Msg:    `成员票据无效`,
		Detail: `排查方法: 确认user_ticket参数来源是否正确。参考接口：[根据code获取成员信息](https://work.weixin.qq.com/api/doc/90000/90139/90313#10028/根据code获取成员信息)`,
	},
	ErrCode84019.Code(): {
		Msg:    `缺少templateid参数`,
		Detail: `排查方法: -`,
	},
	ErrCode84020.Code(): {
		Msg:    `templateid不存在`,
		Detail: `排查方法: 确认参数是否有带，并且已创建`,
	},
	ErrCode84021.Code(): {
		Msg:    `缺少register_code参数`,
		Detail: `排查方法: -`,
	},
	ErrCode84022.Code(): {
		Msg:    `无效的register_code参数`,
		Detail: `排查方法: -`,
	},
	ErrCode84023.Code(): {
		Msg:    `不允许调用设置通讯录同步完成接口`,
		Detail: `排查方法: -`,
	},
	ErrCode84024.Code(): {
		Msg:    `无注册信息`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：84024);无注册信息。可能是以下两种情况：;1）注册流程未完成。;2）注册成功已超过24小时。`,
	},
	ErrCode84025.Code(): {
		Msg:    `不符合的state参数`,
		Detail: `排查方法: 必须是[a-zA-Z0-9]的参数值，长度不可超过128个字节`,
	},
	ErrCode84052.Code(): {
		Msg:    `缺少caller参数`,
		Detail: `排查方法: -`,
	},
	ErrCode84053.Code(): {
		Msg:    `缺少callee参数`,
		Detail: `排查方法: -`,
	},
	ErrCode84054.Code(): {
		Msg:    `缺少auth_corpid参数`,
		Detail: `排查方法: -`,
	},
	ErrCode84055.Code(): {
		Msg:    `超过拨打公费电话频率`,
		Detail: `排查方法: 同一个客服5秒内只能调用api拨打一次公费电话`,
	},
	ErrCode84056.Code(): {
		Msg:    `被拨打用户安装应用时未授权拨打公费电话权限`,
		Detail: `排查方法: -`,
	},
	ErrCode84057.Code(): {
		Msg:    `公费电话余额不足`,
		Detail: `排查方法: -`,
	},
	ErrCode84058.Code(): {
		Msg:    `caller 呼叫号码不支持`,
		Detail: `排查方法: -`,
	},
	ErrCode84059.Code(): {
		Msg:    `号码非法`,
		Detail: `排查方法: -`,
	},
	ErrCode84060.Code(): {
		Msg:    `callee 呼叫号码不支持`,
		Detail: `排查方法: -`,
	},
	ErrCode84061.Code(): {
		Msg:    `不存在外部联系人的关系`,
		Detail: `排查方法: -`,
	},
	ErrCode84062.Code(): {
		Msg:    `未开启公费电话应用`,
		Detail: `排查方法: -`,
	},
	ErrCode84063.Code(): {
		Msg:    `caller不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode84064.Code(): {
		Msg:    `callee不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode84065.Code(): {
		Msg:    `caller跟callee电话号码一致`,
		Detail: `排查方法: 不允许自己拨打给自己`,
	},
	ErrCode84066.Code(): {
		Msg:    `服务商拨打次数超过限制`,
		Detail: `排查方法: 单个企业管理员，在一天（以上午10:00为起始时间）内，对应单个服务商，只能被呼叫【4】次。`,
	},
	ErrCode84067.Code(): {
		Msg:    `管理员收到的服务商公费电话个数超过限制`,
		Detail: `排查方法: 单个企业管理员，在一天（以上午10:00为起始时间）内，一共只能被【3】个服务商成功呼叫。`,
	},
	ErrCode84069.Code(): {
		Msg:    `拨打方被限制拨打公费电话`,
		Detail: `排查方法: -`,
	},
	ErrCode84070.Code(): {
		Msg:    `不支持的电话号码`,
		Detail: `排查方法: 拨打方或者被拨打方电话号码不支持`,
	},
	ErrCode84071.Code(): {
		Msg:    `不合法的外部联系人授权码`,
		Detail: `排查方法: 非法或者已经消费过`,
	},
	ErrCode84072.Code(): {
		Msg:    `应用未配置客服`,
		Detail: `排查方法: -`,
	},
	ErrCode84073.Code(): {
		Msg:    `客服userid不在应用配置的客服列表中`,
		Detail: `排查方法: -`,
	},
	ErrCode84074.Code(): {
		Msg:    `没有外部联系人权限`,
		Detail: `排查方法: -`,
	},
	ErrCode84075.Code(): {
		Msg:    `不合法或过期的authcode`,
		Detail: `排查方法: -`,
	},
	ErrCode84076.Code(): {
		Msg:    `缺失authcode`,
		Detail: `排查方法: -`,
	},
	ErrCode84077.Code(): {
		Msg:    `订单价格过高，无法受理`,
		Detail: `排查方法: -`,
	},
	ErrCode84078.Code(): {
		Msg:    `购买人数不正确`,
		Detail: `排查方法: -`,
	},
	ErrCode84079.Code(): {
		Msg:    `价格策略不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode84080.Code(): {
		Msg:    `订单不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode84081.Code(): {
		Msg:    `存在未支付订单`,
		Detail: `排查方法: -`,
	},
	ErrCode84082.Code(): {
		Msg:    `存在申请退款中的订单`,
		Detail: `排查方法: -`,
	},
	ErrCode84083.Code(): {
		Msg:    `非服务人员`,
		Detail: `排查方法: -`,
	},
	ErrCode84084.Code(): {
		Msg:    `非跟进用户`,
		Detail: `排查方法: -`,
	},
	ErrCode84085.Code(): {
		Msg:    `应用已下架`,
		Detail: `排查方法: -`,
	},
	ErrCode84086.Code(): {
		Msg:    `订单人数超过可购买最大人数`,
		Detail: `排查方法: -`,
	},
	ErrCode84087.Code(): {
		Msg:    `打开订单支付前禁止关闭订单`,
		Detail: `排查方法: -`,
	},
	ErrCode84088.Code(): {
		Msg:    `禁止关闭已支付的订单`,
		Detail: `排查方法: -`,
	},
	ErrCode84089.Code(): {
		Msg:    `订单已支付`,
		Detail: `排查方法: -`,
	},
	ErrCode84090.Code(): {
		Msg:    `缺失user_ticket`,
		Detail: `排查方法: -`,
	},
	ErrCode84091.Code(): {
		Msg:    `订单价格不可低于下限`,
		Detail: `排查方法: -`,
	},
	ErrCode84092.Code(): {
		Msg:    `无法发起代下单操作`,
		Detail: `排查方法: -`,
	},
	ErrCode84093.Code(): {
		Msg:    `代理关系已占用，无法代下单`,
		Detail: `排查方法: -`,
	},
	ErrCode84094.Code(): {
		Msg:    `该应用未配置代理分润规则，请先联系应用服务商处理`,
		Detail: `排查方法: -`,
	},
	ErrCode84095.Code(): {
		Msg:    `免费试用版，无法扩容`,
		Detail: `排查方法: -`,
	},
	ErrCode84096.Code(): {
		Msg:    `免费试用版，无法续期`,
		Detail: `排查方法: -`,
	},
	ErrCode84097.Code(): {
		Msg:    `当前企业有未处理订单`,
		Detail: `排查方法: -`,
	},
	ErrCode84098.Code(): {
		Msg:    `固定总量，无法扩容`,
		Detail: `排查方法: -`,
	},
	ErrCode84099.Code(): {
		Msg:    `非购买状态，无法扩容`,
		Detail: `排查方法: -`,
	},
	ErrCode84100.Code(): {
		Msg:    `未购买过此应用，无法续期`,
		Detail: `排查方法: -`,
	},
	ErrCode84101.Code(): {
		Msg:    `企业已试用付费版本，无法全新购买`,
		Detail: `排查方法: -`,
	},
	ErrCode84102.Code(): {
		Msg:    `企业当前应用状态已过期，无法扩容`,
		Detail: `排查方法: -`,
	},
	ErrCode84103.Code(): {
		Msg:    `仅可修改未支付订单`,
		Detail: `排查方法: -`,
	},
	ErrCode84104.Code(): {
		Msg:    `订单已支付，无法修改`,
		Detail: `排查方法: -`,
	},
	ErrCode84105.Code(): {
		Msg:    `订单已被取消，无法修改`,
		Detail: `排查方法: -`,
	},
	ErrCode84106.Code(): {
		Msg:    `企业含有该应用的待支付订单，无法代下单`,
		Detail: `排查方法: -`,
	},
	ErrCode84107.Code(): {
		Msg:    `企业含有该应用的退款中订单，无法代下单`,
		Detail: `排查方法: -`,
	},
	ErrCode84108.Code(): {
		Msg:    `企业含有该应用的待生效订单，无法代下单`,
		Detail: `排查方法: -`,
	},
	ErrCode84109.Code(): {
		Msg:    `订单定价不能未0`,
		Detail: `排查方法: -`,
	},
	ErrCode84110.Code(): {
		Msg:    `新安装应用不在试用状态，无法升级为付费版`,
		Detail: `排查方法: -`,
	},
	ErrCode84111.Code(): {
		Msg:    `无足够可用优惠券`,
		Detail: `排查方法: -`,
	},
	ErrCode84112.Code(): {
		Msg:    `无法关闭未支付订单`,
		Detail: `排查方法: -`,
	},
	ErrCode84113.Code(): {
		Msg:    `无付费信息`,
		Detail: `排查方法: -`,
	},
	ErrCode84114.Code(): {
		Msg:    `虚拟版本不支持下单`,
		Detail: `排查方法: -`,
	},
	ErrCode84115.Code(): {
		Msg:    `虚拟版本不支持扩容`,
		Detail: `排查方法: -`,
	},
	ErrCode84116.Code(): {
		Msg:    `虚拟版本不支持续期`,
		Detail: `排查方法: -`,
	},
	ErrCode84117.Code(): {
		Msg:    `在虚拟正式版期内不能扩容`,
		Detail: `排查方法: -`,
	},
	ErrCode84118.Code(): {
		Msg:    `虚拟正式版期内不能变更版本`,
		Detail: `排查方法: -`,
	},
	ErrCode84119.Code(): {
		Msg:    `当前企业未报备，无法进行代下单`,
		Detail: `排查方法: -`,
	},
	ErrCode84120.Code(): {
		Msg:    `当前应用版本已删除`,
		Detail: `排查方法: -`,
	},
	ErrCode84121.Code(): {
		Msg:    `应用版本已删除，无法扩容`,
		Detail: `排查方法: -`,
	},
	ErrCode84122.Code(): {
		Msg:    `应用版本已删除，无法续期`,
		Detail: `排查方法: -`,
	},
	ErrCode84123.Code(): {
		Msg:    `非虚拟版本，无法升级`,
		Detail: `排查方法: -`,
	},
	ErrCode84124.Code(): {
		Msg:    `非行业方案订单，不能添加部分应用版本的订单`,
		Detail: `排查方法: -`,
	},
	ErrCode84125.Code(): {
		Msg:    `购买人数不能少于最少购买人数`,
		Detail: `排查方法: -`,
	},
	ErrCode84126.Code(): {
		Msg:    `购买人数不能多于最大购买人数`,
		Detail: `排查方法: -`,
	},
	ErrCode84127.Code(): {
		Msg:    `无应用管理权限`,
		Detail: `排查方法: -`,
	},
	ErrCode84128.Code(): {
		Msg:    `无该行业方案下全部应用的管理权限`,
		Detail: `排查方法: -`,
	},
	ErrCode84129.Code(): {
		Msg:    `付费策略已被删除，无法下单`,
		Detail: `排查方法: -`,
	},
	ErrCode84130.Code(): {
		Msg:    `订单生效时间不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode84200.Code(): {
		Msg:    `文件转译解析错误`,
		Detail: `排查方法: 只支持utf8文件转译，可能是不支持的文件类型或者格式`,
	},
	ErrCode85002.Code(): {
		Msg:    `包含不合法的词语`,
		Detail: `排查方法: -`,
	},
	ErrCode85004.Code(): {
		Msg:    `每企业每个月设置的可信域名不可超过20个`,
		Detail: `排查方法: -`,
	},
	ErrCode85005.Code(): {
		Msg:    `可信域名未通过所有权校验`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：85005);域名未通过所有权校验，仅oauth2生效，jssdk功能将受限，请根据调用者身份按以下不同方式完成校验：;1）若调用者是企业应用，请登录企业微信管理端，进入应用详情，按照指引完成域名的所有权校验。;2）若调用者是第三方服务，请登录企业微信服务管理端，进入第三方应用详情，按照指引完成域名的所有权校验。`,
	},
	ErrCode86001.Code(): {
		Msg:    `参数 chatid 不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode86003.Code(): {
		Msg:    `参数 chatid 不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode86004.Code(): {
		Msg:    `参数 群名不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode86005.Code(): {
		Msg:    `参数 群主不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode86006.Code(): {
		Msg:    `群成员数过多或过少`,
		Detail: `排查方法: -`,
	},
	ErrCode86007.Code(): {
		Msg:    `不合法的群成员`,
		Detail: `排查方法: -`,
	},
	ErrCode86008.Code(): {
		Msg:    `非法操作非自己创建的群`,
		Detail: `排查方法: -`,
	},
	ErrCode86101.Code(): {
		Msg:    `仅群主才有操作权限`,
		Detail: `排查方法: -`,
	},
	ErrCode86201.Code(): {
		Msg:    `参数 需要chatid`,
		Detail: `排查方法: -`,
	},
	ErrCode86202.Code(): {
		Msg:    `参数 需要群名`,
		Detail: `排查方法: -`,
	},
	ErrCode86203.Code(): {
		Msg:    `参数 需要群主`,
		Detail: `排查方法: -`,
	},
	ErrCode86204.Code(): {
		Msg:    `参数 需要群成员`,
		Detail: `排查方法: -`,
	},
	ErrCode86205.Code(): {
		Msg:    `参数 字符串chatid过长`,
		Detail: `排查方法: -`,
	},
	ErrCode86206.Code(): {
		Msg:    `参数 数字chatid过大`,
		Detail: `排查方法: -`,
	},
	ErrCode86207.Code(): {
		Msg:    `群主不在群成员列表`,
		Detail: `排查方法: -`,
	},
	ErrCode86214.Code(): {
		Msg:    `群发类型不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode86215.Code(): {
		Msg:    `会话ID已经存在`,
		Detail: `排查方法: -`,
	},
	ErrCode86216.Code(): {
		Msg:    `存在非法会话成员ID`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：86216);存在非法会话成员ID。确认：;1）添加会话成员时，指定的成员ID不存在通讯录;2）删除会话成员时，指定的成员ID不存在于会话中`,
	},
	ErrCode86217.Code(): {
		Msg:    `会话发送者不在会话成员列表中`,
		Detail: `排查方法: 会话的发送者，必须是会话的成员列表之一`,
	},
	ErrCode86220.Code(): {
		Msg:    `指定的会话参数不合法`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：86220);指定的会话参数不合法。请确认：;1）参数 type 只能指定 single/group;2）参数 msgtype 只能指定 text/image/file/voice/link`,
	},
	ErrCode86224.Code(): {
		Msg:    `不是受限群，不允许使用该接口`,
		Detail: `排查方法: -`,
	},
	ErrCode90001.Code(): {
		Msg:    `未认证摇一摇周边`,
		Detail: `排查方法: -`,
	},
	ErrCode90002.Code(): {
		Msg:    `缺少摇一摇周边ticket参数`,
		Detail: `排查方法: -`,
	},
	ErrCode90003.Code(): {
		Msg:    `摇一摇周边ticket参数不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode90100.Code(): {
		Msg:    `非法的对外属性类型`,
		Detail: `排查方法: -`,
	},
	ErrCode90101.Code(): {
		Msg:    `对外属性：文本类型长度不合法`,
		Detail: `排查方法: 文本长度不可超过12个UTF8字符`,
	},
	ErrCode90102.Code(): {
		Msg:    `对外属性：网页类型标题长度不合法`,
		Detail: `排查方法: 标题长度不可超过12个UTF8字符`,
	},
	ErrCode90103.Code(): {
		Msg:    `对外属性：网页url不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode90104.Code(): {
		Msg:    `对外属性：小程序类型标题长度不合法`,
		Detail: `排查方法: 标题长度不可超过12个UTF8字符`,
	},
	ErrCode90105.Code(): {
		Msg:    `对外属性：小程序类型pagepath不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode90106.Code(): {
		Msg:    `对外属性：请求参数不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode90200.Code(): {
		Msg:    `缺少小程序appid参数`,
		Detail: `排查方法: -`,
	},
	ErrCode90201.Code(): {
		Msg:    `小程序通知的contDetailt_item个数超过限制`,
		Detail: `排查方法: item个数不能超过10个`,
	},
	ErrCode90202.Code(): {
		Msg:    `小程序通知中的key长度不合法`,
		Detail: `排查方法: 不能为空或超过10个汉字`,
	},
	ErrCode90203.Code(): {
		Msg:    `小程序通知中的value长度不合法`,
		Detail: `排查方法: 不能为空或超过30个汉字`,
	},
	ErrCode90204.Code(): {
		Msg:    `小程序通知中的page参数不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode90206.Code(): {
		Msg:    `小程序未关联到企业中`,
		Detail: `排查方法: -`,
	},
	ErrCode90207.Code(): {
		Msg:    `不合法的小程序appid`,
		Detail: `排查方法: -`,
	},
	ErrCode90208.Code(): {
		Msg:    `小程序appid不匹配`,
		Detail: `排查方法: -`,
	},
	ErrCode90300.Code(): {
		Msg:    `orderid 不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode90302.Code(): {
		Msg:    `付费应用已过期`,
		Detail: `排查方法: -`,
	},
	ErrCode90303.Code(): {
		Msg:    `付费应用超过最大使用人数`,
		Detail: `排查方法: -`,
	},
	ErrCode90304.Code(): {
		Msg:    `订单中心服务异常，请稍后重试`,
		Detail: `排查方法: -`,
	},
	ErrCode90305.Code(): {
		Msg:    `参数错误，errmsg中有提示具体哪个参数有问题`,
		Detail: `排查方法: -`,
	},
	ErrCode90306.Code(): {
		Msg:    `商户设置不合法，详情请见errmsg`,
		Detail: `排查方法: -`,
	},
	ErrCode90307.Code(): {
		Msg:    `登录态过期`,
		Detail: `排查方法: -`,
	},
	ErrCode90308.Code(): {
		Msg:    `在开启IP鉴权的前提下，识别为无效的请求IP`,
		Detail: `排查方法: -`,
	},
	ErrCode90309.Code(): {
		Msg:    `订单已经存在，请勿重复下单`,
		Detail: `排查方法: -`,
	},
	ErrCode90310.Code(): {
		Msg:    `找不到订单`,
		Detail: `排查方法: -`,
	},
	ErrCode90311.Code(): {
		Msg:    `关单失败, 可能原因：该单并没被拉起支付页面; 已经关单；已经支付；渠道失败；单处于保护状态；等等`,
		Detail: `排查方法: -`,
	},
	ErrCode90312.Code(): {
		Msg:    `退款请求失败, 详情请看errmsg`,
		Detail: `排查方法: -`,
	},
	ErrCode90313.Code(): {
		Msg:    `退款调用频率限制，超过规定的阈值`,
		Detail: `排查方法: -`,
	},
	ErrCode90314.Code(): {
		Msg:    `订单状态错误，可能未支付，或者当前状态操作受限`,
		Detail: `排查方法: -`,
	},
	ErrCode90315.Code(): {
		Msg:    `退款请求失败，主键冲突，请核实退款refund_id是否已使用`,
		Detail: `排查方法: -`,
	},
	ErrCode90316.Code(): {
		Msg:    `退款原因编号不对`,
		Detail: `排查方法: -`,
	},
	ErrCode90317.Code(): {
		Msg:    `尚未注册成为供应商`,
		Detail: `排查方法: -`,
	},
	ErrCode90318.Code(): {
		Msg:    `参数nonce_str 为空或者重复，判定为重放攻击`,
		Detail: `排查方法: -`,
	},
	ErrCode90319.Code(): {
		Msg:    `时间戳为空或者与系统时间间隔太大`,
		Detail: `排查方法: -`,
	},
	ErrCode90320.Code(): {
		Msg:    `订单tokDetail无效`,
		Detail: `排查方法: -`,
	},
	ErrCode90321.Code(): {
		Msg:    `订单tokDetail已过有效时间`,
		Detail: `排查方法: -`,
	},
	ErrCode90322.Code(): {
		Msg:    `旧套件（包含多个应用的套件）不支持支付系统`,
		Detail: `排查方法: -`,
	},
	ErrCode90323.Code(): {
		Msg:    `单价超过限额`,
		Detail: `排查方法: -`,
	},
	ErrCode90324.Code(): {
		Msg:    `商品数量超过限额`,
		Detail: `排查方法: -`,
	},
	ErrCode90325.Code(): {
		Msg:    `预支单已经存在`,
		Detail: `排查方法: -`,
	},
	ErrCode90326.Code(): {
		Msg:    `预支单单号非法`,
		Detail: `排查方法: -`,
	},
	ErrCode90327.Code(): {
		Msg:    `该预支单已经结算下单`,
		Detail: `排查方法: -`,
	},
	ErrCode90328.Code(): {
		Msg:    `结算下单失败，详情请看errmsg`,
		Detail: `排查方法: -`,
	},
	ErrCode90329.Code(): {
		Msg:    `该订单号已经被预支单占用`,
		Detail: `排查方法: -`,
	},
	ErrCode90330.Code(): {
		Msg:    `创建供应商失败`,
		Detail: `排查方法: -`,
	},
	ErrCode90331.Code(): {
		Msg:    `更新供应商失败`,
		Detail: `排查方法: -`,
	},
	ErrCode90332.Code(): {
		Msg:    `还没签署合同`,
		Detail: `排查方法: -`,
	},
	ErrCode90333.Code(): {
		Msg:    `创建合同失败`,
		Detail: `排查方法: -`,
	},
	ErrCode90338.Code(): {
		Msg:    `已经过了可退款期限`,
		Detail: `排查方法: -`,
	},
	ErrCode90339.Code(): {
		Msg:    `供应商主体名包含非法字符`,
		Detail: `排查方法: -`,
	},
	ErrCode90340.Code(): {
		Msg:    `创建客户失败，可能信息真实性校验失败`,
		Detail: `排查方法: -`,
	},
	ErrCode90341.Code(): {
		Msg:    `退款金额大于付款金额`,
		Detail: `排查方法: -`,
	},
	ErrCode90342.Code(): {
		Msg:    `退款金额超过账户余额`,
		Detail: `排查方法: -`,
	},
	ErrCode90343.Code(): {
		Msg:    `退款单号已经存在`,
		Detail: `排查方法: -`,
	},
	ErrCode90344.Code(): {
		Msg:    `指定的付款渠道无效`,
		Detail: `排查方法: -`,
	},
	ErrCode90345.Code(): {
		Msg:    `超过5w人民币不可指定微信支付渠道`,
		Detail: `排查方法: -`,
	},
	ErrCode90346.Code(): {
		Msg:    `同一单的退款次数超过限制`,
		Detail: `排查方法: -`,
	},
	ErrCode90347.Code(): {
		Msg:    `退款金额不可为0`,
		Detail: `排查方法: -`,
	},
	ErrCode90348.Code(): {
		Msg:    `管理端没配置支付密钥`,
		Detail: `排查方法: -`,
	},
	ErrCode90349.Code(): {
		Msg:    `记录数量太大`,
		Detail: `排查方法: -`,
	},
	ErrCode90350.Code(): {
		Msg:    `银行信息真实性校验失败`,
		Detail: `排查方法: -`,
	},
	ErrCode90351.Code(): {
		Msg:    `应用状态异常`,
		Detail: `排查方法: -`,
	},
	ErrCode90352.Code(): {
		Msg:    `延迟试用期天数超过限制`,
		Detail: `排查方法: -`,
	},
	ErrCode90353.Code(): {
		Msg:    `预支单列表不可为空`,
		Detail: `排查方法: -`,
	},
	ErrCode90354.Code(): {
		Msg:    `预支单列表数量超过限制`,
		Detail: `排查方法: -`,
	},
	ErrCode90355.Code(): {
		Msg:    `关联有退款预支单，不可删除`,
		Detail: `排查方法: -`,
	},
	ErrCode90356.Code(): {
		Msg:    `不能0金额下单`,
		Detail: `排查方法: -`,
	},
	ErrCode90357.Code(): {
		Msg:    `代下单必须指定支付渠道`,
		Detail: `排查方法: -`,
	},
	ErrCode90358.Code(): {
		Msg:    `预支单或代下单，不支持部分退款`,
		Detail: `排查方法: -`,
	},
	ErrCode90359.Code(): {
		Msg:    `预支单与下单者企业不匹配`,
		Detail: `排查方法: -`,
	},
	ErrCode90381.Code(): {
		Msg:    `参数 refunded_credit_orderid 不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode90456.Code(): {
		Msg:    `必须指定组织者`,
		Detail: `排查方法: -`,
	},
	ErrCode90457.Code(): {
		Msg:    `日历ID异常`,
		Detail: `排查方法: -`,
	},
	ErrCode90458.Code(): {
		Msg:    `日历ID列表不能为空`,
		Detail: `排查方法: -`,
	},
	ErrCode90459.Code(): {
		Msg:    `日历已删除`,
		Detail: `排查方法: -`,
	},
	ErrCode90460.Code(): {
		Msg:    `日程已删除`,
		Detail: `排查方法: -`,
	},
	ErrCode90461.Code(): {
		Msg:    `日程ID异常`,
		Detail: `排查方法: -`,
	},
	ErrCode90462.Code(): {
		Msg:    `日程ID列表不能为空`,
		Detail: `排查方法: -`,
	},
	ErrCode90463.Code(): {
		Msg:    `不能变更组织者`,
		Detail: `排查方法: -`,
	},
	ErrCode90464.Code(): {
		Msg:    `参与者数量超过限制`,
		Detail: `排查方法: -`,
	},
	ErrCode90465.Code(): {
		Msg:    `不支持的重复类型`,
		Detail: `排查方法: -`,
	},
	ErrCode90466.Code(): {
		Msg:    `不能操作别的应用创建的日历/日程`,
		Detail: `排查方法: -`,
	},
	ErrCode90467.Code(): {
		Msg:    `星期参数异常`,
		Detail: `排查方法: -`,
	},
	ErrCode90468.Code(): {
		Msg:    `不能变更组织者`,
		Detail: `排查方法: -`,
	},
	ErrCode90469.Code(): {
		Msg:    `每页大小超过限制`,
		Detail: `排查方法: -`,
	},
	ErrCode90470.Code(): {
		Msg:    `页数异常`,
		Detail: `排查方法: -`,
	},
	ErrCode90471.Code(): {
		Msg:    `提醒时间异常`,
		Detail: `排查方法: -`,
	},
	ErrCode90472.Code(): {
		Msg:    `没有日历/日程操作权限`,
		Detail: `排查方法: -`,
	},
	ErrCode90473.Code(): {
		Msg:    `颜色参数异常`,
		Detail: `排查方法: -`,
	},
	ErrCode90474.Code(): {
		Msg:    `组织者不能与参与者重叠`,
		Detail: `排查方法: -`,
	},
	ErrCode90475.Code(): {
		Msg:    `不是组织者的日历`,
		Detail: `排查方法: -`,
	},
	ErrCode90479.Code(): {
		Msg:    `不允许操作用户创建的日程`,
		Detail: `排查方法: -`,
	},
	ErrCode90500.Code(): {
		Msg:    `群主并未离职`,
		Detail: `排查方法: -`,
	},
	ErrCode90501.Code(): {
		Msg:    `该群不是客户群`,
		Detail: `排查方法: -`,
	},
	ErrCode90502.Code(): {
		Msg:    `群主已经离职`,
		Detail: `排查方法: -`,
	},
	ErrCode90503.Code(): {
		Msg:    `满人 & 99个微信成员，没办法踢，要客户端确认`,
		Detail: `排查方法: -`,
	},
	ErrCode90504.Code(): {
		Msg:    `群主没变`,
		Detail: `排查方法: -`,
	},
	ErrCode90507.Code(): {
		Msg:    `离职群正在继承处理中`,
		Detail: `排查方法: -`,
	},
	ErrCode90508.Code(): {
		Msg:    `离职群已经继承`,
		Detail: `排查方法: -`,
	},
	ErrCode91040.Code(): {
		Msg:    `获取ticket的类型无效`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：91040);获取ticket的类型无效。jsapi ticket可以通过以下几种获取：;1）[获取jsapi_ticket](https://opDetail.work.weixin.qq.com/api/doc#10029/获取jsapi_ticket)。这里参数只需要传access_tokDetail，不需要带其余的参数，比如type=jsapi;2）[获取电子发票ticket](https://opDetail.work.weixin.qq.com/api/doc#10029/获取电子发票ticket)。需要同时指定access_tokDetail及type，同时type=wx_card是固定的。`,
	},
	ErrCode92000.Code(): {
		Msg:    `成员不在应用可见范围之内`,
		Detail: `排查方法: -`,
	},
	ErrCode92001.Code(): {
		Msg:    `应用没有敏感信息权限`,
		Detail: `排查方法: -`,
	},
	ErrCode92002.Code(): {
		Msg:    `不允许跨企业调用`,
		Detail: `排查方法: -`,
	},
	ErrCode92006.Code(): {
		Msg:    `该直播已经开始或取消`,
		Detail: `排查方法: -`,
	},
	ErrCode92007.Code(): {
		Msg:    `该直播回放不能被删除`,
		Detail: `排查方法: -`,
	},
	ErrCode92008.Code(): {
		Msg:    `当前应用没权限操作这个直播`,
		Detail: `排查方法: -`,
	},
	ErrCode93000.Code(): {
		Msg:    `机器人webhookurl不合法或者机器人已经被移除出群`,
		Detail: `排查方法: -`,
	},
	ErrCode93004.Code(): {
		Msg:    `机器人被停用`,
		Detail: `排查方法: -`,
	},
	ErrCode93008.Code(): {
		Msg:    `不在群里`,
		Detail: `排查方法: -`,
	},
	ErrCode94000.Code(): {
		Msg:    `应用未开启工作台自定义模式`,
		Detail: `排查方法: 请在管理端后台应用详情里面开启自定义工作台模式`,
	},
	ErrCode94001.Code(): {
		Msg:    `不合法的type类型`,
		Detail: `排查方法: -`,
	},
	ErrCode94002.Code(): {
		Msg:    `缺少keydata字段`,
		Detail: `排查方法: -`,
	},
	ErrCode94003.Code(): {
		Msg:    `keydata的items列表长度超出限制`,
		Detail: `排查方法: -`,
	},
	ErrCode94005.Code(): {
		Msg:    `缺少list字段`,
		Detail: `排查方法: -`,
	},
	ErrCode94006.Code(): {
		Msg:    `list的items列表长度超出限制`,
		Detail: `排查方法: -`,
	},
	ErrCode94007.Code(): {
		Msg:    `缺少webview字段`,
		Detail: `排查方法: -`,
	},
	ErrCode94008.Code(): {
		Msg:    `应用未设置自定义工作台模版类型`,
		Detail: `排查方法: -`,
	},
	ErrCode301002.Code(): {
		Msg:    `无权限操作指定的应用`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：301002);无权限操作指定的应用。access_tokDetail来源需要有指定应用的权限。;比如说，[发消息接口](https://opDetail.work.weixin.qq.com/api/doc#10167) 指定了参数 “agDetailtid”: 14，但使用的 access_tokDetail 是通过应用agDetailtid: 100032 生成的调用凭证，这种就会报该错误码。;access_tokDetail的权限查询，可在 [错误码查询工具](https://opDetail.work.weixin.qq.com/devtool/query) 确认。`,
	},
	ErrCode301005.Code(): {
		Msg:    `不允许删除创建者`,
		Detail: `排查方法: 创建者不允许从通讯录中删除。如果需要删除该成员，需要先在WEB管理端转移创建者身份。`,
	},
	ErrCode301012.Code(): {
		Msg:    `参数 position 不合法`,
		Detail: `排查方法: 长度不允许超过128个字符`,
	},
	ErrCode301013.Code(): {
		Msg:    `参数 telephone 不合法`,
		Detail: `排查方法: telephone必须由1-32位的纯数字或’-‘号组成。`,
	},
	ErrCode301014.Code(): {
		Msg:    `参数 Detailglish_name 不合法`,
		Detail: `排查方法: 参数如果有传递，不允许为空字符串，同时不能超过64字节，只能是由字母、数字、点(.)、减号(-)、空格或下划线(_)组成`,
	},
	ErrCode301015.Code(): {
		Msg:    `参数 mediaid 不合法`,
		Detail: `排查方法: 请检查 mediaid 来源，应该通过[上传临时素材](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112)的图片类型获得mediaid`,
	},
	ErrCode301016.Code(): {
		Msg:    `上传语音文件不符合系统要求`,
		Detail: `排查方法: 语音文件的系统限制，参考[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112/上传的媒体文件限制)`,
	},
	ErrCode301017.Code(): {
		Msg:    `上传语音文件仅支持AMR格式`,
		Detail: `排查方法: 语音文件的系统限制，参考[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112/上传的媒体文件限制)`,
	},
	ErrCode301021.Code(): {
		Msg:    `参数 userid 无效`,
		Detail: `排查方法: 至少有一个userid不存在于通讯录中`,
	},
	ErrCode301022.Code(): {
		Msg:    `获取打卡数据失败`,
		Detail: `排查方法: 系统失败，可重试处理`,
	},
	ErrCode301023.Code(): {
		Msg:    `useridlist非法或超过限额`,
		Detail: `排查方法: 列表数量不能为0且不超过100`,
	},
	ErrCode301024.Code(): {
		Msg:    `获取打卡记录时间间隔超限`,
		Detail: `排查方法: 保证开始时间大于0 且结束时间大于 0 且结束时间大于开始时间，且间隔少于一个月`,
	},
	ErrCode301025.Code(): {
		Msg:    `审批开放接口参数错误`,
		Detail: `排查方法: 请参考参数说明正确填写`,
	},
	ErrCode301036.Code(): {
		Msg:    `不允许更新该用户的userid`,
		Detail: `排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：301036);不允许更新该用户的userid。确认：;只有当userid由系统自动生成时，才被允许修改一次;比如，邀请关注时用户提交登记信息，审批通过后系统会自动分配userid，此时可修改userid`,
	},
	ErrCode301039.Code(): {
		Msg:    `请求参数错误，请检查输入参数`,
		Detail: `排查方法: -`,
	},
	ErrCode301042.Code(): {
		Msg:    `ip白名单限制，请求ip不在设置白名单范围`,
		Detail: `排查方法: -`,
	},
	ErrCode301048.Code(): {
		Msg:    `sdkfileid对应的文件不存在或已过期`,
		Detail: `排查方法: -`,
	},
	ErrCode301052.Code(): {
		Msg:    `会话存档服务已过期`,
		Detail: `排查方法: -`,
	},
	ErrCode301053.Code(): {
		Msg:    `会话存档服务未开启`,
		Detail: `排查方法: -`,
	},
	ErrCode301058.Code(): {
		Msg:    `拉取会话数据请求超过大小限制，可减少limit参数`,
		Detail: `排查方法: -`,
	},
	ErrCode301059.Code(): {
		Msg:    `非内部群，不提供数据`,
		Detail: `排查方法: -`,
	},
	ErrCode301060.Code(): {
		Msg:    `拉取同意情况请求量过大，请减少到100个参数以下`,
		Detail: `排查方法: -`,
	},
	ErrCode301061.Code(): {
		Msg:    `userid或者exteropDetailid用户不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode302003.Code(): {
		Msg:    `批量导入任务的文件中userid有重复`,
		Detail: `排查方法: -`,
	},
	ErrCode302004.Code(): {
		Msg:    `组织架构不合法（1不是一棵树，2 多个一样的partyid，3 partyid空，4 partyid name 空，5 同一个父节点下有两个子节点 部门名字一样 可能是以上情况，请一一排查）`,
		Detail: `排查方法: -`,
	},
	ErrCode302005.Code(): {
		Msg:    `批量导入系统失败，请重新尝试导入`,
		Detail: `排查方法: -`,
	},
	ErrCode302006.Code(): {
		Msg:    `批量导入任务的文件中partyid有重复`,
		Detail: `排查方法: -`,
	},
	ErrCode302007.Code(): {
		Msg:    `批量导入任务的文件中，同一个部门下有两个子部门名字一样`,
		Detail: `排查方法: -`,
	},
	ErrCode2000002.Code(): {
		Msg:    `CorpId参数无效`,
		Detail: `排查方法: 指定的CorpId不存在`,
	},
	ErrCode600001.Code(): {
		Msg:    `不合法的sn`,
		Detail: `排查方法: sn可能尚未进行登记`,
	},
	ErrCode600002.Code(): {
		Msg:    `设备已注册`,
		Detail: `排查方法: 可能设备已经建立过长连接`,
	},
	ErrCode600003.Code(): {
		Msg:    `不合法的硬件activecode`,
		Detail: `排查方法: -`,
	},
	ErrCode600004.Code(): {
		Msg:    `该硬件尚未授权任何企业`,
		Detail: `排查方法: -`,
	},
	ErrCode600005.Code(): {
		Msg:    `硬件Secret无效`,
		Detail: `排查方法: -`,
	},
	ErrCode600007.Code(): {
		Msg:    `缺少硬件sn`,
		Detail: `排查方法: -`,
	},
	ErrCode600008.Code(): {
		Msg:    `缺少nonce参数`,
		Detail: `排查方法: -`,
	},
	ErrCode600009.Code(): {
		Msg:    `缺少timestamp参数`,
		Detail: `排查方法: -`,
	},
	ErrCode600010.Code(): {
		Msg:    `缺少signature参数`,
		Detail: `排查方法: -`,
	},
	ErrCode600011.Code(): {
		Msg:    `签名校验失败`,
		Detail: `排查方法: -`,
	},
	ErrCode600012.Code(): {
		Msg:    `长连接已经注册过设备`,
		Detail: `排查方法: -`,
	},
	ErrCode600013.Code(): {
		Msg:    `缺少activecode参数`,
		Detail: `排查方法: -`,
	},
	ErrCode600014.Code(): {
		Msg:    `设备未网络注册`,
		Detail: `排查方法: -`,
	},
	ErrCode600015.Code(): {
		Msg:    `缺少secret参数`,
		Detail: `排查方法: -`,
	},
	ErrCode600016.Code(): {
		Msg:    `设备未激活`,
		Detail: `排查方法: -`,
	},
	ErrCode600018.Code(): {
		Msg:    `无效的起始结束时间`,
		Detail: `排查方法: -`,
	},
	ErrCode600020.Code(): {
		Msg:    `设备未登录`,
		Detail: `排查方法: -`,
	},
	ErrCode600021.Code(): {
		Msg:    `设备sn已存在`,
		Detail: `排查方法: -`,
	},
	ErrCode600023.Code(): {
		Msg:    `时间戳已失效`,
		Detail: `排查方法: -`,
	},
	ErrCode600024.Code(): {
		Msg:    `固件大小超过5M`,
		Detail: `排查方法: -`,
	},
	ErrCode600025.Code(): {
		Msg:    `固件名为空或者超过20字节`,
		Detail: `排查方法: -`,
	},
	ErrCode600026.Code(): {
		Msg:    `固件信息不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode600027.Code(): {
		Msg:    `非法的固件参数`,
		Detail: `排查方法: -`,
	},
	ErrCode600028.Code(): {
		Msg:    `固件版本已存在`,
		Detail: `排查方法: -`,
	},
	ErrCode600029.Code(): {
		Msg:    `非法的固件版本`,
		Detail: `排查方法: -`,
	},
	ErrCode600030.Code(): {
		Msg:    `缺少固件版本参数`,
		Detail: `排查方法: -`,
	},
	ErrCode600031.Code(): {
		Msg:    `硬件固件不允许升级`,
		Detail: `排查方法: -`,
	},
	ErrCode600032.Code(): {
		Msg:    `无法解析硬件二维码`,
		Detail: `排查方法: -`,
	},
	ErrCode600033.Code(): {
		Msg:    `设备型号id冲突`,
		Detail: `排查方法: -`,
	},
	ErrCode600034.Code(): {
		Msg:    `指纹数据大小超过限制`,
		Detail: `排查方法: -`,
	},
	ErrCode600035.Code(): {
		Msg:    `人脸数据大小超过限制`,
		Detail: `排查方法: -`,
	},
	ErrCode600036.Code(): {
		Msg:    `设备sn冲突`,
		Detail: `排查方法: -`,
	},
	ErrCode600037.Code(): {
		Msg:    `缺失设备型号id`,
		Detail: `排查方法: -`,
	},
	ErrCode600038.Code(): {
		Msg:    `设备型号不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode600039.Code(): {
		Msg:    `不支持的设备类型`,
		Detail: `排查方法: -`,
	},
	ErrCode600040.Code(): {
		Msg:    `打印任务id不存在`,
		Detail: `排查方法: -`,
	},
	ErrCode600041.Code(): {
		Msg:    `无效的offset或limit参数值`,
		Detail: `排查方法: -`,
	},
	ErrCode600042.Code(): {
		Msg:    `无效的设备型号id`,
		Detail: `排查方法: -`,
	},
	ErrCode600043.Code(): {
		Msg:    `门禁规则未设置`,
		Detail: `排查方法: -`,
	},
	ErrCode600044.Code(): {
		Msg:    `门禁规则不合法`,
		Detail: `排查方法: -`,
	},
	ErrCode600045.Code(): {
		Msg:    `设备已订阅企业信息`,
		Detail: `排查方法: -`,
	},
	ErrCode600046.Code(): {
		Msg:    `操作id和用户userid不匹配`,
		Detail: `排查方法: -`,
	},
	ErrCode600047.Code(): {
		Msg:    `secretno的status非法`,
		Detail: `排查方法: 请确认是否是使用统一初始secretno的设备，如果是无有正确执行换secretno的流程`,
	},
	ErrCode600048.Code(): {
		Msg:    `无效的指纹算法`,
		Detail: `排查方法: -`,
	},
	ErrCode600049.Code(): {
		Msg:    `无效的人脸识别算法`,
		Detail: `排查方法: -`,
	},
	ErrCode600050.Code(): {
		Msg:    `无效的算法长度`,
		Detail: `排查方法: -`,
	},
	ErrCode600051.Code(): {
		Msg:    `设备过期`,
		Detail: `排查方法: -`,
	},
	ErrCode600052.Code(): {
		Msg:    `无效的文件分块`,
		Detail: `排查方法: -`,
	},
	ErrCode600053.Code(): {
		Msg:    `该链接已经激活`,
		Detail: `排查方法: -`,
	},
	ErrCode600054.Code(): {
		Msg:    `该链接已经订阅`,
		Detail: `排查方法: -`,
	},
	ErrCode600055.Code(): {
		Msg:    `无效的用户类型`,
		Detail: `排查方法: -`,
	},
	ErrCode600056.Code(): {
		Msg:    `无效的健康状态`,
		Detail: `排查方法: -`,
	},
	ErrCode600057.Code(): {
		Msg:    `缺少体温参数`,
		Detail: `排查方法: -`,
	},
	ErrCode610001.Code(): {
		Msg:    `永久二维码超过每个员工5000的限制`,
		Detail: `排查方法: -`,
	},
	ErrCode610003.Code(): {
		Msg:    `scDetaile参数不合法`,
		Detail: `排查方法: 有效的scDetaile长度为1~64字符，由英文字母、数字、中划线、下划线以及点号构成`,
	},
	ErrCode610004.Code(): {
		Msg:    `userid不在客户联系配置的使用范围内`,
		Detail: `排查方法: 请在管理端后台 客户联系->配置->配置使用范围配置该用户`,
	},
	ErrCode610014.Code(): {
		Msg:    `无效的unionid`,
		Detail: `排查方法: -`,
	},
	ErrCode610015.Code(): {
		Msg:    `小程序对应的开放平台账号未认证`,
		Detail: `排查方法: -`,
	},
	ErrCode610016.Code(): {
		Msg:    `企业未认证`,
		Detail: `排查方法: -`,
	},
	ErrCode610017.Code(): {
		Msg:    `小程序和企业主体不一致`,
		Detail: `排查方法: -`,
	},
	ErrCode640001.Code(): {
		Msg:    `微盘不存在当前空间`,
		Detail: `排查方法: 判断spaceid是否填错`,
	},
	ErrCode640002.Code(): {
		Msg:    `文件不存在`,
		Detail: `排查方法: 判断fileid是否填错`,
	},
	ErrCode640003.Code(): {
		Msg:    `文件已删除`,
		Detail: `排查方法: 判断fileid对应的文件是否已经被删除`,
	},
	ErrCode640004.Code(): {
		Msg:    `无权限访问`,
		Detail: `排查方法: 判断当前用户是否有访问`,
	},
	ErrCode640005.Code(): {
		Msg:    `成员不在空间内`,
		Detail: `排查方法: 判断当前成员是否在空间内`,
	},
	ErrCode640006.Code(): {
		Msg:    `超出当前成员拥有的容量`,
		Detail: `排查方法: 判断当前用户的个人容量是否超出了限制`,
	},
	ErrCode640007.Code(): {
		Msg:    `超出微盘的容量`,
		Detail: `排查方法: 在管理端查看微盘的容量是否要满了`,
	},
	ErrCode640008.Code(): {
		Msg:    `没有空间权限`,
		Detail: `排查方法: 判断当前userid是否有空间权限`,
	},
	ErrCode640009.Code(): {
		Msg:    `非法文件名`,
		Detail: `排查方法: 判断file_name字段是否为空`,
	},
	ErrCode640010.Code(): {
		Msg:    `超出空间的最大成员数`,
		Detail: `排查方法: 判断当前空间的成员数是否超过了管理端设置的空间最大成员数`,
	},
	ErrCode640011.Code(): {
		Msg:    `json格式不匹配`,
		Detail: `排查方法: 判断是否的json格式是否有误`,
	},
	ErrCode640012.Code(): {
		Msg:    `非法的userid`,
		Detail: `排查方法: 判断userid字段是否设置错误`,
	},
	ErrCode640013.Code(): {
		Msg:    `非法的departmDetailtid`,
		Detail: `排查方法: 判断departmDetailtid字段是否设置错误`,
	},
	ErrCode640014.Code(): {
		Msg:    `空间没有有效的管理员`,
		Detail: `排查方法: 判断当前空间是否没有有效的管理员`,
	},
	ErrCode640015.Code(): {
		Msg:    `不支持设置预览权限`,
		Detail: `排查方法: 文件预览权限只有VIP用户才能设置`,
	},
	ErrCode640016.Code(): {
		Msg:    `不支持设置文件水印`,
		Detail: `排查方法: 文件水印只有VIP用户才能设置`,
	},
	ErrCode640017.Code(): {
		Msg:    `微盘管理端未开通API 权限`,
		Detail: `排查方法: 在微盘管理端进行打开`,
	},
	ErrCode640018.Code(): {
		Msg:    `微盘管理端未设置编辑权限`,
		Detail: `排查方法: 在微盘管理端进行打开编辑权限`,
	},
	ErrCode640019.Code(): {
		Msg:    `API 调用次数超出限制`,
		Detail: `排查方法: 免费版：1000次/企业/月; 付费版：100,000次/企业/月`,
	},
	ErrCode640020.Code(): {
		Msg:    `非法的权限类型`,
		Detail: `排查方法: 普通文件：可下载、仅预览; 微文档：可编辑、仅浏览`,
	},
	ErrCode640021.Code(): {
		Msg:    `非法的fatherid`,
		Detail: `排查方法: fatherid应该为：文件所在的目录fileid, 在根目录时为fileid（判断当前字段是否为空）`,
	},
	ErrCode640022.Code(): {
		Msg:    `非法的文件内容的base64`,
		Detail: `排查方法: 文件内容base64，判断本字段是否为空`,
	},
	ErrCode640023.Code(): {
		Msg:    `非法的权限范围`,
		Detail: `排查方法: auth_scope应该为三个中的其中一个：1:指定人 2:企业内 3:企业外`,
	},
	ErrCode640024.Code(): {
		Msg:    `非法的fileid`,
		Detail: `排查方法: 判断fileid字段是否为空`,
	},
	ErrCode640025.Code(): {
		Msg:    `非法的space_name`,
		Detail: `排查方法: 判断space_name字段是否空`,
	},
	ErrCode640026.Code(): {
		Msg:    `非法的spaceid`,
		Detail: `排查方法: 判断spaceid字段是否空`,
	},
	ErrCode640027.Code(): {
		Msg:    `参数错误`,
		Detail: `排查方法: 判断输入的参数是否有误`,
	},
	ErrCode640028.Code(): {
		Msg:    `空间设置了关闭成员邀请链接`,
		Detail: `排查方法: 查看空间的安全设置的成员邀请链接按钮是否处于关闭状态`,
	},
	ErrCode640029.Code(): {
		Msg:    `只支持下载普通文件，不支持下载文件夹等其他非文件实体类型`,
		Detail: `排查方法: 检查fileid对应的文件是否为普通文件`,
	},
	ErrCode844001.Code(): {
		Msg:    `非法的output_file_format`,
		Detail: `排查方法: 判断输出文件格式是否正确`,
	},
}
