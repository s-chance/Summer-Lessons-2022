type Person = {
    name: string;
    age: number;
};

const p: Person = {
    name: "cx",
    // Name: "x",   名称不匹配
    age: 1,
    // weight: 132  出现了规定以外的属性
};

const people: Person[] = [
    {
        name: "mm",
        age: 999,
    },
];