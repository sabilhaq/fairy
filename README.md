# Fairy

[![Go Report Card](https://goreportcard.com/badge/github.com/rl404/fairy)](https://goreportcard.com/report/github.com/rl404/fairy)
![License: MIT](https://img.shields.io/github/license/rl404/fairy.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/rl404/fairy.svg)](https://pkg.go.dev/github.com/rl404/fairy)

_Fairy_ contains several general tools for easier and simpler development.

- Swappable cache
  - [Redis](https://redis.io/)
  - [In-memory](https://github.com/allegro/bigcache)
  - [Memcached](https://memcached.org/)
- Swappable pubsub
  - [Redis](https://redis.io/)
  - [RabbitMQ](https://rabbitmq.com/)
  - [NSQ](https://nsq.io/)
  - [Google PubSub](https://cloud.google.com/pubsub)
- Struct modification and validation
- Logging & http middleware logging
  - [Logrus](https://github.com/sirupsen/logrus)
  - [Zap](https://github.com/uber-go/zap)
  - [Zerolog](https://github.com/rs/zerolog)
  - [Elasticsearch](https://www.elastic.co/)
  - [Newrelic](https://newrelic.com/)
- Error stack tracing
- Monitoring
  - [Prometheus](https://prometheus.io/)
    - Cache
    - Database ([GORM](https://gorm.io/))
    - HTTP & GRPC
  - [Newrelic](https://newrelic.com/)
    - Cache
    - Database ([GORM](https://gorm.io/))
    - HTTP & GRPC
    - Pubsub
- Rate-limiting

All these tools are simplified version of the original ones
and for general use. Your case may need a more complex one
but I hope these tools can help you or at least become a reference
for your own tools.

Good luck.

---

## Installation

```
go get github.com/rl404/fairy
```

## Usage

Please go to the [documentation](https://pkg.go.dev/github.com/rl404/fairy) or `example` folder.

## License

MIT License

Copyright (c) 2021 Axel