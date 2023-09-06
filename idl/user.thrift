namespace go user


//注册rpc请求和回复
struct RegisterReq {
    1:required string email
    2:required string pwd
}
struct RegisterResp {
    1: required i64 uid
}
//登陆rpc请求和回复
struct LoginReq {
    1:required string email
    2:required string pwd
}
struct LoginResp {
    3: required i64 uid
}

//用户信息及其请求回复
struct UserInfo {
        1: required i64 id
        2: required string name
        3: required i64 follow_count
        4: required i64 follower_count
        5: required bool is_follow
        6: required string avatar
        7: required string background_image
        8: required string signature
        9: required i64 total_favorited
        10: required i64 work_count
        11: required i64 favirite_count
}
struct UserInfoReq {
     1: required i64 send_req_user_id
     2: required i64 req_user_id
}
struct UserInfoResp {
     3: required UserInfo userInfo
}

service UserRpcService {
    RegisterResp UserRegister(1:RegisterReq req)
    LoginResp UserLogin(1:LoginReq req)
    UserInfoResp GetUserInfo(1:UserInfoReq req)
}