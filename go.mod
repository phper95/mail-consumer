module product-consumer

go 1.16

require (
	gitee.com/phper95/pkg/cache v0.0.0-20220717100747-3cf477b9a8d7
	gitee.com/phper95/pkg/es v0.0.0-20220717102538-d460e19eb7dc
	gitee.com/phper95/pkg/file v0.0.0-20220717100747-3cf477b9a8d7
	gitee.com/phper95/pkg/logger v0.0.0-20220717100747-3cf477b9a8d7
	gitee.com/phper95/pkg/mq v0.0.0-20220717100747-3cf477b9a8d7
	gitee.com/phper95/pkg/shutdown v0.0.0-20220717100747-3cf477b9a8d7
	gitee.com/phper95/pkg/trace v0.0.0-20220717100747-3cf477b9a8d7
	//github.com/Shopify/sarama v1.29.1  //linux下使用此版本
	github.com/Shopify/sarama v1.19.0 //windows下使用此版本
	github.com/eapache/go-resiliency v1.3.0 // indirect
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/fsnotify/fsnotify v1.5.4
	github.com/go-redis/redis/v7 v7.4.1
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/spf13/viper v1.12.0
	github.com/tebeka/strftime v0.1.5 // indirect
	go.uber.org/zap v1.21.0
)
