package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

// import "github.com/elliotchance/orderedmap/v2"

func main(){
	// Go 中的 map 是一种无序的键值对集合，底层基于哈希表实现，提供了高效的查找、插入和删除操作。

	// 1.声明与初始化
	// var 声明，得到 nil map（不能写入，只能读）
	var m1 map[string]int  // key为string val为map[key]
	v1 := m1["score"] // 读取不存在的 key 返回零值，不会 panic
	fmt.Printf("score: %v\n", v1) // score: 0
	// m1["score"] = 98 // panic: assignment to entry in nil map  m1为nil零值会抛出panic

	// 直接使用make初始化map(常用)
	scores := make(map[string]float64)
	scores["math"] = 100
	scores["english"] = 95
	fmt.Println(scores) // map[english:95 math:100]
	fmt.Println(scores["math"]) // 100
	// 指定初始容量（性能优化）
	// 预计这个 map 大概要存放 1000 个左右的键值对，请提前分配好足够的内存，避免后面频繁扩容。”
	// 但这只是一个建议值，不是严格的容量上限。完全可以往里面存超过 1000 个元素，map 会自动扩容，不会有任何报错，也不会受这个 1000 的限制。
	users := make(map[string]string, 1000)  // 预分配约1000个元素的空间
	users["tom"] = "ok"
	users["jack"] = "no"
	fmt.Println(users, len(users)) //map[jack:no tom:ok] 2 
	// 注意：map只有 len(m) 获取长度，没有 cap()。


	// 字面量初始化
	colors := map[string]string{
		"red": "#FF0000",
		"gree": "#00FF00",
		"blue": "#0000FF",
	}
	fmt.Println(colors["red"])
	// nil map 能读不能写；空 map（make/字面量）能读能写。

	// 2.对map中元素进行增删改查
	scores2 := make(map[string]float64)

	// 新增key和修改
	scores2["math"] = 100.5 // 新增一个不存在key math
	scores2["math"] = 99 // 修改已存在的key，则覆盖
	fmt.Println(scores2["math"]) // 99
	// 获取值
	val := scores2["english"] //  如果 key 不存在，返回零值（int 为 0）
	fmt.Println(val) // 0
	// 使用comma-ok 模式，来判断key是否存在
	if v, ok := scores2["math"]; ok {
		fmt.Printf("math存在：%v\n", v) // math存在：99
	} else {
		fmt.Println("math不存在")
	}
	// 只关心key是否存在
	_, exists := scores2["english"]
	fmt.Println("english是否存在:", exists) // english是否存在: false

	// 删除 直接通过delete函数
	mp := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
	}
	delete(mp, "a")
	if _, ok := mp["a"]; !ok {
		fmt.Println("a不存在")
	}
	fmt.Println(mp) // map[b:1 c:2 d:3]
	// 注意的是，如果值为 NaN，甚至没法删除该键值对。
	mp2 := make(map[float64]string, 10)
	// NaN 是 IEEE 754 浮点数标准中定义的一种特殊值，用来表示"无意义的运算结果" 
	// 0 / 0, 对负数开平方 math.Sqrt(-1),  ∞ - ∞, 0 × ∞
	mp2[math.NaN()] = "a"
	mp2[math.NaN()] = "b"
	mp2[math.NaN()] = "c"
	fmt.Println(mp2)
	// Go 的 map 在查找/插入 key 时，要比较 key 是否相等。每次调用 math.NaN() 都返回一个 NaN，而 NaN != NaN，
	// 所以 map 认为"这是一个全新的 key"，于是连续插入三条记录。
	// 同理，删除也删不掉：
	delete(mp2, math.NaN()) // 无效！因为传入的 NaN != 已存储的 Na
	fmt.Println(mp2) // 仍然是 map[NaN:a NaN:b NaN:c]

	// 通过len获取长度
	scores2["pe"] = 60
	length := len(scores2)
	fmt.Println(length)

	// 3.遍历map
	mp3 := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
	}
	// map 的 for range 遍历顺序是随机的，是由于其底层哈希表实现和遍历时的随机起始位置设计决定的。
	for key, val := range mp3 {
		fmt.Printf("map[\"%s\"]: %d\n", key, val) // map["a"]: 0
	}
	// 只遍历键：
	for key := range mp3 {
		fmt.Printf("map[\"%s\"]\n", key)
	}
	// 需要有序遍历，通过slice切片进行控制
	// 方式一：按key排序遍历
	mp4 := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
	}
	// 将key取出，放置slice切片中
	keys := make([]string, 0, len(mp4)) // 取出容量与mp4长度一致
	for key := range mp4 { // 遍历map4,将key放置slice切片中
		keys = append(keys, key)
	}
	// 使用sort包对slice切片排序
	sort.Strings(keys) // Ints / Float64s 同理
	// 按keys中排序后的顺序进行遍历map
	for _, k := range keys {
		fmt.Printf("map[\"%s\"]:%d\n",k, mp4[k])
	}

	mp5 := map[string]int{
    "banana": 2,
    "apple":  1,
    "cherry": 3,
	}

	// 方式二：按map的value值进行排序
	keys = make([]string, 0, len(mp5)) // 取出容量与mp4长度一致
	for key := range mp5 { // 遍历map4,将key放置slice切片中
		keys = append(keys, key)
	}
	// 在Go 1.21 及以上版本，可以使用标准库 slices 包自定义排序逻辑。
	sort.Slice(keys, func(i, j int) bool {
		// 从小到大
		return mp5[keys[i]] < mp5[keys[j]] // 返回true,不变key[i]和key[j]不交换顺序，返回false,key[i]和key[j]交换顺序
	})
	for _, k := range keys {
		fmt.Printf("map[\"%s\"]:%d\n",k, mp5[k])
	}

	// 方式三：按插入顺序遍历
	orderMap := NewOrderedMap() // orderedMap实例
	orderMap.Set("aaa", 1)
	orderMap.Set("ccc", 3)
	orderMap.Set("bbb", 2) // 在插入时已经已经o.keys中维护好插入顺序
	orderMap.Range() // 遍历
	// 或者使用第三方库 github.com/elliotchance/orderedmap/v2
	// m := orderedmap.NewOrderedMap[string, int]()
	// m.Set("banana", 3)
	// m.Set("apple", 1) // 后插入但在前面
	// m.Set("peach", 2)

	// for _, k := range m.Keys() {
	// 		v, _ := m.Get(k)
	// 		println(k, v)
	// }

	// 4.并发安全 map
	// 使用 sync.Map 配合排序（并发场景）
	// 需并发安全且有序遍历时。
	var sm sync.Map
	// 写入数据
	sm.Store("banana", 3)
	sm.Store("apple", 1)
	sm.Store("peach", 2)

	// Range收集key存放到keys切片中
	keys = make([]string, 0)
	sm.Range(func(key, value any) bool {
		// key 为any进行断言成string
		if k, ok := key.(string); ok { // 类型为string,则添加到切片中
			keys = append(keys, k)
		}
		return true // 返回 false 会提前终止遍历
	})
	// 在通过对 key 排序
	sort.Strings(keys)
	// 按排序后的 key 取值
	for _, k := range keys {
		if v, ok := sm.Load(k); ok {
			fmt.Printf("%s -> %v\n",k, v)
		}
	}
	/*
	为什么非常不推荐？（核心问题）
	① 性能严重劣化（双重锁开销）
	sync.Map 的 Range 内部已经持有锁（或使用原子操作）遍历整个底层数据结构。
	你在外部再调用 Load，每次 Load 又要重新加锁或走原子读。
	相比之下，普通 map + sync.RWMutex 在遍历排序时，只需在 Range 阶段加一次读锁，提取完 Key 后即可释放锁，之后读取值甚至无需再锁（如果 map 不变），性能远超 sync.Map。

	② 并发安全问题（数据漂移）
	从调用 Range 提取 Keys，到排序，再到 Load 取值，这期间是分步的。
	如果其他 goroutine 在这期间删除了某个 Key，你的 Load 会返回 false，导致数据缺失。
	如果其他 goroutine 修改了某个 Key 的 Value，你取到的是最新值，虽然不报错，但严格意义上你遍历的“快照”并不一致（因为你提取 Keys 那一刻的值，和最终打印时的值可能不同）。

	③ 类型断言繁琐
	sync.Map 的 Key 和 Value都是 interface{}，提取 Keys 时必须强制类型断言（如 key.(string)），如果 Key 类型不统一，极易引发 panic。
	*/
	// map + sync.RWMutex 的例子
	safeMap1 := NewSafeMap()
	safeMap1.items["cc"] = 3
	safeMap1.items["aa"] = 1
	safeMap1.items["bb"] = 2
	sKeys := safeMap1.SortedKeys()
	for _, key := range sKeys {
		v := safeMap1.items[key]
		fmt.Printf("key[%s] = %d\n", key, v)
	}

	// map + sync.RWMutex 的并发例子
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	// 5 个 goroutine 并发写
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func (i int)  {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("key%d", i), i)
		}(i)
	}
	// 5 个 goroutine 并发读
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if v, ok := safeMap.Get(fmt.Sprintf("key%d", i)); ok {
				fmt.Printf("key%d = %d\n", i, v)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("全部完成")

}
// 方式三：按插入顺序遍历
// map 本身不记录插入顺序。如果业务需要保持插入顺序（比如配置、日志、LRU），必须额外维护一个 slice 来记录顺序：
type OrderedMap struct {
	keys []string // 维护插入的顺序
	values map[string]int
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{values: make(map[string]int)}
}

func (o *OrderedMap)Set(k string, v int)  {
	// 将传入k存放在order的维护keys切片中
	if _, ok := o.values[k]; !ok {
		o.keys = append(o.keys, k)
	}
	// v值存放o.map中
	o.values[k] = v
}

// 按o.keys中的顺序进行遍历
func (o *OrderedMap)Range()  {
	for _, k := range o.keys {
		fmt.Printf("map[\"%s\"]:%d\n", k, o.values[k])
	}
}

// 方式五：map + sync.RWMutex 的例子
// SafeMap 把普通map 和读写锁封装在一起，实现并发安全
type SafeMap struct {
	mu sync.RWMutex  // 读写锁
	items map[string]int // 普通map
}

// NewSafeMap 构造函数，返回一个实例
func NewSafeMap() *SafeMap{
	return &SafeMap{ items: make(map[string]int) }
}

func (s *SafeMap) SortedKeys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// 创建一个与map容量一样长度的key切片
	keys := make([]string, len(s.items))
	for key := range s.items {
		keys = append(keys, key)
	}
	// 在sort进行排序
	sort.Strings(keys)
	return keys
  // 遍历时：
	// sm.mu.RLock() -> 拿到 keys -> 遍历 keys 读取 sm.m[key] -> sm.mu.RUnlock()
  // 全程加锁，保证一次性快照一致性，且只加一次锁。
}

// Set 写入：lock写入锁，独占，每次只允许一个goroutine写入对于key值
func (s SafeMap)Set(key string, val interface{})  {
	s.mu.Lock() // 上锁
	defer s.mu.Unlock() // 函数如何退出，解锁操作都会被执行
	if v, ok := val.(int); ok {
		s.items[key] = v
	}
}

// Get 读取：用读锁（共享，多个 goroutine 可同时读）
func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.RLock() // 添加读锁
	defer s.mu.RUnlock()
	v, ok := s.items[key]
	return v, ok
}

func (s *SafeMap) delete(key string)  {
	s.mu.Lock() // 添加写锁
	defer s.mu.Unlock()
	delete(s.items, key)
}

