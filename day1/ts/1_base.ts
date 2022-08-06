let a: number;
let b: boolean;
let d: string;
let e;
let f: any;

// 常用类型
// Array, Object, number, boolean, string, any, unknown
// ts规定变量的类型,赋值时只能用对应的类型赋值
a = 1; // ok
// a = "1"; //wrong


// any类型赋值没有类型限制,但不推荐使用
f = 1;
f = "131313";
