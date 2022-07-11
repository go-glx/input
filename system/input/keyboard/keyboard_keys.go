package keyboard

type Key int32

const KeyUnknown Key = 0 // 0x00 ('\0')

// F1-12 block
const (
	KeyF1  Key = 1073741882 // 0x4000003A
	KeyF2  Key = 1073741883 // 0x4000003B
	KeyF3  Key = 1073741884 // 0x4000003C
	KeyF4  Key = 1073741885 // 0x4000003D
	KeyF5  Key = 1073741886 // 0x4000003E
	KeyF6  Key = 1073741887 // 0x4000003F
	KeyF7  Key = 1073741888 // 0x40000040
	KeyF8  Key = 1073741889 // 0x40000041
	KeyF9  Key = 1073741890 // 0x40000042
	KeyF10 Key = 1073741891 // 0x40000043
	KeyF11 Key = 1073741892 // 0x40000044
	KeyF12 Key = 1073741893 // 0x40000045
)

// numeric block (0-9, -, =)
const (
	Key0      Key = 48 // 0x30 ('0')
	Key1      Key = 49 // 0x31 ('1')
	Key2      Key = 50 // 0x32 ('2')
	Key3      Key = 51 // 0x33 ('3')
	Key4      Key = 52 // 0x34 ('4')
	Key5      Key = 53 // 0x35 ('5')
	Key6      Key = 54 // 0x36 ('6')
	Key7      Key = 55 // 0x37 ('7')
	Key8      Key = 56 // 0x38 ('8')
	Key9      Key = 57 // 0x39 ('9')
	KeyMinus  Key = 45 // 0x2D ('-')
	KeyEquals Key = 61 // 0x3D ('=')
)

// chars block (a-z)
const (
	KeyA Key = 97  // 0x61 ('a')
	KeyB Key = 98  // 0x62 ('b')
	KeyC Key = 99  // 0x63 ('c')
	KeyD Key = 100 // 0x64 ('d')
	KeyE Key = 101 // 0x65 ('e')
	KeyF Key = 102 // 0x66 ('f')
	KeyG Key = 103 // 0x67 ('g')
	KeyH Key = 104 // 0x68 ('h')
	KeyI Key = 105 // 0x69 ('i')
	KeyJ Key = 106 // 0x6A ('j')
	KeyK Key = 107 // 0x6B ('k')
	KeyL Key = 108 // 0x6C ('l')
	KeyM Key = 109 // 0x6D ('m')
	KeyN Key = 110 // 0x6E ('n')
	KeyO Key = 111 // 0x6F ('o')
	KeyP Key = 112 // 0x70 ('p')
	KeyQ Key = 113 // 0x71 ('q')
	KeyR Key = 114 // 0x72 ('r')
	KeyS Key = 115 // 0x73 ('s')
	KeyT Key = 116 // 0x74 ('t')
	KeyU Key = 117 // 0x75 ('u')
	KeyV Key = 118 // 0x76 ('v')
	KeyW Key = 119 // 0x77 ('w')
	KeyX Key = 120 // 0x78 ('x')
	KeyY Key = 121 // 0x79 ('y')
	KeyZ Key = 122 // 0x7A ('z')
)

// arrows block
const (
	KeyRight Key = 1073741903 // 0x4000004F
	KeyLeft  Key = 1073741904 // 0x40000050
	KeyDown  Key = 1073741905 // 0x40000051
	KeyUp    Key = 1073741906 // 0x40000052
)

// special keys block
const (
	KeyTilda     Key = 96         // 0x60 ('`')
	KeyBackspace Key = 8          // 0x08 ('\b')
	KeyTab       Key = 9          // 0x09 ('\t')
	KeyCapslock  Key = 1073741881 // 0x40000039
	KeyReturn    Key = 13         // 0x0D ('\r')
	KeyEscape    Key = 27         // 0x1B ('\033')
	KeySpace     Key = 32         // 0x20 (' ')
	KeyLCtrl     Key = 1073742048 // 0x400000E0
	KeyRCtrl     Key = 1073742052 // 0x400000E4
	KeyLShift    Key = 1073742049 // 0x400000E1
	KeyRShift    Key = 1073742053 // 0x400000E5
	KeyLAlt      Key = 1073742050 // 0x400000E2
	KeyRAlt      Key = 1073742054 // 0x400000E6
)

// additional keys
const (
	KeyComma        Key = 44         // 0x2C (',')
	KeyPeriod       Key = 46         // 0x2E ('.')
	KeySlash        Key = 47         // 0x2F ('/')
	KeySemicolon    Key = 59         // 0x3B (';')
	KeyQuote        Key = 39         // 0x27 ('\'')
	KeyLeftBracket  Key = 91         // 0x5B ('[')
	KeyRightBracket Key = 93         // 0x5D (']')
	KeyBackslash    Key = 92         // 0x5C ('\\')
	KeyDelete       Key = 127        // 0x7F ('\177')
	KeyInsert       Key = 1073741897 // 0x40000049
	KeyPrintScreen  Key = 1073741894 // 0x40000046
	KeyScrollLock   Key = 1073741895 // 0x40000047
	KeyPause        Key = 1073741896 // 0x40000048
	KeyHome         Key = 1073741898 // 0x4000004A
	KeyEnd          Key = 1073741901 // 0x4000004D
	KeyPageUp       Key = 1073741899 // 0x4000004B
	KeyPageDown     Key = 1073741902 // 0x4000004E
)

// numpad
const (
	KeyKpDivide   Key = 1073741908 // 0x40000054
	KeyKpMultiply Key = 1073741909 // 0x40000055
	KeyKpMinus    Key = 1073741910 // 0x40000056
	KeyKpPlus     Key = 1073741911 // 0x40000057
	KeyKpEnter    Key = 1073741912 // 0x40000058
	KeyKp1        Key = 1073741913 // 0x40000059
	KeyKp2        Key = 1073741914 // 0x4000005A
	KeyKp3        Key = 1073741915 // 0x4000005B
	KeyKp4        Key = 1073741916 // 0x4000005C
	KeyKp5        Key = 1073741917 // 0x4000005D
	KeyKp6        Key = 1073741918 // 0x4000005E
	KeyKp7        Key = 1073741919 // 0x4000005F
	KeyKp8        Key = 1073741920 // 0x40000060
	KeyKp9        Key = 1073741921 // 0x40000061
	KeyKp0        Key = 1073741922 // 0x40000062
)
