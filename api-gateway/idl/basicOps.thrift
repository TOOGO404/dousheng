namespace go api

//bisic operations -> doc 3.1
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

// 用户注册
struct RegisterRequest {
    1: required string username (api.query="username")
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
    2:required string token (api.query="token")
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
    2: optional string token (api.query="token")
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
    1: required string token (api.form="token")
    2: required binary data (api.form="data")
    3: required string title (api.form="title")
}
struct PublishActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}

// 发布列表
struct PublishListRequest {
    1: required i64 user_id (api.query="user_id")
    2: required string token (api.query="token")
}
struct PublishListResponse {
    1: required i32 status_code
    2: optional string status_msg
    3: optional list<Video> video_list
}

service PublishService {
    PublishActionResponse PublishAction(1: PublishActionRequest req) (api.post="/douyin/publish/action/")
    PublishListResponse GetPublishList(1: PublishListRequest req) (api.get="/douyin/publish/list/")
}
