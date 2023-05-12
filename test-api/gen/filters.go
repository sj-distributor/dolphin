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
	return f.ApplyWithAlias(ctx, TableName("users"), wheres, values, joins)
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
		*joins = append(*joins, "LEFT JOIN "+"tasks"+" "+_alias+" ON "+_alias+"."+"assignee_id"+" = "+alias+".id")
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

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NULL"+" OR "+aliasPrefix+SnakeString("id")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("id")+" <> ''")
		}
	}

	if f.Phone != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" = ?")
		values = append(values, f.Phone)
	}

	if f.PhoneNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" != ?")
		values = append(values, f.PhoneNe)
	}

	if f.PhoneGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" > ?")
		values = append(values, f.PhoneGt)
	}

	if f.PhoneLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" < ?")
		values = append(values, f.PhoneLt)
	}

	if f.PhoneGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" >= ?")
		values = append(values, f.PhoneGte)
	}

	if f.PhoneLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" <= ?")
		values = append(values, f.PhoneLte)
	}

	if f.PhoneIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" IN (?)")
		values = append(values, f.PhoneIn)
	}

	if f.PhoneLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.PhoneLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.PhonePrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PhonePrefix))
	}

	if f.PhoneSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("phone")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PhoneSuffix))
	}

	if f.PhoneNull != nil {
		if *f.PhoneNull {
			conditions = append(conditions, aliasPrefix+SnakeString("phone")+" IS NULL"+" OR "+aliasPrefix+SnakeString("phone")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("phone")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("phone")+" <> ''")
		}
	}

	if f.Password != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" = ?")
		values = append(values, f.Password)
	}

	if f.PasswordNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" != ?")
		values = append(values, f.PasswordNe)
	}

	if f.PasswordGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" > ?")
		values = append(values, f.PasswordGt)
	}

	if f.PasswordLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" < ?")
		values = append(values, f.PasswordLt)
	}

	if f.PasswordGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" >= ?")
		values = append(values, f.PasswordGte)
	}

	if f.PasswordLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" <= ?")
		values = append(values, f.PasswordLte)
	}

	if f.PasswordIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" IN (?)")
		values = append(values, f.PasswordIn)
	}

	if f.PasswordLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.PasswordLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.PasswordPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PasswordPrefix))
	}

	if f.PasswordSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("password")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PasswordSuffix))
	}

	if f.PasswordNull != nil {
		if *f.PasswordNull {
			conditions = append(conditions, aliasPrefix+SnakeString("password")+" IS NULL"+" OR "+aliasPrefix+SnakeString("password")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("password")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("password")+" <> ''")
		}
	}

	if f.Email != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" = ?")
		values = append(values, f.Email)
	}

	if f.EmailNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" != ?")
		values = append(values, f.EmailNe)
	}

	if f.EmailGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" > ?")
		values = append(values, f.EmailGt)
	}

	if f.EmailLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" < ?")
		values = append(values, f.EmailLt)
	}

	if f.EmailGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" >= ?")
		values = append(values, f.EmailGte)
	}

	if f.EmailLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" <= ?")
		values = append(values, f.EmailLte)
	}

	if f.EmailIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" IN (?)")
		values = append(values, f.EmailIn)
	}

	if f.EmailLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.EmailLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.EmailPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailPrefix))
	}

	if f.EmailSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailSuffix))
	}

	if f.EmailNull != nil {
		if *f.EmailNull {
			conditions = append(conditions, aliasPrefix+SnakeString("email")+" IS NULL"+" OR "+aliasPrefix+SnakeString("email")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("email")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("email")+" <> ''")
		}
	}

	if f.Nickname != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" = ?")
		values = append(values, f.Nickname)
	}

	if f.NicknameNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" != ?")
		values = append(values, f.NicknameNe)
	}

	if f.NicknameGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" > ?")
		values = append(values, f.NicknameGt)
	}

	if f.NicknameLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" < ?")
		values = append(values, f.NicknameLt)
	}

	if f.NicknameGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" >= ?")
		values = append(values, f.NicknameGte)
	}

	if f.NicknameLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" <= ?")
		values = append(values, f.NicknameLte)
	}

	if f.NicknameIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" IN (?)")
		values = append(values, f.NicknameIn)
	}

	if f.NicknameLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.NicknameLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.NicknamePrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NicknamePrefix))
	}

	if f.NicknameSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NicknameSuffix))
	}

	if f.NicknameNull != nil {
		if *f.NicknameNull {
			conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" IS NULL"+" OR "+aliasPrefix+SnakeString("nickname")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("nickname")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("nickname")+" <> ''")
		}
	}

	if f.Age != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" = ?")
		values = append(values, f.Age)
	}

	if f.AgeNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" != ?")
		values = append(values, f.AgeNe)
	}

	if f.AgeGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" > ?")
		values = append(values, f.AgeGt)
	}

	if f.AgeLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" < ?")
		values = append(values, f.AgeLt)
	}

	if f.AgeGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" >= ?")
		values = append(values, f.AgeGte)
	}

	if f.AgeLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" <= ?")
		values = append(values, f.AgeLte)
	}

	if f.AgeIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("age")+" IN (?)")
		values = append(values, f.AgeIn)
	}

	if f.AgeNull != nil {
		if *f.AgeNull {
			conditions = append(conditions, aliasPrefix+SnakeString("age")+" IS NULL"+" OR "+aliasPrefix+SnakeString("age")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("age")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("age")+" <> ''")
		}
	}

	if f.LastName != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" = ?")
		values = append(values, f.LastName)
	}

	if f.LastNameNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" != ?")
		values = append(values, f.LastNameNe)
	}

	if f.LastNameGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" > ?")
		values = append(values, f.LastNameGt)
	}

	if f.LastNameLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" < ?")
		values = append(values, f.LastNameLt)
	}

	if f.LastNameGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" >= ?")
		values = append(values, f.LastNameGte)
	}

	if f.LastNameLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" <= ?")
		values = append(values, f.LastNameLte)
	}

	if f.LastNameIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" IN (?)")
		values = append(values, f.LastNameIn)
	}

	if f.LastNameLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.LastNameLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.LastNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNamePrefix))
	}

	if f.LastNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameSuffix))
	}

	if f.LastNameNull != nil {
		if *f.LastNameNull {
			conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" IS NULL"+" OR "+aliasPrefix+SnakeString("lastName")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("lastName")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("lastName")+" <> ''")
		}
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}

	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}

	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}

	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}

	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}

	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}

	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.DeletedByNull != nil {
		if *f.DeletedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" <> ''")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" <> ''")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" <> ''")
		}
	}

	if f.DeletedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" = ?")
		values = append(values, f.DeletedAt)
	}

	if f.DeletedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" != ?")
		values = append(values, f.DeletedAtNe)
	}

	if f.DeletedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" > ?")
		values = append(values, f.DeletedAtGt)
	}

	if f.DeletedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" < ?")
		values = append(values, f.DeletedAtLt)
	}

	if f.DeletedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" >= ?")
		values = append(values, f.DeletedAtGte)
	}

	if f.DeletedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" <= ?")
		values = append(values, f.DeletedAtLte)
	}

	if f.DeletedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IN (?)")
		values = append(values, f.DeletedAtIn)
	}

	if f.DeletedAtNull != nil {
		if *f.DeletedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" <> ''")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" <> ''")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
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
	return f.ApplyWithAlias(ctx, TableName("tasks"), wheres, values, joins)
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

	if f.Assignee != nil {
		_alias := alias + "_assignee"
		*joins = append(*joins, "LEFT JOIN "+"users"+" "+_alias+" ON "+_alias+".id = "+alias+"."+"assignee_id")
		err := f.Assignee.ApplyWithAlias(ctx, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *TaskFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NULL"+" OR "+aliasPrefix+SnakeString("id")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("id")+" <> ''")
		}
	}

	if f.Title != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" = ?")
		values = append(values, f.Title)
	}

	if f.TitleNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" != ?")
		values = append(values, f.TitleNe)
	}

	if f.TitleGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" > ?")
		values = append(values, f.TitleGt)
	}

	if f.TitleLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" < ?")
		values = append(values, f.TitleLt)
	}

	if f.TitleGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" >= ?")
		values = append(values, f.TitleGte)
	}

	if f.TitleLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" <= ?")
		values = append(values, f.TitleLte)
	}

	if f.TitleIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" IN (?)")
		values = append(values, f.TitleIn)
	}

	if f.TitleLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.TitleLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.TitlePrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TitlePrefix))
	}

	if f.TitleSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("title")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TitleSuffix))
	}

	if f.TitleNull != nil {
		if *f.TitleNull {
			conditions = append(conditions, aliasPrefix+SnakeString("title")+" IS NULL"+" OR "+aliasPrefix+SnakeString("title")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("title")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("title")+" <> ''")
		}
	}

	if f.Completed != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" = ?")
		values = append(values, f.Completed)
	}

	if f.CompletedNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" != ?")
		values = append(values, f.CompletedNe)
	}

	if f.CompletedGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" > ?")
		values = append(values, f.CompletedGt)
	}

	if f.CompletedLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" < ?")
		values = append(values, f.CompletedLt)
	}

	if f.CompletedGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" >= ?")
		values = append(values, f.CompletedGte)
	}

	if f.CompletedLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" <= ?")
		values = append(values, f.CompletedLte)
	}

	if f.CompletedIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("completed")+" IN (?)")
		values = append(values, f.CompletedIn)
	}

	if f.CompletedNull != nil {
		if *f.CompletedNull {
			conditions = append(conditions, aliasPrefix+SnakeString("completed")+" IS NULL"+" OR "+aliasPrefix+SnakeString("completed")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("completed")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("completed")+" <> ''")
		}
	}

	if f.DueDate != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" = ?")
		values = append(values, f.DueDate)
	}

	if f.DueDateNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" != ?")
		values = append(values, f.DueDateNe)
	}

	if f.DueDateGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" > ?")
		values = append(values, f.DueDateGt)
	}

	if f.DueDateLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" < ?")
		values = append(values, f.DueDateLt)
	}

	if f.DueDateGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" >= ?")
		values = append(values, f.DueDateGte)
	}

	if f.DueDateLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" <= ?")
		values = append(values, f.DueDateLte)
	}

	if f.DueDateIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" IN (?)")
		values = append(values, f.DueDateIn)
	}

	if f.DueDateNull != nil {
		if *f.DueDateNull {
			conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" IS NULL"+" OR "+aliasPrefix+SnakeString("dueDate")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("dueDate")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("dueDate")+" <> ''")
		}
	}

	if f.AssigneeID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("assigneeId")+" = ?")
		values = append(values, f.AssigneeID)
	}

	if f.AssigneeIDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("assigneeId")+" != ?")
		values = append(values, f.AssigneeIDNe)
	}

	if f.AssigneeIDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("assigneeId")+" > ?")
		values = append(values, f.AssigneeIDGt)
	}

	if f.AssigneeIDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("assigneeId")+" < ?")
		values = append(values, f.AssigneeIDLt)
	}

	if f.AssigneeIDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("assigneeId")+" >= ?")
		values = append(values, f.AssigneeIDGte)
	}

	if f.AssigneeIDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("assigneeId")+" <= ?")
		values = append(values, f.AssigneeIDLte)
	}

	if f.AssigneeIDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("assigneeId")+" IN (?)")
		values = append(values, f.AssigneeIDIn)
	}

	if f.AssigneeIDNull != nil {
		if *f.AssigneeIDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("assigneeId")+" IS NULL"+" OR "+aliasPrefix+SnakeString("assigneeId")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("assigneeId")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("assigneeId")+" <> ''")
		}
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}

	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}

	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}

	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}

	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}

	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}

	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.DeletedByNull != nil {
		if *f.DeletedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" <> ''")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" <> ''")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" <> ''")
		}
	}

	if f.DeletedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" = ?")
		values = append(values, f.DeletedAt)
	}

	if f.DeletedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" != ?")
		values = append(values, f.DeletedAtNe)
	}

	if f.DeletedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" > ?")
		values = append(values, f.DeletedAtGt)
	}

	if f.DeletedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" < ?")
		values = append(values, f.DeletedAtLt)
	}

	if f.DeletedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" >= ?")
		values = append(values, f.DeletedAtGte)
	}

	if f.DeletedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" <= ?")
		values = append(values, f.DeletedAtLte)
	}

	if f.DeletedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IN (?)")
		values = append(values, f.DeletedAtIn)
	}

	if f.DeletedAtNull != nil {
		if *f.DeletedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" <> ''")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" <> ''")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
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

func (f *UploadFileFilterType) IsEmpty(ctx context.Context) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *UploadFileFilterType) Apply(ctx context.Context, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, TableName("upload_files"), wheres, values, joins)
}
func (f *UploadFileFilterType) ApplyWithAlias(ctx context.Context, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
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

	return nil
}

func (f *UploadFileFilterType) WhereContent(aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NULL"+" OR "+aliasPrefix+SnakeString("id")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("id")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("id")+" <> ''")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+SnakeString("name")+" IS NULL"+" OR "+aliasPrefix+SnakeString("name")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("name")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("name")+" <> ''")
		}
	}

	if f.Hash != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("hash")+" = ?")
		values = append(values, f.Hash)
	}

	if f.HashNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("hash")+" != ?")
		values = append(values, f.HashNe)
	}

	if f.HashGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("hash")+" > ?")
		values = append(values, f.HashGt)
	}

	if f.HashLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("hash")+" < ?")
		values = append(values, f.HashLt)
	}

	if f.HashGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("hash")+" >= ?")
		values = append(values, f.HashGte)
	}

	if f.HashLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("hash")+" <= ?")
		values = append(values, f.HashLte)
	}

	if f.HashIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("hash")+" IN (?)")
		values = append(values, f.HashIn)
	}

	if f.HashLike != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("hash")+" LIKE ?")
		values = append(values, "%"+strings.Replace(strings.Replace(*f.HashLike, "?", "_", -1), "*", "%", -1)+"%")
	}

	if f.HashPrefix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("hash")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.HashPrefix))
	}

	if f.HashSuffix != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("hash")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.HashSuffix))
	}

	if f.HashNull != nil {
		if *f.HashNull {
			conditions = append(conditions, aliasPrefix+SnakeString("hash")+" IS NULL"+" OR "+aliasPrefix+SnakeString("hash")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("hash")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("hash")+" <> ''")
		}
	}

	if f.DeletedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" = ?")
		values = append(values, f.DeletedBy)
	}

	if f.DeletedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" != ?")
		values = append(values, f.DeletedByNe)
	}

	if f.DeletedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" > ?")
		values = append(values, f.DeletedByGt)
	}

	if f.DeletedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" < ?")
		values = append(values, f.DeletedByLt)
	}

	if f.DeletedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" >= ?")
		values = append(values, f.DeletedByGte)
	}

	if f.DeletedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" <= ?")
		values = append(values, f.DeletedByLte)
	}

	if f.DeletedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IN (?)")
		values = append(values, f.DeletedByIn)
	}

	if f.DeletedByNull != nil {
		if *f.DeletedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedBy")+" <> ''")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedBy")+" <> ''")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdBy")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdBy")+" <> ''")
		}
	}

	if f.DeletedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" = ?")
		values = append(values, f.DeletedAt)
	}

	if f.DeletedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" != ?")
		values = append(values, f.DeletedAtNe)
	}

	if f.DeletedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" > ?")
		values = append(values, f.DeletedAtGt)
	}

	if f.DeletedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" < ?")
		values = append(values, f.DeletedAtLt)
	}

	if f.DeletedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" >= ?")
		values = append(values, f.DeletedAtGte)
	}

	if f.DeletedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" <= ?")
		values = append(values, f.DeletedAtLte)
	}

	if f.DeletedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IN (?)")
		values = append(values, f.DeletedAtIn)
	}

	if f.DeletedAtNull != nil {
		if *f.DeletedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("deletedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("deletedAt")+" <> ''")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("updatedAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("updatedAt")+" <> ''")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" =''")
		} else {
			conditions = append(conditions, aliasPrefix+SnakeString("createdAt")+" IS NOT NULL"+" OR "+aliasPrefix+SnakeString("createdAt")+" <> ''")
		}
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *UploadFileFilterType) AndWith(f2 ...*UploadFileFilterType) *UploadFileFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UploadFileFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *UploadFileFilterType) OrWith(f2 ...*UploadFileFilterType) *UploadFileFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UploadFileFilterType{
		Or: append(_f2, f),
	}
}
