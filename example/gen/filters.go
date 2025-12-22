package gen

import (
	"context"
	"fmt"
	"strings"
)

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

	if f.Profile != nil {
		_alias := alias + "_profile"
		*joins = append(*joins, "LEFT JOIN "+TableName("profiles", ctx)+" "+_alias+" ON "+_alias+".id = "+alias+"."+"profile_id")
		err := f.Profile.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+TableName("tasks", ctx)+" "+_alias+" ON "+_alias+"."+"user_id"+" = "+alias+".id")
		err := f.Tasks.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	if f.UserRoles != nil {
		_alias := alias + "_userRoles"
		*joins = append(*joins, "LEFT JOIN "+TableName("userRole_users", ctx)+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"user_id"+" LEFT JOIN "+TableName("user_roles", ctx)+" "+_alias+" ON "+_alias+"_jointable"+"."+"user_role_id"+" = "+_alias+".id")
		err := f.UserRoles.ApplyWithAlias(ctx, _alias, wheres, values, joins)
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

	if f.ProfileID != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("profileId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("profileId")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("profileId"))
		values = append(values, f.ProfileID)
	}

	if f.ProfileIDNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("profileId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("profileId")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("profileId"))
		values = append(values, f.ProfileIDNe)
	}

	if f.ProfileIDGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("profileId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("profileId")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("profileId"))
		values = append(values, f.ProfileIDGt)
	}

	if f.ProfileIDLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("profileId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("profileId")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("profileId"))
		values = append(values, f.ProfileIDLt)
	}

	if f.ProfileIDGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("profileId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("profileId")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("profileId"))
		values = append(values, f.ProfileIDGte)
	}

	if f.ProfileIDLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("profileId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("profileId")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("profileId"))
		values = append(values, f.ProfileIDLte)
	}

	if f.ProfileIDIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("profileId")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("profileId")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("profileId"))
		values = append(values, f.ProfileIDIn)
	}

	if f.ProfileIDNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("profileId")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("profileId"))
		if *f.ProfileIDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("profileId")+" IS NULL"+" OR "+aliasPrefix+SnakeString("profileId")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("profileId")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("profileId")+" <> ''")
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
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *UserFilterType) OrWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		Or: append(_f2, f),
	}
}

func (f *ProfileFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ProfileFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("profiles", ctx), wheres, values, joins)
}
func (f *ProfileFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
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

func (f *ProfileFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
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

	if f.Avatar != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("avatar")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("avatar"))
		values = append(values, f.Avatar)
	}

	if f.AvatarNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("avatar")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("avatar"))
		values = append(values, f.AvatarNe)
	}

	if f.AvatarGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("avatar")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("avatar"))
		values = append(values, f.AvatarGt)
	}

	if f.AvatarLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("avatar")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("avatar"))
		values = append(values, f.AvatarLt)
	}

	if f.AvatarGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("avatar")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("avatar"))
		values = append(values, f.AvatarGte)
	}

	if f.AvatarLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("avatar")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("avatar"))
		values = append(values, f.AvatarLte)
	}

	if f.AvatarIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("avatar")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("avatar"))
		values = append(values, f.AvatarIn)
	}

	if f.AvatarLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("avatar")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("avatar"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.AvatarLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.AvatarPrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("avatar")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("avatar"))
		values = append(values, fmt.Sprintf("%s%%", *f.AvatarPrefix))
	}

	if f.AvatarSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("avatar")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("avatar"))
		values = append(values, fmt.Sprintf("%%%s", *f.AvatarSuffix))
	}

	if f.AvatarNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("avatar")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("avatar"))
		if *f.AvatarNull {
			conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" IS NULL"+" OR "+aliasPrefix+SnakeString("avatar")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("avatar")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("avatar")+" <> ''")
		}
	}

	if f.Bio != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("bio")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("bio")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("bio"))
		values = append(values, f.Bio)
	}

	if f.BioNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("bio")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("bio")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("bio"))
		values = append(values, f.BioNe)
	}

	if f.BioGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("bio")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("bio")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("bio"))
		values = append(values, f.BioGt)
	}

	if f.BioLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("bio")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("bio")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("bio"))
		values = append(values, f.BioLt)
	}

	if f.BioGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("bio")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("bio")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("bio"))
		values = append(values, f.BioGte)
	}

	if f.BioLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("bio")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("bio")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("bio"))
		values = append(values, f.BioLte)
	}

	if f.BioIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("bio")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("bio")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("bio"))
		values = append(values, f.BioIn)
	}

	if f.BioLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("bio")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("bio")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("bio"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.BioLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.BioPrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("bio")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("bio")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("bio"))
		values = append(values, fmt.Sprintf("%s%%", *f.BioPrefix))
	}

	if f.BioSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("bio")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("bio")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("bio"))
		values = append(values, fmt.Sprintf("%%%s", *f.BioSuffix))
	}

	if f.BioNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("bio")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("bio"))
		if *f.BioNull {
			conditions = append(conditions, aliasPrefix+SnakeString("bio")+" IS NULL"+" OR "+aliasPrefix+SnakeString("bio")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("bio")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("bio")+" <> ''")
		}
	}

	if f.Birthday != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("birthday")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("birthday")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("birthday"))
		values = append(values, f.Birthday)
	}

	if f.BirthdayNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("birthday")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("birthday")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("birthday"))
		values = append(values, f.BirthdayNe)
	}

	if f.BirthdayGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("birthday")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("birthday")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("birthday"))
		values = append(values, f.BirthdayGt)
	}

	if f.BirthdayLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("birthday")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("birthday")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("birthday"))
		values = append(values, f.BirthdayLt)
	}

	if f.BirthdayGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("birthday")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("birthday")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("birthday"))
		values = append(values, f.BirthdayGte)
	}

	if f.BirthdayLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("birthday")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("birthday")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("birthday"))
		values = append(values, f.BirthdayLte)
	}

	if f.BirthdayIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("birthday")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("birthday")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("birthday"))
		values = append(values, f.BirthdayIn)
	}

	if f.BirthdayNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("birthday")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("birthday"))
		if *f.BirthdayNull {
			conditions = append(conditions, aliasPrefix+SnakeString("birthday")+" IS NULL"+" OR "+aliasPrefix+SnakeString("birthday")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("birthday")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("birthday")+" <> ''")
		}
	}

	if f.Address != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("address")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("address")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("address"))
		values = append(values, f.Address)
	}

	if f.AddressNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("address")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("address")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("address"))
		values = append(values, f.AddressNe)
	}

	if f.AddressGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("address")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("address")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("address"))
		values = append(values, f.AddressGt)
	}

	if f.AddressLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("address")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("address")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("address"))
		values = append(values, f.AddressLt)
	}

	if f.AddressGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("address")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("address")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("address"))
		values = append(values, f.AddressGte)
	}

	if f.AddressLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("address")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("address")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("address"))
		values = append(values, f.AddressLte)
	}

	if f.AddressIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("address")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("address")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("address"))
		values = append(values, f.AddressIn)
	}

	if f.AddressLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("address")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("address")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("address"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.AddressLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.AddressPrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("address")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("address")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("address"))
		values = append(values, fmt.Sprintf("%s%%", *f.AddressPrefix))
	}

	if f.AddressSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("address")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("address")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("address"))
		values = append(values, fmt.Sprintf("%%%s", *f.AddressSuffix))
	}

	if f.AddressNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("address")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("address"))
		if *f.AddressNull {
			conditions = append(conditions, aliasPrefix+SnakeString("address")+" IS NULL"+" OR "+aliasPrefix+SnakeString("address")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("address")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("address")+" <> ''")
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
func (f *ProfileFilterType) AndWith(f2 ...*ProfileFilterType) *ProfileFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ProfileFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ProfileFilterType) OrWith(f2 ...*ProfileFilterType) *ProfileFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ProfileFilterType{
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

	if f.Tags != nil {
		_alias := alias + "_tags"
		*joins = append(*joins, "LEFT JOIN "+TableName("tag_tasks", ctx)+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"task_id"+" LEFT JOIN "+TableName("tags", ctx)+" "+_alias+" ON "+_alias+"_jointable"+"."+"tag_id"+" = "+_alias+".id")
		err := f.Tags.ApplyWithAlias(ctx, _alias, wheres, values, joins)
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

	if f.Description != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.DescriptionPrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+SnakeString("description")+" IS NULL"+" OR "+aliasPrefix+SnakeString("description")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("description")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("description")+" <> ''")
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

	if f.Priority != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("priority")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("priority")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("priority"))
		values = append(values, f.Priority)
	}

	if f.PriorityNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("priority")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("priority")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("priority"))
		values = append(values, f.PriorityNe)
	}

	if f.PriorityGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("priority")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("priority")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("priority"))
		values = append(values, f.PriorityGt)
	}

	if f.PriorityLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("priority")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("priority")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("priority"))
		values = append(values, f.PriorityLt)
	}

	if f.PriorityGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("priority")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("priority")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("priority"))
		values = append(values, f.PriorityGte)
	}

	if f.PriorityLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("priority")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("priority")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("priority"))
		values = append(values, f.PriorityLte)
	}

	if f.PriorityIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("priority")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("priority")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("priority"))
		values = append(values, f.PriorityIn)
	}

	if f.PriorityNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("priority")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("priority"))
		if *f.PriorityNull {
			conditions = append(conditions, aliasPrefix+SnakeString("priority")+" IS NULL"+" OR "+aliasPrefix+SnakeString("priority")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("priority")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("priority")+" <> ''")
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
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *TaskFilterType) OrWith(f2 ...*TaskFilterType) *TaskFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		Or: append(_f2, f),
	}
}

func (f *UserRoleFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *UserRoleFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("user_roles", ctx), wheres, values, joins)
}
func (f *UserRoleFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
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

	if f.Users != nil {
		_alias := alias + "_users"
		*joins = append(*joins, "LEFT JOIN "+TableName("userRole_users", ctx)+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"userRole_id"+" LEFT JOIN "+TableName("users", ctx)+" "+_alias+" ON "+_alias+"_jointable"+"."+"user_id"+" = "+_alias+".id")
		err := f.Users.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *UserRoleFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
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

	if f.Name != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.Name)
	}

	if f.NameNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.NamePrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+SnakeString("name")+" IS NULL"+" OR "+aliasPrefix+SnakeString("name")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("name")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("name")+" <> ''")
		}
	}

	if f.Description != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.Description)
	}

	if f.DescriptionNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionNe)
	}

	if f.DescriptionGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionGt)
	}

	if f.DescriptionLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionLt)
	}

	if f.DescriptionGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionGte)
	}

	if f.DescriptionLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionLte)
	}

	if f.DescriptionIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, f.DescriptionIn)
	}

	if f.DescriptionLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.DescriptionLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.DescriptionPrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, fmt.Sprintf("%s%%", *f.DescriptionPrefix))
	}

	if f.DescriptionSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("description")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		values = append(values, fmt.Sprintf("%%%s", *f.DescriptionSuffix))
	}

	if f.DescriptionNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("description")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("description"))
		if *f.DescriptionNull {
			conditions = append(conditions, aliasPrefix+SnakeString("description")+" IS NULL"+" OR "+aliasPrefix+SnakeString("description")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("description")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("description")+" <> ''")
		}
	}

	if f.Permissions != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("permissions")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("permissions"))
		values = append(values, f.Permissions)
	}

	if f.PermissionsNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("permissions")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("permissions"))
		values = append(values, f.PermissionsNe)
	}

	if f.PermissionsGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("permissions")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("permissions"))
		values = append(values, f.PermissionsGt)
	}

	if f.PermissionsLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("permissions")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("permissions"))
		values = append(values, f.PermissionsLt)
	}

	if f.PermissionsGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("permissions")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("permissions"))
		values = append(values, f.PermissionsGte)
	}

	if f.PermissionsLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("permissions")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("permissions"))
		values = append(values, f.PermissionsLte)
	}

	if f.PermissionsIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("permissions")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("permissions"))
		values = append(values, f.PermissionsIn)
	}

	if f.PermissionsLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("permissions")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("permissions"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.PermissionsLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.PermissionsPrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("permissions")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("permissions"))
		values = append(values, fmt.Sprintf("%s%%", *f.PermissionsPrefix))
	}

	if f.PermissionsSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("permissions")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("permissions"))
		values = append(values, fmt.Sprintf("%%%s", *f.PermissionsSuffix))
	}

	if f.PermissionsNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("permissions")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("permissions"))
		if *f.PermissionsNull {
			conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" IS NULL"+" OR "+aliasPrefix+SnakeString("permissions")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("permissions")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("permissions")+" <> ''")
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
func (f *UserRoleFilterType) AndWith(f2 ...*UserRoleFilterType) *UserRoleFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserRoleFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *UserRoleFilterType) OrWith(f2 ...*UserRoleFilterType) *UserRoleFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserRoleFilterType{
		Or: append(_f2, f),
	}
}

func (f *TagFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *TagFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("tags", ctx), wheres, values, joins)
}
func (f *TagFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
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
		*joins = append(*joins, "LEFT JOIN "+TableName("tag_tasks", ctx)+" "+_alias+"_jointable"+" ON "+alias+".id = "+_alias+"_jointable"+"."+"tag_id"+" LEFT JOIN "+TableName("tasks", ctx)+" "+_alias+" ON "+_alias+"_jointable"+"."+"task_id"+" = "+_alias+".id")
		err := f.Tasks.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *TagFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
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

	if f.Name != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.Name)
	}

	if f.NameNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.NamePrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("name")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("name"))
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+SnakeString("name")+" IS NULL"+" OR "+aliasPrefix+SnakeString("name")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("name")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("name")+" <> ''")
		}
	}

	if f.Color != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("color")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("color")+" = ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("color"))
		values = append(values, f.Color)
	}

	if f.ColorNe != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("color")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("color")+" != ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("color"))
		values = append(values, f.ColorNe)
	}

	if f.ColorGt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("color")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("color")+" > ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("color"))
		values = append(values, f.ColorGt)
	}

	if f.ColorLt != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("color")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("color")+" < ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("color"))
		values = append(values, f.ColorLt)
	}

	if f.ColorGte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("color")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("color")+" >= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("color"))
		values = append(values, f.ColorGte)
	}

	if f.ColorLte != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("color")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("color")+" <= ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("color"))
		values = append(values, f.ColorLte)
	}

	if f.ColorIn != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("color")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("color")+" IN (?)")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("color"))
		values = append(values, f.ColorIn)
	}

	if f.ColorLike != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("color")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("color")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("color"))
		values = append(values, "%"+strings.Replace(strings.Replace(*f.ColorLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.ColorPrefix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("color")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("color")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("color"))
		values = append(values, fmt.Sprintf("%s%%", *f.ColorPrefix))
	}

	if f.ColorSuffix != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("color")) == -1 {
		conditions = append(conditions, aliasPrefix+SnakeString("color")+" LIKE ?")
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("color"))
		values = append(values, fmt.Sprintf("%%%s", *f.ColorSuffix))
	}

	if f.ColorNull != nil && IndexOf(whereConditions, aliasPrefix+SnakeString("color")) == -1 {
		whereConditions = append(whereConditions, aliasPrefix+SnakeString("color"))
		if *f.ColorNull {
			conditions = append(conditions, aliasPrefix+SnakeString("color")+" IS NULL"+" OR "+aliasPrefix+SnakeString("color")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("color")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("color")+" <> ''")
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
func (f *TagFilterType) AndWith(f2 ...*TagFilterType) *TagFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TagFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *TagFilterType) OrWith(f2 ...*TagFilterType) *TagFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TagFilterType{
		Or: append(_f2, f),
	}
}
