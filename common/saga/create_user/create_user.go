package create_user

type User struct {
	Id        string
	IsPrivate bool
}

type CreateUserCommandType int8

const (
	CreateUserConnection CreateUserCommandType = iota
	ApproveUser
	UnknownCommand
)

type CreateUserCommand struct {
	User User
	Type CreateUserCommandType
}

type CreateUserReplyType int8

const (
	UserConnectionCreated CreateUserReplyType = iota
	UserConnectionNotCreated
	UnknownReply
)

type CreateUserReply struct {
	User User
	Type CreateUserReplyType
}
