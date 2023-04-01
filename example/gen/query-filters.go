package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"
)

type TodoQueryFilter struct {
	Query *string
}

func (qf *TodoQueryFilter) Apply(ctx context.Context, db *gorm.DB, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
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
		if err := qf.applyQueryWithFields(db, fields, part, TableName("todos"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *TodoQueryFilter) applyQueryWithFields(db *gorm.DB, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
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

	if _, ok := fieldsMap["age"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("age")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["money"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("money")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["remark"]; ok {

		column := alias + "." + SnakeString("remark")

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

	if fs, ok := fieldsMap["user"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_user"
		*joins = append(*joins, "LEFT JOIN "+"users"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"user_id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := UserQueryFilter{qf.Query}
		err := q.applyQueryWithFields(db, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

type UserQueryFilter struct {
	Query *string
}

func (qf *UserQueryFilter) Apply(ctx context.Context, db *gorm.DB, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
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
		if err := qf.applyQueryWithFields(db, fields, part, TableName("users"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *UserQueryFilter) applyQueryWithFields(db *gorm.DB, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["username"]; ok {

		column := alias + "." + SnakeString("username")

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

	if fs, ok := fieldsMap["todo"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_todo"
		*joins = append(*joins, "LEFT JOIN "+"todos"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"todo_id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := TodoQueryFilter{qf.Query}
		err := q.applyQueryWithFields(db, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
