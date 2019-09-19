函数名

作用

复杂度

zslCreateNode

新建并返回一个跳表节点

O(1)

zslCreate

新建并初始化一个跳跃表

O(L)

zslFreeNode

释放给定的节点

O(1)

zslFree

释放给定的跳跃表

O(N)

zslRandomLevel

得到新节点的层数(抛硬币法的改进)

O(1)

zslInsert

将给定的score与member新建节点并添加到跳表中

O(logN)

zslDeleteNode

删除给定的跳表节点

O(L)

zslDelete

删除给定的score与member在跳表中匹配的节点

O(logN)

zslIsInRange

检查跳表中的元素score值是否在给定的范围内

O(1)

zslFirstInRange

查找第一个符合给定范围的节点

O(logN)

zslLastInRange

查找最后一个符合给定范围的节点

O(logN)

zslDeleteRangeByScore

删除score值在给定范围内的节点

O(logN)+O(M)

zslDeleteRangeByRank

删除排名在给定范围内的节点

O(logN)+O(M)

zslGetRank

返回给定score与member在集合中的排名

O(logN)

zslGetElementByRank

根据给定的rank来查找元素

O(logN)
