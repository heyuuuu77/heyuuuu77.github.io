---
title: "Python 核心知识点"
date: 2025-01-15T11:30:03+00:00
draft: false
description: "Python 装饰器、迭代器、生成器等核心概念详解"
tags: ["Python", "装饰器", "迭代器", "生成器"]
categories: ["Python"]
author: "Heyuuuu"
---

#### Python的数据类型
- int: 整型
- float: 浮点型
- bool: 布尔值
- str: 字符串
- list: 列表，是一种可变的、有序的序列，可以包含不同类型的元素。使用方括号 [] 来定义。
- tuple: 元祖，是一种不可变的、有序的序列，一旦创建，其元素不能被修改。使用圆括号 () 来定义。
- set: 集合，是一种无序的、唯一的数据集合，不允许有重复的元素。使用花括号 {} 或 set() 函数来定义
- dict: 字典，是一种无序的键值对集合，每个键必须是唯一的，且键必须是可哈希的数据类型（如整数、字符串、元组等）。使用花括号 {} 来定义，键和值之间用冒号 : 分隔
- NoneType: 只有一个值 None，通常用于表示变量没有值或函数没有返回值。


#### List(列表) 与 Tuple(元祖) 的区别
List 可变，通过可以通过索引修改、添加或删除， 使用[]定义。 Tuple不可变，使用()定义。


#### 可迭代对象
包含 __iter__ 方法，能够被遍历的都可以称之为可迭代对象。 例如python的容器类数据类型：List, Tuple, Set, Dict。使用 `iter()` 函数可以将可迭代对象转换成迭代器。

#### 迭代器
包含 __next__ 魔法方法的对象。 对于可迭代对象，可以通过 `next()` 方法调用，每次调用，都会返回可迭代对象的下一个元素。直到没有元素返回时抛出StopIteration异常。

#### 生成器 
可以看成是一个特殊的迭代器。使用 yield 方法代替 return。 每次调用，都会记录函数状态并返回当前值。下次调用，会继续从纪录的地方继续执行。

#### 装饰器
包含wrapper方法。旨在不改变函数原有代码的前提下，修改函数/给函数添加功能。 实现原理是因为 将原函数当成参数传入装饰器函数。 涉及到一个概念： Python中，函数是一等公民.

#### 上下文管理器
可以理解成包含__enter__, __exit__魔法方法的对象。 通过with调用，在读取文件，数据库连接，锁的获取和释放中经常使用。作用是确保资源的正确使用和释放。例如：
```
    with open('example.json', 'r'):
        content = file.read()
        print(content)
```
open('example.json', 'r') 会创建一个文件对象，它就是一个上下文管理器。进入with语句时，文件被打开。当with语句结束时，文件会自动关闭。就算读取文件过程中发生异常，文件也会被正常关闭。
下面是数据库上下文管理器示例：
```
import sqlite3

class DatabaseConnection:
    def __init__(self, db_name):
        self.db_name = db_name
        self.connection = None


    def __enter__(self):
        self.connection = sqlite3.connect(self.db_name)
        return self.connection


    def __exit__(self, exc_type, exc_value, traceback):
        if self.connection:
            self.connection.close()


# 使用数据库连接上下文管理器
with DatabaseConnection('example.db') as conn:
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM users")
    rows = cursor.fetchall()
    print(rows)
```

contextlib 提供了一些创建上下文管理器的工具，使代码更简洁。使用 contextmanager 装饰器可以将一个生成器函数转换为上下文管理器：
```
from contextlib import contextmanager

@contextmanager
def simple_context_manager():
    print("进入上下文")
    try: 
        yield "在上下文中的对象"
    finally:
        print("退出上下文")
```



#### GIL
全局解释器锁。python 默认解释器(CPython)的一个全局锁。python解释器只能允许一个线程执行 Python字节码，在多核处理器环境下，Python多线程程序无法充分利用多个CPU的并行计算能力。

**作用**: 保证内存管理安全。Python的内存管理不是线程安全的。在多线程环境中，多个线程如果同时修改Python对象的内存，可能导致数据不一致，内存泄漏等问题。
GIL锁通过限制同一时刻只有一个线程执行Python代码，避免多个线程同时访问和修改Python对象的内存，从而保证了内存管理的安全性。

**解决方案**:
- 由于GIL的限制，无法真正发挥CPU的并行计算能力，解决办法是：使用多进程。 Python的multiprocessing模块可以创建多个进程，每个进程都有自己独立的Python解释器和GIL。因此可以充分利用CPU的并行计算能力。代码示例：
```python
import multiprocessing

#定义一个CPU密集型任务
def cpu_intensive_task():
    num = 0
    for i in range(10**7):
        num += 1

if __name__ == '__main__':
    # 创建两个进程
    process1 = multiprocessing.Process(target=cpu_intensive_task)
    process2 = multiprocessing.Process(target=cpu_intensive_task)

    # 启动进程
    process1.start()
    process2.start()

    # 等待程序执行完毕
    process1.join()
    process2.join()

    print("任务完成")

```

- 使用其他解释器：除了CPython, 其他诸如Jython, IronPython 解释器没有GIL锁，可以发挥多核的并行能力，但是可能存在兼容性问题
- 使用 C 扩展：对于一些挂件的CPU密集型代码， 可以使用 C 语言编写扩展模块，绕开 Python解释器。比如使用 Numpy，Pandas等科学计算库，这些库的底层代码通常是用 C 或 Fortran 编写的。