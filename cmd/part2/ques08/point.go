package ques08

type User struct {
	isSpecialMember bool
	isCardMember    bool
	isMoblieMember  bool
}

func NewUser(isSpecialMember bool, isCardMember bool, isMobileMember bool) User {
	return User{
		isSpecialMember: isSpecialMember,
		isCardMember:    isCardMember,
		isMoblieMember:  isMobileMember,
	}
}

func GetPointMultiplier(user User, day int) int {
	point := 0
	if user.isSpecialMember {
		point += 3
	}
	if user.isCardMember {
		point += 2
	}
	if user.isMoblieMember {
		point += 5
	}
	if day == 15 {
		point += 3
	}
	return point
}
