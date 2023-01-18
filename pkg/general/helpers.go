package sheeters

func GetUser(username string) (*User, error){
  users := map[string]User{}

  err := GetJsonData(USERSDATA, &users)
  if err != nil{
    return nil, err
  }

  if user, ok := users[username]; ok{
    return &user, nil
  }

  return nil, err
}

func SetUser(u *User)(error){
  var usersList map[string]User
  err := GetJsonData(USERSDATA, &usersList)
  if err != nil{
    return err
  }

  usersList[u.Username] = *u

  err = SetJsonData(USERSDATA, usersList)
  if err != nil{
    return err
  }

  return nil
}
