var a = 1
a = "hello"
console.log(a)
//console.log(s+"nihao")
console.log("hello world")
console.log("你好")
console.log(Number.MAX_VALUE)
console.log("32" - 32)
console.log("32" + 32)
console.log(Math.pow(2, 3))
console.log(Math.PI)
var x = NaN
console.log(x != x)
var y = 0.3 - 0.2
var z = 0.2 - 0.1
console.log(y == z)

var date = new Date(2019, 3, 21)
console.log(date.getMonth())
console.log(date.getDate())
var s = "nihao"
console.log(s.length)
console.log(parseInt("10"))
console.log(parseInt("10adsa"))
console.log(parseInt("g10"))
console.log(parseInt("1fdsf"))
console.log(parseInt("10.98"))

console.log(Number("10"))
console.log(Number("10adsa"))
console.log(Number("g10"))
console.log(Number("110.98fdsf"))
console.log(Number("10.98"))

console.log(Boolean(1))
console.log(Boolean(0))
console.log(Boolean(""))
console.log(Boolean("nihao"))
console.log(Boolean(null))
console.log(Boolean(undefined))

var num1 = 10;
var num1 = 10;
var num1 = 10;
console.log(num1.toString());
console.log(String(10));
//如果变量有意义，调用.toString()
//如果变量没有意义，使用String()
var num1 = 20;
num1++;
++num1;
num1 += 1;
console.log(num1)

console.log(1 | 2)

var arr = new Array(5); //单个参数表示长度
arr[0] = 1;
arr[1] = "he";
console.log(arr)

var arr1 = []
console.log(arr1)

var arr2 = new Array(10, 20, 30, 40, 50)//多个参数表示元素
console.log(arr2)

var arr3 = [10,30,40,60,1,2,4]
console.log(arr3)



function sum(x, y) {
    console.log(arguments.length)
    console.log(arguments[0])
    console.log(arguments[1])
    return x+y;
}

var result = sum(10,20)
console.log(result)


var f = function(x, y) {
    return x + y
}
console.log(f(1,2));

//函数自调用
(function() {
    console.log("被调用了");
})();
dog()
function dog () {
    num11 = 22
}

console.log(num11)
//除了函数中定义的变量外，其他位置的都是全局变量

f1()
console.log(c1)
console.log(b1)
//console.log(a1)

function f1() {
    var a1 = b1 = c1 = 9
    //var al
    //b1 = 9
    //c1 = 9
    console.log(a1)
    console.log(b1)
    console.log(c1)
}

//对象
var obj = new Object();
obj.name = "xiaomi"
obj.age = 38
obj.sex = "女"

obj.eat = function() {
    console.log("我喜欢吃")
}

obj.play = function() {
    console.log("我喜欢玩")
}

console.log(obj.name)

//自定义构造函数
function Person(name, age) {
    this.name = name
    this.age = age
    this.eat = function(){
        console.log("我在吃饭")
    }
}

var p1 = new Person("xiaoming", 20)
p1.eat()

/**
 * 在内存中开辟空间，存储创建的新对象
   把this设置为当前对象
   设置对象的属性和方法
   把this这个对象返回
 */
var p2 = {
    name:"xiaoli",
    age:18,
    eat: function(){

    },
    play: function(){

    }
}
p2.address = "zhejiang"
console.log(p2.name)
console.log(p2.address)

var json = {
    "name":"小米",
    "age":"10",
    "sex":"男"
}
//遍历 json
for(var key in json) {
    console.log(json[key])
}

var dt = new Date();
console.log(dt.getDay())
console.log(dt.toString())
console.log(dt.toLocaleString())
console.log(dt.toDateString())

var ss = new String("hello")
console.log(ss.lastIndexOf("l"))
var s2 = "hello"
console.log(s2.lastIndexOf("l"))

const s3 = "hello"
console.log(s3)