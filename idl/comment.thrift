namespace go comment
struct CommentActionRequest {
  1: required i64 uid // 用户鉴权token
  2: required i64 video_id  // 视频id
  3: required i32 action_type  // 1-发布评论，2-删除评论
  4: optional string comment_text  // 用户填写的评论内容，在action_type=1的时候使用
  5: optional i64 comment_id // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse {
  1: required i32 status_code  // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: optional Comment comment  // 评论成功返回评论内容，不需要重新拉取整个列表
}


struct CommentListRequest {
  1: required i64 uid  // 用户鉴权token
  2: required i64 video_id // 视频id
}
struct CommentListResponse {
  1: required i32 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg  // 返回状态描述
  3: list<Comment> comment_list
}

struct Comment {
  1: required i64 id  // 视频评论id
  2: required User user  // 评论用户信息
  3: required string content  // 评论内容
  4: required string create_date // 评论发布日期，格式 mm-dd
}

struct User {
        1: required i64 id
        2: required string name
        3: optional i64 follow_count
        4: required i64 follower_count
        5: required bool is_follow
        6: optional string avatar
        7: optional string background_image
        8: optional string signature
        9: optional i64 total_favorited
        10: optional i64 work_count
        11: optional i64 favirite_count
}
service CommentRpcService {
    CommentListResponse CommentGet(1:CommentListRequest req)
    CommentActionResponse CommentAction(1:CommentActionRequest req)
    i64 GetCommentCnt(1:i64 vid)
}
