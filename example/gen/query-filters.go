package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"
)

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

	if *qf.Query != "" {
		queryParts := strings.Split(*qf.Query, " ")
		for _, part := range queryParts {
			ors := []string{}
			if err := qf.applyQueryWithFields(db, fields, part, TableName("users"), &ors, values, joins); err != nil {
				return err
			}
			*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
		}
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

	if _, ok := fieldsMap["phone"]; ok {

		column := alias + "." + SnakeString("phone")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["password"]; ok {

		column := alias + "." + SnakeString("password")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["email"]; ok {

		column := alias + "." + SnakeString("email")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["nickname"]; ok {

		column := alias + "." + SnakeString("nickname")

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

	if _, ok := fieldsMap["lastName"]; ok {

		column := alias + "." + SnakeString("lastName")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["isDelete"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("isDelete")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["weight"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("weight")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["state"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("state")+" AS %s)", alias+".", cast)

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

	// if fs, ok := fieldsMap["tasks"]; ok {
	// 	_fields := []*ast.Field{}
	// 	_alias := alias + "_tasks"
	// 	*joins = append(*joins,"LEFT JOIN "+"tasks"+" "+_alias+" ON "+_alias+"."+"user_id"+" = "+alias+".id")

	// 	for _, f := range fs {
	// 		for _, s := range f.SelectionSet {
	// 			if f, ok := s.(*ast.Field); ok {
	// 				_fields = append(_fields, f)
	// 			}
	// 		}
	// 	}
	// 	q := TaskQueryFilter{qf.Query}
	// 	err := q.applyQueryWithFields(db, _fields, query, _alias, ors, values, joins)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

type TaskQueryFilter struct {
	Query *string
}

func (qf *TaskQueryFilter) Apply(ctx context.Context, db *gorm.DB, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
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

	if *qf.Query != "" {
		queryParts := strings.Split(*qf.Query, " ")
		for _, part := range queryParts {
			ors := []string{}
			if err := qf.applyQueryWithFields(db, fields, part, TableName("tasks"), &ors, values, joins); err != nil {
				return err
			}
			*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
		}
	}
	return nil
}

func (qf *TaskQueryFilter) applyQueryWithFields(db *gorm.DB, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
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

	if _, ok := fieldsMap["isDelete"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("isDelete")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["weight"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("weight")+" AS %s)", alias+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "%"+query+"%")
	}

	if _, ok := fieldsMap["state"]; ok {

		cast := "TEXT"
		if db.Name() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+SnakeString("state")+" AS %s)", alias+".", cast)

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

	// if fs, ok := fieldsMap["user"]; ok {
	// 	_fields := []*ast.Field{}
	// 	_alias := alias + "_user"
	// 	*joins = append(*joins,"LEFT JOIN "+"users"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"user_id")

	// 	for _, f := range fs {
	// 		for _, s := range f.SelectionSet {
	// 			if f, ok := s.(*ast.Field); ok {
	// 				_fields = append(_fields, f)
	// 			}
	// 		}
	// 	}
	// 	q := UserQueryFilter{qf.Query}
	// 	err := q.applyQueryWithFields(db, _fields, query, _alias, ors, values, joins)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
