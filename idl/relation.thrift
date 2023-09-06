namespace go relationship

struct UserInfo {
        1: required i64 id
        2: required string name
        5: required bool is_follow
        6: required string avatar
}


struct SubActionReq {
    1: required i64 who
    2: required i64 to_user_id
    3: required i32 action_type
}


const i32 Action_Sub = 1
const i32 Action_Cancel = 2
struct SubListReq {
    1: required i64 who
}

struct SubListResp {
    1: optional list<UserInfo> follow_list
}

struct FollowerListReq {
    1: required i64 who
}

struct FollowerListResp {
    1: optional list<UserInfo> follower_list
}

struct FriendlistReq {
    1:required i64 who
}

struct FriendlistResp {
    1:optional list<UserInfo> friend_list
}

struct CheckReq {
    1:required i64 who
    2:required i64 to_user_id
}

service RelationshipRpcService {
    bool Sub(1:SubActionReq req)
    bool CheckSub(1:CheckReq req)
    SubListResp GetSubList(1:SubListReq req)
    FollowerListResp GetFollowerList(1:FollowerListReq req)
    i64 GetFollowCnt(1: i64 uid)
    i64 GetFollowerCnt(1: i64 uid)
    FriendlistResp GetFriendList(1:FriendlistReq req)
}


