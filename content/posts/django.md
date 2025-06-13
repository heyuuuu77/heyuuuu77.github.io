+++
date = '2025-02-05T11:09:45+08:00'
draft = false
title = 'Django的常见问题'
tags = ['Django']
TocOpen = true
+++


#### 什么是ORM

ORM 即对象关系映射，它将数据库中的表映射为 Python 类，将表中的记录映射为类的实例，通过操作类和实例来操作数据库。优点包括提高开发效率、可移植性强、安全性高（自动处理 SQL 注入）等。

#### 什么是 N+1问题

Django 的 ORM（对象关系映射）为开发者提供了便捷的数据库操作方式，允许通过 Python 对象来操作数据库。然而，当进行关联查询时，如果处理不当，就容易引发 N + 1 查询问题。其核心原因在于，在进行多次查询时，没有一次性获取所有关联数据，而是先执行一次查询获取主对象列表，然后针对每个主对象，再分别执行一次查询来获取其关联对象，这就导致查询次数过多，性能下降。

简单来说就是多表关联的时候，如果涉及到一对多的关联关系， 查询完成“一”之后，在遍历这个“一”获取其关联的多条数据时， 就面临了N的问题。

django中提供了 select_related() 和 prefetch_related()

- select_related()


前者主要用来处理 一对一/外键关联关联(ForeignKey和OneToOneField), 它通过sql中的join一次性查询所有相关数据，代码如下
```py
# 由于这里是反向查询，select_related 不适用，下面只是示例其使用场景
# 如果是正向查询，比如从 Book 查询 Author 可以用 select_related
books = Book.objects.select_related('author').all()
for book in books:
    print(f"{book.title} 的作者是 {book.author.name}")
```

- prefetch_related()

prefetch_related() 用于处理多对多和反向关联（如上述示例中从 Author 查询 Book），它通过多次查询并在 Python 层面合并数据。修改最初的代码如下:

```py
authors = Author.objects.prefetch_related('books').all()
for author in authors:
    book_count = author.books.count()
    print(f"{author.name} 有 {book_count} 本书。")

```
这里，prefetch_related('books') 会先执行一次查询获取所有作者，再执行一次查询获取所有相关的书籍，然后在 Python 层面将它们关联起来。总共只执行了 2 次查询，而不是 N + 1 次查询，有效提高了性能。


#### 关联模式


