package glib

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

/* ================================================================================
 * Oauth Qq
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Reflect
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type (
	Reflect struct {
		obj     interface{}
		typeOf  reflect.Type
		valueOf reflect.Value
	}

	ReflectFunc struct {
		funcs map[string]interface{}
	}

	FieldInfo struct {
		Index     []int
		Name      string
		Tag       map[string]interface{}
		Type      string
		Value     reflect.Value
		Interface interface{}
		Childs    []*FieldInfo
	}

	MethodInfo struct {
		Index  int
		Name   string
		Value  reflect.Value
		NumIn  int
		NumOut int
		In     []string
		Out    []string
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化Reflect
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewReflectFunc() *ReflectFunc {
	r := &ReflectFunc{
		funcs: make(map[string]interface{}, 0),
	}

	return r
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据结构体和字段名获取对应的包名和字段值
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetFieldByName(model interface{}, fieldName string) (string, string, error) {
	typeOf := reflect.TypeOf(model)
	if kind := typeOf.Kind(); kind != reflect.Ptr {
		panic("Model is not a pointer type")
	}

	valueElem := reflect.ValueOf(model).Elem()

	if len(fieldName) == 0 {
		fieldName = "Id"
	}

	if _, ok := valueElem.Type().FieldByName(fieldName); !ok {
		return "", "", errors.New("fieldname not exists")
	}

	pkgName := strings.Split(valueElem.String(), " ")[0][1:]
	fieldValue := valueElem.FieldByName(fieldName).String()

	return pkgName, fieldValue, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 注册Reflect
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *ReflectFunc) Register(key string, fn interface{}) error {
	valueOf := reflect.ValueOf(fn)
	if valueOf.Kind() != reflect.Func {
		return errors.New("not func")
	}

	if _, isExists := s.funcs[key]; isExists {
		return errors.New("fn map key is exists")
	}

	s.funcs[key] = fn
	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 动态调用函数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *ReflectFunc) Call(key string, args ...interface{}) ([]reflect.Value, error) {
	fn, isExists := s.funcs[key]
	if !isExists {
		return nil, errors.New("func name not exists")
	}

	valueOf := reflect.ValueOf(fn)
	if len(args) != valueOf.Type().NumIn() {
		return nil, errors.New("func params not match")
	}

	in := make([]reflect.Value, len(args))
	for i, v := range args {
		in[i] = reflect.ValueOf(v)
	}

	return valueOf.Call(in), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化Reflect
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewReflect(obj interface{}) *Reflect {
	r := &Reflect{
		obj: obj,
	}

	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		r.typeOf = reflect.TypeOf(obj).Elem()
	} else {
		r.typeOf = reflect.TypeOf(obj)
	}

	if reflect.ValueOf(obj).Kind() == reflect.Ptr {
		r.valueOf = reflect.ValueOf(obj).Elem()
	} else {
		r.valueOf = reflect.ValueOf(obj)
	}

	return r
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取字段数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) NumField() int {
	if !s.IsStruct() {
		return 0
	}

	return s.typeOf.NumField()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取方法数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) NumMethod() int {
	return s.typeOf.NumMethod()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据索引获取字段reflect.StructTag
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Tag(index int) reflect.StructTag {
	return s.typeOf.Field(index).Tag
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置字段值
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) SetField(name string, value interface{}) error {
	fieldValue, err := s.FieldByName(name)
	if err != nil {
		return err
	}

	if !fieldValue.CanSet() {
		return errors.New("field disabled set")
	}

	switch value.(type) {
	case int, int8, int16, int32, int64:
		fieldValue.SetInt(value.(int64))
	case uint, uint8, uint16, uint32, uint64:
		fieldValue.SetUint(value.(uint64))
	case float32, float64:
		fieldValue.SetFloat(value.(float64))
	case string:
		fieldValue.SetString(value.(string))
	case bool:
		fieldValue.SetBool(value.(bool))
	}

	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据索引获取字段reflect.Value
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Field(index int) (fieldValue reflect.Value, er error) {
	defer func() {
		if err := recover(); err != nil {
			if err, ok := err.(error); ok {
				er = err
			} else {
				er = errors.New("field not found")
			}
		}
	}()

	fieldValue = s.valueOf.Field(index)

	return fieldValue, er
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据索引链获取字段reflect.Value
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) FieldByIndex(index []int) (fieldValue reflect.Value, er error) {
	defer func() {
		if err := recover(); err != nil {
			if err, ok := err.(error); ok {
				er = err
			} else {
				er = errors.New("field not found")
			}
		}
	}()

	fieldValue = s.valueOf.FieldByIndex(index)

	return fieldValue, er
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据名称获取字段字段reflect.Value
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) FieldByName(name string) (reflect.Value, error) {
	if !s.IsStruct() {
		return reflect.Value{}, errors.New("not struct ")
	}

	_, ok := s.valueOf.Type().FieldByName(name)
	if !ok {
		return reflect.Value{}, errors.New("not found fileld")
	}

	return s.valueOf.FieldByName(name), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据索引获取字段*FieldInfo
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) FieldInfo(index int) (fieldInfo *FieldInfo, er error) {
	defer func() {
		if err := recover(); err != nil {
			if err, ok := err.(error); ok {
				er = err
			} else {
				er = errors.New("field not found")
			}
		}
	}()

	field := s.valueOf.Field(index)

	fieldInfo = &FieldInfo{
		Index:     []int{index},
		Name:      s.typeOf.Field(index).Name,
		Type:      fmt.Sprintf("%s", s.typeOf.Field(index).Type),
		Value:     field,
		Interface: field.Interface(),
	}

	return fieldInfo, er
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据索引链获取字段*FieldInfo
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) FieldInfoByIndex(index []int) (fieldInfo *FieldInfo, er error) {
	defer func() {
		if err := recover(); err != nil {
			if err, ok := err.(error); ok {
				er = err
			} else {
				er = errors.New("field not found")
			}
		}
	}()

	fieldValue := s.valueOf.FieldByIndex(index)
	fieldType := s.valueOf.Type().FieldByIndex(index)

	fieldInfo = &FieldInfo{
		Index:     index,
		Name:      fieldType.Name,
		Type:      fmt.Sprintf("%s", fieldType.Type),
		Value:     fieldValue,
		Interface: fieldValue.Interface(),
	}

	return fieldInfo, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据名称获取字段*FieldInfo
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) FieldInfoByName(name string) (*FieldInfo, error) {
	if !s.IsStruct() {
		return nil, errors.New("not struct ")
	}

	field, ok := s.valueOf.Type().FieldByName(name)
	if !ok {
		return nil, errors.New("not found fileld")
	}

	fieldInfo := &FieldInfo{
		Index:     field.Index,
		Name:      field.Name,
		Type:      fmt.Sprintf("%s", field.Type),
		Value:     s.valueOf.FieldByName(name),
		Interface: s.valueOf.FieldByName(name).Interface(),
	}

	return fieldInfo, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取字段集合[]*FieldInfo
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Fields() ([]*FieldInfo, error) {
	if !s.IsStruct() {
		return nil, errors.New("not struct ")
	}

	var fieldInfos []*FieldInfo
	var err error

	for i := 0; i < s.typeOf.NumField(); i++ {
		field := s.typeOf.Field(i)

		tagMap := make(map[string]interface{}, 0)
		tagMap["json"] = field.Tag.Get("json")

		fieldInfo := &FieldInfo{
			Index: field.Index,
			Name:  field.Name,
			Type:  fmt.Sprintf("%s", field.Type),
			Tag:   tagMap,
		}

		if field.Type.Kind() == reflect.Struct {
			s.fieldChilds(fieldInfo.Index, fieldInfo, field)
		} else {
			fieldInfo.Value = s.valueOf.Field(i)
			fieldInfo.Interface = s.valueOf.Field(i).Interface()
		}

		fieldInfos = append(fieldInfos, fieldInfo)
	}

	return fieldInfos, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 遍历字段子集
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) fieldChilds(indexs []int, fieldInfo *FieldInfo, field reflect.StructField) {
	if field.Type.Kind() == reflect.Struct {
		for j := 0; j < field.Type.NumField(); j++ {
			indexs := append(indexs, j)
			newField := s.typeOf.FieldByIndex(indexs)

			newFieldInfo := &FieldInfo{
				Index: indexs,
				Name:  newField.Name,
				Type:  fmt.Sprintf("%s", newField.Type),
			}

			if newField.Type.Kind() == reflect.Struct {
				s.fieldChilds(indexs, newFieldInfo, newField)
			} else {
				newFieldInfo.Value = s.valueOf.FieldByIndex(indexs)
				//newFieldInfo.Interface = s.valueOf.FieldByIndex(indexs).Interface()
			}

			fieldInfo.Childs = append(fieldInfo.Childs, newFieldInfo)
		}
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据索引获取方法reflect.Value
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Method(index int) (reflect.Value, error) {
	var err error
	v := s.valueOf.Method(index)
	if !v.IsValid() {
		err = errors.New("fileld is valid")
	}
	return v, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据名称获取方法reflect.Value
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) MethodByName(name string) (reflect.Value, error) {
	var err error
	v := s.valueOf.MethodByName(name)
	if !v.IsValid() {
		err = errors.New("fileld is valid")
	}
	return v, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据名称获取方法*MethodInfo
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) MethodInfoByName(name string) (*MethodInfo, error) {
	var err error
	method := s.valueOf.MethodByName(name)
	if !method.IsValid() {
		return nil, errors.New("method is valid")
	}

	mTypeOf, _ := s.typeOf.MethodByName(name)

	numIn := method.Type().NumIn()
	numOut := method.Type().NumOut()

	methodInfo := &MethodInfo{
		Index:  mTypeOf.Index,
		Name:   mTypeOf.Name,
		Value:  method,
		NumIn:  numIn,
		NumOut: numOut,
	}

	if numIn > 0 {
		inTypes := make([]string, numIn)
		for i := 0; i < numIn; i++ {
			inType := method.Type().In(i)
			inTypes = append(inTypes, inType.String())
		}

		methodInfo.In = inTypes
	}

	if numOut > 0 {
		outTypes := make([]string, numOut)
		for i := 0; i < numOut; i++ {
			outType := method.Type().Out(i)
			outTypes = append(outTypes, outType.String())
		}

		methodInfo.Out = outTypes
	}

	return methodInfo, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取方法集合[]*MethodInfo
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Methods() []*MethodInfo {
	var methods []*MethodInfo

	for i := 0; i < s.NumMethod(); i++ {
		method := s.valueOf.Method(i)
		numIn := method.Type().NumIn()
		numOut := method.Type().NumOut()

		methodInfo := &MethodInfo{
			Index:  s.typeOf.Method(i).Index,
			Name:   s.typeOf.Method(i).Name,
			Value:  method,
			NumIn:  numIn,
			NumOut: numOut,
		}

		if numIn > 0 {
			inTypes := make([]string, numIn)
			for i := 0; i < numIn; i++ {
				inType := method.Type().In(i)
				inTypes = append(inTypes, inType.String())
			}

			methodInfo.In = inTypes
		}

		if numOut > 0 {
			outTypes := make([]string, numOut)
			for i := 0; i < numOut; i++ {
				outType := method.Type().Out(i)
				outTypes = append(outTypes, outType.String())
			}

			methodInfo.Out = outTypes
		}

		methods = append(methods, methodInfo)
	}

	return methods
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 调用指定名称的方法
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Invoke(methodName string, args ...interface{}) ([]reflect.Value, error) {
	if !s.IsStruct() {
		return nil, errors.New("obj not struct")
	}

	in := make([]reflect.Value, len(args))
	for i := range args {
		in[i] = reflect.ValueOf(args[i])
	}

	method, err := s.MethodByName(methodName)
	if err != nil {
		return nil, err
	}

	return method.Call(in), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象信息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Interfce() interface{} {
	return s.obj
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 类型信息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Type() reflect.Type {
	return s.typeOf
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 值信息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Value() reflect.Value {
	return s.valueOf
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 类型类别
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Kind() reflect.Kind {
	return s.typeOf.Kind()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 类型大小
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Size() uintptr {
	return s.typeOf.Size()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 输出信息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Dump() string {
	dumps := make([]string, 0)

	if fieldInfos, err := s.Fields(); err == nil {
		for _, fieldInfo := range fieldInfos {
			fieldName := fieldInfo.Name
			fieldValue := fieldInfo.Value

			if jsonName, isOk := fieldInfo.Tag["json"]; isOk {
				fieldName = jsonName.(string)
			}

			if fieldName != "-" {
				info := fmt.Sprintf("%s:%v", fieldName, fieldValue)
				dumps = append(dumps, info)
			}
		}
	}

	jsonDumpString := ""
	if jsonDump, err := json.Marshal(dumps); err == nil {
		jsonDumpString = string(jsonDump)
		if len(jsonDumpString) > 0 {
			if isExists := strings.HasPrefix(jsonDumpString, "["); isExists {
				jsonDumpString = strings.TrimLeft(jsonDumpString, "[")
			}

			if isExists := strings.HasSuffix(jsonDumpString, "]"); isExists {
				jsonDumpString = strings.TrimRight(jsonDumpString, "]")
			}
		}
	}

	return fmt.Sprintf("\r\n===== Dump Begin =====\r\n%s\r\n===== Dump End =====", jsonDumpString)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否切片对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsSlice() bool {
	return s.typeOf.Kind() == reflect.Slice
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否结构对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsStruct() bool {
	return s.typeOf.Kind() == reflect.Struct
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否实现了 u 接口。
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsImplements(u reflect.Type) bool {
	return s.typeOf.Implements(u)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断类型的值可否转换为 u 类型
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsConvertibleTo(u reflect.Type) bool {
	return s.typeOf.ConvertibleTo(u)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断类型的值可否赋值给 u 类型
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsAssignableTo(u reflect.Type) bool {
	return s.typeOf.AssignableTo(u)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断类型的值可否进行比较操作
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsComparable() bool {
	return s.typeOf.Comparable()
}
