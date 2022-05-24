package msg

import ()

func CreateMsg(username string) string {

	body_msg := "hello" + username + "\n" + "we welcome you all to this family"

	return body_msg
}
