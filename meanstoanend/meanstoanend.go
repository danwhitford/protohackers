package meanstoanend

import (
	"encoding/binary"
	"log"
	"net"
)

type message struct {
	instruction byte
	a, b        int32
}

func HandleFunc(c net.Conn) {
	defer c.Close()
	prices := make(map[int32]int32)

	for {
		buffer := make([]byte, 9)
		for i := 0; i < 9; i++ {
			b := make([]byte, 1)
			_, err := c.Read(b)
			if err != nil {
				log.Println(err)
				return
			}
			buffer[i] = b[0]
		}

		msg := parseMessage(buffer)
		switch msg.instruction {
		case 'Q':
			mean := handleQuery(prices, msg)
			out := make([]byte, 4)
			binary.BigEndian.PutUint32(out, uint32(mean))
			c.Write(out)
		case 'I':
			handleInsert(prices, msg)
		default:
			log.Println("got a unknown")
			return
		}
	}
}

func handleQuery(prices map[int32]int32, msg message) int32 {
	if msg.a > msg.b {
		return 0
	}
	start := msg.a
	end := msg.b
	var total int64
	var count int64
	for t, p := range prices {
		if start <= t && t <= end {
			total += int64(p)
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return int32(total / count)
}

func handleInsert(mp map[int32]int32, msg message) {
	mp[msg.a] = msg.b
}

func parseMessage(b []byte) message {
	var msg message
	msg.instruction = b[0]
	msg.a = uncomplement(b[1:5])
	msg.b = uncomplement(b[5:9])
	return msg
}

func uncomplement(b []byte) int32 {
	i := binary.BigEndian.Uint32(b)
	ii := int32(i)
	return ii
}
