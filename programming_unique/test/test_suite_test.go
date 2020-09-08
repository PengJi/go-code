package test_test

import (
	"testing"

	// 使用点号导入，把这两个包导入到当前命名空间
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTest(t *testing.T) {
	// 将Ginkgo的Fail函数传递给Gomega，Fail函数用于标记测试失败，这是Ginkgo和Gomega唯一的交互点
	// 如果Gomega断言失败，就会调用Fail进行处理
	RegisterFailHandler(Fail)

	// 启动测试套件
	RunSpecs(t, "Test Suite")
}
