## 执行一条SQL请求的过程是什么？
1. 连接器:建立连接，管理连接、校验用户身份;
2. 查询缓存:查询语句如果命中查询缓存则直接返回，否则继续往下执行。MySQL8.0已删除该模块
3. 解析 SQL，通过解析器对 SQL 查询语句进行词法分析、语法分析，然后构建语法树，方便后续模块读.取表名、字段、语句类型:
4. 执行 SQL:执行 SQL 共有三个阶段:
- 预处理阶段:检查表或字段是否存在;将select*中的*符号扩展为表上的所有列。
- 优化阶段:基于查询成本的考虑，选择查询成本最小的执行计划;
- 执行阶段:根据执行计划执行 SQL查询语句，从存储引擎读取记录，返回给客户端

## 事务的隔离级别有哪些？
1. 读未提交(read uncommitted)，指一个事务还没提交时，它做的变更就能被其他事务看到;
2. 读提交(read committed)，指一个事务提交之后，它做的变更才能被其他事务看到;
3. 可重复读(repeatable read)，指一个事务执行过程中看到的数据，一直跟这个事务启动时看到的数.据是一致的，MySQL InnoDB 引擎的默认隔离级别;
4. 串行化(serializable);会对记录加上读写锁，在多个事务对这条记录进行读写操作时，如果发生了读.写冲突的时候，后访问的事务必须等前一个事务执行完成，才能继续执行，

- 在「读未提交」隔离级别下，可能发生脏读、不可重复读和幻读现象;
- 在「读提交」隔离级别下，可能发生不可重复读和幻读现象，但是不可能发生脏读现象·
- 在「可重复读」隔离级别下，可能发生幻读现象，但是不可能脏读和不可重复读现象;
- 在「串行化」隔离级别下，脏读、不可重复读和幻读现象都不可能会发生。

## mysql的explain有什么作用？
- explain 是查看 sql的执行计划，主要用来分析 sql 语句的执行过程，比如有没有走索引，有没有外部排序，有没有索引覆盖等等。
如下图，就是一个没有使用索引，并且是一个全表扫描的查询语句。
![alt text](..\picture\mysql.png)
- 对于执行计划，参数有:
- - possible_keys 字段表示可能用到的索引,
- -key 字段表示实际用的索引，如果这一项为 NULL，说明没有使用索引:
- - key_len 表示索引的长度,
- - rows 表示扫描的数据行数。
- - type 表示数据扫描类型，我们需要重点看这个

## 给你张表，发现查询速度很慢，你有那些解决方案
1. 分析查询语句:使用EXPLAIN命令分析SQL执行计划，找出查询的原因，比如是否使用了全表扫描,是否存在索引未被利用的情况等，并根据相应情况对索引进行适当修改。
2. 创建或优化索引:根据查询条件创建合适的索引，特别是经常用于WHERE子句的字段、Orderby 排序的字段、Join 连表査询的字典、 group by的字段，并且如果查询中经常涉及多个字段，考虑创建联合索引，使用联合索引要符合最左匹配原则，不然会索引失效
3. 避免索引失效:比如不要用左模糊匹配、函数计算、表达式计算等等。
4. 查询优化:避免使用SELECT*，只查询真正需要的列;使用覆盖索引，即索引包含所有查询的字段;
联表查询最好要以小表驱动大表，并且被驱动表的字段要有索引，当然最好通过冗余字段的设计，避免联表查询。
5. 分页优化:针对 limit ny 深分页的查询优化，可以把Limit查询转换成某个位置的査询:select*from tb sku where id>20000 limit 10，该方案适用于主键自增的表,
6. 优化数据库表:如果单表的数据超过了千万级别，考虑是否需要将大表拆分为小表，减轻单个表的查询压力。也可以将字段多的表分解成多个表，有些字段使用频率高，有些低，数据量大时，会由于使用频率低的存在而变慢，可以考虑分开。
7. 使用缓存技术:引入缓存层，如Redis，存储热点数据和频繁查询的结果，但是要考虑缓存一致性的问题，对于读请求会选择旁路缓存策略，对于写请求会选择先更新 db，再删除缓存的策略。

## update 语句的执行过程。
![alt text](..\picture\update.png)