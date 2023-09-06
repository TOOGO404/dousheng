namespace go feed


//feed rpc 请求和回复，
struct Video {
    1: i64 id
    2: i64 AuthorID
    3: string title
    4: string play_url
    5: string cover_url
}


struct FeedReq {
    1:required i64 latest_time
}
struct FeedResp {
    1: list<Video> video_list
    2: i64 next_time
}

service FeedRpcService {
    FeedResp GetFeed(1:FeedReq req)
}