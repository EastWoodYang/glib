package glib

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

/* ================================================================================
 * 反射
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Reflect
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type (
	Reflect struct {
		target  interface{}
		rawTypeOf reflect.Type
		rawValueOf reflect.Value
		typeOf  reflect.Type
		valueOf reflect.Value
	}

	ReflectFunc struct {
		funcs map[string]interface{}
	}

	FieldInfo struct {
		Index     []int
		Name      string
		Kind      string
		Type      string
		Value     reflect.Value
		Tag       reflect.StructTag
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
func NewReflect(target interface{}) *Reflect {
	tf := reflect.TypeOf(target)
	vf := reflect.ValueOf(target)

	r := &Reflect{
		target:     target,
		typeOf:     tf,
		valueOf:    vf,
		rawTypeOf:  tf,
		rawValueOf: vf,
	}

	if r.rawTypeOf.Kind() == reflect.Ptr {
		r.typeOf = r.rawTypeOf.Elem()
	}

	if r.rawValueOf.Kind() == reflect.Ptr {
		r.valueOf = r.rawValueOf.Elem()
	}

	return r
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化ReflectFunc
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
func GetStructFieldValueByName(target interface{}, fieldName string) (string, interface{}, error) {
	if target == nil || len(fieldName) == 0 {
		return "", nil, errors.New("argument error")
	}

	typeOf := reflect.TypeOf(target)
	valueOf := reflect.ValueOf(target)

	if typeOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}

	if _, ok := valueOf.Type().FieldByName(fieldName); !ok {
		return "", "", errors.New("fieldname not exists")
	}

	pkgName := strings.Split(valueOf.String(), " ")[0][1:]
	fieldValue := valueOf.FieldByName(fieldName).Interface()

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

	fnValue := reflect.ValueOf(fn)
	if len(args) != fnValue.Type().NumIn() {
		return nil, errors.New("func params not match")
	}

	in := make([]reflect.Value, len(args))
	for i, v := range args {
		in[i] = reflect.ValueOf(v)
	}

	return fnValue.Call(in), nil
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
 * 根据字段索引获取reflect.Value
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Field(index int) (fieldValue reflect.Value, er error) {
	return s.FieldByIndex([]int{index})
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据字段名称获取reflect.Value
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) FieldByName(name string) (reflect.Value, error) {
	if fieldType, isOk := s.valueOf.Type().FieldByName(name); isOk {
		return s.FieldByIndex(fieldType.Index)
	}

	return reflect.Value{}, errors.New("field name not found")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据字段索引链获取reflect.Value
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) FieldByIndex(indexs []int) (fieldValue reflect.Value, er error) {
	defer func() {
		if err := recover(); err != nil {
			if err, ok := err.(error); ok {
				er = err
			} else {
				er = errors.New("field not found")
			}
		}
	}()

	fieldValue = s.valueOf.FieldByIndex(indexs)

	return fieldValue, er
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据字段索引获取FieldInfo
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) FieldInfo(index int) (fieldInfo *FieldInfo, er error) {
	return s.FieldInfoByIndex([]int{index})
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据字段名称获取FieldInfo
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) FieldInfoByName(name string) (*FieldInfo, error) {
	if fieldType, isOk := s.valueOf.Type().FieldByName(name); isOk {
		return s.FieldInfoByIndex(fieldType.Index)
	}

	return nil, errors.New("not found fileld")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据字段索引链获取*FieldInfo
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) FieldInfoByIndex(indexs []int) (fieldInfo *FieldInfo, er error) {
	defer func() {
		if err := recover(); err != nil {
			if err, ok := err.(error); ok {
				er = err
			} else {
				er = errors.New("field not found")
			}
		}
	}()

    //fieldType := s.valueOf.Type().FieldByIndex(indexs)
    fieldType := s.typeOf.FieldByIndex(indexs)
	fieldValue := s.valueOf.FieldByIndex(indexs)

	fieldInfo = &FieldInfo{
		Index:     indexs,
		Name:      fieldType.Name,
		Kind:      fmt.Sprintf("%s", fieldType.Type.Kind()),
		Type:      fmt.Sprintf("%s", fieldType.Type),
		Value:     fieldValue,
		Tag:       fieldType.Tag,
		//Interface: fieldValue.Interface(),
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
		fieldType := s.typeOf.Field(i)
		fieldValue := s.valueOf.Field(i)

		fieldInfo := &FieldInfo{
			Index: fieldType.Index,
			Name:  fieldType.Name,
			Kind:  fmt.Sprintf("%s", fieldType.Type.Kind()),
			Type:  fmt.Sprintf("%s", fieldType.Type),
			Tag:   fieldType.Tag,
		}

		if fieldType.Type.Kind() == reflect.Struct {
			s.fieldChilds(fieldInfo.Index, fieldInfo, fieldType)
		} else {
			fieldInfo.Value = fieldValue
			fieldInfo.Interface = fieldValue.Interface()
		}

		fieldInfos = append(fieldInfos, fieldInfo)
	}

	return fieldInfos, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 遍历字段子集
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) fieldChilds(indexs []int, fieldInfo *FieldInfo, field reflect.StructField) {
	for j := 0; j < field.Type.NumField(); j++ {
		indexs := append(indexs, j)
		newField := s.typeOf.FieldByIndex(indexs)

		newFieldInfo := &FieldInfo{
			Index: indexs,
			Name:  newField.Name,
			Kind: fmt.Sprintf("%s", newField.Type.Kind()),
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

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据方法索引获取reflect.Value
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Method(index int) (reflect.Value, error) {
	v := s.valueOf.Method(index)

	if !v.IsValid() {
		return v, errors.New("method is valid")
	}

	return v, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据方法名称获取reflect.Value
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) MethodByName(name string) (reflect.Value, error) {
    methodTypeOf, _ := s.typeOf.MethodByName(name)
    return s.Method(methodTypeOf.Index)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据方法名称获取*MethodInfo
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) MethodInfoByName(name string) (*MethodInfo, error) {
	var err error
	method := s.valueOf.MethodByName(name)
	if !method.IsValid() {
		return nil, errors.New("method is valid")
	}

	methodTypeOf, _ := s.typeOf.MethodByName(name)
	numIn := method.Type().NumIn()
	numOut := method.Type().NumOut()

	methodInfo := &MethodInfo{
		Index:  methodTypeOf.Index,
		Name:   methodTypeOf.Name,
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
		return nil, errors.New("target not struct")
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
 * 获取底层类型
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Kind() reflect.Kind {
	return s.typeOf.Kind()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据索引获取字段reflect.StructTag
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Tag(index int) reflect.StructTag {
	return s.typeOf.Field(index).Tag
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
	return s.rawTypeOf.NumMethod()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 对象信息
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Interface() interface{} {
	return s.target
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 长度
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) Len() int {
	if s.IsSlice() || s.IsArray() {
        return s.valueOf.Len()
	}

	return 0
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

			if jsonName, isOk := fieldInfo.Tag.Lookup("json"); isOk {
				fieldName = jsonName

			    if fieldName != "-" {
				    info := fmt.Sprintf("%s:%v", fieldName, fieldValue)
				    dumps = append(dumps, info)
			    }
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
 * 是否结构对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsStruct() bool {
	return s.typeOf.Kind() == reflect.Struct
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否数组对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsArray() bool {
	return s.typeOf.Kind() == reflect.Array
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否切片对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsSlice() bool {
	return s.typeOf.Kind() == reflect.Slice
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否Map对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsMap() bool {
	return s.typeOf.Kind() == reflect.Map
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否func对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsFunc() bool {
	return s.typeOf.Kind() == reflect.Func
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 是否ptr对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsPtr() bool {
	return s.rawTypeOf.Kind() == reflect.Ptr
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断是否实现指定接口。
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsImplements(inter interface{}) bool {
	interTypeOf := reflect.TypeOf(inter)
	if interTypeOf.Kind() == reflect.Ptr {
		interTypeOf = interTypeOf.Elem()
	}

	return s.typeOf.Implements(interTypeOf)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断类型的值可否转换为指定类型
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsConvertibleTo(inter interface{}) bool {
	interTypeOf := reflect.TypeOf(inter)
	if interTypeOf.Kind() == reflect.Ptr {
		interTypeOf = interTypeOf.Elem()
	}

	return s.typeOf.ConvertibleTo(interTypeOf)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断类型的值可否赋值给指定类型
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsAssignableTo(inter interface{}) bool {
	interTypeOf := reflect.TypeOf(inter)
	if interTypeOf.Kind() == reflect.Ptr {
		interTypeOf = interTypeOf.Elem()
	}

	return s.typeOf.AssignableTo(interTypeOf)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 判断类型可否进行比较操作
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Reflect) IsComparable() bool {
	return s.typeOf.Comparable()
}
