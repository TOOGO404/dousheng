package relationship_impl

import (
	"context"
	"errors"
	"relationship-service/dal/dao"
	relationship "relationship-service/kitex_gen/relationship"
)

// RelationshipRpcServiceImpl implements the last service interface defined in the IDL.
type RelationshipRpcServiceImpl struct{}

// Sub implements the RelationshipRpcServiceImpl interface.
func (s *RelationshipRpcServiceImpl) Sub(ctx context.Context, req *relationship.SubActionReq) (bool, error) {
	switch req.ActionType {
	case relationship.Action_Sub:
		{
			return dao.AddSubRecord(req.Who, req.ToUserId)

		}
	case relationship.Action_Cancel:
		{
			return dao.RemoveSubRecord(req.Who, req.ToUserId)
		}
	}
	return false, errors.New("no such action type")
}

// CheckSub implements the RelationshipRpcServiceImpl interface.
func (s *RelationshipRpcServiceImpl) CheckSub(ctx context.Context, req *relationship.CheckReq) (bool, error) {
	return dao.IsSub(req.Who, req.ToUserId), nil
}

// GetSubList implements the RelationshipRpcServiceImpl interface.
func (s *RelationshipRpcServiceImpl) GetSubList(ctx context.Context, req *relationship.SubListReq) (*relationship.SubListResp, error) {
	var resp = new(relationship.SubListResp)
	ui := dao.GetSubList(req.Who)
	resp.FollowList = ui
	return resp, nil
}

// GetFollowerList implements the RelationshipRpcServiceImpl interface.
func (s *RelationshipRpcServiceImpl) GetFollowerList(ctx context.Context, req *relationship.FollowerListReq) (*relationship.FollowerListResp, error) {
	var resp = new(relationship.FollowerListResp)
	ui := dao.GetFollowerList(req.Who)
	resp.FollowerList = ui
	return resp, nil
}

// GetFriendList implements the RelationshipRpcServiceImpl interface.
func (s *RelationshipRpcServiceImpl) GetFriendList(ctx context.Context, req *relationship.FriendlistReq) (*relationship.FriendlistResp, error) {
	var resp = new(relationship.FriendlistResp)
	ui := dao.GetFriendList(req.Who)
	resp.FriendList = ui
	return resp, nil
}

func (s *RelationshipRpcServiceImpl) GetFollowCnt(ctx context.Context, uid int64) (int64, error) {
	return dao.GetFollowCnt(uid), nil
}
func (s *RelationshipRpcServiceImpl) GetFollowerCnt(ctx context.Context, uid int64) (int64, error) {
	return dao.GetFollowerCnt(uid), nil
}
