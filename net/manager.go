/**
 * User: Jackong
 * Date: 13-11-25
 * Time: 下午8:26
 */
package net

type manager struct {
	connections map[string]*Conn
}
var (
	Anonymous *manager
	Signed *manager
)

func init() {
	Anonymous = &manager{connections: make(map[string]*Conn)}
	Signed = &manager{connections: make(map[string]*Conn)}
}

func (this *manager) Get(id string) *Conn {
	return this.connections[id]
}

func (this *manager) Put(id string, conn *Conn) {
	this.connections[id] = conn
}

func (this *manager) Del(id string) {
	delete(this.connections, id)
}

func (this *manager) CloseAll() {
	for id, conn := range this.connections {
		conn.Close()
		this.Del(id)
	}
}

func (this *manager) Close(id string) {
	if _, ok := this.connections[id]; ok {
		this.connections[id].Close()
		this.Del(id)
	}
}

func SignIn(conn *Conn, sid string) bool {
	if conn == nil {
		return false
	}
	conn.IsSigned = true
	conn.Id = sid
	Signed.Put(sid, conn)
	Anonymous.Del(conn.RemoteAddr().String())
	return true
}

func SignOut(conn *Conn) bool {
	if conn == nil {
		return false
	}
	Signed.Del(conn.Id)
	conn.IsSigned = false
	conn.Id = conn.RemoteAddr().String()
	Anonymous.Put(conn.Id, conn)
	return true
}

func Close(conn *Conn) {
	if conn.IsSigned {
		Signed.Close(conn.Id)
	} else {
		Anonymous.Close(conn.Id)
	}
}
