package feed_service

import (
	"context"
	"feed-service/dal/dao"
	feed "feed-service/kitex_gen/feed"
	"log"
)

// FeedRpcServiceImpl implements the last service interface defined in the IDL.
type FeedRpcServiceImpl struct{}

// GetFeed implements the FeedRpcServiceImpl interface.
func (s *FeedRpcServiceImpl) GetFeed(ctx context.Context, req *feed.FeedReq) (*feed.FeedResp, error) {
	videoList := dao.GetVIdeoListByLatestTime(req.LatestTime)
	var resp = new(feed.FeedResp)
	videos := make([]*feed.Video, len(videoList))
	for index, video := range videoList {
		fv := new(feed.Video)
		fv.Id = video.Id
		fv.AuthorID = video.Author
		fv.Title = video.Title
		fv.PlayUrl = video.PlayUrl
		fv.CoverUrl = video.CoverUrl
		videos[index] = fv
	}
	log.Println("this time get video num:", len(videos))
	resp.VideoList = videos
	if l := len(videos); l != 0 {
		resp.NextTime = videoList[l-1].CreatedAt
	}

	return resp, nil
}
