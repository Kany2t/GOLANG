1.protobuf协议声明服务与消息
2.一元请求与流式请求
3.header与trailer元数据传递

grpc开步骤
1.安装protoc编译器，根据protobuf定义的文件生成grpc客户端与服务器存根
2。安装protoc-gen-go插件，生成go语言代码
3.安装grpc插件，生成grpc代码
4.protobuf定义生成message与service
6.基于服务端存根实现grpc server端
7.将存根共享给客户端，以便rpc调用server端