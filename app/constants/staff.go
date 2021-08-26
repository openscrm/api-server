package constants

// UserGender 用户性别
type UserGender int

const (
	// UserGenderUnspecified 性别未定义
	UserGenderUnspecified UserGender = 0
	// UserGenderMale 男性
	UserGenderMale UserGender = 1
	// UserGenderFemale 女性
	UserGenderFemale UserGender = 2
	// UserGenderUnknown 未知
	UserGenderUnknown UserGender = 3
)

// UserStatus 用户激活信息
//
// 已激活代表已激活企业微信或已关注微工作台（原企业号）。
// 未激活代表既未激活企业微信又未关注微工作台（原企业号）。
type UserStatus int

const (
	// UserStatusActivated 已激活
	UserStatusActivated UserStatus = 1
	// UserStatusDeactivated 已禁用
	UserStatusDeactivated UserStatus = 2
	// UserStatusUnactivated 未激活
	UserStatusUnactivated UserStatus = 4
)

// FollowUserAddWay 该成员添加此客户的来源
//
// 具体含义详见[来源定义](https://work.weixin.qq.com/api/doc/90000/90135/92114#13878/%E6%9D%A5%E6%BA%90%E5%AE%9A%E4%B9%89)
type FollowUserAddWay int

const (
	// FollowUserAddWayUnknown 未知来源
	FollowUserAddWayUnknown FollowUserAddWay = 0
	// FollowUserAddWayQRCode 扫描二维码
	FollowUserAddWayQRCode FollowUserAddWay = 1
	// FollowUserAddWayMobile 搜索手机号
	FollowUserAddWayMobile FollowUserAddWay = 2
	// FollowUserAddWayCard 名片分享
	FollowUserAddWayCard FollowUserAddWay = 3
	// FollowUserAddWayGroupChat 群聊
	FollowUserAddWayGroupChat FollowUserAddWay = 4
	// FollowUserAddWayAddressBook 手机通讯录
	FollowUserAddWayAddressBook FollowUserAddWay = 5
	// FollowUserAddWayWeChatContact 微信联系人
	FollowUserAddWayWeChatContact FollowUserAddWay = 6
	// FollowUserAddWayWeChatFriendApply 来自微信的添加好友申请
	FollowUserAddWayWeChatFriendApply FollowUserAddWay = 7
	// FollowUserAddWayThirdParty 安装第三方应用时自动添加的客服人员
	FollowUserAddWayThirdParty FollowUserAddWay = 8
	// FollowUserAddWayEmail 搜索邮箱
	FollowUserAddWayEmail FollowUserAddWay = 9
	// FollowUserAddWayInternalShare 内部成员共享
	FollowUserAddWayInternalShare FollowUserAddWay = 201
	// FollowUserAddWayAdmin 管理员/负责人分配
	FollowUserAddWayAdmin FollowUserAddWay = 202
)
