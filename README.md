Learning how to write parsers by implementing an expression evaluator.
Uses recursive-descent parsing technique.
Supports variables.

Grammar:
```
Statement :=
    | ID "=" Expression
    | Expression
    ;

Expression :=
    | Term ( ( "+" | "-" ) Term )*
    ;

Term := 
    | Primary ( ( "*" | "/" ) Primary )*;
    ;

Primary :=
    | Factor ( "^" Primary)*
    ;

Factor :=
    | NUMBER
    | ID
    | "-" Factor
    | "|" Expression "|"
    | "(" Expression ")"
    ;
```

TODO:
- Add more operators ( % ! )
- Add global constants
