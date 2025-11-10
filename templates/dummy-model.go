/*
 * @Author: Marlon.M
 * @Email: maiguangyang@163.com
 * @Date: 2025-03-17 14:43:51
 */
package templates

var DummyModel = `
directive @hasRole(role: Role!) on FIELD_DEFINITION

enum Role {
  ALL
  ADMIN
  USER
  GUEST
}

type User @entity(title: "用户管理") {
  phone: String! @column(gorm: "type:varchar(32) comment '账号：使用手机号码';NOT NULL;index:phone;") @validator(required: "true", type: "phone")
  password: String! @column(gorm: "type:varchar(64) comment '登录密码';NOT NULL;") @validator(required: "true", type: "password")
	email: String @column(gorm: "type:varchar(64) comment '用户邮箱地址';default:null;") @validator(required: "true", type: "email")
	nickname: String @column(gorm: "type:varchar(64) comment '昵称';DEFAULT NULL;index:nickname;")
	age: Int @column(gorm: "type:int(3) comment '年龄';default:1;") @validator(type: "int")
	lastName: String @column
	tasks: [Task!]! @relationship(inverse:"user")
}

type Task @entity(title: "任务管理") {
	title: String @column
	completed: Boolean @column
	dueDate: Time @column
	user: User @relationship(inverse:"tasks")
}

extend type Subscription {
  webSocket: Any
}
`
