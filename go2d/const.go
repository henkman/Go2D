package go2d

import "github.com/papplampe/Go2D/sdl"

const (
	BLENDMODE_NONE  = sdl.BLENDMODE_NONE
	BLENDMODE_BLEND = sdl.BLENDMODE_BLEND
	BLENDMODE_ADD   = sdl.BLENDMODE_ADD
	BLENDMODE_MOD   = sdl.BLENDMODE_MOD
)

const (
	KEY_UNKNOWN            = sdl.KEY_UNKNOWN
	KEY_A                  = sdl.KEY_A
	KEY_B                  = sdl.KEY_B
	KEY_C                  = sdl.KEY_C
	KEY_D                  = sdl.KEY_D
	KEY_E                  = sdl.KEY_E
	KEY_F                  = sdl.KEY_F
	KEY_G                  = sdl.KEY_G
	KEY_H                  = sdl.KEY_H
	KEY_I                  = sdl.KEY_I
	KEY_J                  = sdl.KEY_J
	KEY_K                  = sdl.KEY_K
	KEY_L                  = sdl.KEY_L
	KEY_M                  = sdl.KEY_M
	KEY_N                  = sdl.KEY_N
	KEY_O                  = sdl.KEY_O
	KEY_P                  = sdl.KEY_P
	KEY_Q                  = sdl.KEY_Q
	KEY_R                  = sdl.KEY_R
	KEY_S                  = sdl.KEY_S
	KEY_T                  = sdl.KEY_T
	KEY_U                  = sdl.KEY_U
	KEY_V                  = sdl.KEY_V
	KEY_W                  = sdl.KEY_W
	KEY_X                  = sdl.KEY_X
	KEY_Y                  = sdl.KEY_Y
	KEY_Z                  = sdl.KEY_Z
	KEY_1                  = sdl.KEY_1
	KEY_2                  = sdl.KEY_2
	KEY_3                  = sdl.KEY_3
	KEY_4                  = sdl.KEY_4
	KEY_5                  = sdl.KEY_5
	KEY_6                  = sdl.KEY_6
	KEY_7                  = sdl.KEY_7
	KEY_8                  = sdl.KEY_8
	KEY_9                  = sdl.KEY_9
	KEY_0                  = sdl.KEY_0
	KEY_RETURN             = sdl.KEY_RETURN
	KEY_ESCAPE             = sdl.KEY_ESCAPE
	KEY_BACKSPACE          = sdl.KEY_BACKSPACE
	KEY_TAB                = sdl.KEY_TAB
	KEY_SPACE              = sdl.KEY_SPACE
	KEY_MINUS              = sdl.KEY_MINUS
	KEY_EQUALS             = sdl.KEY_EQUALS
	KEY_LEFTBRACKET        = sdl.KEY_LEFTBRACKET
	KEY_RIGHTBRACKET       = sdl.KEY_RIGHTBRACKET
	KEY_BACKSLASH          = sdl.KEY_BACKSLASH
	KEY_NONUSHASH          = sdl.KEY_NONUSHASH
	KEY_SEMICOLON          = sdl.KEY_SEMICOLON
	KEY_APOSTROPHE         = sdl.KEY_APOSTROPHE
	KEY_GRAVE              = sdl.KEY_GRAVE
	KEY_COMMA              = sdl.KEY_COMMA
	KEY_PERIOD             = sdl.KEY_PERIOD
	KEY_SLASH              = sdl.KEY_SLASH
	KEY_CAPSLOCK           = sdl.KEY_CAPSLOCK
	KEY_F1                 = sdl.KEY_F1
	KEY_F2                 = sdl.KEY_F2
	KEY_F3                 = sdl.KEY_F3
	KEY_F4                 = sdl.KEY_F4
	KEY_F5                 = sdl.KEY_F5
	KEY_F6                 = sdl.KEY_F6
	KEY_F7                 = sdl.KEY_F7
	KEY_F8                 = sdl.KEY_F8
	KEY_F9                 = sdl.KEY_F9
	KEY_F10                = sdl.KEY_F10
	KEY_F11                = sdl.KEY_F11
	KEY_F12                = sdl.KEY_F12
	KEY_PRINTSCREEN        = sdl.KEY_PRINTSCREEN
	KEY_SCROLLLOCK         = sdl.KEY_SCROLLLOCK
	KEY_PAUSE              = sdl.KEY_PAUSE
	KEY_INSERT             = sdl.KEY_INSERT
	KEY_HOME               = sdl.KEY_HOME
	KEY_PAGEUP             = sdl.KEY_PAGEUP
	KEY_DELETE             = sdl.KEY_DELETE
	KEY_END                = sdl.KEY_END
	KEY_PAGEDOWN           = sdl.KEY_PAGEDOWN
	KEY_RIGHT              = sdl.KEY_RIGHT
	KEY_LEFT               = sdl.KEY_LEFT
	KEY_DOWN               = sdl.KEY_DOWN
	KEY_UP                 = sdl.KEY_UP
	KEY_NUMLOCKCLEAR       = sdl.KEY_NUMLOCKCLEAR
	KEY_KP_DIVIDE          = sdl.KEY_KP_DIVIDE
	KEY_KP_MULTIPLY        = sdl.KEY_KP_MULTIPLY
	KEY_KP_MINUS           = sdl.KEY_KP_MINUS
	KEY_KP_PLUS            = sdl.KEY_KP_PLUS
	KEY_KP_ENTER           = sdl.KEY_KP_ENTER
	KEY_KP_1               = sdl.KEY_KP_1
	KEY_KP_2               = sdl.KEY_KP_2
	KEY_KP_3               = sdl.KEY_KP_3
	KEY_KP_4               = sdl.KEY_KP_4
	KEY_KP_5               = sdl.KEY_KP_5
	KEY_KP_6               = sdl.KEY_KP_6
	KEY_KP_7               = sdl.KEY_KP_7
	KEY_KP_8               = sdl.KEY_KP_8
	KEY_KP_9               = sdl.KEY_KP_9
	KEY_KP_0               = sdl.KEY_KP_0
	KEY_KP_PERIOD          = sdl.KEY_KP_PERIOD
	KEY_NONUSBACKSLASH     = sdl.KEY_NONUSBACKSLASH
	KEY_APPLICATION        = sdl.KEY_APPLICATION
	KEY_POWER              = sdl.KEY_POWER
	KEY_KP_EQUALS          = sdl.KEY_KP_EQUALS
	KEY_F13                = sdl.KEY_F13
	KEY_F14                = sdl.KEY_F14
	KEY_F15                = sdl.KEY_F15
	KEY_F16                = sdl.KEY_F16
	KEY_F17                = sdl.KEY_F17
	KEY_F18                = sdl.KEY_F18
	KEY_F19                = sdl.KEY_F19
	KEY_F20                = sdl.KEY_F20
	KEY_F21                = sdl.KEY_F21
	KEY_F22                = sdl.KEY_F22
	KEY_F23                = sdl.KEY_F23
	KEY_F24                = sdl.KEY_F24
	KEY_EXECUTE            = sdl.KEY_EXECUTE
	KEY_HELP               = sdl.KEY_HELP
	KEY_MENU               = sdl.KEY_MENU
	KEY_SELECT             = sdl.KEY_SELECT
	KEY_STOP               = sdl.KEY_STOP
	KEY_AGAIN              = sdl.KEY_AGAIN
	KEY_UNDO               = sdl.KEY_UNDO
	KEY_CUT                = sdl.KEY_CUT
	KEY_COPY               = sdl.KEY_COPY
	KEY_PASTE              = sdl.KEY_PASTE
	KEY_FIND               = sdl.KEY_FIND
	KEY_MUTE               = sdl.KEY_MUTE
	KEY_VOLUMEUP           = sdl.KEY_VOLUMEUP
	KEY_VOLUMEDOWN         = sdl.KEY_VOLUMEDOWN
	KEY_KP_COMMA           = sdl.KEY_KP_COMMA
	KEY_KP_EQUALSAS400     = sdl.KEY_KP_EQUALSAS400
	KEY_INTERNATIONAL1     = sdl.KEY_INTERNATIONAL1
	KEY_INTERNATIONAL2     = sdl.KEY_INTERNATIONAL2
	KEY_INTERNATIONAL3     = sdl.KEY_INTERNATIONAL3
	KEY_INTERNATIONAL4     = sdl.KEY_INTERNATIONAL4
	KEY_INTERNATIONAL5     = sdl.KEY_INTERNATIONAL5
	KEY_INTERNATIONAL6     = sdl.KEY_INTERNATIONAL6
	KEY_INTERNATIONAL7     = sdl.KEY_INTERNATIONAL7
	KEY_INTERNATIONAL8     = sdl.KEY_INTERNATIONAL8
	KEY_INTERNATIONAL9     = sdl.KEY_INTERNATIONAL9
	KEY_LANG1              = sdl.KEY_LANG1
	KEY_LANG2              = sdl.KEY_LANG2
	KEY_LANG3              = sdl.KEY_LANG3
	KEY_LANG4              = sdl.KEY_LANG4
	KEY_LANG5              = sdl.KEY_LANG5
	KEY_LANG6              = sdl.KEY_LANG6
	KEY_LANG7              = sdl.KEY_LANG7
	KEY_LANG8              = sdl.KEY_LANG8
	KEY_LANG9              = sdl.KEY_LANG9
	KEY_ALTERASE           = sdl.KEY_ALTERASE
	KEY_SYSREQ             = sdl.KEY_SYSREQ
	KEY_CANCEL             = sdl.KEY_CANCEL
	KEY_CLEAR              = sdl.KEY_CLEAR
	KEY_PRIOR              = sdl.KEY_PRIOR
	KEY_RETURN2            = sdl.KEY_RETURN2
	KEY_SEPARATOR          = sdl.KEY_SEPARATOR
	KEY_OUT                = sdl.KEY_OUT
	KEY_OPER               = sdl.KEY_OPER
	KEY_CLEARAGAIN         = sdl.KEY_CLEARAGAIN
	KEY_CRSEL              = sdl.KEY_CRSEL
	KEY_EXSEL              = sdl.KEY_EXSEL
	KEY_KP_00              = sdl.KEY_KP_00
	KEY_KP_000             = sdl.KEY_KP_000
	KEY_THOUSANDSSEPARATOR = sdl.KEY_THOUSANDSSEPARATOR
	KEY_DECIMALSEPARATOR   = sdl.KEY_DECIMALSEPARATOR
	KEY_CURRENCYUNIT       = sdl.KEY_CURRENCYUNIT
	KEY_CURRENCYSUBUNIT    = sdl.KEY_CURRENCYSUBUNIT
	KEY_KP_LEFTPAREN       = sdl.KEY_KP_LEFTPAREN
	KEY_KP_RIGHTPAREN      = sdl.KEY_KP_RIGHTPAREN
	KEY_KP_LEFTBRACE       = sdl.KEY_KP_LEFTBRACE
	KEY_KP_RIGHTBRACE      = sdl.KEY_KP_RIGHTBRACE
	KEY_KP_TAB             = sdl.KEY_KP_TAB
	KEY_KP_BACKSPACE       = sdl.KEY_KP_BACKSPACE
	KEY_KP_A               = sdl.KEY_KP_A
	KEY_KP_B               = sdl.KEY_KP_B
	KEY_KP_C               = sdl.KEY_KP_C
	KEY_KP_D               = sdl.KEY_KP_D
	KEY_KP_E               = sdl.KEY_KP_E
	KEY_KP_F               = sdl.KEY_KP_F
	KEY_KP_XOR             = sdl.KEY_KP_XOR
	KEY_KP_POWER           = sdl.KEY_KP_POWER
	KEY_KP_PERCENT         = sdl.KEY_KP_PERCENT
	KEY_KP_LESS            = sdl.KEY_KP_LESS
	KEY_KP_GREATER         = sdl.KEY_KP_GREATER
	KEY_KP_AMPERSAND       = sdl.KEY_KP_AMPERSAND
	KEY_KP_DBLAMPERSAND    = sdl.KEY_KP_DBLAMPERSAND
	KEY_KP_VERTICALBAR     = sdl.KEY_KP_VERTICALBAR
	KEY_KP_DBLVERTICALBAR  = sdl.KEY_KP_DBLVERTICALBAR
	KEY_KP_COLON           = sdl.KEY_KP_COLON
	KEY_KP_HASH            = sdl.KEY_KP_HASH
	KEY_KP_SPACE           = sdl.KEY_KP_SPACE
	KEY_KP_AT              = sdl.KEY_KP_AT
	KEY_KP_EXCLAM          = sdl.KEY_KP_EXCLAM
	KEY_KP_MEMSTORE        = sdl.KEY_KP_MEMSTORE
	KEY_KP_MEMRECALL       = sdl.KEY_KP_MEMRECALL
	KEY_KP_MEMCLEAR        = sdl.KEY_KP_MEMCLEAR
	KEY_KP_MEMADD          = sdl.KEY_KP_MEMADD
	KEY_KP_MEMSUBTRACT     = sdl.KEY_KP_MEMSUBTRACT
	KEY_KP_MEMMULTIPLY     = sdl.KEY_KP_MEMMULTIPLY
	KEY_KP_MEMDIVIDE       = sdl.KEY_KP_MEMDIVIDE
	KEY_KP_PLUSMINUS       = sdl.KEY_KP_PLUSMINUS
	KEY_KP_CLEAR           = sdl.KEY_KP_CLEAR
	KEY_KP_CLEARENTRY      = sdl.KEY_KP_CLEARENTRY
	KEY_KP_BINARY          = sdl.KEY_KP_BINARY
	KEY_KP_OCTAL           = sdl.KEY_KP_OCTAL
	KEY_KP_DECIMAL         = sdl.KEY_KP_DECIMAL
	KEY_KP_HEXADECIMAL     = sdl.KEY_KP_HEXADECIMAL
	KEY_LCTRL              = sdl.KEY_LCTRL
	KEY_LSHIFT             = sdl.KEY_LSHIFT
	KEY_LALT               = sdl.KEY_LALT
	KEY_LGUI               = sdl.KEY_LGUI
	KEY_RCTRL              = sdl.KEY_RCTRL
	KEY_RSHIFT             = sdl.KEY_RSHIFT
	KEY_RALT               = sdl.KEY_RALT
	KEY_RGUI               = sdl.KEY_RGUI
	KEY_MODE               = sdl.KEY_MODE
	KEY_AUDIONEXT          = sdl.KEY_AUDIONEXT
	KEY_AUDIOPREV          = sdl.KEY_AUDIOPREV
	KEY_AUDIOSTOP          = sdl.KEY_AUDIOSTOP
	KEY_AUDIOPLAY          = sdl.KEY_AUDIOPLAY
	KEY_AUDIOMUTE          = sdl.KEY_AUDIOMUTE
	KEY_MEDIASELECT        = sdl.KEY_MEDIASELECT
	KEY_WWW                = sdl.KEY_WWW
	KEY_MAIL               = sdl.KEY_MAIL
	KEY_CALCULATOR         = sdl.KEY_CALCULATOR
	KEY_COMPUTER           = sdl.KEY_COMPUTER
	KEY_AC_SEARCH          = sdl.KEY_AC_SEARCH
	KEY_AC_HOME            = sdl.KEY_AC_HOME
	KEY_AC_BACK            = sdl.KEY_AC_BACK
	KEY_AC_FORWARD         = sdl.KEY_AC_FORWARD
	KEY_AC_STOP            = sdl.KEY_AC_STOP
	KEY_AC_REFRESH         = sdl.KEY_AC_REFRESH
	KEY_AC_BOOKMARKS       = sdl.KEY_AC_BOOKMARKS
	KEY_BRIGHTNESSDOWN     = sdl.KEY_BRIGHTNESSDOWN
	KEY_BRIGHTNESSUP       = sdl.KEY_BRIGHTNESSUP
	KEY_DISPLAYSWITCH      = sdl.KEY_DISPLAYSWITCH
	KEY_KBDILLUMTOGGLE     = sdl.KEY_KBDILLUMTOGGLE
	KEY_KBDILLUMDOWN       = sdl.KEY_KBDILLUMDOWN
	KEY_KBDILLUMUP         = sdl.KEY_KBDILLUMUP
	KEY_EJECT              = sdl.KEY_EJECT
	KEY_SLEEP              = sdl.KEY_SLEEP
)
