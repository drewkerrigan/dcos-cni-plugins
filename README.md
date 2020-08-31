我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# Purpose
This repo hosts the CNI plugins specific to DC/OS. A brief description
of each of the plugins is given below:

* [l4lb](l4lb/README.md): A CNI plugin which allows containers in isolated virtual networks to use services provided by [Minuteman](https://github.com/dcos/minuteman) and [Spartan](https://github.com/dcos/spartan).

# Pre-requisites
* GoLang 1.6+
* [Glide](https://github.com/Masterminds/glide)

# Build instructions
The project uses `glide` to capture all the golang dependencies into the vendor folder. In order to build the project just do
```
make
```
To do a clean build run:
```
make clean
```
The build installs the plugins in the `bin` folder.

