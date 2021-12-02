package Builder

// 构造者接口
type Builder interface {
	Build()
}



// 管理者
type Director struct {
	builder Builder
}
// 管理者 - 初始化
func NewBuilder(b Builder) Director {
	return Director{
		builder: b,
	}
}
// 管理者 - 执行创建 builder.Build()
func (d *Director) Construct() {
	d.builder.Build()
}



// Concrete 具体的构造者
type ConcreteBuilder struct {
	built bool
}
func NewConcreteBuilder() ConcreteBuilder{
	return ConcreteBuilder{false}
}
func (c *ConcreteBuilder) Build() {
	c.built = true
}
func (c *ConcreteBuilder) GetResult() Product {
	return Product{c.built}
}


type Product struct {
	Built bool
}


