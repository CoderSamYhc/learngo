package ops

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{} // 订阅者通道
	topicFunc func(v interface{}) bool // 主题过滤器
)

type Publisher struct {
	lock sync.RWMutex // 锁
	buf int // 订阅队列的缓存大小
	timeout time.Duration // 发布超时时间
	subscribers map[subscriber]topicFunc
}

// 创建一个发布者，可设置超时时间和队列长度
func NewPublisher(timeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buf: buffer,
		timeout: timeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// 添加一个新订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscriberTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buf)
	p.lock.Lock()
	p.subscribers[ch] = topic
	p.lock.Unlock()
	return ch
}

// 添加一个新订阅者，订阅全部主题
func (p *Publisher) Subscriber() chan interface{} {
	return p.SubscriberTopic(nil)
}

// 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// 发送主题，容忍超时
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):

	}
}

// 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)

		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 关闭发布者，关闭订阅者通道
func (p *Publisher) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}





