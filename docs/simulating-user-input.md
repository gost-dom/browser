# Simulating user input

> [!NOTE]
>
> This is currently WIP.

User input currently consist of simulating mouse clicks and keyboard
interaction. These interactions triggers default behaviour in many elements.

## Input controller

Keyboard input is simulated through the `KeyboardController` type. This is
associated with a single window, and it will simulate what should happen when
the user types on the keyboard, and the window has input focus in the OS.

Input is directed goes to the currently focused element. To type a value in a
text field, you must first move focus to the text field, e.g. by simulating a
click, or simply calling [Element.focus]

[Element.focus]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus

### Sending a sequence of keys

To simulate the user typing, call `KeybaordConroller.SendKeys()`. Use
`key.StringToKeys()` to generate the proper sequence from a string.

> [!IMPORTANT]  
> 
> It is currently clear that the design is not sufficient, as modifier keys are
> not simulated properly. So the design _might_ change, but the operation
> `SendKeys(key.StringToKeys("input"))` should stay valid as a whole.

A key press is represented by the `Key` type, but generally you 
calling `SendKey`.

A sequence of keys can be simulated using the experimental `SendKeys`, accepting
an `iter.Seq[Key]` as input. `KeysOfString` can be used to create a sequence of
keys from a `string` input.


## Misssing features

Keyboard simulation is in a very early stage. 

- Modifier keys
- Moving focus on <kbd>tab</kbd>/<kbd>shift</kbd>+<kbd>tab</kbd>.
- Clicking buttons on <kbd>space</kbd>.
- Submitting forms on <kbd>enter</kbd>.
