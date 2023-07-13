package dao_test


import(
	"testing"
	"github.com/liujiage/restapi/service/dao"
	"github.com/liujiage/restapi/common"
)


func TestUserDao_UserAdd(t *testing.T){
	//common.GetDBConnect()
	userId := common.GetUUID()
    user := dao.UserModelDao{Id: userId, Name: "test"}
	user.UserAdd()
}