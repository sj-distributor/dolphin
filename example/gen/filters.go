package gen

import (
	"context"
	"fmt"
	"strings"
)

// CompactNils removes nil pointers from a slice.
func CompactNils[T any](items []*T) []*T {
	result := make([]*T, 0, len(items))
	for _, item := range items {
		if item != nil {
			result = append(result, item)
		}
	}
	return result
}

func (f *UserFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *UserFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("users", ctx), wheres, values, joins)
}
func (f *UserFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+TableName("tasks", ctx)+" "+_alias+" ON "+_alias+"."+"user_id"+" = "+alias+".id")
		err := f.Tasks.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *UserFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}
	whereConditions := []string{}

	if f.ID != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.ID)
	}

	if f.IDNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NULL"+" OR "+aliasPrefix+SnakeString("id")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("id")+" <> ''")
		}
	}

	if f.Phone != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("phone")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("phone"))
		values = append(values, f.Phone)
	}

	if f.PhoneNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("phone")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("phone"))
		values = append(values, f.PhoneNe)
	}

	if f.PhoneGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("phone")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("phone"))
		values = append(values, f.PhoneGt)
	}

	if f.PhoneLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("phone")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("phone"))
		values = append(values, f.PhoneLt)
	}

	if f.PhoneGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("phone")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("phone"))
		values = append(values, f.PhoneGte)
	}

	if f.PhoneLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("phone")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("phone"))
		values = append(values, f.PhoneLte)
	}

	if f.PhoneIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("phone")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("phone"))
		values = append(values, f.PhoneIn)
	}

	if f.PhoneLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("phone")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("phone"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.PhoneLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.PhonePrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("phone")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("phone"))
		values = append(values, fmt.Sprintf("%s%%", *f.PhonePrefix))
	}

	if f.PhoneSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("phone")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("phone"))
		values = append(values, fmt.Sprintf("%%%s", *f.PhoneSuffix))
	}

	if f.PhoneNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("phone")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("phone"))
		if *f.PhoneNull {
			conditions = append(conditions, aliasPrefix+SnakeString("phone")+" IS NULL"+" OR "+aliasPrefix+SnakeString("phone")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("phone")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("phone")+" <> ''")
		}
	}

	if f.Password != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("password")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("password"))
		values = append(values, f.Password)
	}

	if f.PasswordNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("password")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("password"))
		values = append(values, f.PasswordNe)
	}

	if f.PasswordGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("password")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("password"))
		values = append(values, f.PasswordGt)
	}

	if f.PasswordLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("password")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("password"))
		values = append(values, f.PasswordLt)
	}

	if f.PasswordGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("password")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("password"))
		values = append(values, f.PasswordGte)
	}

	if f.PasswordLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("password")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("password"))
		values = append(values, f.PasswordLte)
	}

	if f.PasswordIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("password")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("password"))
		values = append(values, f.PasswordIn)
	}

	if f.PasswordLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("password")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("password"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.PasswordLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.PasswordPrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("password")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("password"))
		values = append(values, fmt.Sprintf("%s%%", *f.PasswordPrefix))
	}

	if f.PasswordSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("password")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("password"))
		values = append(values, fmt.Sprintf("%%%s", *f.PasswordSuffix))
	}

	if f.PasswordNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("password")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("password"))
		if *f.PasswordNull {
			conditions = append(conditions, aliasPrefix+SnakeString("password")+" IS NULL"+" OR "+aliasPrefix+SnakeString("password")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("password")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("password")+" <> ''")
		}
	}

	if f.Email != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("email")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("email"))
		values = append(values, f.Email)
	}

	if f.EmailNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("email")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("email"))
		values = append(values, f.EmailNe)
	}

	if f.EmailGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("email")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("email"))
		values = append(values, f.EmailGt)
	}

	if f.EmailLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("email")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("email"))
		values = append(values, f.EmailLt)
	}

	if f.EmailGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("email")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("email"))
		values = append(values, f.EmailGte)
	}

	if f.EmailLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("email")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("email"))
		values = append(values, f.EmailLte)
	}

	if f.EmailIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("email")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("email"))
		values = append(values, f.EmailIn)
	}

	if f.EmailLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("email")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("email"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.EmailLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.EmailPrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("email")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("email"))
		values = append(values, fmt.Sprintf("%s%%", *f.EmailPrefix))
	}

	if f.EmailSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("email")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("email"))
		values = append(values, fmt.Sprintf("%%%s", *f.EmailSuffix))
	}

	if f.EmailNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("email")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("email"))
		if *f.EmailNull {
			conditions = append(conditions, aliasPrefix+SnakeString("email")+" IS NULL"+" OR "+aliasPrefix+SnakeString("email")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("email")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("email")+" <> ''")
		}
	}

	if f.Nickname != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("nickname")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("nickname"))
		values = append(values, f.Nickname)
	}

	if f.NicknameNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("nickname")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("nickname"))
		values = append(values, f.NicknameNe)
	}

	if f.NicknameGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("nickname")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("nickname"))
		values = append(values, f.NicknameGt)
	}

	if f.NicknameLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("nickname")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("nickname"))
		values = append(values, f.NicknameLt)
	}

	if f.NicknameGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("nickname")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("nickname"))
		values = append(values, f.NicknameGte)
	}

	if f.NicknameLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("nickname")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("nickname"))
		values = append(values, f.NicknameLte)
	}

	if f.NicknameIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("nickname")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("nickname"))
		values = append(values, f.NicknameIn)
	}

	if f.NicknameLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("nickname")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("nickname"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.NicknameLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.NicknamePrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("nickname")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("nickname"))
		values = append(values, fmt.Sprintf("%s%%", *f.NicknamePrefix))
	}

	if f.NicknameSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("nickname")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("nickname"))
		values = append(values, fmt.Sprintf("%%%s", *f.NicknameSuffix))
	}

	if f.NicknameNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("nickname")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("nickname"))
		if *f.NicknameNull {
			conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" IS NULL"+" OR "+aliasPrefix+SnakeString("nickname")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("nickname")+" <> ''")
		}
	}

	if f.Age != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("age")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("age"))
		values = append(values, f.Age)
	}

	if f.AgeNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("age")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("age"))
		values = append(values, f.AgeNe)
	}

	if f.AgeGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("age")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("age"))
		values = append(values, f.AgeGt)
	}

	if f.AgeLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("age")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("age"))
		values = append(values, f.AgeLt)
	}

	if f.AgeGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("age")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("age"))
		values = append(values, f.AgeGte)
	}

	if f.AgeLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("age")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("age"))
		values = append(values, f.AgeLte)
	}

	if f.AgeIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("age")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("age"))
		values = append(values, f.AgeIn)
	}

	if f.AgeNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("age")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("age"))
		if *f.AgeNull {
			conditions = append(conditions, aliasPrefix+SnakeString("age")+" IS NULL"+" OR "+aliasPrefix+SnakeString("age")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("age")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("age")+" <> ''")
		}
	}

	if f.LastName != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("lastName")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("lastName"))
		values = append(values, f.LastName)
	}

	if f.LastNameNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("lastName")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("lastName"))
		values = append(values, f.LastNameNe)
	}

	if f.LastNameGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("lastName")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("lastName"))
		values = append(values, f.LastNameGt)
	}

	if f.LastNameLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("lastName")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("lastName"))
		values = append(values, f.LastNameLt)
	}

	if f.LastNameGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("lastName")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("lastName"))
		values = append(values, f.LastNameGte)
	}

	if f.LastNameLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("lastName")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("lastName"))
		values = append(values, f.LastNameLte)
	}

	if f.LastNameIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("lastName")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("lastName"))
		values = append(values, f.LastNameIn)
	}

	if f.LastNameLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("lastName")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("lastName"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.LastNameLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.LastNamePrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("lastName")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("lastName"))
		values = append(values, fmt.Sprintf("%s%%", *f.LastNamePrefix))
	}

	if f.LastNameSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("lastName")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("lastName"))
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameSuffix))
	}

	if f.LastNameNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("lastName")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("lastName"))
		if *f.LastNameNull {
			conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" IS NULL"+" OR "+aliasPrefix+SnakeString("lastName")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("lastName")+" <> ''")
		}
	}

	if f.IsDelete != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDelete)
	}

	if f.IsDeleteNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteNe)
	}

	if f.IsDeleteGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteGt)
	}

	if f.IsDeleteLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteLt)
	}

	if f.IsDeleteGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteGte)
	}

	if f.IsDeleteLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteLte)
	}

	if f.IsDeleteIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteIn)
	}

	if f.IsDeleteNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		if *f.IsDeleteNull {
			conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IS NULL"+" OR "+aliasPrefix+SnakeString("isDelete")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("isDelete")+" <> ''")
		}
	}

	if f.Weight != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.Weight)
	}

	if f.WeightNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightNe)
	}

	if f.WeightGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightGt)
	}

	if f.WeightLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightLt)
	}

	if f.WeightGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightGte)
	}

	if f.WeightLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightLte)
	}

	if f.WeightIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightIn)
	}

	if f.WeightNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		if *f.WeightNull {
			conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IS NULL"+" OR "+aliasPrefix+SnakeString("weight")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("weight")+" <> ''")
		}
	}

	if f.State != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.State)
	}

	if f.StateNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateNe)
	}

	if f.StateGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateGt)
	}

	if f.StateLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateLt)
	}

	if f.StateGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateGte)
	}

	if f.StateLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateLte)
	}

	if f.StateIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateIn)
	}

	if f.StateNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		if *f.StateNull {
			conditions = append(conditions, aliasPrefix+SnakeString("state")+" IS NULL"+" OR "+aliasPrefix+SnakeString("state")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("state")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("state")+" <> ''")
		}
	}

	if f.DeletedBy != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedBy)
	}

	if f.DeletedByNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByNe)
	}

	if f.DeletedByGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByGt)
	}

	if f.DeletedByLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByLt)
	}

	if f.DeletedByGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByGte)
	}

	if f.DeletedByLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByLte)
	}

	if f.DeletedByIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByIn)
	}

	if f.DeletedByNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		if *f.DeletedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" <> ''")
		}
	}

	if f.UpdatedBy != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" <> ''")
		}
	}

	if f.CreatedBy != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" <> ''")
		}
	}

	if f.DeletedAt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAt)
	}

	if f.DeletedAtNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtNe)
	}

	if f.DeletedAtGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtGt)
	}

	if f.DeletedAtLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtLt)
	}

	if f.DeletedAtGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtGte)
	}

	if f.DeletedAtLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtLte)
	}

	if f.DeletedAtIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtIn)
	}

	if f.DeletedAtNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		if *f.DeletedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" <> ''")
		}
	}

	if f.UpdatedAt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" <> ''")
		}
	}

	if f.CreatedAt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" <> ''")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *UserFilterType) AndWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := CompactNils(f2)
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *UserFilterType) OrWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := CompactNils(f2)
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		Or: append(_f2, f),
	}
}

func (f *TaskFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *TaskFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("tasks", ctx), wheres, values, joins)
}
func (f *TaskFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := alias + "."

	_where, _values := f.WhereContent(aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_cs := []string{}
			err := or.ApplyWithAlias(ctx, alias, &_cs, &vs, &js)
			if err != nil {
				return err
			}
			cs = append(cs, strings.Join(_cs, " AND "))
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.User != nil {
		_alias := alias + "_user"
		*joins = append(*joins, "LEFT JOIN "+TableName("users", ctx)+" "+_alias+" ON "+_alias+".id = "+alias+"."+"user_id")
		err := f.User.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *TaskFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}
	whereConditions := []string{}

	if f.ID != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.ID)
	}

	if f.IDNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("id")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("id"))
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NULL"+" OR "+aliasPrefix+SnakeString("id")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("id")+" <> ''")
		}
	}

	if f.Title != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("title")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("title"))
		values = append(values, f.Title)
	}

	if f.TitleNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("title")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("title"))
		values = append(values, f.TitleNe)
	}

	if f.TitleGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("title")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("title"))
		values = append(values, f.TitleGt)
	}

	if f.TitleLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("title")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("title"))
		values = append(values, f.TitleLt)
	}

	if f.TitleGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("title")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("title"))
		values = append(values, f.TitleGte)
	}

	if f.TitleLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("title")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("title"))
		values = append(values, f.TitleLte)
	}

	if f.TitleIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("title")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("title"))
		values = append(values, f.TitleIn)
	}

	if f.TitleLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("title")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("title"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.TitleLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.TitlePrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("title")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("title"))
		values = append(values, fmt.Sprintf("%s%%", *f.TitlePrefix))
	}

	if f.TitleSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("title")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("title"))
		values = append(values, fmt.Sprintf("%%%s", *f.TitleSuffix))
	}

	if f.TitleNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("title")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("title"))
		if *f.TitleNull {
			conditions = append(conditions, aliasPrefix+SnakeString("title")+" IS NULL"+" OR "+aliasPrefix+SnakeString("title")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("title")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("title")+" <> ''")
		}
	}

	if f.Completed != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("completed")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("completed"))
		values = append(values, f.Completed)
	}

	if f.CompletedNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("completed")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("completed"))
		values = append(values, f.CompletedNe)
	}

	if f.CompletedGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("completed")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("completed"))
		values = append(values, f.CompletedGt)
	}

	if f.CompletedLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("completed")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("completed"))
		values = append(values, f.CompletedLt)
	}

	if f.CompletedGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("completed")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("completed"))
		values = append(values, f.CompletedGte)
	}

	if f.CompletedLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("completed")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("completed"))
		values = append(values, f.CompletedLte)
	}

	if f.CompletedIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("completed")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("completed"))
		values = append(values, f.CompletedIn)
	}

	if f.CompletedNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("completed")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("completed"))
		if *f.CompletedNull {
			conditions = append(conditions, aliasPrefix+SnakeString("completed")+" IS NULL"+" OR "+aliasPrefix+SnakeString("completed")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("completed")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("completed")+" <> ''")
		}
	}

	if f.DueDate != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("dueDate")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("dueDate"))
		values = append(values, f.DueDate)
	}

	if f.DueDateNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("dueDate")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("dueDate"))
		values = append(values, f.DueDateNe)
	}

	if f.DueDateGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("dueDate")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("dueDate"))
		values = append(values, f.DueDateGt)
	}

	if f.DueDateLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("dueDate")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("dueDate"))
		values = append(values, f.DueDateLt)
	}

	if f.DueDateGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("dueDate")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("dueDate"))
		values = append(values, f.DueDateGte)
	}

	if f.DueDateLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("dueDate")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("dueDate"))
		values = append(values, f.DueDateLte)
	}

	if f.DueDateIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("dueDate")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("dueDate"))
		values = append(values, f.DueDateIn)
	}

	if f.DueDateNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("dueDate")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("dueDate"))
		if *f.DueDateNull {
			conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" IS NULL"+" OR "+aliasPrefix+SnakeString("dueDate")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("dueDate")+" <> ''")
		}
	}

	if f.UserID != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("userId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("userId")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("userId"))
		values = append(values, f.UserID)
	}

	if f.UserIDNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("userId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("userId")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("userId"))
		values = append(values, f.UserIDNe)
	}

	if f.UserIDGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("userId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("userId")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("userId"))
		values = append(values, f.UserIDGt)
	}

	if f.UserIDLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("userId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("userId")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("userId"))
		values = append(values, f.UserIDLt)
	}

	if f.UserIDGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("userId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("userId")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("userId"))
		values = append(values, f.UserIDGte)
	}

	if f.UserIDLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("userId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("userId")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("userId"))
		values = append(values, f.UserIDLte)
	}

	if f.UserIDIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("userId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("userId")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("userId"))
		values = append(values, f.UserIDIn)
	}

	if f.UserIDNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("userId")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("userId"))
		if *f.UserIDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("userId")+" IS NULL"+" OR "+aliasPrefix+SnakeString("userId")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("userId")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("userId")+" <> ''")
		}
	}

	if f.IsDelete != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDelete)
	}

	if f.IsDeleteNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteNe)
	}

	if f.IsDeleteGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteGt)
	}

	if f.IsDeleteLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteLt)
	}

	if f.IsDeleteGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteGte)
	}

	if f.IsDeleteLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteLte)
	}

	if f.IsDeleteIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		values = append(values, f.IsDeleteIn)
	}

	if f.IsDeleteNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("isDelete")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("isDelete"))
		if *f.IsDeleteNull {
			conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IS NULL"+" OR "+aliasPrefix+SnakeString("isDelete")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("isDelete")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("isDelete")+" <> ''")
		}
	}

	if f.Weight != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.Weight)
	}

	if f.WeightNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightNe)
	}

	if f.WeightGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightGt)
	}

	if f.WeightLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightLt)
	}

	if f.WeightGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightGte)
	}

	if f.WeightLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightLte)
	}

	if f.WeightIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		values = append(values, f.WeightIn)
	}

	if f.WeightNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("weight")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("weight"))
		if *f.WeightNull {
			conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IS NULL"+" OR "+aliasPrefix+SnakeString("weight")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("weight")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("weight")+" <> ''")
		}
	}

	if f.State != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.State)
	}

	if f.StateNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateNe)
	}

	if f.StateGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateGt)
	}

	if f.StateLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateLt)
	}

	if f.StateGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateGte)
	}

	if f.StateLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateLte)
	}

	if f.StateIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("state")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		values = append(values, f.StateIn)
	}

	if f.StateNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("state")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("state"))
		if *f.StateNull {
			conditions = append(conditions, aliasPrefix+SnakeString("state")+" IS NULL"+" OR "+aliasPrefix+SnakeString("state")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("state")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("state")+" <> ''")
		}
	}

	if f.DeletedBy != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedBy)
	}

	if f.DeletedByNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByNe)
	}

	if f.DeletedByGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByGt)
	}

	if f.DeletedByLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByLt)
	}

	if f.DeletedByGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByGte)
	}

	if f.DeletedByLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByLte)
	}

	if f.DeletedByIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		values = append(values, f.DeletedByIn)
	}

	if f.DeletedByNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedBy")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedBy"))
		if *f.DeletedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" <> ''")
		}
	}

	if f.UpdatedBy != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedBy")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedBy"))
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" <> ''")
		}
	}

	if f.CreatedBy != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdBy")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdBy"))
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" <> ''")
		}
	}

	if f.DeletedAt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAt)
	}

	if f.DeletedAtNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtNe)
	}

	if f.DeletedAtGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtGt)
	}

	if f.DeletedAtLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtLt)
	}

	if f.DeletedAtGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtGte)
	}

	if f.DeletedAtLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtLte)
	}

	if f.DeletedAtIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		values = append(values, f.DeletedAtIn)
	}

	if f.DeletedAtNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("deletedAt")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("deletedAt"))
		if *f.DeletedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" <> ''")
		}
	}

	if f.UpdatedAt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("updatedAt")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("updatedAt"))
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" <> ''")
		}
	}

	if f.CreatedAt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("createdAt")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("createdAt"))
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" <> ''")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *TaskFilterType) AndWith(f2 ...*TaskFilterType) *TaskFilterType {
	_f2 := CompactNils(f2)
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *TaskFilterType) OrWith(f2 ...*TaskFilterType) *TaskFilterType {
	_f2 := CompactNils(f2)
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		Or: append(_f2, f),
	}
}
