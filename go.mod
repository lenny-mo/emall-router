module github.com/lenny-mo/router

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/bitly/go-simplejson v0.5.1 // indirect
	github.com/coreos/etcd v3.3.27+incompatible // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20230601102743-20bbbf26f4d8 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/emirpasic/gods v1.12.1 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-gonic/gin v1.9.1
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.8.6 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/lenny-mo/cart-api v0.0.0-20231216061546-0fb05a4c41fa
	github.com/lenny-mo/order-api v0.0.0-20231217140329-72c2787608e7
	github.com/lenny-mo/payment-api v0.0.0-20231212055245-13d741f8de57
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/nats-io/nkeys v0.4.6 // indirect
	google.golang.org/protobuf v1.31.0
)
