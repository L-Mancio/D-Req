package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"net"
	"os"
)

type reqstructureGui struct {
	DownloadName string
	ReqAuthor    string
	Priority     int //number 1-5
}

type HoldPriorityOptions struct {
	rd1 *walk.RadioButton
	rd2 *walk.RadioButton
	rd3 *walk.RadioButton
	rd4 *walk.RadioButton
	rd5 *walk.RadioButton

	hidden_msg *walk.TextLabel
}

func runUI(requests []reqstructureGui) {
	var inputName, inputDwnldName *walk.TextEdit
	mw := &HoldPriorityOptions{}
	_, _ = MainWindow{
		Title:  "D-REQ",
		Size:   Size{300, 400},
		Layout: VBox{},

		Children: []Widget{
			TextLabel{
				Text: "Your name, always use the same one:",

				RowSpan: 2,
			},
			HSplitter{
				MaxSize: Size{20, 20},

				Children: []Widget{
					TextEdit{AssignTo: &inputName},
				},
			},
			TextLabel{
				Text: "Download Name, be specific:",
			},

			HSplitter{

				MaxSize: Size{20, 20},

				Children: []Widget{
					TextEdit{AssignTo: &inputDwnldName},
				},
			},
			TextLabel{
				Text: "Priority, 5->Very urgent:",
			},
			RadioButtonGroup{

				Buttons: []RadioButton{
					{
						Alignment: 1,
						AssignTo:  &mw.rd1,
						Name:      "optionOne",
						Text:      "1",
						Value:     1,
					},
					{
						Alignment: 1,
						AssignTo:  &mw.rd2,
						Name:      "optionTwo",
						Text:      "2",
						Value:     2,
					},
					{
						Alignment: 1,
						AssignTo:  &mw.rd3,
						Name:      "optionThree",
						Text:      "3",
						Value:     3,
					},
					{
						Alignment: 1,
						AssignTo:  &mw.rd4,
						Name:      "optionFour",
						Text:      "4",
						Value:     4,
					},
					{
						Alignment: 1,
						AssignTo:  &mw.rd5,
						Name:      "optionFive",
						Text:      "5",
						Value:     5,
					},
				},
			},

			PushButton{

				Text: "Send",

				MaxSize: Size{Width: 50, Height: 50},
				OnClicked: func() {
					var options []bool
					options = append(options, mw.rd1.Checked(), mw.rd2.Checked(), mw.rd3.Checked(), mw.rd4.Checked(), mw.rd5.Checked())
					var priority int
					for i, rb := range options {
						if rb == true {
							priority = i + 1
						}
					}
					requests = append(requests, reqstructureGui{
						DownloadName: inputDwnldName.Text(),
						ReqAuthor:    inputName.Text(),
						Priority:     priority,
					})
					//fmt.Println(requests[0].Priority,requests[0].ReqAuthor, requests[0].DownloadName)

					buf := new(bytes.Buffer)
					enc := gob.NewEncoder(buf)
					if err := enc.Encode(requests); err != nil {
						fmt.Println(err)
					}
					requestInByteFormat := buf.Bytes()

					servAddr := "192.168.1.51:4444"
					tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
					if err != nil {
						println("ResolveTCPAddr failed:", err.Error())
						os.Exit(1)
					}
					conn, err := net.DialTCP("tcp", nil, tcpAddr)
					if err != nil {
						println("Dial failed:", err.Error())
						os.Exit(1)
					}

					_, err = conn.Write(requestInByteFormat)

					if err != nil {
						println("Write to server failed:", err.Error())
						os.Exit(1)
					}

					println("Sending to server ", servAddr)

					reply := make([]byte, 1024)
					_, err = conn.Read(reply)
					if err != nil {
						println("Write to server failed:", err.Error())
						os.Exit(1)
					}

					_ = conn.Close()
					requests = append(requests[1:])
					mw.hidden_msg.SetVisible(true)
					_ = inputDwnldName.SetText("")
					_ = inputName.SetText("")
					mw.rd1.SetChecked(false)
					mw.rd2.SetChecked(false)
					mw.rd3.SetChecked(false)
					mw.rd4.SetChecked(false)
					mw.rd5.SetChecked(false)

				},
			},

			TextLabel{
				Text:     "Request Sent",
				AssignTo: &mw.hidden_msg,
				Visible:  false,
			},
		},
	}.Run()
}
func main() {
	var reqArray []reqstructureGui

	runUI(reqArray)

}
