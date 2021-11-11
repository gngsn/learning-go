## Maps

Map은 키(Key)에 대응하는 값(Value)을 신속히 찾는 해시테이블(Hash table)을 구현한 자료구조. 



### 정의

> map[Key type] Value type



**예시** 

``` go 
var m map[string]string
// OR
var dictionary = map[string]string{}
// OR
var dictionary = make(map[string]string)
```

<br/>

**key type !**

`integers, floating point, complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs, arrays`

<br/><br/>

### comma ok

> val, err = map_name[key]

<br/>

위와 같이 map의 key값을 접근하면 2개의 인자가 반환 됨.

첫 번째는 key에 해당하는 값

두 번째는 key값에 해당하는 데이터가 존재하는지 존재하지 않은지에 대한 boolean값!

그래서 두 번째 인자를 comma ok라고 관용적으로 부름 ~

<br/>

**Example**

``` go
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
	  ...
}

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
}
```





### 삭제

> 

