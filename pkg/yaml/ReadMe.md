# go解析yaml

# yaml基础语法

## 基本规则
1. 大小写敏感
2. 使用缩进表示层级关系
3. 禁止使用 tab 缩进，只能使用空格键
4. 缩进长度没有限制，只要元素对齐就表示这些元素属于一个层级
5. 使用 # 表示注释
6. 字符串可以不用引号标注

## 三种数据结构
1. map   
使用 : 表示键值对，同一缩进的所有键值对属于一个map
```$xslt
# yaml表示
age : 20
name : name1
# 或者将一个map写在一行

# 对应json表示
{'age':20, 'name':'name1'}
```

2. list
使用 - 表示
```$xslt
# yaml表示
- a
- b
- 20
# 或者写在一行
[a,b,20]

# 对应json表示
['a','b',20]
```

3. scala
数据最小的单位，不可以再分割

## 数据结构嵌套
map和list元素可以是另一个map或者list或是scalar。有以下几种嵌套
1. map 嵌套 map
```$xslt
# yaml表示
websites:
 YAML: yaml.org 
 Ruby: ruby-lang.org 
 Python: python.org 
 Perl: use.perl.org 
 
# 对应json表示
{ websites: 
   { YAML: 'yaml.org',
     Ruby: 'ruby-lang.org',
     Python: 'python.org',
     Perl: 'use.perl.org' } }
```

2. map 嵌套 list
```$xslt
# yaml表示
languages:
 - Ruby
 - Perl
 - Python 
 - c
 
# 对应json表示
{ languages: [ 'Ruby', 'Perl', 'Python', 'c' ] }
```

3. list 嵌套 list
```$xslt
# yaml 表示
-
  - Ruby
  - Perl
  - Python 
- 
  - c
  - c++
  - java
# 或者
- - Ruby
  - Perl
  - Python 
- - c
  - c++
  - java
  
# 或者
- [Ruby,Perl,Python]
- [c,c++,java]
  
# 对应 json 表示
[ [ 'Ruby', 'Perl', 'Python' ], [ 'c', 'c++', 'java' ] ]
```

4. list 嵌套 map
```$xslt
# yaml 表示
-
  id: 1
  name: huang
-
  id: 2
  name: liao
  
# 对应json表示
[ { id: 1, name: 'huang' }, { id: 2, name: 'liao' } ]
```

# references
[将yaml转成json](http://nodeca.github.io/js-yaml/)