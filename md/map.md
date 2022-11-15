## map

+ map bucket 使用了链接为什么扩容的实际要搬迁数据？为了查找效率？方便计算 bucket 位置
+ map 最后一个溢出桶为什么要链接到第一个正常桶上？分配溢出桶时判断预分配溢出桶空间是否使用完
+ map 遍历为什么乱序？容量增加扩容会打乱原有的 bucket 顺序，每次遍历开始 bucket 不确定，如果map元素全部删除会重新生成hash种子，导致计算得到的hash值不一致
+ map key，val 中存的地址还是值？根据 key，val 类型而定
+ map val 是结构体时，赋值为什么不允许对结构体内部变量赋值？
### 使用

#### 声明/初始化
```golang 
    var m map[string]string // 声明 map，m = nil，执行删除、赋值操作会 panic
    var m = make(map[string]string) // 声明并初始化 map
    var m = make(map[string]string, 2) // 声明并预设置容量初始化 map
    var m = map[string]string{
        "name": "Jack",
    } // 声明并初始化内容 map
```

#### 添加元素
```golang 
   var m = make(map[string]string)
   m["name"] = "Jack"
```

#### 遍历
```golang 
    var m = make(map[string]string)
    for key,val := range m { // map 遍历顺序不可控，原因后续解释
    }
```

#### 获取单个元素
```golang 
    var m = map[string]string{
        "name": "Jack",
    }
    val := m["name"] // val = "Jack"，未添加元素获取时会返回 val 类型的零值
    val, ok := m["name"] // val = "Jack", ok = true
    val_, ok := ma["age"] // val = "", ok = false
```

#### 删除
```golang 
   var m = map[string]string{
        "name": "Jack",
   }
   delete(m, "name")
```

### 实现细节

#### 主要结构
```golang 
    // A header for a Go map.
    type hmap struct {
        // Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
        // Make sure this stays in sync with the compiler's definition.
        count     int // # live cells == size of map.  Must be first (used by len() builtin)
        flags     uint8
        B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
        noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
        hash0     uint32 // hash seed
    
        buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
        oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
        nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)
    
        extra *mapextra // optional fields
    }
    
    
    // A bucket for a Go map.
    type bmap struct {
        // tophash generally contains the top byte of the hash value
        // for each key in this bucket. If tophash[0] < minTopHash,
        // tophash[0] is a bucket evacuation state instead.
        tophash [bucketCnt]uint8
        // Followed by bucketCnt keys and then bucketCnt elems.
        // NOTE: packing all the keys together and then all the elems together makes the
        // code a bit more complicated than alternating key/elem/key/elem/... but it allows
        // us to eliminate padding which would be needed for, e.g., map[int64]int8.
        // Followed by an overflow pointer.
        
        // 编译时添加字段
        keys     [8]keytype
        values   [8]valuetype
		overflow uintptr
    }
    
```
#### 创建
    + 预分配溢出桶时和正常桶在连续内存空间上，预分配溢出桶 = 正常桶 / 16
#### 赋值
    + 赋值会触发 map 扩容，触发扩容后在每次删除和每次赋值会迁移两个 bucket 
#### 遍历
    + map 遍历是无序的。1.每次遍历随机选取起始 bucket，起始 bucket 下标开始遍历 2.扩容时 bucket 位置调整 3.元素全部删除后 hash 种子改变

#### PS
+ map 遍历是无序的，因为 map 会扩容，没办法保证有序
+ map value 是不可寻址类型， 因为 map 会扩容，返回地址，在扩容后原 key 所对应


#### 参考文档
源码位置 runtime/map.go