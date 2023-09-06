namespace go favorite
struct FavoriteActionReq {
    1: required i64 uid
    2: required i64 video_id
    3: required i32 action_type
}
struct FavoriteActionResp {
    1: required i32 status_code
    2: optional string status_msg
}
struct FavoriteListReq {
    1: required i64 uid
	2: required i64 user_id
}
struct FavoriteListResp {
    1: required i32 status_code
    2: optional string status_msg
	3: optional list<Video> video_list
}

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
struct User{
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

struct CheckFavoritedReq {
    1: required i64 uid 
    2: required i64 vid 
}
service FavoriteRpcService {
	FavoriteActionResp FavoriteAction(1: FavoriteActionReq req)
	FavoriteListResp GetFavoriteList(1: FavoriteListReq req)
    i64 GetTotalFavorited(1:i64 uid)
    bool IsFavorited(1:CheckFavoritedReq req)
    i64 GetFavorCount(1:i64 uid) 
}