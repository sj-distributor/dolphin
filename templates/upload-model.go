package templates

var UploadModel = `
type UploadFile @entity(title: "上传文件管理") {
	name: String! @column(gorm: "type:varchar(255) comment '文件名称';NOT NULL;index:name;") @validator(required: "true", repeat: "no")
	hash: String! @column(gorm: "type:text comment '文件hash值';NOT NULL;") @validator(required: "true", repeat: "no")
}

extend type Mutation {
  upload(files: [FileField!]!): Any @entity(title: "上传文件", default: 1)
}
`
