package client

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"time"
)

type Client struct {
	client *clientv3.Client
}

func (c *Client) Close() {
	c.client.Close()
}

func New(n []string) (*Client, error) {
	cfg := clientv3.Config{
		Endpoints: n,
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    10 * time.Second,
		DialKeepAliveTimeout: 3 * time.Second,
	}
	c, err := clientv3.New(cfg)
	if err != nil {
		return &Client{}, err
	}
	return &Client{c}, nil
}

func Put(c *Client, k string, v string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := c.client.Put(ctx, k, v)
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			fmt.Println("ctx is canceled by another routine")
		case context.DeadlineExceeded:
			fmt.Println("ctx is attached with a deadline and it exceeded")
		case rpctypes.ErrEmptyKey:
			fmt.Println("client-side error: key is not provided")
		default:
			fmt.Println("bad cluster endpoints, which are not etcd servers")
		}
		return err
	}
	fmt.Println("key:", k, "value:", v)
	return nil
}
