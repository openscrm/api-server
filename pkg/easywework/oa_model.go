package workwx

// OAApplyEvent 提交审批申请
type OAApplyEvent struct {
	// CreatorUserID 申请人userid，此审批申请将以此员工身份提交，申请人需在应用可见范围内
	CreatorUserID string `json:"creator_userid"`
	// TemplateID 模板id。可在“获取审批申请详情”、“审批状态变化回调通知”中获得，也可在审批模板的模板编辑页面链接中获得。暂不支持通过接口提交[打卡补卡][调班]模板审批单。
	TemplateID string `json:"template_id"`
	// UseTemplateApprover 审批人模式：0-通过接口指定审批人、抄送人（此时approver、notifyer等参数可用）; 1-使用此模板在管理后台设置的审批流程，支持条件审批。默认为0
	UseTemplateApprover uint8 `json:"use_template_approver"`
	// Approver 审批流程信息，用于指定审批申请的审批流程，支持单人审批、多人会签、多人或签，可能有多个审批节点，仅use_template_approver为0时生效。
	Approver []OAApprover `json:"approver"`
	// Notifier 抄送人节点userid列表，仅use_template_approver为0时生效。
	Notifier []string `json:"notifyer"`
	// NotifyType 抄送方式：1-提单时抄送（默认值）； 2-单据通过后抄送；3-提单和单据通过后抄送。仅use_template_approver为0时生效。
	NotifyType *uint8 `json:"notify_type"`
	// ApplyData 审批申请数据，可定义审批申请中各个控件的值，其中必填项必须有值，选填项可为空，数据结构同“获取审批申请详情”接口返回值中同名参数“apply_data”
	ApplyData OAContents `json:"apply_data"`
	// SummaryList 摘要信息，用于显示在审批通知卡片、审批列表的摘要信息，最多3行
	SummaryList []OASummaryList `json:"summary_list"`
}

// OAApprover 审批流程信息
type OAApprover struct {
	// Attr 节点审批方式：1-或签；2-会签，仅在节点为多人审批时有效
	Attr uint8 `json:"attr"`
	// UserID 审批节点审批人userid列表，若为多人会签、多人或签，需填写每个人的userid
	UserID []string `json:"userid"`
}

// OAContent 审批申请详情，由多个表单控件及其内容组成，其中包含需要对控件赋值的信息
type OAContent struct {
	// Control 控件类型：Text-文本；Textarea-多行文本；Number-数字；Money-金额；Date-日期/日期+时间；Selector-单选/多选；；Contact-成员/部门；Tips-说明文字；File-附件；Table-明细；
	Control OAControl `json:"control"`
	// ID 控件id：控件的唯一id，可通过“获取审批模板详情”接口获取
	ID string `json:"id"`
	// Title 控件名称 ，若配置了多语言则会包含中英文的控件名称
	Title []OAText `json:"title"`
	// Value 控件值 ，需在此为申请人在各个控件中填写内容不同控件有不同的赋值参数，具体说明详见附录。模板配置的控件属性为必填时，对应value值需要有值。
	Value OAContentValue `json:"value"`
}

// OAContents 审批申请详情，由多个表单控件及其内容组成，其中包含需要对控件赋值的信息
type OAContents struct {
	// Contents 审批申请详情，由多个表单控件及其内容组成，其中包含需要对控件赋值的信息
	Contents []OAContent `json:"contents"`
}

// OAText 通用文本信息
type OAText struct {
	// Text 文字
	Text string `json:"text"`
	// Lang 语言
	Lang string `json:"lang"`
}

// OASummaryList 摘要行信息，用于定义某一行摘要显示的内容
type OASummaryList struct {
	// SummaryInfo 摘要行信息，用于定义某一行摘要显示的内容
	SummaryInfo []OAText `json:"summary_info"`
}

// OAContentValue 控件值 ，需在此为申请人在各个控件中填写内容不同控件有不同的赋值参数，具体说明详见附录。模板配置的控件属性为必填时，对应value值需要有值。
type OAContentValue struct {
	// Text 文本/多行文本控件（control参数为Text或Textarea）
	Text string `json:"text"`
	// Number 数字控件（control参数为Number）
	Number string `json:"new_number"`
	// Money 金额控件（control参数为Money）
	Money string `json:"new_money"`
	// Date 日期/日期+时间控件（control参数为Date）
	Date OAContentDate `json:"date"`
	// Selector 单选/多选控件（control参数为Selector）
	Selector OAContentSelector `json:"selector"`
	// Members 成员控件（control参数为Contact，且value参数为members）
	Members []OAContentMember `json:"members"`
	// Departments 部门控件（control参数为Contact，且value参数为departments）
	Departments []OAContentDepartment `json:"departments"`
	// Files 附件控件（control参数为File，且value参数为files）
	Files []OAContentFile `json:"files"`
	// Table 明细控件（control参数为Table）
	Table []OAContentTableList `json:"children"`
	// Vacation 假勤组件-请假组件（control参数为Vacation）
	Vacation OAContentVacation `json:"vacation"`
	// Location 位置控件（control参数为Location，且value参数为location）
	Location OAContentLocation `json:"location"`
	// RelatedApproval 关联审批单控件（control参数为RelatedApproval，且value参数为related_approval）
	RelatedApproval []OAContentRelatedApproval `json:"related_approval"`
	// Formula 公式控件（control参数为Formula，且value参数为formula）
	Formula OAContentFormula `json:"formula"`
	// DateRange 时长组件（control参数为DateRange，且value参数为date_range）
	DateRange OAContentDateRange `json:"date_range"`
}

// OAContentDate 日期/日期+时间内容
type OAContentDate struct {
	// Type 时间展示类型：day-日期；hour-日期+时间 ，和对应模板控件属性一致
	Type string `json:"type"`
	// Timestamp 时间戳-字符串类型，在此填写日期/日期+时间控件的选择值，以此为准
	Timestamp string `json:"s_timestamp"`
}

// OAContentSelector 类型标志，单选/多选控件的config中会包含此参数
type OAContentSelector struct {
	// Type 选择方式：single-单选；multi-多选
	Type string `json:"type"`
	// Options 多选选项，多选属性的选择控件允许输入多个
	Options []OAContentSelectorOption `json:"options"`
}

// OAContentSelectorOption 多选选项，多选属性的选择控件允许输入多个
type OAContentSelectorOption struct {
	// Key 选项key，可通过“获取审批模板详情”接口获得
	Key string `json:"key"`
}

// OAContentMember 所选成员内容，即申请人在此控件选择的成员，多选模式下可以有多个
type OAContentMember struct {
	// UserID 所选成员的userid
	UserID string `json:"userid"`
	// Name 成员名
	Name string `json:"name"`
}

// OAContentDepartment 所选部门内容，即申请人在此控件选择的部门，多选模式下可能有多个
type OAContentDepartment struct {
	// OpenAPIID 所选部门id
	OpenAPIID string `json:"openapi_id"`
	// Name 所选部门名
	Name string `json:"name"`
}

// OAContentFile 附件
type OAContentFile struct {
	// FileID 文件id，该id为临时素材上传接口返回的的media_id，注：提单后将作为单据内容转换为长期文件存储；目前一个审批申请单，全局仅支持上传6个附件，否则将失败。
	FileID string `json:"file_id"`
}

// OAContentTableList 子明细列表，在此填写子明细的所有子控件的值，子控件的数据结构同一般控件
type OAContentTableList struct {
	// List 子明细列表，在此填写子明细的所有子控件的值，子控件的数据结构同一般控件
	List []OAContent `json:"list"`
}

// OAContentVacation 请假内容，即申请人在此组件内选择的请假信息
type OAContentVacation struct {
	// Selector 请假类型，所选选项与假期管理关联，为假期管理中的假期类型
	Selector OAContentSelector `json:"selector"`
	// Attendance 假勤组件
	Attendance OAContentVacationAttendance `json:"attendance"`
}

// OAContentVacationAttendance 假勤组件
type OAContentVacationAttendance struct {
	// DateRange 假勤组件时间选择范围
	DateRange OAContentVacationAttendanceDateRange `json:"date_range"`
	// Type 假勤组件类型：1-请假；3-出差；4-外出；5-加班
	Type uint8 `json:"type"`
}

// OAContentVacationAttendanceDateRange 假勤组件时间选择范围
type OAContentVacationAttendanceDateRange struct {
	// Type 时间展示类型：day-日期；hour-日期+时间
	Type string `json:"type"`
	//  时长范围
	OAContentDateRange
}

// OAContentLocation 位置控件
type OAContentLocation struct {
	// Latitude 纬度，精确到6位小数
	Latitude string `json:"latitude"`
	// Longitude 经度，精确到6位小数
	Longitude string `json:"longitude"`
	// Title 地点标题
	Title string `json:"title"`
	// Address 地点详情地址
	Address string `json:"address"`
	// Time 选择地点的时间
	Time int `json:"time"`
}

// OAContentRelatedApproval 关联审批单控件
type OAContentRelatedApproval struct {
	// SpNo 关联审批单的审批单号
	SpNo string `json:"sp_no"`
}

// OAContentFormula 公式控件
type OAContentFormula struct {
	// Value 公式的值，提交表单时无需填写，后台自动计算
	Value string `json:"value"`
}

// OAContentDateRange 时长组件
type OAContentDateRange struct {
	// NewBegin 开始时间，unix时间戳
	NewBegin int `json:"new_begin"`
	// NewEnd 结束时间，unix时间戳
	NewEnd int `json:"new_end"`
	// NewDuration 时长范围，单位秒
	NewDuration int `json:"new_duration"`
}

// OATemplateDetail 审批模板详情
type OATemplateDetail struct {
	// TemplateNames 模板名称，若配置了多语言则会包含中英文的模板名称，默认为zh_CN中文
	TemplateNames []OAText `json:"template_names"`
	// TemplateContent 模板控件信息
	TemplateContent OATemplateControls `json:"template_content"`
	// Vacation Vacation控件（假勤控件）
	Vacation OATemplateControlConfigVacation `json:"vacation_list"`
}

// OATemplateControls 模板控件数组。模板详情由多个不同类型的控件组成，控件类型详细说明见附录。
type OATemplateControls struct {
	// Controls 模板名称，若配置了多语言则会包含中英文的模板名称，默认为zh_CN中文
	Controls []OATemplateControl `json:"controls"`
}

// OATemplateControl 模板控件信息
type OATemplateControl struct {
	// Property 模板控件属性，包含了模板内控件的各种属性信息
	Property OATemplateControlProperty `json:"property"`
	// Config 模板控件配置，包含了部分控件类型的附加类型、属性，详见附录说明。目前有配置信息的控件类型有：Date-日期/日期+时间；Selector-单选/多选；Contact-成员/部门；Table-明细；Attendance-假勤组件（请假、外出、出差、加班）
	Config OATemplateControlConfig `json:"config"`
}

// OATemplateControlProperty 模板控件属性
type OATemplateControlProperty struct {
	// Control 模板控件属性，包含了模板内控件的各种属性信息
	Control OAControl `json:"control"`
	// ID 模板控件配置，包含了部分控件类型的附加类型、属性，详见附录说明。目前有配置信息的控件类型有：Date-日期/日期+时间；Selector-单选/多选；Contact-成员/部门；Table-明细；Attendance-假勤组件（请假、外出、出差、加班）
	ID string `json:"id"`
	// Title 模板控件配置，包含了部分控件类型的附加类型、属性，详见附录说明。目前有配置信息的控件类型有：Date-日期/日期+时间；Selector-单选/多选；Contact-成员/部门；Table-明细；Attendance-假勤组件（请假、外出、出差、加班）
	Title []OAText `json:"title"`
	// Placeholder 模板控件配置，包含了部分控件类型的附加类型、属性，详见附录说明。目前有配置信息的控件类型有：Date-日期/日期+时间；Selector-单选/多选；Contact-成员/部门；Table-明细；Attendance-假勤组件（请假、外出、出差、加班）
	Placeholder []OAText `json:"placeholder"`
	// Require 是否必填：1-必填；0-非必填
	Require uint8 `json:"require"`
	// UnPrint 是否参与打印：1-不参与打印；0-参与打印
	UnPrint uint8 `json:"un_print"`
}

// OATemplateControlConfig 模板控件配置
type OATemplateControlConfig struct {
	// Date Date控件（日期/日期+时间控件）
	Date OATemplateControlConfigDate `json:"date"`
	// Selector Selector控件（单选/多选控件）
	Selector OATemplateControlConfigSelector `json:"selector"`
	// Contact Contact控件（成员/部门控件）
	Contact OATemplateControlConfigContact `json:"contact"`
	// Table Table（明细控件）
	Table OATemplateControlConfigTable `json:"table"`
	// Attendance Attendance控件（假勤控件）
	Attendance OATemplateControlConfigAttendance `json:"attendance"`
}

// OATemplateControlConfigDate 类型标志，日期/日期+时间控件的config中会包含此参数
type OATemplateControlConfigDate struct {
	// Type 时间展示类型：day-日期；hour-日期+时间
	Type string `json:"type"`
}

// OATemplateControlConfigSelector 类型标志，单选/多选控件的config中会包含此参数
type OATemplateControlConfigSelector struct {
	// Type 选择类型：single-单选；multi-多选
	Type string `json:"type"`
	// Options 选项，包含单选/多选控件中的所有选项，可能有多个
	Options []OATemplateControlConfigSelectorOption `json:"options"`
}

// OATemplateControlConfigSelectorOption 选项，包含单选/多选控件中的所有选项，可能有多个
type OATemplateControlConfigSelectorOption struct {
	// Key 选项key，选项的唯一id，可用于发起审批申请，为单选/多选控件赋值
	Key string `json:"key"`
	// Value 选项值，若配置了多语言则会包含中英文的选项值，默认为zh_CN中文
	Value []OAText `json:"value"`
}

// OATemplateControlConfigContact 类型标志，单选/多选控件的config中会包含此参数
type OATemplateControlConfigContact struct {
	// Type 选择类型：single-单选；multi-多选
	Type string `json:"type"`
	// Mode 选择对象：user-成员；department-部门
	Mode string `json:"mode"`
}

// OATemplateControlConfigTable 类型标志，明细控件的config中会包含此参数
type OATemplateControlConfigTable struct {
	// Children 明细内的子控件，内部结构同controls
	Children []OATemplateControl `json:"children"`
}

// OATemplateControlConfigAttendance 类型标志，假勤控件的config中会包含此参数
type OATemplateControlConfigAttendance struct {
	// DateRange 假期控件属性
	DateRange OATemplateControlConfigAttendanceDateRange `json:"date_range"`
	// Type 假勤控件类型：1-请假，3-出差，4-外出，5-加班
	Type uint8 `json:"type"`
}

// OATemplateControlConfigAttendanceDateRange 假期控件属性
type OATemplateControlConfigAttendanceDateRange struct {
	// Type 时间刻度：hour-精确到分钟, halfday—上午/下午
	Type string `json:"type"`
}

// OATemplateControlConfigVacation 类型标志，假勤控件的config中会包含此参数
type OATemplateControlConfigVacation struct {
	// Item 单个假期类型属性
	Item []OATemplateControlConfigVacationItem `json:"item"`
}

// OATemplateControlConfigVacationItem 类型标志，假勤控件的config中会包含此参数
type OATemplateControlConfigVacationItem struct {
	// ID 假期类型标识id
	ID int `json:"id"`
	// Name 假期类型名称，默认zh_CN中文名称
	Name []OAText `json:"name"`
}

// OAControl 控件类型
type OAControl string

// OAControlText 文本
const OAControlText OAControl = "Text"

// OAControlTextarea 多行文本
const OAControlTextarea OAControl = "Textarea"

// OAControlNumber 数字
const OAControlNumber OAControl = "Number"

// OAControlMoney 金额
const OAControlMoney OAControl = "Money"

// OAControlDate 日期/日期+时间控件
const OAControlDate OAControl = "Date"

// OAControlSelector 单选/多选控件
const OAControlSelector OAControl = "Selector"

// OAControlContact 成员/部门控件
const OAControlContact OAControl = "Contact"

// OAControlTips 说明文字控件
const OAControlTips OAControl = "Tips"

// OAControlFile 附件控件
const OAControlFile OAControl = "File"

// OAControlTable 明细控件
const OAControlTable OAControl = "Table"

// OAControlLocation 位置控件
const OAControlLocation OAControl = "Location"

// OAControlRelatedApproval 关联审批单控件
const OAControlRelatedApproval OAControl = "RelatedApproval"

// OAControlFormula 公式控件
const OAControlFormula OAControl = "Formula"

// OAControlDateRange 时长控件
const OAControlDateRange OAControl = "DateRange"

// OAControlVacation 假勤组件-请假组件
const OAControlVacation OAControl = "Vacation"

// OAControlAttendance 假勤组件-出差/外出/加班组件
const OAControlAttendance OAControl = "Attendance"

// OAApprovalDetail 审批申请详情
type OAApprovalDetail struct {
	// SpNo 审批编号
	SpNo string `json:"sp_no"`
	// SpName 审批申请类型名称（审批模板名称）
	SpName string `json:"sp_name"`
	// SpStatus 申请单状态：1-审批中；2-已通过；3-已驳回；4-已撤销；6-通过后撤销；7-已删除；10-已支付
	SpStatus uint8 `json:"sp_status"`
	// TemplateID 审批模板id。可在“获取审批申请详情”、“审批状态变化回调通知”中获得，也可在审批模板的模板编辑页面链接中获得。
	TemplateID string `json:"template_id"`
	// ApplyTime 审批申请提交时间,Unix时间戳
	ApplyTime int `json:"apply_time"`
	// Applicant 申请人信息
	Applicant OAApprovalDetailApplicant `json:"applyer"`
	// SpRecord 审批流程信息，可能有多个审批节点。
	SpRecord []OAApprovalDetailSpRecord `json:"sp_record"`
	// Notifier 抄送信息，可能有多个抄送节点
	Notifier []OAApprovalDetailNotifier `json:"notifyer"`
	// ApplyData 审批申请数据
	ApplyData OAContents `json:"apply_data"`
	// Comments 审批申请备注信息，可能有多个备注节点
	Comments []OAApprovalDetailComment `json:"comments"`
}

// OAApprovalDetailApplicant 审批申请详情申请人信息
type OAApprovalDetailApplicant struct {
	// UserID 申请人userid
	UserID string `json:"userid"`
	// PartyID 申请人所在部门id
	PartyID string `json:"partyid"`
}

// OAApprovalDetailSpRecord 审批流程信息，可能有多个审批节点。
type OAApprovalDetailSpRecord struct {
	// SpStatus 审批节点状态：1-审批中；2-已同意；3-已驳回；4-已转审
	SpStatus uint8 `json:"sp_status"`
	// ApproverAttr 节点审批方式：1-或签；2-会签
	ApproverAttr uint8 `json:"approverattr"`
	// Details 审批节点详情,一个审批节点有多个审批人
	Details []OAApprovalDetailSpRecordDetail `json:"details"`
}

// OAApprovalDetailSpRecordDetail 审批节点详情,一个审批节点有多个审批人
type OAApprovalDetailSpRecordDetail struct {
	// Approver 分支审批人
	Approver OAApprovalDetailSpRecordDetailApprover `json:"approver"`
	// Speech 审批意见
	Speech string `json:"speech"`
	// SpStatus 分支审批人审批状态：1-审批中；2-已同意；3-已驳回；4-已转审
	SpStatus uint8 `json:"sp_status"`
	// SpTime 节点分支审批人审批操作时间戳，0表示未操作
	SpTime int `json:"sptime"`
	// MediaID 节点分支审批人审批意见附件，media_id具体使用请参考：文档-获取临时素材
	MediaID []string `json:"media_id"`
}

// OAApprovalDetailSpRecordDetailApprover 分支审批人
type OAApprovalDetailSpRecordDetailApprover struct {
	// UserID 分支审批人userid
	UserID string `json:"userid"`
}

// OAApprovalDetailNotifier 抄送信息，可能有多个抄送节点
type OAApprovalDetailNotifier struct {
	// UserID 节点抄送人userid
	UserID string `json:"userid"`
}

// OAApprovalDetailComment 审批申请备注信息，可能有多个备注节点
type OAApprovalDetailComment struct {
	// CommentUserInfo 备注人信息
	CommentUserInfo OAApprovalDetailCommentUserInfo `json:"commentUserInfo"`
	// CommentTime 备注提交时间戳，Unix时间戳
	CommentTime int `json:"commenttime"`
	// CommentTontent 备注文本内容
	CommentTontent string `json:"commentcontent"`
	// CommentID 备注id
	CommentID string `json:"commentid"`
	// MediaID 备注附件id，可能有多个，media_id具体使用请参考：文档-获取临时素材
	MediaID []string `json:"media_id"`
}

// OAApprovalDetailCommentUserInfo 备注人信息
type OAApprovalDetailCommentUserInfo struct {
	// UserID 备注人userid
	UserID string `json:"userid"`
}

// OAApprovalInfoFilter 备注人信息
type OAApprovalInfoFilter struct {
	// Key 筛选类型，包括：template_id - 模板类型/模板id；creator - 申请人；department - 审批单提单者所在部门；sp_status - 审批状态。注意:仅“部门”支持同时配置多个筛选条件。不同类型的筛选条件之间为“与”的关系，同类型筛选条件之间为“或”的关系
	Key OAApprovalInfoFilterKey `json:"key"`
	// Value 筛选值，对应为：template_id - 模板id；creator - 申请人userid；department - 所在部门id；sp_status - 审批单状态（1-审批中；2-已通过；3-已驳回；4-已撤销；6-通过后撤销；7-已删除；10-已支付）
	Value string `json:"value"`
}

// OAApprovalInfoFilterKey 拉取审批筛选类型
type OAApprovalInfoFilterKey string

// OAApprovalInfoFilterKeyTemplateID 模板类型
const OAApprovalInfoFilterKeyTemplateID OAApprovalInfoFilterKey = "template_id"

// OAApprovalInfoFilterKeyCreator 申请人
const OAApprovalInfoFilterKeyCreator OAApprovalInfoFilterKey = "creator"

// OAApprovalInfoFilterKeyDepartment 审批单提单者所在部门
const OAApprovalInfoFilterKeyDepartment OAApprovalInfoFilterKey = "department"

// OAApprovalInfoFilterKeySpStatus 审批状态
const OAApprovalInfoFilterKeySpStatus OAApprovalInfoFilterKey = "sp_status"
