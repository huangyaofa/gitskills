#GIT学习笔记
[toc]
##创建一个版本库
```
$ mkdir learngit
$ cd learngit
$ pwd
/Users/michael/learngit
```
>pwd命令用于显示当前目录
>如果你使用Windows系统，为了避免遇到各种莫名其妙的问题，请确保目录名（包括父目录）不包含中文。
##把文件添加到版本库
>现在我们编写一个readme.txt文件一定要放到learngit目录下（子目录也行），因为这是一个Git仓库，放到其他地方Git再厉害也找不到这个文件。
>把一个文件放到Git仓库只需要两步。第一步，用命令git add告诉Git，把文件添加到仓库：
```
$ git add readme.txt
```
>执行上面的命令，没有任何显示，这就对了，Unix的哲学是“没有消息就是好消息”，说明添加成功。
>第二步，用命令git commit告诉Git，把文件提交到仓库：
```
$ git commit -m "wrote a readme file"

```
>git commit命令，-m后面输入的是本次提交的说明，可以输入任意内容，当然最好是有意义的，这样你就能从历史记录里方便地找到改动记录。
>为什么Git添加文件需要add，commit一共两步呢？因为commit可以一次提交很多文件，所以你可以多次add不同的文件，比如：
```
$ git add file1.txt
$ git add file2.txt file3.txt
$ git commit -m "add 3 files."
```
>小结
>现在总结一下今天学的两点内容：
>初始化一个Git仓库，使用git init命令。
>添加文件到Git仓库，分两步：
>第一步，使用命令git add <file>，注意，可反复多次使用，添加多个文件；
>第二步，使用命令git commit，完成。
##更改并保存
>我们已经成功地添加并提交了一个readme.txt文件，现在，是时候继续工作了，于是，我们继续修改readme.txt文件，改成如下内容：
>Git is a distributed version control system.
>Git is free software.
>现在，运行git status命令看看结果：

```
$ git status
# On branch master
# Changes not staged for commit:
#   (use "git add <file>..." to update what will be committed)
#   (use "git checkout -- <file>..." to discard changes in working directory)
#
#    modified:   readme.txt
#
no changes added to commit (use "git add" and/or "git commit -a")
```
>git status命令可以让我们时刻掌握仓库当前的状态，上面的命令告诉我们，readme.txt被修改过了，但还没有准备提交的修改。
