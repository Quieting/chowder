package db

import (
	"fmt"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"

	"github.com/Quieting/chowder/script/xerror"
)

// sql 基本功能，执行sql语句，映射结果到结构体

const dbTag = "db"

var source = "root:vspnmysql123@tcp(192.168.0.12:3312)/battlesothebys?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"

func getDb() *sql.DB {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	return db
}

func Insert(sqlStr string, args ...interface{}) (int64, error) {
	db := getDb()
	defer db.Close()
	res, err := db.Exec(sqlStr, args...)
	if err != nil {
		return 0, xerror.New(err, fmt.Sprintf("执行SQL语句失败, sql: %s; arg: %+v", sqlStr, args))
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, xerror.New(err, fmt.Sprintf("执行SQL语句失败, sql: %s; arg: %+v", sqlStr, args))
	}
	return id, nil
}

func Exec(sqlStr string, args ...interface{}) (int64, error) {
	db := getDb()
	defer db.Close()
	res, err := db.Exec(sqlStr, args...)
	if err != nil {
		return 0, xerror.New(err, fmt.Sprintf("执行SQL语句失败, sql: %s; arg: %+v", sqlStr, args))
	}

	num, err := res.RowsAffected()
	if err != nil {
		return 0, xerror.New(err, fmt.Sprintf("执行SQL语句失败, sql: %s; arg: %+v", sqlStr, args))
	}
	return num, nil
}

// QueryRows 查找多条记录
// v : 切片结构体指针类型，形如 *[]struct{}
func QueryRows(v interface{}, sqlStr string, args ...interface{}) error {
	typ := reflect.TypeOf(v) // 指针类型，指向一个切片
	data := reflect.ValueOf(v)

	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()   // slice 类型，
		data = data.Elem() // slice 数据
	}

	data0 := data

	if typ.Kind() != reflect.Slice {
		return xerror.New(nil, fmt.Sprintf("需要一个数组、切片类型, 传入对象类型：%s", typ.Kind()))
	}

	_typ := typ.Elem() // 数组切片元素类型
	if _typ.Kind() != reflect.Struct {
		return xerror.New(nil, fmt.Sprintf("需要一个struct类型, 传入对象类型：%s", typ.Kind()))
	}

	db := getDb()
	defer db.Close()

	rows, err := db.Query(sqlStr, args...)
	if err != nil {
		return xerror.New(err, fmt.Sprintf("执行SQL语句失败, sql: %s; arg: %+v", sqlStr, args))
	}
	defer rows.Close()

	for rows.Next() {
		val := reflect.New(_typ)

		err = rows.Scan(dest(val.Interface(), rows)...)
		if err != nil {
			return xerror.New(err, fmt.Sprintf("解析数据失败"))
		}

		if _typ.Kind() == reflect.Struct {
			val = val.Elem()
		}

		data0 = reflect.Append(data0, val)
	}

	data.Set(data0)

	return nil
}

// QueryRow 查找单条记录，支持结构体和基础类型指针对象
// v : *struct{}、*int 等
func QueryRow(v interface{}, sqlStr string, args ...interface{}) error {
	typ := reflect.TypeOf(v)
	if typ.Kind() != reflect.Pointer {
		return xerror.New(nil, fmt.Sprintf("需要一个指针类型, 传入对象类型：%s", typ.Kind()))
	}

	switch typ.Elem().Kind() {
	case reflect.Int, reflect.Int64, reflect.Int32:
	case reflect.Struct:
	default:
		return xerror.New(nil, fmt.Sprintf("需要一个基本类型或者struct类型, 传入对象类型：%s", typ.Elem().Kind()))
	}

	db := getDb()
	defer db.Close()

	rows, err := db.Query(sqlStr, args...)
	if err != nil {
		return xerror.New(err, fmt.Sprintf("执行SQL语句失败, sql: %s; arg: %+v", sqlStr, args))
	}
	defer rows.Close()
	if !rows.Next() {
		return nil
	}

	if typ.Elem().Kind() == reflect.Struct {
		err = rows.Scan(dest(v, rows)...)
	} else {
		err = rows.Scan(v)
	}
	if err != nil {
		return xerror.New(err, fmt.Sprintf("解析SQL执行结果失败"))
	}
	return nil
}

// RawFieldNames converts golang struct field into slice string.
func RawFieldNames(in interface{}, postgresSql ...bool) []string {
	out := make([]string, 0)
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	var pg bool
	if len(postgresSql) > 0 {
		pg = postgresSql[0]
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		tagv := fi.Tag.Get(dbTag)
		switch tagv {
		case "-":
			continue
		case "":
			if pg {
				out = append(out, fi.Name)
			} else {
				out = append(out, fmt.Sprintf("`%s`", fi.Name))
			}
		default:
			// get tag name with the tag opton, e.g.:
			// `db:"id"`
			// `db:"id,type=char,length=16"`
			// `db:",type=char,length=16"`
			if strings.Contains(tagv, ",") {
				tagv = strings.TrimSpace(strings.Split(tagv, ",")[0])
			}
			if len(tagv) == 0 {
				tagv = fi.Name
			}
			if pg {
				out = append(out, tagv)
			} else {
				out = append(out, fmt.Sprintf("`%s`", tagv))
			}
		}
	}

	return out
}

func dest(v interface{}, rows *sql.Rows) []interface{} {
	val := reflect.ValueOf(v).Elem()
	fmt.Printf("reflect.TypeOf(v): %v\n", reflect.TypeOf(v))

	typ := reflect.TypeOf(v).Elem()
	columns, _ := rows.Columns()
	columnTypes, _ := rows.ColumnTypes()

	fmt.Printf("columns: %v\n", columns)

	values := make([]interface{}, 0, len(columns))
	for i := 0; i < len(columns); i++ {
		var item interface{}

		// 遍历获取 v field 地址
		for j := 0; j < typ.NumField(); j++ {
			if columns[i] == typ.Field(j).Name ||
				columns[i] == typ.Field(j).Tag.Get(dbTag) {
				if typ.Field(j).Type.Kind() == reflect.Pointer {
					item = val.Field(j).Elem().Addr()
				} else {
					item = val.Field(j).Addr().Interface()
				}
				break
			}
		}

		if item == nil {
			item = reflect.New(columnTypes[i].ScanType()).Elem().Addr()
		}

		// fmt.Printf("itemtype: %s\n", item.Kind())
		// fmt.Printf("itemtype: %s\n", reflect.ValueOf(item).Kind())

		values = append(values, item)
	}

	return values
}
