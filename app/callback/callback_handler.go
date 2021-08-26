package callback

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"openscrm/app/callback/customer_event"
	"openscrm/app/callback/department_event"
	"openscrm/app/callback/group_chat_event"
	"openscrm/app/callback/staff_event"
	"openscrm/app/callback/tag_event"
	"openscrm/app/services"
	"openscrm/common/app"
	"openscrm/common/log"
	"openscrm/common/util"
	"openscrm/common/we_work"
	"openscrm/pkg/easywework"
	"runtime"
)

var eventHandlers map[services.Event]services.CallbackHandlerFunc

func init() {
	eventHandlers = map[services.Event]services.CallbackHandlerFunc{
		// 添加客户事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeExternalContact,
			ChangeType:  workwx.ChangeTypeAddExternalContact}: staff_event.EventAddExternalContactHandler,

		//	删除企业客户事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeExternalContact,
			ChangeType:  workwx.ChangeTypeDelExternalContact}: staff_event.EventDelExternalContactHandler,

		//	编辑企业客户事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeExternalContact,
			ChangeType:  workwx.ChangeTypeEditExternalContact}: staff_event.EventEditExternalContactHandler,

		// 更新员工事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeContact,
			ChangeType:  workwx.ChangeTypeUpdateUser}: staff_event.EventUpdateStaffHandler,

		// 客户删除员工
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeExternalContact,
			ChangeType:  workwx.ChangeTypeDelFollowUser}: customer_event.EventDelFollowUserHandler,

		//	新建部门事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeContact,
			ChangeType:  workwx.ChangeTypeCreateParty}: department_event.EventCreateDepartment,

		//	更新部门事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeContact,
			ChangeType:  workwx.ChangeTypeUpdateParty}: department_event.EventUpdateDepartment,

		//	删除部门事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeContact,
			ChangeType:  workwx.ChangeTypeDeleteParty}: department_event.EventDeleteDepartment,

		//	新建标签事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeExternalTag,
			ChangeType:  workwx.ChangeTypeCreateTag}: tag_event.EventCreateExternalTagHandler,

		// 更新标签事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeExternalTag,
			ChangeType:  workwx.ChangeTypeUpdateTag}: tag_event.EventUpdateExternalTagHandler,

		//	删除标签事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeExternalTag,
			ChangeType:  workwx.ChangeTypeDeleteTag}: tag_event.EventDeleteExternalTagHandler,

		//	新建群聊事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeExternalChat,
			ChangeType:  workwx.ChangeTypeCreateChat}: group_chat_event.EventCreateExternalChatHandler,

		//	更新群聊事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeExternalChat,
			ChangeType:  workwx.ChangeTypeUpdateChat}: group_chat_event.EventUpdateExternalChatHandler,

		//	解散群聊事件
		services.Event{
			MessageType: workwx.MessageTypeEvent,
			EventType:   workwx.EventTypeChangeExternalChat,
			ChangeType:  workwx.ChangeTypeDismissChat}: group_chat_event.EventDismissExternalChatHandler,
	}
}

type Handler struct {
	dispatcher map[services.Event]services.CallbackHandlerFunc
}

func NewHandler() *Handler {
	return &Handler{dispatcher: eventHandlers}
}
func (o *Handler) RegisterHandlerFuncs(handlerFuncs map[services.Event]services.CallbackHandlerFunc) {
	for event, handlerFunc := range handlerFuncs {
		o.dispatcher[event] = handlerFunc
	}
}

// RegisterHandlerFunc not thread-safe
func (o *Handler) RegisterHandlerFunc(event services.Event, handler services.CallbackHandlerFunc) {
	o.dispatcher[event] = handler
}

// HandleCallback 回调处理入口
func (o *Handler) HandleCallback(c *gin.Context) {
	// 验证callback url
	if c.Request.Method == http.MethodGet {
		we_work.Callback.EchoTestHandler(c.Writer, c.Request)
		return
	}

	handler := app.NewHandler(c)
	msg, err := we_work.Callback.GetCallBackMsg(c.Request)
	if err != nil {
		handler.ResponseError(err)
		return
	}

	log.Sugar.Debug(util.JsonEncode(msg))

	//err = o.HandleCallbackMsg(msg)
	err = o.ProtectRun(o.HandleCallbackMsg, msg)
	if err != nil {
		handler.ResponseError(err)
		return
	}
	handler.ResponseItem(nil)
}

func (o *Handler) GetEventHandlerFunc(
	msgType workwx.MessageType, event workwx.EventType, changeType workwx.ChangeType) (services.CallbackHandlerFunc, error) {
	e := services.Event{
		MessageType: msgType,
		EventType:   event,
		ChangeType:  changeType,
	}
	handler, ok := o.dispatcher[e]
	if !ok {
		return nil, errors.New("get callback handler failed")
	}
	return handler, nil
}

func (o *Handler) ProtectRun(entry func(msg *workwx.RxMessage) error, msg *workwx.RxMessage) (err error) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			log.Sugar.Error("runtime error:", err)
		default: // 非运行时错误
			log.Sugar.Error("error:", err)
		}
	}()
	err = entry(msg)
	if err != nil {
		return
	}
	return
}

func (o *Handler) HandleCallbackMsg(msg *workwx.RxMessage) error {
	fn, err := o.GetEventHandlerFunc(msg.MsgType, msg.Event, msg.ChangeType)
	if err != nil {
		log.Sugar.Errorw("GetEventHandlerFunc failed", "err", err)
		return err
	}
	return fn(msg)
}
