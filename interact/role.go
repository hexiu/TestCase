package interact

type Permissions int64

// 原子权限
const (
	Apply      Permissions = 1
	Invited    Permissions = 2
	Replied    Permissions = 3
	Close      Permissions = 4
	CloseVoice Permissions = 5
)

// 规则角色拼装
type RoleItem struct {
	// 角色权限细则
	Permissions []Permissions
}

func (r RoleItem) Can(p Permissions) bool {
	for _, v := range r.Permissions {
		if v == p {
			return true
		}
	}
	return false
}

// 房主
var Owner = &RoleItem{
	Permissions: []Permissions{
		Apply,
		Invited,
		Close,
		CloseVoice,
		Replied,
	},
}

var Admin = &RoleItem{
	Permissions: []Permissions{
		Invited,
		Close,
		Replied,
	},
}

var User = &RoleItem{
	Permissions: []Permissions{},
}

type Role interface {
	// 是否角色有这个权限判定
	Can(p Permissions) bool
}

// 角色关联 初始化 role
func GetRole(user interface{}) Role {

	return RoleItem{}
}
