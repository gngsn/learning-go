## Decorator

<small>🔗 <a href="https://youtu.be/4Oml8mbBXgo">Go 로 만드는 웹 - Web Handler</a></small>
<small>🔗 <a href="https://refactoring.guru/design-patterns/decorator/go/example">Decorator in Go</a></small>
<br/>

이 예제 프로젝트는 아래의 과정을 구현한다.

`Data -> Encrypt -> Zip -> Send`

`Data <- Decrypt <- UnZip <- Receive`

<br/><br/>

### Decorator Pattern ?

새로운 동작을 특수 Wrapper 안에 배치하여 동적으로 개체에 추가할 수 있는 구조적 패턴

Target Object와 Decorator가 모두 동일한 인터페이스를 따르기 때문에 Decorator를 사용하면 개체를 수없이 여러 번 포장할 수 있음. 

<br/>

이름 그대로 어떤 개체에 특정한 역할을 하도록 꾸미는 것

만약, `Data`를 내보내는데 **압축**을 하고싶다거나, **암호화**를 하고싶다거나, **로깅**을 하려고 한다고 가정해보자.

`Data+압축`이나, `Data+압축+암호화`, `Data+암호화` 등등... 어떤 기능을 추가한다고 해도 기존의 **Data(Target)**은 변하지 않음.

그래서 Decorator해줄 Wrapper를 만들어서 Data와 Decorator를 묶어주는 구조를 만들어 조합시기 위한 구조.



<br/>



<img src="https://user-images.githubusercontent.com/43839834/141096973-8429d5eb-e3e5-493d-8279-247a4ec6b70c.png" alt="decorator" style="zoom:70%;" />

<br/>



**압축 과정만 자세히 확인해보기**

``` go
type Component interface {
	Operator(string)
}
```

``` go
type SendComponent struct {}

// SendComponent의 Operator Method 구현
func (self *SendComponent) Operator(data string) {
	sentData = data
}
```

``` go
type ZipComponent struct {
	com Component
}

// ZipComponent의 ConcreteDecorator - Operator Method 구현
func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
  // Decorator가 갖고 있는 component의 Operator를 호출
	self.com.Operator(string(zipData))
}
```

<br/>

``` go
// ZipComponent - Decorator 
// SendComponent - Component
sender := &ZipComponent{ 
	com: &SendComponent{},
}

sender.Operator("Hello World")
```

