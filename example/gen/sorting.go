package gen

import (
	"context"
)

func (s UserSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("users", ctx), sorts, joins)
}
func (s UserSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.Phone != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("phone")+" "+s.Phone.String())
	}

	if s.Password != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("password")+" "+s.Password.String())
	}

	if s.Email != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("email")+" "+s.Email.String())
	}

	if s.Nickname != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("nickname")+" "+s.Nickname.String())
	}

	if s.Age != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("age")+" "+s.Age.String())
	}

	if s.ProfileID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("profileId")+" "+s.ProfileID.String())
	}

	if s.IsDelete != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("isDelete")+" "+s.IsDelete.String())
	}

	if s.Weight != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("weight")+" "+s.Weight.String())
	}

	if s.State != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("state")+" "+s.State.String())
	}

	if s.DeletedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedBy")+" "+s.DeletedBy.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdBy")+" "+s.CreatedBy.String())
	}

	if s.DeletedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedAt")+" "+s.DeletedAt.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdAt")+" "+s.CreatedAt.String())
	}

	if s.Profile != nil {
		_alias := alias + "_profile"
		*joins = append(*joins, "LEFT JOIN "+TableName("profiles", ctx)+" "+_alias+" ON "+_alias+".id = "+alias+"."+"profile_id")
		err := s.Profile.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+TableName("tasks", ctx)+" "+_alias+" ON "+_alias+"."+"user_id"+" = "+alias+".id")
		err := s.Tasks.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.UserRoles != nil {
		_alias := alias + "_userRoles"
		*joins = append(*joins, "LEFT JOIN "+TableName("userRole_users", ctx)+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"user_id"+" LEFT JOIN "+TableName("user_roles", ctx)+" "+_alias+" ON "+_alias+"_jointable"+"."+"user_role_id"+" = "+_alias+".id")
		err := s.UserRoles.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s ProfileSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("profiles", ctx), sorts, joins)
}
func (s ProfileSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.Avatar != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("avatar")+" "+s.Avatar.String())
	}

	if s.Bio != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("bio")+" "+s.Bio.String())
	}

	if s.Birthday != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("birthday")+" "+s.Birthday.String())
	}

	if s.Address != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("address")+" "+s.Address.String())
	}

	if s.UserID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("userId")+" "+s.UserID.String())
	}

	if s.IsDelete != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("isDelete")+" "+s.IsDelete.String())
	}

	if s.Weight != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("weight")+" "+s.Weight.String())
	}

	if s.State != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("state")+" "+s.State.String())
	}

	if s.DeletedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedBy")+" "+s.DeletedBy.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdBy")+" "+s.CreatedBy.String())
	}

	if s.DeletedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedAt")+" "+s.DeletedAt.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdAt")+" "+s.CreatedAt.String())
	}

	if s.User != nil {
		_alias := alias + "_user"
		*joins = append(*joins, "LEFT JOIN "+TableName("users", ctx)+" "+_alias+" ON "+_alias+".id = "+alias+"."+"user_id")
		err := s.User.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s TaskSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("tasks", ctx), sorts, joins)
}
func (s TaskSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.Title != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("title")+" "+s.Title.String())
	}

	if s.Description != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("description")+" "+s.Description.String())
	}

	if s.Completed != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("completed")+" "+s.Completed.String())
	}

	if s.DueDate != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("dueDate")+" "+s.DueDate.String())
	}

	if s.Priority != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("priority")+" "+s.Priority.String())
	}

	if s.UserID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("userId")+" "+s.UserID.String())
	}

	if s.IsDelete != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("isDelete")+" "+s.IsDelete.String())
	}

	if s.Weight != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("weight")+" "+s.Weight.String())
	}

	if s.State != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("state")+" "+s.State.String())
	}

	if s.DeletedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedBy")+" "+s.DeletedBy.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdBy")+" "+s.CreatedBy.String())
	}

	if s.DeletedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedAt")+" "+s.DeletedAt.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdAt")+" "+s.CreatedAt.String())
	}

	if s.User != nil {
		_alias := alias + "_user"
		*joins = append(*joins, "LEFT JOIN "+TableName("users", ctx)+" "+_alias+" ON "+_alias+".id = "+alias+"."+"user_id")
		err := s.User.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Tags != nil {
		_alias := alias + "_tags"
		*joins = append(*joins, "LEFT JOIN "+TableName("tag_tasks", ctx)+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"task_id"+" LEFT JOIN "+TableName("tags", ctx)+" "+_alias+" ON "+_alias+"_jointable"+"."+"tag_id"+" = "+_alias+".id")
		err := s.Tags.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s UserRoleSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("user_roles", ctx), sorts, joins)
}
func (s UserRoleSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.Name != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("name")+" "+s.Name.String())
	}

	if s.Description != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("description")+" "+s.Description.String())
	}

	if s.Permissions != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("permissions")+" "+s.Permissions.String())
	}

	if s.IsDelete != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("isDelete")+" "+s.IsDelete.String())
	}

	if s.Weight != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("weight")+" "+s.Weight.String())
	}

	if s.State != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("state")+" "+s.State.String())
	}

	if s.DeletedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedBy")+" "+s.DeletedBy.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdBy")+" "+s.CreatedBy.String())
	}

	if s.DeletedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedAt")+" "+s.DeletedAt.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdAt")+" "+s.CreatedAt.String())
	}

	if s.Users != nil {
		_alias := alias + "_users"
		*joins = append(*joins, "LEFT JOIN "+TableName("userRole_users", ctx)+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"userRole_id"+" LEFT JOIN "+TableName("users", ctx)+" "+_alias+" ON "+_alias+"_jointable"+"."+"user_id"+" = "+_alias+".id")
		err := s.Users.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s TagSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("tags", ctx), sorts, joins)
}
func (s TagSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.Name != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("name")+" "+s.Name.String())
	}

	if s.Color != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("color")+" "+s.Color.String())
	}

	if s.IsDelete != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("isDelete")+" "+s.IsDelete.String())
	}

	if s.Weight != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("weight")+" "+s.Weight.String())
	}

	if s.State != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("state")+" "+s.State.String())
	}

	if s.DeletedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedBy")+" "+s.DeletedBy.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedBy")+" "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdBy")+" "+s.CreatedBy.String())
	}

	if s.DeletedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("deletedAt")+" "+s.DeletedAt.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("updatedAt")+" "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("createdAt")+" "+s.CreatedAt.String())
	}

	if s.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+TableName("tag_tasks", ctx)+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"tag_id"+" LEFT JOIN "+TableName("tasks", ctx)+" "+_alias+" ON "+_alias+"_jointable"+"."+"task_id"+" = "+_alias+".id")
		err := s.Tasks.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
