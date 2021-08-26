package workwx

// ExternalContact 外部联系人
type ExternalContact struct {
	// ExternalUserid 外部联系人的userid
	ExternalUserid string `json:"external_userid"`
	// Name 外部联系人的名称，如果外部联系人为微信用户，则返回外部联系人的名称为其微信昵称；如果外部联系人为企业微信用户，则会按照以下优先级顺序返回：此外部联系人或管理员设置的昵称、认证的实名和账号名称。
	Name string `json:"name"`
	// Position 外部联系人的职位，如果外部企业或用户选择隐藏职位，则不返回，仅当联系人类型是企业微信用户时有此字段
	Position string `json:"position"`
	// Avatar 外部联系人头像，第三方不可获取
	Avatar string `json:"avatar"`
	// CorpName 外部联系人所在企业的简称，仅当联系人类型是企业微信用户时有此字段
	CorpName string `json:"corp_name"`
	// Type 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
	Type ExternalUserType `json:"type"`
	// Gender 外部联系人性别 0-未知 1-男性 2-女性
	Gender UserGender `json:"gender"`
	// Unionid 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当联系人类型是微信用户，且企业或第三方服务商绑定了微信开发者ID有此字段。查看绑定方法 关于返回的unionid，如果是第三方应用调用该接口，则返回的unionid是该第三方服务商所关联的微信开放者帐号下的unionid。也就是说，同一个企业客户，企业自己调用，与第三方服务商调用，所返回的unionid不同；不同的服务商调用，所返回的unionid也不同。
	Unionid string `json:"unionid"`
	// ExternalProfile 成员对外信息
	ExternalProfile ExternalProfile `json:"external_profile"`
}

// ExternalProfile 成员对外信息
type ExternalProfile struct {
	// ExternalCorpName 企业简称
	ExternalCorpName string `json:"external_corp_name"`
	// ExternalAttr 属性列表，目前支持文本、网页、小程序三种类型
	ExternalAttr []ExternalAttr `json:"external_attr"`
}

// ExternalAttr 属性列表，目前支持文本、网页、小程序三种类型
type ExternalAttr struct {
	// Type 属性类型: 0-文本 1-网页 2-小程序
	Type int `json:"type"`
	// Name 属性名称： 需要先确保在管理端有创建该属性，否则会忽略
	Name string `json:"name"`
	// Text 文本类型的属性 ，type为0时必填
	Text ExternalAttrText `json:"text"`
	// Web 网页类型的属性，url和title字段要么同时为空表示清除该属性，要么同时不为空 ，type为1时必填
	Web ExternalAttrWeb `json:"web"`
	// Miniprogram 小程序类型的属性，appid和title字段要么同时为空表示清除改属性，要么同时不为空 ，type为2时必填
	Miniprogram ExternalAttrMiniprogram `json:"miniprogram"`
}

// ExternalAttrText 文本类型的属性
type ExternalAttrText struct {
	// Value 文本属性内容,长度限制12个UTF8字符
	Value string `json:"value"`
}

// ExternalAttrWeb 网页类型的属性，url和title字段要么同时为空表示清除该属性，要么同时不为空 ，type为1时必填
type ExternalAttrWeb struct {
	// Url 网页的url,必须包含http或者https头
	Url string `json:"url"`
	// Title 网页的展示标题,长度限制12个UTF8字符
	Title string `json:"title"`
}

// ExternalAttrMiniprogram 小程序类型的属性，appid和title字段要么同时为空表示清除改属性，要么同时不为空 ，type为2时必填
type ExternalAttrMiniprogram struct {
	// Appid 小程序appid，必须是有在本企业安装授权的小程序，否则会被忽略
	Appid string `json:"appid"`
	// Pagepath 小程序的页面路径
	Pagepath string `json:"pagepath"`
	// Title 企业对外简称，需从已认证的企业简称中选填。可在“我的企业”页中查看企业简称认证状态。
	Title string `json:"title"`
}

// ExternalUserType 外部联系人的类型
//
// 1表示该外部联系人是微信用户
// 2表示该外部联系人是企业微信用户
type ExternalUserType int

const (
	// ExternalUserTypeWeChat 微信用户
	ExternalUserTypeWeChat ExternalUserType = 1
	// ExternalUserTypeWorkWeChat 企业微信用户
	ExternalUserTypeWorkWeChat ExternalUserType = 2
)

// FollowUser 添加了外部联系人的企业成员
type FollowUser struct {
	//  添加了外部联系人的企业成员
	FollowUserInfo
	// Tags 该成员添加此外部联系人所打标签
	Tags []FollowUserTag `json:"tags"`
}

// FollowInfo 企业成员客户跟进信息，可以参考获取客户详情，但标签信息只会返回企业标签的tag_id，个人标签将不再返回
type FollowInfo struct {
	//  添加了外部联系人的企业成员
	FollowUserInfo
	// TagID 该成员添加此外部联系人所打标签
	TagID []string `json:"tag_id"`
}

// FollowUserInfo 添加了外部联系人的企业成员
type FollowUserInfo struct {
	// UserID 	添加了此外部联系人的企业成员userid
	UserID string `json:"userid"`
	// Remark 该成员对此外部联系人的备注
	Remark string `json:"remark"`
	// Description 该成员对此外部联系人的描述
	Description string `json:"description"`
	// Createtime 该成员添加此外部联系人的时间
	Createtime int `json:"createtime"`
	// RemarkCorpName 该成员对此客户备注的企业名称
	RemarkCorpName string `json:"remark_corp_name"`
	// RemarkMobiles 该成员对此客户备注的手机号码，第三方不可获取
	RemarkMobiles []string `json:"remark_mobiles"`
	// AddWay 该成员添加此客户的来源
	AddWay FollowUserAddWay `json:"add_way"`
	// OperUserID 发起添加的userid，如果成员主动添加，为成员的userid；如果是客户主动添加，则为客户的外部联系人userid；如果是内部成员共享/管理员分配，则为对应的成员/管理员userid
	OperUserID string `json:"oper_userid"`
	// State 企业自定义的state参数，用于区分客户具体是通过哪个「联系我」添加，由企业通过创建「联系我」方式指定
	State string `json:"state"`
}

// FollowUserTag 该成员添加此外部联系人所打标签
type FollowUserTag struct {
	// GroupName 该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
	GroupName string `json:"group_name"`
	// TagName 该成员添加此外部联系人所打标签名称
	TagName string `json:"tag_name"`
	// Type 该成员添加此外部联系人所打标签类型, 1-企业设置, 2-用户自定义
	Type FollowUserTagType `json:"type"`
	// TagID 标签id
	TagID string `json:"tag_id"`
}

// FollowUserTagType 该成员添加此外部联系人所打标签类型
//
// 1-企业设置
// 2-用户自定义
type FollowUserTagType int

const (
	// 企业设置
	FollowUserTagTypeWork FollowUserTagType = 1
	// 用户自定义
	FollowUserTagTypeUser FollowUserTagType = 2
)

// FollowUserAddWay 该成员添加此客户的来源
//
// 具体含义详见[来源定义](https://work.weixin.qq.com/api/doc/90000/90135/92114#13878/%E6%9D%A5%E6%BA%90%E5%AE%9A%E4%B9%89)
type FollowUserAddWay int

const (
	// 未知来源
	FollowUserAddWayUnknown FollowUserAddWay = 0
	// 扫描二维码
	FollowUserAddWayQRCode FollowUserAddWay = 1
	// 搜索手机号
	FollowUserAddWayMobile FollowUserAddWay = 2
	// 名片分享
	FollowUserAddWayCard FollowUserAddWay = 3
	// 群聊
	FollowUserAddWayGroupChat FollowUserAddWay = 4
	// 手机通讯录
	FollowUserAddWayAddressBook FollowUserAddWay = 5
	// 微信联系人
	FollowUserAddWayWeChatContact FollowUserAddWay = 6
	// 来自微信的添加好友申请
	FollowUserAddWayWeChatFriendApply FollowUserAddWay = 7
	// 安装第三方应用时自动添加的客服人员
	FollowUserAddWayThirdParty FollowUserAddWay = 8
	// 搜索邮箱
	FollowUserAddWayEmail FollowUserAddWay = 9
	// 内部成员共享
	FollowUserAddWayInternalShare FollowUserAddWay = 201
	// 管理员/负责人分配
	FollowUserAddWayAdmin FollowUserAddWay = 202
)

// ExternalContactRemark 客户备注信息
type ExternalContactRemark struct {
	// Userid 企业成员的userid
	Userid string `json:"userid"`
	// ExternalUserid 外部联系人userid
	ExternalUserid string `json:"external_userid"`
	// Remark 此用户对外部联系人的备注，最多20个字符，remark，description，remark_company，remark_mobiles和remark_pic_mediaid不可同时为空。
	Remark string `json:"remark"`
	// Description 此用户对外部联系人的描述，最多150个字符
	Description string `json:"description"`
	// RemarkCompany 此用户对外部联系人备注的所属公司名称，最多20个字符，remark_company只在此外部联系人为微信用户时有效。
	RemarkCompany string `json:"remark_company"`
	// RemarkMobiles 此用户对外部联系人备注的手机号，如果填写了remark_mobiles，将会覆盖旧的备注手机号。如果要清除所有备注手机号,请在remark_mobiles填写一个空字符串(“”)。
	RemarkMobiles []string `json:"remark_mobiles"`
	// RemarkPicMediaid 备注图片的mediaid，remark_pic_mediaid可以通过素材管理接口获得。
	RemarkPicMediaid string `json:"remark_pic_mediaid"`
}

// ExternalContactCorpTag 企业客户标签
type ExternalContactCorpTag struct {
	// ID 标签id
	ID string `json:"id"`
	// Name 标签名称
	Name string `json:"name"`
	// CreateTime 标签创建时间
	CreateTime int `json:"create_time"`
	// Order 标签排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Order uint32 `json:"order"`
	// Deleted 标签是否已经被删除，只在指定tag_id进行查询时返回
	Deleted bool `json:"deleted"`
}

// ExternalContactCorpTagGroup 企业客户标签
type ExternalContactCorpTagGroup struct {
	// GroupID 标签组id
	GroupID string `json:"group_id"`
	// GroupName 标签组名称
	GroupName string `json:"group_name"`
	// CreateTime 标签组创建时间
	CreateTime int `json:"create_time"`
	// Order 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Order uint32 `json:"order"`
	// Deleted 标签组是否已经被删除，只在指定tag_id进行查询时返回
	Deleted bool `json:"deleted"`
	// Tag 标签组内的标签列表
	Tag []ExternalContactCorpTag `json:"tag"`
}

// ExternalContactMarkTag 企业标记客户标签
type ExternalContactMarkTag struct {
	// UserID 添加外部联系人的userid
	UserID string `json:"userid"`
	// ExternalUserID 外部联系人userid
	ExternalUserID string `json:"external_userid"`
	// AddTag 要标记的标签列表
	AddTag []string `json:"add_tag"`
	// RemoveTag 要移除的标签列表
	RemoveTag []string `json:"remove_tag"`
}

// ExternalContactUnassignedList 离职成员的客户列表
type ExternalContactUnassignedList struct {
	// Info 离职成员的客户
	Info []ExternalContactUnassigned `json:"info"`
	// IsLast 是否是最后一条记录
	IsLast bool `json:"is_last"`
	// NextCursor 分页查询游标,已经查完则返回空("")
	NextCursor string `json:"next_cursor"`
}

// ExternalContactTransferStatus 客户接替结果状态
type ExternalContactTransferStatus uint8

const (
	// ExternalContactTransferStatusSuccess 1-接替完毕
	ExternalContactTransferStatusSuccess ExternalContactTransferStatus = 1
	// ExternalContactTransferStatusWait 2-等待接替
	ExternalContactTransferStatusWait ExternalContactTransferStatus = 2
	// ExternalContactTransferStatusRefused 3-客户拒绝
	ExternalContactTransferStatusRefused ExternalContactTransferStatus = 3
	// ExternalContactTransferStatusExhausted 4-接替成员客户达到上限
	ExternalContactTransferStatusExhausted ExternalContactTransferStatus = 4
	// ExternalContactTransferStatusNoData 5-无接替记录
	ExternalContactTransferStatusNoData ExternalContactTransferStatus = 5
)

// ExternalContactGroupChatTransferFailed 离职成员的群再分配失败
type ExternalContactGroupChatTransferFailed struct {
	// ChatID 没能成功继承的群ID
	ChatID string `json:"chat_id"`
	// ErrCode 没能成功继承的群，错误码
	ErrCode int `json:"errcode"`
	// ErrMsg 没能成功继承的群，错误描述
	ErrMsg string `json:"errmsg"`
}
