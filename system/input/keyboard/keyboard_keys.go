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
	KeyComma        Key = 44  // 0x2C (',')
	KeyPeriod       Key = 46  // 0x2E ('.')
	KeySlash        Key = 47  // 0x2F ('/')
	KeySemicolon    Key = 59  // 0x3B (';')
	KeyQuote        Key = 39  // 0x27 ('\'')
	KeyLeftbracket  Key = 91  // 0x5B ('[')
	KeyRightbracket Key = 93  // 0x5D (']')
	KeyBackslash    Key = 92  // 0x5C ('\\')
	KeyDelete       Key = 127 // 0x7F ('\177')
)

// -------------------------------------------------------------------

// const KeyDoubleQuote = 34 // 0x22 ('"')
//
// const KeyExclaim = 33    // 0x21 ('!')
// const KeyAt = 64         // 0x40 ('@')
// const KeyHash = 35       // 0x23 ('#')
// const KeyDollar = 36     // 0x24 ('$')
// const KeyPercent = 37    // 0x25 ('%')
// const KeyCaret = 94      // 0x5E ('^')
// const KeyAmpersand = 38  // 0x26 ('&')
// const KeyAsterisk = 42   // 0x2A ('*')
// const KeyLeftParen = 40  // 0x28 ('(')
// const KeyRightParen = 41 // 0x29 (')')
// const KeyUnderscore = 95 // 0x5F ('_')
// const KeyPlus = 43       // 0x2B ('+')
// const KeyColon = 58      // 0x3A (':')
// const KeyLess = 60       // 0x3C ('<')
// const KeyGreater = 62    // 0x3E ('>')
// const KeyQuestion = 63   // 0x3F ('?')
//
// const KeyPrintscreen = 1073741894 // 0x40000046
// const KeyScrolllock = 1073741895  // 0x40000047
// const KeyPause = 1073741896       // 0x40000048
// const KeyInsert = 1073741897      // 0x40000049
// const KeyHome = 1073741898        // 0x4000004A
// const KeyPageup = 1073741899      // 0x4000004B
// const KeyEnd = 1073741901         // 0x4000004D
// const KeyPagedown = 1073741902    // 0x4000004E
//
// const KeyNumlockclear = 1073741907       // 0x40000053
// const KeyKp_divide = 1073741908          // 0x40000054
// const KeyKp_multiply = 1073741909        // 0x40000055
// const KeyKp_minus = 1073741910           // 0x40000056
// const KeyKp_plus = 1073741911            // 0x40000057
// const KeyKp_enter = 1073741912           // 0x40000058
// const KeyKp_1 = 1073741913               // 0x40000059
// const KeyKp_2 = 1073741914               // 0x4000005A
// const KeyKp_3 = 1073741915               // 0x4000005B
// const KeyKp_4 = 1073741916               // 0x4000005C
// const KeyKp_5 = 1073741917               // 0x4000005D
// const KeyKp_6 = 1073741918               // 0x4000005E
// const KeyKp_7 = 1073741919               // 0x4000005F
// const KeyKp_8 = 1073741920               // 0x40000060
// const KeyKp_9 = 1073741921               // 0x40000061
// const KeyKp_0 = 1073741922               // 0x40000062
// const KeyKp_period = 1073741923          // 0x40000063
// const KeyApplication = 1073741925        // 0x40000065
// const KeyPower = 1073741926              // 0x40000066
// const KeyKp_equals = 1073741927          // 0x40000067
// const KeyF13 = 1073741928                // 0x40000068
// const KeyF14 = 1073741929                // 0x40000069
// const KeyF15 = 1073741930                // 0x4000006A
// const KeyF16 = 1073741931                // 0x4000006B
// const KeyF17 = 1073741932                // 0x4000006C
// const KeyF18 = 1073741933                // 0x4000006D
// const KeyF19 = 1073741934                // 0x4000006E
// const KeyF20 = 1073741935                // 0x4000006F
// const KeyF21 = 1073741936                // 0x40000070
// const KeyF22 = 1073741937                // 0x40000071
// const KeyF23 = 1073741938                // 0x40000072
// const KeyF24 = 1073741939                // 0x40000073
// const KeyExecute = 1073741940            // 0x40000074
// const KeyHelp = 1073741941               // 0x40000075
// const KeyMenu = 1073741942               // 0x40000076
// const KeySelect = 1073741943             // 0x40000077
// const KeyStop = 1073741944               // 0x40000078
// const KeyAgain = 1073741945              // 0x40000079
// const KeyUndo = 1073741946               // 0x4000007A
// const KeyCut = 1073741947                // 0x4000007B
// const KeyCopy = 1073741948               // 0x4000007C
// const KeyPaste = 1073741949              // 0x4000007D
// const KeyFind = 1073741950               // 0x4000007E
// const KeyMute = 1073741951               // 0x4000007F
// const KeyVolumeup = 1073741952           // 0x40000080
// const KeyVolumedown = 1073741953         // 0x40000081
// const KeyKp_comma = 1073741957           // 0x40000085
// const KeyKp_equalsas400 = 1073741958     // 0x40000086
// const KeyAlterase = 1073741977           // 0x40000099
// const KeySysreq = 1073741978             // 0x4000009A
// const KeyCancel = 1073741979             // 0x4000009B
// const KeyClear = 1073741980              // 0x4000009C
// const KeyPrior = 1073741981              // 0x4000009D
// const KeyReturn2 = 1073741982            // 0x4000009E
// const KeySeparator = 1073741983          // 0x4000009F
// const KeyOut = 1073741984                // 0x400000A0
// const KeyOper = 1073741985               // 0x400000A1
// const KeyClearagain = 1073741986         // 0x400000A2
// const KeyCrsel = 1073741987              // 0x400000A3
// const KeyExsel = 1073741988              // 0x400000A4
// const KeyKp_00 = 1073742000              // 0x400000B0
// const KeyKp_000 = 1073742001             // 0x400000B1
// const KeyThousandsseparator = 1073742002 // 0x400000B2
// const KeyDecimalseparator = 1073742003   // 0x400000B3
// const KeyCurrencyunit = 1073742004       // 0x400000B4
// const KeyCurrencysubunit = 1073742005    // 0x400000B5
// const KeyKp_leftparen = 1073742006       // 0x400000B6
// const KeyKp_rightparen = 1073742007      // 0x400000B7
// const KeyKp_leftbrace = 1073742008       // 0x400000B8
// const KeyKp_rightbrace = 1073742009      // 0x400000B9
// const KeyKp_tab = 1073742010             // 0x400000BA
// const KeyKp_backspace = 1073742011       // 0x400000BB
// const KeyKp_a = 1073742012               // 0x400000BC
// const KeyKp_b = 1073742013               // 0x400000BD
// const KeyKp_c = 1073742014               // 0x400000BE
// const KeyKp_d = 1073742015               // 0x400000BF
// const KeyKp_e = 1073742016               // 0x400000C0
// const KeyKp_f = 1073742017               // 0x400000C1
// const KeyKp_xor = 1073742018             // 0x400000C2
// const KeyKp_power = 1073742019           // 0x400000C3
// const KeyKp_percent = 1073742020         // 0x400000C4
// const KeyKp_less = 1073742021            // 0x400000C5
// const KeyKp_greater = 1073742022         // 0x400000C6
// const KeyKp_ampersand = 1073742023       // 0x400000C7
// const KeyKp_dblampersand = 1073742024    // 0x400000C8
// const KeyKp_verticalbar = 1073742025     // 0x400000C9
// const KeyKp_dblverticalbar = 1073742026  // 0x400000CA
// const KeyKp_colon = 1073742027           // 0x400000CB
// const KeyKp_hash = 1073742028            // 0x400000CC
// const KeyKp_space = 1073742029           // 0x400000CD
// const KeyKp_at = 1073742030              // 0x400000CE
// const KeyKp_exclam = 1073742031          // 0x400000CF
// const KeyKp_memstore = 1073742032        // 0x400000D0
// const KeyKp_memrecall = 1073742033       // 0x400000D1
// const KeyKp_memclear = 1073742034        // 0x400000D2
// const KeyKp_memadd = 1073742035          // 0x400000D3
// const KeyKp_memsubtract = 1073742036     // 0x400000D4
// const KeyKp_memmultiply = 1073742037     // 0x400000D5
// const KeyKp_memdivide = 1073742038       // 0x400000D6
// const KeyKp_plusminus = 1073742039       // 0x400000D7
// const KeyKp_clear = 1073742040           // 0x400000D8
// const KeyKp_clearentry = 1073742041      // 0x400000D9
// const KeyKp_binary = 1073742042          // 0x400000DA
// const KeyKp_octal = 1073742043           // 0x400000DB
// const KeyKp_decimal = 1073742044         // 0x400000DC
// const KeyKp_hexadecimal = 1073742045     // 0x400000DD
// const KeyLgui = 1073742051           // 0x400000E3
// const KeyRgui = 1073742055           // 0x400000E7
// const KeyMode = 1073742081           // 0x40000101
// const KeyAudionext = 1073742082      // 0x40000102
// const KeyAudioprev = 1073742083      // 0x40000103
// const KeyAudiostop = 1073742084      // 0x40000104
// const KeyAudioplay = 1073742085      // 0x40000105
// const KeyAudiomute = 1073742086      // 0x40000106
// const KeyMediaselect = 1073742087    // 0x40000107
// const KeyWww = 1073742088            // 0x40000108
// const KeyMail = 1073742089           // 0x40000109
// const KeyCalculator = 1073742090     // 0x4000010A
// const KeyComputer = 1073742091       // 0x4000010B
// const KeyAc_search = 1073742092      // 0x4000010C
// const KeyAc_home = 1073742093        // 0x4000010D
// const KeyAc_back = 1073742094        // 0x4000010E
// const KeyAc_forward = 1073742095     // 0x4000010F
// const KeyAc_stop = 1073742096        // 0x40000110
// const KeyAc_refresh = 1073742097     // 0x40000111
// const KeyAc_bookmarks = 1073742098   // 0x40000112
// const KeyBrightnessdown = 1073742099 // 0x40000113
// const KeyBrightnessup = 1073742100   // 0x40000114
// const KeyDisplayswitch = 1073742101  // 0x40000115
// const KeyKbdillumtoggle = 1073742102 // 0x40000116
// const KeyKbdillumdown = 1073742103   // 0x40000117
// const KeyKbdillumup = 1073742104     // 0x40000118
// const KeyEject = 1073742105          // 0x40000119
// const KeySleep = 1073742106          // 0x4000011A
//
