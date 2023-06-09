package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

type FuncService interface {
	Invoke(inputs string) string
}

type FuncServiceImpl struct {
}

func (t *FuncServiceImpl) Invoke(inputs string) string {
	return inputs
}

// 目的：根据实现类的名字创建实现类
// golang根据接口实现类的名字创建实现类
// golang 工厂模式 配置化实现
func main() {
	impl3()
}

func impl1() {
	// 实现类类型名字
	implTypeName := "Service"
	// 获取实现类类型
	implType := reflect.TypeOf(FuncServiceImpl{})
	if !implType.Implements(reflect.TypeOf((*FuncService)(nil)).Elem()) {
		fmt.Printf("Implementation type %s does not implement FuncService\n", implTypeName)
		return
	}
	// 创建实现类实例
	implValue := reflect.New(implType.Elem())
	implInstance := implValue.Interface().(FuncService)

	outputs := implInstance.Invoke("1111")

	fmt.Println(outputs)
}

func impl2() {
	// 实现类类型名字
	implTypeName := "FuncServiceImpl"
	// 获取实现类类型
	implType, ok := reflect.TypeOf((*FuncService)(nil)).Elem().FieldByNameFunc(func(name string) bool {
		return name == implTypeName
	})
	if !ok {
		fmt.Printf("Implementation type %s not found\n", implTypeName)
		return
	}
	// 创建实现类实例
	implValue := reflect.New(implType.Type).Elem()
	implInstance := implValue.Addr().Interface().(FuncService)

	outputs := implInstance.Invoke("1111")

	fmt.Println(outputs)
}

func impl3() {
	// 读取配置文件
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析配置文件
	var config map[string]string
	if err := json.Unmarshal(configData, &config); err != nil {
		fmt.Println(err)
		return
	}

	// 获取实现类的名称
	serviceName, ok := config["funcService"]
	if !ok {
		fmt.Println("service name not found in config")
		return
	}

	// 获取实现类的类型
	serviceType, err := getServiceType(serviceName)
	if err != nil {
		fmt.Println("failed to get service type:", err)
		return
	}

	// 创建实现类的实例
	service := reflect.New(serviceType).Interface().(FuncService)

	// 使用实现类的方法示例
	output := service.Invoke("Hello World!")
	fmt.Println(output)
}

func getServiceType(name string) (reflect.Type, error) {
	// 查找实现类的类型
	for _, t := range []reflect.Type{
		reflect.TypeOf((*FuncServiceImpl)(nil)),
	} {
		if t.Name() == name {
			return t, nil
		}
	}
	return nil, fmt.Errorf("service type not found: %s", name)
}
