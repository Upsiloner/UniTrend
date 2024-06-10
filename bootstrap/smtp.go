package bootstrap

import (
	"fmt"
	"log"
	"sync"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

type SMTPClientManager struct {
	server       *mail.SMTPServer
	client       *mail.SMTPClient
	clientMutex  sync.Mutex
	connectMutex sync.Mutex
}

func NewSMTPClientManager(env *Env) *SMTPClientManager {
	server := mail.NewSMTPClient()
	server.Host = env.SMTPHost
	server.Port = env.SMTPPort
	server.Username = env.SMTPUsername
	server.Password = env.SMTPPassword
	server.Encryption = mail.EncryptionSSL
	server.KeepAlive = true
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 3 * time.Second

	manager := &SMTPClientManager{
		server: server,
	}
	manager.connect()
	return manager
}

func (m *SMTPClientManager) connect() {
	m.connectMutex.Lock()
	defer m.connectMutex.Unlock()

	client, err := m.server.Connect()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to SMTP server: %v", err))
	}
	m.client = client
}

func (m *SMTPClientManager) GetClient() *mail.SMTPClient {
	// 保证只有一个协程访问client字段
	m.clientMutex.Lock()
	defer m.clientMutex.Unlock()

	// 检查客户端是否有效，如果无效则重新连接
	if err := m.client.Noop(); err != nil {
		log.Println("SMTP client connection lost, reconnecting...")
		m.connect()
	}
	return m.client
}

func (m *SMTPClientManager) SMTPClose() error {
	m.clientMutex.Lock()
	defer m.clientMutex.Unlock()

	if m.client != nil {
		return m.client.Close()
	}
	return nil
}
