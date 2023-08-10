namespace go api


namespace go common
// 用户
struct User {
    1: required i64 id
    2: required string name
    3: optional i64 follow_count
    4: required i64 follower_count
    5: required bool is_follow
    6: optional string avatar
    7: optional string background_image
    8: optional string signature
    9: optional string total_favorited
    10: optional i64 work_count
    11: optional i64 favirite_count
}

// 视频
struct Video {
    1: required i64 id
    2: required User author
    3: required string play_url
    4: required string cover_url
    5: required i64 favorite_count
    6: required i64 comment_count
    7: required bool is_favorite
    8: required string title
}
// 聊天记录
struct Message {
	1: required i64 id
    2: required i64 to_user_id
	3: required i64 from_user_id
    4: required string content
    5: required i64 create_time
}
//评论操作
struct Comment {
  	1: required i64 id // 视频评论id
	2: required User user // 评论用户信息
   	3: required string content // 评论内容
   	4: required string create_date // 评论发布日期，格式 mm-dd
}

struct FriendUser {
    1: required i64 id
    2: required string name
    3: optional i64 follow_count
    4: required i64 follower_count
    5: required bool is_follow
    6: optional string avatar
    7: optional string background_image
    8: optional string signature
    9: optional string total_favorited
    10: optional i64 work_count
    11: optional i64 favirite_count
    12: optional string message
    13: required i64 msgType
}

//bisic operations -> doc 3.1
//Token
struct Token {
    1: optional string Token (api.query="token",api.form="token")
}
// 用户注册
struct RegisterRequest {
    1: required string email (api.query="username")
    2: required string password (api.query="password")
}
struct RegisterResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
    4: required string token
}

// 用户登录
struct LoginRequest {
    1: required string username (api.query="username")
    2: required string password (api.query="password")
}
struct LoginResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required i64 user_id
    4: required string token
}

// 用户信息
struct UserInfoRequest {
    1:required i64 user_id (api.query="user_id")
}
struct UserInfoResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required User user
}

service UserService {
    RegisterResponse Register(1: RegisterRequest req) (api.post="/douyin/user/register/")
    LoginResponse Login(1: LoginRequest req) (api.post="/douyin/user/login/")
    UserInfoResponse GetUserInfo(1: UserInfoRequest req) (api.get="/douyin/user/")
}

// 获取视频接口
struct FeedRequest {
    1: optional i64 latest_time (api.query="latest_time")
}
struct FeedResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: required list<Video> video_list
    4: optional i64 next_time
}

service FeedService {
    FeedResponse Feed(1: FeedRequest req) (api.get="/douyin/feed/")
}

// 发布视频
struct PublishActionRequest {
    1: required binary data (api.form="data")
    2: required string title (api.form="title")
}
struct PublishActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}

// 发布列表
struct PublishListRequest {
    1: required i64 user_id (api.query="user_id")
}
struct PublishListResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: optional list<Video> video_list
}

service PublishService {
    PublishActionResponse PublishAction(1: PublishActionRequest req) (api.post="/douyin/publish/dao/")
    PublishListResponse GetPublishList(1: PublishListRequest req) (api.get="/douyin/publish/list/")
}


// interaction operations -> doc 3.2
//点赞操作
struct FavoriteActionRequest {
    1: required i64 video_id (api.query="video_id")
    2: required i32 action_type (api.query="action_type")
}
struct FavoriteActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}

// 喜欢列表
struct FavoriteListRequest {
	1: required i64 user_id (api.query="user_id")
}
struct FavoriteListResponse {
    1: required i32 status_code
    2: optional string status_msg
	3: optional list<Video> video_list
}

service FavoriteService {
	FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req) (api.post="/douyin/favorite/dao/")
	FavoriteListResponse GetFavoriteList(1: FavoriteListRequest req) (api.get="/douyin/favorite/list/")
}

struct CommentActionRequest {
	1: required i64 video_id (api.query="video_id")
	2: required i32 action_type (api.query="action_type")
	3: optional string comment_text (api.query="comment_text")
	4: optional i64 comment_id (api.query="comment_id")
}
struct CommentActionResponse {
    1: required i32 status_code
    2: optional string status_msg
	3: optional Comment comment
}

// 评论列表
struct CommentListRequest {
	1: required i64 video_id (api.query="video_id")
}
struct CommentListResponse {
    1: required i32 status_code
    2: optional string status_msg
	3: required list<Comment> comment_list
}

service CommentService {
    CommentActionResponse CommentAction(1: CommentActionRequest req) (api.post="/douyin/comment/dao/")
    CommentListResponse GetComment(1: CommentListRequest req) (api.get="/douyin/comment/list/")
}


// social operations -> doc 3.3
// 关注操作
struct RelationActionRequest {
	1: required i64 to_user_id (api.query="to_user_id")
	2: required i32 action_type (api.query="action_type")
}
struct RelationActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}

// 关注列表
struct RelationFollowListRequest {
	1: required i64 user_id (api.query="user_id")
}
struct RelationFollowListResponse {
    1: required i32 status_code
    2: optional string status_msg
	3: optional list<User> user_list
}

// 粉丝列表
struct RelationFollowerListRequest {
	1: required i64 user_id (api.query="user_id")
}
struct RelationFollowerListResponse {
    1: required i32 status_code
    2: optional string status_msg
	3: optional list<User> user_list
}

// 好友列表



struct RelationFriendListRequest {
	1: required i64 user_id (api.query="user_id")
}
struct RelationFriendListResponse {
    1: required i32 status_code
    2: optional string status_msg
	3: optional list<FriendUser> user_list
}

service RelationService {
    RelationActionResponse RelationAction(1: RelationActionRequest req) (api.post="/douyin/relation/dao/")
    RelationFollowListResponse GetRelationFollowList(1: RelationFollowListRequest req) (api.get="/douyin/relation/follow/list/")
    RelationFollowerListResponse GetRelationFollowerList(1: RelationFollowerListRequest req) (api.get="/douyin/relation/follower/list/")
	RelationFriendListResponse GetRelationFriendList(1:RelationFriendListRequest req) (api.get="/douyin/relation/friend/list/")
}

// 发送消息
struct MessageActionRequest {
	1: required i64 to_user_id (api.query="to_user_id")
	2: required i32 action_type (api.query="action_type")
	3: required string content (api.query="content")
}
struct MessageActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}


struct MessageChatRequest {

	1: required i64 to_user_id (api.query="to_user_id")
    2: required i64 pre_msg_time (api.query="pre_msg_time")
}
struct MessageChatResponse {
    1: required i32 status_code
    2: optional string status_msg
	3: optional list<Message> message_list
}

service MeassgeService {
	MessageActionResponse MessageAction(1: MessageActionRequest req) (api.post="/douyin/message/dao/")
	MessageChatResponse GetMessageChat(1: MessageChatRequest req) (api.get="/douyin/message/chat/")
}
