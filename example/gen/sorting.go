package gen

import (
	"context"
)

func (s UserSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("users"), sorts, joins)
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
		*joins = append(*joins, "LEFT JOIN "+"tasks"+" "+_alias+" ON "+_alias+"."+"user_id"+" = "+alias+".id")
		err := s.Tasks.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s TaskSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("tasks"), sorts, joins)
}
func (s TaskSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.Title != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("title")+" "+s.Title.String())
	}

	if s.UserID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("userId")+" "+s.UserID.String())
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
		*joins = append(*joins, "LEFT JOIN "+"users"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"user_id")
		err := s.User.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
