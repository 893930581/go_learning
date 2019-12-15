package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MySqlConfig struct{
	Host 		string		`ini:"host"`
	Port 		int			`ini:"port"`
	Username 	string		`ini:"username"`
	Password  	int			`ini:"password"`
}

type RedisConfig struct{
	Host 		string		`ini:"host"`
	Port 		int			`ini:"port"`
	Password 	string		`ini:"password"`
	DataBase  	int			`ini:"database"`
}

type Config struct{
	MySqlConfig	`ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func readIni(filename string,x interface{}) (err error) {
	/** x 参数校验*/
	t := reflect.TypeOf(x)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		//确认传入的是结构体指针类型
		err = errors.New("Arg x needs to be a ptr of struct")
		return
	}

	/** 开始文件处理*/
	f, err:= ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	c := string(f)
	//读取文件内容转化为字符串
	ca := strings.Split(c,"\r\n")
	//按行拆分文件内容
	var structName string
	//声明要读取配置文件名的结构体名（mysql/redis）
	for idx,line := range ca {
		//开始按行解析文件
		line = strings.TrimSpace(line)
		fmt.Println(line)
		if strings.HasPrefix(line,"#") || strings.HasPrefix(line,";") || len(line) == 0{
			continue
		}
		//判断是否为注释或者为空
		if strings.HasPrefix(line,"[") {
			//解析分配置标头
			section := strings.TrimSpace(line[1:len(line)-1])
			//获取 标头内容
			if !strings.HasSuffix(line,"]"){
				err = fmt.Errorf("line %d syntax error", idx+1)
				return
			}

			if len(section) == 0 {
				err = fmt.Errorf("line %d syntax error", idx+1)
				return
			}
			//判断标头格式是否正确
			for i := 0;i < t.Elem().NumField(); i++ {
				//通过反射 x 的字段名
				field := t.Elem().Field(i)
				if section == field.Tag.Get("ini") {
					structName = field.Name
					//设置当前要解析的配置文件标题，并且和结构体产生映射。
					fmt.Println("找到了字段",structName)
				}
			}
		}else{
			if len(structName) == 0 {
				continue
			}
			//防止首个配置内容在标题前面出现
			if strings.HasPrefix(line,"=") || strings.HasSuffix(line,"=") || strings.Count(line,"=") != 1{
				err = fmt.Errorf("line %d syntax error", idx+1)
				return 
			}
			breakPoint := strings.Index(line,"=")
			if breakPoint == -1 {
				err = fmt.Errorf("line %d syntax error", idx+1)
				return 
			}
			//判断配置内容是否正确
			key := strings.TrimSpace(line[:breakPoint])
			value := strings.TrimSpace(line[breakPoint+1:])
			//拆分配置文件
			v := reflect.ValueOf(x)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()
			//通过反射拿到要转化的结构体（MysqlConfig/ReadisConfig）的值反射和类型反射。
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("inValid key of x")
				return
			}
			//确定拿到的是结构体。
			var fieldName string
			//声明子结构体字段名
			for i := 0; i< sValue.NumField(); i++ {
				field := sType.Field(i)
				if field.Tag.Get("ini") == key {
					fieldName = field.Name
				}
 			}
			//遍历结构体字段名并且找到和配置文件中key对应的那个。
			tField,_ := sType.FieldByName(fieldName)
			//拿到字段Type ——> StructField
			fmt.Println(fieldName)
			fmt.Println(tField.Type.Kind() )
			switch tField.Type.Kind() {
			case reflect.Int,reflect.Int32,reflect.Int64:
				tValue,_ := strconv.Atoi(value)
				sValue.FieldByName(fieldName).SetInt(int64(tValue))
			case reflect.String:
				sValue.FieldByName(fieldName).SetString(value)
			default:
				sValue.FieldByName(fieldName).SetString(value)
			}
			//通过结构体类型转换配置字段，并且个结构体Field赋值。
		}
	}
	return
}

func mapStruct(x interface{}) (err error) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		//确认传入的是结构体指针类型
		err = errors.New("Arg x needs to be a ptr of struct")
		return
	}
	for i := 0;i < v.Elem().NumField() ;i++  {
		key := v.Elem().Field(i).Type().Name()
		value := v.Elem().Field(i).Interface()
		fmt.Printf("key: %v,value: %v\n",key,value)
	}
	return
}

func main() {
	con := new(Config)
	err := readIni("./config.ini",con)
	fmt.Println(err)
	err = mapStruct(con)
	fmt.Println(err)
}
