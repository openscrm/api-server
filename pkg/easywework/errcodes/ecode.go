package errcodes

// ErrCode 错误码类型
// 全局错误码文档: https://work.weixin.qq.com/api/doc/90000/90139/90313

type ErrCode = int64

// ErrCodeServiceUnavailable 系统繁忙
// 排查方法: 服务器暂不可用，建议稍候重试。建议重试次数不超过3次。
const ErrCodeServiceUnavailable ErrCode = -1

// ErrCodeSuccess 请求成功
// 排查方法: 接口调用成功
const ErrCodeSuccess ErrCode = 0

// ErrCode6000 数据版本冲突
// 排查方法: 可能有多个调用端同时修改数据，稍后重试
const ErrCode6000 ErrCode = 6000

// ErrCode40001 不合法的secret参数
// 排查方法: secret在应用详情/通讯录管理助手可查看
const ErrCode40001 ErrCode = 40001

// ErrCode40003 无效的UserID
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40003)
// 不合法的UserID。确认：
// 1）有效的UserID需要满足：长度1~64字符，由英文字母、数字、中划线、下划线以及点号构成。
// 2）除了创建用户，其余使用UserID的接口，还要保证UserID必须在通讯录中存在。
const ErrCode40003 ErrCode = 40003

// ErrCode40004 不合法的媒体文件类型
// 排查方法: 不满足系统文件要求。参考：[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112)
const ErrCode40004 ErrCode = 40004

// ErrCode40005 不合法的type参数
// 排查方法: 合法的type取值，参考：[上传临时素材](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112)
const ErrCode40005 ErrCode = 40005

// ErrCode40006 不合法的文件大小
// 排查方法: 系统文件要求，参考：[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112)
const ErrCode40006 ErrCode = 40006

// ErrCode40007 不合法的media_id参数
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40007)
// 不合法的媒体文件。确认：
// 1）媒体文件ID的获取方式，是否存在。注：上传临时素材生成的medida_id，有效期是3天。
// 2）媒体文件类型应符合接口要求（比如发送图片消息，此时不能用音频文件的media_id）。
const ErrCode40007 ErrCode = 40007

// ErrCode40008 不合法的msgtype参数
// 排查方法: 合法的msgtype取值，参考：[消息类型](https://work.weixin.qq.com/api/doc/90000/90139/90313#10167)
const ErrCode40008 ErrCode = 40008

// ErrCode40009 上传图片大小不是有效值
// 排查方法: 图片大小的系统限制，参考[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112/上传的媒体文件限制)
const ErrCode40009 ErrCode = 40009

// ErrCode40011 上传视频大小不是有效值
// 排查方法: 视频大小的系统限制，参考[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112/上传的媒体文件限制)
const ErrCode40011 ErrCode = 40011

// ErrCode40013 不合法的CorpID
// 排查方法: 需确认CorpID是否填写正确，在 web管理端-设置 可查看
const ErrCode40013 ErrCode = 40013

// ErrCode40014 不合法的access_token
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40014)
// access_token参数错误。确认：
// 1）access_token的获取方式
// 2）access_token是否已过期
// 可以重新获取一次access_token解决
const ErrCode40014 ErrCode = 40014

// ErrCode40016 不合法的按钮个数
// 排查方法: 菜单按钮1-3个
const ErrCode40016 ErrCode = 40016

// ErrCode40017 不合法的按钮类型
// 排查方法: 支持的类型，参考：[按钮类型](https://work.weixin.qq.com/api/doc/90000/90139/90313#10786)
const ErrCode40017 ErrCode = 40017

// ErrCode40018 不合法的按钮名字长度
// 排查方法: 长度应不超过16个字节
const ErrCode40018 ErrCode = 40018

// ErrCode40019 不合法的按钮KEY长度
// 排查方法: 长度应不超过128字节
const ErrCode40019 ErrCode = 40019

// ErrCode40020 不合法的按钮URL长度
// 排查方法: 长度应不超过1024字节
const ErrCode40020 ErrCode = 40020

// ErrCode40022 不合法的子菜单级数
// 排查方法: 只能包含一级菜单和二级菜单
const ErrCode40022 ErrCode = 40022

// ErrCode40023 不合法的子菜单按钮个数
// 排查方法: 子菜单按钮1-5个
const ErrCode40023 ErrCode = 40023

// ErrCode40024 不合法的子菜单按钮类型
// 排查方法: 支持的类型，参考：[按钮类型](https://work.weixin.qq.com/api/doc/90000/90139/90313#10786)
const ErrCode40024 ErrCode = 40024

// ErrCode40025 不合法的子菜单按钮名字长度
// 排查方法: 支持的类型，参考：[按钮类型](https://work.weixin.qq.com/api/doc/90000/90139/90313#10786)
const ErrCode40025 ErrCode = 40025

// ErrCode40026 不合法的子菜单按钮KEY长度
// 排查方法: -
const ErrCode40026 ErrCode = 40026

// ErrCode40027 不合法的子菜单按钮URL长度
// 排查方法: 长度应不超过1024字节
const ErrCode40027 ErrCode = 40027

// ErrCode40029 不合法的oauth_code
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40029)
// oauth_code参数错误。确认：
// 1）code只能消费一次，不能重复消费。比如说，是否存在多个服务器同时消费同一code情况。
// 2）code需要在有效期间消费（5分钟），过期会自动失效。
const ErrCode40029 ErrCode = 40029

// ErrCode40031 不合法的UserID列表
// 排查方法: 指定的UserID列表，部分UserID不在通讯录中
const ErrCode40031 ErrCode = 40031

// ErrCode40032 不合法的UserID列表长度
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40032)
// 不合法的UserID列表长度。确认：
// 1）[发消息接口](https://open.work.weixin.qq.com/api/doc#10167)，最多指定1000人。
// 2）[批量删除成员接口](https://open.work.weixin.qq.com/api/doc#10060)，最多指定200人。
const ErrCode40032 ErrCode = 40032

// ErrCode40033 不合法的请求字符
// 排查方法: 不能包含\uxxxx格式的字符
const ErrCode40033 ErrCode = 40033

// ErrCode40035 不合法的参数
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40035)
// 不合法的参数。确认：
// 1）userlist和partylist不能同时为空
// 2）userlist包含的成员个数不能大于1000
// 3）partylist包含的部门个数不能大于100
// 4）指定的userlist和partylist为数组格式，不是字符串格式。比如说， “userlist”:[ “user1”,”user2”]，而不是 “userlist”: “user1|user2”
const ErrCode40035 ErrCode = 40035

// ErrCode40036 不合法的模板id长度
// 排查方法: -
const ErrCode40036 ErrCode = 40036

// ErrCode40037 无效的模板id
// 排查方法: -
const ErrCode40037 ErrCode = 40037

// ErrCode40039 不合法的url长度
// 排查方法: url长度限制1024个字节
const ErrCode40039 ErrCode = 40039

// ErrCode40050 chatid不存在
// 排查方法: 会话需要先创建后，才可修改会话详情或者发起聊天
const ErrCode40050 ErrCode = 40050

// ErrCode40054 不合法的子菜单url域名
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40054 40055)
// 菜单设置URL不合法。确认：
// 1）链接需要带上协议头。以 http:// 或者 https:// 开头。比如：https://work.weixin.qq.com
// 2）微信支付的链接，必须以 weixin://wxpay/bizpayurl 开头
const ErrCode40054 ErrCode = 40054

// ErrCode40055 不合法的菜单url域名
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40054 40055)
// 菜单设置URL不合法。确认：
// 1）链接需要带上协议头。以 http:// 或者 https:// 开头。比如：https://work.weixin.qq.com
// 2）微信支付的链接，必须以 weixin://wxpay/bizpayurl 开头
const ErrCode40055 ErrCode = 40055

// ErrCode40056 不合法的agentid
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40056)
// agentid不合法。确认：
// 1）agentid为整型数字
// 2）在web管理端存在该应用
const ErrCode40056 ErrCode = 40056

// ErrCode40057 不合法的callbackurl或者callbackurl验证失败
// 排查方法: 可自助到[开发调试工具](https://work.weixin.qq.com/api/devtools/devtool.php)重现
const ErrCode40057 ErrCode = 40057

// ErrCode40058 不合法的参数
// 排查方法: 传递参数不符合系统要求，需要参照具体API接口说明
const ErrCode40058 ErrCode = 40058

// ErrCode40059 不合法的上报地理位置标志位
// 排查方法: 开关标志位只能填 0 或者 1
const ErrCode40059 ErrCode = 40059

// ErrCode40063 参数为空
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40063)
// 必填的参数缺少，需要参照具体API接口说明。同时确认：
// 1）Http请求方法，是否正确。比如说接口要求以Post方法，就不能使用Get方式
// 2）Http请求参数，是否正确。比如说，接口内容要求json结构体，就不能以url参数传递或者form-data方式。
const ErrCode40063 ErrCode = 40063

// ErrCode40066 不合法的部门列表
// 排查方法: 部门列表为空，或者至少存在一个部门ID不存在于通讯录中
const ErrCode40066 ErrCode = 40066

// ErrCode40068 不合法的标签ID
// 排查方法: 标签ID未指定，或者指定的标签ID不存在
const ErrCode40068 ErrCode = 40068

// ErrCode40070 指定的标签范围结点全部无效
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40070)
// 指定的标签范围结点全部无效。确认：
// 1）指定的参数格式是否正确。比如，”userlist”:[ “user1”]，而不是指定为 “userlist” : “user1”。
// 2）指定的成员或者部门，是否存在于通讯录中。
const ErrCode40070 ErrCode = 40070

// ErrCode40071 不合法的标签名字
// 排查方法: 标签名字已经存在
const ErrCode40071 ErrCode = 40071

// ErrCode40072 不合法的标签名字长度
// 排查方法: 不允许为空，最大长度限制为32个字（汉字或英文字母）
const ErrCode40072 ErrCode = 40072

// ErrCode40073 不合法的openid
// 排查方法: openid不存在，需确认获取来源
const ErrCode40073 ErrCode = 40073

// ErrCode40074 news消息不支持保密消息类型
// 排查方法: 图文消息支持保密类型需改用mpnews
const ErrCode40074 ErrCode = 40074

// ErrCode40077 不合法的pre_auth_code参数
// 排查方法: 预授权码不存在，参考：[获取预授权码](https://work.weixin.qq.com/api/doc/90000/90139/90313#10975/获取预授权码)
const ErrCode40077 ErrCode = 40077

// ErrCode40078 不合法的auth_code参数
// 排查方法: 需确认获取来源，并且只能消费一次
const ErrCode40078 ErrCode = 40078

// ErrCode40080 不合法的suite_secret
// 排查方法: 套件secret可在第三方管理端套件详情查看
const ErrCode40080 ErrCode = 40080

// ErrCode40082 不合法的suite_token
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40082)
// suite_token参数错误。确认：
// 1）suite_token的获取方式
// 2）suite_token是否已过期
// 可以重新获取一次suite_token解决
const ErrCode40082 ErrCode = 40082

// ErrCode40083 不合法的suite_id
// 排查方法: suite_id不存在
const ErrCode40083 ErrCode = 40083

// ErrCode40084 不合法的permanent_code参数
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40084)
// 不合法的永久授权码。确认：
// 1）是否填写有误
// 2）企业是否已取消授权该套件
// 3）永久授权码不能跨服务商使用
const ErrCode40084 ErrCode = 40084

// ErrCode40085 不合法的的suite_ticket参数
// 排查方法: suite_ticket不存在或者已失效
const ErrCode40085 ErrCode = 40085

// ErrCode40086 不合法的第三方应用appid
// 排查方法: 至少有一个不存在应用id
const ErrCode40086 ErrCode = 40086

// ErrCode40088 jobid不存在
// 排查方法: 请检查 jobid 来源
const ErrCode40088 ErrCode = 40088

// ErrCode40089 批量任务的结果已清理
// 排查方法: 系统仅保存最近5次批量任务的结果。可在通讯录查看实际导入情况
const ErrCode40089 ErrCode = 40089

// ErrCode40091 secret不合法
// 排查方法: 可能用了别的企业的secret
const ErrCode40091 ErrCode = 40091

// ErrCode40092 导入文件存在不合法的内容
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：40092)
// 导入文件存在不合法的内容。确认：
// 1）不允许上传空文件
// 2）文件内容缺少必填字段，比如：手机/邮箱，姓名，UserID或者部门。
const ErrCode40092 ErrCode = 40092

// ErrCode40093 jsapi签名错误
// 排查方法: 请检查用于签名的jsapi_ticket是否是正确的，是否过期。可以通过获取相应jsapi_ticket接口获取当前的jsapi_ticket跟用于签名的jsapi_ticket比对是否一致，若jsapi_ticket还在有效期内，当前获取到的jsapi_ticket是一致的。若jsapi_ticket没问题，请检查用于签名的url参数是不是正确的， url（当前网页的URL， 不包含#及其后面部分）。
const ErrCode40093 ErrCode = 40093

// ErrCode40094 不合法的URL
// 排查方法: 缺少主页URL参数，或者URL不合法（链接需要带上协议头，以 http:// 或者 https:// 开头）
const ErrCode40094 ErrCode = 40094

// ErrCode40096 不合法的外部联系人userid
// 排查方法: -
const ErrCode40096 ErrCode = 40096

// ErrCode40097 该成员尚未离职
// 排查方法: 离职成员外部联系人转移接口要求转出用户必须已经离职
const ErrCode40097 ErrCode = 40097

// ErrCode40098 成员尚未实名认证
// 排查方法: 确认传入的userid是已经过实名认证成员的
const ErrCode40098 ErrCode = 40098

// ErrCode40099 外部联系人的数量已达上限
// 排查方法: -
const ErrCode40099 ErrCode = 40099

// ErrCode40100 此用户的外部联系人已经在转移流程中
// 排查方法: -
const ErrCode40100 ErrCode = 40100

// ErrCode40102 域名或IP不可与应用市场上架应用重复
// 排查方法: -
const ErrCode40102 ErrCode = 40102

// ErrCode40123 上传临时图片素材，图片格式非法
// 排查方法: 请确认上传的内容是否为合法的图片内容
const ErrCode40123 ErrCode = 40123

// ErrCode40124 推广活动里的sn禁止绑定
// 排查方法: -
const ErrCode40124 ErrCode = 40124

// ErrCode40125 无效的openuserid参数
// 排查方法: -
const ErrCode40125 ErrCode = 40125

// ErrCode40126 企业标签个数达到上限，最多为3000个
// 排查方法: -
const ErrCode40126 ErrCode = 40126

// ErrCode40127 不支持的uri schema
// 排查方法: 检查uri链接的schema是否符合参数要求
const ErrCode40127 ErrCode = 40127

// ErrCode40128 客户转接过于频繁（90天内只允许转接一次，同一个客户最多只能转接两次）
// 排查方法: -
const ErrCode40128 ErrCode = 40128

// ErrCode40129 当前客户正在转接中
// 排查方法: -
const ErrCode40129 ErrCode = 40129

// ErrCode40130 原跟进人与接手人一样，不可继承
// 排查方法: -
const ErrCode40130 ErrCode = 40130

// ErrCode40131 handover_userid 并不是外部联系人的跟进人
// 排查方法: -
const ErrCode40131 ErrCode = 40131

// ErrCode41001 缺少access_token参数
// 排查方法: -
const ErrCode41001 ErrCode = 41001

// ErrCode41002 缺少corpid参数
// 排查方法: -
const ErrCode41002 ErrCode = 41002

// ErrCode41004 缺少secret参数
// 排查方法: -
const ErrCode41004 ErrCode = 41004

// ErrCode41006 缺少media_id参数
// 排查方法: media_id为调用接口必填参数，请确认是否有传递
const ErrCode41006 ErrCode = 41006

// ErrCode41008 缺少auth code参数
// 排查方法: -
const ErrCode41008 ErrCode = 41008

// ErrCode41009 缺少userid参数
// 排查方法: -
const ErrCode41009 ErrCode = 41009

// ErrCode41010 缺少url参数
// 排查方法: -
const ErrCode41010 ErrCode = 41010

// ErrCode41011 缺少agentid参数
// 排查方法: -
const ErrCode41011 ErrCode = 41011

// ErrCode41016 缺少title参数
// 排查方法: 发送图文消息，标题是必填参数。请确认参数是否有传递。
const ErrCode41016 ErrCode = 41016

// ErrCode41019 缺少 department 参数
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：41019)
// 缺少 department 参数。确认：
// 1）创建成员接口，成员所属部门是必填信息。
// 2）所属部门是数字数组格式，不是字符串。如：”department: [1, 2]
const ErrCode41019 ErrCode = 41019

// ErrCode41017 缺少tagid参数
// 排查方法: -
const ErrCode41017 ErrCode = 41017

// ErrCode41021 缺少suite_id参数
// 排查方法: -
const ErrCode41021 ErrCode = 41021

// ErrCode41022 缺少suite_access_token参数
// 排查方法: -
const ErrCode41022 ErrCode = 41022

// ErrCode41023 缺少suite_ticket参数
// 排查方法: -
const ErrCode41023 ErrCode = 41023

// ErrCode41024 缺少secret参数
// 排查方法: -
const ErrCode41024 ErrCode = 41024

// ErrCode41025 缺少permanent_code参数
// 排查方法: -
const ErrCode41025 ErrCode = 41025

// ErrCode41033 缺少 description 参数
// 排查方法: [发送文本卡片消息接口](https://work.weixin.qq.com/api/doc/90000/90139/90313#10167/文本卡片消息)，description 是必填字段
const ErrCode41033 ErrCode = 41033

// ErrCode41035 缺少外部联系人userid参数
// 排查方法: -
const ErrCode41035 ErrCode = 41035

// ErrCode41036 不合法的企业对外简称
// 排查方法: 企业对外简称必须是认证过的，如果要改回默认简称，传空字符串把对外简称清除就可以了
const ErrCode41036 ErrCode = 41036

// ErrCode41037 缺少「联系我」type参数
// 排查方法: -
const ErrCode41037 ErrCode = 41037

// ErrCode41038 缺少「联系我」scene参数
// 排查方法: -
const ErrCode41038 ErrCode = 41038

// ErrCode41039 无效的「联系我」type参数
// 排查方法: -
const ErrCode41039 ErrCode = 41039

// ErrCode41040 无效的「联系我」scene参数
// 排查方法: -
const ErrCode41040 ErrCode = 41040

// ErrCode41041 「联系我」使用人数超过限制
// 排查方法: 默认限制不超过100人(包括部门展开后的人数)
const ErrCode41041 ErrCode = 41041

// ErrCode41042 无效的「联系我」style参数
// 排查方法: -
const ErrCode41042 ErrCode = 41042

// ErrCode41043 缺少「联系我」config_id参数
// 排查方法: -
const ErrCode41043 ErrCode = 41043

// ErrCode41044 无效的「联系我」config_id参数
// 排查方法: -
const ErrCode41044 ErrCode = 41044

// ErrCode41045 API添加「联系我」达到数量上限
// 排查方法: -
const ErrCode41045 ErrCode = 41045

// ErrCode41046 缺少企业群发消息id
// 排查方法: -
const ErrCode41046 ErrCode = 41046

// ErrCode41047 无效的企业群发消息id
// 排查方法: -
const ErrCode41047 ErrCode = 41047

// ErrCode41048 无可发送的客户
// 排查方法: -
const ErrCode41048 ErrCode = 41048

// ErrCode41049 缺少欢迎语code参数
// 排查方法: -
const ErrCode41049 ErrCode = 41049

// ErrCode41050 无效的欢迎语code
// 排查方法: 欢迎语code(welcome_code)具有时效性，须在添加好友后20秒内使用
const ErrCode41050 ErrCode = 41050

// ErrCode41051 客户和服务人员已经开始聊天了
// 排查方法: 已经开始的聊天的客户不能发送欢迎语
const ErrCode41051 ErrCode = 41051

// ErrCode41052 无效的发送时间
// 排查方法: -
const ErrCode41052 ErrCode = 41052

// ErrCode41053 客户未同意聊天存档
// 排查方法: 须外部联系人同意服务须知后，成员才可发送欢迎语
const ErrCode41053 ErrCode = 41053

// ErrCode41054 该用户尚未激活
// 排查方法: -
const ErrCode41054 ErrCode = 41054

// ErrCode41055 群欢迎语模板数量达到上限
// 排查方法: -
const ErrCode41055 ErrCode = 41055

// ErrCode41056 外部联系人id类型不正确
// 排查方法: -
const ErrCode41056 ErrCode = 41056

// ErrCode41057 企业或服务商未绑定微信开发者账号
// 排查方法: -
const ErrCode41057 ErrCode = 41057

// ErrCode41059 缺少moment_id参数
// 排查方法: -
const ErrCode41059 ErrCode = 41059

// ErrCode41060 不合法的moment_id参数
// 排查方法: -
const ErrCode41060 ErrCode = 41060

// ErrCode41061 不合法朋友圈发送成员userid，当前朋友圈并非此用户发表
// 排查方法: -
const ErrCode41061 ErrCode = 41061

// ErrCode41062 企业创建的朋友圈尚未被成员userid发表
// 排查方法: -
const ErrCode41062 ErrCode = 41062

// ErrCode41063 群发消息正在被派发中，请稍后再试
// 排查方法: [创建企业群发](https://work.weixin.qq.com/api/doc/90000/90139/90313#15836)后，立刻调用[获取企业的全部群发记录](https://work.weixin.qq.com/api/doc/90000/90139/90313#25429)的相关接口，将可能出现该错误
const ErrCode41063 ErrCode = 41063

// ErrCode41064 附件大小超过限制
// 排查方法: -
const ErrCode41064 ErrCode = 41064

// ErrCode41065 无效的附件类型
// 排查方法: -
const ErrCode41065 ErrCode = 41065

// ErrCode41066 用户视频号名称错误
// 排查方法: -
const ErrCode41066 ErrCode = 41066

// ErrCode41102 缺少菜单名
// 排查方法: -
const ErrCode41102 ErrCode = 41102

// ErrCode42001 access_token已过期
// 排查方法: access_token有时效性，需要重新获取一次
const ErrCode42001 ErrCode = 42001

// ErrCode42007 pre_auth_code已过期
// 排查方法: pre_auth_code有时效性，需要重新获取一次
const ErrCode42007 ErrCode = 42007

// ErrCode42009 suite_access_token已过期
// 排查方法: suite_access_token有时效性，需要重新获取一次
const ErrCode42009 ErrCode = 42009

// ErrCode42012 jsapi_ticket不可用，一般是没有正确调用接口来创建jsapi_ticket
// 排查方法: 如果是agentConfig使用，请特别注意是否是使用”[获取应用身份的ticket](https://work.weixin.qq.com/api/doc/90000/90139/90313#10029/获取应用的jsapi_ticket)“来获取jsapi_ticket
const ErrCode42012 ErrCode = 42012

// ErrCode42013 小程序未登陆或登录态已经过期
// 排查方法: 需要重新走登陆流程
const ErrCode42013 ErrCode = 42013

// ErrCode42014 任务卡片消息的task_id不合法
// 排查方法: -
const ErrCode42014 ErrCode = 42014

// ErrCode42015 更新的消息的应用与发送消息的应用不匹配
// 排查方法: -
const ErrCode42015 ErrCode = 42015

// ErrCode42016 更新的task_id不存在
// 排查方法: -
const ErrCode42016 ErrCode = 42016

// ErrCode42017 按钮key值不存在
// 排查方法: -
const ErrCode42017 ErrCode = 42017

// ErrCode42018 按钮key值不合法
// 排查方法: -
const ErrCode42018 ErrCode = 42018

// ErrCode42019 缺少按钮key值不合法
// 排查方法: -
const ErrCode42019 ErrCode = 42019

// ErrCode42020 缺少按钮名称
// 排查方法: -
const ErrCode42020 ErrCode = 42020

// ErrCode42021 device_access_token 过期
// 排查方法: -
const ErrCode42021 ErrCode = 42021

// ErrCode42022 code已经被使用过。只能使用一次
// 排查方法: -
const ErrCode42022 ErrCode = 42022

// ErrCode43004 指定的userid未绑定微信或未关注微工作台（原企业号）
// 排查方法: 需要成员使用微信登录企业微信或者关注微工作台才能获取openid
const ErrCode43004 ErrCode = 43004

// ErrCode43009 企业未验证主体
// 排查方法: -
const ErrCode43009 ErrCode = 43009

// ErrCode43012 应用需配置回调url
// 排查方法: -
const ErrCode43012 ErrCode = 43012

// ErrCode44001 多媒体文件为空
// 排查方法: 上传格式参考：[上传临时素材](https://work.weixin.qq.com/api/doc#10112)，确认header和body的内容正确。
const ErrCode44001 ErrCode = 44001

// ErrCode44004 文本消息content参数为空
// 排查方法: 发文本消息content为必填参数，且不能为空
const ErrCode44004 ErrCode = 44004

// ErrCode45001 多媒体文件大小超过限制
// 排查方法: 图片不可超过5M；音频不可超过5M；文件不可超过20M
const ErrCode45001 ErrCode = 45001

// ErrCode45002 消息内容大小超过限制
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：45002)
// 消息内容大小超过限制。确认：
// 1）文本消息类型：最长不超过2048个字节。
// 2）图文消息类型：最长不超过666k个字节
const ErrCode45002 ErrCode = 45002

// ErrCode45004 应用description参数长度不符合系统限制
// 排查方法: 设置应用若带有description参数，则长度必须为4至120个字符
const ErrCode45004 ErrCode = 45004

// ErrCode45007 语音播放时间超过限制
// 排查方法: 语音播放时长不能超过60秒
const ErrCode45007 ErrCode = 45007

// ErrCode45008 图文消息的文章数量不符合系统限制
// 排查方法: 图文消息的文章数量不能超过8条
const ErrCode45008 ErrCode = 45008

// ErrCode45009 接口调用超过限制
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：45009)
// 接口调用超过限制。
// 1) 具体频率策略，参考：[主动调用频率限制](https://open.work.weixin.qq.com/api/doc#10785)
// 2) 频率拦截时长一般与调用的限制时长相同，比如说是分钟级别的限制，则在中频率后的1分钟后自动解除。小时、天、以及月份，也是以此类推。
// 3) 我们对接口调用的频率限制是比较宽松的。对于接口中频率的调用，考虑以下优化：
// * 接口实现时，仅系统失败需要重试。其余错误码，应该排查下调用失败原因
// * 发消息应该控制合理调用，对于单个成员来说，一天收到大量的推送，体验是不好的
// 4) 部分频率拦截，可自助解封，访问：[频率自助解封工具](https://open.work.weixin.qq.com/wwopen/devtool/checkCorpSpamBlock)
// 5) 发送应用消息的频率拦截，可用api接口查询各个应用消息发送统计，访问：[查询应用消息发送统计](https://open.work.weixin.qq.com/api/doc/90000/90135/92369)
const ErrCode45009 ErrCode = 45009

// ErrCode45022 应用name参数长度不符合系统限制
// 排查方法: 设置应用若带有name参数，则不允许为空，且不超过32个字符
const ErrCode45022 ErrCode = 45022

// ErrCode45024 帐号数量超过上限
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：45024)
// 帐号数量超过上限。请确认：
// 1）通讯录是否有无效或者无用的帐号，可以删除，让出额度
// 2）提高帐号上限，可以提交重新认证或者申请扩容
const ErrCode45024 ErrCode = 45024

// ErrCode45026 触发删除用户数的保护
// 排查方法: 限制参考：[全量覆盖成员](https://work.weixin.qq.com/api/doc/90000/90139/90313#10138/全量覆盖成员)
const ErrCode45026 ErrCode = 45026

// ErrCode45029 回包大小超过上限
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：45029)
// 回包大小超过上限。请确认：
// 1）/cgi-bin/user/list：由于通讯录组织架构庞大，建议按部门分别拉取，同时不要指定fetch_child=1。
const ErrCode45029 ErrCode = 45029

// ErrCode45032 图文消息author参数长度超过限制
// 排查方法: 最长64个字节
const ErrCode45032 ErrCode = 45032

// ErrCode45033 接口并发调用超过限制
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：45033)
// 接口并发调用超过限制。出现这种拦截限制，一般是开发者的程序有bug，导致对同一份资源有过高的并发且持续不断的请求，例如对一个media_id一直持续不断请求“获取临时素材”接口。
const ErrCode45033 ErrCode = 45033

// ErrCode45034 url必须有协议头
// 排查方法: 在url前面加上协议头 http:// 或 https:const ErrCode45034 ErrCode = 45034

// ErrCode46003 菜单未设置
// 排查方法: 菜单需发布后才能获取到数据
const ErrCode46003 ErrCode = 46003

// ErrCode46004 指定的用户不存在
// 排查方法: 需要确认指定的用户存在于通讯录中
const ErrCode46004 ErrCode = 46004

// ErrCode48002 API接口无权限调用
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：48002)
// API接口无权限调用。请确认：
// 1）写通讯录接口，只能由通讯录同步助手的access_token来调用。同时需要保证通讯录同步功能是开启的。
// 2）通讯录同步助手的access_token，仅用于同步通讯录，不能用于发消息
// 3）设置应用可见范围，仅支持注册定制化安装情况，详情见：[设置授权应用可见范围](https://open.work.weixin.qq.com/api/doc#14936)
// 4）客户联系相关的接口，只能由系统应用“客户联系”，或配置到“可调用应用”列表中的自建应用的access_token来调用。
// 5) 小程序应用仅支持发送[小程序通知消息](https://work.weixin.qq.com/api/doc/90000/90135/90236#%E5%B0%8F%E7%A8%8B%E5%BA%8F%E9%80%9A%E7%9F%A5%E6%B6%88%E6%81%AF)，暂不支持文本、图片、语音、视频、图文等其他类型的消息。
const ErrCode48002 ErrCode = 48002

// ErrCode48003 不合法的suite_id
// 排查方法: 确认suite_access_token由指定的suite_id生成
const ErrCode48003 ErrCode = 48003

// ErrCode48004 授权关系无效
// 排查方法: 可能是无授权或授权已被取消
const ErrCode48004 ErrCode = 48004

// ErrCode48005 API接口已废弃
// 排查方法: 接口已不再支持，建议改用新接口或者新方案
const ErrCode48005 ErrCode = 48005

// ErrCode48006 接口权限被收回
// 排查方法: 由于企业长时间未使用应用，接口权限被收回，需企业管理员重新启用
const ErrCode48006 ErrCode = 48006

// ErrCode49004 签名不匹配
// 排查方法: -
const ErrCode49004 ErrCode = 49004

// ErrCode49008 群已经解散
// 排查方法: 群主已经解散群聊
const ErrCode49008 ErrCode = 49008

// ErrCode50001 redirect_url未登记可信域名
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：50001)
// redirect_url未登记可信域名。请确认：
// 1）颁发code的场景在哪个应用点击的。消费code使用的access_token是否有该应用权限。（通过[查询access_token权限](https://open.work.weixin.qq.com/devtool/query)可确认）
// 2）secret的获取来源
// * 来源于应用：url的域名，需设置到应用可信域名中。
// * 来源于通讯录同步助手：仅可同步通讯录，不可用于发消息或者消费code
// * 来源于第三方套件授权：套件中至少有一个应用，设置了该url域名为可信域名
// * 来源于管理组：管理组配置的应用列表，至少有一个应用设置了该url域名为可信域名
// 3）url填写的域名，必须与设置的可信域名 **完全匹配**（包括端口号）。比如：填可信域名填qq.com，访问url域名为www.qq.com，就不匹配；或者可信域名填www.qq.com，访问url域名为www.qq.com:8008，也不匹配。
const ErrCode50001 ErrCode = 50001

// ErrCode50002 成员不在权限范围
// 排查方法: 请检查应用或管理组的权限范围
const ErrCode50002 ErrCode = 50002

// ErrCode50003 应用已禁用
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：50003)
// 应用禁用之后，将无法再调用api，可在”管理端-应用管理”重新启用该应用。
// <img src="https://p.qpic.cn/pic_wework/1283325914/090fad19ebf740ce78b1e49f2cc0f0f8d1791c3262b8e946/0" style="width:800px;"/>
const ErrCode50003 ErrCode = 50003

// ErrCode50100 分页查询的游标无效
// 排查方法: -
const ErrCode50100 ErrCode = 50100

// ErrCode60001 部门长度不符合限制
// 排查方法: 部门名称不能为空且长度不能超过32个字
const ErrCode60001 ErrCode = 60001

// ErrCode60003 部门ID不存在
// 排查方法: 需要确认部门ID是否有带，并且存在通讯录中
const ErrCode60003 ErrCode = 60003

// ErrCode60004 父部门不存在
// 排查方法: 需要确认父亲部门ID是否有带，并且存在通讯录中
const ErrCode60004 ErrCode = 60004

// ErrCode60005 部门下存在成员
// 排查方法: 不允许删除有成员的部门
const ErrCode60005 ErrCode = 60005

// ErrCode60006 部门下存在子部门
// 排查方法: 不允许删除有子部门的部门
const ErrCode60006 ErrCode = 60006

// ErrCode60007 不允许删除根部门
// 排查方法: -
const ErrCode60007 ErrCode = 60007

// ErrCode60008 部门已存在
// 排查方法: 部门ID或者部门名称已存在
const ErrCode60008 ErrCode = 60008

// ErrCode60009 部门名称含有非法字符
// 排查方法: 不能含有 \:?*“<>| 等字符
const ErrCode60009 ErrCode = 60009

// ErrCode60010 部门存在循环关系
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：60010)
// 部门存在循环关系。请确认：
// 1）创建部门和更新部门时，指定的parentid参数不能是 部门id 或者 子部门id
const ErrCode60010 ErrCode = 60010

// ErrCode60011 指定的成员/部门/标签参数无权限
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：60011)
// 指定的成员/部门/标签参数无权限。请确认：
// 1) 变更通讯录接口，需要有通讯录编辑权限。
// * 普通应用的secret仅有只读权限，可使用通讯录同步助手的secret同步。
// 2) 其它接口，需要满足配置的通讯录范围。
// * 成员：通讯录同步助手access_token可指定任意成员id；应用access_token仅能指定可见范围配置的成员，以及部门/标签包含的成员（递归展开）
// * 部门：通讯录同步助手access_token可指定任意部门id；应用access_token仅能指定可见范围配置的部门id(创建或移动部门，还需要具有父部门的管理权限)，标签包括的部门id，以及上述部门的子部门id
// * 标签：通讯录同步助手access_token可指定超级管理组及通讯录同步助手创建的标签；应用access_token仅能由应用API创建的标签
const ErrCode60011 ErrCode = 60011

// ErrCode60012 不允许删除默认应用
// 排查方法: 默认应用的id为0
const ErrCode60012 ErrCode = 60012

// ErrCode60020 访问ip不在白名单之中
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：60020)
// 访问ip不在白名单之中。请确认：
// 1）请确认访问ip是否在服务商白名单IP列表。
// 登录 [服务商管理后台](https://open.work.weixin.qq.com/wwopen/login)，在“服务商信息” - “基本信息” - “IP白名单”配置
const ErrCode60020 ErrCode = 60020

// ErrCode60021 userid不在应用可见范围内
// 排查方法: -
const ErrCode60021 ErrCode = 60021

// ErrCode60028 不允许修改第三方应用的主页 URL
// 排查方法: 第三方应用类型，不允许通过接口修改该应用的主页 URL
const ErrCode60028 ErrCode = 60028

// ErrCode60102 UserID已存在
// 排查方法: -
const ErrCode60102 ErrCode = 60102

// ErrCode60103 手机号码不合法
// 排查方法: 长度不超过32位，字符仅支持数字，加号和减号
const ErrCode60103 ErrCode = 60103

// ErrCode60104 手机号码已存在
// 排查方法: 同一个企业内，成员的手机号不能重复。建议更换手机号，或者更新已有的手机记录。
const ErrCode60104 ErrCode = 60104

// ErrCode60105 邮箱不合法
// 排查方法: 长度不超过64位，且为有效的email格式
const ErrCode60105 ErrCode = 60105

// ErrCode60106 邮箱已存在
// 排查方法: 同一个企业内，成员的邮箱不能重复。建议更换邮箱，或者更新已有的邮箱记录。
const ErrCode60106 ErrCode = 60106

// ErrCode60107 微信号不合法
// 排查方法: 微信号格式由字母、数字、”-“、”_“组成，长度为 3-20 字节，首字符必须是字母或”-“或”_“
const ErrCode60107 ErrCode = 60107

// ErrCode60110 用户所属部门数量超过限制
// 排查方法: 用户同时归属部门不超过20个
const ErrCode60110 ErrCode = 60110

// ErrCode60111 UserID不存在
// 排查方法: UserID参数为空，或者不存在通讯录中
const ErrCode60111 ErrCode = 60111

// ErrCode60112 成员name参数不合法
// 排查方法: 不能为空，且不能超过64字符
const ErrCode60112 ErrCode = 60112

// ErrCode60123 无效的部门id
// 排查方法: 部门不存在通讯录中
const ErrCode60123 ErrCode = 60123

// ErrCode60124 无效的父部门id
// 排查方法: 父部门不存在通讯录中
const ErrCode60124 ErrCode = 60124

// ErrCode60125 非法部门名字
// 排查方法: 不能为空，且不能超过64字节，且不能含有\:*?”<>|等字符
const ErrCode60125 ErrCode = 60125

// ErrCode60127 缺少department参数
// 排查方法: -
const ErrCode60127 ErrCode = 60127

// ErrCode60129 成员手机和邮箱都为空
// 排查方法: 成员手机和邮箱至少有个非空
const ErrCode60129 ErrCode = 60129

// ErrCode60132 is_leader_in_dept和department的元素个数不一致
// 排查方法: -
const ErrCode60132 ErrCode = 60132

// ErrCode60136 记录不存在
// 排查方法: -
const ErrCode60136 ErrCode = 60136

// ErrCode60137 家长手机号重复
// 排查方法: 同一个家校通讯录中，家长的手机号不能重复。建议更换手机号，或者更新已有的手机记录。
const ErrCode60137 ErrCode = 60137

// ErrCode60203 不合法的模版ID
// 排查方法: -
const ErrCode60203 ErrCode = 60203

// ErrCode60204 模版状态不可用
// 排查方法: -
const ErrCode60204 ErrCode = 60204

// ErrCode60205 模版关键词不匹配
// 排查方法: -
const ErrCode60205 ErrCode = 60205

// ErrCode60206 该种类型的消息只支持第三方独立应用使用
// 排查方法: -
const ErrCode60206 ErrCode = 60206

// ErrCode60207 第三方独立应用只允许发送模板消息
// 排查方法: -
const ErrCode60207 ErrCode = 60207

// ErrCode60208 第三方独立应用不支持指定@all，不支持参数toparty和totag
// 排查方法: -
const ErrCode60208 ErrCode = 60208

// ErrCode60209 缺少操作者vid
// 排查方法: -
const ErrCode60209 ErrCode = 60209

// ErrCode60210 选择成员列表为空
// 排查方法: -
const ErrCode60210 ErrCode = 60210

// ErrCode60211 SelectedTicket为空
// 排查方法: -
const ErrCode60211 ErrCode = 60211

// ErrCode60214 仅支持第三方应用调用
// 排查方法: -
const ErrCode60214 ErrCode = 60214

// ErrCode60215 传入SelectedTicket数量超过最大限制（10个）
// 排查方法: -
const ErrCode60215 ErrCode = 60215

// ErrCode60217 当前操作者无权限，操作者需要授权或者在可见范围内
// 排查方法: -
const ErrCode60217 ErrCode = 60217

// ErrCode60218 仅支持成员授权模式的应用可调用
// 排查方法: -
const ErrCode60218 ErrCode = 60218

// ErrCode60219 消费SelectedTicket和创建SelectedTicket的应用appid不匹配
// 排查方法: -
const ErrCode60219 ErrCode = 60219

// ErrCode60220 缺少corpappid
// 排查方法: -
const ErrCode60220 ErrCode = 60220

// ErrCode60221 open_userid对应的服务商不是当前服务商
// 排查方法: -
const ErrCode60221 ErrCode = 60221

// ErrCode60222 非法SelectedTicket
// 排查方法: -
const ErrCode60222 ErrCode = 60222

// ErrCode60223 非法BundleId
// 排查方法: -
const ErrCode60223 ErrCode = 60223

// ErrCode60224 非法PackageName
// 排查方法: -
const ErrCode60224 ErrCode = 60224

// ErrCode60225 当前操作者并非SelectedTicket相关人，不能创建群聊
// 排查方法: -
const ErrCode60225 ErrCode = 60225

// ErrCode60226 选人数量超过最大限制（2000）
// 排查方法: -
const ErrCode60226 ErrCode = 60226

// ErrCode60227 缺少ServiceCorpid
// 排查方法: -
const ErrCode60227 ErrCode = 60227

// ErrCode65000 学校已经迁移
// 排查方法: -
const ErrCode65000 ErrCode = 65000

// ErrCode65001 无效的关注模式
// 排查方法: -
const ErrCode65001 ErrCode = 65001

// ErrCode65002 导入家长信息数量过多
// 排查方法: 批量导入家长每次最多1000个
const ErrCode65002 ErrCode = 65002

// ErrCode65003 学校尚未迁移
// 排查方法: -
const ErrCode65003 ErrCode = 65003

// ErrCode65004 组织架构不存在
// 排查方法: -
const ErrCode65004 ErrCode = 65004

// ErrCode65005 无效的同步模式
// 排查方法: -
const ErrCode65005 ErrCode = 65005

// ErrCode65006 无效的管理员类型
// 排查方法: -
const ErrCode65006 ErrCode = 65006

// ErrCode65007 无效的家校部门类型
// 排查方法: -
const ErrCode65007 ErrCode = 65007

// ErrCode65008 无效的入学年份
// 排查方法: -
const ErrCode65008 ErrCode = 65008

// ErrCode65009 无效的标准年级类型
// 排查方法: -
const ErrCode65009 ErrCode = 65009

// ErrCode65010 此userid并不是学生
// 排查方法: -
const ErrCode65010 ErrCode = 65010

// ErrCode65011 家长userid数量超过限制
// 排查方法: 每次最多批量处理100个家长
const ErrCode65011 ErrCode = 65011

// ErrCode65012 学生userid数量超过限制
// 排查方法: 每次最多批量处理10个学生
const ErrCode65012 ErrCode = 65012

// ErrCode65013 学生已有家长
// 排查方法: -
const ErrCode65013 ErrCode = 65013

// ErrCode65014 非学校企业
// 排查方法: -
const ErrCode65014 ErrCode = 65014

// ErrCode65015 父部门类型不匹配
// 排查方法: 添加学校部门，需满足层级关机，班级需要以年级为父部门
const ErrCode65015 ErrCode = 65015

// ErrCode65018 家长人数达到上限
// 排查方法: 未验证的学校\企业最多可添加2000名家长，验证过的学校\企业最多可添加20000名家长
const ErrCode65018 ErrCode = 65018

// ErrCode660001 无效的商户号
// 排查方法: 请检查商户号是否正确
const ErrCode660001 ErrCode = 660001

// ErrCode660002 无效的企业收款人id
// 排查方法: 请检查payee_userid是否正确
const ErrCode660002 ErrCode = 660002

// ErrCode660003 userid不在应用的可见范围
// 排查方法: -
const ErrCode660003 ErrCode = 660003

// ErrCode660004 partyid不在应用的可见范围
// 排查方法: -
const ErrCode660004 ErrCode = 660004

// ErrCode660005 tagid不在应用的可见范围
// 排查方法: -
const ErrCode660005 ErrCode = 660005

// ErrCode660006 找不到该商户号
// 排查方法: -
const ErrCode660006 ErrCode = 660006

// ErrCode660007 申请已经存在
// 排查方法: 不需要重复申请
const ErrCode660007 ErrCode = 660007

// ErrCode660008 商户号已经绑定
// 排查方法: 不需要重新提交申请
const ErrCode660008 ErrCode = 660008

// ErrCode660009 商户号主体和商户主体不一致
// 排查方法: -
const ErrCode660009 ErrCode = 660009

// ErrCode660010 超过商户号绑定数量限制
// 排查方法: -
const ErrCode660010 ErrCode = 660010

// ErrCode660011 商户号未绑定
// 排查方法: -
const ErrCode660011 ErrCode = 660011

// ErrCode670001 应用不在共享范围
// 排查方法: -
const ErrCode670001 ErrCode = 670001

// ErrCode72023 发票已被其他公众号锁定
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：72023)
// 一般为发票已进入后续报销流程，报销企业公众号/企业微信/App锁定了发票。
const ErrCode72023 ErrCode = 72023

// ErrCode72024 发票状态错误
// 排查方法: reimburse_status状态错误，参考：[更新发票状态](https://work.weixin.qq.com/api/doc/90000/90139/90313#11633)
const ErrCode72024 ErrCode = 72024

// ErrCode72037 存在发票不属于该用户
// 排查方法: 只能批量更新该openid的发票，参考：[批量更新发票状态](https://work.weixin.qq.com/api/doc/90000/90139/90313#11634)
const ErrCode72037 ErrCode = 72037

// ErrCode80001 可信域名不正确，或者无ICP备案
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：80001)
// 可信域名不正确，未校验域名所有权归属或者可信域名没有ICP备案。请确认：
// 1）可信域名，只支持全域名匹配，无法通过配置父域来让所有子域都成为可信域名。
// 2）可信域名，不支持IP地址、端口号及短链域名。
// 3）如果确认域名已经通过ICP备案，但依然提示这个错误，请尝试重新设置。
const ErrCode80001 ErrCode = 80001

// ErrCode81001 部门下的结点数超过限制（3W）
// 排查方法: -
const ErrCode81001 ErrCode = 81001

// ErrCode81002 部门最多15层
// 排查方法: -
const ErrCode81002 ErrCode = 81002

// ErrCode81003 标签下节点个数超过30000个
// 排查方法: -
const ErrCode81003 ErrCode = 81003

// ErrCode81011 无权限操作标签
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：81011)
// 无权限操作标签。请确认：
// 1）除了通讯录同步助手和通讯录应用，其他应用和管理组都只能操作自己创建的标签。
// 2）通讯录同步助手或者通讯录应用，除了能管理自己的标签，还能操作超级管理组创建的标签。
const ErrCode81011 ErrCode = 81011

// ErrCode81012 缺失可见范围
// 排查方法: 请求没有填写UserID、部门ID、标签ID
const ErrCode81012 ErrCode = 81012

// ErrCode81013 UserID、部门ID、标签ID全部非法或无权限
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：81013)
// UserID、部门ID、标签ID全部非法或无权限。一般有以下两种原因：
// 1）成员、部门或标签已被删除，此种情况需要调整调用接口的接收人参数。
// 2）成员、部门或标签被移出应用的可见范围，可在管理端将接收人添加到应用的可见范围内。
// <img src="https://p.qpic.cn/pic_wework/1283325914/4e48eb5ea14e98977a1687c6cceab0d61ad57eea9af062f0/0" style="width:800px;"/>
const ErrCode81013 ErrCode = 81013

// ErrCode81014 标签添加成员，单次添加user或party过多
// 排查方法: -
const ErrCode81014 ErrCode = 81014

// ErrCode81015 邮箱域名需要跟企业邮箱域名一致
// 排查方法: -
const ErrCode81015 ErrCode = 81015

// ErrCode81016 logined_userid字段缺失
// 排查方法: -
const ErrCode81016 ErrCode = 81016

// ErrCode81017 请求个数超过限制
// 排查方法: -
const ErrCode81017 ErrCode = 81017

// ErrCode81018 该服务商可获取名字数量配额不足
// 排查方法: -
const ErrCode81018 ErrCode = 81018

// ErrCode81019 items数组成员缺少id字段
// 排查方法: -
const ErrCode81019 ErrCode = 81019

// ErrCode81020 items数组成员缺少type字段
// 排查方法: -
const ErrCode81020 ErrCode = 81020

// ErrCode81021 items数组成员的type字段不合法
// 排查方法: -
const ErrCode81021 ErrCode = 81021

// ErrCode82001 指定的成员/部门/标签全部为空
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：82001)
// 指定的成员/部门/标签全部为空。请确认：
// 参数是否有传递，且至少有一个参数非空。
const ErrCode82001 ErrCode = 82001

// ErrCode82002 不合法的PartyID列表长度
// 排查方法: 发消息，单次不能超过100个部门
const ErrCode82002 ErrCode = 82002

// ErrCode82003 不合法的TagID列表长度
// 排查方法: 发消息，单次不能超过100个标签
const ErrCode82003 ErrCode = 82003

// ErrCode82004 不合法的消息内容
// 排查方法: 消息内容中可能存在使客户端crash的内容
const ErrCode82004 ErrCode = 82004

// ErrCode84014 成员票据过期
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：84014)
// 成员票据过期。确认：
// 1）user_ticket 有时效性，有效时长由expires_in指定。参考接口：[根据code获取成员信息](https://open.work.weixin.qq.com/api/doc#10028/根据code获取成员信息)
// 2）若需再次获取用户详情，需要用户重新点击链接后，根据新的code获取新的user_ticket
const ErrCode84014 ErrCode = 84014

// ErrCode84015 成员票据无效
// 排查方法: 确认user_ticket参数来源是否正确。参考接口：[根据code获取成员信息](https://work.weixin.qq.com/api/doc/90000/90139/90313#10028/根据code获取成员信息)
const ErrCode84015 ErrCode = 84015

// ErrCode84019 缺少templateid参数
// 排查方法: -
const ErrCode84019 ErrCode = 84019

// ErrCode84020 templateid不存在
// 排查方法: 确认参数是否有带，并且已创建
const ErrCode84020 ErrCode = 84020

// ErrCode84021 缺少register_code参数
// 排查方法: -
const ErrCode84021 ErrCode = 84021

// ErrCode84022 无效的register_code参数
// 排查方法: -
const ErrCode84022 ErrCode = 84022

// ErrCode84023 不允许调用设置通讯录同步完成接口
// 排查方法: -
const ErrCode84023 ErrCode = 84023

// ErrCode84024 无注册信息
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：84024)
// 无注册信息。可能是以下两种情况：
// 1）注册流程未完成。
// 2）注册成功已超过24小时。
const ErrCode84024 ErrCode = 84024

// ErrCode84025 不符合的state参数
// 排查方法: 必须是[a-zA-Z0-9]的参数值，长度不可超过128个字节
const ErrCode84025 ErrCode = 84025

// ErrCode84052 缺少caller参数
// 排查方法: -
const ErrCode84052 ErrCode = 84052

// ErrCode84053 缺少callee参数
// 排查方法: -
const ErrCode84053 ErrCode = 84053

// ErrCode84054 缺少auth_corpid参数
// 排查方法: -
const ErrCode84054 ErrCode = 84054

// ErrCode84055 超过拨打公费电话频率
// 排查方法: 同一个客服5秒内只能调用api拨打一次公费电话
const ErrCode84055 ErrCode = 84055

// ErrCode84056 被拨打用户安装应用时未授权拨打公费电话权限
// 排查方法: -
const ErrCode84056 ErrCode = 84056

// ErrCode84057 公费电话余额不足
// 排查方法: -
const ErrCode84057 ErrCode = 84057

// ErrCode84058 caller 呼叫号码不支持
// 排查方法: -
const ErrCode84058 ErrCode = 84058

// ErrCode84059 号码非法
// 排查方法: -
const ErrCode84059 ErrCode = 84059

// ErrCode84060 callee 呼叫号码不支持
// 排查方法: -
const ErrCode84060 ErrCode = 84060

// ErrCode84061 不存在外部联系人的关系
// 排查方法: -
const ErrCode84061 ErrCode = 84061

// ErrCode84062 未开启公费电话应用
// 排查方法: -
const ErrCode84062 ErrCode = 84062

// ErrCode84063 caller不存在
// 排查方法: -
const ErrCode84063 ErrCode = 84063

// ErrCode84064 callee不存在
// 排查方法: -
const ErrCode84064 ErrCode = 84064

// ErrCode84065 caller跟callee电话号码一致
// 排查方法: 不允许自己拨打给自己
const ErrCode84065 ErrCode = 84065

// ErrCode84066 服务商拨打次数超过限制
// 排查方法: 单个企业管理员，在一天（以上午10:00为起始时间）内，对应单个服务商，只能被呼叫【4】次。
const ErrCode84066 ErrCode = 84066

// ErrCode84067 管理员收到的服务商公费电话个数超过限制
// 排查方法: 单个企业管理员，在一天（以上午10:00为起始时间）内，一共只能被【3】个服务商成功呼叫。
const ErrCode84067 ErrCode = 84067

// ErrCode84069 拨打方被限制拨打公费电话
// 排查方法: -
const ErrCode84069 ErrCode = 84069

// ErrCode84070 不支持的电话号码
// 排查方法: 拨打方或者被拨打方电话号码不支持
const ErrCode84070 ErrCode = 84070

// ErrCode84071 不合法的外部联系人授权码
// 排查方法: 非法或者已经消费过
const ErrCode84071 ErrCode = 84071

// ErrCode84072 应用未配置客服
// 排查方法: -
const ErrCode84072 ErrCode = 84072

// ErrCode84073 客服userid不在应用配置的客服列表中
// 排查方法: -
const ErrCode84073 ErrCode = 84073

// ErrCode84074 没有外部联系人权限
// 排查方法: -
const ErrCode84074 ErrCode = 84074

// ErrCode84075 不合法或过期的authcode
// 排查方法: -
const ErrCode84075 ErrCode = 84075

// ErrCode84076 缺失authcode
// 排查方法: -
const ErrCode84076 ErrCode = 84076

// ErrCode84077 订单价格过高，无法受理
// 排查方法: -
const ErrCode84077 ErrCode = 84077

// ErrCode84078 购买人数不正确
// 排查方法: -
const ErrCode84078 ErrCode = 84078

// ErrCode84079 价格策略不存在
// 排查方法: -
const ErrCode84079 ErrCode = 84079

// ErrCode84080 订单不存在
// 排查方法: -
const ErrCode84080 ErrCode = 84080

// ErrCode84081 存在未支付订单
// 排查方法: -
const ErrCode84081 ErrCode = 84081

// ErrCode84082 存在申请退款中的订单
// 排查方法: -
const ErrCode84082 ErrCode = 84082

// ErrCode84083 非服务人员
// 排查方法: -
const ErrCode84083 ErrCode = 84083

// ErrCode84084 非跟进用户
// 排查方法: -
const ErrCode84084 ErrCode = 84084

// ErrCode84085 应用已下架
// 排查方法: -
const ErrCode84085 ErrCode = 84085

// ErrCode84086 订单人数超过可购买最大人数
// 排查方法: -
const ErrCode84086 ErrCode = 84086

// ErrCode84087 打开订单支付前禁止关闭订单
// 排查方法: -
const ErrCode84087 ErrCode = 84087

// ErrCode84088 禁止关闭已支付的订单
// 排查方法: -
const ErrCode84088 ErrCode = 84088

// ErrCode84089 订单已支付
// 排查方法: -
const ErrCode84089 ErrCode = 84089

// ErrCode84090 缺失user_ticket
// 排查方法: -
const ErrCode84090 ErrCode = 84090

// ErrCode84091 订单价格不可低于下限
// 排查方法: -
const ErrCode84091 ErrCode = 84091

// ErrCode84092 无法发起代下单操作
// 排查方法: -
const ErrCode84092 ErrCode = 84092

// ErrCode84093 代理关系已占用，无法代下单
// 排查方法: -
const ErrCode84093 ErrCode = 84093

// ErrCode84094 该应用未配置代理分润规则，请先联系应用服务商处理
// 排查方法: -
const ErrCode84094 ErrCode = 84094

// ErrCode84095 免费试用版，无法扩容
// 排查方法: -
const ErrCode84095 ErrCode = 84095

// ErrCode84096 免费试用版，无法续期
// 排查方法: -
const ErrCode84096 ErrCode = 84096

// ErrCode84097 当前企业有未处理订单
// 排查方法: -
const ErrCode84097 ErrCode = 84097

// ErrCode84098 固定总量，无法扩容
// 排查方法: -
const ErrCode84098 ErrCode = 84098

// ErrCode84099 非购买状态，无法扩容
// 排查方法: -
const ErrCode84099 ErrCode = 84099

// ErrCode84100 未购买过此应用，无法续期
// 排查方法: -
const ErrCode84100 ErrCode = 84100

// ErrCode84101 企业已试用付费版本，无法全新购买
// 排查方法: -
const ErrCode84101 ErrCode = 84101

// ErrCode84102 企业当前应用状态已过期，无法扩容
// 排查方法: -
const ErrCode84102 ErrCode = 84102

// ErrCode84103 仅可修改未支付订单
// 排查方法: -
const ErrCode84103 ErrCode = 84103

// ErrCode84104 订单已支付，无法修改
// 排查方法: -
const ErrCode84104 ErrCode = 84104

// ErrCode84105 订单已被取消，无法修改
// 排查方法: -
const ErrCode84105 ErrCode = 84105

// ErrCode84106 企业含有该应用的待支付订单，无法代下单
// 排查方法: -
const ErrCode84106 ErrCode = 84106

// ErrCode84107 企业含有该应用的退款中订单，无法代下单
// 排查方法: -
const ErrCode84107 ErrCode = 84107

// ErrCode84108 企业含有该应用的待生效订单，无法代下单
// 排查方法: -
const ErrCode84108 ErrCode = 84108

// ErrCode84109 订单定价不能未0
// 排查方法: -
const ErrCode84109 ErrCode = 84109

// ErrCode84110 新安装应用不在试用状态，无法升级为付费版
// 排查方法: -
const ErrCode84110 ErrCode = 84110

// ErrCode84111 无足够可用优惠券
// 排查方法: -
const ErrCode84111 ErrCode = 84111

// ErrCode84112 无法关闭未支付订单
// 排查方法: -
const ErrCode84112 ErrCode = 84112

// ErrCode84113 无付费信息
// 排查方法: -
const ErrCode84113 ErrCode = 84113

// ErrCode84114 虚拟版本不支持下单
// 排查方法: -
const ErrCode84114 ErrCode = 84114

// ErrCode84115 虚拟版本不支持扩容
// 排查方法: -
const ErrCode84115 ErrCode = 84115

// ErrCode84116 虚拟版本不支持续期
// 排查方法: -
const ErrCode84116 ErrCode = 84116

// ErrCode84117 在虚拟正式版期内不能扩容
// 排查方法: -
const ErrCode84117 ErrCode = 84117

// ErrCode84118 虚拟正式版期内不能变更版本
// 排查方法: -
const ErrCode84118 ErrCode = 84118

// ErrCode84119 当前企业未报备，无法进行代下单
// 排查方法: -
const ErrCode84119 ErrCode = 84119

// ErrCode84120 当前应用版本已删除
// 排查方法: -
const ErrCode84120 ErrCode = 84120

// ErrCode84121 应用版本已删除，无法扩容
// 排查方法: -
const ErrCode84121 ErrCode = 84121

// ErrCode84122 应用版本已删除，无法续期
// 排查方法: -
const ErrCode84122 ErrCode = 84122

// ErrCode84123 非虚拟版本，无法升级
// 排查方法: -
const ErrCode84123 ErrCode = 84123

// ErrCode84124 非行业方案订单，不能添加部分应用版本的订单
// 排查方法: -
const ErrCode84124 ErrCode = 84124

// ErrCode84125 购买人数不能少于最少购买人数
// 排查方法: -
const ErrCode84125 ErrCode = 84125

// ErrCode84126 购买人数不能多于最大购买人数
// 排查方法: -
const ErrCode84126 ErrCode = 84126

// ErrCode84127 无应用管理权限
// 排查方法: -
const ErrCode84127 ErrCode = 84127

// ErrCode84128 无该行业方案下全部应用的管理权限
// 排查方法: -
const ErrCode84128 ErrCode = 84128

// ErrCode84129 付费策略已被删除，无法下单
// 排查方法: -
const ErrCode84129 ErrCode = 84129

// ErrCode84130 订单生效时间不合法
// 排查方法: -
const ErrCode84130 ErrCode = 84130

// ErrCode84200 文件转译解析错误
// 排查方法: 只支持utf8文件转译，可能是不支持的文件类型或者格式
const ErrCode84200 ErrCode = 84200

// ErrCode85002 包含不合法的词语
// 排查方法: -
const ErrCode85002 ErrCode = 85002

// ErrCode85004 每企业每个月设置的可信域名不可超过20个
// 排查方法: -
const ErrCode85004 ErrCode = 85004

// ErrCode85005 可信域名未通过所有权校验
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：85005)
// 域名未通过所有权校验，仅oauth2生效，jssdk功能将受限，请根据调用者身份按以下不同方式完成校验：
// 1）若调用者是企业应用，请登录企业微信管理端，进入应用详情，按照指引完成域名的所有权校验。
// 2）若调用者是第三方服务，请登录企业微信服务管理端，进入第三方应用详情，按照指引完成域名的所有权校验。
const ErrCode85005 ErrCode = 85005

// ErrCode86001 参数 chatid 不合法
// 排查方法: -
const ErrCode86001 ErrCode = 86001

// ErrCode86003 参数 chatid 不存在
// 排查方法: -
const ErrCode86003 ErrCode = 86003

// ErrCode86004 参数 群名不合法
// 排查方法: -
const ErrCode86004 ErrCode = 86004

// ErrCode86005 参数 群主不合法
// 排查方法: -
const ErrCode86005 ErrCode = 86005

// ErrCode86006 群成员数过多或过少
// 排查方法: -
const ErrCode86006 ErrCode = 86006

// ErrCode86007 不合法的群成员
// 排查方法: -
const ErrCode86007 ErrCode = 86007

// ErrCode86008 非法操作非自己创建的群
// 排查方法: -
const ErrCode86008 ErrCode = 86008

// ErrCode86101 仅群主才有操作权限
// 排查方法: -
const ErrCode86101 ErrCode = 86101

// ErrCode86201 参数 需要chatid
// 排查方法: -
const ErrCode86201 ErrCode = 86201

// ErrCode86202 参数 需要群名
// 排查方法: -
const ErrCode86202 ErrCode = 86202

// ErrCode86203 参数 需要群主
// 排查方法: -
const ErrCode86203 ErrCode = 86203

// ErrCode86204 参数 需要群成员
// 排查方法: -
const ErrCode86204 ErrCode = 86204

// ErrCode86205 参数 字符串chatid过长
// 排查方法: -
const ErrCode86205 ErrCode = 86205

// ErrCode86206 参数 数字chatid过大
// 排查方法: -
const ErrCode86206 ErrCode = 86206

// ErrCode86207 群主不在群成员列表
// 排查方法: -
const ErrCode86207 ErrCode = 86207

// ErrCode86214 群发类型不合法
// 排查方法: -
const ErrCode86214 ErrCode = 86214

// ErrCode86215 会话ID已经存在
// 排查方法: -
const ErrCode86215 ErrCode = 86215

// ErrCode86216 存在非法会话成员ID
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：86216)
// 存在非法会话成员ID。确认：
// 1）添加会话成员时，指定的成员ID不存在通讯录
// 2）删除会话成员时，指定的成员ID不存在于会话中
const ErrCode86216 ErrCode = 86216

// ErrCode86217 会话发送者不在会话成员列表中
// 排查方法: 会话的发送者，必须是会话的成员列表之一
const ErrCode86217 ErrCode = 86217

// ErrCode86220 指定的会话参数不合法
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：86220)
// 指定的会话参数不合法。请确认：
// 1）参数 type 只能指定 single/group
// 2）参数 msgtype 只能指定 text/image/file/voice/link
const ErrCode86220 ErrCode = 86220

// ErrCode86224 不是受限群，不允许使用该接口
// 排查方法: -
const ErrCode86224 ErrCode = 86224

// ErrCode90001 未认证摇一摇周边
// 排查方法: -
const ErrCode90001 ErrCode = 90001

// ErrCode90002 缺少摇一摇周边ticket参数
// 排查方法: -
const ErrCode90002 ErrCode = 90002

// ErrCode90003 摇一摇周边ticket参数不合法
// 排查方法: -
const ErrCode90003 ErrCode = 90003

// ErrCode90100 非法的对外属性类型
// 排查方法: -
const ErrCode90100 ErrCode = 90100

// ErrCode90101 对外属性：文本类型长度不合法
// 排查方法: 文本长度不可超过12个UTF8字符
const ErrCode90101 ErrCode = 90101

// ErrCode90102 对外属性：网页类型标题长度不合法
// 排查方法: 标题长度不可超过12个UTF8字符
const ErrCode90102 ErrCode = 90102

// ErrCode90103 对外属性：网页url不合法
// 排查方法: -
const ErrCode90103 ErrCode = 90103

// ErrCode90104 对外属性：小程序类型标题长度不合法
// 排查方法: 标题长度不可超过12个UTF8字符
const ErrCode90104 ErrCode = 90104

// ErrCode90105 对外属性：小程序类型pagepath不合法
// 排查方法: -
const ErrCode90105 ErrCode = 90105

// ErrCode90106 对外属性：请求参数不合法
// 排查方法: -
const ErrCode90106 ErrCode = 90106

// ErrCode90200 缺少小程序appid参数
// 排查方法: -
const ErrCode90200 ErrCode = 90200

// ErrCode90201 小程序通知的content_item个数超过限制
// 排查方法: item个数不能超过10个
const ErrCode90201 ErrCode = 90201

// ErrCode90202 小程序通知中的key长度不合法
// 排查方法: 不能为空或超过10个汉字
const ErrCode90202 ErrCode = 90202

// ErrCode90203 小程序通知中的value长度不合法
// 排查方法: 不能为空或超过30个汉字
const ErrCode90203 ErrCode = 90203

// ErrCode90204 小程序通知中的page参数不合法
// 排查方法: -
const ErrCode90204 ErrCode = 90204

// ErrCode90206 小程序未关联到企业中
// 排查方法: -
const ErrCode90206 ErrCode = 90206

// ErrCode90207 不合法的小程序appid
// 排查方法: -
const ErrCode90207 ErrCode = 90207

// ErrCode90208 小程序appid不匹配
// 排查方法: -
const ErrCode90208 ErrCode = 90208

// ErrCode90300 orderid 不合法
// 排查方法: -
const ErrCode90300 ErrCode = 90300

// ErrCode90302 付费应用已过期
// 排查方法: -
const ErrCode90302 ErrCode = 90302

// ErrCode90303 付费应用超过最大使用人数
// 排查方法: -
const ErrCode90303 ErrCode = 90303

// ErrCode90304 订单中心服务异常，请稍后重试
// 排查方法: -
const ErrCode90304 ErrCode = 90304

// ErrCode90305 参数错误，errmsg中有提示具体哪个参数有问题
// 排查方法: -
const ErrCode90305 ErrCode = 90305

// ErrCode90306 商户设置不合法，详情请见errmsg
// 排查方法: -
const ErrCode90306 ErrCode = 90306

// ErrCode90307 登录态过期
// 排查方法: -
const ErrCode90307 ErrCode = 90307

// ErrCode90308 在开启IP鉴权的前提下，识别为无效的请求IP
// 排查方法: -
const ErrCode90308 ErrCode = 90308

// ErrCode90309 订单已经存在，请勿重复下单
// 排查方法: -
const ErrCode90309 ErrCode = 90309

// ErrCode90310 找不到订单
// 排查方法: -
const ErrCode90310 ErrCode = 90310

// ErrCode90311 关单失败, 可能原因：该单并没被拉起支付页面; 已经关单；已经支付；渠道失败；单处于保护状态；等等
// 排查方法: -
const ErrCode90311 ErrCode = 90311

// ErrCode90312 退款请求失败, 详情请看errmsg
// 排查方法: -
const ErrCode90312 ErrCode = 90312

// ErrCode90313 退款调用频率限制，超过规定的阈值
// 排查方法: -
const ErrCode90313 ErrCode = 90313

// ErrCode90314 订单状态错误，可能未支付，或者当前状态操作受限
// 排查方法: -
const ErrCode90314 ErrCode = 90314

// ErrCode90315 退款请求失败，主键冲突，请核实退款refund_id是否已使用
// 排查方法: -
const ErrCode90315 ErrCode = 90315

// ErrCode90316 退款原因编号不对
// 排查方法: -
const ErrCode90316 ErrCode = 90316

// ErrCode90317 尚未注册成为供应商
// 排查方法: -
const ErrCode90317 ErrCode = 90317

// ErrCode90318 参数nonce_str 为空或者重复，判定为重放攻击
// 排查方法: -
const ErrCode90318 ErrCode = 90318

// ErrCode90319 时间戳为空或者与系统时间间隔太大
// 排查方法: -
const ErrCode90319 ErrCode = 90319

// ErrCode90320 订单token无效
// 排查方法: -
const ErrCode90320 ErrCode = 90320

// ErrCode90321 订单token已过有效时间
// 排查方法: -
const ErrCode90321 ErrCode = 90321

// ErrCode90322 旧套件（包含多个应用的套件）不支持支付系统
// 排查方法: -
const ErrCode90322 ErrCode = 90322

// ErrCode90323 单价超过限额
// 排查方法: -
const ErrCode90323 ErrCode = 90323

// ErrCode90324 商品数量超过限额
// 排查方法: -
const ErrCode90324 ErrCode = 90324

// ErrCode90325 预支单已经存在
// 排查方法: -
const ErrCode90325 ErrCode = 90325

// ErrCode90326 预支单单号非法
// 排查方法: -
const ErrCode90326 ErrCode = 90326

// ErrCode90327 该预支单已经结算下单
// 排查方法: -
const ErrCode90327 ErrCode = 90327

// ErrCode90328 结算下单失败，详情请看errmsg
// 排查方法: -
const ErrCode90328 ErrCode = 90328

// ErrCode90329 该订单号已经被预支单占用
// 排查方法: -
const ErrCode90329 ErrCode = 90329

// ErrCode90330 创建供应商失败
// 排查方法: -
const ErrCode90330 ErrCode = 90330

// ErrCode90331 更新供应商失败
// 排查方法: -
const ErrCode90331 ErrCode = 90331

// ErrCode90332 还没签署合同
// 排查方法: -
const ErrCode90332 ErrCode = 90332

// ErrCode90333 创建合同失败
// 排查方法: -
const ErrCode90333 ErrCode = 90333

// ErrCode90338 已经过了可退款期限
// 排查方法: -
const ErrCode90338 ErrCode = 90338

// ErrCode90339 供应商主体名包含非法字符
// 排查方法: -
const ErrCode90339 ErrCode = 90339

// ErrCode90340 创建客户失败，可能信息真实性校验失败
// 排查方法: -
const ErrCode90340 ErrCode = 90340

// ErrCode90341 退款金额大于付款金额
// 排查方法: -
const ErrCode90341 ErrCode = 90341

// ErrCode90342 退款金额超过账户余额
// 排查方法: -
const ErrCode90342 ErrCode = 90342

// ErrCode90343 退款单号已经存在
// 排查方法: -
const ErrCode90343 ErrCode = 90343

// ErrCode90344 指定的付款渠道无效
// 排查方法: -
const ErrCode90344 ErrCode = 90344

// ErrCode90345 超过5w人民币不可指定微信支付渠道
// 排查方法: -
const ErrCode90345 ErrCode = 90345

// ErrCode90346 同一单的退款次数超过限制
// 排查方法: -
const ErrCode90346 ErrCode = 90346

// ErrCode90347 退款金额不可为0
// 排查方法: -
const ErrCode90347 ErrCode = 90347

// ErrCode90348 管理端没配置支付密钥
// 排查方法: -
const ErrCode90348 ErrCode = 90348

// ErrCode90349 记录数量太大
// 排查方法: -
const ErrCode90349 ErrCode = 90349

// ErrCode90350 银行信息真实性校验失败
// 排查方法: -
const ErrCode90350 ErrCode = 90350

// ErrCode90351 应用状态异常
// 排查方法: -
const ErrCode90351 ErrCode = 90351

// ErrCode90352 延迟试用期天数超过限制
// 排查方法: -
const ErrCode90352 ErrCode = 90352

// ErrCode90353 预支单列表不可为空
// 排查方法: -
const ErrCode90353 ErrCode = 90353

// ErrCode90354 预支单列表数量超过限制
// 排查方法: -
const ErrCode90354 ErrCode = 90354

// ErrCode90355 关联有退款预支单，不可删除
// 排查方法: -
const ErrCode90355 ErrCode = 90355

// ErrCode90356 不能0金额下单
// 排查方法: -
const ErrCode90356 ErrCode = 90356

// ErrCode90357 代下单必须指定支付渠道
// 排查方法: -
const ErrCode90357 ErrCode = 90357

// ErrCode90358 预支单或代下单，不支持部分退款
// 排查方法: -
const ErrCode90358 ErrCode = 90358

// ErrCode90359 预支单与下单者企业不匹配
// 排查方法: -
const ErrCode90359 ErrCode = 90359

// ErrCode90381 参数 refunded_credit_orderid 不合法
// 排查方法: -
const ErrCode90381 ErrCode = 90381

// ErrCode90456 必须指定组织者
// 排查方法: -
const ErrCode90456 ErrCode = 90456

// ErrCode90457 日历ID异常
// 排查方法: -
const ErrCode90457 ErrCode = 90457

// ErrCode90458 日历ID列表不能为空
// 排查方法: -
const ErrCode90458 ErrCode = 90458

// ErrCode90459 日历已删除
// 排查方法: -
const ErrCode90459 ErrCode = 90459

// ErrCode90460 日程已删除
// 排查方法: -
const ErrCode90460 ErrCode = 90460

// ErrCode90461 日程ID异常
// 排查方法: -
const ErrCode90461 ErrCode = 90461

// ErrCode90462 日程ID列表不能为空
// 排查方法: -
const ErrCode90462 ErrCode = 90462

// ErrCode90463 不能变更组织者
// 排查方法: -
const ErrCode90463 ErrCode = 90463

// ErrCode90464 参与者数量超过限制
// 排查方法: -
const ErrCode90464 ErrCode = 90464

// ErrCode90465 不支持的重复类型
// 排查方法: -
const ErrCode90465 ErrCode = 90465

// ErrCode90466 不能操作别的应用创建的日历/日程
// 排查方法: -
const ErrCode90466 ErrCode = 90466

// ErrCode90467 星期参数异常
// 排查方法: -
const ErrCode90467 ErrCode = 90467

// ErrCode90468 不能变更组织者
// 排查方法: -
const ErrCode90468 ErrCode = 90468

// ErrCode90469 每页大小超过限制
// 排查方法: -
const ErrCode90469 ErrCode = 90469

// ErrCode90470 页数异常
// 排查方法: -
const ErrCode90470 ErrCode = 90470

// ErrCode90471 提醒时间异常
// 排查方法: -
const ErrCode90471 ErrCode = 90471

// ErrCode90472 没有日历/日程操作权限
// 排查方法: -
const ErrCode90472 ErrCode = 90472

// ErrCode90473 颜色参数异常
// 排查方法: -
const ErrCode90473 ErrCode = 90473

// ErrCode90474 组织者不能与参与者重叠
// 排查方法: -
const ErrCode90474 ErrCode = 90474

// ErrCode90475 不是组织者的日历
// 排查方法: -
const ErrCode90475 ErrCode = 90475

// ErrCode90479 不允许操作用户创建的日程
// 排查方法: -
const ErrCode90479 ErrCode = 90479

// ErrCode90500 群主并未离职
// 排查方法: -
const ErrCode90500 ErrCode = 90500

// ErrCode90501 该群不是客户群
// 排查方法: -
const ErrCode90501 ErrCode = 90501

// ErrCode90502 群主已经离职
// 排查方法: -
const ErrCode90502 ErrCode = 90502

// ErrCode90503 满人 & 99个微信成员，没办法踢，要客户端确认
// 排查方法: -
const ErrCode90503 ErrCode = 90503

// ErrCode90504 群主没变
// 排查方法: -
const ErrCode90504 ErrCode = 90504

// ErrCode90507 离职群正在继承处理中
// 排查方法: -
const ErrCode90507 ErrCode = 90507

// ErrCode90508 离职群已经继承
// 排查方法: -
const ErrCode90508 ErrCode = 90508

// ErrCode91040 获取ticket的类型无效
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：91040)
// 获取ticket的类型无效。jsapi ticket可以通过以下几种获取：
// 1）[获取jsapi_ticket](https://open.work.weixin.qq.com/api/doc#10029/获取jsapi_ticket)。这里参数只需要传access_token，不需要带其余的参数，比如type=jsapi
// 2）[获取电子发票ticket](https://open.work.weixin.qq.com/api/doc#10029/获取电子发票ticket)。需要同时指定access_token及type，同时type=wx_card是固定的。
const ErrCode91040 ErrCode = 91040

// ErrCode92000 成员不在应用可见范围之内
// 排查方法: -
const ErrCode92000 ErrCode = 92000

// ErrCode92001 应用没有敏感信息权限
// 排查方法: -
const ErrCode92001 ErrCode = 92001

// ErrCode92002 不允许跨企业调用
// 排查方法: -
const ErrCode92002 ErrCode = 92002

// ErrCode92006 该直播已经开始或取消
// 排查方法: -
const ErrCode92006 ErrCode = 92006

// ErrCode92007 该直播回放不能被删除
// 排查方法: -
const ErrCode92007 ErrCode = 92007

// ErrCode92008 当前应用没权限操作这个直播
// 排查方法: -
const ErrCode92008 ErrCode = 92008

// ErrCode93000 机器人webhookurl不合法或者机器人已经被移除出群
// 排查方法: -
const ErrCode93000 ErrCode = 93000

// ErrCode93004 机器人被停用
// 排查方法: -
const ErrCode93004 ErrCode = 93004

// ErrCode93008 不在群里
// 排查方法: -
const ErrCode93008 ErrCode = 93008

// ErrCode94000 应用未开启工作台自定义模式
// 排查方法: 请在管理端后台应用详情里面开启自定义工作台模式
const ErrCode94000 ErrCode = 94000

// ErrCode94001 不合法的type类型
// 排查方法: -
const ErrCode94001 ErrCode = 94001

// ErrCode94002 缺少keydata字段
// 排查方法: -
const ErrCode94002 ErrCode = 94002

// ErrCode94003 keydata的items列表长度超出限制
// 排查方法: -
const ErrCode94003 ErrCode = 94003

// ErrCode94005 缺少list字段
// 排查方法: -
const ErrCode94005 ErrCode = 94005

// ErrCode94006 list的items列表长度超出限制
// 排查方法: -
const ErrCode94006 ErrCode = 94006

// ErrCode94007 缺少webview字段
// 排查方法: -
const ErrCode94007 ErrCode = 94007

// ErrCode94008 应用未设置自定义工作台模版类型
// 排查方法: -
const ErrCode94008 ErrCode = 94008

// ErrCode301002 无权限操作指定的应用
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：301002)
// 无权限操作指定的应用。access_token来源需要有指定应用的权限。
// 比如说，[发消息接口](https://open.work.weixin.qq.com/api/doc#10167) 指定了参数 “agentid”: 14，但使用的 access_token 是通过应用agentid: 100032 生成的调用凭证，这种就会报该错误码。
// access_token的权限查询，可在 [错误码查询工具](https://open.work.weixin.qq.com/devtool/query) 确认。
const ErrCode301002 ErrCode = 301002

// ErrCode301005 不允许删除创建者
// 排查方法: 创建者不允许从通讯录中删除。如果需要删除该成员，需要先在WEB管理端转移创建者身份。
const ErrCode301005 ErrCode = 301005

// ErrCode301012 参数 position 不合法
// 排查方法: 长度不允许超过128个字符
const ErrCode301012 ErrCode = 301012

// ErrCode301013 参数 telephone 不合法
// 排查方法: telephone必须由1-32位的纯数字或’-‘号组成。
const ErrCode301013 ErrCode = 301013

// ErrCode301014 参数 english_name 不合法
// 排查方法: 参数如果有传递，不允许为空字符串，同时不能超过64字节，只能是由字母、数字、点(.)、减号(-)、空格或下划线(_)组成
const ErrCode301014 ErrCode = 301014

// ErrCode301015 参数 mediaid 不合法
// 排查方法: 请检查 mediaid 来源，应该通过[上传临时素材](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112)的图片类型获得mediaid
const ErrCode301015 ErrCode = 301015

// ErrCode301016 上传语音文件不符合系统要求
// 排查方法: 语音文件的系统限制，参考[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112/上传的媒体文件限制)
const ErrCode301016 ErrCode = 301016

// ErrCode301017 上传语音文件仅支持AMR格式
// 排查方法: 语音文件的系统限制，参考[上传的媒体文件限制](https://work.weixin.qq.com/api/doc/90000/90139/90313#10112/上传的媒体文件限制)
const ErrCode301017 ErrCode = 301017

// ErrCode301021 参数 userid 无效
// 排查方法: 至少有一个userid不存在于通讯录中
const ErrCode301021 ErrCode = 301021

// ErrCode301022 获取打卡数据失败
// 排查方法: 系统失败，可重试处理
const ErrCode301022 ErrCode = 301022

// ErrCode301023 useridlist非法或超过限额
// 排查方法: 列表数量不能为0且不超过100
const ErrCode301023 ErrCode = 301023

// ErrCode301024 获取打卡记录时间间隔超限
// 排查方法: 保证开始时间大于0 且结束时间大于 0 且结束时间大于开始时间，且间隔少于一个月
const ErrCode301024 ErrCode = 301024

// ErrCode301025 审批开放接口参数错误
// 排查方法: 请参考参数说明正确填写
const ErrCode301025 ErrCode = 301025

// ErrCode301036 不允许更新该用户的userid
// 排查方法: [查看帮助](https://work.weixin.qq.com/api/doc/90000/90139/90313#10649/错误码：301036)
// 不允许更新该用户的userid。确认：
// 只有当userid由系统自动生成时，才被允许修改一次
// 比如，邀请关注时用户提交登记信息，审批通过后系统会自动分配userid，此时可修改userid
const ErrCode301036 ErrCode = 301036

// ErrCode301039 请求参数错误，请检查输入参数
// 排查方法: -
const ErrCode301039 ErrCode = 301039

// ErrCode301042 ip白名单限制，请求ip不在设置白名单范围
// 排查方法: -
const ErrCode301042 ErrCode = 301042

// ErrCode301048 sdkfileid对应的文件不存在或已过期
// 排查方法: -
const ErrCode301048 ErrCode = 301048

// ErrCode301052 会话存档服务已过期
// 排查方法: -
const ErrCode301052 ErrCode = 301052

// ErrCode301053 会话存档服务未开启
// 排查方法: -
const ErrCode301053 ErrCode = 301053

// ErrCode301058 拉取会话数据请求超过大小限制，可减少limit参数
// 排查方法: -
const ErrCode301058 ErrCode = 301058

// ErrCode301059 非内部群，不提供数据
// 排查方法: -
const ErrCode301059 ErrCode = 301059

// ErrCode301060 拉取同意情况请求量过大，请减少到100个参数以下
// 排查方法: -
const ErrCode301060 ErrCode = 301060

// ErrCode301061 userid或者exteropenid用户不存在
// 排查方法: -
const ErrCode301061 ErrCode = 301061

// ErrCode302003 批量导入任务的文件中userid有重复
// 排查方法: -
const ErrCode302003 ErrCode = 302003

// ErrCode302004 组织架构不合法（1不是一棵树，2 多个一样的partyid，3 partyid空，4 partyid name 空，5 同一个父节点下有两个子节点 部门名字一样 可能是以上情况，请一一排查）
// 排查方法: -
const ErrCode302004 ErrCode = 302004

// ErrCode302005 批量导入系统失败，请重新尝试导入
// 排查方法: -
const ErrCode302005 ErrCode = 302005

// ErrCode302006 批量导入任务的文件中partyid有重复
// 排查方法: -
const ErrCode302006 ErrCode = 302006

// ErrCode302007 批量导入任务的文件中，同一个部门下有两个子部门名字一样
// 排查方法: -
const ErrCode302007 ErrCode = 302007

// ErrCode2000002 CorpId参数无效
// 排查方法: 指定的CorpId不存在
const ErrCode2000002 ErrCode = 2000002

// ErrCode600001 不合法的sn
// 排查方法: sn可能尚未进行登记
const ErrCode600001 ErrCode = 600001

// ErrCode600002 设备已注册
// 排查方法: 可能设备已经建立过长连接
const ErrCode600002 ErrCode = 600002

// ErrCode600003 不合法的硬件activecode
// 排查方法: -
const ErrCode600003 ErrCode = 600003

// ErrCode600004 该硬件尚未授权任何企业
// 排查方法: -
const ErrCode600004 ErrCode = 600004

// ErrCode600005 硬件Secret无效
// 排查方法: -
const ErrCode600005 ErrCode = 600005

// ErrCode600007 缺少硬件sn
// 排查方法: -
const ErrCode600007 ErrCode = 600007

// ErrCode600008 缺少nonce参数
// 排查方法: -
const ErrCode600008 ErrCode = 600008

// ErrCode600009 缺少timestamp参数
// 排查方法: -
const ErrCode600009 ErrCode = 600009

// ErrCode600010 缺少signature参数
// 排查方法: -
const ErrCode600010 ErrCode = 600010

// ErrCode600011 签名校验失败
// 排查方法: -
const ErrCode600011 ErrCode = 600011

// ErrCode600012 长连接已经注册过设备
// 排查方法: -
const ErrCode600012 ErrCode = 600012

// ErrCode600013 缺少activecode参数
// 排查方法: -
const ErrCode600013 ErrCode = 600013

// ErrCode600014 设备未网络注册
// 排查方法: -
const ErrCode600014 ErrCode = 600014

// ErrCode600015 缺少secret参数
// 排查方法: -
const ErrCode600015 ErrCode = 600015

// ErrCode600016 设备未激活
// 排查方法: -
const ErrCode600016 ErrCode = 600016

// ErrCode600018 无效的起始结束时间
// 排查方法: -
const ErrCode600018 ErrCode = 600018

// ErrCode600020 设备未登录
// 排查方法: -
const ErrCode600020 ErrCode = 600020

// ErrCode600021 设备sn已存在
// 排查方法: -
const ErrCode600021 ErrCode = 600021

// ErrCode600023 时间戳已失效
// 排查方法: -
const ErrCode600023 ErrCode = 600023

// ErrCode600024 固件大小超过5M
// 排查方法: -
const ErrCode600024 ErrCode = 600024

// ErrCode600025 固件名为空或者超过20字节
// 排查方法: -
const ErrCode600025 ErrCode = 600025

// ErrCode600026 固件信息不存在
// 排查方法: -
const ErrCode600026 ErrCode = 600026

// ErrCode600027 非法的固件参数
// 排查方法: -
const ErrCode600027 ErrCode = 600027

// ErrCode600028 固件版本已存在
// 排查方法: -
const ErrCode600028 ErrCode = 600028

// ErrCode600029 非法的固件版本
// 排查方法: -
const ErrCode600029 ErrCode = 600029

// ErrCode600030 缺少固件版本参数
// 排查方法: -
const ErrCode600030 ErrCode = 600030

// ErrCode600031 硬件固件不允许升级
// 排查方法: -
const ErrCode600031 ErrCode = 600031

// ErrCode600032 无法解析硬件二维码
// 排查方法: -
const ErrCode600032 ErrCode = 600032

// ErrCode600033 设备型号id冲突
// 排查方法: -
const ErrCode600033 ErrCode = 600033

// ErrCode600034 指纹数据大小超过限制
// 排查方法: -
const ErrCode600034 ErrCode = 600034

// ErrCode600035 人脸数据大小超过限制
// 排查方法: -
const ErrCode600035 ErrCode = 600035

// ErrCode600036 设备sn冲突
// 排查方法: -
const ErrCode600036 ErrCode = 600036

// ErrCode600037 缺失设备型号id
// 排查方法: -
const ErrCode600037 ErrCode = 600037

// ErrCode600038 设备型号不存在
// 排查方法: -
const ErrCode600038 ErrCode = 600038

// ErrCode600039 不支持的设备类型
// 排查方法: -
const ErrCode600039 ErrCode = 600039

// ErrCode600040 打印任务id不存在
// 排查方法: -
const ErrCode600040 ErrCode = 600040

// ErrCode600041 无效的offset或limit参数值
// 排查方法: -
const ErrCode600041 ErrCode = 600041

// ErrCode600042 无效的设备型号id
// 排查方法: -
const ErrCode600042 ErrCode = 600042

// ErrCode600043 门禁规则未设置
// 排查方法: -
const ErrCode600043 ErrCode = 600043

// ErrCode600044 门禁规则不合法
// 排查方法: -
const ErrCode600044 ErrCode = 600044

// ErrCode600045 设备已订阅企业信息
// 排查方法: -
const ErrCode600045 ErrCode = 600045

// ErrCode600046 操作id和用户userid不匹配
// 排查方法: -
const ErrCode600046 ErrCode = 600046

// ErrCode600047 secretno的status非法
// 排查方法: 请确认是否是使用统一初始secretno的设备，如果是无有正确执行换secretno的流程
const ErrCode600047 ErrCode = 600047

// ErrCode600048 无效的指纹算法
// 排查方法: -
const ErrCode600048 ErrCode = 600048

// ErrCode600049 无效的人脸识别算法
// 排查方法: -
const ErrCode600049 ErrCode = 600049

// ErrCode600050 无效的算法长度
// 排查方法: -
const ErrCode600050 ErrCode = 600050

// ErrCode600051 设备过期
// 排查方法: -
const ErrCode600051 ErrCode = 600051

// ErrCode600052 无效的文件分块
// 排查方法: -
const ErrCode600052 ErrCode = 600052

// ErrCode600053 该链接已经激活
// 排查方法: -
const ErrCode600053 ErrCode = 600053

// ErrCode600054 该链接已经订阅
// 排查方法: -
const ErrCode600054 ErrCode = 600054

// ErrCode600055 无效的用户类型
// 排查方法: -
const ErrCode600055 ErrCode = 600055

// ErrCode600056 无效的健康状态
// 排查方法: -
const ErrCode600056 ErrCode = 600056

// ErrCode600057 缺少体温参数
// 排查方法: -
const ErrCode600057 ErrCode = 600057

// ErrCode610001 永久二维码超过每个员工5000的限制
// 排查方法: -
const ErrCode610001 ErrCode = 610001

// ErrCode610003 scene参数不合法
// 排查方法: 有效的scene长度为1~64字符，由英文字母、数字、中划线、下划线以及点号构成
const ErrCode610003 ErrCode = 610003

// ErrCode610004 userid不在客户联系配置的使用范围内
// 排查方法: 请在管理端后台 客户联系->配置->配置使用范围配置该用户
const ErrCode610004 ErrCode = 610004

// ErrCode610014 无效的unionid
// 排查方法: -
const ErrCode610014 ErrCode = 610014

// ErrCode610015 小程序对应的开放平台账号未认证
// 排查方法: -
const ErrCode610015 ErrCode = 610015

// ErrCode610016 企业未认证
// 排查方法: -
const ErrCode610016 ErrCode = 610016

// ErrCode610017 小程序和企业主体不一致
// 排查方法: -
const ErrCode610017 ErrCode = 610017

// ErrCode640001 微盘不存在当前空间
// 排查方法: 判断spaceid是否填错
const ErrCode640001 ErrCode = 640001

// ErrCode640002 文件不存在
// 排查方法: 判断fileid是否填错
const ErrCode640002 ErrCode = 640002

// ErrCode640003 文件已删除
// 排查方法: 判断fileid对应的文件是否已经被删除
const ErrCode640003 ErrCode = 640003

// ErrCode640004 无权限访问
// 排查方法: 判断当前用户是否有访问
const ErrCode640004 ErrCode = 640004

// ErrCode640005 成员不在空间内
// 排查方法: 判断当前成员是否在空间内
const ErrCode640005 ErrCode = 640005

// ErrCode640006 超出当前成员拥有的容量
// 排查方法: 判断当前用户的个人容量是否超出了限制
const ErrCode640006 ErrCode = 640006

// ErrCode640007 超出微盘的容量
// 排查方法: 在管理端查看微盘的容量是否要满了
const ErrCode640007 ErrCode = 640007

// ErrCode640008 没有空间权限
// 排查方法: 判断当前userid是否有空间权限
const ErrCode640008 ErrCode = 640008

// ErrCode640009 非法文件名
// 排查方法: 判断file_name字段是否为空
const ErrCode640009 ErrCode = 640009

// ErrCode640010 超出空间的最大成员数
// 排查方法: 判断当前空间的成员数是否超过了管理端设置的空间最大成员数
const ErrCode640010 ErrCode = 640010

// ErrCode640011 json格式不匹配
// 排查方法: 判断是否的json格式是否有误
const ErrCode640011 ErrCode = 640011

// ErrCode640012 非法的userid
// 排查方法: 判断userid字段是否设置错误
const ErrCode640012 ErrCode = 640012

// ErrCode640013 非法的departmentid
// 排查方法: 判断departmentid字段是否设置错误
const ErrCode640013 ErrCode = 640013

// ErrCode640014 空间没有有效的管理员
// 排查方法: 判断当前空间是否没有有效的管理员
const ErrCode640014 ErrCode = 640014

// ErrCode640015 不支持设置预览权限
// 排查方法: 文件预览权限只有VIP用户才能设置
const ErrCode640015 ErrCode = 640015

// ErrCode640016 不支持设置文件水印
// 排查方法: 文件水印只有VIP用户才能设置
const ErrCode640016 ErrCode = 640016

// ErrCode640017 微盘管理端未开通API 权限
// 排查方法: 在微盘管理端进行打开
const ErrCode640017 ErrCode = 640017

// ErrCode640018 微盘管理端未设置编辑权限
// 排查方法: 在微盘管理端进行打开编辑权限
const ErrCode640018 ErrCode = 640018

// ErrCode640019 API 调用次数超出限制
// 排查方法: 免费版：1000次/企业/月; 付费版：100,000次/企业/月
const ErrCode640019 ErrCode = 640019

// ErrCode640020 非法的权限类型
// 排查方法: 普通文件：可下载、仅预览; 微文档：可编辑、仅浏览
const ErrCode640020 ErrCode = 640020

// ErrCode640021 非法的fatherid
// 排查方法: fatherid应该为：文件所在的目录fileid, 在根目录时为fileid（判断当前字段是否为空）
const ErrCode640021 ErrCode = 640021

// ErrCode640022 非法的文件内容的base64
// 排查方法: 文件内容base64，判断本字段是否为空
const ErrCode640022 ErrCode = 640022

// ErrCode640023 非法的权限范围
// 排查方法: auth_scope应该为三个中的其中一个：1:指定人 2:企业内 3:企业外
const ErrCode640023 ErrCode = 640023

// ErrCode640024 非法的fileid
// 排查方法: 判断fileid字段是否为空
const ErrCode640024 ErrCode = 640024

// ErrCode640025 非法的space_name
// 排查方法: 判断space_name字段是否空
const ErrCode640025 ErrCode = 640025

// ErrCode640026 非法的spaceid
// 排查方法: 判断spaceid字段是否空
const ErrCode640026 ErrCode = 640026

// ErrCode640027 参数错误
// 排查方法: 判断输入的参数是否有误
const ErrCode640027 ErrCode = 640027

// ErrCode640028 空间设置了关闭成员邀请链接
// 排查方法: 查看空间的安全设置的成员邀请链接按钮是否处于关闭状态
const ErrCode640028 ErrCode = 640028

// ErrCode640029 只支持下载普通文件，不支持下载文件夹等其他非文件实体类型
// 排查方法: 检查fileid对应的文件是否为普通文件
const ErrCode640029 ErrCode = 640029

// ErrCode844001 非法的output_file_format
// 排查方法: 判断输出文件格式是否正确
const ErrCode844001 ErrCode = 844001
