package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// 1、实现3节点选举
// 2、改造代码成为分布式选举代码，加入RPC调用
// 3、演示完整代码，自主选举，日志复制

// 定义3节点
const raftCount = 3

// 声明leader对象
type Leader struct {
	// 任期
	Term int
	// Leader Id编号
	LeaderId int
}

// 声明raft
type Raft struct {
	// 锁
	mu sync.Mutex
	// 节点编号
	me int
	// 当前任期
	currentTerm int
	// 为那个节点投票
	votedFor int
	// 3个状态
	// 0 follower 1 candidate 2 leader
	state int
	// 发送最后一条数据的时间
	lastMessageTime int64
	// 设置当前节点的领导
	currentLeader int
	// 节点间发送信息的通道
	message chan bool
	// 选举的通道
	electCh chan bool
	// 心跳信号通道
	heartBeat chan bool
	// 返回心跳信号的通道
	heartbeatRe chan bool
	// 超时时间
	timeout int
}

// 0 没有上任，-1 没有编号
var leader = Leader{0, -1}

func main() {
	// 过程： 有3个节点，最初为follower
	// 若有candidate状态，进行投票拉票
	// 会产生leader

	// 创建3个节点
	for i := 0; i < raftCount; i++ {
		// 创建3个raft节点
		Make(i)
	}
	// 不能让main结束
	for {

	}
}

func Make(me int) *Raft {
	fmt.Println("+++++Make+++++start......")
	rf := &Raft{}
	rf.me = me
	// -1 代表谁都不投，此时节点刚创建
	rf.votedFor = -1
	// 0 follower
	rf.state = 0
	rf.timeout = 0
	rf.currentLeader = -1
	// 节点任期
	rf.setTeam(0)
	rf.message = make(chan bool)
	rf.electCh = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartbeatRe = make(chan bool)
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 选举的协程
	go rf.election()
	// 心跳检测的协程
	go rf.sendLeaderHearBeat()
	return rf
}

// 设置节点任期
func (rf *Raft) setTeam(term int) {
	rf.currentTerm = term
}

func (rf *Raft) election() {
	// 设置标记，判断是否选出了leader
	var result bool
	for {
		// 设置超时，150到300随机数
		timeout := randRange(150, 300)
		rf.lastMessageTime = millisecond()
		fmt.Println("开始election-timeout...", timeout, rf.me)
		select {
		// 延迟毫秒
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("当前节点状态为：", rf.state)
		}
		result = false
		for !result {
			// 选主逻辑
			result = rf.election_on_round(&leader)
		}
	}
}

// 实现选主逻辑
func (rf *Raft) election_on_round(leader *Leader) bool {
	fmt.Println("开始election_on_round...", rf.me)
	// 定义超时
	var timeout int64 = 100
	// 投票数量
	var vote int
	// 定义是否开始心跳信号的产生
	var triggerHeartbeat bool
	// 时间
	last := millisecond()
	// 用于返回值
	success := false
	// 给当前节点变成candidate
	rf.mu.Lock()
	// 修改状态
	rf.becomeCandidate()
	rf.mu.Unlock()
	fmt.Println("start electing leader")
	for {
		// 遍历所有节点拉选票
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				// 拉选票
				go func() {
					if leader.LeaderId < 0 {
						// 设置投票
						rf.electCh <- true
					}
				}()
			}
		}
		// 设置投票数量
		vote = 1
		// 遍历
		for i := 0; i < raftCount; i++ {
			// 计算投票数量
			select {
			case ok := <-rf.electCh:
				if ok {
					// 投票数量加1
					vote++
					// 若选票个数，大于节点个数/2，则成功
					success = vote > raftCount/2
					if success && !triggerHeartbeat {
						// 变化成主节点，选主成功了
						// 开始触发心跳信号检测
						triggerHeartbeat = true
						rf.mu.Lock()
						// 变主
						rf.becomeLeader()
						rf.mu.Unlock()
						// 由leader向其他节点发送心跳信号
						rf.heartBeat <- true
						fmt.Println(rf.me, "号节点成了leader")
						fmt.Println("leader开始发送心跳信号了")
					}
				}
			}
		}
		// 做最后的校验工作
		// 若不超时，且票数大于一半，则选举成功，break
		if timeout+last < millisecond() || (vote > raftCount/2) || rf.currentLeader > -1 {
			break
		} else {
			// 等待操作
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return success
}

// 修改状态candidate
func (rf *Raft) becomeCandidate() {
	rf.state = 1
	rf.setTeam(rf.currentTerm + 1)
	rf.votedFor = rf.me
	rf.currentTerm = -1
}

// 修改状态Leader
func (rf *Raft) becomeLeader() {
	rf.state = 2
	rf.currentLeader = rf.me
}

// leader节点发送心跳信号
// 随便完成数据同步
// 看小弟挂没有挂
func (rf *Raft) sendLeaderHearBeat() {
	// 死循环
	for {
		select {
		case <-rf.heartBeat:
			rf.sendAppendEntriesImpl()
		}

	}
}

// 用于返回给leader的确认信号
func (rf *Raft) sendAppendEntriesImpl() {
	// 是主就别跑了
	if rf.currentLeader == rf.me {
		var success_count = 0
		// 此时是leader
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				go func() {
					rf.heartbeatRe <- true
				}()
			}
		}
		// 计算返回确认信号个数
		for i := 0; i < raftCount; i++ {
			select {
			case ok := <-rf.heartbeatRe:
				if ok {
					success_count++
					if success_count > raftCount/2 {
						fmt.Println("投票选举成功，心跳信号ok")
						log.Fatal("程序结束")
					}
				}
			}
		}
	}

}

func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// 获取当前时间，发送最后一条数据的时间
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
