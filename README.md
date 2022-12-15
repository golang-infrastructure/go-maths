# Go Maths 

修改自己的定位，只提供一些helper方法，不做比较专业的优化，搞不了...

# API

泛型的类型约束依赖：
```
https://github.com/golang-infrastructure/go-gtypes
```

## 绝对值

对数字求绝对值，输入类型和输出类型一致，不需要强制转换类型

```go
func Abs[T gtypes.Number](v T) T
```

## 平均值

对数字类型的切片数组求平均值，结果以float64类型返回

```go
func Avg[T gtypes.Number](slice []T) float64
```

对数字类型的切片数组求平均值，结果以指定的类型返回，这个R类型要支持从float64直接强转成功，否则会panic

```go
func AvgR[T, R gtypes.Number](slice []T) R 
```

## 相等

比较两个数字切片是否相等 

```go
func Equals[T1, T2 gtypes.Number](a T1, b T2) bool
```

比较两个浮点数是否相等：

```go
func FloatEquals[T gtypes.Float](a, b T) bool
```

## Float

把64位的浮点型数据转为uint64数据

```go
func Float64SliceToUint64Slice(float64Slice []float64) []uint64
```

把64位的无符号数组转为浮点型数组

```go
func Uint64SliceToFloat64Slice(uint64Slice []uint64) []float64
```



```go
func Float32SliceToUint32Slice(float32Slice []float32) []uint32
```



```go
func Uint32SliceToFloat32Slice(uint32Slice []uint32) []float32
```

## 求最值

最大值：

```go
func Max[T gtypes.Number](values ...T) T 
```

最小值：

```go
func Min[T gtypes.Number](values ...T) T
```

是否等于0：

```go
func IsZero[T gtypes.Number](value T) bool
```















