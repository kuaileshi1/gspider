// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/26 10:59 下午
package task

type Status int

const (
	StatusUnknown Status = iota
	StatusRunning
	StatusStopped
	StatusCompleted
	StatusUnExceptedExited
)

var StatusMessage = map[Status]string{
	StatusUnknown:          "未知状态",
	StatusRunning:          "运行中",
	StatusStopped:          "停止",
	StatusCompleted:        "完成",
	StatusUnExceptedExited: "异常退出",
}

func GetStatusMessage(status Status) string {
	if v, ok := StatusMessage[status]; ok {
		return v
	}

	return ""
}
