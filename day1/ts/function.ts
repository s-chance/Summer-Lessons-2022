function foo(n: number) {
    console.log(n);
}
// 直接规定size的值只能为"large","small","middle"其中一个
function foo2(size: "large" | "small" | "middle") {
    console.log(size);
}

function foo3(size: number | string) {

}

foo(2);
// foo("2");

foo2("large");
// foo2("big");