---
title: "设计模式详解"
date: 2025-02-05T10:00:19+08:00
draft: false
description: "23种设计模式详解：创建型、结构型、行为型"
tags: ["设计模式", "单例模式", "工厂模式"]
categories: ["软件工程"]
author: "Heyuuuu"
---


设计模式是软件开发中针对各种反复出现的问题归纳出的通用的解决方案，总共有23种设计模式，分别分成创建型，结构性和行为型三大类

#### 创建型
1. 单例模式
确保一个类只有一个实例，并提供一个全局访问点来访问这个实例。
Python中的实现方式是通过__new__魔术方法实现。
```python

class Singleton:
    _instance = None

    def __new__(cls, *args, **kwargs):
        if not cls._instance:
            cls._instance = super().__new__(cls, *args, **kwargs)
        return cls._instance

# 使用示例
s1 = Singleton()
s2 = Singleton()
print(s1 is s2)  # 输出: True


```

2. 工厂模式

工厂模式可细分为简单工厂模式、工厂方法模式和抽象工厂模式。


- **简单工厂模式**

定义一个工厂类，该类有一个创建对象的方法，通过传入不同的参数来创建不同类型的对象。
```python
class Animal:
    def speak(self):
        pass

class Dog(Animal):
    def speak(self):
        return "Woof!"

class Cat(Animal):
    def speak(self):
        return "Meow!"

class AnimalFactory:
    @staticmethod
    def create_animal(animal_type):
        if animal_type == 'dog':
            return Dog()
        elif animal_type == 'cat':
            return Cat()
        return None

# 使用示例
dog = AnimalFactory.create_animal('dog')
print(dog.speak())  # 输出: Woof!
```

- **工厂模式**

将对象的创建延迟到子类中进行。定义一个创建对象的抽象方法，由具体的子类来实现该方法，从而创建具体的对象

```python
class Animal:
    def speak(self):
        pass

class Dog(Animal):
    def speak(self):
        return "Woof!"

class Cat(Animal):
    def speak(self):
        return "Meow!"

class AnimalFactory:
    def create_animal(self):
        pass

class DogFactory(AnimalFactory):
    def create_animal(self):
        return Dog()

class CatFactory(AnimalFactory):
    def create_animal(self):
        return Cat()

# 使用示例
dog_factory = DogFactory()
dog = dog_factory.create_animal()
print(dog.speak())  # 输出: Woof!
```

- **抽象工厂模式**

提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类
```py
class Pet:
    def play(self):
        pass

class Guard:
    def protect(self):
        pass

class DogPet(Pet):
    def play(self):
        return "Dog is playing."

class DogGuard(Guard):
    def protect(self):
        return "Dog is guarding."

class CatPet(Pet):
    def play(self):
        return "Cat is playing."

class CatGuard(Guard):
    def protect(self):
        return "Cat is guarding (weakly)."

class AnimalFactory:
    def create_pet(self):
        pass

    def create_guard(self):
        pass

class DogFactory(AnimalFactory):
    def create_pet(self):
        return DogPet()

    def create_guard(self):
        return DogGuard()

class CatFactory(AnimalFactory):
    def create_pet(self):
        return CatPet()

    def create_guard(self):
        return CatGuard()

# 使用示例
dog_factory = DogFactory()
dog_pet = dog_factory.create_pet()
dog_guard = dog_factory.create_guard()
print(dog_pet.play())  # 输出: Dog is playing.
print(dog_guard.protect())  # 输出: Dog is guarding.
```



#### 结构型

1. 代理模式
允许通过代理对象来控制另一个对象的访问。以下是一个简单的python实现：
```python
# 定义主题接口
class Subject:
    def request(self):
        pass

# 定义真实主题类
class RealSubject(Subject):
    def request(self):
        print("Real Subject: handling request.")


# 定义代理类
class Proxy(Subject):
    def __init__(self. real_subject):
        self.real_subject = real_subject

    def request(self):
        # 在调用真实主题的方法之前可以进行一些预处理
        print("Proxy: logging the request before forwarding")
        self.real_subject.request()
        print("Proxy: logging the request after forwarding")


# 客户端代码
if __name__ == '__main__':
    real_subject = RealSubject()
    proxy = Proxy(real_subject)
    proxy.request()

```

python没有java中的动态代理机制。但是可以通过Python的动态特性是的在运行时创建对象和修改对象的行为。以下是两种常见的方式:

第一种装饰器模式: 
```python
def proxy_decorator(func):
    def wrapper(*args, **kwargs):
        print("Proxy: before func call")
        result = func(*args, **kwargs)
        print("Proxy: after func call")
        return result
    
    return wrapper

# 定义一个目标函数
@proxy_decorator
def target_function():
    print("Targe func is running.")

```

第二种 __getattr__: 
```python
class DynamicProxy:
    def __init__(self, target):
        self.target = target

    def __getattr__(self, name):
        def wrapper(*args, **kwargs):
            print(f"Proxy: Before call {name}")
            result = getattr(self.target, name)(*args, **kwargs)
            print(f"Proxy: After call {name}")

            return result
        return wrapper

# 定义一个目标类
class Target:
    def method(self):
        print("Target method is running")

# 创建目标对象和代理对象
target = Target()
proxy = DynamicProxy(target)

# 通过代理对象调用目标方法
proxy.method()
.
```

2. 装饰器模式

允许向一个现有的对象添加新的功能，同时又不改变其结构，在Python中，装饰器是一种特殊函数，接受一个函数作为参数，并返回一个新的函数。新函数会在原函数基础上添加额外的功能。

基础的装饰器示例:
```python
# 定义一个装饰器函数
def my_decorator(func):
    def wrapper():
        print("原函数之前进行操作")
        func()
        print("原函数之后进行操作")

    return wrapper

# 定义一个普通函数
@my_decorator
def say_hello():
    print("Hello")

# 调用被装饰后的函数
say_hello()

```

**装饰 带参数的函数**
```python
def decorator_with_args(func):
    def wrapper(*args, **kwargs):
        print("Before")
        result = func(*args, **kwargs)
        print("After")
        return result
    
    return wrapper

@decorator_with_args
def add_numbers(a, b):
    return a + b


# 调用被装饰带有参数的函数
result = add_numbers(3, 5)
print(result)
```


**带参数的装饰器**
```python
def repeat(n):
    def decorator(func):
        def wrapper(*args, **kwargs):
            for _ in range(n):
                result = func(*args, **kwargs)
            return result
        return wrapper
    return decorator

@repeat(3)
def greet():
    print("Hi")

# 调用 带参数的装饰器
greet()

```

**类装饰器**
```python
class MyDecorator:
    def __init__(self, func):
        self.func = func

    def __call__(self, *args, **kwargs):
        print("类装饰器, Before")
        result = self.func(*args, **kwargs)
        print("类装饰器，After")
        return result

@MyDecorator
def print_message():
    print("这是一条消息")

# 调用被类装饰器装饰的函数
print_message()

```


3. 组合模式

它允许你将对象组合成树形结构以表示"部分-整体"的层次关系。使用组合模式，客户端代码可以统一处理单个对象和对象组合，无需区分它们。 这种模式是的客户端代码可以一致的处理对象的组合结构，而不必关心具体是单个对象还是对象组合。
从而简化客户端代码的编写。

组合模式通常有以下几个角色:
- 抽象组件(Component): 定义了组合中对象的公共接口，声明了用于管理子组件和访问子组件的方法，单个对象和组合对象都要实现这个接口。
- 叶子组件(Leaf): 表示组合中的叶子结点
- 复合组件(Composite): 表示组合中的分支节点，包含子节点，可以是叶子结点或其他复合组件。 它实现了抽象组件接口中管理子组件的方法。


代码示例:
```python
# 抽象组件类
class FileSystemComponent:
    def display(self, depth=0):
        pass

    def add(self, component):
        pass

    def remove(self, component):
        pass

    def get_child(self, index):
        pass

# 叶子组件类：文件
class File(FileSystemComponent):
    def __init__(self, name):
        self.name = name

    def display(self, depth=0):
        print("  " * depth + self.name)

# 复合组件类：文件夹
class Folder(FileSystemComponent):
    def __init__(self, name):
        self.name = name
        self.children = []

    def display(self, depth=0):
        print("  " * depth + self.name)
        for child in self.children:
            child.display(depth + 1)

    def add(self, component):
        self.children.append(component)

    def remove(self, component):
        if component in self.children:
            self.children.remove(component)

    def get_child(self, index):
        if 0 <= index < len(self.children):
            return self.children[index]
        return None

# 客户端代码
if __name__ == "__main__":
    # 创建文件夹和文件
    root_folder = Folder("Root")
    sub_folder1 = Folder("Subfolder 1")
    sub_folder2 = Folder("Subfolder 2")
    file1 = File("File 1")
    file2 = File("File 2")
    file3 = File("File 3")

    # 构建文件系统结构
    sub_folder1.add(file1)
    sub_folder1.add(file2)
    sub_folder2.add(file3)
    root_folder.add(sub_folder1)
    root_folder.add(sub_folder2)

    # 显示文件系统结构
    root_folder.display()
    
```




#### 获取用户详情接口
