package test_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// 为了方便，被测试包被导入当前命名空间
	. "github.com/pengji/go-code/programming_unique/test/books"
)

// 顶级的Describe容器

// Describe块用于组织Specs，其中可以包含任意数量的：
//    BeforeEach：在Spec（It块）运行之前执行，嵌套Describe时最外层BeforeEach先执行
//    AfterEach：在Spec运行之后执行，嵌套Describe时最内层AfterEach先执行
//    JustBeforeEach：在It块，所有BeforeEach之后执行
//    Measurement

// 可以在Describe块内嵌套Describe、Context、When块
var _ = Describe("Book", func() {
	var (
		// 通过闭包在 BeforeEach 和 It 之间共享数据
		longBook Book
		shortBook Book
	)

	// 此函数用于初始化Spec的状态，在It块之前运行。如果存在嵌套Describe，则最外面的BeforeEach最先运行
	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			// 通过It来创建一个Spec
			It("should be a novel", func() {
				// Gomega的Expect用于断言
				Expect(longBook.AuthorLastName()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.AuthorLastName()).To(Equal("SHORT STORY"))
			})
		})
	})
})

// 这是关于Book服务测试
var _ = Describe("Book", func() {
	var (
		book Book
		err error
	)

	BeforeEach(func() {
		book, err = NewBookFromJSON(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
        }`)
	})

	// 测试加载Book行为
	Describe("loading from JSON", func() {
		// 如果正常解析JSON
		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				// 期望                相等
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Pages).To(Equal(1488))
			})

			It("should not error", func() {
				// 期望      没有发生错误
				Expect(err).NotTo(HaveOccurred())
			})
		})
		// 如果无法解析JSON
		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				// 这是一个BDD反模式，可以用JustBeforeEach
				book, err = NewBookFromJSON(`{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":1488oops
                }`)
			})

			It("should return the zero-value for the book", func() {
				// 期望          为零
				Expect(book).To(BeZero())
			})

			It("should error", func() {
				// 期望        发生了错误
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Extracting the author's last name", func() {
		It("should correctly identify and return the last name", func() {
			Expect(book.AuthorLastName()).To(Equal("Hugo"))
		})
	})
})

var _ = Describe("Book", func() {
	var (
		book Book
		err error
		json string
	)
	// 准备默认JSON
	BeforeEach(func() {
		json = `{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
        }`
	})

	JustBeforeEach(func() {
		// 按需，根据默认数据/无效JSON创建book，避免NewBookFromJSON的重复调用（如果代价很高的话……）
		book, err = NewBookFromJSON(json)
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses succesfully", func() {
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				// 覆盖默认JSON为无效JSON
				json = `{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":1488oops
                }`
			})
		})
	})
})
