package ascii

const (
	Fedora = `
             .',;::::;,'.
         .';:cccccccccccc:;,.
      .;cccccccccccccccccccccc;.
    .:cccccccccccccccccccccccccc:.
  .;ccccccccccccc;.:dddl:.;ccccccc;.
 .:ccccccccccccc;OWMKOOXMWd;ccccccc:.
.:ccccccccccccc;KMMc;cc;xMMc;ccccccc:.
,cccccccccccccc;MMM.;cc;;WW:;cccccccc,
:cccccccccccccc;MMM.;cccccccccccccccc:
:ccccccc;oxOOOo;MMM000k.;cccccccccccc:
cccccc;0MMKxdd:;MMMkddc.;cccccccccccc;
ccccc;XMO';cccc;MMM.;cccccccccccccccc'
ccccc;MMo;ccccc;MMW.;ccccccccccccccc;
ccccc;0MNc.ccc.xMMd;ccccccccccccccc;
cccccc;dNMWXXXWM0:;cccccccccccccc:,
cccccccc;.:odl:.;cccccccccccccc:,.
ccccccccccccccccccccccccccccc:'.
:ccccccccccccccccccccccc:;,..
 ':cccccccccccccccc::;,.
`
	Ubuntu = `
                             ....
              .',:clooo:  .:looooo:.
           .;looooooooc  .oooooooooo'
        .;looooool:,''.  :ooooooooooc
       ;looool;.         'oooooooooo,
      ;clool'             .cooooooc.  ,,
         ...                ......  .:oo,
  .;clol:,.                        .loooo'
 :ooooooooo,                        'ooool
'ooooooooooo.                        loooo.
'ooooooooool                         coooo.
 ,loooooooc.                        .loooo.
   .,;;;'.                          ;ooooc
       ...                         ,ooool.
    .cooooc.              ..',,'.  .cooo.
      ;ooooo:.           ;oooooooc.  :l.
       .coooooc,..      coooooooooo.
         .:ooooooolc:. .ooooooooooo'
           .':loooooo;  ,oooooooooc
               ..';::c'  .;loooo:'
`
	Debian = `
        _,met$$$$$$$$$$gg.
     ,g$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$P.
   ,g$$$$P""       """Y$$$$.".
  ,$$$$P'              '$$$$$$.
',$$$$P       ,ggs.     '$$$$b:
'd$$$$'     ,$P"'   .    $$$$$$
 $$$$P      d$'     ,    $$$$P
 $$$$:      $$$.   -    ,d$$$$'
 $$$$;      Y$b._   _,d$P'
 Y$$$$.    '.  "Y$$$$$$$$P"'
 '$$$$b      "-.__
  'Y$$$$b
   'Y$$$$.
     '$$$$b.
       'Y$$$$b.
         '"Y$$b._
             '""""
`
	Arch = `
                  -'
                 .o+'
                'ooo/
               '+oooo:
              '+oooooo:
              -+oooooo+:
            '/:-:++oooo+:
           '/++++/+++++++:
          '/++++++++++++++:
         '/+++ooooooooooooo/'
        ./ooosssso++osssssso+'
       .oossssso-''''/ossssss+'
      -osssssso.      :ssssssso.
     :osssssss/        osssso+++.
    /ossssssss/        +ssssooo/-
  '/ossssso+/:-        -:/+osssso+-
 '+sso+:-'                 '.-/+oso:
'++:.                           '-/+/
.'                                 '/
`
	Alpine = `
       .hddddddddddddddddddddddh.
      :dddddddddddddddddddddddddd:
     /dddddddddddddddddddddddddddd/
    +dddddddddddddddddddddddddddddd+
  'sdddddddddddddddddddddddddddddddds'
 'ydddddddddddd++hdddddddddddddddddddy'
.hddddddddddd+'  '+ddddh:-sdddddddddddh.
hdddddddddd+'      '+y:    .sddddddddddh
ddddddddh+'   '//'   '.'     -sddddddddd
ddddddh+'   '/hddh/'   ':s-    -sddddddd
ddddh+'   '/+/dddddh/'   '+s-    -sddddd
ddd+'   '/o' :dddddddh/'   'oy-    .yddd
hdddyo+ohddyosdddddddddho+oydddy++ohdddh
.hddddddddddddddddddddddddddddddddddddh.
 'yddddddddddddddddddddddddddddddddddy'
  'sdddddddddddddddddddddddddddddddds'
    +dddddddddddddddddddddddddddddd+
     /dddddddddddddddddddddddddddd/
      :dddddddddddddddddddddddddd:
       .hddddddddddddddddddddddh.
`
	MacOS = `
                     ..'
                 ,xNMM.
               .OMMMMo
               lMM"
     .;loddo:.  .olloddol;.
   cKMMMMMMMMMMNWMMMMMMMMMM0:
 .KMMMMMMMMMMMMMMMMMMMMMMMWd.
 XMMMMMMMMMMMMMMMMMMMMMMMX.
;MMMMMMMMMMMMMMMMMMMMMMMM:
:MMMMMMMMMMMMMMMMMMMMMMMM:
.MMMMMMMMMMMMMMMMMMMMMMMMX.
 kMMMMMMMMMMMMMMMMMMMMMMMMWd.
 'XMMMMMMMMMMMMMMMMMMMMMMMMMMk
  'XMMMMMMMMMMMMMMMMMMMMMMMMK.
    kMMMMMMMMMMMMMMMMMMMMMMd
     ;KMMMMMMMWXXWMMMMMMMk.
       "cooc*"    "*coo'"
`
	Unknown = `
       ________
   _jgN########Ngg_
 _N##N@@""  ""9NN##Np_
d###P            N####p
"^^"              T####
                  d###P
               _g###@F
            _gN##@P
          gN###F"
         d###F
        0###F
        0###F
        0###F
        "NN@'

         ___
        q###r
         ""
`
)

var logoMap = map[string]string{
	"fedora": Fedora,
	"ubuntu": Ubuntu,
	"debian": Debian,
	"arch":   Arch,
	"alpine": Alpine,
	"macos":  MacOS,
}

func Get(osID string) string {
	if l, ok := logoMap[osID]; ok {
		return l
	}
	return Ubuntu
}
