## chan

不要通过共享内存来通信，而应该通过通信来共享内存

### 创建
```
    var ch = make(chan int) // 无缓冲 chan
    var ch = make(chan int, 2) // 有缓冲 chan
```
### 使用
  #### 写入
  ```
    ch <- 1
  ```
  #### 读取
  ```
    for val := range ch {  
    }
    
    _ = <- ch
    _, _ = <- ch 
    
    select {
        case <-ch:
    }
  ```

### 错误
+ 无缓冲 chan 使用 for range 方式读取需要在读取后关闭 chan，否则会产生 'fatal error: all goroutines are asleep - deadlock!' 错误
+ 同一协程读写无缓冲 chan 会产生 'fatal error: all goroutines are asleep - deadlock!' 错误
+ chan 先读取再写入会产生 'fatal error: all goroutines are asleep - deadlock!' 错误
+ 向已关闭的 chan 内写入数据会产生 'panic: send on closed channel' 错误
+ 重复关闭 chan 会产生 'panic: close of closed channel' 错误

### 注意事项
+ 向已关闭协程读取数据会返回协程类型零值
+ 当 chan 为空时，读取会一直阻塞
+ chan 关闭应该由写入方操作

### 底层实现
  #### 创建
  调用 make 方法初始化时，分为三种不同的情况申请内存
  + 无缓冲 chan 申请了一块连续内存存储 runtime.hchan 结构，并将 buf 指向自己
  + 有缓冲 chan，chan 类型不包含指针对象，也申请一块连续内存存储 runtime.hchan 结构和 buf 数据，前面一部分存储 hchan，后面存储 buf
  + 有缓冲 chan，chan 类型包含执行对象，buf 和 hchan 分别存储

  #### 写入
  + 写入时有等待的接收方，将直接将值写入接收方，不需要写入缓存
  + 写入时执行值拷贝操作，而不是直接使用原值地址
  + 缓存 buf 循环使用
  + select 操作 buf 区满的时候直接返回，其他情况会将其加入 send 队列
  + 会唤醒 recvq 中的协程（缓冲区已满）

  #### 读取
  + 无缓冲读取 send 队列
  + 有缓冲 chan 读取 buf 数据后会将 send 队列的数据添加到 buf
  + 会唤醒 sendq 中的协程（缓存区已满）
 
