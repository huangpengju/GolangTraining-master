# GolangTraining-master
Golang知识训练合集

## GolangTraining-master （Golang训练-大师）
### 01_getting-started  
<details>
<summary>start包</summary>

##### fmt.Println()
##### fmt.Printf()↓
1. %d（十进制）  如：60 的十进制写法 60
2. %b（二进制）  如：60 的二进制写法 111100
3. %x（十六进制a-f）  如：60 的十六进制写法 3c
4. %X（十六进制A-F）  如：60 的十六进制写法 3C
5. %#x（0x开头，十六进制a-f） 如：60 的十六进制写法 0x3c
6. %#X（0X开头，十六进制a-f） 如：60 的十六进制写法 0X3c
7. %q（单引号括起来的go语法字符字面值） 如：60 的字符字面值法 '<'
</details>

### 02_package
<details>
<summary>amain包</summary>

#### 同一包下的函数调用方式及要求:
1. 调用：函数名()
2. 要求：需要注意此时对标识符（函数名）的首字母无要求
#### 不同包下的变量和函数调用方式及要求:
1. 调用：包名.函数名()
2. 要求：需要注意此时对标识符（函数名）的首字母需要大写
</details>

### 03_variables
<details>
<summary>variables包</summary>

#### 变量的声明和初始化:
1. 使用关键字 “var” 
2. “:=” 运算符
3. 变量的作用域（全局变量：在函数外部声明可以在多个函数内部使用；局部变量：在函数内部声明只在该函数内部使用）
#### 单引号、双引号、反引号:
1. 双引号""：里面可以是单个字符也可以是字符串，双引号里面可以有转义字符，如\n、\r等，对应go语言中的string类型
2. 单引号''：单引号在go语言中表示golang中的rune（int32）类型，单引号里面的单个字符，对应的是该字符的ASCII码
3. 反引号``：像双引号一样工作，但是对转义字符无效，内容按照原格式输出
</details>

### 04_scope
<details>
<summary>scope包</summary>

#### 变量和函数 在同一包和不同包中的使用:
1. 需要注意标识符（变量名和函数名）的首字母大小写问题
#### 变量在函数内部声明
1. 作用域：在该函数内部
#### 变量在函数中声明，在{}括号内声明
1. 函数中声明-作用域：在该函数内部
2. 函数内-新的{}中声明-作用域：在{}内部
#### 同一包下，不同.go文件中声明的全局变量
1. 作用域：该包下的所有.go文件中均可使用
2. 注意：因为是同一包下的.go文件使用，所以全局变量首字母无需大写
#### 匿名函数
1. 没有名称的函数
2. 匿名函数可以赋值给变量(函数表达式:将函数赋值给变量)
3. 闭包帮助我们限制多个函数使用的变量作用域，如果没有闭包，两个或多个函数要访问同一个变量，该变量需要是包作用域

#### 顺序很重要
1. 针对函数内部声明的变量而言，必须先声明，然后再使用
2. 针对函数内部使用函数外部声明的变量时，变量声明所在行（位置顺序）不重要,代码编译时，会优先编译全局变量，然后编译函数

#### 变量跟踪
1. 变量与函数同名，函数赋值给变量后，此时标识符表示的是变量。

</details>

### 05_blank-identifier
<details>
<summary>blank包</summary>

#### 占位符 _
1. 在go语言中，声明而不使用的变量，会报错
2. 下划线 _:是一个特殊的标识符，被用作占位符。
3. 在不需要使用变量的地方,可以使用下划线作为变量名,这样可以避免编译器报出未使用变量的警告,同时也表明了这个变量是不需要使用的。 
4. 下划线还可以用于多个返回值的函数或方法中,当我们只需要其中的某些返回值时,可以使用下划线来忽略不需要的返回值,而只关心需要的返回值。
</details>

### 06_constants
<details>
<summary>constants包</summary>

#### 常量
1. 常量是一个简单不变的值
2. 使用关键字const声明

#### iota 
1. iota是go语言的常量计数器,只能在常量表达式中使用，例如：const a = iota //0
2. iota在const关键字出现时将被重置为0,const中每新增一行常量声明将使iota计数一次 可理解为const语句块中的行索引。
例如：  
const a = iota //0  
const b = iota //0  
const (  
    c = iota // 0  
    d = iota // 1  
    e = iota // 2  
)

#### 左移运算符 <<
1. 左移运算符（<<）将一个运算对象的各二进制位全部左移若干位（左边的二进制位丢弃，右边的补0）  
例如：  
KB = 1 << (iota * 10) // 1 << (1 * 10)  
KB = 1<<10   // 1的二进制是 1 ，左边舍弃10位，右边补上10个0 得到：100 0000 0000，将100 0000 0000 转换为十进制是 1 * 2^10 = 1024  
KB= 1024
</details>

### 07_memory-address
<details>
<summary>address包</summary>

#### 内存地址
1. &：取地址符
2. 使用&获取变量的地址输出的是十六进制数
3. 内存地址的使用：fmt.Scan(&变量)，Scan从标准输入扫描文本，将成功读取的空白分隔的值保存进成功传递给本函数的参数
</details>

### 08_pointers
<details>
<summary>pointers包</summary>

#### 指针
1. &：取地址符
2. *：根据地址取值运算符（根据地址取出指向的值）
3. 总结：对变量进行取地址(&)操作，可以获得这个变量的地址（指针变量）
4. 总结：指针变量的值是原变量的地址
5. 总结：对指针变量进取取值(*)操作，可以获得指针变量指向的原变量的值。
6. Go语言指针学习地址：https://blog.csdn.net/weixin_44211968/article/details/121343717
</details>

### 09_remainder
<details>
<summary>remainder包</summary>

#### 取余数 %
1. 算数运算符 % 用来取余运算
</details>

### 10_for-loop
<details>
<summary>forloop包</summary>

#### for 循环
1. for循环可以初始化循环条件i,例如：for i:=0;i<=10;i++{}
2. for循环可以嵌套
#### for 循环的 While 用法
1. for循环可以写成while的方式(带循环条件)，例如：i:=0 for i<10{i++}
2. for循环可以写成while的方式(不带循环条件，变成死循环)，例如：i:=0 for {i++}
#### for 循环中使用 break 结束循环
1. for循环可以写成while的方式(不带循环条件，变成死循环，但使用break可以结束循环)，例如：i:=0 for {if i>=10{break} i++}
#### for 循环中使用 continue 跳出循环
1. for循环可以写成while的方式(不带循环条件，变成死循环，但使用break可以结束循环，加上continue可以跳过一次循环，开始下次个循环)，例如：```i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}```
#### for 循环中输出字符（ UTF-8是Go使用的文本编码方案）
1. "string()" 可以把int类型的参数转换为string类型，并且输出参数（数值）对应的ASCII码，UTF-8是Go使用的文本编码方案。
2. "[]byte()" 可以把string类型的参数转换为[]uint8类型的切片,比如，" []byte("T") "，输出为：" [84] "，数值84对应的ASCII码是T
</details>


### 11_switchstatements
<details>
<summary>switchstatements包</summary>

#### switch 判断
1. 语法：```switch 表达式{ case 条件1:  代码块  case 条件2:  代码块  default  代码块}```
2. fallthrough的用法：用于穿透switch，当switch中某个case匹配成功之后,就执行该case语句，如果遇到fallthrough,那么后面紧邻的case无需匹配, 执行穿透执行。fallthrough应该位于某个case的最后一行
3. case后面同时具有多个条件：```switch 表达式{ case 条件1,条件2:  代码块  case 条件3:  代码块  default  代码块}```
4. switch后面可以没有表达式：```switch { case 条件1:  代码块  case 条件2,条件3:  代码块  default  代码块}```
5. switch中允许使用变量的类型：```switch x.(type){//这是一个断言;断言"x是这种类型"  case int:  代码块  case string:  代码块  case 结构体类型  代码块  default  代码块}```
</details>


### 12_ifelse
<details>
<summary>ifelse包</summary>

#### true 为真 | !false 为真
1. 判断条件为真，就是true
2. if 条件为真，则执行紧邻if的{}中的代码

#### false 为假 | !true 为假
1. 判断条件为假，就是false
2. if 条件为假，则执行else的{}中的代码

#### 表达式中的初始化语句
1. 语法：``` b := true  if food := "apple"; b {  fmt.Println(food)  }  …… ``` 
2. 表达式中初始化的变量的作用域：变量food的作用域在 if else 分支语句中

#### 单分支
1. if

#### 双分支
1. if | else

#### 多分支
1. if | else if | else
2. if | else if | else if | else

</details>

### 13_exercise_solution
<details>
<summary>solution包</summary>

#### fmt.Println
1. fmt.Println() 返回2个结果,第1个结果表示写入标准输出的字节数,第2个结果表示写入标准输出的错误。
2. fmt.Println("Hello") 打印参数"Hello"
3. fmt.Println("Hello"[1]) 打印参数"Hello"的第2个元素e对应的ASCII码 101

#### fmt.Print 输出
1. 输出参数后，不换行

#### fmt.Scan 输入
1. 参数："&变量"
2. 可以通过终端输入，为变量赋值

#### 取余数 %
1. 例如：4%2 输出 0

####  if else 分支
1. 条件为真时，if
2. 条件为假时，else

#### for 
1. for 循环可以声明循环变量i:=0;加上循环条件i<=100;i++

#### += 累加运算
1. a:=0  a+=1  //a此时为1

</details>


### 14_function
<details>
<summary>function包</summary>

#### main() 
1. main()函数是程序的入口

#### 函数的参数
1. 参数的类型可以是任意类型
2. 参数的个数可以是1个，可以是2个
3. 参数的个数大于2个，并且是同类型时，可以省略第一个参数的类型，只写最后一个参数的类型
4. 函数的参数可以是可变的，语法：func test(sf ...int)，此时sf的类型是切片。在调用此函数时，需要传入多个参数test(1,2,3)或者传入一个切片，test([]{1,2,3}...)
5. 函数的参数可以是切片类型的，语法：func test(sf []int),调用方式test([]{1,2,3})

#### 函数的返回值
1. 函数可以返回一个指定类型的值
2. 函数的返回值，可以命名为一个变量，但是要避免这样使用
3. 函数可以返回多个值

#### 函数作为表达式
1. 函数或者匿名函数可以作为表达式，赋值给一个变量。 当函数或者匿名函数作为表达式赋值给一个变量时，这个变量的类型是：func()
2. 函数的返回值也可以是一个(匿名)函数类型

#### 闭包
1. 闭包帮助我们限制没有闭包的多个函数使用的变量的作用域，对于两个或多个函数访问同一个变量，该变量需要是包作用域
2. 闭包 = 函数 + 引用环境
2. 函数体内嵌套了另外一个函数，并且返回值是一个函数
3. 内层函数被暴露出去，被外层函数以外的地方引用着，形成了闭包。
4. 闭包每次返回都是一个新的实例，每个实例都有一份自己的环境。
5. 同一个实例多次执行，会使用相同的环境。
</details>