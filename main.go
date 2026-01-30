package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	robotgo.KeySleep = 10

	fmt.Println("Double Space → . + Space   (нажми Ctrl + Shift + Q для выхода)")
	fmt.Println("Работает...")

	evChan := hook.Start()
	defer hook.End()

	var lastSpaceTime time.Time
	var wasPreviousSpace bool

	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("Выход...")
		hook.End()
	})

	for ev := range evChan {
		if ev.Kind == hook.KeyDown {
			// Пробел на Windows → Rawcode 32
			if ev.Rawcode == 32 {
				fmt.Printf("Пробел пойман | Rawcode=%d  Keycode=%d  When=%v\n",
					ev.Rawcode, ev.Keycode, ev.When)

				now := ev.When

				if wasPreviousSpace && now.Sub(lastSpaceTime) < 500*time.Millisecond {
					// fmt.Println("→ УСЛОВИЕ СРАБОТАЛО: два быстрых пробела")

					// Замена: удаляем два пробела, вставляем . + space
					robotgo.KeyTap("backspace")
					robotgo.MilliSleep(20)

					robotgo.KeyTap("backspace")
					robotgo.MilliSleep(20)

					robotgo.KeyTap(".")
					robotgo.MilliSleep(20)

					robotgo.KeyTap("space")

					fmt.Println("→ Вставлена точка + пробел (если не видно — проверьте права или фокус)")

					wasPreviousSpace = false
					lastSpaceTime = time.Time{} // Обнуляем время
				} else {
					wasPreviousSpace = true
					lastSpaceTime = now
				}

				// Нe сбрасываем флаг здесь — continue не нужен, т.к. ниже else сбросит только для не-пробелов
			} else {
				// Сброс для любой другой KeyDown (не пробел)
				wasPreviousSpace = false
			}
		}

	}
}