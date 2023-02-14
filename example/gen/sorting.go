package gen

import (
	"context"
)

func (s TodoSortType) Apply(ctx context.Context, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, TableName("todos"), sorts, joins)
}
func (s TodoSortType) ApplyWithAlias(ctx context.Context, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := alias + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("id")+" "+s.ID.String())
	}

	if s.Title != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("title")+" "+s.Title.String())
	}

	if s.Age != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("age")+" "+s.Age.String())
	}

	if s.Money != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("money")+" "+s.Money.String())
	}

	if s.Remark != nil {
		*sorts = append(*sorts, aliasPrefix+SnakeString("remark")+" "+s.Remark.String())
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

	return nil
}
