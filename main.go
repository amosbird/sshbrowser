package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/sys/windows/registry"
)

func params() string {
	if len(os.Args) == 1 {
		return ""
	}
	return os.Args[1]
}

func main() {
	pkey, err := ioutil.ReadFile("C:/Users/Administrator/.ssh/id_rsa")
	// pkey, err := ioutil.ReadFile("/home/amos/.ssh/id_rsa")
	if err != nil {
		panic("ioutil.ReadFile(sshkey):" + err.Error())
	}

	s, err := ssh.ParsePrivateKey(pkey)
	if err != nil {
		panic("ssh.ParsePrivateKey(): " + err.Error())
	}

	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	hostport, _, err := k.GetStringValue("CLIENTNAME")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ssh.Dial("tcp", hostport, &ssh.ClientConfig{
		User: "amos",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(s),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}

	defer session.Close()

	cmd := fmt.Sprintf("env DISPLAY=:0 /home/amos/scripts/luakit %s", params())
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(cmd); err != nil {
		panic("Failed to run: " + err.Error())
	}
}
