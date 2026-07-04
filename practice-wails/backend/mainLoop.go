package backend

import (
	"log"
	"time"
)

const (
	initialized = false
)

type MainLoopState int

const (
	MainLoopState_Uninitialized MainLoopState = iota
	MainLoopState_Running
)

// シングルトン
type MainLoop struct {
	// 状態を保持するフィールド
	state MainLoopState
}

// シングルトンのインスタンス
var mainLoop *MainLoop = nil

func sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func NewMainLoop() *MainLoop {
	if mainLoop != nil {
		return mainLoop
	}
	mainLoop = &MainLoop{state: MainLoopState_Uninitialized}
	return mainLoop
}

func (m *MainLoop) StartGoroutine() {
	m.state = MainLoopState_Running
	for i := 0; true; i++ {
		sleep(1) // 1秒待機
		log.Println("main loop running...", i)
		// メインループの処理をここに追加する
	}
}
