package libs

import (
	"context"
	"fmt"
	"gin_frame/config"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var EtcdCli *clientv3.Client

func init() {
	var err error
	EtcdCli, err = clientv3.New(clientv3.Config{
		Endpoints:   config.LoadConfig().Etcd.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Print(err)
	}
	//defer cli.Close()
}
func GetOne(key string) string {
	var timeout = time.Duration(10) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	values, err := EtcdCli.Get(ctx, key)
	if err != nil {
		return ""
	}
	var value string
	for _, ev := range values.Kvs {
		value = string(ev.Value)
		break
	}
	return value
}

func GetLock(key string, expire int64, maxLeaseCount int) bool {
	//获取key为空，则获取到锁，加锁，有值则返回false,加锁失败
	var (
		resp                   *clientv3.LeaseGrantResponse
		leaseKeepAliveChan     <-chan *clientv3.LeaseKeepAliveResponse
		err                    error
		leaseKeepAliveResponse *clientv3.LeaseKeepAliveResponse
		leaseCount             int
		lease                  clientv3.Lease
	)
	lease = clientv3.NewLease(EtcdCli)
	resp, err = lease.Grant(context.TODO(), expire)
	if err != nil {
		return false
	}
	var leaseId = resp.ID
	ctx, cancelFunc := context.WithCancel(context.TODO())
	if leaseKeepAliveChan, err = lease.KeepAlive(ctx, leaseId); err != nil {
		cancelFunc() //取消续租
		lease.Revoke(context.TODO(), leaseId)
		return false
	}

	//启动续租
	go func() {
		for {
			select {
			case leaseKeepAliveResponse = <-leaseKeepAliveChan:
				if leaseKeepAliveResponse != nil { //续租成功
					fmt.Print("续约:" + string(leaseCount))
					leaseCount++
					if leaseCount >= maxLeaseCount {
						fmt.Print("续约三次")
						cancelFunc() //取消续租
						lease.Revoke(context.TODO(), leaseId)
						goto END
					}
				} else { //续租失败
					fmt.Print("续约失败")
					cancelFunc() //取消续租
					lease.Revoke(context.TODO(), leaseId)
					goto END
				}
			}
		}
	END:
	}()

	kv := clientv3.NewKV(EtcdCli)
	//创建事务
	txn := kv.Txn(context.TODO())
	txn.If(clientv3.Compare(clientv3.CreateRevision("/lock/"+key), "=", 0)).
		Then(clientv3.OpPut("/lock/"+key, "1", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("/lock/" + key)) //否则抢锁失败

	//提交事务
	if txtResp, err := txn.Commit(); err != nil {
		fmt.Println(err)
		return true
	} else {
		//判断是否抢锁
		if !txtResp.Succeeded {
			cancelFunc() //取消续租
			lease.Revoke(context.TODO(), leaseId)
			fmt.Println("锁被占用：", string(txtResp.Responses[0].GetResponseRange().Kvs[0].Value))
			return false
		}
	}

	fmt.Println("处理任务")

	// time.Sleep(20 * time.Second)

	return true
}
func Unlock(key string) bool {
	return true
}
