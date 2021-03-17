#etcd注册组件接口开发
##分析，类似于session中间件
###定义服务注册总接口Registry,定义方法
####Name():插件名，例如传etcd
####Init(opts ...Option)：初始化，里面用选项设计模式做初始化
####Register():服务注册
####Unregister():服务反注册，例如服务端停了，注册列表销毁
####GetService:服务发现（ip port[] string）
###抽象出一些结构体
#####Node:单个节点的结构体，包含id ip port weight（权重）
####Service:里面有服务名，还有节点列表，一个服务多台服务器支撑
###选项设计模式，实现参数初始化
###插件管理类
####可以用一个大的map管理，key字符串，value是Registry接口对象
####用户自定义去调用，自定义插件
####实现注册中心的初始化，供系统使用
