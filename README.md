# goperf
TCP、UDP、SCTP Performance Testing Tools
# 指标
> 1. https://community.f5.com/kb/technicalarticles/understanding-performance-metrics-and-network-traffic/286109
> 2. https://help.aliyun.com/zh/slb/classic-load-balancer/user-guide/faq-about-clb-2?spm=a2c4g.11186623.0.0.7dc13c96Rncdgh#concept-umd-czv-tdb

| 指标                            | 说明   |
|---------------------------------|------------|
| BPS(bits per-second)            | 每秒bit数，即带宽 |
| Latency                         | 时延，单个请求响应所消耗的时间 |
| PPS(packet per-second)          | 每秒数据包数 |
| CPS(connections per-sencond)    | 每秒新建链接数 |
| MCC(max concurrent connections) | 最大并发连接数 |

## 带宽（bps）

## 时延（latency）

## 每秒数据包数（pps）

## 每秒新建链接数（cps）
1. 启动线程池
2. 每间隔周期(t1-ms,t1<1s>)建立特定数量(n)链接
3. 每个链接维持特定时间(t2-ms)，与服务端完成一次通信(tcp/udp/sctp)
4. 调整周期(100ms)，如果未出现链接失败数(f1)处于可增强范围(0~1%)，下一周期n增加特定步长(s1);如果链接失败数高于特定值(3%)则需要抑制，下一周期n减小特定步长(s2)；如果链接失败数处于性能边缘(1~3%)，则n不变，认为是一个极限的状态，继续维持m1个周期后，增加特定步长(s3)。
5. 链接失败数处于性能边缘的数据统计，则为最终结果集；收集连续的m2个周期，取平均数。

核心思想：通过s1、s2快速逼近大致的值，通过s3微调到一个较大的稳定值

## 最大并发链接数（mcc）
1. 启动线程池
2. 每间隔周期(t1-ms,t1>300s>)建立特定数量(n)链接
3. 每个链接维持特定时间(t2-ms)，与服务端完成一次通信(tcp/udp/sctp)
4. 调整周期(100ms)，如果出现失败，则下个周期n减少特定步长(s1)；如果未出现失败，则下个周期n增加特定步长(s2)；直到n=0(即一个都创建不出来)



