package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

func (*UserInfoAtomicService) Save(u *UserInfo) error {
	session := myEngine.NewSession()
	defer session.Close()
	err := session.Begin()
	checkErr(err)

	dao := userInfoDao{session}
	err = dao.Save(u)

	if err == nil {
		session.Commit()
	} else {
		session.Rollback()
	}
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	dao := userInfoDao{myEngine}
	return dao.FindAll()
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	dao := userInfoDao{myEngine}
	return dao.FindByID(id)
}