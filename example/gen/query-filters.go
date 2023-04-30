package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"
)

type BookCategoryQueryFilter struct {
	Query *string
}

func (qf *BookCategoryQueryFilter) Apply(ctx context.Context, db *gorm.DB, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field.")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(db, fields, part, TableName("book_categories"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *BookCategoryQueryFilter) applyQueryWithFields(db *gorm.DB, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["name"]; ok {

		column := alias + "." + SnakeString("name")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["description"]; ok {

		column := alias + "." + SnakeString("description")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["deletedAt"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("deletedAt")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["updatedAt"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("updatedAt")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["createdAt"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("createdAt")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if fs, ok := fieldsMap["books"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_books"
		*joins = append(*joins, "LEFT JOIN "+"books"+" "+_alias+" ON "+_alias+"."+"category_id"+" = "+alias+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := BookQueryFilter{qf.Query}
		err := q.applyQueryWithFields(db, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

type BookQueryFilter struct {
	Query *string
}

func (qf *BookQueryFilter) Apply(ctx context.Context, db *gorm.DB, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field.")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(db, fields, part, TableName("books"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *BookQueryFilter) applyQueryWithFields(db *gorm.DB, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["title"]; ok {

		column := alias + "." + SnakeString("title")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["author"]; ok {

		column := alias + "." + SnakeString("author")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["price"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("price")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["publishDateAt"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("publishDateAt")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["deletedAt"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("deletedAt")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["updatedAt"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("updatedAt")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["createdAt"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("createdAt")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if fs, ok := fieldsMap["category"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_category"
		*joins = append(*joins, "LEFT JOIN "+"book_categories"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"category_id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := BookCategoryQueryFilter{qf.Query}
		err := q.applyQueryWithFields(db, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
