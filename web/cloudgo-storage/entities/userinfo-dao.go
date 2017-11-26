package entities

type userInfoDao DaoSource

var userInfoInsertStmt = "INSERT UserInfo SET username=?,departname=?,created=?"

func (dao *userInfoDao) Save(u *UserInfo) error {
	res, err := dao.Exec(userInfoInsertStmt,u.UserName, u.DepartName, u.Created)
	checkErr(err)
	if err != nil {
        return err
	}
	id, err := res.LastInsertId()
    if err != nil {
        return err
    }
    u.UID = int(id)
	return nil
}

var userInfoQueryByID = "uid = ?"

func (dao *userInfoDao) FindAll() []UserInfo {
	user := new(UserInfo)
	rows, err := dao.Rows(user)
	checkErr(err)
	defer rows.Close()

	ulist := make([]UserInfo, 0, 0)
	for rows.Next() {
		u := &UserInfo{}
		err := rows.Scan(u)
        checkErr(err)
        ulist = append(ulist, *u)
	}
	return ulist
}

func (dao *userInfoDao) FindByID(id int) *UserInfo {
	user := new(UserInfo)
	row, err := dao.Where(userInfoQueryByID,id).Rows(user)
	checkErr(err)
	defer row.Close()

	u := &UserInfo{}
	for row.Next() {
		err2 := row.Scan(u)
		checkErr(err2)
	}

	return u
}