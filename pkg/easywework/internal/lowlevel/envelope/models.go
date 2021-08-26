package envelope

import (
	"encoding/xml"
)

type xmlRxEnvelope struct {
	ToUserName string `xml:"ToUserName"`
	AgentID    string `xml:"AgentID"`
	Encrypt    string `xml:"Encrypt"`
}

type cdataNode struct {
	CData string `xml:",cdata"`
}

type xmlTxEnvelope struct {
	XMLName      xml.Name  `xml:"xml"`
	Encrypt      cdataNode `xml:"Encrypt"`
	MsgSignature cdataNode `xml:"MsgSignature"`
	Timestamp    int64     `xml:"Timestamp"`
	Nonce        cdataNode `xml:"Nonce"`
}

type Envelope struct {
	ToUserName string
	AgentID    string
	Msg        []byte
	ReceiveID  []byte
}
