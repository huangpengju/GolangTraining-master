# GolangTraining-master
Golang知识训练合集

## GolangTraining-master （Golang训练-大师）
### 01_getting-started  
<details>
<summary>start包</summary>
主要练习↓

##### fmt.Println()
##### fmt.Printf()↓
%d（十进制）  %b（二进制）  %x（十六进制a-f）  %X（十六进制A-F）  %#x（0x开头，十六进制a-f） %#X（0X开头，十六进制a-f） %q（单引号括起来的go语法字符字面值） 
</details>

### 02_package
<details>
<summary>amain包</summary>
主要练习↓

#### 同一包下的函数调用方式及要求 ↓
标识符首字母无需大写
#### 不同包下的变量和函数调用方式及要求 ↓
标识符首字母需要大写
</details>
</details>

### 02_variables
<details>
<summary>variables包</summary>
主要练习↓

#### 变量的声明和初始化
1. 使用关键字 “var” 
2. “:=” 运算符
3. 变量的作用域（全局变量：在函数外部声明可以在多个函数内部使用；局部变量：在函数内部声明只在该函数内部使用）
#### 单引号、双引号、反引号
1. 双引号""：里面可以是单个字符也可以是字符串，双引号里面可以有转义字符，如\n、\r等，对应go语言中的string类型
2. 单引号''：单引号在go语言中表示golang中的rune（int32）类型，单引号里面的单个字符，对应的是该字符的ASCII码
3. 反引号``：像双引号一样工作，但是对转义字符无效，内容按照原格式输出
</details>

