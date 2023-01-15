package mq

import "testing"

func TestCreateConnUrl(t *testing.T) {
	// Create connection url
	connUrl := CreateConnUrl("guest", "guest", "localhost", "5672", "/")
	if connUrl != "amqp://guest:guest@localhost:5672/" {
		t.Errorf("CreateConnUrl() = %s; want amqp://guest:guest@localhost:5672/", connUrl)
	}
}

func TestConnectToMq(t *testing.T) {
	connUrl := CreateConnUrl("guest", "guest", "localhost", "5672", "/")
	conn, err := ConnectToMq(connUrl)
	if err != nil {
		t.Errorf("ConnectToMq() = %s", err)
	}
	defer conn.Close()
}

func TestCreateChannel(t *testing.T) {
	connUrl := CreateConnUrl("guest", "guest", "localhost", "5672", "/")
	conn, err := ConnectToMq(connUrl)
	if err != nil {
		t.Errorf("ConnectToMq() = %s", err)
	}
	defer conn.Close()
	ch, err := CreateChannel(*conn)
	if err != nil {
		t.Errorf("CreateChannel() = %s", err)
	}
	defer ch.Close()
}

func TestDeclareQueue(t *testing.T) {
	connUrl := CreateConnUrl("guest", "guest", "localhost", "5672", "/")
	conn, err := ConnectToMq(connUrl)
	if err != nil {
		t.Errorf("ConnectToMq() = %s", err)
	}
	defer conn.Close()
	ch, err := CreateChannel(*conn)
	if err != nil {
		t.Errorf("CreateChannel() = %s", err)
	}
	defer ch.Close()
	q, err := DeclareQueue(*ch, "test")
	if err != nil {
		t.Errorf("DeclareQueue() = %s", err)
	}
	defer ch.QueueDelete(q.Name, false, false, false)
}

func TestCloseChannel(t *testing.T) {
	connUrl := CreateConnUrl("guest", "guest", "localhost", "5672", "/")
	conn, err := ConnectToMq(connUrl)
	if err != nil {
		t.Errorf("ConnectToMq() = %s", err)
	}
	defer conn.Close()
	ch, err := CreateChannel(*conn)
	if err != nil {
		t.Errorf("CreateChannel() = %s", err)
	}
	defer ch.Close()
	CloseChannel(*ch)
}

func TestCloseConnection(t *testing.T) {
	connUrl := CreateConnUrl("guest", "guest", "localhost", "5672", "/")
	conn, err := ConnectToMq(connUrl)
	if err != nil {
		t.Errorf("ConnectToMq() = %s", err)
	}
	defer conn.Close()
	CloseConnection(*conn)
}

func TestPublishMessage(t *testing.T) {
	connUrl := CreateConnUrl("guest", "guest", "localhost", "5672", "/")
	conn, err := ConnectToMq(connUrl)
	if err != nil {
		t.Errorf("ConnectToMq() = %s", err)
	}
	defer conn.Close()
	ch, err := CreateChannel(*conn)
	if err != nil {
		t.Errorf("CreateChannel() = %s", err)
	}
	defer ch.Close()
	q, err := DeclareQueue(*ch, "test")
	if err != nil {
		t.Errorf("DeclareQueue() = %s", err)
	}
	defer ch.QueueDelete(q.Name, false, false, false)
	//PublishMessage(*ch, q.Name, "test message")
}

func TestConsumeMessage(t *testing.T) {
	connUrl := CreateConnUrl("guest", "guest", "localhost", "5672", "/")
	conn, err := ConnectToMq(connUrl)
	if err != nil {
		t.Errorf("ConnectToMq() = %s", err)
	}
	defer conn.Close()
	ch, err := CreateChannel(*conn)
	if err != nil {
		t.Errorf("CreateChannel() = %s", err)
	}
	defer ch.Close()
	q, err := DeclareQueue(*ch, "test")
	if err != nil {
		t.Errorf("DeclareQueue() = %s", err)
	}
	defer ch.QueueDelete(q.Name, false, false, false)
	//PublishMessage(*ch, q.Name, "test message")
	//ConsumeMessage(*ch, q.Name)
}
