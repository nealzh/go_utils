package ssh_utils

import (
	"bytes"
	"fmt"
	"go_utils/regx_utils"
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
	"time"
)

type ServerCmdInfo struct {
	Host        string
	Port        uint16
	UserName    string
	Passwd      string
	Cmd         string
	Timeout     time.Duration
	OutputLines *[]string
	Err         error
}

func getSshCmdResult(sci *ServerCmdInfo) (string, error) {

	//var hostKey ssh.PublicKey

	//创建sshp登陆配置
	sshConfig := &ssh.ClientConfig{
		Timeout: sci.Timeout, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:    sci.UserName,
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}

	sshConfig.Auth = []ssh.AuthMethod{ssh.Password(sci.Passwd)}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sci.Host, sci.Port)
	sshClient, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		log.Fatalln("get client", err)
		return "", err
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatalln("get session", err)
		return "", err
	}
	defer session.Close()

	//执行远程命令
	current_cmd := fmt.Sprintf("bash -l -c '%s'", sci.Cmd)
	//combo, err := session.CombinedOutput(current_cmd)
	//if err != nil {
	//	log.Fatalln("exec cmd", err)
	//	return "", err
	//}

	var buf bytes.Buffer
	session.Stdout = &buf
	if err := session.Run(current_cmd); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}

	return buf.String(), err
}

func GetSshCmdResultLines(sci *ServerCmdInfo) {

	resultStr, err := getSshCmdResult(sci)

	if err != nil {
		sci.Err = err
		return
	}

	var resultStrArray []string

	for _, lineStr := range regx_utils.LineBreakRegx.Split(resultStr, -1) {

		cleanLineStr := strings.Trim(lineStr, " ")

		if len(cleanLineStr) > 0 {
			resultStrArray = append(resultStrArray, cleanLineStr)
			//println("line: ", i, cleanLineStr)
		}
	}

	sci.OutputLines = &resultStrArray
	sci.Err = nil
}
