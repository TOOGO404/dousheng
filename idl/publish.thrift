namespace go publish

struct Video {
    1: i64 id
    2: i64 AuthorID
    3: string title
    4: string play_url
    5: string cover_url
}

struct VideoData {
    1: required binary data
    2: required string title
    3: required i64 uid
    4: required string file_type
}
struct PublishActionResp {
    1: required i64 video_id
}
struct PublishListReq {
    1: required i64 user_id
}
struct PublishListResp {
    1: optional list<Video> video_list
}

service PublishRpcService {
    PublishActionResp PublishAction(1:VideoData data)
    PublishListResp GetPublishLish(1:PublishListReq req)
}