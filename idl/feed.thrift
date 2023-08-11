namespace go feed

struct FeedReq {
    1: i64 latest_time
}

struct Video {
    1: i64 id
    2: i64 AuthorID
    3: string title
    4: string play_url
    5: string cover_url
}

struct FeedResp {
    1: list<Video> video_list
}

service FeedRpcService {
    FeedResp GetFeed(1:FeedReq req)
}