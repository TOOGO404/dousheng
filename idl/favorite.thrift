namespace go favorite
struct FavoriteActionRequest {
    1: required i64 uid
    2: required i64 video_id
    3: required i32 action_type
}
struct FavoriteActionResponse {
    1: required i32 status_code
    2: optional string status_msg
}
struct FavoriteListRequest {
    1: required i64 uid
	2: required i64 user_id
}
struct FavoriteListResponse {
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
service FavoriteRpcService {
	FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req)
	FavoriteListResponse GetFavoriteList(1: FavoriteListRequest req)
}
