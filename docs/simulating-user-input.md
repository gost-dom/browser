# Simulating user input

> [!NOTE]
>
> This is currently WIP.

User input currently consist of simulating mouse clicks and keyboard
interaction. These interactions triggers default behaviour in many elements.

## `KeyboardController`

To simulate user input, use the `KeyboardController` type. Example:

```Go
b := browser.New()
win := b.Open("http://example.com")
input := win.Document().QueryElement("#fullname-input")
input.(html.HTMLElement).Focus()

ctrl := KeyboardController{Window: win}
ctrl.SendKeys(keys.StringToKeys("John Smith"))
assert.Equal(t, "John Smith", input.Value())
```

Input is directed goes to the currently focused element. here [Element.focus] is
explicitly called to bring focus to the element.

[Element.focus]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus

> [!IMPORTANT]  
> 
> It is currently clear that the design is not sufficient, as modifier keys have
> not been taken into consideration. So the design _might_ change. But the
> operation `SendKeys(key.StringToKeys("input"))` should stay valid as a whole.

## Misssing features

Keyboard simulation is in a very early stage. 

- Modifier keys
- Moving focus on <kbd>tab</kbd>/<kbd>shift</kbd>+<kbd>tab</kbd>.
- Clicking buttons on <kbd>space</kbd>.
- Submitting forms on <kbd>enter</kbd>.
