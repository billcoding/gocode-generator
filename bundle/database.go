package bundle

import (
	"fmt"
	. "github.com/billcoding/gocode-generator/config"
	. "github.com/billcoding/gocode-generator/model"
	"strings"
)

func Tables(database string, c *Configuration) []*Table {
	whereSql := ""
	if c.IncludeTables != nil && len(c.IncludeTables) > 0 {
		whereSql = fmt.Sprintf("AND t.`TABLE_NAME` IN('%s')", strings.Join(c.IncludeTables, "','"))
	} else if c.ExcludeTables != nil && len(c.ExcludeTables) > 0 {
		whereSql = fmt.Sprintf("AND t.`TABLE_NAME` NOT IN('%s')", strings.Join(c.ExcludeTables, "','"))
	}
	tableList := SelectTableListSelectMapper.Prepare(map[string]interface{}{
		"DBName": database,
		"Where":  whereSql,
	}).Exec().List(new(Table))
	ts := make([]*Table, len(tableList))
	for i, t := range tableList {
		tt := t.(*Table)
		tt.Columns = make([]*Column, 0)
		tt.Indexes = make([]*Index, 0)
		ts[i] = tt
	}
	return ts
}

func Columns(database string) []*Column {
	columnList := SelectTableColumnListSelectMapper.Prepare(database).Exec().List(new(Column))
	cs := make([]*Column, len(columnList))
	for i, c := range columnList {
		cc := c.(*Column)
		pk := "F"
		insertIgnore := "F"
		updateIgnore := "F"
		autoIncrement := ""
		notNull := "NULL"
		defaultStr := ""
		defaultPre := ""
		defaultSuf := ""
		if cc.AutoIncrement == 1 {
			pk = "T"
			insertIgnore = "T"
			autoIncrement = " auto_increment"
		}
		if strings.Contains(cc.ColumnKey, "PRI") {
			pk = "T"
		}
		if cc.NotNull == 1 {
			notNull = "NOT NULL"
		}
		if cc.Default != "__NULL__" {
			if cc.Default != "CURRENT_TIMESTAMP" {
				defaultPre = "'"
				defaultSuf = "'"
			} else {
				insertIgnore = "T"
				updateIgnore = "T"
			}
			defaultStr = " default " + defaultPre + strings.ToLower(cc.Default) + defaultSuf
		}
		cc.OrmTag = fmt.Sprintf("orm:\"pk{%s} c{%s} ig{%s} ug{%s} def{%s}\"",
			pk, cc.Name, insertIgnore, updateIgnore, fmt.Sprintf("%s %s %s%s%s comment '%s'", cc.Name, cc.DataType, notNull, autoIncrement, defaultStr, cc.Comment))
		cs[i] = cc
	}
	return cs
}

func Indexes(database string) []*Index {
	indexList := SelectTableIndexListSelectMapper.Prepare(database).Exec().List(new(Index))
	is := make([]*Index, len(indexList))
	for i, c := range indexList {
		ii := c.(*Index)
		is[i] = ii
	}
	return is
}

func TransformTables(tables []*Table) map[string]*Table {
	tableMap := make(map[string]*Table, len(tables))
	for _, t := range tables {
		tableMap[t.Name] = t
	}
	return tableMap
}

func TransformColumns(columns []*Column) map[string]*[]*Column {
	columnMap := make(map[string]*[]*Column, 0)
	for _, c := range columns {
		cs, have := columnMap[c.Table]
		if have {
			*cs = append(*cs, c)
		} else {
			csp := make([]*Column, 1)
			csp[0] = c
			columnMap[c.Table] = &csp
		}
	}
	return columnMap
}

func TransformIndexes(indexes []*Index) map[string]*[]*Index {
	indexMap := make(map[string]*[]*Index, 0)
	for _, i := range indexes {
		cs, have := indexMap[i.Table]
		if have {
			*cs = append(*cs, i)
		} else {
			csp := make([]*Index, 1)
			csp[0] = i
			indexMap[i.Table] = &csp
		}
	}
	return indexMap
}

func SetTableColumns(tableMap map[string]*Table, columnMap map[string]*[]*Column) {
	for k, v := range tableMap {
		if cc, have := columnMap[k]; have {
			v.Columns = append(v.Columns, *cc...)
		}
	}
}

func SetTableIndexes(tableMap map[string]*Table, indexMap map[string]*[]*Index) {
	for k, v := range tableMap {
		if vv, have := indexMap[k]; have {
			v.Indexes = append(v.Indexes, *vv...)
		}
	}
}
