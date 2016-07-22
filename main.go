package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	
	"log"
	"net/http"
	"io"
	"os"
	"time"
	"encoding/binary"
	"bytes"
)

func wsH264(ws *websocket.Conn) {
	fmt.Printf("new socket\n")

	fi, err := os.Open("./test.h264")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
  
	msg := make([]byte, 1024*512)
	lenBytes := make([]byte, 4)
	for {
		time.Sleep(40 * time.Millisecond)
		
		lenNum, err := fi.Read(lenBytes)
		if (err != nil && err != io.EOF) || lenNum != 4 {
			log.Println(err)
			time.Sleep(1 * time.Second)
			break
		}
		
		b_buf := bytes.NewBuffer(lenBytes)
    var lenreal int32
    binary.Read(b_buf, binary.LittleEndian, &lenreal)
    
		
		n, err := fi.Read(msg[0:lenreal])
		if (err != nil && err != io.EOF) || n != int(lenreal) {
			log.Println(err)
			time.Sleep(1 * time.Second)
			break
		}
		
		err = websocket.Message.Send(ws, msg[:n])
		if err != nil {
			log.Println(err)
			break
		}
	}

	log.Println("send over socket\n")
}

func wsMpeg1(ws *websocket.Conn) {
	fmt.Printf("new socket\n")

	buf := make([]byte, 10)
	buf[0] = 'j'
	buf[1] = 's'
	buf[2] = 'm'
	buf[3] = 'p'
	buf[4] = 0x01
	buf[5] = 0x40
	buf[6] = 0x0
	buf[7] = 0xf0
	websocket.Message.Send(ws, buf[:8])

	fi, err := os.Open("./test.mpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	msg := make([]byte, 1024*1)
	for {
		time.Sleep(40 * time.Millisecond)
		n, err := fi.Read(msg)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if 0 == n {
			time.Sleep(1 * time.Second)
			break
		}
		err = websocket.Message.Send(ws, msg[:n])
		if err != nil {
		   log.Println(err)
		   break
		}
	}
	fmt.Printf("send over socket\n")	
}


func main() {
	http.Handle("/wsh264", websocket.Handler(wsH264))
	http.Handle("/wsmpeg", websocket.Handler(wsMpeg1))
	http.Handle("/", http.FileServer(http.Dir("./public")))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
