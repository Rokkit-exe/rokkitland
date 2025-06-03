package controller

import (
	"time"

	"github.com/Rokkit-exe/rokkitland/models"
	"github.com/Rokkit-exe/rokkitland/view"
)

type MainController struct {
	Renderer          *view.Renderer
	InputController   *InputController
	ConsoleController *ConsoleController
	State             *models.State
}

func NewMainController(state *models.State, stateController *StateController, consoleController *ConsoleController) MainController {
	return MainController{
		Renderer:          view.NewRenderer(state),
		InputController:   NewInputController(state, stateController, consoleController),
		ConsoleController: consoleController,
		State:             state,
	}
}

func (m *MainController) Start() {
	go m.inputLoop()
	m.renderLoop()
}

func (m *MainController) inputLoop() {
	for {
		err := m.InputController.RecordKeys() // works directly with state
		if err != nil {
			m.State.Console.Add("Input error: " + err.Error())
			break
		}
	}
}

func (m *MainController) renderLoop() {
	for {
		m.State.Mu.Lock()
		m.Renderer.Render()
		m.State.Mu.Unlock()
		time.Sleep(time.Second / 60) // ~60fps
	}
}
