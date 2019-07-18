grammar FmtParser;

expression
    : text* ((simp_exp
                 | bool_exp
                 | cmp_exp
                 | text) text*)*
    ;
yes_stmt
    : text* ((simp_exp
                 | bool_exp
                 | cmp_exp
                 | text) text*)*
    ;
no_stmt
    : text* ((simp_exp
                 | bool_exp
                 | cmp_exp
                 | text) text*)*
    ;
cmp_exp_body
    : param eq value QM yes_stmt (BAR no_stmt)?
    ;
cmp_exp
    : EXP_OPEN cmp_exp_body EXP_CLOSE
    ;
pattern_exp_body
    :param PATTERN value QM yes_stmt (BAR no_stmt)?
    ;
bool_exp_body
    :param QM yes_stmt (BAR no_stmt)?
    ;
bool_exp
    : EXP_OPEN (bool_exp_body|pattern_exp_body) EXP_CLOSE
    ;
simp_exp
    : EXP_OPEN param EXP_CLOSE
    ;

text
    : TSTR | STRING | QSTRING
    ;
param
    : TSTR | STRING | QSTRING
    ;
value:QSTRING;
eq   :DEQ;

EXP_OPEN    : '%(' ;
EXP_CLOSE: ')';
QM : '?';
DEQ: '==' | '!=';
PATTERN:'~';
BAR: '|';
AT: '@';
TSTR:('<color red>' | '<color green>') (~[%<]+ | '%' ~[(]+) | '</color>' | STRING TSTR;
QSTRING: '\'' STRING* '\'' | '/' ~[/]* '/' ;
STRING :  SEQ | ESCAPE;
SEQ: ~[%)?!|='\\~]+ | '%' ~'('*? | '%)';
ESCAPE:'\\|'|'\\' | '\\?' | '\\)' | '\\~';
