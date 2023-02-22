package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

//关注和取消关注需要返回结构。
type RelationActionResp struct {
	Response
}

//获取关注列表需要返回的结构。
type FollowingResp struct {
	Response
	UserList []service.User `json:"user_list,omitempty"`
}

//获取粉丝列表需要返回的结构。
type FollowersResp struct {
	Response
	// 序列化
	UserList []service.User `json:"user_list,omitempty"`
}

//处理关注和取消关注请求
func RelationActionFuction(c *gin.Context) {
	//获取用户id
	userId, err1 := strconv.ParseInt(c.GetString("userId"), 10, 64)
	toUserId, err2 := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	//请求参数类型
	actionType, err3 := strconv.ParseInt(c.Query("action_type"), 10, 64)
	// fmt.Println(userId, toUserId, actionType)
	// 传入参数格式有问题。
	if nil != err1 || nil != err2 || nil != err3 || actionType < 1 || actionType > 2 {
		c.JSON(http.StatusOK, RelationActionResp{
			Response{
				StatusCode: -1,
				StatusMsg:  "用户id格式错误",
			},
		})
		return
	}
	// 正常处理
	fsi := service.NewFSIInstance()
	switch {
	// 关注
	case 1 == actionType:
		go fsi.AddFollowRelation(userId, toUserId)
	// 取关
	case 2 == actionType:
		go fsi.DeleteFollowRelation(userId, toUserId)
	}
	//设置返回状态码
	c.JSON(http.StatusOK, RelationActionResp{
		Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
	})
}

// 处理获取关注列表请求。
func GetFollowing(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	// 用户id解析出错。
	if nil != err {
		c.JSON(http.StatusOK, FollowingResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "用户id格式错误。",
			},
			UserList: nil,
		})
		return
	}
	// 正常获取关注列表
	fsi := service.NewFSIInstance()
	users, err := fsi.GetFollowing(userId)
	// 获取关注列表时出错。
	if err != nil {
		c.JSON(http.StatusOK, FollowingResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "获取关注列表时出错。",
			},
			UserList: nil,
		})
		return
	}
	// 成功获取到关注列表
	//设置状态码
	c.JSON(http.StatusOK, FollowingResp{
		UserList: users,
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
	})
}

// 处理获取关注列表请求
func GetFollowers(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	// 用户id解析出错
	if nil != err {
		c.JSON(http.StatusOK, FollowersResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "用户id格式错误。",
			},
			UserList: nil,
		})
		return
	}
	// 正常获取粉丝列表
	fsi := service.NewFSIInstance()
	users, err := fsi.GetFollowers(userId)
	// 获取关注列表时出错。
	if err != nil {
		c.JSON(http.StatusOK, FollowersResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "获取粉丝列表时出错。",
			},
			UserList: nil,
		})
		return
	}
	// 成功获取到粉丝列表
	//设置状态数据
	c.JSON(http.StatusOK, FollowersResp{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
		UserList: users,
	})
}
