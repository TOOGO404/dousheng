namespace go message

// 发送消息
struct MessageActionReq {
	1:required i64 user_id
	2:required i64 to_user_id
	3:required i32 action_type
	4:optional string msg
}

// 聊天记录
struct Message {
	1:required i64 id
    2:required string msg
	3:required i64 create_time
    4:required i64 from_user_id
    5:required i64 to_user_id
}

struct MessageChatReq {
	1:required i64 user_id
	2:required i64 to_user_id
    3:required i64 pre_msg_time
}

struct MessageChatResp {
	3:required list<Message> message_list
}

service MeassgeRpcService {
	bool MessageAction(1: MessageActionReq req)
	MessageChatResp GetMessageChat(1: MessageChatReq req)
}