## ES6概要
### let、const、var关键字及一些符号的使用
1. let、const、var声明变量
let、const的作用域为块作用域,即该关键字所在的{}包围的范围内,var作用域则是整个文件
2. let和const区别
const声明的是常量、let声明的是变量;**可变的数值或布尔值使用let声明;对象、数组、字符串建议使用const来声明**,因为对象、数组、字符串在赋值时它们的地址本身就没有发生变化。另外,字符串在js中本身就不可变,要改变字符串内容只能再返回一个新字符串。
3. "$==$"和"$===$"的使用
由于js是弱类型语言,ES6标准中的"$===$"是比"$==$"更加严格的等于比较,只要类型不同则返回false,而"$==$"在比较时则会先进行一次类型转换再比较。**在对变量进行比较时建议用"$===$"**
4. "!!"的作用将任意变量转换为布尔值
### 解构赋值
1. 数组解构、对象解构
```javascript
// 数组解构
let [a, b, c] = [1, 2, "cxnb"];
相当于
/*
let a = 1;
let b = 2;
let c = "cxnb"; 
*/
另外下面的写法也没有问题
// let [, b, c] = [1, 2, "cxnb"];
// let [a, b] = [1, 2, "cxnb"];

// 对象 (Object) 解构
let person = {
  name: "cx",
  sex: {
    child: "male",
    adult: "unknow",
  },
};
// 不解构获取person对象中sex属性中的adult属性的值
const person_sex1 = person.sex.adult;

// 一次解构获取person对象中sex属性,变量名要和属性名一致
const { sex } = person; // const { sex: sex } = person;的简写形式

// 为防止重名,可自定义变量名
// const { sex: sex1 } = person;
// console.log(sex1);
// const { sex: sex2 } = person;
// console.log(sex2);

// 二次解构获取sex属性中的adult属性的值,需要先有一次解构
const { adult: person_sex2 } = sex;
```
2. 函数参数解构
```javascript
const fruit = {
  apple: true,
  peach: false,
  grape: true,
};

const animal = ["pdx", "sd"];

function add([{ apple, grape }, visitor]) {
  console.log(apple, grape);
  console.log(visitor[0]);
}

add([fruit, animal]); // 传入一个数组
// 调用add传递参数的过程相当于
// const [{ apple, grape }, visitor] = [fruit, animal];
// 相当于
// const { apple, grape } = fruit;
// const visitor = animal;

// 原始方式
// function add(obj) {
//   console.log(obj[0].apple, obj[0].grape);
//   console.log(obj[1][0]);
// 相比之下, 解构方式的代码可读性更强
```
### 三目运算符与模板字符串
```javascript
// 三目运算符
// 返回布尔值的语句 ? true时执行的语句 : false时执行的语句;

// 模板字符串
console.log("apple is " + apple ? "good" : "bad"); // 会省略掉"apple is "
// 可以修改为以下两种语句
console.log(apple ? "apple is good" : "apple is bad");
console.log(`apple is ${apple ? "good" : "bad"}`);
```
### 函数参数默认值
```javascript
// 调用函数时未规定x的参数则默认函数参数x为2, 否则由规定的x的参数覆盖默认值 
function foo(x = 2, y) {
    console.log(x, y);
}

foo(1, 2); // 1 2
foo(2); // 2 undefined
// 另外在调用函数时使用赋值表达式无法设定函数参数默认值
foo((x=1)); // 2 undefined
foo((y=2)); // 2 undefined
```
### 箭头函数定义方式
```javascript
const obj = {
    // 常规函数定义方式
    f0: function() {},
    // 简写方式省略:和function
    f1() {
        console.log("hello");
        console.log(this);
        // this指向obj的地址
    },

    // 箭头函数
    f2: () => {
        console.log("hello");
        console.log(this);
        // 箭头函数中中的this无法指向obj
    },
};

obj.f1();
obj.f2();

function func() {};

const func1 = () => {};

function func2() {};
```
### js对象操作
```javascript
let obj;
const x = 1;
const y = 2;

obj = {x123: x, y456: y};
console.log(obj);
//访问成员属性,[]中的字符可拼接对应具体的对象, 以此实现动态访问
console.log(obj["x" + "123"]);
console.log(obj["y" + "456"])
```
### 扩展运算符"..."
```javascript
const a = [1, 2, 3];
const b = [3, 4, 5];
const c = a.concat(b);
// concat函数将a、b数组拼接成一个新数组
console.log(c);
// ...和[]组合:...理解为把数组中每个元素取出来,[]则把这些元素包起来重新形成一个新数组
const d = [...a, ...b];
console.log(d);
// Set集合不允许重复元素
const e = new Set([...a, ...b]);
console.log(e);
// 利用...和[]将Set集合对象转换为数组
console.log([...e]);

const obj1 = {
  a: 3,
  c: 5,
};
const obj2 = {
  a: 2,
  b: 2,
};
// 将两个对象组合成一个对象,重复的元素会覆盖
const obj3 = { ...obj1, ...obj2 };
console.log(obj3);
```
### 遍历
```javascript
const arr = [12, 1, 5, 2, 1];
// 普通for循环可以通过break中途退出循环
for (let i = 0; i < arr.length; i++) {
  console.log(arr[i]);
  // if (i > 2) break;
}

// 箭头函数 const fn = () => {};
arr.forEach((item, index) => {
  item++; // item表示的是一个数组元素的值,而不是元素本身, index表示索引(下标)
  console.log(index, item);
});
console.log(arr); // 数组元素没有改变

// map相对forEach,能够在修改完数组后返回一个新的数组
const arr2 = arr.map((item) => item + 1);
console.log(arr2); // 新数组,数组元素理论上并没有被修改,而是新生成的

// 遍历的是对象数组,则可根据地址对值修改
const objArr = [{ a: 1 }, { a: 2 }];
objArr.forEach((item) => (item["b"] = 2)); // item表示的是对象数组元素的地址,这里新增了一个属性
console.log(objArr);
```
### filter筛子
```javascript
const arr = [1, 2, 3, 4, 5, 6];
// filter类似于map返回了一个新的数组
// 但相对原数组,map一般是改变了数组的值,而filter是改变了数组元素的个数
const after = arr.filter((item) => {
  if (item % 2 === 1) return true;
  // return item % 2 === 1;
});

// const after = arr.filter((item) => item % 2 === 1);

console.log(after);
```