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

	if s.TID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("tId")+" "+s.TID.String())
	}

	if s.TtID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("ttId")+" "+s.TtID.String())
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

	if s.T != nil {
		_alias := alias + "_t"
		*joins = append(*joins, "LEFT JOIN "+"tasks"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"t_id")
		err := s.T.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Tt != nil {
		_alias := alias + "_tt"
		*joins = append(*joins, "LEFT JOIN "+"tasks"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"tt_id")
		err := s.Tt.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Ttt != nil {
		_alias := alias + "_ttt"
		*joins = append(*joins, "LEFT JOIN "+"tasks"+" "+_alias+" ON "+_alias+"."+"uuu_id"+" = "+alias+".id")
		err := s.Ttt.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Tttt != nil {
		_alias := alias + "_tttt"
		*joins = append(*joins, "LEFT JOIN "+"task_uuuu"+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"uuuu_id"+" LEFT JOIN "+TableName("tasks")+" "+_alias+" ON "+_alias+"_jointable"+"."+"tttt_id"+" = "+_alias+".id")
		err := s.Tttt.ApplyWithAlias(ctx, _alias, sorts, joins)
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

	if s.UID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("uId")+" "+s.UID.String())
	}

	if s.UuuID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("uuuId")+" "+s.UuuID.String())
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

	if s.U != nil {
		_alias := alias + "_u"
		*joins = append(*joins, "LEFT JOIN "+"users"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"u_id")
		err := s.U.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Uu != nil {
		_alias := alias + "_uu"
		*joins = append(*joins, "LEFT JOIN "+"users"+" "+_alias+" ON "+_alias+"."+"tt_id"+" = "+alias+".id")
		err := s.Uu.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Uuu != nil {
		_alias := alias + "_uuu"
		*joins = append(*joins, "LEFT JOIN "+"users"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"uuu_id")
		err := s.Uuu.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Uuuu != nil {
		_alias := alias + "_uuuu"
		*joins = append(*joins, "LEFT JOIN "+"task_uuuu"+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"tttt_id"+" LEFT JOIN "+TableName("users")+" "+_alias+" ON "+_alias+"_jointable"+"."+"uuuu_id"+" = "+_alias+".id")
		err := s.Uuuu.ApplyWithAlias(ctx, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
