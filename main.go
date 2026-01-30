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
		// Внутри цикла for ev := range evChan {

if ev.Kind == hook.KeyDown {
    if ev.Rawcode == 32 {  // пробел
        now := ev.When

        // Проверяем, не слишком ли быстро после предыдущего KeyDown (удержание/автоповтор)
        if wasPreviousSpace && now.Sub(lastSpaceTime) < 500*time.Millisecond {
            // Дополнительно: если интервал очень маленький — вероятно удержание, пропускаем
            if now.Sub(lastSpaceTime) > 60*time.Millisecond {  // ← настройка по предпочтениям (50–100 мс)
                fmt.Println("→ ДВОЙНОЙ ТАП: вставляем . ")

                // Удаляем два пробела
                robotgo.KeyTap("backspace")
                robotgo.MilliSleep(25)

                robotgo.KeyTap("backspace")
                robotgo.MilliSleep(25)

                // Вставляем ". " надёжно, независимо от раскладки
                robotgo.TypeStr(". ")

                wasPreviousSpace = false
                lastSpaceTime = time.Time{}
            } else {
                // Слишком быстро → вероятно удержание, просто обновляем время
                lastSpaceTime = now
            }
        } else {
            wasPreviousSpace = true
            lastSpaceTime = now
        }
    } else {
        // Другая клавиша → сброс
        wasPreviousSpace = false
    }
}

// Для сброса только KeyDown других клавиш

	}
}