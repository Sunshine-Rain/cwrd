# cwrd
HAM CW Random Text Generator.
This tool is created for the HAM amateurs to create random characters.
The output text can be played by CWPlayer.

# Usage
1. clone the project
2. type in terminal
```
cd cwrd
go build ./cmd/cwrd
```
3. run the ```cwrd``` exec with flag.

# --help

```
Usage: HAM CW Random Text Generator <command> [flags]

Used for HAM amateurs to create random characters

Flags:
  -h, --help            Show context-sensitive help.
  -s, --start="a"       Start char in alphabet, default is a ($HAM CW RANDOM TEXT GENERATOR_START).
  -e, --end="z"         End char in alphabet, default is z ($HAM CW RANDOM TEXT GENERATOR_END).
  -c, --charCount=5     Chars count in each word, default is 5 ($HAM CW RANDOM TEXT GENERATOR_CHAR_COUNT).
  -w, --wordCount=10    Words count in each word, default is 10 ($HAM CW RANDOM TEXT GENERATOR_WORD_COUNT).
  -l, --lines=10        Total lines, default is 10 ($HAM CW RANDOM TEXT GENERATOR_LINES).
      --charSet=""      Input character set, if not empty only use character in charset, default is empty ($HAM CW RANDOM TEXT GENERATOR_CHAR_SET).
  -p, --prefix=""       Output prefix, default is empty ($HAM CW RANDOM TEXT GENERATOR_PREFIX).
  -x, --suffix=""       Output suffix, default is empty ($HAM CW RANDOM TEXT GENERATOR_SUFFIX).
  -v, --version         Print version information and quit ($HAM CW RANDOM TEXT GENERATOR_VERSION)

Commands:
  rand    Used to create.

Run "HAM CW Random Text Generator <command> --help" for more information on a command.
```

# Examples

```
❯ ./cwrd 
uchcn rkwin usath tbmhz etanv osdlm xznhg iutge hdzwk hwfaz 
fyyya yakeb ikkcv kfuhc fubpp dwona wtadt lbqof avyzb zufzg 
kxkzk tgufw gqqwx wwdsh urovk kmfty pftvi gxyrz nwrtq ynhrh 
rcast fshiz eakut esxfn ouyuh cligv ugjeb qwqoy mvgwi xmiok 
lfwiv ebwxk oqwye qujvf gjoeh xvdex oihkj xuyhe lhrgi kktcy 
onknc ofexa qvndc oaqho doyyl qdmtg utmrw maxam rakjy uykem 
kgiqg xneky bsuca pqezw frvza pzbun jbwdm qxwsi gxxbw dzqqf 
nrohz ccgja qryyc vhbnq catnp zbgvg vfkmw ydbyv qzdnd ytjjf 
tobbc wsvdo jgary wyfca kcfok lvtao jgpql agozm dfscr katef 
zdstp meebt kxnbt niglp eopcx ohbbv vsfsw behgz tvkvv yiqjw
```

```
❯ ./cwrd -s x -e y -l 3
xxyxx xxyxx yyxxx yyxyy yyxyx yyyxx xyyyx yyxxx xyxyy yyyxy 
xyyyy yyyyy xxyyx yyxyx xyyxx xyxxy yyyyx xxxyy xyyxx yyyyy 
xxyxx yyyyy xxyyx yyxxx xyxyy xxyyx yyxyx yxxxy yxyxx xyxxy 
```

```
❯ ./cwrd -l 3 --charSet="axym"
xmmyx aaaax yaxym ymxyx myxmx yyxya aaaam yxymm mmxya mymyy 
amaya mayax yyyay xaxya mmyxx xamyy xayay mmaym yyyaa ayxam 
xayaa xyamx yxxxa xxmyy mamyx mxyxa yaymm mxymm xyymx xxyax 
```
